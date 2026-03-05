import service from '@/utils/request'

// @Tags ShopPayment
// @Summary 管理员充值钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopPayment true "管理员充值钱包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shopPayment/adminRechargeWallet [post]
export const adminRechargeWallet = (data) => {
  return service({
    url: '/shopPayment/adminRechargeWallet',
    method: 'post',
    data
  })
}

// @Tags ShopPayment
// @Summary 分页获取钱包日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取钱包日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopPayment/getWalletLogList [get]
export const getWalletLogList = (params) => {
  return service({
    url: '/shopPayment/getWalletLogList',
    method: 'get',
    params
  })
}
