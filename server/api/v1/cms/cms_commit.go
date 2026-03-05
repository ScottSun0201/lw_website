package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CmsCommitApi struct {
}

var cmsCommitService = service.ServiceGroupApp.CmsServiceGroup.CmsCommitService

// CreateCmsCommit 创建评论
// @Tags CmsCommit
// @Summary 创建评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsCommit true "创建评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsCommit/createCmsCommit [post]
func (cmsCommitApi *CmsCommitApi) CreateCmsCommit(c *gin.Context) {
	var cmsCommit cms.CmsCommit
	err := c.ShouldBindJSON(&cmsCommit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	cmsCommit.UserID = utils.GetUserID(c)

	if err := cmsCommitService.CreateCmsCommit(&cmsCommit); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsCommit 删除评论
// @Tags CmsCommit
// @Summary 删除评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsCommit true "删除评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsCommit/deleteCmsCommit [delete]
func (cmsCommitApi *CmsCommitApi) DeleteCmsCommit(c *gin.Context) {
	ID := c.Query("ID")
	if err := cmsCommitService.DeleteCmsCommit(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsCommitByIds 批量删除评论
// @Tags CmsCommit
// @Summary 批量删除评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsCommit/deleteCmsCommitByIds [delete]
func (cmsCommitApi *CmsCommitApi) DeleteCmsCommitByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := cmsCommitService.DeleteCmsCommitByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsCommit 更新评论
// @Tags CmsCommit
// @Summary 更新评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsCommit true "更新评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsCommit/updateCmsCommit [put]
func (cmsCommitApi *CmsCommitApi) UpdateCmsCommit(c *gin.Context) {
	var cmsCommit cms.CmsCommit
	err := c.ShouldBindJSON(&cmsCommit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsCommitService.UpdateCmsCommit(cmsCommit); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsCommit 用id查询评论
// @Tags CmsCommit
// @Summary 用id查询评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.CmsCommit true "用id查询评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsCommit/findCmsCommit [get]
func (cmsCommitApi *CmsCommitApi) FindCmsCommit(c *gin.Context) {
	ID := c.Query("ID")
	if recmsCommit, err := cmsCommitService.GetCmsCommit(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmsCommit": recmsCommit}, c)
	}
}

// GetCmsCommitList 分页获取评论列表
// @Tags CmsCommit
// @Summary 分页获取评论列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.CmsCommitSearch true "分页获取评论列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCommit/getCmsCommitList [get]
func (cmsCommitApi *CmsCommitApi) GetCmsCommitList(c *gin.Context) {
	var pageInfo cmsReq.CmsCommitSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmsCommitService.GetCmsCommitInfoList(pageInfo); err != nil {
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
