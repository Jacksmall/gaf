package entry

type NftInit struct {
	Name              string `form:"name" json:"name" binding:"required"`
	PicUrl            string `form:"pic_url" json:"pic_url" binding:"required"`
	Number            int    `form:"number" json:"number" binding:"required"`
	IssuerInfo        string `form:"issuer_info" json:"issuer_info" binding:"required"`
	IssuerDescription string `form:"issuer_description" json:"issuer_description" binding:"required"`
	IsOpenCardLimit   int8   `form:"is_open_card_limit" json:"is_open_card_limit"`
	CardLimit         int8   `form:"card_limit" json:"card_limit"`
	Desc              string `form:"desc" json:"desc" binding:"required"`
	SupplierId        int    `form:"supplier_id" json:"supplier_id" binding:"required"`
	AppId             int    `form:"app_id" json:"app_app" binding:"required"`
	PrizeList         []PrizeDetail
	BaseTimeField
}

type PrizeDetail struct {
	GoodsId int  `form:"goods_id" json:"goods_id"`
	Type    int8 `form:"type" json:"type" binding:"required"`
	Number  int  `form:"number" json:"number" binding:"required"`
}
