package client

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initClientUser struct{}

const initOrderClientUser = initOrderSEO + 1

// auto run
func init() {
	system.RegisterInit(initOrderClientUser, &initClientUser{})
}

func (i initClientUser) InitializerName() string {
	return client.ClientUser{}.TableName()
}

func (i *initClientUser) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&client.ClientUser{})
}

func (i *initClientUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&client.ClientUser{})
}

func (i *initClientUser) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	u, _ := uuid.NewV4()
	entities := []client.ClientUser{
		{Username: "a303176530", Password: "$2a$10$7XykKHy3Zy2RAKN.wptxdule3/k23XjOPwV/vuwRoj2JUKddyub56", Nickname: "NewUser—3d", Avatar: "uploads/file/96d0f363087ac4d233b64eb0c65ef59d_20240323200351.png", Age: utils.Pointer(18), Gender: "1", About: "so smart", FirstName: "miao", LastName: "qi", Email: "xxxx@qq.com", UUID: u},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, client.ClientUser{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initClientUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&client.ClientUser{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
