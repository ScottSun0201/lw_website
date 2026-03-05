package cms

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initArticle struct{}

const initOrderArticle = system.InitOrderCms + 1

var Content = "<p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusan<span style=\"color: rgb(225, 60, 57);\">tium praesentium eius, ut atque fuga cu</span>lpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.</p><p>Lorem ip<em>sum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eiu</em>s, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accus<u>antium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.</u></p><p><strong>Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.</strong></p><p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.</p><p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum.</p>"
var Desc = "Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto accusantium praesentium eius, ut atque fuga culpa, similique sequi cum eos quis dolorum."

// auto run
func init() {
	system.RegisterInit(initOrderArticle, &initArticle{})
}

func (i initArticle) InitializerName() string {
	return cms.CmsArticle{}.TableName()
}

func (i *initArticle) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&cms.CmsArticle{})
}

func (i *initArticle) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&cms.CmsArticle{})
}

func (i *initArticle) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []cms.CmsArticle{
		{Title: "Boost your conversion rate", Content: Content, Desc: Desc, Image: "uploads/file/bb1e2ee645e6579d4dfd1f6eea33d4cb_20240323210824.png", Author: 1, ReadingTime: 10},
		{Title: "How to use search engine optimization to drive sales", Content: Content, Desc: Desc, Image: "uploads/file/80d35b6a3b9e15a7d544ec778fcb1009_20240323210822.png", Author: 1, ReadingTime: 10},
		{Title: "Improve your customer experience", Content: Content, Desc: Desc, Image: "uploads/file/4318cb63533860e83296cd8cf09f6281_20240323210820.png", Author: 1, ReadingTime: 10},
		{Title: "Boost your conversion rate", Content: Content, Desc: Desc, Image: "uploads/file/bb1e2ee645e6579d4dfd1f6eea33d4cb_20240323210824.png", Author: 1, ReadingTime: 10},
		{Title: "Boost your conversion rate", Content: Content, Desc: Desc, Image: "uploads/file/bb1e2ee645e6579d4dfd1f6eea33d4cb_20240323210824.png", Author: 1, ReadingTime: 10},
		{Title: "optimization ", Content: Content, Desc: Desc, Image: "uploads/file/80d35b6a3b9e15a7d544ec778fcb1009_20240323210822.png", Author: 1, ReadingTime: 10},
		{Title: "Improve your customer experience", Content: Content, Desc: Desc, Image: "uploads/file/4318cb63533860e83296cd8cf09f6281_20240323210820.png", Author: 1, ReadingTime: 10},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, cms.CmsArticle{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initArticle) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&cms.CmsArticle{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
