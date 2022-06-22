package controllers

import (
	"github.com/Jacksmall/go-api-framework/entry"
	"github.com/gin-gonic/gin"
)

type NftController struct{}

func (nc *NftController) CreateInit(c *gin.Context) (initId int, err error) {
	nftInit := entry.NftInit{}
	if err := c.ShouldBind(&nftInit); err != nil {
		return 0, err
	}
	initId, err = AdminNftInitService.CreateInit(c, nftInit)
	if err != nil {
		return 0, err
	}
	return initId, nil
}
