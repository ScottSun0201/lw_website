package cms

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initBanner struct{}

const initOrderBanner = initOrderArticleTag + 1

// auto run
func init() {
	system.RegisterInit(initOrderBanner, &initBanner{})
}

func (i initBanner) InitializerName() string {
	return cms.CmsBanner{}.TableName()
}

func (i *initBanner) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&cms.CmsBanner{})
}

func (i *initBanner) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&cms.CmsBanner{})
}

func (i *initBanner) InitializeData(ctx context.Context) (context.Context, error) {
	_, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []cms.CmsBanner{}
	//if err := db.Create(&entities).Error; err != nil {
	//	return ctx, errors.Wrap(err, cms.CmsBanner{}.TableName()+"表数据初始化失败!")
	//}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initBanner) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&cms.CmsBanner{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
