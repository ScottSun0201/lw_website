package client

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClientSEORouter struct {
}

// InitClientSEORouter 初始化 客户端SEO 路由信息
func (s *ClientSEORouter) InitClientSEORouter(Router, PublicRouter *gin.RouterGroup) {
	clientSEORouter := Router.Group("clientSEO").Use(middleware.OperationRecord())
	publicClientSEO := PublicRouter.Group("clientSEO")

	var clientSEOApi = v1.ApiGroupApp.ClientApiGroup.ClientSEOApi
	{
		clientSEORouter.POST("createClientSEO", clientSEOApi.CreateClientSEO) // 新建客户端SEO
	}
	{
		publicClientSEO.GET("findClientSEO", clientSEOApi.FindClientSEO) // 根据ID获取客户端SEO
	}
}
