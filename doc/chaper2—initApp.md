## 设计一个应用
现在回想一下一个典型的 web 应用有哪些模块？

* 路由
* 配置 
* 数据库操作
* Service
* Api

在一些高并发应用下还有 `repository`，处在 service 和 db 层之间， 给 service 提供一层缓存。

新建 internal 目录，目录结构：

```php
internal git:(master) ✗ tree -L 1
├── application
├── middleware
├── models
├── repository
├── service
└── storage
```
调用链路为 API -> service -> repository -> ( DB：storage -> models)。

从目录结构可以看到要设计的模块， 从最里层开始设计即 `storage`，`models` 是数据库和 `struct` 的映射，看一下即可。说到这里，想起有人开发了一些好用的工具：[SQL2Struct](https://printlove.cn/tools/sql2gorm/) 很好用。

各个模块初始化放在 `application` 里，返回一个 `*App` 对象， 先看一下 app 里的整体样子，再各模块分别设计:

## 模块概览
-  file: application/app.go
```php
type App struct {
	config *config.Config
	*Repository
	*storage.Storage
	*Service
	*Api
}

func NewApp() *App {
	app := &App{
		config: config.InitConfig(),
	}
    
	app.Storage = storage.NewStorage(app.config.MySqlConfig)

	app.Repository = NewRepository(app.Storage)

	app.Service = NewService(app.Repository)

	app.Api = NewApi(app.Service)

	return app
}
```

## 配置

使用一个包 `github.com/spf13/viper`, 功能就是把配置文件内容直接映射为一个结构体。为什么用这个包，我想 star 数说明一切了。
- file: config/config.go
```php 
func InitConfig() *Config {
	config := &Config{}

	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	conf := "./config/conf.yml"
	file := helper.Must(os.Open(conf)).(*os.File)
	defer file.Close()
	helper.Must(nil, viper.ReadConfig(file))

	helper.Must(nil, viper.Unmarshal(config))
	helper.Must(nil, Validate(config))
	return config
}

```

## storage

新建 `./internal/storage/storage.go`，从这里开始配置数据库， Mysql, Redis，ES... , 本次只配置 Mysql

可以看到有个 `helper.Must` 这是属于包 `github.com/thinkeridea/go-extend/helper"`，配置有错误的时候，直接 `panic` 终止程序。
以配置数据库示例：
- file: config/config.go
```php
type DB struct {
	DataSourceName string `mapstructure:"data_source_name" validate:"required"`
	MaxOpenConns   int    `mapstructure:"max_open_conns" validate:"required,min=1"`
	MaxIdleConns   int    `mapstructure:"max_idle_conns" validate:"required,min=1,ltefield=MaxOpenConns"`
}
type MySqlConfig struct {
	Default DB `mapstructure:"default"`
}

type Config struct {
	MySqlConfig MySqlConfig `mapstructure:"mysql_config"'`
}
```

数据库一般会有多个，示例中定义了一个 `default` 数据库。 多数情况下，多库并非是为了读写分离，现在的实际生产环境，使用阿里云的时候，数据库连接就是读写的链接，会有 `rwlb` 标识。
即使是自建库也会使用三方的中间件来做，而不会在业务代码做读写分离。

### 配置数据库
数据库使用 `gorm`，增加如下配置，并配置最大链接数,最大等待链接数
- file: storage/default_db/mysql.go
```php
type Mysql struct {
	DB *gorm.DB
}
func NewMysql(db config.DB) *Mysql {
	client := helper.Must(gorm.Open(mysql.Open(db.DataSourceName))).(*gorm.DB)
	sqlDB := helper.Must(client.DB()).(*sql.DB)
	// 赋值为指针类型，修改生效
	sqlDB.SetMaxIdleConns(db.MaxIdleConns)
	sqlDB.SetMaxOpenConns(db.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Minute)
	client.Debug()
	s := &Mysql{
		DB: client,
	}
	return s
}
```

可以看到 `Mysql.DB` 为大写，这个是为了支持在 service 可直接读取数据库。

再配置 `storage` ，作为数据库的统一入。以后再扩展数据库，统一在这里配置，Redis,Es.. 等等
- file : storage/storage.go

```php
type Storage struct {
	Default *default_db.Mysql
}

func NewStorage(mysql config.MySqlConfig) *Storage {
	return &Storage{
		Default: default_db.NewMysql(mysql.Default),
	}
}
```

## repository
`repository` 为数据层, 以 `blog` demo 说明： 
- file: ./repository/blog/repo.go
```php
type RepositoryInterface interface {
	FindById(id int) (default_db.Blog, error)
}
type BlogRepo struct {
	DB *gorm.DB
}
func NewBlogRepo(storage *storage.Storage) RepositoryInterface {
	return &BlogRepo{storage.Default.DB}
}
```

可以看到，返回的是一个 `interface`， 这样的目的是在 `service` 层直接注入这个 interface， 在新增 repository 方法的时候，无需改动 service 的注入

## service 
还是以 blog demo 示例：把数据层注入到属性
- file : service/blog_service/blog

```php

type BlogService struct {
	repo *blog.BlogRepo
}

func NewBlogService(repo *blog.BlogRepo) *BlogService {
	return &BlogService{
		repo: repo,
	}
}
```
## api
最后就是 api 层了。 使用的 gin 框架， 那 api 的方法就都是 router 对应的 handleFunction
- file: api/blog_api/blog.go
```php
type BlogApi struct {
	s       *blog_service.BlogService
}

func NewBlogApi(s *blog_service.BlogService) *BlogApi {
	return &BlogApi{
		s: s,
	}
}
func (b *BlogApi) List(c *gin.Context) {
	req := protocol.BlogListReq{}
	var err error
	err = c.ShouldBind(&req)
	if err != nil {
		api.NewResponse(c, err, nil)
		return
	}

	data, err := b.s.List(req.Id)
	if err != nil {
		return
	}
	api.NewResponse(c, nil, data)
}
```

新增 api/helper.go , 增加 api 的统一输出，`api.NewResponse(c, err, nil)`
```php
const SuccessCode = 0
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
func NewResponse(c *gin.Context, err error, data interface{}) {
	res := &Response{Data: data}
	if err != nil {
		res.Code = -1
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	res.Code = SuccessCode
	res.Message = "操作成功"
	c.JSON(http.StatusOK, res)
}
```

一个 api 配套一个 protocol， 参数和结构的映射关系定义在 protocol 目录:

```php
package protocol

type BlogListReq struct {
	Id int `json:"id" form:"id"`
}

type BlogResp struct {
	Name string `json:"name"`
}
```

最后添加一个路由, 以 blogList 为例：
- file: router/router.go
```php
engine.GET("/blog/list", app.Api.Blog.List)
```

## 最后
这样一个 web 请求就完成了。后续根据业务需求，再对这个脚手架做增加。完成代码再 [Github](https://github.com/sky-L/gin-web-layout)

### feature
- 预警
- 中间件
- 日志
- 缓存
- ...






 
