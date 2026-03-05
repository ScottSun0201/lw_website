package client

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClientUserRouter struct {
}

// InitClientUserRouter 初始化 客户端用户 路由信息
func (s *ClientUserRouter) InitClientUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	clientUserRouter := Router.Group("clientUser").Use(middleware.OperationRecord())
	clientUserRouterWithoutRecord := Router.Group("clientUser")
	clientUserRouterPublic := PublicRouter.Group("clientUser")
	var clientUserApi = v1.ApiGroupApp.ClientApiGroup.ClientUserApi
	{
		clientUserRouter.POST("createClientUser", clientUserApi.CreateClientUser)             // 新建客户端用户
		clientUserRouter.DELETE("deleteClientUser", clientUserApi.DeleteClientUser)           // 删除客户端用户
		clientUserRouter.DELETE("deleteClientUserByIds", clientUserApi.DeleteClientUserByIds) // 批量删除客户端用户
		clientUserRouter.PUT("updateClientUser", clientUserApi.UpdateClientUser)              // 更新客户端用户
	}
	{
		clientUserRouterWithoutRecord.GET("findClientUser", clientUserApi.FindClientUser)       // 根据ID获取客户端用户
		clientUserRouterWithoutRecord.GET("getClientUserList", clientUserApi.GetClientUserList) // 获取客户端用户列表
	}
	{
		clientUserRouterPublic.POST("register", clientUserApi.Register)          // 注册
		clientUserRouterPublic.POST("login", clientUserApi.Login)                // 登录
		clientUserRouterPublic.GET("getUserInfo", clientUserApi.GetUserInfo)     // 获取用户信息
		clientUserRouterPublic.PUT("setClientUser", clientUserApi.SetClientUser) // 更新客户端用户
	}
}
