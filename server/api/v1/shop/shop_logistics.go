package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShopLogisticsApi struct{}
var shopLogisticsService = service.ServiceGroupApp.ShopServiceGroup.ShopLogisticsService

func (a *ShopLogisticsApi) GetTraces(c *gin.Context) {
	orderNo := c.Query("orderNo")
	if orderNo == "" { response.FailWithMessage("参数错误", c); return }
	if list, err := shopLogisticsService.GetTraces(orderNo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithData(gin.H{"list": list}, c) }
}

func (a *ShopLogisticsApi) AddTrace(c *gin.Context) {
	var req shopReq.ShopLogisticsAddReq
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	if err := shopLogisticsService.AddTrace(req.OrderNo, req.ShipNo, req.ShipCompany, "", req.Detail); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("添加成功", c) }
}
