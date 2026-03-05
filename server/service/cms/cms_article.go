package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
	"gorm.io/gorm"
	"math"
)

type CmsArticleService struct {
}

// CreateCmsArticle 创建文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsArticleService *CmsArticleService) CreateCmsArticle(cmsArticle *cms.CmsArticle) (err error) {
	// 取整
	cmsArticle.ReadingTime = math.Ceil(float64(len(cmsArticle.Content) / 100))

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		e := tx.Create(cmsArticle).Error
		if e != nil {
			return e
		}
		var cmsArticleTags []cms.CmsArticleTag
		// 创建文章标签关联
		for _, tagID := range cmsArticle.Tagids {
			cmsArticleTags = append(cmsArticleTags, cms.CmsArticleTag{CmsArticleID: cmsArticle.ID, CmsTagID: uint(tagID)})
		}
		return tx.Create(&cmsArticleTags).Error
	})
}

// DeleteCmsArticle 删除文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsArticleService *CmsArticleService) DeleteCmsArticle(ID string) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if e := tx.Delete(&cms.CmsArticle{}, "id = ?", ID).Error; e != nil {
			return e
		}
		return tx.Delete(&cms.CmsArticleTag{}, "cms_article_id = ?", ID).Error
	})
}

// DeleteCmsArticleByIds 批量删除文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsArticleService *CmsArticleService) DeleteCmsArticleByIds(IDs []string) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if e := tx.Delete(&cms.CmsArticle{}, "id in ?", IDs).Error; e != nil {
			return e
		}
		return tx.Delete(&cms.CmsArticleTag{}, "cms_article_id in ?", IDs).Error
	})
}

// UpdateCmsArticle 更新文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsArticleService *CmsArticleService) UpdateCmsArticle(cmsArticle cms.CmsArticle) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if e := tx.Save(&cmsArticle).Error; e != nil {
			return e
		}
		if e := tx.Delete(&cms.CmsArticleTag{}, "cms_article_id = ?", cmsArticle.ID).Error; e != nil {
			return e
		}
		var cmsArticleTags []cms.CmsArticleTag
		for _, tagID := range cmsArticle.Tagids {
			cmsArticleTags = append(cmsArticleTags, cms.CmsArticleTag{CmsArticleID: cmsArticle.ID, CmsTagID: uint(tagID)})
		}
		return tx.Create(&cmsArticleTags).Error
	})
}

// GetCmsArticle 根据ID获取文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsArticleService *CmsArticleService) GetCmsArticle(ID string) (cmsArticle cms.CmsArticle, err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if e := tx.Where("id = ?", ID).Preload("AuthorInfo").First(&cmsArticle).Error; e != nil {
			return e
		}
		if e := tx.Model(&cms.CmsArticleTag{}).
			Where("cms_article_id = ?", ID).
			Pluck("cms_tag_id", &cmsArticle.Tagids).
			Error; e != nil {
			return e
		}
		return nil
	})
	return
}

// GetCmsArticleInfoList 分页获取文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsArticleService *CmsArticleService) GetCmsArticleInfoList(info cmsReq.CmsArticleSearch) (list []cms.CmsArticle, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cms.CmsArticle{}).Preload("AuthorInfo").Omit("content")
	var cmsArticles []cms.CmsArticle
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Tag != "" {
		var cmsArticleIDS []uint
		er := global.GVA_DB.Model(&cms.CmsArticleTag{}).Where("cms_tag_id = ?", info.Tag).Pluck("cms_article_id", &cmsArticleIDS).Error
		if er != nil {
			return nil, 0, er
		}
		db = db.Where("id in (?)", cmsArticleIDS)
	}
	if info.Author != 0 {
		db = db.Where("author = ?", info.Author)
	}
	if info.Pid != nil {
		db = db.Where("pid = ?", info.Pid)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cmsArticles).Error
	// 查出关联的tabs
	for i := range cmsArticles {
		var cmsArticleTags []cms.CmsArticleTag
		err = global.GVA_DB.Preload("Tags").Where("cms_article_id = ?", cmsArticles[i].ID).Find(&cmsArticleTags).Error
		if err != nil {
			return
		}
		cmsArticles[i].Tag = make([]cms.CmsTag, 0)
		for _, cmsArticleTag := range cmsArticleTags {
			cmsArticles[i].Tag = append(cmsArticles[i].Tag, cmsArticleTag.Tags)
			cmsArticles[i].Tagids = append(cmsArticles[i].Tagids, int(cmsArticleTag.CmsTagID))
		}
	}
	return cmsArticles, total, err
}
