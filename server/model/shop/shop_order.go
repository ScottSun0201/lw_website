package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Status: 0待付款 1已付款待发货 2已发货 3已收货 4已完成 5已取消 6退款中 7已退款
type ShopOrder struct {
	global.GVA_MODEL
	OrderNo         string          `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号;uniqueIndex"`
	UserID          uint            `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;index"`
	Status          int             `json:"status" form:"status" gorm:"column:status;comment:订单状态;default:0"`
	TotalAmount     float64         `json:"totalAmount" form:"totalAmount" gorm:"column:total_amount;comment:商品总额;type:decimal(10,2)"`
	FreightAmount   float64         `json:"freightAmount" form:"freightAmount" gorm:"column:freight_amount;comment:运费;type:decimal(10,2);default:0"`
	DiscountAmount  float64         `json:"discountAmount" form:"discountAmount" gorm:"column:discount_amount;comment:优惠金额;type:decimal(10,2);default:0"`
	PayAmount       float64         `json:"payAmount" form:"payAmount" gorm:"column:pay_amount;comment:实付金额;type:decimal(10,2)"`
	PayMethod       string          `json:"payMethod" form:"payMethod" gorm:"column:pay_method;comment:支付方式"`
	ReceiverName    string          `json:"receiverName" form:"receiverName" gorm:"column:receiver_name;comment:收货人"`
	ReceiverPhone   string          `json:"receiverPhone" form:"receiverPhone" gorm:"column:receiver_phone;comment:收货电话"`
	ReceiverAddress string          `json:"receiverAddress" form:"receiverAddress" gorm:"column:receiver_address;comment:收货地址"`
	ShipCompany     string          `json:"shipCompany" form:"shipCompany" gorm:"column:ship_company;comment:物流公司"`
	ShipNo          string          `json:"shipNo" form:"shipNo" gorm:"column:ship_no;comment:物流单号"`
	Remark          string          `json:"remark" form:"remark" gorm:"column:remark;comment:买家备注"`
	PayTime         *time.Time      `json:"payTime" form:"payTime" gorm:"column:pay_time;comment:支付时间"`
	ShipTime        *time.Time      `json:"shipTime" form:"shipTime" gorm:"column:ship_time;comment:发货时间"`
	ReceiveTime     *time.Time      `json:"receiveTime" form:"receiveTime" gorm:"column:receive_time;comment:收货时间"`
	CompleteTime    *time.Time      `json:"completeTime" form:"completeTime" gorm:"column:complete_time;comment:完成时间"`
	CancelTime      *time.Time      `json:"cancelTime" form:"cancelTime" gorm:"column:cancel_time;comment:取消时间"`
	ExpireTime      *time.Time      `json:"expireTime" form:"expireTime" gorm:"column:expire_time;comment:过期时间"`
	CouponID        *uint           `json:"couponId" form:"couponId" gorm:"column:coupon_id;comment:优惠券ID"`
	CouponCode      string          `json:"couponCode" form:"couponCode" gorm:"column:coupon_code;comment:优惠券码"`
	Items           []ShopOrderItem `json:"items" gorm:"foreignKey:OrderNo;references:OrderNo"`
}

func (ShopOrder) TableName() string {
	return "shop_order"
}
