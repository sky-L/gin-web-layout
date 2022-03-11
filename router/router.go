package router

import (
	"github.com/gin-gonic/gin"
	"github.com/skylee/gin-web-layout/internal/application"
)

func InitRouter(engine *gin.Engine, app *application.App) {
	engine.GET("/blog/list", app.Api.Blog.List)

	engine.GET("/start", app.Api.Blog.StartCmd)

	engine.GET("/stop", app.Api.Blog.StopCmd)
}
