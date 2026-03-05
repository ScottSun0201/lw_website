-- =============================================
-- 骊威商城管理端菜单初始化脚本
-- 使用方式：在已初始化的 GVA 数据库上执行
-- =============================================

-- 1. 插入顶级菜单：商城管理
INSERT INTO sys_base_menus (created_at, updated_at, menu_level, parent_id, path, name, hidden, component, sort, title, icon, keep_alive, close_tab, default_menu, active_name)
VALUES (NOW(), NOW(), 0, 0, 'shopAdmin', 'shopAdmin', 0, 'view/routerHolder.vue', 900, '商城管理', 'shop', 0, 0, 0, '');

SET @shop_parent_id = LAST_INSERT_ID();

-- 2. 插入子菜单
-- 2.1 商品分类
INSERT INTO sys_base_menus (created_at, updated_at, menu_level, parent_id, path, name, hidden, component, sort, title, icon, keep_alive, close_tab, default_menu, active_name)
VALUES (NOW(), NOW(), 0, @shop_parent_id, 'shopCategory', 'shopCategory', 0, 'view/shopCategory/shopCategory.vue', 0, '商品分类', 'grid', 0, 0, 0, '');

SET @menu_category_id = LAST_INSERT_ID();

-- 2.2 品牌管理
INSERT INTO sys_base_menus (created_at, updated_at, menu_level, parent_id, path, name, hidden, component, sort, title, icon, keep_alive, close_tab, default_menu, active_name)
VALUES (NOW(), NOW(), 0, @shop_parent_id, 'shopBrand', 'shopBrand', 0, 'view/shopBrand/shopBrand.vue', 1, '品牌管理', 'stamp', 0, 0, 0, '');

SET @menu_brand_id = LAST_INSERT_ID();

-- 2.3 商品管理
INSERT INTO sys_base_menus (created_at, updated_at, menu_level, parent_id, path, name, hidden, component, sort, title, icon, keep_alive, close_tab, default_menu, active_name)
VALUES (NOW(), NOW(), 0, @shop_parent_id, 'shopProduct', 'shopProduct', 0, 'view/shopProduct/shopProduct.vue', 2, '商品管理', 'goods', 0, 0, 0, '');

SET @menu_product_id = LAST_INSERT_ID();

-- 2.4 商品编辑（隐藏菜单，从商品列表跳转）
INSERT INTO sys_base_menus (created_at, updated_at, menu_level, parent_id, path, name, hidden, component, sort, title, icon, keep_alive, close_tab, default_menu, active_name)
VALUES (NOW(), NOW(), 0, @shop_parent_id, 'shopProductForm', 'shopProductForm', 1, 'view/shopProduct/shopProductForm.vue', 3, '商品编辑', 'goods', 0, 1, 0, 'shopProduct');

SET @menu_product_form_id = LAST_INSERT_ID();

-- 2.5 库存管理
INSERT INTO sys_base_menus (created_at, updated_at, menu_level, parent_id, path, name, hidden, component, sort, title, icon, keep_alive, close_tab, default_menu, active_name)
VALUES (NOW(), NOW(), 0, @shop_parent_id, 'shopInventory', 'shopInventory', 0, 'view/shopInventory/shopInventory.vue', 4, '库存管理', 'box', 0, 0, 0, '');

SET @menu_inventory_id = LAST_INSERT_ID();

-- 2.6 订单管理
INSERT INTO sys_base_menus (created_at, updated_at, menu_level, parent_id, path, name, hidden, component, sort, title, icon, keep_alive, close_tab, default_menu, active_name)
VALUES (NOW(), NOW(), 0, @shop_parent_id, 'shopOrder', 'shopOrder', 0, 'view/shopOrder/shopOrder.vue', 5, '订单管理', 'document', 0, 0, 0, '');

SET @menu_order_id = LAST_INSERT_ID();

-- 3. 为超级管理员角色(888)分配所有商城菜单
INSERT INTO sys_authority_menus (sys_base_menu_id, sys_authority_authority_id) VALUES
(@shop_parent_id, '888'),
(@menu_category_id, '888'),
(@menu_brand_id, '888'),
(@menu_product_id, '888'),
(@menu_product_form_id, '888'),
(@menu_inventory_id, '888'),
(@menu_order_id, '888');

-- 完成提示
SELECT CONCAT('商城菜单初始化完成，共插入 7 条菜单记录，父菜单ID: ', @shop_parent_id) AS result;
