package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopInventorySetting struct {
	global.GVA_MODEL
	SkuID          uint `json:"skuId" form:"skuId" gorm:"column:sku_id;comment:SKU ID(0为全局);uniqueIndex"`
	LowThreshold   int  `json:"lowThreshold" form:"lowThreshold" gorm:"column:low_threshold;comment:低库存阈值;default:10"`
	EmptyThreshold int  `json:"emptyThreshold" form:"emptyThreshold" gorm:"column:empty_threshold;comment:缺货阈值;default:0"`
}

func (ShopInventorySetting) TableName() string {
	return "shop_inventory_setting"
}
