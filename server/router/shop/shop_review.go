package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopReviewRouter struct{}

func (s *ShopReviewRouter) InitShopReviewRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	shopReviewRouter := Router.Group("shopReview").Use(middleware.OperationRecord())
	var api = v1.ApiGroupApp.ShopApiGroup.ShopReviewApi
	{
		shopReviewRouter.GET("getReviewList", api.GetReviewList)
		shopReviewRouter.PUT("auditReview", api.AuditReview)
		shopReviewRouter.PUT("replyReview", api.ReplyReview)
	}

	shopReviewPublicRouter := PublicRouter.Group("shopReview")
	{
		shopReviewPublicRouter.POST("createReview", api.CreateReview)
		shopReviewPublicRouter.GET("getProductReviews", api.GetProductReviews)
		shopReviewPublicRouter.GET("getReviewStat", api.GetReviewStat)
		shopReviewPublicRouter.GET("getMyReviews", api.GetMyReviews)
	}
}
