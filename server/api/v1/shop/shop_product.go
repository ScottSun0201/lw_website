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

type ShopProductApi struct{}

var shopProductService = service.ServiceGroupApp.ShopServiceGroup.ShopProductService

func (a *ShopProductApi) CreateSpu(c *gin.Context) {
	var req shopReq.CreateSpuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	spu := shop.ShopSpu{
		Name:       req.Name,
		SubTitle:   req.SubTitle,
		CategoryID: req.CategoryID,
		BrandID:    req.BrandID,
		MainImage:  req.MainImage,
		Images:     req.Images,
		Detail:     req.Detail,
		Status:     req.Status,
		Sort:       req.Sort,
	}
	for _, s := range req.Skus {
		sku := shop.ShopSku{
			Name:        s.Name,
			SkuCode:     s.SkuCode,
			SpecData:    s.SpecData,
			Price:       s.Price,
			MarketPrice: s.MarketPrice,
			CostPrice:   s.CostPrice,
			Image:       s.Image,
			Weight:      s.Weight,
			Status:      s.Status,
			Stock:       s.Stock,
		}
		spu.Skus = append(spu.Skus, sku)
	}
	if err := shopProductService.CreateSpu(&spu); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (a *ShopProductApi) UpdateSpu(c *gin.Context) {
	var req shopReq.UpdateSpuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	spu := shop.ShopSpu{
		Name:       req.Name,
		SubTitle:   req.SubTitle,
		CategoryID: req.CategoryID,
		BrandID:    req.BrandID,
		MainImage:  req.MainImage,
		Images:     req.Images,
		Detail:     req.Detail,
		Status:     req.Status,
		Sort:       req.Sort,
	}
	spu.ID = req.ID
	for _, s := range req.Skus {
		sku := shop.ShopSku{
			Name:        s.Name,
			SkuCode:     s.SkuCode,
			SpecData:    s.SpecData,
			Price:       s.Price,
			MarketPrice: s.MarketPrice,
			CostPrice:   s.CostPrice,
			Image:       s.Image,
			Weight:      s.Weight,
			Status:      s.Status,
			Stock:       s.Stock,
		}
		sku.ID = s.ID
		spu.Skus = append(spu.Skus, sku)
	}
	if err := shopProductService.UpdateSpu(&spu); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (a *ShopProductApi) DeleteSpu(c *gin.Context) {
	ID := c.Query("ID")
	if ID == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := shopProductService.DeleteSpu(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (a *ShopProductApi) GetSpu(c *gin.Context) {
	ID := c.Query("ID")
	if spu, err := shopProductService.GetSpu(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"data": spu}, c)
	}
}

func (a *ShopProductApi) GetSpuList(c *gin.Context) {
	var pageInfo shopReq.ShopSpuSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopProductService.GetSpuList(pageInfo); err != nil {
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

func (a *ShopProductApi) UpdateSpuStatus(c *gin.Context) {
	var req shopReq.ShopSpuStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopProductService.UpdateSpuStatus(req); err != nil {
		global.GVA_LOG.Error("状态更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("状态更新成功", c)
	}
}

// 客户端接口
func (a *ShopProductApi) GetClientProductList(c *gin.Context) {
	var pageInfo shopReq.ShopSpuSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopProductService.GetClientProductList(pageInfo); err != nil {
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

func (a *ShopProductApi) GetClientProductDetail(c *gin.Context) {
	ID := c.Query("ID")
	if spu, err := shopProductService.GetClientProductDetail(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"data": spu}, c)
	}
}
