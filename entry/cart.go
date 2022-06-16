package entry

type APICartAddReq struct {
	GoodsID uint64 `form:"goods_id" json:"goods_id" binding:"required" comment:"商品id"`
	SkuID   uint64 `form:"sku_id" json:"sku_id" binding:"required" comment:"商品sku_id"`
	Num     uint32 `form:"num" json:"num" binding:"required,min=1" comment:"商品数量"`
}
