package blog_api

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/skylee/gin-web-layout/api"
	"github.com/skylee/gin-web-layout/api/blog_api/protocol"
	"github.com/skylee/gin-web-layout/internal/service/blog_service"
	"os/exec"
)

type BlogApi struct {
	s       *blog_service.BlogService
	CmdList []context.CancelFunc
}

func NewBlogApi(s *blog_service.BlogService) *BlogApi {
	return &BlogApi{
		s: s,
		CmdList: make([]context.CancelFunc, 10, 20),
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

func (b *BlogApi) StartCmd(c *gin.Context) {

	ctx, cancel := context.WithCancel(context.Background())

	b.CmdList[1] = cancel

	cmd := exec.CommandContext(ctx, "php", "/tmp/index.php")

	var d bytes.Buffer

	cmd.Stdout = &d

	err := cmd.Run()

	if err != nil {
		api.NewResponse(c, err, d.String())
		return
	}
	api.NewResponse(c, nil, d.String())
}

func (b *BlogApi) StopCmd(c *gin.Context) {
	b.CmdList[1]()
	api.NewResponse(c, nil, nil)
}
