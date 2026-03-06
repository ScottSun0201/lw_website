package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopOrderRouter struct{}

func (s *ShopOrderRouter) InitShopOrderRouter(Router *gin.RouterGroup, ClientPrivateRouter *gin.RouterGroup) {
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

	// 客户端路由（需要 JWT 认证）
	shopOrderClientRouter := ClientPrivateRouter.Group("shopOrder")
	{
		shopOrderClientRouter.POST("createOrder", api.CreateOrder)
		shopOrderClientRouter.PUT("cancelOrder", api.CancelOrder)
		shopOrderClientRouter.PUT("confirmReceive", api.ConfirmReceive)
		shopOrderClientRouter.GET("getUserOrderList", api.GetUserOrderList)
		shopOrderClientRouter.GET("getUserOrder", api.GetUserOrder)
	}
}
