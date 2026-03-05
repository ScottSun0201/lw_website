import service from '@/utils/request'

// @Tags ShopBrand
// @Summary 创建品牌
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopBrand true "创建品牌"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shopBrand/createShopBrand [post]
export const createShopBrand = (data) => {
  return service({
    url: '/shopBrand/createShopBrand',
    method: 'post',
    data
  })
}

// @Tags ShopBrand
// @Summary 删除品牌
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopBrand true "删除品牌"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopBrand/deleteShopBrand [delete]
export const deleteShopBrand = (params) => {
  return service({
    url: '/shopBrand/deleteShopBrand',
    method: 'delete',
    params
  })
}

// @Tags ShopBrand
// @Summary 更新品牌
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopBrand true "更新品牌"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopBrand/updateShopBrand [put]
export const updateShopBrand = (data) => {
  return service({
    url: '/shopBrand/updateShopBrand',
    method: 'put',
    data
  })
}

// @Tags ShopBrand
// @Summary 用id查询品牌
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ShopBrand true "用id查询品牌"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopBrand/findShopBrand [get]
export const findShopBrand = (params) => {
  return service({
    url: '/shopBrand/findShopBrand',
    method: 'get',
    params
  })
}

// @Tags ShopBrand
// @Summary 分页获取品牌列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取品牌列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopBrand/getShopBrandList [get]
export const getShopBrandList = (params) => {
  return service({
    url: '/shopBrand/getShopBrandList',
    method: 'get',
    params
  })
}

// @Tags ShopBrand
// @Summary 获取全部品牌
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "获取全部品牌"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopBrand/getAllBrands [get]
export const getAllBrands = (params) => {
  return service({
    url: '/shopBrand/getAllBrands',
    method: 'get',
    params
  })
}
