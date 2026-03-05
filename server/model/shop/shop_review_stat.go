package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopReviewStat struct {
	global.GVA_MODEL
	SpuID      uint    `json:"spuId" form:"spuId" gorm:"column:spu_id;comment:商品ID;uniqueIndex"`
	TotalCount int     `json:"totalCount" form:"totalCount" gorm:"column:total_count;comment:总评价数;default:0"`
	AvgRating  float64 `json:"avgRating" form:"avgRating" gorm:"column:avg_rating;comment:平均评分;type:decimal(3,2);default:0"`
	Star1Count int     `json:"star1Count" form:"star1Count" gorm:"column:star1_count;comment:1星数;default:0"`
	Star2Count int     `json:"star2Count" form:"star2Count" gorm:"column:star2_count;comment:2星数;default:0"`
	Star3Count int     `json:"star3Count" form:"star3Count" gorm:"column:star3_count;comment:3星数;default:0"`
	Star4Count int     `json:"star4Count" form:"star4Count" gorm:"column:star4_count;comment:4星数;default:0"`
	Star5Count int     `json:"star5Count" form:"star5Count" gorm:"column:star5_count;comment:5星数;default:0"`
}

func (ShopReviewStat) TableName() string {
	return "shop_review_stat"
}
