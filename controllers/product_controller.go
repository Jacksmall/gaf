package controllers

import (
	"github.com/Jacksmall/go-api-framework/entry"
	"github.com/Jacksmall/go-api-framework/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type ProductController struct{}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	var respMsg utils.RespMsg
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
	respMsg.Suc(ctx, 0, data)
}

func (p *ProductController) GetProduct(ctx *gin.Context) {
	AdminProductService.GetProduct(ctx)
}

func (p *ProductController) createProduct(ctx *gin.Context) {

}
