package blog_service

import (
	"github.com/skylee/gin-web-layout/internal/models/default_db"
	"github.com/skylee/gin-web-layout/internal/repository/blog"
)

type BlogService struct {
	repo blog.RepositoryInterface
}

func NewBlogService(repo blog.RepositoryInterface) *BlogService {
	return &BlogService{
		repo: repo,
	}
}

func (b *BlogService) List(id int) (default_db.Blog, error) {
	// 1 直接调用 model
	// b.repo.DB.Model().Find().Error
	// 2 使用repo
	return b.repo.FindById(id)
}
