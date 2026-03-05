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

type CmsArticleApi struct {
}

var cmsArticleService = service.ServiceGroupApp.CmsServiceGroup.CmsArticleService

// CreateCmsArticle 创建文章
// @Tags CmsArticle
// @Summary 创建文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsArticle true "创建文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsArticle/createCmsArticle [post]
func (cmsArticleApi *CmsArticleApi) CreateCmsArticle(c *gin.Context) {
	var cmsArticle cms.CmsArticle
	err := c.ShouldBindJSON(&cmsArticle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cmsArticle.Author = utils.GetUserID(c)
	if err := cmsArticleService.CreateCmsArticle(&cmsArticle); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsArticle 删除文章
// @Tags CmsArticle
// @Summary 删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsArticle true "删除文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsArticle/deleteCmsArticle [delete]
func (cmsArticleApi *CmsArticleApi) DeleteCmsArticle(c *gin.Context) {
	ID := c.Query("ID")
	if err := cmsArticleService.DeleteCmsArticle(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsArticleByIds 批量删除文章
// @Tags CmsArticle
// @Summary 批量删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsArticle/deleteCmsArticleByIds [delete]
func (cmsArticleApi *CmsArticleApi) DeleteCmsArticleByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := cmsArticleService.DeleteCmsArticleByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsArticle 更新文章
// @Tags CmsArticle
// @Summary 更新文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsArticle true "更新文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsArticle/updateCmsArticle [put]
func (cmsArticleApi *CmsArticleApi) UpdateCmsArticle(c *gin.Context) {
	var cmsArticle cms.CmsArticle
	err := c.ShouldBindJSON(&cmsArticle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cmsArticleService.UpdateCmsArticle(cmsArticle); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsArticle 用id查询文章
// @Tags CmsArticle
// @Summary 用id查询文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.CmsArticle true "用id查询文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsArticle/findCmsArticle [get]
func (cmsArticleApi *CmsArticleApi) FindCmsArticle(c *gin.Context) {
	ID := c.Query("ID")
	if recmsArticle, err := cmsArticleService.GetCmsArticle(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmsArticle": recmsArticle}, c)
	}
}

// GetCmsArticleList 分页获取文章列表
// @Tags CmsArticle
// @Summary 分页获取文章列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.CmsArticleSearch true "分页获取文章列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsArticle/getCmsArticleList [get]
func (cmsArticleApi *CmsArticleApi) GetCmsArticleList(c *gin.Context) {
	var pageInfo cmsReq.CmsArticleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmsArticleService.GetCmsArticleInfoList(pageInfo); err != nil {
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
