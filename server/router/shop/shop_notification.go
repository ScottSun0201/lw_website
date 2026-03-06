package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ShopNotificationRouter struct{}

func (s *ShopNotificationRouter) InitShopNotificationRouter(Router *gin.RouterGroup, ClientPrivateRouter *gin.RouterGroup) {
	shopNotificationRouter := Router.Group("shopNotification")
	var api = v1.ApiGroupApp.ShopApiGroup.ShopNotificationApi
	{
		shopNotificationRouter.GET("getAdminNotifications", api.GetAdminNotifications)
		shopNotificationRouter.PUT("markAdminRead", api.MarkAdminRead)
		shopNotificationRouter.PUT("markAllAdminRead", api.MarkAllAdminRead)
		shopNotificationRouter.GET("getAdminUnreadCount", api.GetAdminUnreadCount)
	}

	// 客户端路由（需要 JWT 认证）
	shopNotificationClientRouter := ClientPrivateRouter.Group("shopNotification")
	{
		shopNotificationClientRouter.GET("getMyNotifications", api.GetMyNotifications)
		shopNotificationClientRouter.PUT("markRead", api.MarkRead)
		shopNotificationClientRouter.PUT("markAllRead", api.MarkAllRead)
		shopNotificationClientRouter.GET("getUnreadCount", api.GetUnreadCount)
	}
}
