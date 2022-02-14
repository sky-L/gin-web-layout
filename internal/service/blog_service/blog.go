package blog_service

import "github.com/skylee/gin-web-layout/internal/repository/blog"

type BlogService struct {
	repo *blog.BlogRepo
}

func NewBlogService(repo *blog.BlogRepo) *BlogService {
	return &BlogService{
		repo: repo,
	}
}
