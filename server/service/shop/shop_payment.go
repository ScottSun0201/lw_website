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

type ShopPaymentService struct{}

// GetWallet 获取用户钱包，不存在则创建
func (s *ShopPaymentService) GetWallet(userID uint) (shop.ShopUserWallet, error) {
	var wallet shop.ShopUserWallet
	err := global.GVA_DB.Where("user_id = ?", userID).First(&wallet).Error
	if err == gorm.ErrRecordNotFound {
		wallet = shop.ShopUserWallet{
			UserID:  userID,
			Balance: 0,
			Version: 0,
		}
		err = global.GVA_DB.Create(&wallet).Error
	}
	return wallet, err
}

// BalancePay 余额支付 - 事务+乐观锁
func (s *ShopPaymentService) BalancePay(userID uint, req shopReq.ShopPayReq) error {
	var order shop.ShopOrder
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 1. 校验订单状态和归属
		if err := tx.Where("order_no = ? AND user_id = ?", req.OrderNo, userID).First(&order).Error; err != nil {
			return fmt.Errorf("订单不存在")
		}
		if order.Status != 0 {
			return fmt.Errorf("订单状态不允许支付")
		}
		// 校验订单是否已过期
		if order.ExpireTime != nil && time.Now().After(*order.ExpireTime) {
			return fmt.Errorf("订单已超时，请重新下单")
		}

		// 2. 获取钱包并校验余额（乐观锁）
		var wallet shop.ShopUserWallet
		if err := tx.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
			return fmt.Errorf("钱包不存在，请先充值")
		}
		if wallet.Balance < order.PayAmount {
			return fmt.Errorf("余额不足")
		}

		// 3. 扣减余额（乐观锁）
		result := tx.Model(&shop.ShopUserWallet{}).
			Where("id = ? AND version = ?", wallet.ID, wallet.Version).
			Updates(map[string]interface{}{
				"balance": gorm.Expr("balance - ?", order.PayAmount),
				"version": gorm.Expr("version + 1"),
			})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return fmt.Errorf("支付失败，请重试")
		}

		// 4. 创建钱包流水
		walletLog := shop.ShopWalletLog{
			UserID:    userID,
			Type:      2, // 消费
			Amount:    order.PayAmount,
			BeforeBal: wallet.Balance,
			AfterBal:  wallet.Balance - order.PayAmount,
			OrderNo:   order.OrderNo,
		}
		if err := tx.Create(&walletLog).Error; err != nil {
			return err
		}

		// 5. 创建支付记录
		now := time.Now()
		randNum, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		payNo := fmt.Sprintf("PAY%s%06d", now.Format("20060102150405"), randNum.Int64())
		payment := shop.ShopPayment{
			PayNo:   payNo,
			OrderNo: order.OrderNo,
			UserID:  userID,
			Amount:  order.PayAmount,
			Method:  "余额",
			Status:  1, // 已支付
			PaidAt:  &now,
		}
		if err := tx.Create(&payment).Error; err != nil {
			return err
		}

		// 6. 更新订单状态
		if err := tx.Model(&shop.ShopOrder{}).
			Where("order_no = ?", order.OrderNo).
			Updates(map[string]interface{}{
				"status":     1,
				"pay_time":   &now,
				"pay_method": "余额",
			}).Error; err != nil {
			return err
		}

		// 7. 记录订单日志
		orderLog := shop.ShopOrderLog{
			OrderNo:    order.OrderNo,
			FromStatus: 0,
			ToStatus:   1,
			OperatorID: userID,
			Remark:     "余额支付成功",
		}
		return tx.Create(&orderLog).Error
	})
	if err != nil {
		return err
	}

	// 支付成功钩子（异步）
	go func() {
		var notificationService ShopNotificationService
		var emailService ShopEmailService
		notificationService.CreateNotification(userID, 1, "支付成功", "您的订单 "+req.OrderNo+" 已支付成功", "order", req.OrderNo)
		emailService.SendOrderPaidEmail(userID, req.OrderNo, order.PayAmount)
		// 管理端通知
		notificationService.CreateAdminNotification(3, "新订单支付", "订单 "+req.OrderNo+" 已完成支付", "order", req.OrderNo)
	}()

	return nil
}

// AdminRechargeWallet 管理员充值
func (s *ShopPaymentService) AdminRechargeWallet(req shopReq.ShopWalletRechargeReq, operatorID uint) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 获取或创建钱包
		var wallet shop.ShopUserWallet
		err := tx.Where("user_id = ?", req.UserID).First(&wallet).Error
		if err == gorm.ErrRecordNotFound {
			wallet = shop.ShopUserWallet{
				UserID:  req.UserID,
				Balance: 0,
				Version: 0,
			}
			if err := tx.Create(&wallet).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		// 乐观锁更新余额
		result := tx.Model(&shop.ShopUserWallet{}).
			Where("id = ? AND version = ?", wallet.ID, wallet.Version).
			Updates(map[string]interface{}{
				"balance": gorm.Expr("balance + ?", req.Amount),
				"version": gorm.Expr("version + 1"),
			})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return fmt.Errorf("充值失败，请重试")
		}

		// 创建钱包流水
		walletLog := shop.ShopWalletLog{
			UserID:    req.UserID,
			Type:      1, // 充值
			Amount:    req.Amount,
			BeforeBal: wallet.Balance,
			AfterBal:  wallet.Balance + req.Amount,
		}
		return tx.Create(&walletLog).Error
	})
}

// RefundToBalance 退款到余额
func (s *ShopPaymentService) RefundToBalance(tx *gorm.DB, userID uint, amount float64, orderNo string) error {
	var wallet shop.ShopUserWallet
	if err := tx.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return fmt.Errorf("钱包不存在")
	}

	result := tx.Model(&shop.ShopUserWallet{}).
		Where("id = ? AND version = ?", wallet.ID, wallet.Version).
		Updates(map[string]interface{}{
			"balance": gorm.Expr("balance + ?", amount),
			"version": gorm.Expr("version + 1"),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("退款失败，请重试")
	}

	walletLog := shop.ShopWalletLog{
		UserID:    userID,
		Type:      3, // 退款
		Amount:    amount,
		BeforeBal: wallet.Balance,
		AfterBal:  wallet.Balance + amount,
		OrderNo:   orderNo,
	}
	return tx.Create(&walletLog).Error
}

// GetWalletLogList 获取钱包流水列表
func (s *ShopPaymentService) GetWalletLogList(info shopReq.ShopWalletLogSearch) ([]shop.ShopWalletLog, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopWalletLog{})
	if info.UserID != nil {
		db = db.Where("user_id = ?", *info.UserID)
	}
	if info.Type != nil {
		db = db.Where("type = ?", *info.Type)
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var list []shop.ShopWalletLog
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err := db.Order("id DESC").Find(&list).Error
	return list, total, err
}
