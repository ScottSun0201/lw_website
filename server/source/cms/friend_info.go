package cms

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type initFriendInfo struct{}

const initOrderFriendInfo = initOrderFriend + 1

// auto run
func init() {
	system.RegisterInit(initOrderFriendInfo, &initFriendInfo{})
}

func (i initFriendInfo) InitializerName() string {
	return cms.CmsFriendInfo{}.TableName()
}

func (i *initFriendInfo) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&cms.CmsFriendInfo{})
}

func (i *initFriendInfo) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&cms.CmsFriendInfo{})
}

func (i *initFriendInfo) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	var jsonInfo = `{"desc": "Anim aute id magna aliqua ad ad non deserunt sunt. Qui irure qui lorem cupidatat commodo. Elit sunt amet fugiat veniam occaecat fugiat aliqua.", "links": [{"href": "#", "name": "Open roles"}, {"href": "#", "name": "Internship program"}, {"href": "#", "name": "Our values"}, {"href": "#", "name": "Meet our leadership"}], "stats": [{"name": "Offices worldwide", "value": "12"}, {"name": "Full-time colleagues", "value": "300+"}, {"name": "Hours per week", "value": "40"}, {"name": "Paid time off", "value": "Unlimited"}], "title": "Work with us", "friendDesc": "Libero fames augue nisl porttitor nisi, quis. Id ac elit odio vitae elementum enim vitae ullamcorper suspendisse.", "friendTitle": "Meet our leadership"}`

	jsonByte := []byte(jsonInfo)

	var j datatypes.JSON

	err := j.UnmarshalJSON(jsonByte)

	if err != nil {
		return ctx, errors.Wrap(err, cms.CmsFriendInfo{}.TableName()+"json解析失败")
	}

	entities := []cms.CmsFriendInfo{
		{FriendInfo: j},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, cms.CmsFriendInfo{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initFriendInfo) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&cms.CmsFriendInfo{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
