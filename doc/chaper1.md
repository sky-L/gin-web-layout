### 快速启动一个 api
目录
```
├── api  \\ 服务入口
├── cmd   \\ 服务启动入口
├── config
├── doc
    └── chaper1.md
├── go.mod
├── internal \\ 业务逻辑
├── main.go
├── pkg   \\ 三方包初始化
└── router \\ Api 路由
```

### GIN && Cobra

一般 WEB 服务，都会包含多个模块: API 接口、定时任务、消费 MQ 常驻进程等等，在这种情况下，很显然直接使用 GIN 来启动就只能开启 API 模块，十分不便。

我们用 Cobra 来管理项目的启动，现在不用关心 Cobra 如何使用，现在要的是满足我们需求。

> 很多时候人会陷入到细节里，就无法宏观的把控全局设计。无论是做需求还是设计框架，都应该梳理整个流程，每个流程需要什么样的技术，而技术细节反而不是最需要关心的。互联网发展到今天，你遇到的问题一定有人遇到过
要把关注点放到你的设计上。 

初始化一个 rootCmd 来管理项目中所有的模块 

```
// main.go

func main() {
	cmd.Execute()
}


// cmd/root.go
var rootCmd = &cobra.Command{
	Use:   "提供WebApi服务",
	Short: "webApi",
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("[错误]启动失败:", err)
	}
}

// cmd/api.go
var httpServer *http.Server

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "apiCmd",
	Long:  `apiCmd 提供api接口服务`,
	Run: func(cmd *cobra.Command, args []string) {
		address := fmt.Sprintf("%v:%v", "0.0.0.0", 8080)
		engine := gin.New()
		httpServer = &http.Server{
			Addr:        address,
			Handler:     engine,
			IdleTimeout: time.Minute,
		}
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("启动失败", err)
		}
	},
}
```

这个时候启动一下, 可以看到需要传一个命令行参数：

```
➜  gin-web-layout git:(master) ✗ go run main.go                  
webApi

Usage:
  提供WebApi服务 [command]

Available Commands:
  api         apiCmd
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

```

使用 `go run main.go api` 就可以启动服务了

```
➜  gin-web-layout git:(master) ✗ go run main.go api              
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)
```

首先开始接入路由， 所见即所得，能快速的看到自己写的成果

```
router/router.go

func InitRouter(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
}

在 cmd/api.go 增加以下代码

engine := gin.New()
router.InitRouter(engine)
```
这样一个 hello world 就完成了，这个也是 gin 快速开始的内容。 启动后访问一下:

```
➜  gin-web-layout git:(master) ✗ curl http://0.0.0.0:8080                                                                                                                                     
"ok"%  
```

这个时候我们来完善一下启动模块的代码，加上平滑重启，设置 5 秒后退出

```
// cmd/api.go
// 只展示核心代码，完整代码可以在 github 查看

// 等待 5 秒 timeout = 5 *time.sencond
func OnServiceStop(timeout time.Duration) {
	quit := make(chan os.Signal)
	signal.Notify(quit, signals.Get()...)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := httpServer.Shutdown(ctx)
	if err != nil {
		log.Fatalf("service stop error:%v", err)
	}
	log.Println("service is stopping")
}
```
这样就支持了热启动，然后就编译启动，ctrl+c， 他为什么没阻塞 5 秒，直接就退出了？

因为 ctrl+c 的时候，会检查是否有待处理的请求，如没有就会直接退出。我们可以模拟一个耗时请求:
```
// router/router.go
    engine.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.JSON(200, "ok")
	})
```
重新启动后，再次 ctrl+c 会发现 5 秒后项目才退出完成。 

### 题外话，线上的服务是如何发布部署的？

一般线上服务都会用网关调度流量，当我们一个服务接受到 kill(重启脚本一般用 kill，杀掉 pid) 信号，就不再接受新请求。

这一点可以用我们刚配置的热重启验证一下，把 timeout 设置 10s， 伪造一个耗时 10s 的请求，启动后执行退出（用 ctrl+c 或者 kill， 本质都是发送一个信号), 然后再访问服务，
会得到
```
➜  gin-web-layout git:(master) ✗ curl http://0.0.0.0:8080
curl: (7) Failed to connect to 0.0.0.0 port 8080: Connection refused
```
网关和服务会有心跳监测，无法提供服务后，网关自动踢掉服务，不再发流量，待恢复后再重新发流量。但是实际部署部署是另有方案，因为心跳是有间隔的，这个间隔期间服务退出了，就会造成大量的 502

实际线上操作为，当一台服务要退出的时候，会先到网关摘流量，再执行平滑退出，启动新服务，到网关挂上流量。 网关一般用的是阿里的 slb，也有人用 kong，都是一样的套路。









