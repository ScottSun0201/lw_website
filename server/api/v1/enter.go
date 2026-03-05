package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/client"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/cms"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	ClientApiGroup  client.ApiGroup
	CmsApiGroup     cms.ApiGroup
	ShopApiGroup    shop.ApiGroup
}
