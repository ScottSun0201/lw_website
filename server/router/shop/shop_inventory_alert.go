package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ShopInventoryAlertRouter struct{}

func (s *ShopInventoryAlertRouter) InitShopInventoryAlertRouter(Router *gin.RouterGroup) {
	shopAlertRouter := Router.Group("shopInventoryAlert")
	var api = v1.ApiGroupApp.ShopApiGroup.ShopInventoryAlertApi
	{
		shopAlertRouter.GET("getAlertList", api.GetAlertList)
		shopAlertRouter.PUT("handleAlert", api.HandleAlert)
		shopAlertRouter.GET("getSetting", api.GetSetting)
		shopAlertRouter.PUT("updateSetting", api.UpdateSetting)
	}
}
