package services

import "github.com/Jacksmall/go-api-framework/entry"

type CartService struct{}

func (s *CartService) Add(userID uint32, openUserID string, req entry.APICartAddReq) (id int, err error) {
	// 防止重复操作

	// 购物车中商品排序

	// 添加相同物品则更新数据否则插入数据

	return 0, nil
}
