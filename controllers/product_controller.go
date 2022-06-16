package controllers

import (
	"log"
	"net/http"

	"github.com/Jacksmall/go-api-framework/entry"
	"github.com/gin-gonic/gin"
)

type ProductController struct{}

// 分页获取商品列表
func (p *ProductController) GetProducts(ctx *gin.Context) {
	var req entry.AdminProductListReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		log.Fatalf("query params error: %v", err)
	}

	list, total, err := AdminProductService.GetProducts(req)
	if err != nil {
		log.Fatalf("admin product controller get products error: %v", err)
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

// 获取指定商品
func (p *ProductController) GetProduct(ctx *gin.Context) {
	AdminProductService.GetProduct(ctx)
}

// 创建商品
func (p *ProductController) AddProduct(ctx *gin.Context) {
	var req entry.AdminAddProductReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Fatalf("添加商品参数错误: %v", err)
	}

	if err := AdminProductService.AddProduct(req); err != nil {
		log.Fatalf("admin add product error: %v", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": req,
	})
}

// 删除指定商品
func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	if err := AdminProductService.DeleteProduct(ctx); err != nil {
		log.Fatalf("admin delete product error: %v", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": true,
	})
}

// 更新指定商品
func (p *ProductController) UpdateProduct(ctx *gin.Context) {
	if err := AdminProductService.UpdateProduct(ctx); err != nil {
		log.Fatalf("admin update product error: %v", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": true,
	})

}
