package cmd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skylee/gin-web-layout/internal/application"
	"github.com/skylee/gin-web-layout/pkg/signals"
	"github.com/skylee/gin-web-layout/router"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var httpServer *http.Server

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "apiCmd",
	Long:  `apiCmd 提供api接口服务`,
	Run: func(cmd *cobra.Command, args []string) {
		address := fmt.Sprintf("%v:%v", "0.0.0.0", 8080)
		log.Printf("listening and serving HTTP on %s\n", address)

		app := application.NewApp()

		engine := gin.New()
		router.InitRouter(engine, app)

		go Run(engine)

		OnServiceStop(5 * time.Second)
	},
}

func Run(httpHandler http.Handler) {
	httpServer = &http.Server{
		Addr:        "0.0.0.0:8080",
		Handler:     httpHandler,
		IdleTimeout: time.Minute,
	}
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("service start error:%v", err)
	}
}

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
	log.Printf("service is stopping after %v", timeout.String())
}
