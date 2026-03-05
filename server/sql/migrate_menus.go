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

	// 检查是否已经存在商城菜单
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM sys_base_menus WHERE path = 'shopAdmin'").Scan(&count)
	if err != nil {
		log.Fatal("查询失败:", err)
	}
	if count > 0 {
		fmt.Println("商城菜单已存在，跳过初始化")
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("开启事务失败:", err)
	}

	// 1. 顶级菜单：商城管理
	res, err := tx.Exec(`INSERT INTO sys_base_menus (created_at, updated_at, menu_level, parent_id, path, name, hidden, component, sort, title, icon, keep_alive, close_tab, default_menu, active_name)
		VALUES (NOW(), NOW(), 0, 0, 'shopAdmin', 'shopAdmin', 0, 'view/routerHolder.vue', 900, '商城管理', 'shop', 0, 0, 0, '')`)
	if err != nil {
		tx.Rollback()
		log.Fatal("插入顶级菜单失败:", err)
	}
	parentID, _ := res.LastInsertId()
	fmt.Printf("顶级菜单 '商城管理' 插入成功, ID=%d\n", parentID)

	// 子菜单定义
	menus := []struct {
		path, name, component, title, icon string
		sort                               int
		hidden                             int
		closeTab                           int
		activeName                         string
	}{
		{"shopCategory", "shopCategory", "view/shopCategory/shopCategory.vue", "商品分类", "grid", 0, 0, 0, ""},
		{"shopBrand", "shopBrand", "view/shopBrand/shopBrand.vue", "品牌管理", "stamp", 1, 0, 0, ""},
		{"shopProduct", "shopProduct", "view/shopProduct/shopProduct.vue", "商品管理", "goods", 2, 0, 0, ""},
		{"shopProductForm", "shopProductForm", "view/shopProduct/shopProductForm.vue", "商品编辑", "goods", 3, 1, 1, "shopProduct"},
		{"shopInventory", "shopInventory", "view/shopInventory/shopInventory.vue", "库存管理", "box", 4, 0, 0, ""},
		{"shopOrder", "shopOrder", "view/shopOrder/shopOrder.vue", "订单管理", "document", 5, 0, 0, ""},
	}

	menuIDs := []int64{parentID}

	for _, m := range menus {
		r, err := tx.Exec(`INSERT INTO sys_base_menus (created_at, updated_at, menu_level, parent_id, path, name, hidden, component, sort, title, icon, keep_alive, close_tab, default_menu, active_name)
			VALUES (NOW(), NOW(), 0, ?, ?, ?, ?, ?, ?, ?, ?, 0, ?, 0, ?)`,
			parentID, m.path, m.name, m.hidden, m.component, m.sort, m.title, m.icon, m.closeTab, m.activeName)
		if err != nil {
			tx.Rollback()
			log.Fatalf("插入菜单 '%s' 失败: %v", m.title, err)
		}
		id, _ := r.LastInsertId()
		menuIDs = append(menuIDs, id)
		fmt.Printf("  子菜单 '%s' 插入成功, ID=%d\n", m.title, id)
	}

	// 为超级管理员角色(888)分配所有商城菜单
	for _, menuID := range menuIDs {
		_, err := tx.Exec(`INSERT INTO sys_authority_menus (sys_base_menu_id, sys_authority_authority_id) VALUES (?, '888')`, menuID)
		if err != nil {
			tx.Rollback()
			log.Fatalf("分配菜单权限失败(menuID=%d): %v", menuID, err)
		}
	}
	fmt.Printf("已为角色888分配 %d 个菜单权限\n", len(menuIDs))

	if err := tx.Commit(); err != nil {
		log.Fatal("提交事务失败:", err)
	}

	fmt.Println("商城菜单初始化完成!")
}
