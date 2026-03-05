package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopUserWallet struct {
	global.GVA_MODEL
	UserID  uint    `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;uniqueIndex"`
	Balance float64 `json:"balance" form:"balance" gorm:"column:balance;comment:余额;type:decimal(10,2);default:0"`
	Version int     `json:"version" form:"version" gorm:"column:version;comment:乐观锁版本;default:0"`
}

func (ShopUserWallet) TableName() string {
	return "shop_user_wallet"
}
