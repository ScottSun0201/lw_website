package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopInventoryRouter struct{}

func (s *ShopInventoryRouter) InitShopInventoryRouter(Router *gin.RouterGroup) {
	shopInventoryRouter := Router.Group("shopInventory").Use(middleware.OperationRecord())
	var api = v1.ApiGroupApp.ShopApiGroup.ShopInventoryApi
	{
		shopInventoryRouter.PUT("setStock", api.SetStock)
		shopInventoryRouter.GET("getInventoryList", api.GetInventoryList)
		shopInventoryRouter.GET("getInventoryLogList", api.GetInventoryLogList)
	}
}
