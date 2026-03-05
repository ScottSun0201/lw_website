package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
)

type CmsBannerService struct {
}

// CreateCmsBanner 创建轮播图记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsBannerService *CmsBannerService) CreateCmsBanner(cmsBanner *cms.CmsBanner) (err error) {
	err = global.GVA_DB.Create(cmsBanner).Error
	return err
}

// DeleteCmsBanner 删除轮播图记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsBannerService *CmsBannerService) DeleteCmsBanner(ID string) (err error) {
	err = global.GVA_DB.Delete(&cms.CmsBanner{}, "id = ?", ID).Error
	return err
}

// DeleteCmsBannerByIds 批量删除轮播图记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsBannerService *CmsBannerService) DeleteCmsBannerByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cms.CmsBanner{}, "id in ?", IDs).Error
	return err
}

// UpdateCmsBanner 更新轮播图记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsBannerService *CmsBannerService) UpdateCmsBanner(cmsBanner cms.CmsBanner) (err error) {
	err = global.GVA_DB.Save(&cmsBanner).Error
	return err
}

// GetCmsBanner 根据ID获取轮播图记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsBannerService *CmsBannerService) GetCmsBanner(ID string) (cmsBanner cms.CmsBanner, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cmsBanner).Error
	return
}

// GetCmsBannerInfoList 分页获取轮播图记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsBannerService *CmsBannerService) GetCmsBannerInfoList(info cmsReq.CmsBannerSearch) (list []cms.CmsBanner, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cms.CmsBanner{})
	var cmsBanners []cms.CmsBanner
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

	err = db.Find(&cmsBanners).Error
	return cmsBanners, total, err
}
