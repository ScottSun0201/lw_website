// 自动生成模板CmsPlate
package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 板块 结构体  CmsPlate
type CmsPlate struct {
	global.GVA_MODEL
	Pid      uint       `json:"pid" form:"pid" gorm:"column:pid;comment:上级id;"`            //上级id
	Name     string     `json:"name" form:"name" gorm:"column:name;comment:名称;"`           //名称
	Desc     string     `json:"desc" form:"desc" gorm:"column:desc;comment:描述;"`           //描述
	Icon     string     `json:"icon" form:"icon" gorm:"column:icon;comment:图标;"`           //图标
	Router   string     `json:"router" form:"router" gorm:"column:router;comment:路由path;"` //详情路由
	Sort     int        `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`           //排序
	Open     *bool      `json:"open" form:"open" gorm:"column:open;comment:启用;"`           //启用
	Children []CmsPlate `json:"children" gorm:"-"`
}

// TableName 板块 CmsPlate自定义表名 cms_plate
func (CmsPlate) TableName() string {
	return "cms_plate"
}
