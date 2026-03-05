package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
)

type ShopNotificationService struct{}

// CreateNotification 创建用户通知
func (s *ShopNotificationService) CreateNotification(userID uint, nType int, title, content, bizType, bizID string) error {
	notification := shop.ShopNotification{
		UserID:  userID,
		Type:    nType,
		Title:   title,
		Content: content,
		Status:  0,
		BizType: bizType,
		BizID:   bizID,
	}
	return global.GVA_DB.Create(&notification).Error
}

// GetUserNotifications 获取用户通知列表
func (s *ShopNotificationService) GetUserNotifications(userID uint, info shopReq.ShopNotificationSearch) (list []shop.ShopNotification, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopNotification{}).Where("user_id = ?", userID)
	if info.Type != nil {
		db = db.Where("type = ?", *info.Type)
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&list).Error
	return
}

// MarkAsRead 标记通知已读
func (s *ShopNotificationService) MarkAsRead(userID uint, id uint) error {
	return global.GVA_DB.Model(&shop.ShopNotification{}).Where("id = ? AND user_id = ?", id, userID).Update("status", 1).Error
}

// MarkAllAsRead 标记所有通知已读
func (s *ShopNotificationService) MarkAllAsRead(userID uint) error {
	return global.GVA_DB.Model(&shop.ShopNotification{}).Where("user_id = ? AND status = 0", userID).Update("status", 1).Error
}

// GetUnreadCount 获取未读通知数
func (s *ShopNotificationService) GetUnreadCount(userID uint) (count int64, err error) {
	err = global.GVA_DB.Model(&shop.ShopNotification{}).Where("user_id = ? AND status = 0", userID).Count(&count).Error
	return
}

// CreateAdminNotification 创建管理端通知
func (s *ShopNotificationService) CreateAdminNotification(nType int, title, content, bizType, bizID string) error {
	notification := shop.ShopAdminNotification{
		Type:    nType,
		Title:   title,
		Content: content,
		Status:  0,
		BizType: bizType,
		BizID:   bizID,
	}
	return global.GVA_DB.Create(&notification).Error
}

// GetAdminNotifications 获取管理端通知列表
func (s *ShopNotificationService) GetAdminNotifications(info shopReq.ShopNotificationSearch) (list []shop.ShopAdminNotification, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopAdminNotification{})
	if info.Type != nil {
		db = db.Where("type = ?", *info.Type)
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&list).Error
	return
}

// MarkAdminRead 标记管理端通知已读
func (s *ShopNotificationService) MarkAdminRead(id uint) error {
	return global.GVA_DB.Model(&shop.ShopAdminNotification{}).Where("id = ?", id).Update("status", 1).Error
}

// MarkAllAdminRead 标记管理端所有通知已读
func (s *ShopNotificationService) MarkAllAdminRead() error {
	return global.GVA_DB.Model(&shop.ShopAdminNotification{}).Where("status = 0").Update("status", 1).Error
}

// GetAdminUnreadCount 获取管理端未读通知数
func (s *ShopNotificationService) GetAdminUnreadCount() (count int64, err error) {
	err = global.GVA_DB.Model(&shop.ShopAdminNotification{}).Where("status = 0").Count(&count).Error
	return
}
