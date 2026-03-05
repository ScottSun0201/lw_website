package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopLogisticsRouter struct{}

func (s *ShopLogisticsRouter) InitShopLogisticsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	shopLogisticsRouter := Router.Group("shopLogistics").Use(middleware.OperationRecord())
	var api = v1.ApiGroupApp.ShopApiGroup.ShopLogisticsApi
	{
		shopLogisticsRouter.POST("addTrace", api.AddTrace)
	}
	shopLogisticsPublicRouter := PublicRouter.Group("shopLogistics")
	{
		shopLogisticsPublicRouter.GET("getTraces", api.GetTraces)
	}
}
