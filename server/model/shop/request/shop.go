package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type ShopCategorySearch struct {
	Name     string `json:"name" form:"name"`
	ParentID *uint  `json:"parentId" form:"parentId"`
	Status   *int   `json:"status" form:"status"`
	request.PageInfo
}

type ShopBrandSearch struct {
	Name   string `json:"name" form:"name"`
	Status *int   `json:"status" form:"status"`
	request.PageInfo
}

type ShopSpuSearch struct {
	Name       string `json:"name" form:"name"`
	CategoryID *uint  `json:"categoryId" form:"categoryId"`
	BrandID    *uint  `json:"brandId" form:"brandId"`
	Status     *int   `json:"status" form:"status"`
	request.PageInfo
}

type CreateSpuReq struct {
	Name       string         `json:"name" binding:"required"`
	SubTitle   string         `json:"subTitle"`
	CategoryID uint           `json:"categoryId"`
	BrandID    uint           `json:"brandId"`
	MainImage  string         `json:"mainImage"`
	Images     string         `json:"images"`
	Detail     string         `json:"detail"`
	Status     int            `json:"status"`
	Sort       int            `json:"sort"`
	Skus       []CreateSkuReq `json:"skus"`
}

type CreateSkuReq struct {
	Name        string  `json:"name"`
	SkuCode     string  `json:"skuCode"`
	SpecData    string  `json:"specData"`
	Price       float64 `json:"price"`
	MarketPrice float64 `json:"marketPrice"`
	CostPrice   float64 `json:"costPrice"`
	Image       string  `json:"image"`
	Weight      float64 `json:"weight"`
	Status      *int    `json:"status"`
	Stock       int     `json:"stock"`
}

type UpdateSpuReq struct {
	ID         uint           `json:"ID" binding:"required"`
	Name       string         `json:"name" binding:"required"`
	SubTitle   string         `json:"subTitle"`
	CategoryID uint           `json:"categoryId"`
	BrandID    uint           `json:"brandId"`
	MainImage  string         `json:"mainImage"`
	Images     string         `json:"images"`
	Detail     string         `json:"detail"`
	Status     int            `json:"status"`
	Sort       int            `json:"sort"`
	Skus       []UpdateSkuReq `json:"skus"`
}

type UpdateSkuReq struct {
	ID          uint    `json:"ID"`
	Name        string  `json:"name"`
	SkuCode     string  `json:"skuCode"`
	SpecData    string  `json:"specData"`
	Price       float64 `json:"price"`
	MarketPrice float64 `json:"marketPrice"`
	CostPrice   float64 `json:"costPrice"`
	Image       string  `json:"image"`
	Weight      float64 `json:"weight"`
	Status      *int    `json:"status"`
	Stock       int     `json:"stock"`
}

type ShopSpuStatusReq struct {
	ID     uint `json:"ID" form:"ID" binding:"required"`
	Status int  `json:"status" form:"status"`
}

type ShopInventorySetReq struct {
	SkuID uint `json:"skuId" form:"skuId" binding:"required"`
	Stock int  `json:"stock" form:"stock"`
}

type ShopCartAddReq struct {
	SkuID    uint `json:"skuId" form:"skuId" binding:"required"`
	Quantity int  `json:"quantity" form:"quantity" binding:"required"`
}

type ShopCartUpdateReq struct {
	ID       uint `json:"id" form:"id" binding:"required"`
	Quantity int  `json:"quantity" form:"quantity" binding:"required"`
}

type ShopCartDeleteReq struct {
	IDs []uint `json:"ids" form:"ids" binding:"required"`
}

type ShopOrderCreateReq struct {
	AddressID  uint   `json:"addressId" form:"addressId" binding:"required"`
	CartIDs    []uint `json:"cartIds" form:"cartIds"`
	SkuID      uint   `json:"skuId" form:"skuId"`
	Quantity   int    `json:"quantity" form:"quantity"`
	Remark     string `json:"remark" form:"remark"`
	CouponCode string `json:"couponCode" form:"couponCode"`
}

type ShopOrderSearch struct {
	Status  *int   `json:"status" form:"status"`
	OrderNo string `json:"orderNo" form:"orderNo"`
	UserID  *uint  `json:"userId" form:"userId"`
	request.PageInfo
}

type ShopOrderShipReq struct {
	OrderNo     string `json:"orderNo" form:"orderNo" binding:"required"`
	ShipCompany string `json:"shipCompany" form:"shipCompany"`
	ShipNo      string `json:"shipNo" form:"shipNo"`
}

type ShopPayReq struct {
	OrderNo string `json:"orderNo" form:"orderNo" binding:"required"`
}

type ShopWalletRechargeReq struct {
	UserID uint    `json:"userId" form:"userId" binding:"required"`
	Amount float64 `json:"amount" form:"amount" binding:"required"`
}

type ShopInventorySearch struct {
	SkuCode string `json:"skuCode" form:"skuCode"`
	SpuName string `json:"spuName" form:"spuName"`
	request.PageInfo
}

type ShopInventoryLogSearch struct {
	SkuID uint `json:"skuId" form:"skuId"`
	request.PageInfo
}

type ShopAddressSearch struct {
	request.PageInfo
}

type ShopWalletLogSearch struct {
	UserID *uint `json:"userId" form:"userId"`
	Type   *int  `json:"type" form:"type"`
	request.PageInfo
}

// ========== 通知 ==========

type ShopNotificationSearch struct {
	Type   *int `json:"type" form:"type"`
	Status *int `json:"status" form:"status"`
	request.PageInfo
}

// ========== 评价 ==========

type ShopReviewCreateReq struct {
	OrderNo string `json:"orderNo" form:"orderNo" binding:"required"`
	Rating  int    `json:"rating" form:"rating" binding:"required"`
	Content string `json:"content" form:"content"`
	Images  string `json:"images" form:"images"`
	IsAnon  int    `json:"isAnon" form:"isAnon"`
}

type ShopReviewSearch struct {
	Status *int  `json:"status" form:"status"`
	SpuID  *uint `json:"spuId" form:"spuId"`
	Rating *int  `json:"rating" form:"rating"`
	request.PageInfo
}

type ShopReviewAuditReq struct {
	ID     uint `json:"id" form:"id" binding:"required"`
	Status int  `json:"status" form:"status" binding:"required"`
}

type ShopReviewReplyReq struct {
	ID    uint   `json:"id" form:"id" binding:"required"`
	Reply string `json:"reply" form:"reply" binding:"required"`
}

// ========== 优惠券 ==========

type ShopCouponSearch struct {
	Name   string `json:"name" form:"name"`
	Status *int   `json:"status" form:"status"`
	request.PageInfo
}

type ShopCouponCreateReq struct {
	Name        string  `json:"name" form:"name" binding:"required"`
	Code        string  `json:"code" form:"code" binding:"required"`
	Type        int     `json:"type" form:"type" binding:"required"`
	Value       float64 `json:"value" form:"value" binding:"required"`
	MinAmount   float64 `json:"minAmount" form:"minAmount"`
	MaxDiscount float64 `json:"maxDiscount" form:"maxDiscount"`
	Scope       int     `json:"scope" form:"scope"`
	ScopeIDs    string  `json:"scopeIds" form:"scopeIds"`
	TotalCount  int     `json:"totalCount" form:"totalCount"`
	PerLimit    int     `json:"perLimit" form:"perLimit"`
	StartTime   string  `json:"startTime" form:"startTime"`
	EndTime     string  `json:"endTime" form:"endTime"`
	Status      int     `json:"status" form:"status"`
}

type ShopCouponUpdateReq struct {
	ID          uint    `json:"id" form:"id" binding:"required"`
	Name        string  `json:"name" form:"name"`
	Type        int     `json:"type" form:"type"`
	Value       float64 `json:"value" form:"value"`
	MinAmount   float64 `json:"minAmount" form:"minAmount"`
	MaxDiscount float64 `json:"maxDiscount" form:"maxDiscount"`
	Scope       int     `json:"scope" form:"scope"`
	ScopeIDs    string  `json:"scopeIds" form:"scopeIds"`
	TotalCount  int     `json:"totalCount" form:"totalCount"`
	PerLimit    int     `json:"perLimit" form:"perLimit"`
	StartTime   string  `json:"startTime" form:"startTime"`
	EndTime     string  `json:"endTime" form:"endTime"`
	Status      int     `json:"status" form:"status"`
}

type ShopUserCouponSearch struct {
	Status *int `json:"status" form:"status"`
	request.PageInfo
}

type ShopCalcDiscountReq struct {
	CouponID uint   `json:"couponId" form:"couponId" binding:"required"`
	CartIDs  []uint `json:"cartIds" form:"cartIds"`
}

// ========== 退款 ==========

type ShopRefundCreateReq struct {
	OrderNo     string  `json:"orderNo" form:"orderNo" binding:"required"`
	Type        int     `json:"type" form:"type" binding:"required"`
	Reason      string  `json:"reason" form:"reason" binding:"required"`
	Description string  `json:"description" form:"description"`
	Images      string  `json:"images" form:"images"`
	Amount      float64 `json:"amount" form:"amount" binding:"required"`
}

type ShopRefundSearch struct {
	Status   *int   `json:"status" form:"status"`
	RefundNo string `json:"refundNo" form:"refundNo"`
	OrderNo  string `json:"orderNo" form:"orderNo"`
	request.PageInfo
}

type ShopRefundAuditReq struct {
	RefundNo    string `json:"refundNo" form:"refundNo" binding:"required"`
	Approved    bool   `json:"approved" form:"approved"`
	AdminRemark string `json:"adminRemark" form:"adminRemark"`
}

type ShopRefundShipReq struct {
	RefundNo    string `json:"refundNo" form:"refundNo" binding:"required"`
	ShipCompany string `json:"shipCompany" form:"shipCompany" binding:"required"`
	ShipNo      string `json:"shipNo" form:"shipNo" binding:"required"`
}

// ========== 收藏 ==========

type ShopFavoriteReq struct {
	SpuID uint `json:"spuId" form:"spuId" binding:"required"`
}

type ShopFavoriteSearch struct {
	request.PageInfo
}

type ShopBatchFavoriteReq struct {
	SpuIDs []uint `json:"spuIds" form:"spuIds" binding:"required"`
}

// ========== 物流 ==========

type ShopLogisticsSearch struct {
	OrderNo string `json:"orderNo" form:"orderNo"`
	request.PageInfo
}

type ShopLogisticsAddReq struct {
	OrderNo     string `json:"orderNo" form:"orderNo" binding:"required"`
	ShipNo      string `json:"shipNo" form:"shipNo"`
	ShipCompany string `json:"shipCompany" form:"shipCompany"`
	Detail      string `json:"detail" form:"detail" binding:"required"`
}

// ========== 库存预警 ==========

type ShopInventoryAlertSearch struct {
	Type   *int `json:"type" form:"type"`
	Status *int `json:"status" form:"status"`
	request.PageInfo
}

type ShopInventorySettingReq struct {
	SkuID          uint `json:"skuId" form:"skuId"`
	LowThreshold   int  `json:"lowThreshold" form:"lowThreshold"`
	EmptyThreshold int  `json:"emptyThreshold" form:"emptyThreshold"`
}

// ========== 推荐 ==========

type ShopRecommendationSearch struct {
	Type *int `json:"type" form:"type"`
	request.PageInfo
}

type ShopRecommendationSetReq struct {
	Type   int  `json:"type" form:"type" binding:"required"`
	SpuID  uint `json:"spuId" form:"spuId" binding:"required"`
	Sort   int  `json:"sort" form:"sort"`
	Status int  `json:"status" form:"status"`
}

type ShopRecommendationUpdateSortReq struct {
	ID   uint `json:"id" form:"id" binding:"required"`
	Sort int  `json:"sort" form:"sort"`
}

// ========== 数据分析 ==========

type ShopAnalyticsDateRange struct {
	StartDate string `json:"startDate" form:"startDate"`
	EndDate   string `json:"endDate" form:"endDate"`
	GroupBy   string `json:"groupBy" form:"groupBy"`
	Limit     int    `json:"limit" form:"limit"`
}
