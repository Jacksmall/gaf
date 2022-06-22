package models

import (
	"net/http"

	"github.com/Jacksmall/go-api-framework/database"
	"github.com/Jacksmall/go-api-framework/entry"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string `form:"code" json:"code" xml:"code" binding:"required"`
	Price uint
}

func GetProducts(req entry.AdminProductListReq) (list []Product, total int64, err error) {
	db := database.DBConn
	err = db.Model(&Product{}).Count(&total).Error
	if total > 0 {
		err = db.Limit(int(req.GetLimit())).Offset(int(req.GetOffset())).Find(&list).Error
	}
	return
}

func GetProduct(ctx *gin.Context) {
	db := database.DBConn
	var product Product
	db.First(&product, ctx.Param("id"))
	ctx.JSON(http.StatusOK, product)
}

func CreateProduct(req entry.AdminAddProductReq) error {
	db := database.DBConn
	tx := db.Begin()
	product := &Product{
		Code:  req.Code,
		Price: uint(req.Price),
	}
	if err := db.Create(product).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func UpdateProduct(ctx *gin.Context) error {
	db := database.DBConn
	var product Product
	db = db.Take(&product, ctx.Param("id"))
	if db.Error != nil {
		return db.Error
	}
	if err := ctx.ShouldBind(&product); err != nil {
		return err
	}
	if err := db.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(ctx *gin.Context) error {
	db := database.DBConn
	var product Product
	db = db.Take(&product, ctx.Param("id"))
	if db.Error != nil {
		return db.Error
	}
	if err := db.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
