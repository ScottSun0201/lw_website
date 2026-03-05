package client

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
	clientReq "github.com/flipped-aurora/gin-vue-admin/server/model/client/request"
)

type ClientAboutService struct {
}

func (clientAboutService *ClientAboutService) SetClientAboutInfo(clientAbout *client.ClientAboutInfo) (err error) {
	err = global.GVA_DB.Save(clientAbout).Error
	return err
}

func (clientAboutService *ClientAboutService) GetClientAboutInfo() (clientAbout *client.ClientAboutInfo, err error) {
	err = global.GVA_DB.First(&clientAbout).Error
	return
}

// CreateClientAbout 创建关于我们记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientAboutService *ClientAboutService) CreateClientAbout(clientAbout *client.ClientAbout) (err error) {
	err = global.GVA_DB.Create(clientAbout).Error
	return err
}

// DeleteClientAbout 删除关于我们记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientAboutService *ClientAboutService) DeleteClientAbout(ID string) (err error) {
	err = global.GVA_DB.Delete(&client.ClientAbout{}, "id = ?", ID).Error
	return err
}

// DeleteClientAboutByIds 批量删除关于我们记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientAboutService *ClientAboutService) DeleteClientAboutByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]client.ClientAbout{}, "id in ?", IDs).Error
	return err
}

// UpdateClientAbout 更新关于我们记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientAboutService *ClientAboutService) UpdateClientAbout(clientAbout client.ClientAbout) (err error) {
	err = global.GVA_DB.Save(&clientAbout).Error
	return err
}

// GetClientAbout 根据ID获取关于我们记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientAboutService *ClientAboutService) GetClientAbout(ID string) (clientAbout client.ClientAbout, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&clientAbout).Error
	return
}

// GetClientAboutInfoList 分页获取关于我们记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientAboutService *ClientAboutService) GetClientAboutInfoList(info clientReq.ClientAboutSearch) (list []client.ClientAbout, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&client.ClientAbout{})
	var clientAbouts []client.ClientAbout
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

	err = db.Find(&clientAbouts).Error
	return clientAbouts, total, err
}
