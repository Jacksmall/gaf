package entry

type AdminProductListReq struct {
	PageReq
}

type AdminAddProductReq struct {
	Code  string `form:"code" json:"code" binding:"required"`
	Price int64  `form:"price" json:"price" binding:"required"`
}
