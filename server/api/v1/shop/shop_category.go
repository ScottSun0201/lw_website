package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShopCategoryApi struct{}

var shopCategoryService = service.ServiceGroupApp.ShopServiceGroup.ShopCategoryService

func (a *ShopCategoryApi) CreateShopCategory(c *gin.Context) {
	var category shop.ShopCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopCategoryService.CreateShopCategory(&category); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (a *ShopCategoryApi) DeleteShopCategory(c *gin.Context) {
	ID := c.Query("ID")
	if ID == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := shopCategoryService.DeleteShopCategory(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (a *ShopCategoryApi) UpdateShopCategory(c *gin.Context) {
	var category shop.ShopCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopCategoryService.UpdateShopCategory(category); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (a *ShopCategoryApi) FindShopCategory(c *gin.Context) {
	ID := c.Query("ID")
	if category, err := shopCategoryService.GetShopCategory(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"data": category}, c)
	}
}

func (a *ShopCategoryApi) GetShopCategoryList(c *gin.Context) {
	var pageInfo shopReq.ShopCategorySearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopCategoryService.GetShopCategoryList(pageInfo); err != nil {
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

func (a *ShopCategoryApi) GetCategoryTree(c *gin.Context) {
	if tree, err := shopCategoryService.GetCategoryTree(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": tree}, c)
	}
}
