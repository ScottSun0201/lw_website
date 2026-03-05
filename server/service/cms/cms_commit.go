package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
)

type CmsCommitService struct {
}

// CreateCmsCommit 创建评论记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCommitService *CmsCommitService) CreateCmsCommit(cmsCommit *cms.CmsCommit) (err error) {
	err = global.GVA_DB.Model(&cms.CmsCommit{}).Create(cmsCommit).Error
	return err
}

// DeleteCmsCommit 删除评论记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCommitService *CmsCommitService) DeleteCmsCommit(ID string) (err error) {
	err = global.GVA_DB.Delete(&cms.CmsCommit{}, "id = ?", ID).Error
	return err
}

// DeleteCmsCommitByIds 批量删除评论记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCommitService *CmsCommitService) DeleteCmsCommitByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cms.CmsCommit{}, "id in ?", IDs).Error
	return err
}

// UpdateCmsCommit 更新评论记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCommitService *CmsCommitService) UpdateCmsCommit(cmsCommit cms.CmsCommit) (err error) {
	err = global.GVA_DB.Save(&cmsCommit).Error
	return err
}

// GetCmsCommit 根据ID获取评论记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCommitService *CmsCommitService) GetCmsCommit(ID string) (cmsCommit cms.CmsCommit, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cmsCommit).Error
	return
}

// GetCmsCommitInfoList 分页获取评论记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCommitService *CmsCommitService) GetCmsCommitInfoList(info cmsReq.CmsCommitSearch) (list []cms.CmsCommit, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cms.CmsCommit{}).Preload("User")
	var cmsCommits []cms.CmsCommit
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.CmsArticleID != 0 {
		db = db.Where("cms_article_id = ?", info.CmsArticleID)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cmsCommits).Error
	return cmsCommits, total, err
}
