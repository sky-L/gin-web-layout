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

一般的 WEB 服务，都会包含 API 接口，定时任务，消费 MQ 的常驻进程等等，在这种情况下，很显然直接使用

我们用 Cobra 来管理项目的启动， 



我