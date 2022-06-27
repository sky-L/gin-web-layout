package blog_storage

import (
	"database/sql"
	"github.com/skylee/gin-web-layout/config"
	"github.com/skylee/gin-web-layout/internal/models/blog"
	"github.com/thinkeridea/go-extend/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Storage struct {
	db *gorm.DB
}

func NewMysql(db config.DB) *Storage {
	client := helper.Must(gorm.Open(mysql.Open(db.DataSourceName))).(*gorm.DB)
	sqlDB := helper.Must(client.DB()).(*sql.DB)
	sqlDB.SetMaxIdleConns(db.MaxIdleConns)
	sqlDB.SetMaxOpenConns(db.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Minute)

	client.Debug()

	s := &Storage{
		db: client,
	}

	return s
}

func (s *Storage) FindById(id int) (*blog.Blog, error) {
	resp := blog.Blog{}
	err := s.db.Where("id = ?", id).First(&resp).Error

	// never reach
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &resp, nil
}
