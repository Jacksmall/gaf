package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Jacksmall/go-api-framework/conf"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		reqID := time.Now().UnixNano()
		// 设置req_id 变量
		ctx.Set("req_id", reqID)

		// 请求前
		// 1.跨域
		// 2.options 请求类型 直接拦截输出204
		if conf.Config.Server.Cors {
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Headers", "Origin,Content-Type,Accept,App-Client,Authorization,x-tools-app-id")
			ctx.Header("Access-Control-Allow-Methods", "GET,OPTIONS,POST,PUT,DELETE,PATCH")
			if ctx.Request.Method == http.MethodOptions {
				ctx.JSON(http.StatusNoContent, "ok")
				ctx.Abort()
				return
			}
		}
		// todo 记录请求body体

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
