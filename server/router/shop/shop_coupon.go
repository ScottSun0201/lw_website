package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopCouponRouter struct{}

func (s *ShopCouponRouter) InitShopCouponRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	shopCouponRouter := Router.Group("shopCoupon").Use(middleware.OperationRecord())
	var api = v1.ApiGroupApp.ShopApiGroup.ShopCouponApi
	{
		shopCouponRouter.POST("createCoupon", api.CreateCoupon)
		shopCouponRouter.PUT("updateCoupon", api.UpdateCoupon)
		shopCouponRouter.DELETE("deleteCoupon", api.DeleteCoupon)
		shopCouponRouter.GET("getCouponList", api.GetCouponList)
		shopCouponRouter.GET("getCouponDetail", api.GetCouponDetail)
	}
	shopCouponPublicRouter := PublicRouter.Group("shopCoupon")
	{
		shopCouponPublicRouter.POST("claimCoupon", api.ClaimCoupon)
		shopCouponPublicRouter.GET("getMyCoupons", api.GetMyCoupons)
		shopCouponPublicRouter.GET("getAvailableCoupons", api.GetAvailableCoupons)
	}
}
