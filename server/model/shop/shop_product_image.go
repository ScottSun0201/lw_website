package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopProductImage struct {
	global.GVA_MODEL
	SpuID uint   `json:"spuId" form:"spuId" gorm:"column:spu_id;comment:SPU ID"`
	URL   string `json:"url" form:"url" gorm:"column:url;comment:图片地址"`
	Sort  int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;default:0"`
}

func (ShopProductImage) TableName() string {
	return "shop_product_image"
}
