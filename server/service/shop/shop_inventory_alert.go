package shop

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"gorm.io/gorm"
)

type ShopInventoryAlertService struct{}

func (s *ShopInventoryAlertService) CheckInventoryAlerts() {
	globalSetting := s.GetGlobalSetting()
	lowThreshold := globalSetting.LowThreshold
	emptyThreshold := globalSetting.EmptyThreshold

	var inventories []shop.ShopInventory
	global.GVA_DB.Find(&inventories)

	for _, inv := range inventories {
		setting := s.getSkuSetting(inv.SkuID, lowThreshold, emptyThreshold)
		if inv.AvailableStock <= setting.EmptyThreshold {
			s.createAlertIfNotExists(inv.SkuID, 2, setting.EmptyThreshold, inv.AvailableStock)
		} else if inv.AvailableStock <= setting.LowThreshold {
			s.createAlertIfNotExists(inv.SkuID, 1, setting.LowThreshold, inv.AvailableStock)
		}
	}
}

func (s *ShopInventoryAlertService) createAlertIfNotExists(skuID uint, alertType, threshold, current int) {
	var count int64
	global.GVA_DB.Model(&shop.ShopInventoryAlert{}).Where("sku_id = ? AND type = ? AND status = 0", skuID, alertType).Count(&count)
	if count > 0 {
		return
	}
	alert := shop.ShopInventoryAlert{SkuID: skuID, Type: alertType, Threshold: threshold, Current: current, Status: 0}
	global.GVA_DB.Create(&alert)

	typeStr := "低库存"
	if alertType == 2 {
		typeStr = "缺货"
	}
	var notificationService ShopNotificationService
	go notificationService.CreateAdminNotification(1, typeStr+"预警", fmt.Sprintf("SKU %d 当前库存 %d", skuID, current), "inventory_alert", fmt.Sprintf("%d", alert.ID))
}

func (s *ShopInventoryAlertService) getSkuSetting(skuID uint, defaultLow, defaultEmpty int) shop.ShopInventorySetting {
	var setting shop.ShopInventorySetting
	if err := global.GVA_DB.Where("sku_id = ?", skuID).First(&setting).Error; err == nil {
		return setting
	}
	return shop.ShopInventorySetting{LowThreshold: defaultLow, EmptyThreshold: defaultEmpty}
}

func (s *ShopInventoryAlertService) GetGlobalSetting() shop.ShopInventorySetting {
	var setting shop.ShopInventorySetting
	if err := global.GVA_DB.Where("sku_id = 0").First(&setting).Error; err == gorm.ErrRecordNotFound {
		setting = shop.ShopInventorySetting{SkuID: 0, LowThreshold: 10, EmptyThreshold: 0}
		global.GVA_DB.Create(&setting)
	}
	return setting
}

func (s *ShopInventoryAlertService) GetAlertList(info shopReq.ShopInventoryAlertSearch) (list []shop.ShopInventoryAlert, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopInventoryAlert{})
	if info.Type != nil {
		db = db.Where("type = ?", *info.Type)
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	err = db.Count(&total).Error
	if err != nil { return }
	if limit != 0 { db = db.Limit(limit).Offset(offset) }
	err = db.Order("created_at desc").Find(&list).Error
	return
}

func (s *ShopInventoryAlertService) HandleAlert(id uint) error {
	return global.GVA_DB.Model(&shop.ShopInventoryAlert{}).Where("id = ?", id).Update("status", 1).Error
}

func (s *ShopInventoryAlertService) UpdateSetting(req shopReq.ShopInventorySettingReq) error {
	var setting shop.ShopInventorySetting
	err := global.GVA_DB.Where("sku_id = ?", req.SkuID).First(&setting).Error
	if err == gorm.ErrRecordNotFound {
		setting = shop.ShopInventorySetting{SkuID: req.SkuID, LowThreshold: req.LowThreshold, EmptyThreshold: req.EmptyThreshold}
		return global.GVA_DB.Create(&setting).Error
	}
	return global.GVA_DB.Model(&setting).Updates(map[string]interface{}{
		"low_threshold": req.LowThreshold, "empty_threshold": req.EmptyThreshold,
	}).Error
}
