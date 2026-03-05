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

type ShopPaymentApi struct{}

var shopPaymentService = service.ServiceGroupApp.ShopServiceGroup.ShopPaymentService

// 客户端 - 余额支付
func (a *ShopPaymentApi) BalancePay(c *gin.Context) {
	var req shopReq.ShopPayReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if err := shopPaymentService.BalancePay(userID, req); err != nil {
		global.GVA_LOG.Error("支付失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("支付成功", c)
	}
}

// 客户端 - 获取钱包余额
func (a *ShopPaymentApi) GetWallet(c *gin.Context) {
	userID := utils.GetUserID(c)
	if wallet, err := shopPaymentService.GetWallet(userID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"data": wallet}, c)
	}
}

// 管理端 - 充值
func (a *ShopPaymentApi) AdminRechargeWallet(c *gin.Context) {
	var req shopReq.ShopWalletRechargeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	operatorID := utils.GetUserID(c)
	if err := shopPaymentService.AdminRechargeWallet(req, operatorID); err != nil {
		global.GVA_LOG.Error("充值失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("充值成功", c)
	}
}

// 管理端 - 钱包流水列表
func (a *ShopPaymentApi) GetWalletLogList(c *gin.Context) {
	var pageInfo shopReq.ShopWalletLogSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopPaymentService.GetWalletLogList(pageInfo); err != nil {
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
