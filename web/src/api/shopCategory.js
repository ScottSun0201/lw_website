import service from '@/utils/request'

// @Tags ShopCategory
// @Summary 创建商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopCategory true "创建商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shopCategory/createShopCategory [post]
export const createShopCategory = (data) => {
  return service({
    url: '/shopCategory/createShopCategory',
    method: 'post',
    data
  })
}

// @Tags ShopCategory
// @Summary 删除商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopCategory true "删除商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopCategory/deleteShopCategory [delete]
export const deleteShopCategory = (params) => {
  return service({
    url: '/shopCategory/deleteShopCategory',
    method: 'delete',
    params
  })
}

// @Tags ShopCategory
// @Summary 更新商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopCategory true "更新商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopCategory/updateShopCategory [put]
export const updateShopCategory = (data) => {
  return service({
    url: '/shopCategory/updateShopCategory',
    method: 'put',
    data
  })
}

// @Tags ShopCategory
// @Summary 用id查询商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ShopCategory true "用id查询商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopCategory/findShopCategory [get]
export const findShopCategory = (params) => {
  return service({
    url: '/shopCategory/findShopCategory',
    method: 'get',
    params
  })
}

// @Tags ShopCategory
// @Summary 分页获取商品分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopCategory/getShopCategoryList [get]
export const getShopCategoryList = (params) => {
  return service({
    url: '/shopCategory/getShopCategoryList',
    method: 'get',
    params
  })
}

// @Tags ShopCategory
// @Summary 获取分类树
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "获取分类树"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopCategory/getCategoryTree [get]
export const getCategoryTree = (params) => {
  return service({
    url: '/shopCategory/getCategoryTree',
    method: 'get',
    params
  })
}
