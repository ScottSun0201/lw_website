package shop

import (
	"strconv"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShopCouponApi struct{}
var shopCouponService = service.ServiceGroupApp.ShopServiceGroup.ShopCouponService

func (a *ShopCouponApi) CreateCoupon(c *gin.Context) {
	var req shopReq.ShopCouponCreateReq
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	if err := shopCouponService.CreateCoupon(req); err != nil {
		global.GVA_LOG.Error("创建优惠券失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("创建成功", c) }
}

func (a *ShopCouponApi) UpdateCoupon(c *gin.Context) {
	var req shopReq.ShopCouponUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage(err.Error(), c); return }
	if err := shopCouponService.UpdateCoupon(req); err != nil {
		global.GVA_LOG.Error("更新优惠券失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("更新成功", c) }
}

func (a *ShopCouponApi) DeleteCoupon(c *gin.Context) {
	idStr := c.Query("id")
	id, parseErr := strconv.ParseUint(idStr, 10, 64)
	if parseErr != nil || id == 0 { response.FailWithMessage("无效的ID参数", c); return }
	if err := shopCouponService.DeleteCoupon(uint(id)); err != nil {
		global.GVA_LOG.Error("删除优惠券失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("删除成功", c) }
}

func (a *ShopCouponApi) GetCouponList(c *gin.Context) {
	var info shopReq.ShopCouponSearch
	if err := c.ShouldBindQuery(&info); err != nil { response.FailWithMessage(err.Error(), c); return }
	if list, total, err := shopCouponService.GetCouponList(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, "获取成功", c) }
}

func (a *ShopCouponApi) GetCouponDetail(c *gin.Context) {
	idStr := c.Query("id")
	id, parseErr := strconv.ParseUint(idStr, 10, 64)
	if parseErr != nil || id == 0 { response.FailWithMessage("无效的ID参数", c); return }
	if coupon, err := shopCouponService.GetCouponDetail(uint(id)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithData(gin.H{"data": coupon}, c) }
}

func (a *ShopCouponApi) ClaimCoupon(c *gin.Context) {
	var req struct{ CouponID uint `json:"couponId" binding:"required"` }
	if err := c.ShouldBindJSON(&req); err != nil { response.FailWithMessage("参数错误", c); return }
	userID := utils.GetUserID(c)
	if userID == 0 { response.FailWithMessage("请先登录", c); return }
	if err := shopCouponService.ClaimCoupon(userID, req.CouponID); err != nil {
		global.GVA_LOG.Error("领取失败!", zap.Error(err)); response.FailWithMessage(err.Error(), c)
	} else { response.OkWithMessage("领取成功", c) }
}

func (a *ShopCouponApi) GetMyCoupons(c *gin.Context) {
	var info shopReq.ShopUserCouponSearch
	if err := c.ShouldBindQuery(&info); err != nil { response.FailWithMessage("参数错误", c); return }
	userID := utils.GetUserID(c)
	if userID == 0 { response.FailWithMessage("请先登录", c); return }
	if list, total, err := shopCouponService.GetMyCoupons(userID, info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: info.Page, PageSize: info.PageSize}, "获取成功", c) }
}

func (a *ShopCouponApi) GetAvailableCoupons(c *gin.Context) {
	userID := utils.GetUserID(c)
	if userID == 0 { response.FailWithMessage("请先登录", c); return }
	if list, err := shopCouponService.GetAvailableCoupons(userID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err)); response.FailWithMessage("获取失败", c)
	} else { response.OkWithData(gin.H{"list": list}, c) }
}
