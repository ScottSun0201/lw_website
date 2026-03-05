package client

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
	clientReq "github.com/flipped-aurora/gin-vue-admin/server/model/client/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ClientAboutApi struct {
}

var clientAboutService = service.ServiceGroupApp.ClientServiceGroup.ClientAboutService

func (clientAboutApi *ClientAboutApi) SetClientAboutInfo(c *gin.Context) {
	var clientAbout client.ClientAboutInfo
	err := c.ShouldBindJSON(&clientAbout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clientAboutService.SetClientAboutInfo(&clientAbout); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (clientAboutApi *ClientAboutApi) GetClientAboutInfo(c *gin.Context) {
	if req, err := clientAboutService.GetClientAboutInfo(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(req, c)
	}
}

// CreateClientAbout 创建关于我们
// @Tags ClientAbout
// @Summary 创建关于我们
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body client.ClientAbout true "创建关于我们"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /clientAbout/createClientAbout [post]
func (clientAboutApi *ClientAboutApi) CreateClientAbout(c *gin.Context) {
	var clientAbout client.ClientAbout
	err := c.ShouldBindJSON(&clientAbout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clientAboutService.CreateClientAbout(&clientAbout); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteClientAbout 删除关于我们
// @Tags ClientAbout
// @Summary 删除关于我们
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body client.ClientAbout true "删除关于我们"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clientAbout/deleteClientAbout [delete]
func (clientAboutApi *ClientAboutApi) DeleteClientAbout(c *gin.Context) {
	ID := c.Query("ID")
	if err := clientAboutService.DeleteClientAbout(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteClientAboutByIds 批量删除关于我们
// @Tags ClientAbout
// @Summary 批量删除关于我们
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /clientAbout/deleteClientAboutByIds [delete]
func (clientAboutApi *ClientAboutApi) DeleteClientAboutByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := clientAboutService.DeleteClientAboutByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateClientAbout 更新关于我们
// @Tags ClientAbout
// @Summary 更新关于我们
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body client.ClientAbout true "更新关于我们"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clientAbout/updateClientAbout [put]
func (clientAboutApi *ClientAboutApi) UpdateClientAbout(c *gin.Context) {
	var clientAbout client.ClientAbout
	err := c.ShouldBindJSON(&clientAbout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clientAboutService.UpdateClientAbout(clientAbout); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindClientAbout 用id查询关于我们
// @Tags ClientAbout
// @Summary 用id查询关于我们
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query client.ClientAbout true "用id查询关于我们"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clientAbout/findClientAbout [get]
func (clientAboutApi *ClientAboutApi) FindClientAbout(c *gin.Context) {
	ID := c.Query("ID")
	if reclientAbout, err := clientAboutService.GetClientAbout(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reclientAbout": reclientAbout}, c)
	}
}

// GetClientAboutList 分页获取关于我们列表
// @Tags ClientAbout
// @Summary 分页获取关于我们列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clientReq.ClientAboutSearch true "分页获取关于我们列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clientAbout/getClientAboutList [get]
func (clientAboutApi *ClientAboutApi) GetClientAboutList(c *gin.Context) {
	var pageInfo clientReq.ClientAboutSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := clientAboutService.GetClientAboutInfoList(pageInfo); err != nil {
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
