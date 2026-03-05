package client

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
	clientReq "github.com/flipped-aurora/gin-vue-admin/server/model/client/request"
	clientRes "github.com/flipped-aurora/gin-vue-admin/server/model/client/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

type ClientUserApi struct {
}

var clientUserService = service.ServiceGroupApp.ClientServiceGroup.ClientUserService

// CreateClientUser 创建客户端用户
// @Tags ClientUser
// @Summary 创建客户端用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body client.ClientUser true "创建客户端用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /clientUser/createClientUser [post]
func (clientUserApi *ClientUserApi) CreateClientUser(c *gin.Context) {
	var clientUser client.ClientUser
	err := c.ShouldBindJSON(&clientUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clientUserService.CreateClientUser(&clientUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteClientUser 删除客户端用户
// @Tags ClientUser
// @Summary 删除客户端用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body client.ClientUser true "删除客户端用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clientUser/deleteClientUser [delete]
func (clientUserApi *ClientUserApi) DeleteClientUser(c *gin.Context) {
	ID := c.Query("ID")
	if err := clientUserService.DeleteClientUser(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteClientUserByIds 批量删除客户端用户
// @Tags ClientUser
// @Summary 批量删除客户端用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /clientUser/deleteClientUserByIds [delete]
func (clientUserApi *ClientUserApi) DeleteClientUserByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := clientUserService.DeleteClientUserByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateClientUser 更新客户端用户
// @Tags ClientUser
// @Summary 更新客户端用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body client.ClientUser true "更新客户端用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clientUser/updateClientUser [put]
func (clientUserApi *ClientUserApi) UpdateClientUser(c *gin.Context) {
	var clientUser clientReq.SetClientUser
	err := c.ShouldBindJSON(&clientUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clientUserService.UpdateClientUser(clientUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (clientUserApi *ClientUserApi) SetClientUser(c *gin.Context) {
	var clientUser clientReq.SetClientUser
	err := c.ShouldBindJSON(&clientUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	id := utils.GetUserID(c)
	if id == 0 {
		response.FailWithMessage("获取用户ID失败", c)
		return
	}
	clientUser.ID = id
	if err := clientUserService.SetClientUser(clientUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindClientUser 用id查询客户端用户
// @Tags ClientUser
// @Summary 用id查询客户端用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query client.ClientUser true "用id查询客户端用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clientUser/findClientUser [get]
func (clientUserApi *ClientUserApi) FindClientUser(c *gin.Context) {
	ID := c.Query("ID")
	if reclientUser, err := clientUserService.GetClientUser(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reclientUser": reclientUser}, c)
	}
}

// GetClientUserList 分页获取客户端用户列表
// @Tags ClientUser
// @Summary 分页获取客户端用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clientReq.ClientUserSearch true "分页获取客户端用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clientUser/getClientUserList [get]
func (clientUserApi *ClientUserApi) GetClientUserList(c *gin.Context) {
	var pageInfo clientReq.ClientUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := clientUserService.GetClientUserInfoList(pageInfo); err != nil {
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

// Register 客户端用户登录
func (clientUserApi *ClientUserApi) Register(c *gin.Context) {
	var L clientReq.ClientUser
	_ = c.ShouldBindJSON(&L)
	if err, user := clientUserService.Register(L); err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Any("err", err))
		response.FailWithMessage("注册失败,"+err.Error(), c)
	} else {
		clientUserApi.TokenNext(c, *user)
	}
}

// Login 客户端用户登录
func (clientUserApi *ClientUserApi) Login(c *gin.Context) {
	var L clientReq.ClientUser
	_ = c.ShouldBindJSON(&L)
	if err, user := clientUserService.Login(L); err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Any("err", err))
		response.FailWithMessage("注册失败,"+err.Error(), c)
	} else {
		clientUserApi.TokenNext(c, *user)
	}
}

// Login 客户端用户登录
func (clientUserApi *ClientUserApi) GetUserInfo(c *gin.Context) {
	id := utils.GetUserID(c)
	if err, user := clientUserService.GetUserInfo(id); err != nil {
		global.GVA_LOG.Error("获取个人信息失败!", zap.Any("err", err))
		response.FailWithMessage("获取个人信息失败,"+err.Error(), c)
	} else {
		response.OkWithData(user, c)
	}
}

// TokenNext 登录以后签发jwt
func (clientUserApi *ClientUserApi) TokenNext(c *gin.Context, user client.ClientUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.Nickname,
		Username:    user.Username,
		AuthorityId: 8080,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(clientRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); errors.Is(err, redis.Nil) {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(clientRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(clientRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}
