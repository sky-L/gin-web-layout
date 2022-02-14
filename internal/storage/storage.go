package storage

import (
	"github.com/skylee/gin-web-layout/config"
	"github.com/skylee/gin-web-layout/internal/storage/default_db"
)

type Storage struct {
	Default *default_db.Mysql
}

func NewStorage(mysql config.MySqlConfig) *Storage {
	return &Storage{
		Default: default_db.NewMysql(mysql.Default),
	}
}
