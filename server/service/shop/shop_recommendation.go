package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
)

type ShopRecommendationService struct{}

func (s *ShopRecommendationService) SetRecommendation(req shopReq.ShopRecommendationSetReq) error {
	rec := shop.ShopRecommendation{Type: req.Type, SpuID: req.SpuID, Sort: req.Sort, Status: req.Status}
	return global.GVA_DB.Create(&rec).Error
}

func (s *ShopRecommendationService) RemoveRecommendation(id uint) error {
	return global.GVA_DB.Delete(&shop.ShopRecommendation{}, id).Error
}

func (s *ShopRecommendationService) GetRecommendationList(info shopReq.ShopRecommendationSearch) (list []shop.ShopRecommendation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopRecommendation{})
	if info.Type != nil {
		db = db.Where("type = ?", *info.Type)
	}
	err = db.Count(&total).Error
	if err != nil { return }
	if limit != 0 { db = db.Limit(limit).Offset(offset) }
	err = db.Order("sort asc, created_at desc").Find(&list).Error
	return
}

func (s *ShopRecommendationService) UpdateSort(id uint, sort int) error {
	return global.GVA_DB.Model(&shop.ShopRecommendation{}).Where("id = ?", id).Update("sort", sort).Error
}

func (s *ShopRecommendationService) GetHotProducts(limit int) (list []shop.ShopSpu, err error) {
	if limit <= 0 { limit = 10 }
	subQuery := global.GVA_DB.Model(&shop.ShopRecommendation{}).Select("spu_id").Where("type = 1 AND status = 1").Order("sort asc").Limit(limit)
	err = global.GVA_DB.Where("id IN (?) AND status = 1", subQuery).Preload("Skus").Find(&list).Error
	return
}

func (s *ShopRecommendationService) GetNewProducts(limit int) (list []shop.ShopSpu, err error) {
	if limit <= 0 { limit = 10 }
	subQuery := global.GVA_DB.Model(&shop.ShopRecommendation{}).Select("spu_id").Where("type = 2 AND status = 1").Order("sort asc").Limit(limit)
	err = global.GVA_DB.Where("id IN (?) AND status = 1", subQuery).Preload("Skus").Find(&list).Error
	return
}

func (s *ShopRecommendationService) GetRelatedProducts(spuID uint, limit int) (list []shop.ShopSpu, err error) {
	if limit <= 0 { limit = 8 }
	var spu shop.ShopSpu
	if err = global.GVA_DB.Where("id = ?", spuID).First(&spu).Error; err != nil {
		return
	}
	err = global.GVA_DB.Where("category_id = ? AND id != ? AND status = 1", spu.CategoryID, spuID).Limit(limit).Preload("Skus").Find(&list).Error
	return
}

func (s *ShopRecommendationService) RefreshHotProducts() {
	global.GVA_DB.Where("type = 1").Delete(&shop.ShopRecommendation{})
	var spuIDs []uint
	global.GVA_DB.Model(&shop.ShopSpu{}).Where("status = 1").Order("sales_count desc").Limit(20).Pluck("id", &spuIDs)
	for i, id := range spuIDs {
		rec := shop.ShopRecommendation{Type: 1, SpuID: id, Sort: i, Status: 1}
		global.GVA_DB.Create(&rec)
	}
}

func (s *ShopRecommendationService) RefreshNewProducts() {
	global.GVA_DB.Where("type = 2").Delete(&shop.ShopRecommendation{})
	var spuIDs []uint
	global.GVA_DB.Model(&shop.ShopSpu{}).Where("status = 1").Order("created_at desc").Limit(20).Pluck("id", &spuIDs)
	for i, id := range spuIDs {
		rec := shop.ShopRecommendation{Type: 2, SpuID: id, Sort: i, Status: 1}
		global.GVA_DB.Create(&rec)
	}
}
