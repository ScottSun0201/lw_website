package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Status: 0未使用 1已使用 2已过期
type ShopUserCoupon struct {
	global.GVA_MODEL
	UserID   uint       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;uniqueIndex:idx_user_coupon"`
	CouponID uint       `json:"couponId" form:"couponId" gorm:"column:coupon_id;comment:优惠券ID;uniqueIndex:idx_user_coupon"`
	OrderNo  string     `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:使用的订单号"`
	Status   int        `json:"status" form:"status" gorm:"column:status;comment:状态0未使用1已使用2已过期;default:0"`
	UsedAt   *time.Time `json:"usedAt" form:"usedAt" gorm:"column:used_at;comment:使用时间"`
	Coupon   ShopCoupon `json:"coupon" gorm:"foreignKey:CouponID;references:ID"`
}

func (ShopUserCoupon) TableName() string {
	return "shop_user_coupon"
}
