// +build ignore

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type apiDef struct {
	path, method, group, desc string
}

func main() {
	dsn := "lw_website:CXwKNa2NcA8TN7rS@tcp(192.168.100.164:3306)/lw_website?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check idempotency
	var count int
	db.QueryRow("SELECT COUNT(*) FROM sys_apis WHERE api_group = 'shopCategory'").Scan(&count)
	if count > 0 {
		fmt.Println("商城API已存在，跳过")
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// All shop APIs (private + public)
	apis := []apiDef{
		// shopCategory
		{"/shopCategory/createShopCategory", "POST", "shopCategory", "创建商品分类"},
		{"/shopCategory/deleteShopCategory", "DELETE", "shopCategory", "删除商品分类"},
		{"/shopCategory/updateShopCategory", "PUT", "shopCategory", "更新商品分类"},
		{"/shopCategory/findShopCategory", "GET", "shopCategory", "查询商品分类"},
		{"/shopCategory/getShopCategoryList", "GET", "shopCategory", "获取分类列表"},
		{"/shopCategory/getCategoryTree", "GET", "shopCategory", "获取分类树"},
		// shopBrand
		{"/shopBrand/createShopBrand", "POST", "shopBrand", "创建品牌"},
		{"/shopBrand/deleteShopBrand", "DELETE", "shopBrand", "删除品牌"},
		{"/shopBrand/updateShopBrand", "PUT", "shopBrand", "更新品牌"},
		{"/shopBrand/findShopBrand", "GET", "shopBrand", "查询品牌"},
		{"/shopBrand/getShopBrandList", "GET", "shopBrand", "获取品牌列表"},
		{"/shopBrand/getAllBrands", "GET", "shopBrand", "获取全部品牌"},
		// shopProduct
		{"/shopProduct/createSpu", "POST", "shopProduct", "创建商品"},
		{"/shopProduct/updateSpu", "PUT", "shopProduct", "更新商品"},
		{"/shopProduct/deleteSpu", "DELETE", "shopProduct", "删除商品"},
		{"/shopProduct/getSpu", "GET", "shopProduct", "查询商品详情"},
		{"/shopProduct/getSpuList", "GET", "shopProduct", "获取商品列表"},
		{"/shopProduct/updateSpuStatus", "PUT", "shopProduct", "更新商品状态"},
		{"/shopProduct/getClientProductList", "GET", "shopProduct", "客户端商品列表"},
		{"/shopProduct/getClientProductDetail", "GET", "shopProduct", "客户端商品详情"},
		// shopInventory
		{"/shopInventory/setStock", "PUT", "shopInventory", "设置库存"},
		{"/shopInventory/getInventoryList", "GET", "shopInventory", "获取库存列表"},
		{"/shopInventory/getInventoryLogList", "GET", "shopInventory", "获取库存日志"},
		// shopAddress
		{"/shopAddress/createAddress", "POST", "shopAddress", "创建收货地址"},
		{"/shopAddress/deleteAddress", "DELETE", "shopAddress", "删除收货地址"},
		{"/shopAddress/updateAddress", "PUT", "shopAddress", "更新收货地址"},
		{"/shopAddress/getAddressList", "GET", "shopAddress", "获取地址列表"},
		{"/shopAddress/setDefaultAddress", "PUT", "shopAddress", "设置默认地址"},
		// shopCart
		{"/shopCart/addToCart", "POST", "shopCart", "加入购物车"},
		{"/shopCart/updateCartQuantity", "PUT", "shopCart", "更新购物车数量"},
		{"/shopCart/removeFromCart", "POST", "shopCart", "删除购物车商品"},
		{"/shopCart/getCartList", "GET", "shopCart", "获取购物车列表"},
		{"/shopCart/getCartCount", "GET", "shopCart", "获取购物车数量"},
		// shopOrder
		{"/shopOrder/createOrder", "POST", "shopOrder", "创建订单"},
		{"/shopOrder/cancelOrder", "PUT", "shopOrder", "取消订单"},
		{"/shopOrder/confirmReceive", "PUT", "shopOrder", "确认收货"},
		{"/shopOrder/getUserOrderList", "GET", "shopOrder", "用户订单列表"},
		{"/shopOrder/getUserOrder", "GET", "shopOrder", "用户订单详情"},
		{"/shopOrder/getOrderList", "GET", "shopOrder", "管理端订单列表"},
		{"/shopOrder/getOrderDetail", "GET", "shopOrder", "管理端订单详情"},
		{"/shopOrder/shipOrder", "PUT", "shopOrder", "订单发货"},
		{"/shopOrder/adminCancelOrder", "PUT", "shopOrder", "管理端取消订单"},
		{"/shopOrder/getOrderLogs", "GET", "shopOrder", "获取订单日志"},
		// shopPayment
		{"/shopPayment/balancePay", "POST", "shopPayment", "余额支付"},
		{"/shopPayment/getWallet", "GET", "shopPayment", "获取钱包"},
		{"/shopPayment/adminRechargeWallet", "POST", "shopPayment", "管理员充值"},
		{"/shopPayment/getWalletLogList", "GET", "shopPayment", "钱包流水列表"},
	}

	// 1. Insert APIs
	for _, a := range apis {
		_, err := tx.Exec(`INSERT INTO sys_apis (created_at, updated_at, path, description, api_group, method) VALUES (NOW(), NOW(), ?, ?, ?, ?)`,
			a.path, a.desc, a.group, a.method)
		if err != nil {
			tx.Rollback()
			log.Fatalf("插入API %s 失败: %v", a.path, err)
		}
	}
	fmt.Printf("插入 %d 个商城API\n", len(apis))

	// 2. Insert casbin rules for admin role 888 (private routes only)
	privateAPIs := []apiDef{
		// shopCategory admin
		{"/shopCategory/createShopCategory", "POST", "", ""},
		{"/shopCategory/deleteShopCategory", "DELETE", "", ""},
		{"/shopCategory/updateShopCategory", "PUT", "", ""},
		// shopBrand admin
		{"/shopBrand/createShopBrand", "POST", "", ""},
		{"/shopBrand/deleteShopBrand", "DELETE", "", ""},
		{"/shopBrand/updateShopBrand", "PUT", "", ""},
		// shopProduct admin
		{"/shopProduct/createSpu", "POST", "", ""},
		{"/shopProduct/updateSpu", "PUT", "", ""},
		{"/shopProduct/deleteSpu", "DELETE", "", ""},
		{"/shopProduct/getSpu", "GET", "", ""},
		{"/shopProduct/getSpuList", "GET", "", ""},
		{"/shopProduct/updateSpuStatus", "PUT", "", ""},
		// shopInventory admin
		{"/shopInventory/setStock", "PUT", "", ""},
		{"/shopInventory/getInventoryList", "GET", "", ""},
		{"/shopInventory/getInventoryLogList", "GET", "", ""},
		// shopOrder admin
		{"/shopOrder/getOrderList", "GET", "", ""},
		{"/shopOrder/getOrderDetail", "GET", "", ""},
		{"/shopOrder/shipOrder", "PUT", "", ""},
		{"/shopOrder/adminCancelOrder", "PUT", "", ""},
		{"/shopOrder/getOrderLogs", "GET", "", ""},
		// shopPayment admin
		{"/shopPayment/adminRechargeWallet", "POST", "", ""},
		{"/shopPayment/getWalletLogList", "GET", "", ""},
	}

	for _, a := range privateAPIs {
		_, err := tx.Exec(`INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES ('p', '888', ?, ?)`,
			a.path, a.method)
		if err != nil {
			tx.Rollback()
			log.Fatalf("插入casbin规则 %s 失败: %v", a.path, err)
		}
	}
	fmt.Printf("插入 %d 条casbin规则(角色888)\n", len(privateAPIs))

	if err := tx.Commit(); err != nil {
		log.Fatal("提交事务失败:", err)
	}
	fmt.Println("商城API和权限初始化完成!")
}
