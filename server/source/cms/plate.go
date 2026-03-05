package cms

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initPlate struct{}

const initOrderPlate = initOrderFriendInfo + 1

// auto run
func init() {
	system.RegisterInit(initOrderPlate, &initPlate{})
}

func (i initPlate) InitializerName() string {
	return cms.CmsPlate{}.TableName()
}

func (i *initPlate) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&cms.CmsPlate{})
}

func (i *initPlate) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&cms.CmsPlate{})
}

func (i *initPlate) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []cms.CmsPlate{
		{Pid: 0, Name: "Tag", Desc: "this is tag", Icon: "", Router: "/blog/tags", Sort: 10, Open: utils.Pointer(true)},
		{Pid: 0, Name: "Blog", Desc: "Blog", Icon: "", Router: "/", Sort: 0, Open: utils.Pointer(true)},
		{Pid: 0, Name: "Friends", Desc: "Friends", Icon: "", Router: "/friends", Sort: 20, Open: utils.Pointer(true)},
		{Pid: 0, Name: "About Us", Desc: "About Us", Icon: "", Router: "/about", Sort: 30, Open: utils.Pointer(true)},
		{Pid: 2, Name: "Trending", Desc: "Get a better understanding of your traffic", Icon: "ChartPieIcon", Router: "/blog", Sort: 0, Open: utils.Pointer(true)},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, cms.CmsPlate{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initPlate) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&cms.CmsPlate{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
