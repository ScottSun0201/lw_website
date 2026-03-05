package shop

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	shopRes "github.com/flipped-aurora/gin-vue-admin/server/model/shop/response"
	"gorm.io/gorm"
)

type ShopCartService struct{}

// AddToCart 添加商品到购物车
func (s *ShopCartService) AddToCart(userID uint, req shopReq.ShopCartAddReq) error {
	// 验证SKU是否存在且启用
	var sku shop.ShopSku
	if err := global.GVA_DB.Where("id = ?", req.SkuID).First(&sku).Error; err != nil {
		return errors.New("商品规格不存在")
	}
	if sku.Status == nil || *sku.Status != 1 {
		return errors.New("商品规格已下架")
	}

	// 查询库存
	var inventory shop.ShopInventory
	if err := global.GVA_DB.Where("sku_id = ?", req.SkuID).First(&inventory).Error; err != nil {
		return errors.New("库存信息不存在")
	}
	if req.Quantity > inventory.AvailableStock {
		return fmt.Errorf("库存不足，当前可用库存为%d", inventory.AvailableStock)
	}

	// 检查是否已存在相同SKU（user_id + sku_id 唯一，包含软删除记录）
	var existingCart shop.ShopCart
	err := global.GVA_DB.Unscoped().Where("user_id = ? AND sku_id = ?", userID, req.SkuID).First(&existingCart).Error

	if err == nil {
		if existingCart.DeletedAt.Valid {
			// 软删除的记录：恢复并设置新数量
			return global.GVA_DB.Unscoped().Model(&existingCart).Updates(map[string]interface{}{
				"quantity":   req.Quantity,
				"deleted_at": nil,
				"selected":   1,
			}).Error
		}
		// 未删除的记录：合并数量
		newQuantity := existingCart.Quantity + req.Quantity
		if newQuantity > 99 {
			return errors.New("单个商品数量不能超过99")
		}
		if newQuantity > inventory.AvailableStock {
			return fmt.Errorf("库存不足，当前可用库存为%d", inventory.AvailableStock)
		}
		return global.GVA_DB.Model(&existingCart).Update("quantity", newQuantity).Error
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 不存在，检查购物车种类数量
	var cartCount int64
	if err := global.GVA_DB.Model(&shop.ShopCart{}).Where("user_id = ?", userID).Count(&cartCount).Error; err != nil {
		return err
	}
	if cartCount >= 99 {
		return errors.New("购物车商品种类已达上限(99种)")
	}

	// 检查添加数量
	if req.Quantity > 99 {
		return errors.New("单个商品数量不能超过99")
	}

	// 新增购物车记录
	cart := shop.ShopCart{
		UserID:   userID,
		SkuID:    req.SkuID,
		Quantity: req.Quantity,
	}
	return global.GVA_DB.Create(&cart).Error
}

// UpdateCartQuantity 更新购物车商品数量
func (s *ShopCartService) UpdateCartQuantity(userID uint, req shopReq.ShopCartUpdateReq) error {
	// 验证购物车记录归属
	var cart shop.ShopCart
	if err := global.GVA_DB.Where("id = ? AND user_id = ?", req.ID, userID).First(&cart).Error; err != nil {
		return errors.New("购物车记录不存在或无权操作")
	}

	// 验证SKU库存是否充足
	var inventory shop.ShopInventory
	if err := global.GVA_DB.Where("sku_id = ?", cart.SkuID).First(&inventory).Error; err != nil {
		return errors.New("库存信息不存在")
	}
	if inventory.AvailableStock < req.Quantity {
		return fmt.Errorf("库存不足，当前可用库存为%d", inventory.AvailableStock)
	}

	return global.GVA_DB.Model(&cart).Update("quantity", req.Quantity).Error
}

// RemoveFromCart 批量删除购物车商品
func (s *ShopCartService) RemoveFromCart(userID uint, req shopReq.ShopCartDeleteReq) error {
	// 验证所有记录都属于该用户
	var count int64
	global.GVA_DB.Model(&shop.ShopCart{}).Where("id IN ? AND user_id = ?", req.IDs, userID).Count(&count)
	if count != int64(len(req.IDs)) {
		return errors.New("存在无权操作的购物车记录")
	}

	return global.GVA_DB.Unscoped().Where("id IN ? AND user_id = ?", req.IDs, userID).Delete(&shop.ShopCart{}).Error
}

// GetCartList 获取购物车列表
func (s *ShopCartService) GetCartList(userID uint) (list []shopRes.ShopCartItem, err error) {
	var carts []shop.ShopCart
	err = global.GVA_DB.Where("user_id = ?", userID).
		Preload("Sku").
		Order("created_at DESC").
		Find(&carts).Error
	if err != nil {
		return nil, err
	}

	list = s.enrichCartItems(carts)
	return
}

// GetCartCount 获取购物车商品数量（用于header角标）
func (s *ShopCartService) GetCartCount(userID uint) (count int64, err error) {
	err = global.GVA_DB.Model(&shop.ShopCart{}).Where("user_id = ?", userID).Count(&count).Error
	return
}

// GetSelectedCartItems 根据ID列表获取购物车商品详情（结算页面使用）
func (s *ShopCartService) GetSelectedCartItems(userID uint, cartIDs []uint) (list []shopRes.ShopCartItem, err error) {
	var carts []shop.ShopCart
	err = global.GVA_DB.Where("id IN ? AND user_id = ?", cartIDs, userID).
		Preload("Sku").
		Find(&carts).Error
	if err != nil {
		return nil, err
	}

	list = s.enrichCartItems(carts)
	return
}

// enrichCartItems 填充购物车商品详情信息
func (s *ShopCartService) enrichCartItems(carts []shop.ShopCart) []shopRes.ShopCartItem {
	if len(carts) == 0 {
		return []shopRes.ShopCartItem{}
	}

	// 收集所有SPU ID
	spuIDSet := make(map[uint]struct{})
	for _, cart := range carts {
		spuIDSet[cart.Sku.SpuID] = struct{}{}
	}
	spuIDs := make([]uint, 0, len(spuIDSet))
	for id := range spuIDSet {
		spuIDs = append(spuIDs, id)
	}

	// 批量查询SPU信息
	var spuList []shop.ShopSpu
	global.GVA_DB.Where("id IN ?", spuIDs).Find(&spuList)
	spuMap := make(map[uint]shop.ShopSpu)
	for _, spu := range spuList {
		spuMap[spu.ID] = spu
	}

	// 收集所有SKU ID，批量查询库存
	skuIDs := make([]uint, 0, len(carts))
	for _, cart := range carts {
		skuIDs = append(skuIDs, cart.SkuID)
	}
	var inventories []shop.ShopInventory
	global.GVA_DB.Where("sku_id IN ?", skuIDs).Find(&inventories)
	inventoryMap := make(map[uint]shop.ShopInventory)
	for _, inv := range inventories {
		inventoryMap[inv.SkuID] = inv
	}

	// 组装结果
	list := make([]shopRes.ShopCartItem, 0, len(carts))
	for _, cart := range carts {
		item := shopRes.ShopCartItem{
			ShopCart:  cart,
			SkuName:   cart.Sku.Name,
			SkuImage:  cart.Sku.Image,
			SkuPrice:  cart.Sku.Price,
			SpecData:  cart.Sku.SpecData,
			SkuStatus: 0,
		}

		// SKU状态
		if cart.Sku.Status != nil {
			item.SkuStatus = *cart.Sku.Status
		}

		// SPU信息
		if spu, ok := spuMap[cart.Sku.SpuID]; ok {
			item.SpuName = spu.Name
			item.SpuImage = spu.MainImage
			item.SpuStatus = spu.Status
		}

		// 库存信息
		if inv, ok := inventoryMap[cart.SkuID]; ok {
			item.Stock = inv.AvailableStock
		}

		// 判断是否失效：SPU状态不为1(上架) 或 SKU状态不为1(启用)
		if item.SpuStatus != 1 || item.SkuStatus != 1 {
			item.Invalid = true
		}

		list = append(list, item)
	}

	return list
}
