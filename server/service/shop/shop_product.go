package shop

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"gorm.io/gorm"
)

type ShopProductService struct{}

// CreateSpu 创建商品SPU（含SKU和库存初始化）
func (s *ShopProductService) CreateSpu(spu *shop.ShopSpu) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 保存SKU列表，先从SPU中取出
		skus := spu.Skus
		spu.Skus = nil

		// 创建SPU
		if err := tx.Create(spu).Error; err != nil {
			return err
		}

		// 批量创建SKU并初始化库存
		if len(skus) > 0 {
			for i := range skus {
				skus[i].SpuID = spu.ID
			}
			if err := tx.Create(&skus).Error; err != nil {
				return err
			}

			// 为每个SKU初始化库存
			var inventoryService ShopInventoryService
			for _, sku := range skus {
				if err := inventoryService.InitInventory(tx, sku.ID, sku.Stock); err != nil {
					return err
				}
			}
		}

		// 将SKU重新赋值回SPU
		spu.Skus = skus
		return nil
	})
}

// UpdateSpu 更新商品SPU（含SKU差异更新）
func (s *ShopProductService) UpdateSpu(spu *shop.ShopSpu) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 取出提交的SKU列表
		newSkus := spu.Skus
		spu.Skus = nil

		// 更新SPU基本信息（使用Select指定字段，避免覆盖sales_count等不应被编辑覆盖的字段）
		if err := tx.Model(&shop.ShopSpu{}).Where("id = ?", spu.ID).Updates(map[string]interface{}{
			"name":        spu.Name,
			"sub_title":   spu.SubTitle,
			"category_id": spu.CategoryID,
			"brand_id":    spu.BrandID,
			"main_image":  spu.MainImage,
			"images":      spu.Images,
			"detail":      spu.Detail,
			"status":      spu.Status,
			"sort":        spu.Sort,
		}).Error; err != nil {
			return err
		}

		// 查询当前数据库中的SKU列表
		var existingSkus []shop.ShopSku
		if err := tx.Where("spu_id = ?", spu.ID).Find(&existingSkus).Error; err != nil {
			return err
		}

		// 构建已有SKU的ID映射
		existingMap := make(map[uint]bool)
		for _, sku := range existingSkus {
			existingMap[sku.ID] = true
		}

		// 构建提交SKU的ID映射
		submittedMap := make(map[uint]bool)
		for _, sku := range newSkus {
			if sku.ID > 0 {
				submittedMap[sku.ID] = true
			}
		}

		// 删除不再存在的SKU（软删除）
		for _, sku := range existingSkus {
			if !submittedMap[sku.ID] {
				if err := tx.Delete(&shop.ShopSku{}, sku.ID).Error; err != nil {
					return err
				}
			}
		}

		// 新增或更新SKU
		var inventoryService ShopInventoryService
		for i := range newSkus {
			newSkus[i].SpuID = spu.ID
			if newSkus[i].ID > 0 && existingMap[newSkus[i].ID] {
				// 更新已有SKU（使用Select指定字段）
				if err := tx.Model(&shop.ShopSku{}).Where("id = ?", newSkus[i].ID).Updates(map[string]interface{}{
					"spu_id":       newSkus[i].SpuID,
					"sku_code":     newSkus[i].SkuCode,
					"name":         newSkus[i].Name,
					"spec_data":    newSkus[i].SpecData,
					"price":        newSkus[i].Price,
					"market_price": newSkus[i].MarketPrice,
					"cost_price":   newSkus[i].CostPrice,
					"image":        newSkus[i].Image,
					"weight":       newSkus[i].Weight,
					"status":       newSkus[i].Status,
				}).Error; err != nil {
					return err
				}
			} else {
				// 新增SKU
				newSkus[i].ID = 0
				if err := tx.Create(&newSkus[i]).Error; err != nil {
					return err
				}
				// 为新SKU初始化库存
				if err := inventoryService.InitInventory(tx, newSkus[i].ID, 0); err != nil {
					return err
				}
			}
		}

		spu.Skus = newSkus
		return nil
	})
}

// DeleteSpu 删除商品SPU及关联SKU（软删除）
func (s *ShopProductService) DeleteSpu(ID string) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 软删除关联SKU
		if err := tx.Where("spu_id = ?", ID).Delete(&shop.ShopSku{}).Error; err != nil {
			return err
		}
		// 软删除SPU
		if err := tx.Delete(&shop.ShopSpu{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetSpu 根据ID获取商品SPU详情
func (s *ShopProductService) GetSpu(ID string) (spu shop.ShopSpu, err error) {
	err = global.GVA_DB.Preload("Skus").Preload("Category").Preload("Brand").
		Where("id = ?", ID).First(&spu).Error
	if err != nil {
		return
	}

	// 填充每个SKU的库存信息
	s.fillSkuStock(&spu)

	// 计算价格区间
	s.calcPriceRange(&spu)

	return
}

// GetSpuList 分页获取商品SPU列表（管理端）
func (s *ShopProductService) GetSpuList(info shopReq.ShopSpuSearch) (list []shop.ShopSpu, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopSpu{})

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.CategoryID != nil {
		db = db.Where("category_id = ?", *info.CategoryID)
	}
	if info.BrandID != nil {
		db = db.Where("brand_id = ?", *info.BrandID)
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

	err = db.Order("sort asc, id desc").Preload("Category").Preload("Brand").Preload("Skus").Find(&list).Error
	if err != nil {
		return
	}

	// 为每个SPU计算价格区间
	for i := range list {
		s.calcPriceRange(&list[i])
	}

	return
}

// UpdateSpuStatus 更新商品上下架状态（状态机）
func (s *ShopProductService) UpdateSpuStatus(req shopReq.ShopSpuStatusReq) error {
	var spu shop.ShopSpu
	if err := global.GVA_DB.Preload("Skus").Where("id = ?", req.ID).First(&spu).Error; err != nil {
		return err
	}

	// 状态机校验: draft(0)→online(1), online(1)⇄offline(2)
	switch spu.Status {
	case 0: // 草稿
		if req.Status != 1 {
			return errors.New("草稿状态只能上架")
		}
		// 上架前校验：必须有启用的SKU且有库存
		if err := s.checkOnlineCondition(&spu); err != nil {
			return err
		}
	case 1: // 上架
		if req.Status != 2 {
			return errors.New("上架状态只能下架")
		}
	case 2: // 下架
		if req.Status != 1 {
			return errors.New("下架状态只能上架")
		}
		// 重新上架也需要校验
		if err := s.checkOnlineCondition(&spu); err != nil {
			return err
		}
	default:
		return errors.New("无效的当前状态")
	}

	return global.GVA_DB.Model(&shop.ShopSpu{}).Where("id = ?", req.ID).Update("status", req.Status).Error
}

// checkOnlineCondition 校验上架条件
func (s *ShopProductService) checkOnlineCondition(spu *shop.ShopSpu) error {
	// 检查是否有启用的SKU
	hasEnabledSku := false
	var enabledSkuIDs []uint
	for _, sku := range spu.Skus {
		if sku.Status != nil && *sku.Status == 1 {
			hasEnabledSku = true
			enabledSkuIDs = append(enabledSkuIDs, sku.ID)
		}
	}
	if !hasEnabledSku {
		return errors.New("上架失败：没有启用的SKU")
	}

	// 检查启用的SKU是否有库存
	var count int64
	global.GVA_DB.Model(&shop.ShopInventory{}).
		Where("sku_id IN ? AND available_stock > 0", enabledSkuIDs).
		Count(&count)
	if count == 0 {
		return errors.New("上架失败：启用的SKU均无可用库存")
	}

	return nil
}

// GetClientProductList 获取客户端商品列表（仅上架商品）
func (s *ShopProductService) GetClientProductList(info shopReq.ShopSpuSearch) (list []shop.ShopSpu, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopSpu{}).Where("status = ?", 1)

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.CategoryID != nil {
		db = db.Where("category_id = ?", *info.CategoryID)
	}
	if info.BrandID != nil {
		db = db.Where("brand_id = ?", *info.BrandID)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("sort asc, id desc").Preload("Category").Preload("Brand").Preload("Skus").Find(&list).Error
	if err != nil {
		return
	}

	// 计算价格区间
	for i := range list {
		s.calcPriceRange(&list[i])
	}

	return
}

// GetClientProductDetail 获取客户端商品详情
func (s *ShopProductService) GetClientProductDetail(ID string) (spu shop.ShopSpu, err error) {
	err = global.GVA_DB.Preload("Skus").Preload("Category").Preload("Brand").
		Where("id = ? AND status = ?", ID, 1).First(&spu).Error
	if err != nil {
		return
	}

	// 填充SKU库存
	s.fillSkuStock(&spu)

	// 计算价格区间
	s.calcPriceRange(&spu)

	return
}

// fillSkuStock 填充SKU的库存信息
func (s *ShopProductService) fillSkuStock(spu *shop.ShopSpu) {
	if len(spu.Skus) == 0 {
		return
	}

	skuIDs := make([]uint, len(spu.Skus))
	for i, sku := range spu.Skus {
		skuIDs[i] = sku.ID
	}

	var inventories []shop.ShopInventory
	global.GVA_DB.Where("sku_id IN ?", skuIDs).Find(&inventories)

	// 构建库存映射
	stockMap := make(map[uint]int)
	for _, inv := range inventories {
		stockMap[inv.SkuID] = inv.AvailableStock
	}

	// 赋值到SKU
	for i := range spu.Skus {
		spu.Skus[i].Stock = stockMap[spu.Skus[i].ID]
	}
}

// calcPriceRange 计算商品价格区间
func (s *ShopProductService) calcPriceRange(spu *shop.ShopSpu) {
	if len(spu.Skus) == 0 {
		spu.PriceRange = "¥0"
		return
	}

	minPrice := spu.Skus[0].Price
	maxPrice := spu.Skus[0].Price

	for _, sku := range spu.Skus {
		if sku.Price < minPrice {
			minPrice = sku.Price
		}
		if sku.Price > maxPrice {
			maxPrice = sku.Price
		}
	}

	if minPrice == maxPrice {
		spu.PriceRange = fmt.Sprintf("¥%.2f", minPrice)
	} else {
		spu.PriceRange = fmt.Sprintf("¥%.2f - ¥%.2f", minPrice, maxPrice)
	}
}
