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

func CreateProduct(ctx *gin.Context) {
	var product Product
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	db := database.DBConn
	db.Create(&product)
	ctx.JSON(http.StatusOK, product)
}

// func UpdateProduct(ctx *fiber.Ctx) error {
// 	db := database.DBConn
// 	var product Product
// 	db = db.Take(&product, ctx.Params("id"))
// 	if db.Error != nil {
// 		return ctx.SendStatus(http.StatusNotFound)
// 	}
// 	if err := ctx.BodyParser(&product); err != nil {
// 		return ctx.SendStatus(http.StatusBadRequest)
// 	}
// 	db.Save(&product)
// 	return ctx.JSON(product)
// }

// func DeleteProduct(ctx *fiber.Ctx) error {
// 	db := database.DBConn
// 	var product Product
// 	db = db.Take(&product, ctx.Params("id"))
// 	if db.Error != nil {
// 		return ctx.SendStatus(http.StatusNotFound)
// 	}
// 	db.Delete(&product)
// 	return ctx.SendStatus(http.StatusOK)
// }
