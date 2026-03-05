package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopUserAddress struct {
	global.GVA_MODEL
	UserID       uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;index"`
	ReceiverName string `json:"receiverName" form:"receiverName" gorm:"column:receiver_name;comment:收货人姓名" binding:"required"`
	Phone        string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号" binding:"required"`
	Province     string `json:"province" form:"province" gorm:"column:province;comment:省"`
	City         string `json:"city" form:"city" gorm:"column:city;comment:市"`
	District     string `json:"district" form:"district" gorm:"column:district;comment:区"`
	DetailAddr   string `json:"detailAddr" form:"detailAddr" gorm:"column:detail_addr;comment:详细地址"`
	IsDefault    *int   `json:"isDefault" form:"isDefault" gorm:"column:is_default;comment:是否默认0否1是;default:0"`
	Label        string `json:"label" form:"label" gorm:"column:label;comment:标签(家/公司)"`
}

func (ShopUserAddress) TableName() string {
	return "shop_user_address"
}
