package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type ShopAddressApi struct{}

var shopAddressService = service.ServiceGroupApp.ShopServiceGroup.ShopAddressService

func (a *ShopAddressApi) CreateAddress(c *gin.Context) {
	var address shop.ShopUserAddress
	if err := c.ShouldBindJSON(&address); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	address.UserID = utils.GetUserID(c)
	if err := shopAddressService.CreateAddress(&address); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (a *ShopAddressApi) DeleteAddress(c *gin.Context) {
	idStr := c.Query("ID")
	if idStr == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	userID := utils.GetUserID(c)
	if err := shopAddressService.DeleteAddress(uint(id), userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (a *ShopAddressApi) UpdateAddress(c *gin.Context) {
	var address shop.ShopUserAddress
	if err := c.ShouldBindJSON(&address); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	address.UserID = utils.GetUserID(c)
	if err := shopAddressService.UpdateAddress(address); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (a *ShopAddressApi) GetAddressList(c *gin.Context) {
	userID := utils.GetUserID(c)
	if list, err := shopAddressService.GetAddressList(userID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}

func (a *ShopAddressApi) SetDefaultAddress(c *gin.Context) {
	idStr := c.Query("ID")
	if idStr == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	userID := utils.GetUserID(c)
	if err := shopAddressService.SetDefaultAddress(uint(id), userID); err != nil {
		global.GVA_LOG.Error("设置默认地址失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}
