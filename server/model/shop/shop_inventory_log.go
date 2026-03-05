package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// Type: 1入库 2出库 3锁定 4释放 5扣减
type ShopInventoryLog struct {
	global.GVA_MODEL
	SkuID      uint   `json:"skuId" form:"skuId" gorm:"column:sku_id;comment:SKU ID;index"`
	OrderNo    string `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号"`
	Type       int    `json:"type" form:"type" gorm:"column:type;comment:类型1入库2出库3锁定4释放5扣减"`
	Quantity   int    `json:"quantity" form:"quantity" gorm:"column:quantity;comment:数量"`
	BeforeQty  int    `json:"beforeQty" form:"beforeQty" gorm:"column:before_qty;comment:操作前数量"`
	AfterQty   int    `json:"afterQty" form:"afterQty" gorm:"column:after_qty;comment:操作后数量"`
	OperatorID uint   `json:"operatorId" form:"operatorId" gorm:"column:operator_id;comment:操作人ID"`
}

func (ShopInventoryLog) TableName() string {
	return "shop_inventory_log"
}
