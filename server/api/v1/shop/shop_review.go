package shop

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShopReviewApi struct{}

var shopReviewService = service.ServiceGroupApp.ShopServiceGroup.ShopReviewService

// 客户端 - 创建评价
func (a *ShopReviewApi) CreateReview(c *gin.Context) {
	var req shopReq.ShopReviewCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if review, err := shopReviewService.CreateReview(userID, req); err != nil {
		global.GVA_LOG.Error("创建评价失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"data": review}, c)
	}
}

// 客户端 - 获取商品评价列表
func (a *ShopReviewApi) GetProductReviews(c *gin.Context) {
	spuIDStr := c.Query("spuId")
	spuID, err := strconv.ParseUint(spuIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	var pageInfo shopReq.ShopReviewSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopReviewService.GetProductReviews(uint(spuID), pageInfo); err != nil {
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

// 客户端 - 获取商品评价统计
func (a *ShopReviewApi) GetReviewStat(c *gin.Context) {
	spuIDStr := c.Query("spuId")
	spuID, err := strconv.ParseUint(spuIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if stat, err := shopReviewService.GetReviewStat(uint(spuID)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"data": stat}, c)
	}
}

// 客户端 - 获取我的评价
func (a *ShopReviewApi) GetMyReviews(c *gin.Context) {
	var pageInfo shopReq.ShopReviewSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if list, total, err := shopReviewService.GetUserReviews(userID, pageInfo); err != nil {
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

// 管理端 - 获取评价列表
func (a *ShopReviewApi) GetReviewList(c *gin.Context) {
	var pageInfo shopReq.ShopReviewSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopReviewService.GetReviewList(pageInfo); err != nil {
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

// 管理端 - 审核评价
func (a *ShopReviewApi) AuditReview(c *gin.Context) {
	var req shopReq.ShopReviewAuditReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopReviewService.AuditReview(req.ID, req.Status); err != nil {
		global.GVA_LOG.Error("审核失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("审核成功", c)
	}
}

// 管理端 - 回复评价
func (a *ShopReviewApi) ReplyReview(c *gin.Context) {
	var req shopReq.ShopReviewReplyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopReviewService.ReplyReview(req.ID, req.Reply); err != nil {
		global.GVA_LOG.Error("回复失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("回复成功", c)
	}
}
