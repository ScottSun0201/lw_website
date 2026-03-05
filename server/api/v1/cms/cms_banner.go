package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CmsBannerApi struct {
}

var cmsBannerService = service.ServiceGroupApp.CmsServiceGroup.CmsBannerService

// CreateCmsBanner 创建轮播图
// @Tags CmsBanner
// @Summary 创建轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsBanner true "创建轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsBanner/createCmsBanner [post]
func (cmsBannerApi *CmsBannerApi) CreateCmsBanner(c *gin.Context) {
	var cmsBanner cms.CmsBanner
	err := c.ShouldBindJSON(&cmsBanner)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsBannerService.CreateCmsBanner(&cmsBanner); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsBanner 删除轮播图
// @Tags CmsBanner
// @Summary 删除轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsBanner true "删除轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsBanner/deleteCmsBanner [delete]
func (cmsBannerApi *CmsBannerApi) DeleteCmsBanner(c *gin.Context) {
	ID := c.Query("ID")
	if err := cmsBannerService.DeleteCmsBanner(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsBannerByIds 批量删除轮播图
// @Tags CmsBanner
// @Summary 批量删除轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsBanner/deleteCmsBannerByIds [delete]
func (cmsBannerApi *CmsBannerApi) DeleteCmsBannerByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := cmsBannerService.DeleteCmsBannerByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsBanner 更新轮播图
// @Tags CmsBanner
// @Summary 更新轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsBanner true "更新轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsBanner/updateCmsBanner [put]
func (cmsBannerApi *CmsBannerApi) UpdateCmsBanner(c *gin.Context) {
	var cmsBanner cms.CmsBanner
	err := c.ShouldBindJSON(&cmsBanner)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsBannerService.UpdateCmsBanner(cmsBanner); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsBanner 用id查询轮播图
// @Tags CmsBanner
// @Summary 用id查询轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.CmsBanner true "用id查询轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsBanner/findCmsBanner [get]
func (cmsBannerApi *CmsBannerApi) FindCmsBanner(c *gin.Context) {
	ID := c.Query("ID")
	if recmsBanner, err := cmsBannerService.GetCmsBanner(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmsBanner": recmsBanner}, c)
	}
}

// GetCmsBannerList 分页获取轮播图列表
// @Tags CmsBanner
// @Summary 分页获取轮播图列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.CmsBannerSearch true "分页获取轮播图列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsBanner/getCmsBannerList [get]
func (cmsBannerApi *CmsBannerApi) GetCmsBannerList(c *gin.Context) {
	var pageInfo cmsReq.CmsBannerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmsBannerService.GetCmsBannerInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
