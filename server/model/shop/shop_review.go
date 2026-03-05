package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Status: 0待审 1通过 2拒绝
type ShopReview struct {
	global.GVA_MODEL
	UserID   uint       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;index"`
	SpuID    uint       `json:"spuId" form:"spuId" gorm:"column:spu_id;comment:商品ID;index"`
	SkuID    uint       `json:"skuId" form:"skuId" gorm:"column:sku_id;comment:SKU ID"`
	OrderNo  string     `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号;uniqueIndex"`
	Rating   int        `json:"rating" form:"rating" gorm:"column:rating;comment:评分1-5"`
	Content  string     `json:"content" form:"content" gorm:"column:content;comment:评价内容;type:text"`
	Images   string     `json:"images" form:"images" gorm:"column:images;comment:评价图片JSON;type:text"`
	Status   int        `json:"status" form:"status" gorm:"column:status;comment:状态0待审1通过2拒绝;default:0"`
	IsAnon   int        `json:"isAnon" form:"isAnon" gorm:"column:is_anon;comment:是否匿名;default:0"`
	Reply    string     `json:"reply" form:"reply" gorm:"column:reply;comment:商家回复;type:text"`
	ReplyAt  *time.Time `json:"replyAt" form:"replyAt" gorm:"column:reply_at;comment:回复时间"`
	Nickname string     `json:"nickname" gorm:"-"`
	Avatar   string     `json:"avatar" gorm:"-"`
}

func (ShopReview) TableName() string {
	return "shop_review"
}
