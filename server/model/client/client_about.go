// 自动生成模板ClientAbout
package client

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 关于我们 结构体  ClientAbout
type ClientAbout struct {
	global.GVA_MODEL
	Name   string `json:"name" form:"name" gorm:"column:name;comment:员工名称;" binding:"required"`          //员工名称
	Desc   string `json:"desc" form:"desc" gorm:"column:desc;comment:员工介绍;size:500;" binding:"required"` //员工介绍
	Title  string `json:"title" form:"title" gorm:"column:title;comment:员工职位;" binding:"required"`       //员工职位
	Avatar string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;"`                         //头像
}

type ClientAboutInfo struct {
	global.GVA_MODEL
	Info datatypes.JSON `json:"info" form:"info" gorm:"column:info;comment:关于我们信息;type:json"`
}

// TableName 关于我们 ClientAbout自定义表名 client_about
func (ClientAbout) TableName() string {
	return "client_about"
}

// TableName 关于我们 ClientAbout自定义表名 client_about
func (ClientAboutInfo) TableName() string {
	return "client_about_info"
}
