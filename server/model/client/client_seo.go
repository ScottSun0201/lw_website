// 自动生成模板ClientSEO
package client

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 客户端SEO 结构体  ClientSEO
type ClientSEO struct {
	global.GVA_MODEL
	Title       string `json:"title" form:"title" gorm:"column:title;comment:title;" binding:"required"`                                  //title
	Description string `json:"description" form:"description" gorm:"column:description;comment:description;size:512;" binding:"required"` //description
	Keywords    string `json:"keywords" form:"keywords" gorm:"column:keywords;comment:keywords;" binding:"required"`                      //keywords
	Name        string `json:"name" form:"name" gorm:"column:name;comment:网站名称;" binding:"required"`                                      //name
	Logo        string `json:"logo" form:"logo" gorm:"column:logo;comment:logo;" binding:"required"`                                      //Logo
}

// TableName 客户端SEO ClientSEO自定义表名 client_seo
func (ClientSEO) TableName() string {
	return "client_seo"
}
