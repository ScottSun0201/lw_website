package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopRefundRouter struct{}

func (s *ShopRefundRouter) InitShopRefundRouter(Router *gin.RouterGroup, ClientPrivateRouter *gin.RouterGroup) {
	shopRefundRouter := Router.Group("shopRefund").Use(middleware.OperationRecord())
	var api = v1.ApiGroupApp.ShopApiGroup.ShopRefundApi
	{
		shopRefundRouter.GET("getRefundList", api.GetRefundList)
		shopRefundRouter.PUT("auditRefund", api.AuditRefund)
		shopRefundRouter.PUT("confirmReturn", api.ConfirmReturn)
	}

	// 客户端路由（需要 JWT 认证）
	shopRefundClientRouter := ClientPrivateRouter.Group("shopRefund")
	{
		shopRefundClientRouter.POST("createRefund", api.CreateRefund)
		shopRefundClientRouter.PUT("shipReturn", api.ShipReturn)
		shopRefundClientRouter.PUT("cancelRefund", api.CancelRefund)
		shopRefundClientRouter.GET("getUserRefunds", api.GetUserRefunds)
		shopRefundClientRouter.GET("getRefundDetail", api.GetRefundDetail)
	}
}
