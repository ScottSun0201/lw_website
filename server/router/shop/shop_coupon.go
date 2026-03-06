package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopCouponRouter struct{}

func (s *ShopCouponRouter) InitShopCouponRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, ClientPrivateRouter *gin.RouterGroup) {
	shopCouponRouter := Router.Group("shopCoupon").Use(middleware.OperationRecord())
	var api = v1.ApiGroupApp.ShopApiGroup.ShopCouponApi
	{
		shopCouponRouter.POST("createCoupon", api.CreateCoupon)
		shopCouponRouter.PUT("updateCoupon", api.UpdateCoupon)
		shopCouponRouter.DELETE("deleteCoupon", api.DeleteCoupon)
		shopCouponRouter.GET("getCouponList", api.GetCouponList)
		shopCouponRouter.GET("getCouponDetail", api.GetCouponDetail)
	}

	// 公开路由（无需登录即可查看优惠券列表）
	shopCouponPublicRouter := PublicRouter.Group("shopCoupon")
	{
		shopCouponPublicRouter.GET("getCouponList", api.GetCouponList)
	}

	// 客户端路由（需要 JWT 认证）
	shopCouponClientRouter := ClientPrivateRouter.Group("shopCoupon")
	{
		shopCouponClientRouter.POST("claimCoupon", api.ClaimCoupon)
		shopCouponClientRouter.GET("getMyCoupons", api.GetMyCoupons)
		shopCouponClientRouter.GET("getAvailableCoupons", api.GetAvailableCoupons)
	}
}
