package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopCategory struct {
	global.GVA_MODEL
	Name     string         `json:"name" form:"name" gorm:"column:name;comment:分类名称" binding:"required"`
	ParentID uint           `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父级ID;default:0"`
	Level    int            `json:"level" form:"level" gorm:"column:level;comment:层级1/2/3;default:1"`
	Sort     int            `json:"sort" form:"sort" gorm:"column:sort;comment:排序;default:0"`
	Icon     string         `json:"icon" form:"icon" gorm:"column:icon;comment:图标"`
	Status   *int           `json:"status" form:"status" gorm:"column:status;comment:状态0禁用1启用;default:1"`
	Children []ShopCategory `json:"children" gorm:"-"`
}

func (ShopCategory) TableName() string {
	return "shop_category"
}
