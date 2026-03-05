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

	// Clean up test orders and related data
	tables := []string{
		"DELETE FROM shop_order_item",
		"DELETE FROM shop_order_log",
		"DELETE FROM shop_order",
		"DELETE FROM shop_payment",
		"DELETE FROM shop_wallet_log",
		"DELETE FROM shop_cart",
		// Reset inventory locks
		"UPDATE shop_inventory SET locked_stock = 0, available_stock = total_stock",
		// Reset wallet
		"UPDATE shop_user_wallet SET balance = 500, version = version + 1 WHERE user_id = 2",
	}

	for _, sql := range tables {
		result, err := db.Exec(sql)
		if err != nil {
			fmt.Printf("FAIL: %s: %v\n", sql, err)
		} else {
			rows, _ := result.RowsAffected()
			fmt.Printf("OK: %s (%d rows)\n", sql, rows)
		}
	}
	fmt.Println("清理完成!")
}
