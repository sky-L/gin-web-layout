package api

import (
	"github.com/gin-gonic/gin"
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
	if err != nil {
		res.Code = -1
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	res.Code = SuccessCode
	res.Message = "操作成功"
	c.JSON(http.StatusOK, res)
}
