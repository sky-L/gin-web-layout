package blog_service

import (
	"github.com/skylee/gin-web-layout/internal/models/blog"
	service "github.com/skylee/gin-web-layout/internal/repository/blog"
)

type BlogService struct {
	repo service.RepositoryInterface
}

func NewBlogService(repo service.RepositoryInterface) *BlogService {
	return &BlogService{
		repo: repo,
	}
}

func (b *BlogService) List(id int) (*blog.Blog, error) {
	return b.repo.FindById(id)
}
