package shop

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"gorm.io/gorm"
)

type ShopOrderService struct{}

// generateOrderNo 生成订单号（使用纳秒+crypto/rand避免碰撞）
func generateOrderNo() string {
	now := time.Now()
	randNum, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("ORD%s%03d%06d", now.Format("20060102150405"), now.Nanosecond()/1000000, randNum.Int64())
}

// CreateOrder 创建订单（核心事务）
func (s *ShopOrderService) CreateOrder(userID uint, req shopReq.ShopOrderCreateReq) (order shop.ShopOrder, err error) {
	var inventoryService ShopInventoryService

	// 1. 验证收货地址属于该用户
	var address shop.ShopUserAddress
	if err = global.GVA_DB.Where("id = ? AND user_id = ?", req.AddressID, userID).First(&address).Error; err != nil {
		return order, fmt.Errorf("收货地址不存在")
	}

	// 2. 构建购买项列表：skuID -> quantity
	type buyItem struct {
		SkuID    uint
		Quantity int
	}
	var buyItems []buyItem

	if len(req.CartIDs) > 0 {
		// 从购物车获取
		var carts []shop.ShopCart
		if err = global.GVA_DB.Where("id IN ? AND user_id = ?", req.CartIDs, userID).Find(&carts).Error; err != nil {
			return order, fmt.Errorf("获取购物车失败")
		}
		if len(carts) == 0 {
			return order, fmt.Errorf("购物车为空")
		}
		for _, cart := range carts {
			buyItems = append(buyItems, buyItem{SkuID: cart.SkuID, Quantity: cart.Quantity})
		}
	} else if req.SkuID > 0 && req.Quantity > 0 {
		// 直接购买
		buyItems = append(buyItems, buyItem{SkuID: req.SkuID, Quantity: req.Quantity})
	} else {
		return order, fmt.Errorf("请选择商品")
	}

	// 3. 校验每个SKU的价格、状态，并构建订单项
	var orderItems []shop.ShopOrderItem
	var totalAmount float64

	for _, item := range buyItems {
		// 从数据库重新读取SKU当前价格（防止篡改）
		var sku shop.ShopSku
		if err = global.GVA_DB.Where("id = ?", item.SkuID).First(&sku).Error; err != nil {
			return order, fmt.Errorf("商品SKU不存在: %d", item.SkuID)
		}
		if sku.Status == nil || *sku.Status != 1 {
			return order, fmt.Errorf("商品SKU已下架: %s", sku.Name)
		}

		// 校验SPU状态
		var spu shop.ShopSpu
		if err = global.GVA_DB.Where("id = ?", sku.SpuID).First(&spu).Error; err != nil {
			return order, fmt.Errorf("商品不存在: %d", sku.SpuID)
		}
		if spu.Status != 1 {
			return order, fmt.Errorf("商品已下架: %s", spu.Name)
		}

		itemTotal := sku.Price * float64(item.Quantity)
		totalAmount += itemTotal

		orderItems = append(orderItems, shop.ShopOrderItem{
			SpuID:      sku.SpuID,
			SkuID:      sku.ID,
			SpuName:    spu.Name,
			SkuName:    sku.Name,
			SkuImage:   sku.Image,
			SpecData:   sku.SpecData,
			Price:      sku.Price,
			Quantity:   item.Quantity,
			TotalPrice: itemTotal,
		})
	}

	// 4. 生成订单号
	orderNo := generateOrderNo()

	// 4a. 优惠券处理
	var discountAmount float64
	var couponID *uint
	var couponCode string
	if req.CouponCode != "" {
		var couponService ShopCouponService
		discount, cID, cErr := couponService.CalcDiscountByCode(userID, req.CouponCode, totalAmount)
		if cErr != nil {
			return order, fmt.Errorf("优惠券无效: %v", cErr)
		}
		discountAmount = discount
		couponID = &cID
		couponCode = req.CouponCode
	}

	// 5. 拼装收货地址
	receiverAddress := address.Province + address.City + address.District + address.DetailAddr

	// 6. 设置过期时间
	expireTime := time.Now().Add(30 * time.Minute)

	// 7. 开启事务
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 锁定库存
		for _, item := range orderItems {
			if lockErr := inventoryService.LockStock(tx, item.SkuID, item.Quantity, orderNo); lockErr != nil {
				return fmt.Errorf("库存不足")
			}
		}

		// 创建订单
		order = shop.ShopOrder{
			OrderNo:         orderNo,
			UserID:          userID,
			Status:          0,
			TotalAmount:     totalAmount,
			FreightAmount:   0,
			DiscountAmount:  discountAmount,
			PayAmount:       totalAmount - discountAmount,
			CouponID:        couponID,
			CouponCode:      couponCode,
			ReceiverName:    address.ReceiverName,
			ReceiverPhone:   address.Phone,
			ReceiverAddress: receiverAddress,
			Remark:          req.Remark,
			ExpireTime:      &expireTime,
		}
		if createErr := tx.Create(&order).Error; createErr != nil {
			return createErr
		}

		// 创建订单项
		for i := range orderItems {
			orderItems[i].OrderNo = orderNo
		}
		if createErr := tx.Create(&orderItems).Error; createErr != nil {
			return createErr
		}

		// 锁定优惠券
		if couponID != nil {
			var couponService ShopCouponService
			if couponErr := couponService.UseCoupon(tx, userID, *couponID, orderNo); couponErr != nil {
				return fmt.Errorf("优惠券使用失败: %v", couponErr)
			}
		}

		// 硬删除已购买的购物车项（避免唯一索引与软删除冲突）
		if len(req.CartIDs) > 0 {
			if delErr := tx.Unscoped().Where("id IN ? AND user_id = ?", req.CartIDs, userID).Delete(&shop.ShopCart{}).Error; delErr != nil {
				return delErr
			}
		}

		// 创建订单日志
		orderLog := shop.ShopOrderLog{
			OrderNo:    orderNo,
			FromStatus: -1,
			ToStatus:   0,
			OperatorID: userID,
			Remark:     "创建订单",
		}
		if logErr := tx.Create(&orderLog).Error; logErr != nil {
			return logErr
		}

		return nil
	})

	if err != nil {
		return shop.ShopOrder{}, err
	}

	order.Items = orderItems

	// 发送通知和邮件（异步）
	go func() {
		var notificationService ShopNotificationService
		var emailService ShopEmailService
		notificationService.CreateNotification(userID, 1, "订单创建成功", "您的订单 "+orderNo+" 已创建，请尽快支付", "order", orderNo)
		emailService.SendOrderCreatedEmail(userID, orderNo, order.PayAmount)
	}()

	return order, nil
}

// GetOrderList 管理端获取订单列表（支持筛选，预加载Items）
func (s *ShopOrderService) GetOrderList(info shopReq.ShopOrderSearch) (list []shop.ShopOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopOrder{})

	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	if info.OrderNo != "" {
		db = db.Where("order_no = ?", info.OrderNo)
	}
	if info.UserID != nil {
		db = db.Where("user_id = ?", *info.UserID)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("created_at desc").Preload("Items").Find(&list).Error
	return
}

// GetUserOrderList 用户端获取订单列表
func (s *ShopOrderService) GetUserOrderList(userID uint, info shopReq.ShopOrderSearch) (list []shop.ShopOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopOrder{}).Where("user_id = ?", userID)

	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	if info.OrderNo != "" {
		db = db.Where("order_no = ?", info.OrderNo)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("created_at desc").Preload("Items").Find(&list).Error
	return
}

// GetOrder 根据订单号获取订单详情（预加载Items）
func (s *ShopOrderService) GetOrder(orderNo string) (order shop.ShopOrder, err error) {
	err = global.GVA_DB.Where("order_no = ?", orderNo).Preload("Items").First(&order).Error
	return
}

// GetUserOrder 根据订单号获取用户订单详情（校验归属）
func (s *ShopOrderService) GetUserOrder(orderNo string, userID uint) (order shop.ShopOrder, err error) {
	err = global.GVA_DB.Where("order_no = ? AND user_id = ?", orderNo, userID).Preload("Items").First(&order).Error
	return
}

// CancelOrder 取消订单（用户/通用）
func (s *ShopOrderService) CancelOrder(orderNo string, userID uint) error {
	var inventoryService ShopInventoryService

	var order shop.ShopOrder
	if err := global.GVA_DB.Where("order_no = ?", orderNo).First(&order).Error; err != nil {
		return fmt.Errorf("订单不存在")
	}

	// 验证订单状态
	if order.Status != 0 {
		return fmt.Errorf("当前订单状态不允许取消")
	}

	// 验证所有权（userID > 0 时校验）
	if userID > 0 && order.UserID != userID {
		return fmt.Errorf("无权操作此订单")
	}

	// 获取订单项用于释放库存
	var items []shop.ShopOrderItem
	if err := global.GVA_DB.Where("order_no = ?", orderNo).Find(&items).Error; err != nil {
		return fmt.Errorf("获取订单项失败")
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 释放库存
		for _, item := range items {
			if err := inventoryService.ReleaseStock(tx, item.SkuID, item.Quantity, orderNo); err != nil {
				return fmt.Errorf("释放库存失败: %v", err)
			}
		}

		// 更新订单状态
		now := time.Now()
		if err := tx.Model(&shop.ShopOrder{}).Where("order_no = ?", orderNo).Updates(map[string]interface{}{
			"status":      5,
			"cancel_time": now,
		}).Error; err != nil {
			return err
		}

		// 创建订单日志
		operatorID := userID
		orderLog := shop.ShopOrderLog{
			OrderNo:    orderNo,
			FromStatus: 0,
			ToStatus:   5,
			OperatorID: operatorID,
			Remark:     "取消订单",
		}
		if logErr := tx.Create(&orderLog).Error; logErr != nil {
			return logErr
		}
		// 退还优惠券
		if order.CouponID != nil {
			var couponService ShopCouponService
			_ = couponService.ReturnCoupon(tx, userID, orderNo)
		}
		return nil
	})
}

// ShipOrder 发货
func (s *ShopOrderService) ShipOrder(req shopReq.ShopOrderShipReq, operatorID uint) error {
	var inventoryService ShopInventoryService

	var order shop.ShopOrder
	if err := global.GVA_DB.Where("order_no = ?", req.OrderNo).First(&order).Error; err != nil {
		return fmt.Errorf("订单不存在")
	}

	if order.Status != 1 {
		return fmt.Errorf("当前订单状态不允许发货")
	}

	// 获取订单项用于扣减库存
	var items []shop.ShopOrderItem
	if err := global.GVA_DB.Where("order_no = ?", req.OrderNo).Find(&items).Error; err != nil {
		return fmt.Errorf("获取订单项失败")
	}

	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 扣减库存（锁定库存 → 已售出）
		for _, item := range items {
			if err := inventoryService.DeductStock(tx, item.SkuID, item.Quantity, req.OrderNo); err != nil {
				return fmt.Errorf("扣减库存失败: %v", err)
			}
		}

		// 更新订单状态
		now := time.Now()
		if err := tx.Model(&shop.ShopOrder{}).Where("order_no = ?", req.OrderNo).Updates(map[string]interface{}{
			"status":       2,
			"ship_time":    now,
			"ship_company": req.ShipCompany,
			"ship_no":      req.ShipNo,
		}).Error; err != nil {
			return err
		}

		// 创建订单日志
		orderLog := shop.ShopOrderLog{
			OrderNo:    req.OrderNo,
			FromStatus: 1,
			ToStatus:   2,
			OperatorID: operatorID,
			Remark:     fmt.Sprintf("发货 物流公司:%s 物流单号:%s", req.ShipCompany, req.ShipNo),
		}
		return tx.Create(&orderLog).Error
	})
	if err != nil {
		return err
	}

	// 发货后钩子（异步）
	go func() {
		var notificationService ShopNotificationService
		var emailService ShopEmailService
		var logisticsService ShopLogisticsService
		notificationService.CreateNotification(order.UserID, 2, "订单已发货", "您的订单 "+req.OrderNo+" 已发货，物流公司: "+req.ShipCompany+"，单号: "+req.ShipNo, "order", req.OrderNo)
		emailService.SendOrderShippedEmail(order.UserID, req.OrderNo, req.ShipCompany, req.ShipNo)
		logisticsService.AddTrace(req.OrderNo, req.ShipNo, req.ShipCompany, "已揽收", "包裹已由 "+req.ShipCompany+" 揽收")
	}()

	return nil
}

// ConfirmReceive 确认收货
func (s *ShopOrderService) ConfirmReceive(orderNo string, userID uint) error {
	var order shop.ShopOrder
	if err := global.GVA_DB.Where("order_no = ?", orderNo).First(&order).Error; err != nil {
		return fmt.Errorf("订单不存在")
	}

	if order.Status != 2 {
		return fmt.Errorf("当前订单状态不允许确认收货")
	}

	if order.UserID != userID {
		return fmt.Errorf("无权操作此订单")
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		now := time.Now()
		if err := tx.Model(&shop.ShopOrder{}).Where("order_no = ?", orderNo).Updates(map[string]interface{}{
			"status":       3,
			"receive_time": now,
		}).Error; err != nil {
			return err
		}

		orderLog := shop.ShopOrderLog{
			OrderNo:    orderNo,
			FromStatus: 2,
			ToStatus:   3,
			OperatorID: userID,
			Remark:     "确认收货",
		}
		return tx.Create(&orderLog).Error
	})
}

// AdminCancelOrder 管理员取消订单（支持取消待付款和已付款订单）
func (s *ShopOrderService) AdminCancelOrder(orderNo string, operatorID uint) error {
	var inventoryService ShopInventoryService
	var paymentService ShopPaymentService

	var order shop.ShopOrder
	if err := global.GVA_DB.Where("order_no = ?", orderNo).First(&order).Error; err != nil {
		return fmt.Errorf("订单不存在")
	}

	// 只允许取消 待付款(0) 和 已付款待发货(1) 的订单
	if order.Status != 0 && order.Status != 1 {
		return fmt.Errorf("当前订单状态不允许取消")
	}

	// 获取订单项用于释放库存
	var items []shop.ShopOrderItem
	if err := global.GVA_DB.Where("order_no = ?", orderNo).Find(&items).Error; err != nil {
		return fmt.Errorf("获取订单项失败")
	}

	originalStatus := order.Status

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if originalStatus == 0 {
			// 待付款：释放锁定库存
			for _, item := range items {
				if err := inventoryService.ReleaseStock(tx, item.SkuID, item.Quantity, orderNo); err != nil {
					return fmt.Errorf("释放库存失败: %v", err)
				}
			}
		} else if originalStatus == 1 {
			// 已付款：释放锁定库存 + 退款到余额
			for _, item := range items {
				if err := inventoryService.ReleaseStock(tx, item.SkuID, item.Quantity, orderNo); err != nil {
					return fmt.Errorf("释放库存失败: %v", err)
				}
			}
			// 退款到余额
			if err := paymentService.RefundToBalance(tx, order.UserID, order.PayAmount, orderNo); err != nil {
				return fmt.Errorf("退款失败: %v", err)
			}
		}

		now := time.Now()
		if err := tx.Model(&shop.ShopOrder{}).Where("order_no = ?", orderNo).Updates(map[string]interface{}{
			"status":      5,
			"cancel_time": now,
		}).Error; err != nil {
			return err
		}

		remark := "管理员取消订单"
		if originalStatus == 1 {
			remark = "管理员取消订单，已退款至余额"
		}

		orderLog := shop.ShopOrderLog{
			OrderNo:    orderNo,
			FromStatus: originalStatus,
			ToStatus:   5,
			OperatorID: operatorID,
			Remark:     remark,
		}
		if logErr := tx.Create(&orderLog).Error; logErr != nil {
			return logErr
		}
		// 退还优惠券
		if order.CouponID != nil {
			var couponService ShopCouponService
			_ = couponService.ReturnCoupon(tx, 0, orderNo)
		}
		return nil
	})
}

// GetOrderLogs 获取订单操作日志
func (s *ShopOrderService) GetOrderLogs(orderNo string) (logs []shop.ShopOrderLog, err error) {
	err = global.GVA_DB.Where("order_no = ?", orderNo).Order("created_at asc").Find(&logs).Error
	return
}

// AutoCancelExpiredOrders 自动取消过期未支付订单（status=0 且 expire_time < now）
func (s *ShopOrderService) AutoCancelExpiredOrders() {
	var inventoryService ShopInventoryService

	var orders []shop.ShopOrder
	now := time.Now()
	if err := global.GVA_DB.Where("status = 0 AND expire_time < ?", now).Find(&orders).Error; err != nil {
		return
	}

	for _, order := range orders {
		_ = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			// 事务内重新检查订单状态（防止与支付并发竞态）
			var freshOrder shop.ShopOrder
			if err := tx.Where("order_no = ? AND status = 0", order.OrderNo).First(&freshOrder).Error; err != nil {
				return err // 状态已变更，跳过
			}

			// 获取订单项
			var items []shop.ShopOrderItem
			if err := tx.Where("order_no = ?", order.OrderNo).Find(&items).Error; err != nil {
				return err
			}

			// 释放库存
			for _, item := range items {
				if err := inventoryService.ReleaseStock(tx, item.SkuID, item.Quantity, order.OrderNo); err != nil {
					return err
				}
			}

			// 更新订单状态
			cancelTime := time.Now()
			if err := tx.Model(&shop.ShopOrder{}).Where("order_no = ? AND status = 0", order.OrderNo).Updates(map[string]interface{}{
				"status":      5,
				"cancel_time": cancelTime,
			}).Error; err != nil {
				return err
			}

			// 创建订单日志
			orderLog := shop.ShopOrderLog{
				OrderNo:    order.OrderNo,
				FromStatus: 0,
				ToStatus:   5,
				OperatorID: 0,
				Remark:     "系统自动取消：订单超时未支付",
			}
			return tx.Create(&orderLog).Error
		})
	}
}

// AutoConfirmReceived 自动确认收货（status=2 且 ship_time < now-15天）
func (s *ShopOrderService) AutoConfirmReceived() {
	var orders []shop.ShopOrder
	deadline := time.Now().Add(-15 * 24 * time.Hour)
	if err := global.GVA_DB.Where("status = 2 AND ship_time < ?", deadline).Find(&orders).Error; err != nil {
		return
	}

	for _, order := range orders {
		_ = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			// 事务内重新检查状态
			var freshOrder shop.ShopOrder
			if err := tx.Where("order_no = ? AND status = 2", order.OrderNo).First(&freshOrder).Error; err != nil {
				return err
			}

			now := time.Now()
			if err := tx.Model(&shop.ShopOrder{}).Where("order_no = ? AND status = 2", order.OrderNo).Updates(map[string]interface{}{
				"status":       3,
				"receive_time": now,
			}).Error; err != nil {
				return err
			}

			orderLog := shop.ShopOrderLog{
				OrderNo:    order.OrderNo,
				FromStatus: 2,
				ToStatus:   3,
				OperatorID: 0,
				Remark:     "系统自动确认收货：发货超过15天",
			}
			return tx.Create(&orderLog).Error
		})
	}
}

// AutoComplete 自动完成订单（status=3 且 receive_time < now-7天）
func (s *ShopOrderService) AutoComplete() {
	var orders []shop.ShopOrder
	deadline := time.Now().Add(-7 * 24 * time.Hour)
	if err := global.GVA_DB.Where("status = 3 AND receive_time < ?", deadline).Find(&orders).Error; err != nil {
		return
	}

	for _, order := range orders {
		txErr := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			// 事务内重新检查状态
			var freshOrder shop.ShopOrder
			if err := tx.Where("order_no = ? AND status = 3", order.OrderNo).First(&freshOrder).Error; err != nil {
				return err
			}

			now := time.Now()
			if err := tx.Model(&shop.ShopOrder{}).Where("order_no = ? AND status = 3", order.OrderNo).Updates(map[string]interface{}{
				"status":        4,
				"complete_time": now,
			}).Error; err != nil {
				return err
			}

			orderLog := shop.ShopOrderLog{
				OrderNo:    order.OrderNo,
				FromStatus: 3,
				ToStatus:   4,
				OperatorID: 0,
				Remark:     "系统自动完成：确认收货超过7天",
			}
			return tx.Create(&orderLog).Error
		})
		if txErr == nil {
			go func(o shop.ShopOrder) {
				var notificationService ShopNotificationService
				var emailService ShopEmailService
				notificationService.CreateNotification(o.UserID, 1, "订单已完成", "您的订单 "+o.OrderNo+" 已自动完成", "order", o.OrderNo)
				emailService.SendOrderCompletedEmail(o.UserID, o.OrderNo)
			}(order)
		}
	}
}
