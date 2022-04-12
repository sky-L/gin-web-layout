package api

import (
	"github.com/gin-gonic/gin"
	"github.com/skylee/gin-web-layout/pkg/appcode"
	"net/http"
)

const SuccessCode = 0

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func NewResponse(c *gin.Context, err error, data interface{}) {
	res := &Response{Data: data}

	switch typed := err.(type) {
	case appcode.AppError:
		res.Code = typed.Code
		res.Message = typed.Msg
	default:
		if typed != nil {
			res.Code = -1
			res.Message = err.Error()
		} else {
			res.Code = SuccessCode
			res.Message = "success"
		}
	}
	c.JSON(http.StatusOK, res)
}
