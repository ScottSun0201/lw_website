package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// Type: 1库存预警 2退款申请 3新订单 4新评价
// Status: 0未读 1已读
type ShopAdminNotification struct {
	global.GVA_MODEL
	Type    int    `json:"type" form:"type" gorm:"column:type;comment:通知类型1库存预警2退款申请3新订单4新评价"`
	Title   string `json:"title" form:"title" gorm:"column:title;comment:标题"`
	Content string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text"`
	Status  int    `json:"status" form:"status" gorm:"column:status;comment:状态0未读1已读;default:0"`
	BizType string `json:"bizType" form:"bizType" gorm:"column:biz_type;comment:业务类型"`
	BizID   string `json:"bizId" form:"bizId" gorm:"column:biz_id;comment:业务ID"`
}

func (ShopAdminNotification) TableName() string {
	return "shop_admin_notification"
}
