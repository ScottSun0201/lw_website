package client

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initSEO struct{}

const initOrderSEO = initOrderAboutInfo + 1

// auto run
func init() {
	system.RegisterInit(initOrderSEO, &initSEO{})
}

func (i initSEO) InitializerName() string {
	return client.ClientSEO{}.TableName()
}

func (i *initSEO) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&client.ClientSEO{})
}

func (i *initSEO) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&client.ClientSEO{})
}

func (i *initSEO) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []client.ClientSEO{
		{Name: "GVA-CMS", Title: "支持SEO的CMS", Description: "基于gva和gva客户端框架搭建的支持SEO的CMS", Keywords: "keywords1,keywords2,keywords3,keywords4", Logo: "uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20240328212603.png"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, client.ClientSEO{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initSEO) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&client.ClientSEO{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
