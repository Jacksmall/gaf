package services

import (
	"log"

	"github.com/Jacksmall/go-api-framework/entry"
	"github.com/Jacksmall/go-api-framework/models"
	"github.com/gin-gonic/gin"
)

type ProductService struct{}

func (s *ProductService) GetProducts(req entry.AdminProductListReq) (list []models.Product, total int64, err error) {
	list, total, err = models.GetProducts(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func (s *ProductService) GetProduct(ctx *gin.Context) {
	models.GetProduct(ctx)
}
