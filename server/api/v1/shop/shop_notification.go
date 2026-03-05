package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type ShopNotificationApi struct{}

var shopNotificationService = service.ServiceGroupApp.ShopServiceGroup.ShopNotificationService

// 客户端 - 获取我的通知列表
func (a *ShopNotificationApi) GetMyNotifications(c *gin.Context) {
	var pageInfo shopReq.ShopNotificationSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if list, total, err := shopNotificationService.GetUserNotifications(userID, pageInfo); err != nil {
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

// 客户端 - 标记通知已读
func (a *ShopNotificationApi) MarkRead(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	userID := utils.GetUserID(c)
	if err := shopNotificationService.MarkAsRead(userID, uint(id)); err != nil {
		global.GVA_LOG.Error("标记失败!", zap.Error(err))
		response.FailWithMessage("标记失败", c)
	} else {
		response.OkWithMessage("标记成功", c)
	}
}

// 客户端 - 标记所有通知已读
func (a *ShopNotificationApi) MarkAllRead(c *gin.Context) {
	userID := utils.GetUserID(c)
	if err := shopNotificationService.MarkAllAsRead(userID); err != nil {
		global.GVA_LOG.Error("标记失败!", zap.Error(err))
		response.FailWithMessage("标记失败", c)
	} else {
		response.OkWithMessage("全部标记已读", c)
	}
}

// 客户端 - 获取未读通知数
func (a *ShopNotificationApi) GetUnreadCount(c *gin.Context) {
	userID := utils.GetUserID(c)
	if count, err := shopNotificationService.GetUnreadCount(userID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"count": count}, c)
	}
}

// 管理端 - 获取管理通知列表
func (a *ShopNotificationApi) GetAdminNotifications(c *gin.Context) {
	var pageInfo shopReq.ShopNotificationSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopNotificationService.GetAdminNotifications(pageInfo); err != nil {
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

// 管理端 - 标记管理通知已读
func (a *ShopNotificationApi) MarkAdminRead(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := shopNotificationService.MarkAdminRead(uint(id)); err != nil {
		global.GVA_LOG.Error("标记失败!", zap.Error(err))
		response.FailWithMessage("标记失败", c)
	} else {
		response.OkWithMessage("标记成功", c)
	}
}

// 管理端 - 标记所有管理通知已读
func (a *ShopNotificationApi) MarkAllAdminRead(c *gin.Context) {
	if err := shopNotificationService.MarkAllAdminRead(); err != nil {
		global.GVA_LOG.Error("标记失败!", zap.Error(err))
		response.FailWithMessage("标记失败", c)
	} else {
		response.OkWithMessage("全部标记已读", c)
	}
}

// 管理端 - 获取管理未读通知数
func (a *ShopNotificationApi) GetAdminUnreadCount(c *gin.Context) {
	if count, err := shopNotificationService.GetAdminUnreadCount(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"count": count}, c)
	}
}
