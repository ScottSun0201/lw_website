package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopPaymentRouter struct{}

func (s *ShopPaymentRouter) InitShopPaymentRouter(Router *gin.RouterGroup, ClientPrivateRouter *gin.RouterGroup) {
	// 管理端路由
	shopPaymentRouter := Router.Group("shopPayment").Use(middleware.OperationRecord())
	var api = v1.ApiGroupApp.ShopApiGroup.ShopPaymentApi
	{
		shopPaymentRouter.POST("adminRechargeWallet", api.AdminRechargeWallet)
		shopPaymentRouter.GET("getWalletLogList", api.GetWalletLogList)
	}

	// 客户端路由（需要 JWT 认证）
	shopPaymentClientRouter := ClientPrivateRouter.Group("shopPayment")
	{
		shopPaymentClientRouter.POST("balancePay", api.BalancePay)
		shopPaymentClientRouter.GET("getWallet", api.GetWallet)
	}
}
