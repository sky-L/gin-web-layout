package default_db

import (
	"database/sql"
	"github.com/skylee/gin-web-layout/config"
	"github.com/thinkeridea/go-extend/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Mysql struct {
	DB *gorm.DB
}

func NewMysql(db config.DB) *Mysql {
	client := helper.Must(gorm.Open(mysql.Open(db.DataSourceName))).(*gorm.DB)
	sqlDB := helper.Must(client.DB()).(*sql.DB)
	sqlDB.SetMaxIdleConns(db.MaxIdleConns)
	sqlDB.SetMaxOpenConns(db.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Minute)

	client.Debug()

	s := &Mysql{
		DB: client,
	}

	return s
}
