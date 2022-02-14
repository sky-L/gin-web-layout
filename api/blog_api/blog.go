package blog_api

import (
	"github.com/gin-gonic/gin"
	"github.com/skylee/gin-web-layout/internal/service/blog_service"
)

type BlogApi struct {
	s *blog_service.BlogService
}

func NewBlogApi(s *blog_service.BlogService) *BlogApi {
	return &BlogApi{
		s: s,
	}
}

func (b *BlogApi) List(c *gin.Context) {

}
