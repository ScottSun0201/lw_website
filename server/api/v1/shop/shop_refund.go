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

type ShopRefundApi struct{}
var shopRefundService = service.ServiceGroupApp.ShopServiceGroup.ShopRefundService

func (a *ShopRefundApi) CreateRefund(c *gin.Context) {
	var req shopReq.ShopRefundCreateReq
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	userID := utils.GetUserID(c)
	if refund, err := shopRefundService.CreateRefund(userID, req); err != nil {
		global.GVA_LOG.Error("申请退款失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithData(gin.H{"data": refund}, c) }
}

func (a *ShopRefundApi) AuditRefund(c *gin.Context) {
	var req shopReq.ShopRefundAuditReq
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	operatorID := utils.GetUserID(c)
	if err := shopRefundService.AuditRefund(req, operatorID); err != nil {
		global.GVA_LOG.Error("审核失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("审核完成", c) }
}

func (a *ShopRefundApi) ShipReturn(c *gin.Context) {
	var req shopReq.ShopRefundShipReq
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	userID := utils.GetUserID(c)
	if err := shopRefundService.ShipReturn(userID, req); err != nil {
		global.GVA_LOG.Error("提交退货物流失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("提交成功", c) }
}

func (a *ShopRefundApi) ConfirmReturn(c *gin.Context) {
	refundNo := c.Query("refundNo")
	if refundNo == "" { response.FailWithMessage("参数错误", c); return }
	operatorID := utils.GetUserID(c)
	if err := shopRefundService.ConfirmReturn(refundNo, operatorID); err != nil {
		global.GVA_LOG.Error("确认收货失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("确认成功", c) }
}

func (a *ShopRefundApi) CancelRefund(c *gin.Context) {
	refundNo := c.Query("refundNo")
	if refundNo == "" { response.FailWithMessage("参数错误", c); return }
	userID := utils.GetUserID(c)
	if err := shopRefundService.CancelRefund(refundNo, userID); err != nil {
		global.GVA_LOG.Error("取消退款失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("取消成功", c) }
}

func (a *ShopRefundApi) GetRefundList(c *gin.Context) {
	var info shopReq.ShopRefundSearch
	if err := c.ShouldBindQuery(&info); err != nil { response.FailWithMessage(err.Error(), c); return }
	if list, total, err := shopRefundService.GetRefundList(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, "获取成功", c) }
}

func (a *ShopRefundApi) GetUserRefunds(c *gin.Context) {
	var info shopReq.ShopRefundSearch
	if err := c.ShouldBindQuery(&info); err != nil { response.FailWithMessage(err.Error(), c); return }
	userID := utils.GetUserID(c)
	if list, total, err := shopRefundService.GetUserRefunds(userID, info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, "获取成功", c) }
}

func (a *ShopRefundApi) GetRefundDetail(c *gin.Context) {
	refundNo := c.Query("refundNo")
	if refundNo == "" { response.FailWithMessage("参数错误", c); return }
	if refund, err := shopRefundService.GetRefundDetail(refundNo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else {
		logs, _ := shopRefundService.GetRefundLogs(refundNo)
		response.OkWithData(gin.H{"data": refund, "logs": logs}, c)
	}
}
