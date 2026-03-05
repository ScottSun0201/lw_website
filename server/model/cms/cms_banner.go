// 自动生成模板CmsBanner
package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 轮播图 结构体  CmsBanner
type CmsBanner struct {
	global.GVA_MODEL
	Title     string `json:"title" form:"title" gorm:"column:title;comment:图片Title属性;"`          //图片Title属性
	Path      string `json:"path" form:"path" gorm:"column:path;comment:点击跳转路径;"`                //点击跳转路径
	Image     string `json:"image" form:"image" gorm:"column:image;comment:图片;"`                 //图片
	NewWindow *bool  `json:"newWindow" form:"newWindow" gorm:"column:new_window;comment:新窗口打开;"` //新窗口打开
	Open      *bool  `json:"open" form:"open" gorm:"column:open;comment:启用;"`                    //启用
}

// TableName 轮播图 CmsBanner自定义表名 cms_banner
func (CmsBanner) TableName() string {
	return "cms_banner"
}
