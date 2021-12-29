package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

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
