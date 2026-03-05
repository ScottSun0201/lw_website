package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopFavorite struct {
	global.GVA_MODEL
	UserID uint `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;uniqueIndex:idx_user_spu"`
	SpuID  uint `json:"spuId" form:"spuId" gorm:"column:spu_id;comment:商品ID;uniqueIndex:idx_user_spu"`
}

func (ShopFavorite) TableName() string {
	return "shop_favorite"
}
