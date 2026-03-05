package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopRecommendationRouter struct{}

func (s *ShopRecommendationRouter) InitShopRecommendationRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	shopRecRouter := Router.Group("shopRecommendation").Use(middleware.OperationRecord())
	var api = v1.ApiGroupApp.ShopApiGroup.ShopRecommendationApi
	{
		shopRecRouter.POST("setRecommendation", api.SetRecommendation)
		shopRecRouter.DELETE("removeRecommendation", api.RemoveRecommendation)
		shopRecRouter.GET("getRecommendationList", api.GetRecommendationList)
		shopRecRouter.PUT("updateSort", api.UpdateSort)
	}
	shopRecPublicRouter := PublicRouter.Group("shopRecommendation")
	{
		shopRecPublicRouter.GET("getHotProducts", api.GetHotProducts)
		shopRecPublicRouter.GET("getNewProducts", api.GetNewProducts)
		shopRecPublicRouter.GET("getRelatedProducts", api.GetRelatedProducts)
	}
}
