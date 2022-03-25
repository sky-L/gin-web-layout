package blog

import (
	"github.com/skylee/gin-web-layout/internal/models/default_db"
	"github.com/skylee/gin-web-layout/internal/storage"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	FindById(id int) (default_db.Blog, error)
}

type BlogRepo struct {
	DB *gorm.DB
}

func NewBlogRepo(storage *storage.Storage) RepositoryInterface {
	return &BlogRepo{storage.Default.DB}
}

func (b *BlogRepo) FindById(id int) (default_db.Blog, error) {
	resp := default_db.Blog{}
	err := b.DB.Where("id = ?", id).Find(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}
