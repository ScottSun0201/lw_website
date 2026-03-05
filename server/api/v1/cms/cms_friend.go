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

type CmsFriendApi struct {
}

var cmsFriendService = service.ServiceGroupApp.CmsServiceGroup.CmsFriendService

func (cmsFriendApi *CmsFriendApi) SetCmsFriendInfo(c *gin.Context) {
	var cmsFriend cms.CmsFriendInfo
	err := c.ShouldBindJSON(&cmsFriend)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsFriendService.SetCmsFriendInfo(&cmsFriend); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (cmsFriendApi *CmsFriendApi) GetCmsFriendInfo(c *gin.Context) {
	if info, err := cmsFriendService.GetCmsFriendInfo(); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(info, c)
	}
}

// CreateCmsFriend 创建友情链接
// @Tags CmsFriend
// @Summary 创建友情链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsFriend true "创建友情链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsFriend/createCmsFriend [post]
func (cmsFriendApi *CmsFriendApi) CreateCmsFriend(c *gin.Context) {
	var cmsFriend cms.CmsFriend
	err := c.ShouldBindJSON(&cmsFriend)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsFriendService.CreateCmsFriend(&cmsFriend); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsFriend 删除友情链接
// @Tags CmsFriend
// @Summary 删除友情链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsFriend true "删除友情链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsFriend/deleteCmsFriend [delete]
func (cmsFriendApi *CmsFriendApi) DeleteCmsFriend(c *gin.Context) {
	ID := c.Query("ID")
	if err := cmsFriendService.DeleteCmsFriend(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsFriendByIds 批量删除友情链接
// @Tags CmsFriend
// @Summary 批量删除友情链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsFriend/deleteCmsFriendByIds [delete]
func (cmsFriendApi *CmsFriendApi) DeleteCmsFriendByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := cmsFriendService.DeleteCmsFriendByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsFriend 更新友情链接
// @Tags CmsFriend
// @Summary 更新友情链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.CmsFriend true "更新友情链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsFriend/updateCmsFriend [put]
func (cmsFriendApi *CmsFriendApi) UpdateCmsFriend(c *gin.Context) {
	var cmsFriend cms.CmsFriend
	err := c.ShouldBindJSON(&cmsFriend)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsFriendService.UpdateCmsFriend(cmsFriend); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsFriend 用id查询友情链接
// @Tags CmsFriend
// @Summary 用id查询友情链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.CmsFriend true "用id查询友情链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsFriend/findCmsFriend [get]
func (cmsFriendApi *CmsFriendApi) FindCmsFriend(c *gin.Context) {
	ID := c.Query("ID")
	if recmsFriend, err := cmsFriendService.GetCmsFriend(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmsFriend": recmsFriend}, c)
	}
}

// GetCmsFriendList 分页获取友情链接列表
// @Tags CmsFriend
// @Summary 分页获取友情链接列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.CmsFriendSearch true "分页获取友情链接列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsFriend/getCmsFriendList [get]
func (cmsFriendApi *CmsFriendApi) GetCmsFriendList(c *gin.Context) {
	var pageInfo cmsReq.CmsFriendSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmsFriendService.GetCmsFriendInfoList(pageInfo); err != nil {
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
