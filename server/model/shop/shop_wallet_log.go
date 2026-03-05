package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// Type: 1充值 2消费 3退款 4调整
type ShopWalletLog struct {
	global.GVA_MODEL
	UserID    uint    `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;index"`
	Type      int     `json:"type" form:"type" gorm:"column:type;comment:类型1充值2消费3退款4调整"`
	Amount    float64 `json:"amount" form:"amount" gorm:"column:amount;comment:金额;type:decimal(10,2)"`
	BeforeBal float64 `json:"beforeBal" form:"beforeBal" gorm:"column:before_bal;comment:操作前余额;type:decimal(10,2)"`
	AfterBal  float64 `json:"afterBal" form:"afterBal" gorm:"column:after_bal;comment:操作后余额;type:decimal(10,2)"`
	OrderNo   string  `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:关联订单号"`
}

func (ShopWalletLog) TableName() string {
	return "shop_wallet_log"
}
