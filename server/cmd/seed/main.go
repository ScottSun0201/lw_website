// +build ignore

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"

	// Trigger source package init() registrations
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/client"
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/cms"
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/example"
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/system"
)

func main() {
	// Must run from server/ directory
	if _, err := os.Stat("config.yaml"); err != nil {
		log.Fatal("请在 server/ 目录下运行此程序")
	}

	// 1. Initialize Viper (loads config.yaml, sets global.GVA_CONFIG)
	global.GVA_VP = core.Viper()
	fmt.Println("Viper 初始化完成")

	// 2. Call InitDB with MySQL credentials
	conf := request.InitDB{
		AdminPassword: "123456",
		DBType:        "mysql",
		Host:          "192.168.100.164",
		Port:          "3306",
		UserName:      "lw_website",
		Password:      "CXwKNa2NcA8TN7rS",
		DBName:        "lw_website",
	}

	svc := &system.InitDBService{}
	fmt.Println("开始初始化数据库...")
	if err := svc.InitDB(conf); err != nil {
		log.Fatalf("InitDB 失败: %v", err)
	}

	fmt.Println("数据库初始化完成! 管理员账号: admin, 密码: 123456")
}
