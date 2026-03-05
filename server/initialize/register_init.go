package initialize

import (
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/client"
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/cms"
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/example"
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
