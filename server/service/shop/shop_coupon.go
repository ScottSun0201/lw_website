package shop

import (
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"gorm.io/gorm"
)

type ShopCouponService struct{}

func (s *ShopCouponService) CreateCoupon(req shopReq.ShopCouponCreateReq) error {
	var startTime, endTime *time.Time
	if req.StartTime != "" {
		t, _ := time.Parse("2006-01-02 15:04:05", req.StartTime)
		startTime = &t
	}
	if req.EndTime != "" {
		t, _ := time.Parse("2006-01-02 15:04:05", req.EndTime)
		endTime = &t
	}
	coupon := shop.ShopCoupon{
		Name: req.Name, Code: req.Code, Type: req.Type, Value: req.Value,
		MinAmount: req.MinAmount, MaxDiscount: req.MaxDiscount,
		Scope: req.Scope, ScopeIDs: req.ScopeIDs,
		TotalCount: req.TotalCount, PerLimit: req.PerLimit,
		StartTime: startTime, EndTime: endTime, Status: req.Status,
	}
	return global.GVA_DB.Create(&coupon).Error
}

func (s *ShopCouponService) UpdateCoupon(req shopReq.ShopCouponUpdateReq) error {
	updates := map[string]interface{}{
		"name": req.Name, "type": req.Type, "value": req.Value,
		"min_amount": req.MinAmount, "max_discount": req.MaxDiscount,
		"scope": req.Scope, "scope_ids": req.ScopeIDs,
		"total_count": req.TotalCount, "per_limit": req.PerLimit, "status": req.Status,
	}
	if req.StartTime != "" {
		t, _ := time.Parse("2006-01-02 15:04:05", req.StartTime)
		updates["start_time"] = &t
	}
	if req.EndTime != "" {
		t, _ := time.Parse("2006-01-02 15:04:05", req.EndTime)
		updates["end_time"] = &t
	}
	return global.GVA_DB.Model(&shop.ShopCoupon{}).Where("id = ?", req.ID).Updates(updates).Error
}

func (s *ShopCouponService) DeleteCoupon(id uint) error {
	return global.GVA_DB.Delete(&shop.ShopCoupon{}, id).Error
}

func (s *ShopCouponService) GetCouponList(info shopReq.ShopCouponSearch) (list []shop.ShopCoupon, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopCoupon{})
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
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

func (s *ShopCouponService) GetCouponDetail(id uint) (coupon shop.ShopCoupon, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&coupon).Error
	return
}

func (s *ShopCouponService) ClaimCoupon(userID, couponID uint) error {
	var coupon shop.ShopCoupon
	if err := global.GVA_DB.Where("id = ? AND status = 1", couponID).First(&coupon).Error; err != nil {
		return fmt.Errorf("优惠券不存在或已下架")
	}
	now := time.Now()
	if coupon.EndTime != nil && now.After(*coupon.EndTime) {
		return fmt.Errorf("优惠券已过期")
	}
	if coupon.StartTime != nil && now.Before(*coupon.StartTime) {
		return fmt.Errorf("优惠券未开始")
	}
	if coupon.TotalCount > 0 && coupon.UsedCount >= coupon.TotalCount {
		return fmt.Errorf("优惠券已领完")
	}
	var claimCount int64
	global.GVA_DB.Model(&shop.ShopUserCoupon{}).Where("user_id = ? AND coupon_id = ?", userID, couponID).Count(&claimCount)
	if coupon.PerLimit > 0 && int(claimCount) >= coupon.PerLimit {
		return fmt.Errorf("已达领取上限")
	}
	uc := shop.ShopUserCoupon{UserID: userID, CouponID: couponID, Status: 0}
	return global.GVA_DB.Create(&uc).Error
}

func (s *ShopCouponService) GetMyCoupons(userID uint, info shopReq.ShopUserCouponSearch) (list []shop.ShopUserCoupon, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopUserCoupon{}).Where("user_id = ?", userID)
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
	err = db.Order("created_at desc").Preload("Coupon").Find(&list).Error
	return
}

func (s *ShopCouponService) GetAvailableCoupons(userID uint) (list []shop.ShopUserCoupon, err error) {
	now := time.Now()
	err = global.GVA_DB.Where("user_id = ? AND status = 0", userID).
		Preload("Coupon", "status = 1 AND (end_time IS NULL OR end_time > ?)", now).
		Find(&list).Error
	var valid []shop.ShopUserCoupon
	for _, uc := range list {
		if uc.Coupon.ID > 0 {
			valid = append(valid, uc)
		}
	}
	return valid, err
}

func (s *ShopCouponService) CalcDiscount(coupon shop.ShopCoupon, totalAmount float64) float64 {
	if totalAmount < coupon.MinAmount {
		return 0
	}
	var discount float64
	switch coupon.Type {
	case 1: // 满减
		discount = coupon.Value
	case 2: // 折扣
		discount = totalAmount * (1 - coupon.Value/100)
	case 3: // 固定
		discount = coupon.Value
	}
	if coupon.MaxDiscount > 0 && discount > coupon.MaxDiscount {
		discount = coupon.MaxDiscount
	}
	if discount > totalAmount {
		discount = totalAmount
	}
	return discount
}

// CalcDiscountByCode 根据优惠券码计算折扣金额，返回折扣金额和优惠券ID
func (s *ShopCouponService) CalcDiscountByCode(userID uint, code string, totalAmount float64) (float64, uint, error) {
	var coupon shop.ShopCoupon
	if err := global.GVA_DB.Where("code = ? AND status = 1", code).First(&coupon).Error; err != nil {
		return 0, 0, fmt.Errorf("优惠券不存在或已下架")
	}
	now := time.Now()
	if coupon.EndTime != nil && now.After(*coupon.EndTime) {
		return 0, 0, fmt.Errorf("优惠券已过期")
	}
	if coupon.StartTime != nil && now.Before(*coupon.StartTime) {
		return 0, 0, fmt.Errorf("优惠券未开始")
	}
	// 验证用户是否领取了该优惠券
	var uc shop.ShopUserCoupon
	if err := global.GVA_DB.Where("user_id = ? AND coupon_id = ? AND status = 0", userID, coupon.ID).First(&uc).Error; err != nil {
		return 0, 0, fmt.Errorf("您未领取该优惠券或已使用")
	}
	discount := s.CalcDiscount(coupon, totalAmount)
	if discount <= 0 {
		return 0, 0, fmt.Errorf("未满足优惠券使用条件")
	}
	return discount, coupon.ID, nil
}

func (s *ShopCouponService) UseCoupon(tx *gorm.DB, userID uint, couponID uint, orderNo string) error {
	now := time.Now()
	result := tx.Model(&shop.ShopUserCoupon{}).
		Where("user_id = ? AND coupon_id = ? AND status = 0", userID, couponID).
		Updates(map[string]interface{}{"status": 1, "order_no": orderNo, "used_at": &now})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("优惠券不可用")
	}
	tx.Model(&shop.ShopCoupon{}).Where("id = ?", couponID).Update("used_count", gorm.Expr("used_count + 1"))
	return nil
}

func (s *ShopCouponService) ReturnCoupon(tx *gorm.DB, userID uint, orderNo string) error {
	var uc shop.ShopUserCoupon
	if err := tx.Where("user_id = ? AND order_no = ? AND status = 1", userID, orderNo).First(&uc).Error; err != nil {
		return nil // no coupon used
	}
	tx.Model(&uc).Updates(map[string]interface{}{"status": 0, "order_no": "", "used_at": nil})
	tx.Model(&shop.ShopCoupon{}).Where("id = ?", uc.CouponID).Update("used_count", gorm.Expr("used_count - 1"))
	return nil
}

func (s *ShopCouponService) AutoExpireUserCoupons() {
	now := time.Now()
	global.GVA_DB.Exec(
		"UPDATE shop_user_coupon uc INNER JOIN shop_coupon c ON uc.coupon_id = c.id SET uc.status = 2 WHERE uc.status = 0 AND c.end_time IS NOT NULL AND c.end_time < ?",
		now,
	)
}
