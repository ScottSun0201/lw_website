package client

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ClientSEOApi struct {
}

var clientSEOService = service.ServiceGroupApp.ClientServiceGroup.ClientSEOService

// CreateClientSEO 创建客户端SEO
// @Tags ClientSEO
// @Summary 创建客户端SEO
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body client.ClientSEO true "创建客户端SEO"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /clientSEO/createClientSEO [post]
func (clientSEOApi *ClientSEOApi) CreateClientSEO(c *gin.Context) {
	var clientSEO client.ClientSEO
	err := c.ShouldBindJSON(&clientSEO)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clientSEOService.CreateClientSEO(&clientSEO); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// FindClientSEO 用id查询客户端SEO
// @Tags ClientSEO
// @Summary 用id查询客户端SEO
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query client.ClientSEO true "用id查询客户端SEO"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clientSEO/findClientSEO [get]
func (clientSEOApi *ClientSEOApi) FindClientSEO(c *gin.Context) {
	if reclientSEO, err := clientSEOService.GetClientSEO(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reclientSEO": reclientSEO}, c)
	}
}
