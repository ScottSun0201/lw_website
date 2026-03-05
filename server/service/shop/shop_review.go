package shop

import (
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"gorm.io/gorm"
)

type ShopReviewService struct{}

// CreateReview 创建评价
func (s *ShopReviewService) CreateReview(userID uint, req shopReq.ShopReviewCreateReq) (review shop.ShopReview, err error) {
	// 校验订单归属+状态>=3
	var order shop.ShopOrder
	if err = global.GVA_DB.Where("order_no = ? AND user_id = ?", req.OrderNo, userID).First(&order).Error; err != nil {
		return review, fmt.Errorf("订单不存在")
	}
	if order.Status < 3 {
		return review, fmt.Errorf("订单状态不允许评价")
	}

	// 校验未重复评价
	var count int64
	global.GVA_DB.Model(&shop.ShopReview{}).Where("order_no = ?", req.OrderNo).Count(&count)
	if count > 0 {
		return review, fmt.Errorf("该订单已评价")
	}

	// 获取订单第一个商品的spuID和skuID
	var item shop.ShopOrderItem
	if err = global.GVA_DB.Where("order_no = ?", req.OrderNo).First(&item).Error; err != nil {
		return review, fmt.Errorf("获取订单商品失败")
	}

	review = shop.ShopReview{
		UserID:  userID,
		SpuID:   item.SpuID,
		SkuID:   item.SkuID,
		OrderNo: req.OrderNo,
		Rating:  req.Rating,
		Content: req.Content,
		Images:  req.Images,
		Status:  0,
		IsAnon:  req.IsAnon,
	}

	if err = global.GVA_DB.Create(&review).Error; err != nil {
		return review, err
	}

	// 更新评价统计（包含待审核的也计入统计，审核拒绝时重新计算）
	go s.recalcStats(item.SpuID)

	return review, nil
}

// GetProductReviews 获取商品评价列表（仅通过审核的）
func (s *ShopReviewService) GetProductReviews(spuID uint, info shopReq.ShopReviewSearch) (list []shop.ShopReview, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopReview{}).Where("spu_id = ? AND status = 1", spuID)

	if info.Rating != nil {
		db = db.Where("rating = ?", *info.Rating)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&list).Error
	if err != nil {
		return
	}

	// 填充用户昵称头像
	for i := range list {
		if list[i].IsAnon == 1 {
			list[i].Nickname = "匿名用户"
			list[i].Avatar = ""
		} else {
			var user client.ClientUser
			if global.GVA_DB.Select("nickname, avatar").Where("id = ?", list[i].UserID).First(&user).Error == nil {
				list[i].Nickname = user.Nickname
				list[i].Avatar = user.Avatar
			}
		}
	}
	return
}

// GetReviewStat 获取商品评价统计
func (s *ShopReviewService) GetReviewStat(spuID uint) (stat shop.ShopReviewStat, err error) {
	err = global.GVA_DB.Where("spu_id = ?", spuID).First(&stat).Error
	if err == gorm.ErrRecordNotFound {
		return shop.ShopReviewStat{SpuID: spuID}, nil
	}
	return
}

// GetUserReviews 获取用户评价列表
func (s *ShopReviewService) GetUserReviews(userID uint, info shopReq.ShopReviewSearch) (list []shop.ShopReview, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopReview{}).Where("user_id = ?", userID)

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&list).Error
	return
}

// GetReviewList 管理端获取评价列表
func (s *ShopReviewService) GetReviewList(info shopReq.ShopReviewSearch) (list []shop.ShopReview, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopReview{})

	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	if info.SpuID != nil {
		db = db.Where("spu_id = ?", *info.SpuID)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&list).Error
	if err != nil {
		return
	}

	// 填充用户信息
	for i := range list {
		var user client.ClientUser
		if global.GVA_DB.Select("nickname, avatar").Where("id = ?", list[i].UserID).First(&user).Error == nil {
			list[i].Nickname = user.Nickname
			list[i].Avatar = user.Avatar
		}
	}
	return
}

// AuditReview 审核评价
func (s *ShopReviewService) AuditReview(id uint, status int) error {
	var review shop.ShopReview
	if err := global.GVA_DB.Where("id = ?", id).First(&review).Error; err != nil {
		return fmt.Errorf("评价不存在")
	}
	if err := global.GVA_DB.Model(&shop.ShopReview{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	go s.recalcStats(review.SpuID)
	return nil
}

// ReplyReview 回复评价
func (s *ShopReviewService) ReplyReview(id uint, reply string) error {
	now := time.Now()
	return global.GVA_DB.Model(&shop.ShopReview{}).Where("id = ?", id).Updates(map[string]interface{}{
		"reply":    reply,
		"reply_at": &now,
	}).Error
}

// recalcStats 重算评价统计
func (s *ShopReviewService) recalcStats(spuID uint) {
	var stat shop.ShopReviewStat

	// 只统计审核通过的
	var totalCount int64
	global.GVA_DB.Model(&shop.ShopReview{}).Where("spu_id = ? AND status = 1", spuID).Count(&totalCount)

	var avgRating float64
	global.GVA_DB.Model(&shop.ShopReview{}).Where("spu_id = ? AND status = 1", spuID).Select("COALESCE(AVG(rating), 0)").Scan(&avgRating)

	var star1, star2, star3, star4, star5 int64
	global.GVA_DB.Model(&shop.ShopReview{}).Where("spu_id = ? AND status = 1 AND rating = 1", spuID).Count(&star1)
	global.GVA_DB.Model(&shop.ShopReview{}).Where("spu_id = ? AND status = 1 AND rating = 2", spuID).Count(&star2)
	global.GVA_DB.Model(&shop.ShopReview{}).Where("spu_id = ? AND status = 1 AND rating = 3", spuID).Count(&star3)
	global.GVA_DB.Model(&shop.ShopReview{}).Where("spu_id = ? AND status = 1 AND rating = 4", spuID).Count(&star4)
	global.GVA_DB.Model(&shop.ShopReview{}).Where("spu_id = ? AND status = 1 AND rating = 5", spuID).Count(&star5)

	err := global.GVA_DB.Where("spu_id = ?", spuID).First(&stat).Error
	if err == gorm.ErrRecordNotFound {
		stat = shop.ShopReviewStat{
			SpuID:      spuID,
			TotalCount: int(totalCount),
			AvgRating:  avgRating,
			Star1Count: int(star1),
			Star2Count: int(star2),
			Star3Count: int(star3),
			Star4Count: int(star4),
			Star5Count: int(star5),
		}
		global.GVA_DB.Create(&stat)
	} else {
		global.GVA_DB.Model(&stat).Updates(map[string]interface{}{
			"total_count": totalCount,
			"avg_rating":  avgRating,
			"star1_count": star1,
			"star2_count": star2,
			"star3_count": star3,
			"star4_count": star4,
			"star5_count": star5,
		})
	}
}
