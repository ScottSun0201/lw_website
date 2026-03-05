package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmsPlateRouter struct {
}

// InitCmsPlateRouter 初始化 板块 路由信息
func (s *CmsPlateRouter) InitCmsPlateRouter(Router, PubRouter *gin.RouterGroup) {
	cmsPlateRouter := Router.Group("cmsPlate").Use(middleware.OperationRecord())
	cmsPlateRouterWithoutRecord := Router.Group("cmsPlate")
	cmsPlateRouterWithoutRole := PubRouter.Group("cmsPlate")
	var cmsPlateApi = v1.ApiGroupApp.CmsApiGroup.CmsPlateApi
	{
		cmsPlateRouter.POST("createCmsPlate", cmsPlateApi.CreateCmsPlate)             // 新建板块
		cmsPlateRouter.DELETE("deleteCmsPlate", cmsPlateApi.DeleteCmsPlate)           // 删除板块
		cmsPlateRouter.DELETE("deleteCmsPlateByIds", cmsPlateApi.DeleteCmsPlateByIds) // 批量删除板块
		cmsPlateRouter.PUT("updateCmsPlate", cmsPlateApi.UpdateCmsPlate)              // 更新板块
	}
	{
		cmsPlateRouterWithoutRecord.GET("findCmsPlate", cmsPlateApi.FindCmsPlate)     // 根据ID获取板块
		cmsPlateRouterWithoutRole.GET("getCmsPlateList", cmsPlateApi.GetCmsPlateList) // 获取板块列表
	}
}
