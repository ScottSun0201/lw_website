package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopRefundRouter struct{}

func (s *ShopRefundRouter) InitShopRefundRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	shopRefundRouter := Router.Group("shopRefund").Use(middleware.OperationRecord())
	var api = v1.ApiGroupApp.ShopApiGroup.ShopRefundApi
	{
		shopRefundRouter.GET("getRefundList", api.GetRefundList)
		shopRefundRouter.PUT("auditRefund", api.AuditRefund)
		shopRefundRouter.PUT("confirmReturn", api.ConfirmReturn)
	}
	shopRefundPublicRouter := PublicRouter.Group("shopRefund")
	{
		shopRefundPublicRouter.POST("createRefund", api.CreateRefund)
		shopRefundPublicRouter.PUT("shipReturn", api.ShipReturn)
		shopRefundPublicRouter.PUT("cancelRefund", api.CancelRefund)
		shopRefundPublicRouter.GET("getUserRefunds", api.GetUserRefunds)
		shopRefundPublicRouter.GET("getRefundDetail", api.GetRefundDetail)
	}
}
