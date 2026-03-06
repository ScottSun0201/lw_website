package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ShopCartRouter struct{}

func (s *ShopCartRouter) InitShopCartRouter(ClientPrivateRouter *gin.RouterGroup) {
	shopCartRouter := ClientPrivateRouter.Group("shopCart")
	var api = v1.ApiGroupApp.ShopApiGroup.ShopCartApi
	{
		shopCartRouter.POST("addToCart", api.AddToCart)
		shopCartRouter.PUT("updateCartQuantity", api.UpdateCartQuantity)
		shopCartRouter.POST("removeFromCart", api.RemoveFromCart)
		shopCartRouter.GET("getCartList", api.GetCartList)
		shopCartRouter.GET("getCartCount", api.GetCartCount)
	}
}
