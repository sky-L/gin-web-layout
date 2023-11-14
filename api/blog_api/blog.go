package blog_api

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skylee/gin-web-layout/api"
	"github.com/skylee/gin-web-layout/api/blog_api/protocol"
	"github.com/skylee/gin-web-layout/internal/service/blog_service"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type BlogApi struct {
	s       *blog_service.BlogService
	CmdList []context.CancelFunc
}

func NewBlogApi(s *blog_service.BlogService) *BlogApi {
	return &BlogApi{
		s:       s,
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

func downloadHandler(c *gin.Context) {
	filePath := "/path/to/file" // Replace with the actual file path
	file, err := os.Open(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	fileSize := fileInfo.Size()
	rangeHeader := c.GetHeader("Range")
	if rangeHeader != "" {
		rangeParts := strings.Split(rangeHeader, "=")
		rangeValue := strings.Split(rangeParts[1], "-")
		start, _ := strconv.ParseInt(rangeValue[0], 10, 64)
		end := fileSize - 1
		if rangeValue[1] != "" {
			end, _ = strconv.ParseInt(rangeValue[1], 10, 64)
		}

		contentLength := end - start + 1
		c.Header("Content-Length", fmt.Sprintf("%d", contentLength))
		c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
		c.Status(http.StatusPartialContent)

		_, err = file.Seek(start, io.SeekStart)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		io.CopyN(c.Writer, file, contentLength)
	} else {
		c.Header("Content-Length", fmt.Sprintf("%d", fileSize))
		c.Status(http.StatusOK)
		io.Copy(c.Writer, file)
	}
}
