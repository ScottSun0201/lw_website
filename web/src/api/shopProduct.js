import service from '@/utils/request'

// @Tags ShopProduct
// @Summary 创建SPU商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopProduct true "创建SPU商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shopProduct/createSpu [post]
export const createSpu = (data) => {
  return service({
    url: '/shopProduct/createSpu',
    method: 'post',
    data
  })
}

// @Tags ShopProduct
// @Summary 更新SPU商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopProduct true "更新SPU商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopProduct/updateSpu [put]
export const updateSpu = (data) => {
  return service({
    url: '/shopProduct/updateSpu',
    method: 'put',
    data
  })
}

// @Tags ShopProduct
// @Summary 删除SPU商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopProduct true "删除SPU商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopProduct/deleteSpu [delete]
export const deleteSpu = (params) => {
  return service({
    url: '/shopProduct/deleteSpu',
    method: 'delete',
    params
  })
}

// @Tags ShopProduct
// @Summary 获取SPU商品详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ShopProduct true "获取SPU商品详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopProduct/getSpu [get]
export const getSpu = (params) => {
  return service({
    url: '/shopProduct/getSpu',
    method: 'get',
    params
  })
}

// @Tags ShopProduct
// @Summary 分页获取SPU商品列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SPU商品列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopProduct/getSpuList [get]
export const getSpuList = (params) => {
  return service({
    url: '/shopProduct/getSpuList',
    method: 'get',
    params
  })
}

// @Tags ShopProduct
// @Summary 更新SPU商品状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopProduct true "更新SPU商品状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopProduct/updateSpuStatus [put]
export const updateSpuStatus = (data) => {
  return service({
    url: '/shopProduct/updateSpuStatus',
    method: 'put',
    data
  })
}
