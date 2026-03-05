package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Status: 0待支付 1已支付 2已退款 3已关闭
type ShopPayment struct {
	global.GVA_MODEL
	PayNo   string     `json:"payNo" form:"payNo" gorm:"column:pay_no;comment:支付单号;uniqueIndex"`
	OrderNo string     `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号;index"`
	UserID  uint       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID"`
	Amount  float64    `json:"amount" form:"amount" gorm:"column:amount;comment:支付金额;type:decimal(10,2)"`
	Method  string     `json:"method" form:"method" gorm:"column:method;comment:支付方式"`
	Status  int        `json:"status" form:"status" gorm:"column:status;comment:状态0待支付1已支付2已退款3已关闭;default:0"`
	PaidAt  *time.Time `json:"paidAt" form:"paidAt" gorm:"column:paid_at;comment:支付时间"`
}

func (ShopPayment) TableName() string {
	return "shop_payment"
}
