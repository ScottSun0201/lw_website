// 自动生成模板CmsFriend
package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 友情链接 结构体  CmsFriend
type CmsFriend struct {
	global.GVA_MODEL
	Title     string `json:"title" form:"title" gorm:"column:title;comment:标题;"`                 //标题
	Image     string `json:"image" form:"image" gorm:"column:image;comment:友链图片;"`               //友链图片
	Link      string `json:"link" form:"link" gorm:"column:link;comment:链接;"`                    //链接
	NewWindow *bool  `json:"newWindow" form:"newWindow" gorm:"column:new_window;comment:新窗口打开;"` //新窗口打开
	Open      *bool  `json:"open" form:"open" gorm:"column:open;comment:启用;"`                    //启用
}

// TableName 友情链接 CmsFriend自定义表名 cms_friend
func (CmsFriend) TableName() string {
	return "cms_friend"
}

type CmsFriendInfo struct {
	global.GVA_MODEL
	FriendInfo datatypes.JSON `json:"friendInfo" form:"friendInfo" gorm:"column:friend_info;comment:友链信息;type:json"`
}

// TableName 友情链接 CmsFriend自定义表名 cms_friend
func (CmsFriendInfo) TableName() string {
	return "cms_friend_info"
}
