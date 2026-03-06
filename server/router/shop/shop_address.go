package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ShopAddressRouter struct{}

func (s *ShopAddressRouter) InitShopAddressRouter(ClientPrivateRouter *gin.RouterGroup) {
	shopAddressRouter := ClientPrivateRouter.Group("shopAddress")
	var api = v1.ApiGroupApp.ShopApiGroup.ShopAddressApi
	{
		shopAddressRouter.POST("createAddress", api.CreateAddress)
		shopAddressRouter.DELETE("deleteAddress", api.DeleteAddress)
		shopAddressRouter.PUT("updateAddress", api.UpdateAddress)
		shopAddressRouter.GET("getAddressList", api.GetAddressList)
		shopAddressRouter.PUT("setDefaultAddress", api.SetDefaultAddress)
	}
}
