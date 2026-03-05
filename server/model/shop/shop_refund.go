package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Type: 1仅退款 2退货退款
// Status: 0待审 1同意 2拒绝 3退款中 4已退 5取消
type ShopRefund struct {
	global.GVA_MODEL
	RefundNo    string     `json:"refundNo" form:"refundNo" gorm:"column:refund_no;comment:退款单号;uniqueIndex"`
	OrderNo     string     `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号;index"`
	UserID      uint       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;index"`
	Type        int        `json:"type" form:"type" gorm:"column:type;comment:退款类型1仅退款2退货退款"`
	Reason      string     `json:"reason" form:"reason" gorm:"column:reason;comment:退款原因"`
	Description string     `json:"description" form:"description" gorm:"column:description;comment:详细描述;type:text"`
	Images      string     `json:"images" form:"images" gorm:"column:images;comment:图片JSON;type:text"`
	Amount      float64    `json:"amount" form:"amount" gorm:"column:amount;comment:退款金额;type:decimal(10,2)"`
	Status      int        `json:"status" form:"status" gorm:"column:status;comment:状态0待审1同意2拒绝3退款中4已退5取消;default:0"`
	AdminRemark string     `json:"adminRemark" form:"adminRemark" gorm:"column:admin_remark;comment:管理员备注;type:text"`
	RefundAt    *time.Time `json:"refundAt" form:"refundAt" gorm:"column:refund_at;comment:退款时间"`
	ShipCompany string     `json:"shipCompany" form:"shipCompany" gorm:"column:ship_company;comment:退货物流公司"`
	ShipNo      string     `json:"shipNo" form:"shipNo" gorm:"column:ship_no;comment:退货物流单号"`
	ShipAt      *time.Time `json:"shipAt" form:"shipAt" gorm:"column:ship_at;comment:退货发货时间"`
}

func (ShopRefund) TableName() string {
	return "shop_refund"
}
