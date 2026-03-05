package shop

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"gorm.io/gorm"
)

type ShopInventoryService struct{}

// InitInventory 初始化SKU库存记录
func (s *ShopInventoryService) InitInventory(tx *gorm.DB, skuID uint, stock int) error {
	inventory := shop.ShopInventory{
		SkuID:          skuID,
		TotalStock:     stock,
		LockedStock:    0,
		AvailableStock: stock,
		Version:        0,
	}
	if err := tx.Create(&inventory).Error; err != nil {
		return err
	}

	// 记录库存日志
	if stock > 0 {
		log := shop.ShopInventoryLog{
			SkuID:     skuID,
			Type:      1, // 入库
			Quantity:  stock,
			BeforeQty: 0,
			AfterQty:  stock,
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}
	}

	return nil
}

// SetStock 管理员设置库存
func (s *ShopInventoryService) SetStock(req shopReq.ShopInventorySetReq, operatorID uint) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var inventory shop.ShopInventory
		if err := tx.Where("sku_id = ?", req.SkuID).First(&inventory).Error; err != nil {
			return err
		}

		// 计算差值
		delta := req.Stock - inventory.TotalStock
		beforeAvailable := inventory.AvailableStock
		newTotalStock := req.Stock
		newAvailableStock := inventory.AvailableStock + delta

		if newAvailableStock < 0 {
			return errors.New("设置失败：可用库存不能为负数")
		}

		// 乐观锁更新库存
		result := tx.Model(&shop.ShopInventory{}).
			Where("sku_id = ? AND version = ?", req.SkuID, inventory.Version).
			Updates(map[string]interface{}{
				"total_stock":     newTotalStock,
				"available_stock": newAvailableStock,
				"version":         gorm.Expr("version + 1"),
			})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("设置失败：数据已被修改，请重试")
		}

		// 记录日志
		logType := 1 // 入库
		quantity := delta
		if delta < 0 {
			logType = 2 // 出库
			quantity = -delta
		}

		inventoryLog := shop.ShopInventoryLog{
			SkuID:      req.SkuID,
			Type:       logType,
			Quantity:   quantity,
			BeforeQty:  beforeAvailable,
			AfterQty:   newAvailableStock,
			OperatorID: operatorID,
		}
		return tx.Create(&inventoryLog).Error
	})
}

// LockStock 锁定库存（乐观锁）
func (s *ShopInventoryService) LockStock(tx *gorm.DB, skuID uint, quantity int, orderNo string) error {
	// 先查询当前库存记录获取version
	var inventory shop.ShopInventory
	if err := tx.Where("sku_id = ?", skuID).First(&inventory).Error; err != nil {
		return err
	}

	beforeAvailable := inventory.AvailableStock

	// 乐观锁更新
	result := tx.Model(&shop.ShopInventory{}).
		Where("sku_id = ? AND available_stock >= ? AND version = ?", skuID, quantity, inventory.Version).
		Updates(map[string]interface{}{
			"locked_stock":    gorm.Expr("locked_stock + ?", quantity),
			"available_stock": gorm.Expr("available_stock - ?", quantity),
			"version":         gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("库存不足")
	}

	// 记录库存日志
	inventoryLog := shop.ShopInventoryLog{
		SkuID:     skuID,
		OrderNo:   orderNo,
		Type:      3, // 锁定
		Quantity:  quantity,
		BeforeQty: beforeAvailable,
		AfterQty:  beforeAvailable - quantity,
	}
	return tx.Create(&inventoryLog).Error
}

// DeductStock 扣减库存（支付成功后）
func (s *ShopInventoryService) DeductStock(tx *gorm.DB, skuID uint, quantity int, orderNo string) error {
	var inventory shop.ShopInventory
	if err := tx.Where("sku_id = ?", skuID).First(&inventory).Error; err != nil {
		return err
	}

	if inventory.LockedStock < quantity {
		return errors.New("锁定库存不足，无法扣减")
	}

	beforeLocked := inventory.LockedStock

	// 乐观锁扣减锁定库存，总库存也减少
	result := tx.Model(&shop.ShopInventory{}).
		Where("sku_id = ? AND locked_stock >= ? AND version = ?", skuID, quantity, inventory.Version).
		Updates(map[string]interface{}{
			"locked_stock": gorm.Expr("locked_stock - ?", quantity),
			"total_stock":  gorm.Expr("total_stock - ?", quantity),
			"version":      gorm.Expr("version + 1"),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("扣减库存失败，数据已被修改，请重试")
	}

	// 记录库存日志
	inventoryLog := shop.ShopInventoryLog{
		SkuID:     skuID,
		OrderNo:   orderNo,
		Type:      5, // 扣减
		Quantity:  quantity,
		BeforeQty: beforeLocked,
		AfterQty:  beforeLocked - quantity,
	}
	return tx.Create(&inventoryLog).Error
}

// ReleaseStock 释放库存（取消订单）
func (s *ShopInventoryService) ReleaseStock(tx *gorm.DB, skuID uint, quantity int, orderNo string) error {
	var inventory shop.ShopInventory
	if err := tx.Where("sku_id = ?", skuID).First(&inventory).Error; err != nil {
		return err
	}

	if inventory.LockedStock < quantity {
		return errors.New("锁定库存不足，无法释放")
	}

	beforeAvailable := inventory.AvailableStock

	// 乐观锁释放：锁定库存减少，可用库存恢复
	result := tx.Model(&shop.ShopInventory{}).
		Where("sku_id = ? AND locked_stock >= ? AND version = ?", skuID, quantity, inventory.Version).
		Updates(map[string]interface{}{
			"locked_stock":    gorm.Expr("locked_stock - ?", quantity),
			"available_stock": gorm.Expr("available_stock + ?", quantity),
			"version":         gorm.Expr("version + 1"),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("释放库存失败，数据已被修改，请重试")
	}

	// 记录库存日志
	inventoryLog := shop.ShopInventoryLog{
		SkuID:     skuID,
		OrderNo:   orderNo,
		Type:      4, // 释放
		Quantity:  quantity,
		BeforeQty: beforeAvailable,
		AfterQty:  beforeAvailable + quantity,
	}
	return tx.Create(&inventoryLog).Error
}

// GetInventoryList 分页获取库存列表
func (s *ShopInventoryService) GetInventoryList(info shopReq.ShopInventorySearch) (list []shop.ShopInventory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopInventory{})

	// 需要关联SKU表时统一JOIN一次
	if info.SkuCode != "" || info.SpuName != "" {
		db = db.Joins("JOIN shop_sku ON shop_sku.id = shop_inventory.sku_id AND shop_sku.deleted_at IS NULL")
	}
	if info.SkuCode != "" {
		db = db.Where("shop_sku.sku_code LIKE ?", "%"+info.SkuCode+"%")
	}
	if info.SpuName != "" {
		db = db.Joins("JOIN shop_spu ON shop_spu.id = shop_sku.spu_id AND shop_spu.deleted_at IS NULL").
			Where("shop_spu.name LIKE ?", "%"+info.SpuName+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Preload("Sku").Find(&list).Error
	return
}

// GetInventoryLogList 分页获取库存操作日志列表
func (s *ShopInventoryService) GetInventoryLogList(info shopReq.ShopInventoryLogSearch) (list []shop.ShopInventoryLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopInventoryLog{})

	if info.SkuID > 0 {
		db = db.Where("sku_id = ?", info.SkuID)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id desc").Find(&list).Error
	return
}
