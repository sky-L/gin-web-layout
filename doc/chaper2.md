### 设计一个应用
现在回想一下一个典型的 web 应用有哪些模块？

* 路由
* 配置 
* 数据库操作
* Service
* Api

在一些高并发应用下还有 `repository`，处在 service 和 db 层， 给 service 提供一层缓存。

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

从目录结构可以看到要设计的模块， 从最里层开始设计即 `storage`。各个模块初始化放在 `application` 里，返回一个 App 对象:

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

### 配置

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
数据库使用 `gorm`， 





 
