package client

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type initAboutInfo struct{}

const initOrderAboutInfo = initOrderAbout + 1

// auto run
func init() {
	system.RegisterInit(initOrderAboutInfo, &initAboutInfo{})
}

func (i initAboutInfo) InitializerName() string {
	return client.ClientAboutInfo{}.TableName()
}

func (i *initAboutInfo) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&client.ClientAboutInfo{})
}

func (i *initAboutInfo) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&client.ClientAboutInfo{})
}

func (i *initAboutInfo) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	var jsonInfo = `{"desc": "Mattis amet hendrerit dolor, quisque lorem pharetra. Pellentesque lacus nisi urna, arcu sociis eu. Orci vel lectus nisl eget eget ut consectetur. Sit justo viverra non adipisicing elit distinctio.", "title": "Get in touch", "teamDesc": "Nulla quam felis, enim faucibus proin velit, ornare id pretium. Augue ultrices sed arcu condimentum vestibulum suspendisse. Volutpat eu faucibus vivamus eget bibendum cras.", "teamTitle": "Our Team"}`

	jsonByte := []byte(jsonInfo)

	var j datatypes.JSON

	err := j.UnmarshalJSON(jsonByte)
	if err != nil {
		return ctx, errors.Wrap(err, client.ClientAboutInfo{}.TableName()+"json解析失败")
	}

	entities := []client.ClientAboutInfo{
		{Info: j},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, client.ClientAboutInfo{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initAboutInfo) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&client.ClientAboutInfo{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
