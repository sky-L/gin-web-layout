package application

import (
	"github.com/skylee/gin-web-layout/internal/repository/blog"
	"github.com/skylee/gin-web-layout/internal/storage"
)

// 数据原子方法

type Repository struct {
	Blog *blog.BlogRepo
}

func NewRepository(storage *storage.Storage) *Repository  {
	return &Repository{
		Blog: blog.NewBlogRepo(storage),
	}
}