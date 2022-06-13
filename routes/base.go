package routes

import (
	"net/http"

	"github.com/Jacksmall/go-api-framework/controllers"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine

	adminProductController controllers.ProductController
)

func init() {
	router = gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  http.StatusText(http.StatusNotFound),
			"data": "",
		})
	})
}

func InitRoutes() *gin.Engine {
	AdminRoute()
	return router
}
