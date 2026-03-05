package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
)

type ShopBrandService struct{}

// CreateShopBrand 创建品牌
func (s *ShopBrandService) CreateShopBrand(brand *shop.ShopBrand) error {
	return global.GVA_DB.Create(brand).Error
}

// DeleteShopBrand 删除品牌
func (s *ShopBrandService) DeleteShopBrand(ID string) error {
	return global.GVA_DB.Delete(&shop.ShopBrand{}, "id = ?", ID).Error
}

// UpdateShopBrand 更新品牌
func (s *ShopBrandService) UpdateShopBrand(brand shop.ShopBrand) error {
	return global.GVA_DB.Save(&brand).Error
}

// GetShopBrand 根据ID获取品牌
func (s *ShopBrandService) GetShopBrand(ID string) (brand shop.ShopBrand, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&brand).Error
	return
}

// GetShopBrandList 分页获取品牌列表
func (s *ShopBrandService) GetShopBrandList(info shopReq.ShopBrandSearch) (list []shop.ShopBrand, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopBrand{})

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("sort asc").Find(&list).Error
	return
}

// GetAllBrands 获取所有启用的品牌（不分页）
func (s *ShopBrandService) GetAllBrands() (list []shop.ShopBrand, err error) {
	err = global.GVA_DB.Where("status = ?", 1).Order("sort asc").Find(&list).Error
	return
}
