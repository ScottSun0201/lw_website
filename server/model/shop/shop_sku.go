package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopSku struct {
	global.GVA_MODEL
	SpuID       uint    `json:"spuId" form:"spuId" gorm:"column:spu_id;comment:SPU ID"`
	SkuCode     string  `json:"skuCode" form:"skuCode" gorm:"column:sku_code;comment:SKU编码;uniqueIndex"`
	Name        string  `json:"name" form:"name" gorm:"column:name;comment:SKU名称"`
	SpecData    string  `json:"specData" form:"specData" gorm:"column:spec_data;comment:规格数据JSON;type:text"`
	Price       float64 `json:"price" form:"price" gorm:"column:price;comment:销售价;type:decimal(10,2)"`
	MarketPrice float64 `json:"marketPrice" form:"marketPrice" gorm:"column:market_price;comment:市场价;type:decimal(10,2)"`
	CostPrice   float64 `json:"costPrice" form:"costPrice" gorm:"column:cost_price;comment:成本价;type:decimal(10,2)"`
	Image       string  `json:"image" form:"image" gorm:"column:image;comment:SKU图片"`
	Weight      float64 `json:"weight" form:"weight" gorm:"column:weight;comment:重量kg;type:decimal(10,2);default:0"`
	Status      *int    `json:"status" form:"status" gorm:"column:status;comment:状态0禁用1启用;default:1"`
	Stock       int     `json:"stock" gorm:"-"`
}

func (ShopSku) TableName() string {
	return "shop_sku"
}
