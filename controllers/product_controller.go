package controllers

import (
	"log"
	"net/http"

	"github.com/Jacksmall/go-api-framework/entry"
	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	var req entry.AdminProductListReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		log.Fatalf("query params error: %v", err)
		return
	}

	list, total, err := AdminProductService.GetProducts(req)
	if err != nil {
		log.Fatalf("admin product controller get products error: %v", err)
		return
	}

	data := entry.PageRes{
		Total: total,
		List:  list,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func (p *ProductController) GetProduct(ctx *gin.Context) {
	AdminProductService.GetProduct(ctx)
}
