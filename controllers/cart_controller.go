package controllers

import (
	"github.com/Jacksmall/go-api-framework/entry"
	"github.com/Jacksmall/go-api-framework/helper/response"
	"github.com/Jacksmall/go-api-framework/services"
	"github.com/gin-gonic/gin"
	"log"
)

type CartController struct{}

// Add 加入购物车
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

	rs := response.NewResponse(0, "SUCCESS", id)
	rs.SuccessJSON(ctx)
}
