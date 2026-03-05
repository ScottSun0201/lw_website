package client

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initAbout struct{}

const initOrderAbout = system.InitOrderClient + 1

// auto run
func init() {
	system.RegisterInit(initOrderAbout, &initAbout{})
}

func (i initAbout) InitializerName() string {
	return client.ClientAbout{}.TableName()
}

func (i *initAbout) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&client.ClientAbout{})
}

func (i *initAbout) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&client.ClientAbout{})
}

func (i *initAbout) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []client.ClientAbout{
		{Name: "Emma Dorsey", Title: "Senior Front-end Developer", Desc: "Ultricies massa malesuada viverra cras lobortis. Tempor orci hac ligula dapibus mauris sit ut eu. Eget turpis urna maecenas cras. Nisl dictum.", Avatar: "uploads/file/0769349b49ca0e303c703c7ae35db5ae_20240325213758.jpg"},
		{Name: "Emma Dorsey", Title: "Senior Front-end Developer", Desc: "Ultricies massa malesuada viverra cras lobortis. Tempor orci hac ligula dapibus mauris sit ut eu. Eget turpis urna maecenas cras. Nisl dictum.", Avatar: "uploads/file/0769349b49ca0e303c703c7ae35db5ae_20240325213758.jpg"},
		{Name: "Emma Dorsey", Title: "Senior Front-end Developer", Desc: "Ultricies massa malesuada viverra cras lobortis. Tempor orci hac ligula dapibus mauris sit ut eu. Eget turpis urna maecenas cras. Nisl dictum.", Avatar: "uploads/file/0769349b49ca0e303c703c7ae35db5ae_20240325213758.jpg"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, client.ClientAbout{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initAbout) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&client.ClientAbout{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
