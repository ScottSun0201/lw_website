// 自动生成模板CmsArticle
package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// 文章 结构体  CmsArticle
type CmsArticle struct {
	global.GVA_MODEL
	Title       string         `json:"title" form:"title" gorm:"column:title;comment:标题;"binding:"required"`                 //标题
	Content     string         `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"binding:"required"` //内容
	Desc        string         `json:"desc" form:"desc" gorm:"column:desc;type:text;comment:简介;"`                            //简介
	Image       string         `json:"image" form:"image" gorm:"column:image;comment:缩略图;"`                                  //缩略图
	Flag        string         `json:"flag" form:"flag" gorm:"column:flag;comment:;"`                                        //文章标识
	Tagids      []int          `json:"tagids" form:"tagids" gorm:"-"`                                                        //标签
	Tag         []CmsTag       `json:"tag" gorm:"-"`
	AuthorInfo  system.SysUser `json:"authorInfo" gorm:"foreignKey:Author;references:ID"`
	Author      uint           `json:"author" form:"author" gorm:"column:author;comment:作者"` //作者
	Plate       CmsPlate       `json:"plate" gorm:"foreignKey:Pid;references:ID"`
	ReadingTime float64        `json:"readingTime" form:"readingTime" gorm:"column:reading_time;comment:阅读时间;"` //阅读时间
}

type CmsArticleTag struct {
	CmsArticleID uint   `json:"cmsArticleID" form:"cmsArticleID" gorm:"column:cms_article_id;comment:文章ID;"` //文章ID
	CmsTagID     uint   `json:"cmsTagID" form:"cmsTagID" gorm:"column:cms_tag_id;comment:标签ID;"`             //标签ID
	Tags         CmsTag `json:"tags" gorm:"foreignKey:CmsTagID;references:ID"`
}

// TableName 文章 CmsArticle自定义表名 cms_article
func (CmsArticleTag) TableName() string {
	return "cms_article_tag"
}

// TableName 文章 CmsArticle自定义表名 cms_article
func (CmsArticle) TableName() string {
	return "cms_article"
}
