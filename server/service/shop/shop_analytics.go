package shop

import (
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	shopRes "github.com/flipped-aurora/gin-vue-admin/server/model/shop/response"
)

type ShopAnalyticsService struct{}

func (s *ShopAnalyticsService) GetDashboardOverview() (overview shopRes.DashboardOverview, err error) {
	today := time.Now().Format("2006-01-02")
	global.GVA_DB.Model(&shop.ShopOrder{}).Where("DATE(created_at) = ? AND status > 0", today).Count(&overview.TodayOrders)
	global.GVA_DB.Model(&shop.ShopOrder{}).Where("DATE(pay_time) = ? AND status >= 1 AND status != 5", today).Select("COALESCE(SUM(pay_amount), 0)").Scan(&overview.TodayRevenue)
	global.GVA_DB.Model(&shop.ShopOrder{}).Where("status = 1").Count(&overview.PendingOrders)
	global.GVA_DB.Model(&shop.ShopRefund{}).Where("status = 0").Count(&overview.PendingRefunds)
	global.GVA_DB.Model(&shop.ShopSpu{}).Where("status = 1").Count(&overview.TotalProducts)
	global.GVA_DB.Model(&shop.ShopInventory{}).Where("available_stock <= 10 AND available_stock > 0").Count(&overview.LowStockCount)
	return
}

func (s *ShopAnalyticsService) GetSalesReport(info shopReq.ShopAnalyticsDateRange) (list []shopRes.SalesReportItem, err error) {
	groupBy := info.GroupBy
	if groupBy == "" { groupBy = "day" }
	var dateFormat string
	switch groupBy {
	case "month":
		dateFormat = "%Y-%m"
	case "week":
		dateFormat = "%Y-%u"
	default:
		dateFormat = "%Y-%m-%d"
	}
	query := global.GVA_DB.Model(&shop.ShopOrder{}).
		Select(fmt.Sprintf("DATE_FORMAT(pay_time, '%s') as date, COUNT(*) as order_count, COALESCE(SUM(pay_amount), 0) as revenue, COALESCE(AVG(pay_amount), 0) as avg_price", dateFormat)).
		Where("status >= 1 AND status != 5 AND status != 7")
	if info.StartDate != "" {
		query = query.Where("pay_time >= ?", info.StartDate)
	}
	if info.EndDate != "" {
		query = query.Where("pay_time <= ?", info.EndDate+" 23:59:59")
	}
	err = query.Group("date").Order("date asc").Find(&list).Error
	return
}

func (s *ShopAnalyticsService) GetTopProducts(info shopReq.ShopAnalyticsDateRange) (list []shopRes.TopProductItem, err error) {
	limit := info.Limit
	if limit <= 0 { limit = 10 }
	query := global.GVA_DB.Model(&shop.ShopOrderItem{}).
		Select("spu_id, spu_name, SUM(quantity) as sales_count, SUM(total_price) as revenue").
		Joins("INNER JOIN shop_order ON shop_order_item.order_no = shop_order.order_no").
		Where("shop_order.status >= 1 AND shop_order.status != 5 AND shop_order.status != 7")
	if info.StartDate != "" {
		query = query.Where("shop_order.pay_time >= ?", info.StartDate)
	}
	if info.EndDate != "" {
		query = query.Where("shop_order.pay_time <= ?", info.EndDate+" 23:59:59")
	}
	err = query.Group("spu_id, spu_name").Order("sales_count desc").Limit(limit).Find(&list).Error
	return
}

func (s *ShopAnalyticsService) GetCategorySales(info shopReq.ShopAnalyticsDateRange) (list []shopRes.CategorySalesItem, err error) {
	query := global.GVA_DB.Model(&shop.ShopOrderItem{}).
		Select("shop_spu.category_id, shop_category.name as category_name, SUM(shop_order_item.quantity) as sales_count, SUM(shop_order_item.total_price) as revenue").
		Joins("INNER JOIN shop_order ON shop_order_item.order_no = shop_order.order_no").
		Joins("INNER JOIN shop_spu ON shop_order_item.spu_id = shop_spu.id").
		Joins("LEFT JOIN shop_category ON shop_spu.category_id = shop_category.id").
		Where("shop_order.status >= 1 AND shop_order.status != 5 AND shop_order.status != 7")
	if info.StartDate != "" {
		query = query.Where("shop_order.pay_time >= ?", info.StartDate)
	}
	if info.EndDate != "" {
		query = query.Where("shop_order.pay_time <= ?", info.EndDate+" 23:59:59")
	}
	err = query.Group("shop_spu.category_id, shop_category.name").Order("revenue desc").Find(&list).Error
	return
}

func (s *ShopAnalyticsService) GetOrderStatusDistribution() (result map[string]int64, err error) {
	result = make(map[string]int64)
	statusNames := map[int]string{0: "待付款", 1: "已付款", 2: "已发货", 3: "已收货", 4: "已完成", 5: "已取消", 6: "退款中", 7: "已退款"}
	for status, name := range statusNames {
		var count int64
		global.GVA_DB.Model(&shop.ShopOrder{}).Where("status = ?", status).Count(&count)
		result[name] = count
	}
	return
}
