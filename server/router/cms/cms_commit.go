package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmsCommitRouter struct {
}

// InitCmsCommitRouter 初始化 评论 路由信息
func (s *CmsCommitRouter) InitCmsCommitRouter(Router, PublicRouter *gin.RouterGroup) {
	cmsCommitRouter := Router.Group("cmsCommit").Use(middleware.OperationRecord())
	cmsCommitRouterWithoutRole := PublicRouter.Group("cmsCommit")

	var cmsCommitApi = v1.ApiGroupApp.CmsApiGroup.CmsCommitApi
	{
		cmsCommitRouter.POST("createCmsCommit", cmsCommitApi.CreateCmsCommit)             // 新建评论
		cmsCommitRouter.DELETE("deleteCmsCommit", cmsCommitApi.DeleteCmsCommit)           // 删除评论
		cmsCommitRouter.DELETE("deleteCmsCommitByIds", cmsCommitApi.DeleteCmsCommitByIds) // 批量删除评论
		cmsCommitRouter.PUT("updateCmsCommit", cmsCommitApi.UpdateCmsCommit)              // 更新评论
	}
	{
		cmsCommitRouterWithoutRole.GET("findCmsCommit", cmsCommitApi.FindCmsCommit)       // 根据ID获取评论
		cmsCommitRouterWithoutRole.GET("getCmsCommitList", cmsCommitApi.GetCmsCommitList) // 获取评论列表
	}
}
