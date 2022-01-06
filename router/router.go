package router

import (
	"github.com/gin-gonic/gin"
	"time"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.JSON(200, "ok")
	})
}
