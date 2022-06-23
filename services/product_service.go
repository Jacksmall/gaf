package services

import (
	"github.com/Jacksmall/go-api-framework/entry"
	"github.com/Jacksmall/go-api-framework/models"
	"github.com/gin-gonic/gin"
)

type ProductService struct{}

func (s ProductService) GetProducts(req entry.AdminProductListReq) (list []models.Product, total int64, err error) {
	where := make([]models.Where, 0)
	list = make([]models.Product, 0)
	// 以下两种写法 将 model 的处理放在 service 层
	total, err = models.M(ProductModel).ListPageWithCount(where, req.GetOffset(), req.GetLimit(), "id desc", &list)
	// 放到 model 层
	// list, total, err = models.GetProducts(req)
	return
}

func (s ProductService) GetProduct(ctx *gin.Context) {
	models.GetProduct(ctx)
}

func (s ProductService) AddProduct(req entry.AdminAddProductReq) (err error) {
	if err = models.CreateProduct(req); err != nil {
		return err
	}
	return nil
}

func (s ProductService) UpdateProduct(ctx *gin.Context) (err error) {
	if err = models.UpdateProduct(ctx); err != nil {
		return err
	}
	return nil
}

func (s ProductService) DeleteProduct(ctx *gin.Context) (err error) {
	if err = models.DeleteProduct(ctx); err != nil {
		return err
	}
	return nil
}
