package cms

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initCommit struct{}

const initOrderCommit = initOrderBanner + 1

// auto run
func init() {
	system.RegisterInit(initOrderCommit, &initCommit{})
}

func (i initCommit) InitializerName() string {
	return cms.CmsCommit{}.TableName()
}

func (i *initCommit) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&cms.CmsCommit{})
}

func (i *initCommit) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&cms.CmsCommit{})
}

func (i *initCommit) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []cms.CmsCommit{
		{Info: "测试一条评论", CmsArticleID: 1, UserID: 1},
		{Info: "测试第二条评论", CmsArticleID: 1, UserID: 1},
		{Info: "测试第三条评论", CmsArticleID: 2, UserID: 1},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, cms.CmsCommit{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initCommit) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&cms.CmsCommit{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
