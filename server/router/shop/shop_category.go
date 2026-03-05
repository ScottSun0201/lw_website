package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopCategoryRouter struct{}

func (s *ShopCategoryRouter) InitShopCategoryRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	shopCategoryRouter := Router.Group("shopCategory").Use(middleware.OperationRecord())
	shopCategoryPublicRouter := PublicRouter.Group("shopCategory")
	var api = v1.ApiGroupApp.ShopApiGroup.ShopCategoryApi
	{
		shopCategoryRouter.POST("createShopCategory", api.CreateShopCategory)
		shopCategoryRouter.DELETE("deleteShopCategory", api.DeleteShopCategory)
		shopCategoryRouter.PUT("updateShopCategory", api.UpdateShopCategory)
	}
	{
		shopCategoryPublicRouter.GET("findShopCategory", api.FindShopCategory)
		shopCategoryPublicRouter.GET("getShopCategoryList", api.GetShopCategoryList)
		shopCategoryPublicRouter.GET("getCategoryTree", api.GetCategoryTree)
	}
}
