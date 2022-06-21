package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code uint, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func (rs *Response) SuccessJSON(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &rs)
}
