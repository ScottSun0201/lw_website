package shop

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
)

type ShopLogisticsService struct{}

func (s *ShopLogisticsService) AddTrace(orderNo, shipNo, shipCompany, status, detail string) error {
	trace := shop.ShopLogisticsTrace{
		OrderNo: orderNo, ShipNo: shipNo, ShipCompany: shipCompany,
		Status: status, Detail: detail, TraceTime: time.Now(),
	}
	return global.GVA_DB.Create(&trace).Error
}

func (s *ShopLogisticsService) GetTraces(orderNo string) (list []shop.ShopLogisticsTrace, err error) {
	err = global.GVA_DB.Where("order_no = ?", orderNo).Order("trace_time desc").Find(&list).Error
	return
}

func (s *ShopLogisticsService) BatchAddTraces(traces []shop.ShopLogisticsTrace) error {
	if len(traces) == 0 { return nil }
	return global.GVA_DB.Create(&traces).Error
}
