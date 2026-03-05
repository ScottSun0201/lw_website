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

type ShopFavoriteApi struct{}
var shopFavoriteService = service.ServiceGroupApp.ShopServiceGroup.ShopFavoriteService

func (a *ShopFavoriteApi) AddFavorite(c *gin.Context) {
	var req shopReq.ShopFavoriteReq
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	userID := utils.GetUserID(c)
	if err := shopFavoriteService.AddFavorite(userID, req.SpuID); err != nil {
		global.GVA_LOG.Error("收藏失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("收藏成功", c) }
}

func (a *ShopFavoriteApi) RemoveFavorite(c *gin.Context) {
	var req shopReq.ShopFavoriteReq
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	userID := utils.GetUserID(c)
	if err := shopFavoriteService.RemoveFavorite(userID, req.SpuID); err != nil {
		global.GVA_LOG.Error("取消收藏失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("取消收藏成功", c) }
}

func (a *ShopFavoriteApi) GetFavoriteList(c *gin.Context) {
	var info shopReq.ShopFavoriteSearch
	if err := c.ShouldBindQuery(&info); err != nil { response.FailWithMessage(err.Error(), c); return }
	userID := utils.GetUserID(c)
	if list, total, err := shopFavoriteService.GetFavoriteList(userID, info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, "获取成功", c) }
}

func (a *ShopFavoriteApi) IsFavorite(c *gin.Context) {
	var req shopReq.ShopFavoriteReq
	if err := c.ShouldBindQuery(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	userID := utils.GetUserID(c)
	isFav := shopFavoriteService.IsFavorite(userID, req.SpuID)
	response.OkWithData(gin.H{"isFavorite": isFav}, c)
}

func (a *ShopFavoriteApi) BatchCheckFavorite(c *gin.Context) {
	var req shopReq.ShopBatchFavoriteReq
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	userID := utils.GetUserID(c)
	result := shopFavoriteService.BatchCheckFavorite(userID, req.SpuIDs)
	response.OkWithData(gin.H{"data": result}, c)
}
