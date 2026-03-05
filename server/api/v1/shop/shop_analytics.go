package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShopAnalyticsApi struct{}
var shopAnalyticsService = service.ServiceGroupApp.ShopServiceGroup.ShopAnalyticsService

func (a *ShopAnalyticsApi) GetDashboardOverview(c *gin.Context) {
	if overview, err := shopAnalyticsService.GetDashboardOverview(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithData(gin.H{"data": overview}, c) }
}

func (a *ShopAnalyticsApi) GetSalesReport(c *gin.Context) {
	var info shopReq.ShopAnalyticsDateRange
	if err := c.ShouldBindQuery(&info); err != nil { response.FailWithMessage(err.Error(), c); return }
	if list, err := shopAnalyticsService.GetSalesReport(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithData(gin.H{"list": list}, c) }
}

func (a *ShopAnalyticsApi) GetTopProducts(c *gin.Context) {
	var info shopReq.ShopAnalyticsDateRange
	if err := c.ShouldBindQuery(&info); err != nil { response.FailWithMessage(err.Error(), c); return }
	if list, err := shopAnalyticsService.GetTopProducts(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithData(gin.H{"list": list}, c) }
}

func (a *ShopAnalyticsApi) GetCategorySales(c *gin.Context) {
	var info shopReq.ShopAnalyticsDateRange
	if err := c.ShouldBindQuery(&info); err != nil { response.FailWithMessage(err.Error(), c); return }
	if list, err := shopAnalyticsService.GetCategorySales(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithData(gin.H{"list": list}, c) }
}

func (a *ShopAnalyticsApi) GetOrderStatusDistribution(c *gin.Context) {
	if result, err := shopAnalyticsService.GetOrderStatusDistribution(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithData(gin.H{"data": result}, c) }
}
