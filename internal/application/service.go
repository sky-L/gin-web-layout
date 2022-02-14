package application

import "github.com/skylee/gin-web-layout/internal/service/blog_service"

type Service struct {
	BlogService *blog_service.BlogService
}

func NewService(repo *Repository) *Service {
	return &Service{
		BlogService: blog_service.NewBlogService(repo.Blog),
	}
}
