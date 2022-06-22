package services

import (
	"github.com/Jacksmall/go-api-framework/database"
	"github.com/Jacksmall/go-api-framework/entry"
	"github.com/Jacksmall/go-api-framework/models"
	"github.com/gin-gonic/gin"
)

type GoodsInterface interface {
	PlaceGoods(ctx *gin.Context, data entry.GoodsForm)
}

type GoodsService struct{}

func (g *GoodsService) PlaceGoods(ctx *gin.Context, data entry.GoodsForm) (goods *models.Goods, err error) {
	goods = &models.Goods{}
	// 将上下文参数绑定到goods结构体指针中
	if err := ctx.ShouldBind(goods); err != nil {
		return nil, err
	}

	// 开始事务
	tx := database.DBConn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 开启事务失败
	if err := tx.Error; err != nil {
		return nil, err
	}
	// 创建商品
	if err := tx.Create(goods).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return goods, tx.Commit().Error
}
