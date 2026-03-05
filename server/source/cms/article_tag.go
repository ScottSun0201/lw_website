package cms

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initArticleTag struct{}

const initOrderArticleTag = initOrderArticle + 1

// auto run
func init() {
	system.RegisterInit(initOrderArticleTag, &initArticleTag{})
}

func (i initArticleTag) InitializerName() string {
	return cms.CmsArticleTag{}.TableName()
}

func (i *initArticleTag) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&cms.CmsArticleTag{})
}

func (i *initArticleTag) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&cms.CmsArticleTag{})
}

func (i *initArticleTag) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []cms.CmsArticleTag{
		{CmsArticleID: 1, CmsTagID: 1},
		{CmsArticleID: 1, CmsTagID: 2},
		{CmsArticleID: 2, CmsTagID: 1},
		{CmsArticleID: 2, CmsTagID: 2},
		{CmsArticleID: 3, CmsTagID: 1},
		{CmsArticleID: 3, CmsTagID: 2},
		{CmsArticleID: 4, CmsTagID: 1},
		{CmsArticleID: 5, CmsTagID: 2},
		{CmsArticleID: 6, CmsTagID: 1},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, cms.CmsArticleTag{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initArticleTag) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&cms.CmsArticleTag{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
