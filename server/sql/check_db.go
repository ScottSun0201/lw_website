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
		log.Fatal("连接数据库失败:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("数据库不可达:", err)
	}
	fmt.Println("数据库连接成功")

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal("查询失败:", err)
	}
	defer rows.Close()

	fmt.Println("\n现有表:")
	count := 0
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Printf("  - %s\n", name)
		count++
	}
	if count == 0 {
		fmt.Println("  (空数据库，无任何表)")
	} else {
		fmt.Printf("\n共 %d 张表\n", count)
	}
}
