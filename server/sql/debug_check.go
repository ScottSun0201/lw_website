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

	fmt.Println("=== shop_inventory ===")
	rows, _ := db.Query("SELECT id, sku_id, total_stock, locked_stock, available_stock, version FROM shop_inventory")
	defer rows.Close()
	for rows.Next() {
		var id, skuID, total, locked, available, version int
		rows.Scan(&id, &skuID, &total, &locked, &available, &version)
		fmt.Printf("  id=%d sku_id=%d total=%d locked=%d available=%d version=%d\n", id, skuID, total, locked, available, version)
	}

	fmt.Println("\n=== shop_cart ===")
	rows2, _ := db.Query("SELECT id, user_id, sku_id, quantity, selected FROM shop_cart")
	defer rows2.Close()
	for rows2.Next() {
		var id, userID, skuID, qty, selected int
		rows2.Scan(&id, &userID, &skuID, &qty, &selected)
		fmt.Printf("  id=%d user_id=%d sku_id=%d qty=%d selected=%d\n", id, userID, skuID, qty, selected)
	}

	fmt.Println("\n=== shop_user_wallet ===")
	rows3, _ := db.Query("SELECT id, user_id, balance, version FROM shop_user_wallet")
	defer rows3.Close()
	for rows3.Next() {
		var id, userID, version int
		var balance float64
		rows3.Scan(&id, &userID, &balance, &version)
		fmt.Printf("  id=%d user_id=%d balance=%.2f version=%d\n", id, userID, balance, version)
	}

	fmt.Println("\n=== shop_user_address ===")
	rows4, _ := db.Query("SELECT id, user_id, receiver_name, is_default FROM shop_user_address")
	defer rows4.Close()
	for rows4.Next() {
		var id, userID, isDefault int
		var name string
		rows4.Scan(&id, &userID, &name, &isDefault)
		fmt.Printf("  id=%d user_id=%d name=%s is_default=%d\n", id, userID, name, isDefault)
	}

	fmt.Println("\n=== client_user ===")
	rows5, _ := db.Query("SELECT id, username, uuid FROM client_user LIMIT 5")
	defer rows5.Close()
	for rows5.Next() {
		var id int
		var username, uuid string
		rows5.Scan(&id, &username, &uuid)
		fmt.Printf("  id=%d username=%s uuid=%s\n", id, username, uuid)
	}
}
