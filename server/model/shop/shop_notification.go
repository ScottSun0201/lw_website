package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// Type: 1订单 2物流 3库存 4系统 5促销
// Status: 0未读 1已读
type ShopNotification struct {
	global.GVA_MODEL
	UserID  uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;index"`
	Type    int    `json:"type" form:"type" gorm:"column:type;comment:通知类型1订单2物流3库存4系统5促销"`
	Title   string `json:"title" form:"title" gorm:"column:title;comment:标题"`
	Content string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text"`
	Status  int    `json:"status" form:"status" gorm:"column:status;comment:状态0未读1已读;default:0"`
	BizType string `json:"bizType" form:"bizType" gorm:"column:biz_type;comment:业务类型"`
	BizID   string `json:"bizId" form:"bizId" gorm:"column:biz_id;comment:业务ID"`
}

func (ShopNotification) TableName() string {
	return "shop_notification"
}
