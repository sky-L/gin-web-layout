package blog

import (
	"github.com/skylee/gin-web-layout/internal/models/blog"
	"github.com/skylee/gin-web-layout/internal/storage"
	"github.com/skylee/gin-web-layout/internal/storage/blog_storage"
)

type RepositoryInterface interface {
	FindById(id int) (*blog.Blog, error)
}

type BlogRepo struct {
	RepositoryInterface
	Storage *blog_storage.Storage
}

func NewBlogRepo(storage *storage.Storage) RepositoryInterface {
	return &BlogRepo{nil, storage.Default}
}

func (b *BlogRepo) FindById(id int) (*blog.Blog, error) {
	return b.Storage.FindById(id)
}
