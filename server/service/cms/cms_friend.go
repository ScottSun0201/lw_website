package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
)

type CmsFriendService struct {
}

func (cmsFriendService *CmsFriendService) SetCmsFriendInfo(cmsFriend *cms.CmsFriendInfo) (err error) {
	err = global.GVA_DB.Save(cmsFriend).Error
	return err
}

func (cmsFriendService *CmsFriendService) GetCmsFriendInfo() (cmsFriend *cms.CmsFriendInfo, err error) {
	err = global.GVA_DB.First(&cmsFriend).Error
	return
}

// CreateCmsFriend 创建友情链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsFriendService *CmsFriendService) CreateCmsFriend(cmsFriend *cms.CmsFriend) (err error) {
	err = global.GVA_DB.Create(cmsFriend).Error
	return err
}

// DeleteCmsFriend 删除友情链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsFriendService *CmsFriendService) DeleteCmsFriend(ID string) (err error) {
	err = global.GVA_DB.Delete(&cms.CmsFriend{}, "id = ?", ID).Error
	return err
}

// DeleteCmsFriendByIds 批量删除友情链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsFriendService *CmsFriendService) DeleteCmsFriendByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cms.CmsFriend{}, "id in ?", IDs).Error
	return err
}

// UpdateCmsFriend 更新友情链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsFriendService *CmsFriendService) UpdateCmsFriend(cmsFriend cms.CmsFriend) (err error) {
	err = global.GVA_DB.Save(&cmsFriend).Error
	return err
}

// GetCmsFriend 根据ID获取友情链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsFriendService *CmsFriendService) GetCmsFriend(ID string) (cmsFriend cms.CmsFriend, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cmsFriend).Error
	return
}

// GetCmsFriendInfoList 分页获取友情链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsFriendService *CmsFriendService) GetCmsFriendInfoList(info cmsReq.CmsFriendSearch) (list []cms.CmsFriend, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cms.CmsFriend{})
	var cmsFriends []cms.CmsFriend
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

	err = db.Find(&cmsFriends).Error
	return cmsFriends, total, err
}
