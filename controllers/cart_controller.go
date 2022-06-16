package controllers

import (
	"log"
	"net/http"

	"github.com/Jacksmall/go-api-framework/entry"
	"github.com/Jacksmall/go-api-framework/services"
	"github.com/gin-gonic/gin"
)

type CartController struct{}

// 加入购物车
// route POST /api/v1/cart/add
func (c *CartController) Add(ctx *gin.Context) {
	var req entry.APICartAddReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Fatalf("Cart add request: %v", err)
	}
	userID := 10001
	OpenUserID := "10001"
	id, err := services.ApiCartService.Add(uint32(userID), OpenUserID, req)
	if err != nil {
		log.Fatalf("Error cart service add: %v", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": id,
	})
}
