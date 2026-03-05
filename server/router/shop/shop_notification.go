package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ShopNotificationRouter struct{}

func (s *ShopNotificationRouter) InitShopNotificationRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	shopNotificationRouter := Router.Group("shopNotification")
	var api = v1.ApiGroupApp.ShopApiGroup.ShopNotificationApi
	{
		shopNotificationRouter.GET("getAdminNotifications", api.GetAdminNotifications)
		shopNotificationRouter.PUT("markAdminRead", api.MarkAdminRead)
		shopNotificationRouter.PUT("markAllAdminRead", api.MarkAllAdminRead)
		shopNotificationRouter.GET("getAdminUnreadCount", api.GetAdminUnreadCount)
	}

	shopNotificationPublicRouter := PublicRouter.Group("shopNotification")
	{
		shopNotificationPublicRouter.GET("getMyNotifications", api.GetMyNotifications)
		shopNotificationPublicRouter.PUT("markRead", api.MarkRead)
		shopNotificationPublicRouter.PUT("markAllRead", api.MarkAllRead)
		shopNotificationPublicRouter.GET("getUnreadCount", api.GetUnreadCount)
	}
}
