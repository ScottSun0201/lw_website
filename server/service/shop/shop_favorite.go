package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
)

type ShopFavoriteService struct{}

func (s *ShopFavoriteService) AddFavorite(userID, spuID uint) error {
	fav := shop.ShopFavorite{UserID: userID, SpuID: spuID}
	return global.GVA_DB.FirstOrCreate(&fav, shop.ShopFavorite{UserID: userID, SpuID: spuID}).Error
}

func (s *ShopFavoriteService) RemoveFavorite(userID, spuID uint) error {
	return global.GVA_DB.Where("user_id = ? AND spu_id = ?", userID, spuID).Delete(&shop.ShopFavorite{}).Error
}

func (s *ShopFavoriteService) GetFavoriteList(userID uint, info shopReq.ShopFavoriteSearch) (list []shop.ShopSpu, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	subQuery := global.GVA_DB.Model(&shop.ShopFavorite{}).Select("spu_id").Where("user_id = ?", userID)
	db := global.GVA_DB.Model(&shop.ShopSpu{}).Where("id IN (?)", subQuery)
	err = db.Count(&total).Error
	if err != nil { return }
	if limit != 0 { db = db.Limit(limit).Offset(offset) }
	err = db.Order("id desc").Find(&list).Error
	return
}

func (s *ShopFavoriteService) IsFavorite(userID, spuID uint) bool {
	var count int64
	global.GVA_DB.Model(&shop.ShopFavorite{}).Where("user_id = ? AND spu_id = ?", userID, spuID).Count(&count)
	return count > 0
}

func (s *ShopFavoriteService) GetFavoriteCount(userID uint) (count int64) {
	global.GVA_DB.Model(&shop.ShopFavorite{}).Where("user_id = ?", userID).Count(&count)
	return
}

func (s *ShopFavoriteService) BatchCheckFavorite(userID uint, spuIDs []uint) map[uint]bool {
	result := make(map[uint]bool)
	if len(spuIDs) == 0 { return result }
	var favs []shop.ShopFavorite
	global.GVA_DB.Where("user_id = ? AND spu_id IN ?", userID, spuIDs).Find(&favs)
	for _, f := range favs {
		result[f.SpuID] = true
	}
	return result
}
