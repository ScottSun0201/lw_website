package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ShopAnalyticsRouter struct{}

func (s *ShopAnalyticsRouter) InitShopAnalyticsRouter(Router *gin.RouterGroup) {
	shopAnalyticsRouter := Router.Group("shopAnalytics")
	var api = v1.ApiGroupApp.ShopApiGroup.ShopAnalyticsApi
	{
		shopAnalyticsRouter.GET("getDashboardOverview", api.GetDashboardOverview)
		shopAnalyticsRouter.GET("getSalesReport", api.GetSalesReport)
		shopAnalyticsRouter.GET("getTopProducts", api.GetTopProducts)
		shopAnalyticsRouter.GET("getCategorySales", api.GetCategorySales)
		shopAnalyticsRouter.GET("getOrderStatusDistribution", api.GetOrderStatusDistribution)
	}
}
