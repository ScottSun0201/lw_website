package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

type ShopLogisticsTrace struct {
	global.GVA_MODEL
	OrderNo     string    `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号;index"`
	ShipNo      string    `json:"shipNo" form:"shipNo" gorm:"column:ship_no;comment:物流单号"`
	ShipCompany string    `json:"shipCompany" form:"shipCompany" gorm:"column:ship_company;comment:物流公司"`
	Status      string    `json:"status" form:"status" gorm:"column:status;comment:物流状态"`
	Detail      string    `json:"detail" form:"detail" gorm:"column:detail;comment:轨迹详情;type:text"`
	TraceTime   time.Time `json:"traceTime" form:"traceTime" gorm:"column:trace_time;comment:轨迹时间"`
}

func (ShopLogisticsTrace) TableName() string {
	return "shop_logistics_trace"
}
