package default_db

import (
	"github.com/skylee/gin-web-layout/config"
	"github.com/thinkeridea/go-extend/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	DB *gorm.DB
}

func NewMysql(db config.DB) *Mysql {
	s := &Mysql{
		DB: helper.Must(gorm.Open(mysql.Open(db.DataSourceName))).(*gorm.DB),
	}

	s.DB.Debug()

	return s
}
