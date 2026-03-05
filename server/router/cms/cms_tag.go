package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmsTagRouter struct {
}

// InitCmsTagRouter 初始化 标签 路由信息
func (s *CmsTagRouter) InitCmsTagRouter(Router, PubRouter *gin.RouterGroup) {
	cmsTagRouter := Router.Group("cmsTag").Use(middleware.OperationRecord())
	cmsTagRouterWithoutRecord := Router.Group("cmsTag")
	cmsTagRouterWithoutRole := PubRouter.Group("cmsTag")
	var cmsTagApi = v1.ApiGroupApp.CmsApiGroup.CmsTagApi
	{
		cmsTagRouter.POST("createCmsTag", cmsTagApi.CreateCmsTag)             // 新建标签
		cmsTagRouter.DELETE("deleteCmsTag", cmsTagApi.DeleteCmsTag)           // 删除标签
		cmsTagRouter.DELETE("deleteCmsTagByIds", cmsTagApi.DeleteCmsTagByIds) // 批量删除标签
		cmsTagRouter.PUT("updateCmsTag", cmsTagApi.UpdateCmsTag)              // 更新标签
	}
	{
		cmsTagRouterWithoutRecord.GET("findCmsTag", cmsTagApi.FindCmsTag)     // 根据ID获取标签
		cmsTagRouterWithoutRole.GET("getCmsTagList", cmsTagApi.GetCmsTagList) // 获取标签列表
	}
}
