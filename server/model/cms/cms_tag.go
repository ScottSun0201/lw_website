// 自动生成模板CmsTag
package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 标签 结构体  CmsTag
type CmsTag struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"` //标签名称
	Desc string `json:"desc" form:"desc" gorm:"column:desc;comment:;"` //描述
}

// TableName 标签 CmsTag自定义表名 cms_tag
func (CmsTag) TableName() string {
	return "cms_tag"
}
