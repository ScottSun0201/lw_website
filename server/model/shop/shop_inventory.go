package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopInventory struct {
	global.GVA_MODEL
	SkuID          uint `json:"skuId" form:"skuId" gorm:"column:sku_id;comment:SKU ID;uniqueIndex"`
	TotalStock     int  `json:"totalStock" form:"totalStock" gorm:"column:total_stock;comment:总库存;default:0"`
	LockedStock    int  `json:"lockedStock" form:"lockedStock" gorm:"column:locked_stock;comment:锁定库存;default:0"`
	AvailableStock int  `json:"availableStock" form:"availableStock" gorm:"column:available_stock;comment:可用库存;default:0"`
	Version        int  `json:"version" form:"version" gorm:"column:version;comment:乐观锁版本;default:0"`
	Sku            ShopSku `json:"sku" gorm:"foreignKey:SkuID;references:ID"`
}

func (ShopInventory) TableName() string {
	return "shop_inventory"
}
