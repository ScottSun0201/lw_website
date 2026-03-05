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

type ShopRefundService struct{}

func generateRefundNo() string {
	now := time.Now()
	randNum, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("REF%s%06d", now.Format("20060102150405"), randNum.Int64())
}

func (s *ShopRefundService) CreateRefund(userID uint, req shopReq.ShopRefundCreateReq) (refund shop.ShopRefund, err error) {
	var order shop.ShopOrder
	if err = global.GVA_DB.Where("order_no = ? AND user_id = ?", req.OrderNo, userID).First(&order).Error; err != nil {
		return refund, fmt.Errorf("订单不存在")
	}
	if order.Status != 1 && order.Status != 2 && order.Status != 3 {
		return refund, fmt.Errorf("当前订单状态不允许退款")
	}
	var activeCount int64
	global.GVA_DB.Model(&shop.ShopRefund{}).Where("order_no = ? AND status IN (0, 1, 3)", req.OrderNo).Count(&activeCount)
	if activeCount > 0 {
		return refund, fmt.Errorf("该订单已有进行中的退款")
	}
	if req.Amount > order.PayAmount {
		return refund, fmt.Errorf("退款金额不能超过实付金额")
	}

	refundNo := generateRefundNo()
	refund = shop.ShopRefund{
		RefundNo: refundNo, OrderNo: req.OrderNo, UserID: userID,
		Type: req.Type, Reason: req.Reason, Description: req.Description,
		Images: req.Images, Amount: req.Amount, Status: 0,
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&refund).Error; err != nil {
			return err
		}
		if err := tx.Model(&shop.ShopOrder{}).Where("order_no = ?", req.OrderNo).Update("status", 6).Error; err != nil {
			return err
		}
		log := shop.ShopRefundLog{RefundNo: refundNo, FromStatus: -1, ToStatus: 0, OperatorID: userID, Remark: "申请退款"}
		return tx.Create(&log).Error
	})

	if err == nil {
		var notificationService ShopNotificationService
		go notificationService.CreateAdminNotification(2, "新退款申请", fmt.Sprintf("订单 %s 申请退款 ¥%.2f", req.OrderNo, req.Amount), "refund", refundNo)
	}
	return
}

func (s *ShopRefundService) AuditRefund(req shopReq.ShopRefundAuditReq, operatorID uint) error {
	var refund shop.ShopRefund
	if err := global.GVA_DB.Where("refund_no = ? AND status = 0", req.RefundNo).First(&refund).Error; err != nil {
		return fmt.Errorf("退款单不存在或已处理")
	}
	var order shop.ShopOrder
	if err := global.GVA_DB.Where("order_no = ?", refund.OrderNo).First(&order).Error; err != nil {
		return fmt.Errorf("订单不存在")
	}

	if req.Approved {
		if refund.Type == 1 {
			return s.executeRefund(refund, order, operatorID)
		}
		return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			tx.Model(&refund).Updates(map[string]interface{}{"status": 1, "admin_remark": req.AdminRemark})
			log := shop.ShopRefundLog{RefundNo: req.RefundNo, FromStatus: 0, ToStatus: 1, OperatorID: operatorID, Remark: "审核通过，等待退货"}
			tx.Create(&log)
			var notificationService ShopNotificationService
			go notificationService.CreateNotification(refund.UserID, 1, "退款审核通过", "您的退款申请已通过，请寄回商品", "refund", req.RefundNo)
			return nil
		})
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&refund).Updates(map[string]interface{}{"status": 2, "admin_remark": req.AdminRemark})
		tx.Model(&shop.ShopOrder{}).Where("order_no = ?", refund.OrderNo).Update("status", order.Status)
		log := shop.ShopRefundLog{RefundNo: req.RefundNo, FromStatus: 0, ToStatus: 2, OperatorID: operatorID, Remark: "审核拒绝: " + req.AdminRemark}
		tx.Create(&log)
		var notificationService ShopNotificationService
		var emailService ShopEmailService
		go notificationService.CreateNotification(refund.UserID, 1, "退款被拒绝", req.AdminRemark, "refund", req.RefundNo)
		go emailService.SendRefundRejectedEmail(refund.UserID, refund.OrderNo, req.AdminRemark)
		return nil
	})
}

func (s *ShopRefundService) executeRefund(refund shop.ShopRefund, order shop.ShopOrder, operatorID uint) error {
	var paymentService ShopPaymentService
	var inventoryService ShopInventoryService
	var couponService ShopCouponService

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := paymentService.RefundToBalance(tx, refund.UserID, refund.Amount, refund.OrderNo); err != nil {
			return err
		}
		var items []shop.ShopOrderItem
		tx.Where("order_no = ?", refund.OrderNo).Find(&items)
		for _, item := range items {
			inventoryService.ReleaseStock(tx, item.SkuID, item.Quantity, refund.OrderNo)
		}
		couponService.ReturnCoupon(tx, refund.UserID, refund.OrderNo)
		now := time.Now()
		tx.Model(&refund).Updates(map[string]interface{}{"status": 4, "refund_at": &now})
		tx.Model(&shop.ShopOrder{}).Where("order_no = ?", refund.OrderNo).Update("status", 7)
		log := shop.ShopRefundLog{RefundNo: refund.RefundNo, FromStatus: refund.Status, ToStatus: 4, OperatorID: operatorID, Remark: "退款完成"}
		tx.Create(&log)
		var notificationService ShopNotificationService
		var emailService ShopEmailService
		go notificationService.CreateNotification(refund.UserID, 1, "退款成功", fmt.Sprintf("¥%.2f 已退回钱包", refund.Amount), "refund", refund.RefundNo)
		go emailService.SendRefundApprovedEmail(refund.UserID, refund.OrderNo, refund.Amount)
		return nil
	})
}

func (s *ShopRefundService) ShipReturn(userID uint, req shopReq.ShopRefundShipReq) error {
	var refund shop.ShopRefund
	if err := global.GVA_DB.Where("refund_no = ? AND user_id = ? AND status = 1", req.RefundNo, userID).First(&refund).Error; err != nil {
		return fmt.Errorf("退款单不存在或状态不允许")
	}
	now := time.Now()
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&refund).Updates(map[string]interface{}{
			"status": 3, "ship_company": req.ShipCompany, "ship_no": req.ShipNo, "ship_at": &now,
		})
		log := shop.ShopRefundLog{RefundNo: req.RefundNo, FromStatus: 1, ToStatus: 3, OperatorID: userID, Remark: fmt.Sprintf("退货发货: %s %s", req.ShipCompany, req.ShipNo)}
		return tx.Create(&log).Error
	})
}

func (s *ShopRefundService) ConfirmReturn(refundNo string, operatorID uint) error {
	var refund shop.ShopRefund
	if err := global.GVA_DB.Where("refund_no = ? AND status = 3", refundNo).First(&refund).Error; err != nil {
		return fmt.Errorf("退款单不存在或状态不允许")
	}
	var order shop.ShopOrder
	if err := global.GVA_DB.Where("order_no = ?", refund.OrderNo).First(&order).Error; err != nil {
		return fmt.Errorf("订单不存在")
	}
	return s.executeRefund(refund, order, operatorID)
}

func (s *ShopRefundService) CancelRefund(refundNo string, userID uint) error {
	var refund shop.ShopRefund
	if err := global.GVA_DB.Where("refund_no = ? AND user_id = ? AND status IN (0, 1)", refundNo, userID).First(&refund).Error; err != nil {
		return fmt.Errorf("退款单不存在或无法取消")
	}
	var order shop.ShopOrder
	global.GVA_DB.Where("order_no = ?", refund.OrderNo).First(&order)
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		oldStatus := refund.Status
		tx.Model(&refund).Update("status", 5)
		// Restore order status: if was 6(退款中), restore to previous status stored before refund
		if order.Status == 6 {
			tx.Model(&order).Update("status", 1) // default restore to paid
		}
		log := shop.ShopRefundLog{RefundNo: refundNo, FromStatus: oldStatus, ToStatus: 5, OperatorID: userID, Remark: "用户取消退款"}
		return tx.Create(&log).Error
	})
}

func (s *ShopRefundService) GetRefundList(info shopReq.ShopRefundSearch) (list []shop.ShopRefund, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopRefund{})
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	if info.RefundNo != "" {
		db = db.Where("refund_no = ?", info.RefundNo)
	}
	if info.OrderNo != "" {
		db = db.Where("order_no = ?", info.OrderNo)
	}
	err = db.Count(&total).Error
	if err != nil { return }
	if limit != 0 { db = db.Limit(limit).Offset(offset) }
	err = db.Order("created_at desc").Find(&list).Error
	return
}

func (s *ShopRefundService) GetUserRefunds(userID uint, info shopReq.ShopRefundSearch) (list []shop.ShopRefund, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopRefund{}).Where("user_id = ?", userID)
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	err = db.Count(&total).Error
	if err != nil { return }
	if limit != 0 { db = db.Limit(limit).Offset(offset) }
	err = db.Order("created_at desc").Find(&list).Error
	return
}

func (s *ShopRefundService) GetRefundDetail(refundNo string) (refund shop.ShopRefund, err error) {
	err = global.GVA_DB.Where("refund_no = ?", refundNo).First(&refund).Error
	return
}

func (s *ShopRefundService) GetRefundLogs(refundNo string) (logs []shop.ShopRefundLog, err error) {
	err = global.GVA_DB.Where("refund_no = ?", refundNo).Order("created_at asc").Find(&logs).Error
	return
}
