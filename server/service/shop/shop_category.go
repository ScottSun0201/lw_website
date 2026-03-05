package shop

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
)

type ShopCategoryService struct{}

// CreateShopCategory 创建商品分类
func (s *ShopCategoryService) CreateShopCategory(category *shop.ShopCategory) error {
	if category.ParentID > 0 {
		var parent shop.ShopCategory
		if err := global.GVA_DB.Where("id = ?", category.ParentID).First(&parent).Error; err != nil {
			return fmt.Errorf("父级分类不存在")
		}
		category.Level = parent.Level + 1
	} else {
		category.Level = 1
	}
	return global.GVA_DB.Create(category).Error
}

// DeleteShopCategory 删除商品分类
func (s *ShopCategoryService) DeleteShopCategory(ID string) error {
	var count int64
	global.GVA_DB.Model(&shop.ShopCategory{}).Where("parent_id = ?", ID).Count(&count)
	if count > 0 {
		return fmt.Errorf("该分类下有子分类，无法删除")
	}
	return global.GVA_DB.Delete(&shop.ShopCategory{}, "id = ?", ID).Error
}

// UpdateShopCategory 更新商品分类
func (s *ShopCategoryService) UpdateShopCategory(category shop.ShopCategory) error {
	return global.GVA_DB.Save(&category).Error
}

// GetShopCategory 根据ID获取商品分类
func (s *ShopCategoryService) GetShopCategory(ID string) (category shop.ShopCategory, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&category).Error
	return
}

// GetShopCategoryList 分页获取商品分类列表
func (s *ShopCategoryService) GetShopCategoryList(info shopReq.ShopCategorySearch) (list []shop.ShopCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&shop.ShopCategory{})

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	if info.ParentID != nil {
		db = db.Where("parent_id = ?", *info.ParentID)
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

// GetCategoryTree 获取分类树形结构
func (s *ShopCategoryService) GetCategoryTree() (tree []shop.ShopCategory, err error) {
	var allCategories []shop.ShopCategory
	err = global.GVA_DB.Order("sort asc").Find(&allCategories).Error
	if err != nil {
		return nil, err
	}

	// 按 parentID 分组
	childrenMap := make(map[uint][]shop.ShopCategory)
	for _, cat := range allCategories {
		childrenMap[cat.ParentID] = append(childrenMap[cat.ParentID], cat)
	}

	// 构建树：从顶级分类（parentID=0）开始
	tree = childrenMap[0]
	for i := range tree {
		s.buildChildren(&tree[i], childrenMap)
	}
	return
}

// buildChildren 递归构建子分类
func (s *ShopCategoryService) buildChildren(category *shop.ShopCategory, childrenMap map[uint][]shop.ShopCategory) {
	children, ok := childrenMap[category.ID]
	if !ok {
		return
	}
	category.Children = children
	for i := range category.Children {
		s.buildChildren(&category.Children[i], childrenMap)
	}
}
