package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmsFriendRouter struct {
}

// InitCmsFriendRouter 初始化 友情链接 路由信息
func (s *CmsFriendRouter) InitCmsFriendRouter(Router, PubRouter *gin.RouterGroup) {
	cmsFriendRouter := Router.Group("cmsFriend").Use(middleware.OperationRecord())
	cmsFriendRouterWithoutRecord := Router.Group("cmsFriend")
	cmsFriendRouterWithoutRole := PubRouter.Group("cmsFriend")

	var cmsFriendApi = v1.ApiGroupApp.CmsApiGroup.CmsFriendApi
	{
		cmsFriendRouter.POST("createCmsFriend", cmsFriendApi.CreateCmsFriend)             // 新建友情链接
		cmsFriendRouter.DELETE("deleteCmsFriend", cmsFriendApi.DeleteCmsFriend)           // 删除友情链接
		cmsFriendRouter.DELETE("deleteCmsFriendByIds", cmsFriendApi.DeleteCmsFriendByIds) // 批量删除友情链接
		cmsFriendRouter.PUT("updateCmsFriend", cmsFriendApi.UpdateCmsFriend)              // 更新友情链接
		cmsFriendRouter.POST("setCmsFriendInfo", cmsFriendApi.SetCmsFriendInfo)           // 更新链接信息

	}
	{
		cmsFriendRouterWithoutRecord.GET("findCmsFriend", cmsFriendApi.FindCmsFriend) // 根据ID获取友情链接
	}
	{
		cmsFriendRouterWithoutRole.GET("getCmsFriendList", cmsFriendApi.GetCmsFriendList) // 获取友情链接列表
		cmsFriendRouterWithoutRole.GET("getCmsFriendInfo", cmsFriendApi.GetCmsFriendInfo) // 获取链接信息
	}
}
