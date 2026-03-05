package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CmsArticleSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Title  string `json:"title" form:"title" `
	Tag    string `json:"tag" form:"tag" `
	Author uint   `json:"author" form:"author"`
	Pid    *int   `json:"pid" form:"pid" `
	request.PageInfo
}
