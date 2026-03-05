package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopOrderItem struct {
	global.GVA_MODEL
	OrderNo    string  `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号;index"`
	SpuID      uint    `json:"spuId" form:"spuId" gorm:"column:spu_id;comment:SPU ID"`
	SkuID      uint    `json:"skuId" form:"skuId" gorm:"column:sku_id;comment:SKU ID"`
	SpuName    string  `json:"spuName" form:"spuName" gorm:"column:spu_name;comment:商品名称快照"`
	SkuName    string  `json:"skuName" form:"skuName" gorm:"column:sku_name;comment:SKU名称快照"`
	SkuImage   string  `json:"skuImage" form:"skuImage" gorm:"column:sku_image;comment:SKU图片快照"`
	SpecData   string  `json:"specData" form:"specData" gorm:"column:spec_data;comment:规格数据快照;type:text"`
	Price      float64 `json:"price" form:"price" gorm:"column:price;comment:单价;type:decimal(10,2)"`
	Quantity   int     `json:"quantity" form:"quantity" gorm:"column:quantity;comment:数量"`
	TotalPrice float64 `json:"totalPrice" form:"totalPrice" gorm:"column:total_price;comment:小计;type:decimal(10,2)"`
}

func (ShopOrderItem) TableName() string {
	return "shop_order_item"
}
