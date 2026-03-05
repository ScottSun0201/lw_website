package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShopOrderApi struct{}

var shopOrderService = service.ServiceGroupApp.ShopServiceGroup.ShopOrderService

// 客户端 - 创建订单
func (a *ShopOrderApi) CreateOrder(c *gin.Context) {
	var req shopReq.ShopOrderCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("请先登录", c)
		return
	}
	if order, err := shopOrderService.CreateOrder(userID, req); err != nil {
		global.GVA_LOG.Error("创建订单失败!", zap.Error(err))
		response.FailWithMessage("创建订单失败", c)
	} else {
		response.OkWithData(gin.H{"order": order}, c)
	}
}

// 客户端 - 取消订单
func (a *ShopOrderApi) CancelOrder(c *gin.Context) {
	orderNo := c.Query("orderNo")
	if orderNo == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("请先登录", c)
		return
	}
	if err := shopOrderService.CancelOrder(orderNo, userID); err != nil {
		global.GVA_LOG.Error("取消订单失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("取消成功", c)
	}
}

// 客户端 - 确认收货
func (a *ShopOrderApi) ConfirmReceive(c *gin.Context) {
	orderNo := c.Query("orderNo")
	if orderNo == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("请先登录", c)
		return
	}
	if err := shopOrderService.ConfirmReceive(orderNo, userID); err != nil {
		global.GVA_LOG.Error("确认收货失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("确认收货成功", c)
	}
}

// 客户端 - 我的订单列表
func (a *ShopOrderApi) GetUserOrderList(c *gin.Context) {
	var pageInfo shopReq.ShopOrderSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("请先登录", c)
		return
	}
	if list, total, err := shopOrderService.GetUserOrderList(userID, pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// 客户端 - 订单详情
func (a *ShopOrderApi) GetUserOrder(c *gin.Context) {
	orderNo := c.Query("orderNo")
	if orderNo == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("请先登录", c)
		return
	}
	if order, err := shopOrderService.GetUserOrder(orderNo, userID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"data": order}, c)
	}
}

// 管理端 - 订单列表
func (a *ShopOrderApi) GetOrderList(c *gin.Context) {
	var pageInfo shopReq.ShopOrderSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopOrderService.GetOrderList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// 管理端 - 订单详情
func (a *ShopOrderApi) GetOrderDetail(c *gin.Context) {
	orderNo := c.Query("orderNo")
	if orderNo == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	if order, err := shopOrderService.GetOrder(orderNo); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		logs, _ := shopOrderService.GetOrderLogs(orderNo)
		response.OkWithData(gin.H{"data": order, "logs": logs}, c)
	}
}

// 管理端 - 发货
func (a *ShopOrderApi) ShipOrder(c *gin.Context) {
	var req shopReq.ShopOrderShipReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	operatorID := utils.GetUserID(c)
	if err := shopOrderService.ShipOrder(req, operatorID); err != nil {
		global.GVA_LOG.Error("发货失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("发货成功", c)
	}
}

// 管理端 - 取消订单
func (a *ShopOrderApi) AdminCancelOrder(c *gin.Context) {
	orderNo := c.Query("orderNo")
	if orderNo == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	operatorID := utils.GetUserID(c)
	if err := shopOrderService.AdminCancelOrder(orderNo, operatorID); err != nil {
		global.GVA_LOG.Error("取消失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("取消成功", c)
	}
}

// 管理端 - 订单日志
func (a *ShopOrderApi) GetOrderLogs(c *gin.Context) {
	orderNo := c.Query("orderNo")
	if orderNo == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	if logs, err := shopOrderService.GetOrderLogs(orderNo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": logs}, c)
	}
}
