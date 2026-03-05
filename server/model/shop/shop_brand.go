package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopBrand struct {
	global.GVA_MODEL
	Name   string `json:"name" form:"name" gorm:"column:name;comment:品牌名称" binding:"required"`
	Logo   string `json:"logo" form:"logo" gorm:"column:logo;comment:品牌Logo"`
	Desc   string `json:"desc" form:"desc" gorm:"column:desc;comment:品牌描述;type:text"`
	Sort   int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;default:0"`
	Status *int   `json:"status" form:"status" gorm:"column:status;comment:状态0禁用1启用;default:1"`
}

func (ShopBrand) TableName() string {
	return "shop_brand"
}
