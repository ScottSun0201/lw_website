package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/shop"

type ShopCartItem struct {
	shop.ShopCart
	SpuName   string  `json:"spuName"`
	SpuImage  string  `json:"spuImage"`
	SkuName   string  `json:"skuName"`
	SkuImage  string  `json:"skuImage"`
	SkuPrice  float64 `json:"skuPrice"`
	SpecData  string  `json:"specData"`
	Stock     int     `json:"stock"`
	SpuStatus int     `json:"spuStatus"`
	SkuStatus int     `json:"skuStatus"`
	Invalid   bool    `json:"invalid"`
}

type ShopCheckoutInfo struct {
	Items       []ShopCartItem        `json:"items"`
	Address     *shop.ShopUserAddress `json:"address"`
	TotalAmount float64              `json:"totalAmount"`
}

// ========== 优惠券计算结果 ==========

type ShopCouponCalcResult struct {
	CouponID       uint    `json:"couponId"`
	CouponName     string  `json:"couponName"`
	DiscountAmount float64 `json:"discountAmount"`
	PayAmount      float64 `json:"payAmount"`
}

// ========== 数据分析 ==========

type DashboardOverview struct {
	TodayOrders     int64   `json:"todayOrders"`
	TodayRevenue    float64 `json:"todayRevenue"`
	PendingOrders   int64   `json:"pendingOrders"`
	PendingRefunds  int64   `json:"pendingRefunds"`
	TotalProducts   int64   `json:"totalProducts"`
	LowStockCount   int64   `json:"lowStockCount"`
}

type SalesReportItem struct {
	Date       string  `json:"date"`
	OrderCount int64   `json:"orderCount"`
	Revenue    float64 `json:"revenue"`
	AvgPrice   float64 `json:"avgPrice"`
}

type TopProductItem struct {
	SpuID      uint    `json:"spuId"`
	SpuName    string  `json:"spuName"`
	SalesCount int64   `json:"salesCount"`
	Revenue    float64 `json:"revenue"`
}

type CategorySalesItem struct {
	CategoryID   uint    `json:"categoryId"`
	CategoryName string  `json:"categoryName"`
	SalesCount   int64   `json:"salesCount"`
	Revenue      float64 `json:"revenue"`
}
