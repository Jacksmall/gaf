package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		reqID := time.Now().UnixNano()

		// 设置req_id 变量
		ctx.Set("req_id", reqID)

		// 请求前

		ctx.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的status
		status := ctx.Writer.Status()
		log.Println(status)
	}
}

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := recover(); err != nil {
			var eErr error
			switch t := err.(type) {
			case error:
				eErr = fmt.Errorf("%v", err)
			case string:
				eErr = fmt.Errorf("%s", t)
			default:
				eErr = fmt.Errorf("%s", t)
			}
			// todo log panic
			panic(eErr)
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)

		ctx.Next()
	}
}
