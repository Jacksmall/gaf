package models

import "time"

type NftGoods struct {
	ID                      uint      `json:"id" gorm:"primaryKey;comment:主键id"`
	NftName                 string    `json:"nftName" gorm:"column:nftName;size:64;default:'';comment:名称（可重复）"`
	NftSymbol               string    `json:"nftSymbol" gorm:"column:nftSymbol;size:64;default:'';comment:简称（不可重复）"`
	NftType                 uint8     `json:"nftType" gorm:"column:nftType;default:0;comment:nft类型 0为数字会员卡 1为数字藏品"`
	IssuerInfo              string    `json:"issuerInfo" gorm:"column:issuerInfo;size:256;default:'';comment:发行方"`
	NftOwnerAddress         string    `json:"nftOwnerAddress" gorm:"column:nftOwnerAddress;size:64;default:'';comment:发行方地址"`
	MaxSerialNum            uint      `json:"maxSerialNum" gorm:"column:maxSerialNum;default:0;comment:开卡排名最大序号"`
	DistributeUniqueUserNum uint      `json:"distributeUniqueUserNum" gorm:"column:distributeUniqueUserNum;default:0;comment:已分配去重用户数量"`
	SaleStatus              uint8     `json:"saleStatus" gorm:"column:saleStatus;default:1;comment:在售状态 0 仓库中 1 售卖中 2 已售罄"`
	SupplierId              uint      `json:"supplierId" gorm:"column:supplierId;default:0;comment:商家id"`
	AppId                   uint      `json:"appId" gorm:"column:appId;default:0;comment:应用id"`
	TempTxId                string    `json:"tempTxId" gorm:"column:tempTxId;size:100;default:'';comment:生成的临时交易单号"`
	TxId                    string    `json:"txId" gorm:"column:txId;size:100;default:'';comment:交易哈希"`
	IsDeleted               uint8     `json:"isDeleted" gorm:"column:isDeleted;default:0;comment:是否删除:1是;0否"`
	CreatedTime             time.Time `json:"createdTime" gorm:"column:createdTime;comment:创建时间"`
	UpdatedTime             time.Time `json:"updatedTime" gorm:"column:updatedTime;comment:更新时间"`
}
