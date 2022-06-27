package storage

import (
	"github.com/skylee/gin-web-layout/config"
	"github.com/skylee/gin-web-layout/internal/storage/blog_storage"
)

type Storage struct {
	Default *blog_storage.Storage
}

func NewStorage(mysql config.MySqlConfig) *Storage {
	return &Storage{
		Default: blog_storage.NewMysql(mysql.Default),
	}
}
