package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// Type: 1热销 2新品 3猜你喜欢 4相关 5人工
type ShopRecommendation struct {
	global.GVA_MODEL
	Type   int  `json:"type" form:"type" gorm:"column:type;comment:推荐类型1热销2新品3猜你喜欢4相关5人工;index"`
	SpuID  uint `json:"spuId" form:"spuId" gorm:"column:spu_id;comment:商品ID;index"`
	Sort   int  `json:"sort" form:"sort" gorm:"column:sort;comment:排序;default:0"`
	Status int  `json:"status" form:"status" gorm:"column:status;comment:状态0禁用1启用;default:1"`
}

func (ShopRecommendation) TableName() string {
	return "shop_recommendation"
}
