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

type ShopBrandApi struct{}

var shopBrandService = service.ServiceGroupApp.ShopServiceGroup.ShopBrandService

func (a *ShopBrandApi) CreateShopBrand(c *gin.Context) {
	var brand shop.ShopBrand
	if err := c.ShouldBindJSON(&brand); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopBrandService.CreateShopBrand(&brand); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (a *ShopBrandApi) DeleteShopBrand(c *gin.Context) {
	ID := c.Query("ID")
	if ID == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := shopBrandService.DeleteShopBrand(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (a *ShopBrandApi) UpdateShopBrand(c *gin.Context) {
	var brand shop.ShopBrand
	if err := c.ShouldBindJSON(&brand); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopBrandService.UpdateShopBrand(brand); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (a *ShopBrandApi) FindShopBrand(c *gin.Context) {
	ID := c.Query("ID")
	if brand, err := shopBrandService.GetShopBrand(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"data": brand}, c)
	}
}

func (a *ShopBrandApi) GetShopBrandList(c *gin.Context) {
	var pageInfo shopReq.ShopBrandSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopBrandService.GetShopBrandList(pageInfo); err != nil {
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

func (a *ShopBrandApi) GetAllBrands(c *gin.Context) {
	if list, err := shopBrandService.GetAllBrands(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}
