package services

import (
	"github.com/Jacksmall/go-api-framework/entry"
	"github.com/gin-gonic/gin"
)

type NftInitService struct{}

func (s *NftInitService) CreateInit(c *gin.Context, ntfInit entry.NftInit) (initId int, err error) {
	// todo 参数验证
	// todo 调用models 插入db-mysql 基础信息
	// todo 调用mongodb models 插入db-mongodb 所有配置 包括奖品
	return 0, nil
}
