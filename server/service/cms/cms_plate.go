package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
)

type CmsPlateService struct {
}

// CreateCmsPlate 创建板块记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsPlateService *CmsPlateService) CreateCmsPlate(cmsPlate *cms.CmsPlate) (err error) {
	err = global.GVA_DB.Create(cmsPlate).Error
	return err
}

// DeleteCmsPlate 删除板块记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsPlateService *CmsPlateService) DeleteCmsPlate(ID string) (err error) {
	err = global.GVA_DB.Delete(&cms.CmsPlate{}, "id = ?", ID).Error
	return err
}

// DeleteCmsPlateByIds 批量删除板块记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsPlateService *CmsPlateService) DeleteCmsPlateByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cms.CmsPlate{}, "id in ?", IDs).Error
	return err
}

// UpdateCmsPlate 更新板块记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsPlateService *CmsPlateService) UpdateCmsPlate(cmsPlate cms.CmsPlate) (err error) {
	err = global.GVA_DB.Save(&cmsPlate).Error
	return err
}

// GetCmsPlate 根据ID获取板块记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsPlateService *CmsPlateService) GetCmsPlate(ID string) (cmsPlate cms.CmsPlate, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cmsPlate).Error
	return
}

// GetCmsPlateInfoList 分页获取板块记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsPlateService *CmsPlateService) GetCmsPlateInfoList() (list []cms.CmsPlate, err error) {
	// 创建db
	db := global.GVA_DB.Model(&cms.CmsPlate{}).Order("sort")
	var cmsPlates []cms.CmsPlate
	if err != nil {
		return
	}

	err = db.Find(&cmsPlates).Error

	var rootPlates []cms.CmsPlate

	for i := range cmsPlates {
		if cmsPlates[i].Pid == 0 {
			rootPlates = append(rootPlates, cmsPlates[i])
		}
	}
	list = SetChildren(rootPlates, cmsPlates)

	return list, err
}

func SetChildren(rootPlates, cmsPlates []cms.CmsPlate) (list []cms.CmsPlate) {
	for i := range rootPlates {
		for j := range cmsPlates {
			if rootPlates[i].ID == cmsPlates[j].Pid {
				rootPlates[i].Children = append(rootPlates[i].Children, cmsPlates[j])
			}
		}
		list = append(list, rootPlates[i])
	}
	return list
}
