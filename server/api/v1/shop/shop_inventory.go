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

type ShopInventoryApi struct{}

var shopInventoryService = service.ServiceGroupApp.ShopServiceGroup.ShopInventoryService

func (a *ShopInventoryApi) SetStock(c *gin.Context) {
	var req shopReq.ShopInventorySetReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	operatorID := utils.GetUserID(c)
	if err := shopInventoryService.SetStock(req, operatorID); err != nil {
		global.GVA_LOG.Error("设置库存失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

func (a *ShopInventoryApi) GetInventoryList(c *gin.Context) {
	var pageInfo shopReq.ShopInventorySearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopInventoryService.GetInventoryList(pageInfo); err != nil {
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

func (a *ShopInventoryApi) GetInventoryLogList(c *gin.Context) {
	var pageInfo shopReq.ShopInventoryLogSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopInventoryService.GetInventoryLogList(pageInfo); err != nil {
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
