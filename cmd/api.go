package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

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
