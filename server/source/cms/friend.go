package cms

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initFriend struct{}

const initOrderFriend = initOrderCommit + 1

// auto run
func init() {
	system.RegisterInit(initOrderFriend, &initFriend{})
}

func (i initFriend) InitializerName() string {
	return cms.CmsFriend{}.TableName()
}

func (i *initFriend) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&cms.CmsFriend{})
}

func (i *initFriend) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&cms.CmsFriend{})
}

func (i *initFriend) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []cms.CmsFriend{
		{Title: "Leslie Alexander", Image: "uploads/file/0769349b49ca0e303c703c7ae35db5ae_20240325213758.jpg", Link: "https://www.gin-vue-admin.com", NewWindow: utils.Pointer(true), Open: utils.Pointer(true)},
		{Title: "Leslie Alexander", Image: "uploads/file/0769349b49ca0e303c703c7ae35db5ae_20240325213758.jpg", Link: "https://www.gin-vue-admin.com", NewWindow: utils.Pointer(true), Open: utils.Pointer(true)},
		{Title: "Leslie Alexander", Image: "uploads/file/0769349b49ca0e303c703c7ae35db5ae_20240325213758.jpg", Link: "https://www.gin-vue-admin.com", NewWindow: utils.Pointer(true), Open: utils.Pointer(true)},
		{Title: "Leslie Alexander", Image: "uploads/file/0769349b49ca0e303c703c7ae35db5ae_20240325213758.jpg", Link: "https://www.gin-vue-admin.com", NewWindow: utils.Pointer(true), Open: utils.Pointer(true)},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, cms.CmsFriend{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initFriend) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&cms.CmsFriend{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
