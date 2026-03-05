package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CmsPlateApi struct {
}

var cmsPlateService = service.ServiceGroupApp.CmsServiceGroup.CmsPlateService

// CreateCmsPlate 创建板块
// @Tags CmsPlate
// @Summary 创建板块
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsPlate true "创建板块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsPlate/createCmsPlate [post]
func (cmsPlateApi *CmsPlateApi) CreateCmsPlate(c *gin.Context) {
	var cmsPlate cms.CmsPlate
	err := c.ShouldBindJSON(&cmsPlate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsPlateService.CreateCmsPlate(&cmsPlate); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsPlate 删除板块
// @Tags CmsPlate
// @Summary 删除板块
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsPlate true "删除板块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsPlate/deleteCmsPlate [delete]
func (cmsPlateApi *CmsPlateApi) DeleteCmsPlate(c *gin.Context) {
	ID := c.Query("ID")
	if err := cmsPlateService.DeleteCmsPlate(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsPlateByIds 批量删除板块
// @Tags CmsPlate
// @Summary 批量删除板块
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsPlate/deleteCmsPlateByIds [delete]
func (cmsPlateApi *CmsPlateApi) DeleteCmsPlateByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := cmsPlateService.DeleteCmsPlateByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsPlate 更新板块
// @Tags CmsPlate
// @Summary 更新板块
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsPlate true "更新板块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsPlate/updateCmsPlate [put]
func (cmsPlateApi *CmsPlateApi) UpdateCmsPlate(c *gin.Context) {
	var cmsPlate cms.CmsPlate
	err := c.ShouldBindJSON(&cmsPlate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsPlateService.UpdateCmsPlate(cmsPlate); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsPlate 用id查询板块
// @Tags CmsPlate
// @Summary 用id查询板块
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.CmsPlate true "用id查询板块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsPlate/findCmsPlate [get]
func (cmsPlateApi *CmsPlateApi) FindCmsPlate(c *gin.Context) {
	ID := c.Query("ID")
	if recmsPlate, err := cmsPlateService.GetCmsPlate(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmsPlate": recmsPlate}, c)
	}
}

// GetCmsPlateList 分页获取板块列表
// @Tags CmsPlate
// @Summary 分页获取板块列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.CmsPlateSearch true "分页获取板块列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsPlate/getCmsPlateList [get]
func (cmsPlateApi *CmsPlateApi) GetCmsPlateList(c *gin.Context) {
	if list, err := cmsPlateService.GetCmsPlateInfoList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
