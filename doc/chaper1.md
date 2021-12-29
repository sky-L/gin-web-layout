### 
项目目录
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

## GIN && Cobra

一般的 WEB 服务，都会包含多个模块， API 接口、定时任务、消费 MQ 的常驻进程等等，在这种情况下，很显然直接使用 GIN 来启动就只能开启 API 模块，十分不便。

我们用 Cobra 来管理项目的启动，初始化一个 rootCmd 来管理项目中所有的模块 

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
		address := fmt.Sprintf("%v:%v", "0.0.0.0", 80)
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

这个时候启动一下, 可以看到需要传一个命令行参数

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

