package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopProductRouter struct{}

func (s *ShopProductRouter) InitShopProductRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	shopProductRouter := Router.Group("shopProduct").Use(middleware.OperationRecord())
	shopProductPublicRouter := PublicRouter.Group("shopProduct")
	var api = v1.ApiGroupApp.ShopApiGroup.ShopProductApi
	{
		shopProductRouter.POST("createSpu", api.CreateSpu)
		shopProductRouter.PUT("updateSpu", api.UpdateSpu)
		shopProductRouter.DELETE("deleteSpu", api.DeleteSpu)
		shopProductRouter.GET("getSpu", api.GetSpu)
		shopProductRouter.GET("getSpuList", api.GetSpuList)
		shopProductRouter.PUT("updateSpuStatus", api.UpdateSpuStatus)
	}
	{
		shopProductPublicRouter.GET("getClientProductList", api.GetClientProductList)
		shopProductPublicRouter.GET("getClientProductDetail", api.GetClientProductDetail)
	}
}
