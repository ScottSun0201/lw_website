package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// Type: 1低库存 2缺货
// Status: 0未处理 1已处理
type ShopInventoryAlert struct {
	global.GVA_MODEL
	SkuID     uint `json:"skuId" form:"skuId" gorm:"column:sku_id;comment:SKU ID;index"`
	Type      int  `json:"type" form:"type" gorm:"column:type;comment:预警类型1低库存2缺货"`
	Threshold int  `json:"threshold" form:"threshold" gorm:"column:threshold;comment:阈值"`
	Current   int  `json:"current" form:"current" gorm:"column:current;comment:当前库存"`
	Status    int  `json:"status" form:"status" gorm:"column:status;comment:状态0未处理1已处理;default:0"`
}

func (ShopInventoryAlert) TableName() string {
	return "shop_inventory_alert"
}
