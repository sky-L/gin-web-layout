package router

import (
	"github.com/gin-gonic/gin"
	"github.com/skylee/gin-web-layout/internal/application"
	"time"
)

func InitRouter(engine *gin.Engine, app *application.App) {
	engine.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.JSON(200, "ok")
	})
}
