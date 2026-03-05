package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopBrandRouter struct{}

func (s *ShopBrandRouter) InitShopBrandRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	shopBrandRouter := Router.Group("shopBrand").Use(middleware.OperationRecord())
	shopBrandPublicRouter := PublicRouter.Group("shopBrand")
	var api = v1.ApiGroupApp.ShopApiGroup.ShopBrandApi
	{
		shopBrandRouter.POST("createShopBrand", api.CreateShopBrand)
		shopBrandRouter.DELETE("deleteShopBrand", api.DeleteShopBrand)
		shopBrandRouter.PUT("updateShopBrand", api.UpdateShopBrand)
	}
	{
		shopBrandPublicRouter.GET("findShopBrand", api.FindShopBrand)
		shopBrandPublicRouter.GET("getShopBrandList", api.GetShopBrandList)
		shopBrandPublicRouter.GET("getAllBrands", api.GetAllBrands)
	}
}
