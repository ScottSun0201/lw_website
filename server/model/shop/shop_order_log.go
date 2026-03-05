package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopOrderLog struct {
	global.GVA_MODEL
	OrderNo    string `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号;index"`
	FromStatus int    `json:"fromStatus" form:"fromStatus" gorm:"column:from_status;comment:原状态"`
	ToStatus   int    `json:"toStatus" form:"toStatus" gorm:"column:to_status;comment:新状态"`
	OperatorID uint   `json:"operatorId" form:"operatorId" gorm:"column:operator_id;comment:操作人ID"`
	Remark     string `json:"remark" form:"remark" gorm:"column:remark;comment:备注"`
}

func (ShopOrderLog) TableName() string {
	return "shop_order_log"
}
