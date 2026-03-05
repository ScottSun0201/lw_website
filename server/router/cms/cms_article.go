package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmsArticleRouter struct {
}

// InitCmsArticleRouter 初始化 文章 路由信息
func (s *CmsArticleRouter) InitCmsArticleRouter(Router, PubRouter *gin.RouterGroup) {
	cmsArticleRouter := Router.Group("cmsArticle").Use(middleware.OperationRecord())
	cmsArticleRouterWithoutRole := PubRouter.Group("cmsArticle")
	var cmsArticleApi = v1.ApiGroupApp.CmsApiGroup.CmsArticleApi
	{
		cmsArticleRouter.POST("createCmsArticle", cmsArticleApi.CreateCmsArticle)             // 新建文章
		cmsArticleRouter.DELETE("deleteCmsArticle", cmsArticleApi.DeleteCmsArticle)           // 删除文章
		cmsArticleRouter.DELETE("deleteCmsArticleByIds", cmsArticleApi.DeleteCmsArticleByIds) // 批量删除文章
		cmsArticleRouter.PUT("updateCmsArticle", cmsArticleApi.UpdateCmsArticle)              // 更新文章
	}
	{
		cmsArticleRouterWithoutRole.GET("findCmsArticle", cmsArticleApi.FindCmsArticle)       // 根据ID获取文章
		cmsArticleRouterWithoutRole.GET("getCmsArticleList", cmsArticleApi.GetCmsArticleList) // 获取文章列表
	}
}
