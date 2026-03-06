package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopReviewRouter struct{}

func (s *ShopReviewRouter) InitShopReviewRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, ClientPrivateRouter *gin.RouterGroup) {
	shopReviewRouter := Router.Group("shopReview").Use(middleware.OperationRecord())
	var api = v1.ApiGroupApp.ShopApiGroup.ShopReviewApi
	{
		shopReviewRouter.GET("getReviewList", api.GetReviewList)
		shopReviewRouter.PUT("auditReview", api.AuditReview)
		shopReviewRouter.PUT("replyReview", api.ReplyReview)
	}

	// 公开路由（无需登录即可查看评价）
	shopReviewPublicRouter := PublicRouter.Group("shopReview")
	{
		shopReviewPublicRouter.GET("getProductReviews", api.GetProductReviews)
		shopReviewPublicRouter.GET("getReviewStat", api.GetReviewStat)
	}

	// 客户端路由（需要 JWT 认证）
	shopReviewClientRouter := ClientPrivateRouter.Group("shopReview")
	{
		shopReviewClientRouter.POST("createReview", api.CreateReview)
		shopReviewClientRouter.GET("getMyReviews", api.GetMyReviews)
	}
}
