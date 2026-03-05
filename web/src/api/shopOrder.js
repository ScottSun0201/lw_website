import service from '@/utils/request'

// @Tags ShopOrder
// @Summary 分页获取订单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopOrder/getOrderList [get]
export const getOrderList = (params) => {
  return service({
    url: '/shopOrder/getOrderList',
    method: 'get',
    params
  })
}

// @Tags ShopOrder
// @Summary 获取订单详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ShopOrder true "获取订单详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopOrder/getOrderDetail [get]
export const getOrderDetail = (params) => {
  return service({
    url: '/shopOrder/getOrderDetail',
    method: 'get',
    params
  })
}

// @Tags ShopOrder
// @Summary 订单发货
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopOrder true "订单发货"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopOrder/shipOrder [put]
export const shipOrder = (data) => {
  return service({
    url: '/shopOrder/shipOrder',
    method: 'put',
    data
  })
}

// @Tags ShopOrder
// @Summary 管理员取消订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ShopOrder true "管理员取消订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopOrder/adminCancelOrder [put]
export const adminCancelOrder = (params) => {
  return service({
    url: '/shopOrder/adminCancelOrder',
    method: 'put',
    params
  })
}

// @Tags ShopOrder
// @Summary 获取订单日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "获取订单日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopOrder/getOrderLogs [get]
export const getOrderLogs = (params) => {
  return service({
    url: '/shopOrder/getOrderLogs',
    method: 'get',
    params
  })
}
