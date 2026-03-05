package shop

import (
	"strconv"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShopRecommendationApi struct{}
var shopRecommendationService = service.ServiceGroupApp.ShopServiceGroup.ShopRecommendationService

func (a *ShopRecommendationApi) SetRecommendation(c *gin.Context) {
	var req shopReq.ShopRecommendationSetReq
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	if err := shopRecommendationService.SetRecommendation(req); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("设置成功", c) }
}

func (a *ShopRecommendationApi) RemoveRecommendation(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	if err := shopRecommendationService.RemoveRecommendation(uint(id)); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("删除成功", c) }
}

func (a *ShopRecommendationApi) GetRecommendationList(c *gin.Context) {
	var info shopReq.ShopRecommendationSearch
	if err := c.ShouldBindQuery(&info); err != nil { response.FailWithMessage(err.Error(), c); return }
	if list, total, err := shopRecommendationService.GetRecommendationList(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, "获取成功", c) }
}

func (a *ShopRecommendationApi) UpdateSort(c *gin.Context) {
	var req shopReq.ShopRecommendationUpdateSortReq
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	if err := shopRecommendationService.UpdateSort(req.ID, req.Sort); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("更新成功", c) }
}

func (a *ShopRecommendationApi) GetHotProducts(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)
	if list, err := shopRecommendationService.GetHotProducts(limit); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithData(gin.H{"list": list}, c) }
}

func (a *ShopRecommendationApi) GetNewProducts(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)
	if list, err := shopRecommendationService.GetNewProducts(limit); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithData(gin.H{"list": list}, c) }
}

func (a *ShopRecommendationApi) GetRelatedProducts(c *gin.Context) {
	spuIDStr := c.Query("spuId")
	spuID, _ := strconv.ParseUint(spuIDStr, 10, 64)
	limitStr := c.DefaultQuery("limit", "8")
	limit, _ := strconv.Atoi(limitStr)
	if list, err := shopRecommendationService.GetRelatedProducts(uint(spuID), limit); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithData(gin.H{"list": list}, c) }
}
