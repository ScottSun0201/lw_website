// +build ignore

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "lw_website:CXwKNa2NcA8TN7rS@tcp(192.168.100.164:3306)/lw_website?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tables := []string{"sys_base_menus", "sys_users", "sys_authorities", "sys_apis"}
	for _, t := range tables {
		var count int
		db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", t)).Scan(&count)
		fmt.Printf("%-25s: %d 条记录\n", t, count)
	}

	// 检查 shop 表
	shopTables := []string{"shop_spu", "shop_sku", "shop_category", "shop_brand", "shop_order", "shop_inventory"}
	fmt.Println()
	for _, t := range shopTables {
		var count int
		db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", t)).Scan(&count)
		fmt.Printf("%-25s: %d 条记录\n", t, count)
	}
}
