package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Type: 1满减 2折扣 3固定
// Scope: 0全场 1分类 2商品
type ShopCoupon struct {
	global.GVA_MODEL
	Name        string     `json:"name" form:"name" gorm:"column:name;comment:优惠券名称"`
	Code        string     `json:"code" form:"code" gorm:"column:code;comment:优惠券码;uniqueIndex"`
	Type        int        `json:"type" form:"type" gorm:"column:type;comment:类型1满减2折扣3固定"`
	Value       float64    `json:"value" form:"value" gorm:"column:value;comment:优惠值;type:decimal(10,2)"`
	MinAmount   float64    `json:"minAmount" form:"minAmount" gorm:"column:min_amount;comment:最低消费;type:decimal(10,2);default:0"`
	MaxDiscount float64    `json:"maxDiscount" form:"maxDiscount" gorm:"column:max_discount;comment:最大折扣;type:decimal(10,2);default:0"`
	Scope       int        `json:"scope" form:"scope" gorm:"column:scope;comment:适用范围0全场1分类2商品;default:0"`
	ScopeIDs    string     `json:"scopeIds" form:"scopeIds" gorm:"column:scope_ids;comment:适用ID列表JSON;type:text"`
	TotalCount  int        `json:"totalCount" form:"totalCount" gorm:"column:total_count;comment:发行总量"`
	UsedCount   int        `json:"usedCount" form:"usedCount" gorm:"column:used_count;comment:已使用数量;default:0"`
	PerLimit    int        `json:"perLimit" form:"perLimit" gorm:"column:per_limit;comment:每人限领;default:1"`
	StartTime   *time.Time `json:"startTime" form:"startTime" gorm:"column:start_time;comment:开始时间"`
	EndTime     *time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:结束时间"`
	Status      int        `json:"status" form:"status" gorm:"column:status;comment:状态0禁用1启用;default:1"`
}

func (ShopCoupon) TableName() string {
	return "shop_coupon"
}
