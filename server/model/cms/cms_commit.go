// 自动生成模板CmsCommit
package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
)

// 评论 结构体  CmsCommit
type CmsCommit struct {
	global.GVA_MODEL
	Info         string             `json:"info" form:"info" gorm:"column:info;comment:评论内容;" binding:"required"`                           //评论内容
	CmsArticleID uint               `json:"cmsArticleID" form:"cmsArticleID" gorm:"column:cms_article_id;comment:文章ID;" binding:"required"` //文章ID
	User         *client.ClientUser `json:"user" gorm:"->;foreignKey:UserID;references:ID"`
	UserID       uint               `json:"userID" form:"userID" gorm:"column:user_id;comment:评论人ID;"` //评论人ID
}

// TableName 评论 CmsCommit自定义表名 cms_commit
func (CmsCommit) TableName() string {
	return "cms_commit"
}
