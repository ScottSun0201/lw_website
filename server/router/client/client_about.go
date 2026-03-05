package client

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClientAboutRouter struct {
}

// InitClientAboutRouter 初始化 关于我们 路由信息
func (s *ClientAboutRouter) InitClientAboutRouter(Router, PubRouter *gin.RouterGroup) {
	clientAboutRouter := Router.Group("clientAbout").Use(middleware.OperationRecord())
	clientAboutRouterWithoutRole := PubRouter.Group("clientAbout")

	var clientAboutApi = v1.ApiGroupApp.ClientApiGroup.ClientAboutApi
	{
		clientAboutRouter.POST("createClientAbout", clientAboutApi.CreateClientAbout)             // 新建关于我们
		clientAboutRouter.DELETE("deleteClientAbout", clientAboutApi.DeleteClientAbout)           // 删除关于我们
		clientAboutRouter.DELETE("deleteClientAboutByIds", clientAboutApi.DeleteClientAboutByIds) // 批量删除关于我们
		clientAboutRouter.PUT("updateClientAbout", clientAboutApi.UpdateClientAbout)              // 更新关于我们
		clientAboutRouter.POST("setClientAboutInfo", clientAboutApi.SetClientAboutInfo)           // 设置关于我们基础信息
	}
	{
		clientAboutRouterWithoutRole.GET("getClientAboutInfo", clientAboutApi.GetClientAboutInfo) // 获取关于我们基础信息
		clientAboutRouterWithoutRole.GET("findClientAbout", clientAboutApi.FindClientAbout)       // 根据ID获取关于我们
		clientAboutRouterWithoutRole.GET("getClientAboutList", clientAboutApi.GetClientAboutList) // 获取关于我们列表
	}
}
