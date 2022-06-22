package utils

import (
	"github.com/gin-gonic/gin"
)

type RespMsg struct{}

func (r *RespMsg) Suc(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, gin.H{
		"code": code,
		"msg":  "SUCCESS",
		"data": data,
	})
}

func (r *RespMsg) Err(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, gin.H{
		"code": code,
		"msg":  "FAIL",
		"data": data,
	})
}
