package application

import (
	"github.com/skylee/gin-web-layout/api/blog_api"
)

type Api struct {
	Blog *blog_api.BlogApi
}

func NewApi(s *Service) *Api {
	return &Api{
		Blog: blog_api.NewBlogApi(s.BlogService),
	}
}
