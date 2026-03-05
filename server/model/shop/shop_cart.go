package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopCart struct {
	global.GVA_MODEL
	UserID   uint `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;uniqueIndex:idx_user_sku"`
	SkuID    uint `json:"skuId" form:"skuId" gorm:"column:sku_id;comment:SKU ID;uniqueIndex:idx_user_sku"`
	Quantity int  `json:"quantity" form:"quantity" gorm:"column:quantity;comment:数量;default:1"`
	Selected *int `json:"selected" form:"selected" gorm:"column:selected;comment:是否选中0否1是;default:1"`
	Sku      ShopSku `json:"sku" gorm:"foreignKey:SkuID;references:ID"`
	Spu      ShopSpu `json:"spu" gorm:"-"`
}

func (ShopCart) TableName() string {
	return "shop_cart"
}
