package shop

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ShopSpu struct {
	global.GVA_MODEL
	Name       string       `json:"name" form:"name" gorm:"column:name;comment:商品名称" binding:"required"`
	SubTitle   string       `json:"subTitle" form:"subTitle" gorm:"column:sub_title;comment:副标题"`
	CategoryID uint         `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:分类ID"`
	BrandID    uint         `json:"brandId" form:"brandId" gorm:"column:brand_id;comment:品牌ID"`
	MainImage  string       `json:"mainImage" form:"mainImage" gorm:"column:main_image;comment:主图"`
	Images     string       `json:"images" form:"images" gorm:"column:images;comment:图片列表JSON;type:text"`
	Detail     string       `json:"detail" form:"detail" gorm:"column:detail;comment:商品详情;type:longtext"`
	Status     int          `json:"status" form:"status" gorm:"column:status;comment:状态0草稿1上架2下架;default:0"`
	Sort       int          `json:"sort" form:"sort" gorm:"column:sort;comment:排序;default:0"`
	SalesCount int          `json:"salesCount" form:"salesCount" gorm:"column:sales_count;comment:销量;default:0"`
	Category   ShopCategory `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	Brand      ShopBrand    `json:"brand" gorm:"foreignKey:BrandID;references:ID"`
	Skus       []ShopSku    `json:"skus" gorm:"foreignKey:SpuID;references:ID"`
	PriceRange string       `json:"priceRange" gorm:"-"`
}

func (ShopSpu) TableName() string {
	return "shop_spu"
}
