package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmsBannerRouter struct {
}

// InitCmsBannerRouter 初始化 轮播图 路由信息
func (s *CmsBannerRouter) InitCmsBannerRouter(Router *gin.RouterGroup) {
	cmsBannerRouter := Router.Group("cmsBanner").Use(middleware.OperationRecord())
	cmsBannerRouterWithoutRecord := Router.Group("cmsBanner")
	var cmsBannerApi = v1.ApiGroupApp.CmsApiGroup.CmsBannerApi
	{
		cmsBannerRouter.POST("createCmsBanner", cmsBannerApi.CreateCmsBanner)             // 新建轮播图
		cmsBannerRouter.DELETE("deleteCmsBanner", cmsBannerApi.DeleteCmsBanner)           // 删除轮播图
		cmsBannerRouter.DELETE("deleteCmsBannerByIds", cmsBannerApi.DeleteCmsBannerByIds) // 批量删除轮播图
		cmsBannerRouter.PUT("updateCmsBanner", cmsBannerApi.UpdateCmsBanner)              // 更新轮播图
	}
	{
		cmsBannerRouterWithoutRecord.GET("findCmsBanner", cmsBannerApi.FindCmsBanner)       // 根据ID获取轮播图
		cmsBannerRouterWithoutRecord.GET("getCmsBannerList", cmsBannerApi.GetCmsBannerList) // 获取轮播图列表
	}
}
