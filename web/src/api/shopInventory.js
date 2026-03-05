import service from '@/utils/request'

// @Tags ShopInventory
// @Summary 设置库存
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopInventory true "设置库存"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopInventory/setStock [put]
export const setStock = (data) => {
  return service({
    url: '/shopInventory/setStock',
    method: 'put',
    data
  })
}

// @Tags ShopInventory
// @Summary 分页获取库存列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取库存列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopInventory/getInventoryList [get]
export const getInventoryList = (params) => {
  return service({
    url: '/shopInventory/getInventoryList',
    method: 'get',
    params
  })
}

// @Tags ShopInventory
// @Summary 分页获取库存日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取库存日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopInventory/getInventoryLogList [get]
export const getInventoryLogList = (params) => {
  return service({
    url: '/shopInventory/getInventoryLogList',
    method: 'get',
    params
  })
}
