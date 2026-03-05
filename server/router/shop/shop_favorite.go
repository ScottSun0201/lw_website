package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ShopFavoriteRouter struct{}

func (s *ShopFavoriteRouter) InitShopFavoriteRouter(PublicRouter *gin.RouterGroup) {
	shopFavoritePublicRouter := PublicRouter.Group("shopFavorite")
	var api = v1.ApiGroupApp.ShopApiGroup.ShopFavoriteApi
	{
		shopFavoritePublicRouter.POST("addFavorite", api.AddFavorite)
		shopFavoritePublicRouter.POST("removeFavorite", api.RemoveFavorite)
		shopFavoritePublicRouter.GET("getFavoriteList", api.GetFavoriteList)
		shopFavoritePublicRouter.GET("isFavorite", api.IsFavorite)
		shopFavoritePublicRouter.POST("batchCheckFavorite", api.BatchCheckFavorite)
	}
}
