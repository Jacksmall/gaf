package models

type Goods struct {
	GoodsId             uint   `json:"goodsId" gorm:"column:goodsId;primaryKey;comment:商品id"`
	ThirdGoodsId        string `json:"thirdGoodsId" gorm:"column:thirdGoodsId;size:100;not null;comment:第三方商品id"`
	Channel             string `json:"channel" gorm:"column:channel;size:60;not null;comment:商品渠道标识"`
	Status              uint   `json:"status" gorm:"column:status;size:1;default:0;comment:商品上下架状态：1上架;0:下架;index:index_s_pt_gt_ge,priority:1"`
	Name                string `json:"name" gorm:"column:name;size:255;not null;comment:商品名称"`
	ImageUrl            string `json:"imageUrl" gorm:"column:imageUrl;size:255;not null;comment:商品缩略图"`
	MarketPrice         string `json:"marketPrice" gorm:"column:marketPrice;size:100;not null;comment:商品市场价 多sku情况 最小值和最大值"`
	CostPrice           string `json:"costPrice" gorm:"column:costPrice;size:100;not null;comment:成本价 我们需要付给第三方的价格"`
	SupplyPrice         string `json:"supplyPrice" gorm:"column:supplyPrice;size:100;not null;comment:供货价 商家付给通兑吧的价格"`
	SaleNumber          uint   `json:"saleNumber" gorm:"column:saleNumber;not null;comment:销量"`
	LocalCatId          string `json:"LocalCatId" gorm:"column:localCatId;size:60;not null;comment:本地分类"`
	ThirdCatId          string `json:"thirdCatId" gorm:"column:thirdCatId;size:60;not null;comment:第三方分类"`
	GoodsType           uint   `json:"goodsType" gorm:"column:goodsType;size:1;not null;comment:商品类型  1实物商品 2虚拟商品 8虚拟潮玩 10会员卡;index:index_s_pt_gt_ge,priority:2"`
	UseType             int    `json:"useType" gorm:"column:useType;size:1;not null;comment:商品用途 1商品  2奖品"`
	IsDeleted           int    `json:"isDeleted" gorm:"column:isDeleted;size:1;not null;default:0;comment:是否删除  1表示删除"`
	IsTodo              int    `json:"isTodo" gorm:"column:isTodo;size:1;not null;default:0;comment:是否待修改 1表示待修改"`
	IsSingle            int    `json:"isSingle" gorm:"column:isSingle;size:1;not null;comment:是否是单规格商品 1表示单规格"`
	IsRecommend         int    `json:"isRecommend" gorm:"column:isRecommend;size:1;not null;default:0;comment:是否首推 1是 0否"`
	State               int    `json:"state" gorm:"column:state;size:1;not null;comment:第三方商城商品上架状态1：已上架 2：已下架"`
	PutAwayTime         int    `json:"putAwayTime" gorm:"column:putAwayTime;not null;comment:最近上架时间"`
	UnShelveTime        int    `json:"unShelveTime" gorm:"column:unShelveTime;not null;comment:最终下架时间"`
	Sort                int    `json:"sort" gorm:"column:sort;not null;comment:商品排序"`
	IsAutoAdd           int    `json:"isAutoAdd" gorm:"column:isAutoAdd;not null;comment:是否自动添加 1自动添加"`
	CreatedAt           int    `json:"createdAt" gorm:"column:createdAt;not null;comment:创建时间"`
	UpdatedAt           int    `json:"updatedAt" gorm:"column:updatedAt;not null;comment:更新时间"`
	Profit              uint   `json:"profit" gorm:"column:profit;not null;default:0;comment:商品利润 单位fen profitType = 2时 表示百分比"`
	ProfitType          uint   `json:"profitType" gorm:"column:profitType;size:1;not null;default:0;comment:商品利润类型 1 现金 2百分比"`
	SupplierChannel     string `json:"supplierChannel" gorm:"column:supplierChannel;size:100;comment:官方自建商品子渠道标识"`
	Stock               int    `json:"stock" gorm:"column:stock;size:1;not null;comment:是否有货：0：表示无货；1:有货"`
	OfficialCouponType  int    `json:"officialCouponType" gorm:"column:officialCouponType;size:1;not null;default:1;comment:优惠券类型 无限：1.卡密兑换 2卡号充值 3手机充值 4账号充值 娱尚：1卡密兑换 2油卡充值 3手机号充值 4Q币充值 官方自建或供应链自建:1：优惠劵码"`
	SupplierLimitNumber int    `json:"supplierLimitNumber" gorm:"column:supplierLimitNumber;not null;default:0;comment:限制商家添加商品数量 0表示不限制"`
	IsCopy              int    `json:"isCopy" gorm:"column:isCopy;not null;default:0;comment:是否复制商品:0:否；1:是"`
}
