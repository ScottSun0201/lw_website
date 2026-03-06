package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ShopFavoriteRouter struct{}

func (s *ShopFavoriteRouter) InitShopFavoriteRouter(ClientPrivateRouter *gin.RouterGroup) {
	shopFavoriteRouter := ClientPrivateRouter.Group("shopFavorite")
	var api = v1.ApiGroupApp.ShopApiGroup.ShopFavoriteApi
	{
		shopFavoriteRouter.POST("addFavorite", api.AddFavorite)
		shopFavoriteRouter.POST("removeFavorite", api.RemoveFavorite)
		shopFavoriteRouter.GET("getFavoriteList", api.GetFavoriteList)
		shopFavoriteRouter.GET("isFavorite", api.IsFavorite)
		shopFavoriteRouter.POST("batchCheckFavorite", api.BatchCheckFavorite)
	}
}
