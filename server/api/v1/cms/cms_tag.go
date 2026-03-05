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

type CmsTagApi struct {
}

var cmsTagService = service.ServiceGroupApp.CmsServiceGroup.CmsTagService

// CreateCmsTag 创建标签
// @Tags CmsTag
// @Summary 创建标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsTag true "创建标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsTag/createCmsTag [post]
func (cmsTagApi *CmsTagApi) CreateCmsTag(c *gin.Context) {
	var cmsTag cms.CmsTag
	err := c.ShouldBindJSON(&cmsTag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsTagService.CreateCmsTag(&cmsTag); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsTag 删除标签
// @Tags CmsTag
// @Summary 删除标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsTag true "删除标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsTag/deleteCmsTag [delete]
func (cmsTagApi *CmsTagApi) DeleteCmsTag(c *gin.Context) {
	ID := c.Query("ID")
	if err := cmsTagService.DeleteCmsTag(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsTagByIds 批量删除标签
// @Tags CmsTag
// @Summary 批量删除标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsTag/deleteCmsTagByIds [delete]
func (cmsTagApi *CmsTagApi) DeleteCmsTagByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := cmsTagService.DeleteCmsTagByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsTag 更新标签
// @Tags CmsTag
// @Summary 更新标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsTag true "更新标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsTag/updateCmsTag [put]
func (cmsTagApi *CmsTagApi) UpdateCmsTag(c *gin.Context) {
	var cmsTag cms.CmsTag
	err := c.ShouldBindJSON(&cmsTag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsTagService.UpdateCmsTag(cmsTag); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsTag 用id查询标签
// @Tags CmsTag
// @Summary 用id查询标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.CmsTag true "用id查询标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsTag/findCmsTag [get]
func (cmsTagApi *CmsTagApi) FindCmsTag(c *gin.Context) {
	ID := c.Query("ID")
	if recmsTag, err := cmsTagService.GetCmsTag(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmsTag": recmsTag}, c)
	}
}

// GetCmsTagList 分页获取标签列表
// @Tags CmsTag
// @Summary 分页获取标签列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.CmsTagSearch true "分页获取标签列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsTag/getCmsTagList [get]
func (cmsTagApi *CmsTagApi) GetCmsTagList(c *gin.Context) {
	var pageInfo cmsReq.CmsTagSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmsTagService.GetCmsTagInfoList(pageInfo); err != nil {
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
