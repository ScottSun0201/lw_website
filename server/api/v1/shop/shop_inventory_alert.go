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

type ShopInventoryAlertApi struct{}
var shopInventoryAlertService = service.ServiceGroupApp.ShopServiceGroup.ShopInventoryAlertService

func (a *ShopInventoryAlertApi) GetAlertList(c *gin.Context) {
	var info shopReq.ShopInventoryAlertSearch
	if err := c.ShouldBindQuery(&info); err != nil { response.FailWithMessage(err.Error(), c); return }
	if list, total, err := shopInventoryAlertService.GetAlertList(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, "获取成功", c) }
}

func (a *ShopInventoryAlertApi) HandleAlert(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	if err := shopInventoryAlertService.HandleAlert(uint(id)); err != nil {
		global.GVA_LOG.Error("处理失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("处理成功", c) }
}

func (a *ShopInventoryAlertApi) GetSetting(c *gin.Context) {
	setting := shopInventoryAlertService.GetGlobalSetting()
	response.OkWithData(gin.H{"data": setting}, c)
}

func (a *ShopInventoryAlertApi) UpdateSetting(c *gin.Context) {
	var req shopReq.ShopInventorySettingReq
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	if err := shopInventoryAlertService.UpdateSetting(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("更新成功", c) }
}
