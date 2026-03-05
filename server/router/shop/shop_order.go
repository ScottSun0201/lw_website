package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopOrderRouter struct{}

func (s *ShopOrderRouter) InitShopOrderRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	// 管理端路由（需要后台管理员权限）
	shopOrderRouter := Router.Group("shopOrder").Use(middleware.OperationRecord())
	var api = v1.ApiGroupApp.ShopApiGroup.ShopOrderApi
	{
		shopOrderRouter.GET("getOrderList", api.GetOrderList)
		shopOrderRouter.GET("getOrderDetail", api.GetOrderDetail)
		shopOrderRouter.PUT("shipOrder", api.ShipOrder)
		shopOrderRouter.PUT("adminCancelOrder", api.AdminCancelOrder)
		shopOrderRouter.GET("getOrderLogs", api.GetOrderLogs)
	}

	// 客户端路由（通过 x-token 获取用户ID）
	shopOrderPublicRouter := PublicRouter.Group("shopOrder")
	{
		shopOrderPublicRouter.POST("createOrder", api.CreateOrder)
		shopOrderPublicRouter.PUT("cancelOrder", api.CancelOrder)
		shopOrderPublicRouter.PUT("confirmReceive", api.ConfirmReceive)
		shopOrderPublicRouter.GET("getUserOrderList", api.GetUserOrderList)
		shopOrderPublicRouter.GET("getUserOrder", api.GetUserOrder)
	}
}
