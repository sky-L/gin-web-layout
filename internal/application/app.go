package application

import (
	"github.com/skylee/gin-web-layout/config"
	"github.com/skylee/gin-web-layout/internal/storage"
)

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
