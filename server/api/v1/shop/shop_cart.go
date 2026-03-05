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

type ShopCartApi struct{}

var shopCartService = service.ServiceGroupApp.ShopServiceGroup.ShopCartService

func (a *ShopCartApi) AddToCart(c *gin.Context) {
	var req shopReq.ShopCartAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if err := shopCartService.AddToCart(userID, req); err != nil {
		global.GVA_LOG.Error("加入购物车失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("加入购物车成功", c)
	}
}

func (a *ShopCartApi) UpdateCartQuantity(c *gin.Context) {
	var req shopReq.ShopCartUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if err := shopCartService.UpdateCartQuantity(userID, req); err != nil {
		global.GVA_LOG.Error("更新数量失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (a *ShopCartApi) RemoveFromCart(c *gin.Context) {
	var req shopReq.ShopCartDeleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if err := shopCartService.RemoveFromCart(userID, req); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (a *ShopCartApi) GetCartList(c *gin.Context) {
	userID := utils.GetUserID(c)
	if list, err := shopCartService.GetCartList(userID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}

func (a *ShopCartApi) GetCartCount(c *gin.Context) {
	userID := utils.GetUserID(c)
	if count, err := shopCartService.GetCartCount(userID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"count": count}, c)
	}
}
