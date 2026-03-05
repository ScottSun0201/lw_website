package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
)

type CmsTagService struct {
}

// CreateCmsTag 创建标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsTagService *CmsTagService) CreateCmsTag(cmsTag *cms.CmsTag) (err error) {
	err = global.GVA_DB.Create(cmsTag).Error
	return err
}

// DeleteCmsTag 删除标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsTagService *CmsTagService) DeleteCmsTag(ID string) (err error) {
	err = global.GVA_DB.Delete(&cms.CmsTag{}, "id = ?", ID).Error
	return err
}

// DeleteCmsTagByIds 批量删除标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsTagService *CmsTagService) DeleteCmsTagByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cms.CmsTag{}, "id in ?", IDs).Error
	return err
}

// UpdateCmsTag 更新标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsTagService *CmsTagService) UpdateCmsTag(cmsTag cms.CmsTag) (err error) {
	err = global.GVA_DB.Save(&cmsTag).Error
	return err
}

// GetCmsTag 根据ID获取标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsTagService *CmsTagService) GetCmsTag(ID string) (cmsTag cms.CmsTag, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cmsTag).Error
	return
}

// GetCmsTagInfoList 分页获取标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsTagService *CmsTagService) GetCmsTagInfoList(info cmsReq.CmsTagSearch) (list []cms.CmsTag, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cms.CmsTag{})
	var cmsTags []cms.CmsTag
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cmsTags).Error
	return cmsTags, total, err
}
