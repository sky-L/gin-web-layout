package blog_api

import (
	"github.com/gin-gonic/gin"
	"github.com/skylee/gin-web-layout/api"
	"github.com/skylee/gin-web-layout/api/blog_api/protocol"
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
	req := protocol.BlogListReq{}
	var err error
	err = c.ShouldBind(&req)
	if err != nil {
		api.NewResponse(c, err, nil)
		return
	}

	data, err := b.s.List(req.Id)
	if err != nil {
		return
	}
	api.NewResponse(c, nil, data)
}
