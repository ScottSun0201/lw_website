import service from '@/utils/request'

// @Tags ClientAbout
// @Summary 创建关于我们
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClientAbout true "创建关于我们"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /clientAbout/createClientAbout [post]
export const createClientAbout = (data) => {
  return service({
    url: '/clientAbout/createClientAbout',
    method: 'post',
    data
  })
}

// @Tags ClientAbout
// @Summary 删除关于我们
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClientAbout true "删除关于我们"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clientAbout/deleteClientAbout [delete]
export const deleteClientAbout = (params) => {
  return service({
    url: '/clientAbout/deleteClientAbout',
    method: 'delete',
    params
  })
}

// @Tags ClientAbout
// @Summary 批量删除关于我们
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除关于我们"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clientAbout/deleteClientAbout [delete]
export const deleteClientAboutByIds = (params) => {
  return service({
    url: '/clientAbout/deleteClientAboutByIds',
    method: 'delete',
    params
  })
}

// @Tags ClientAbout
// @Summary 更新关于我们
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClientAbout true "更新关于我们"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clientAbout/updateClientAbout [put]
export const updateClientAbout = (data) => {
  return service({
    url: '/clientAbout/updateClientAbout',
    method: 'put',
    data
  })
}

// @Tags ClientAbout
// @Summary 用id查询关于我们
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ClientAbout true "用id查询关于我们"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clientAbout/findClientAbout [get]
export const findClientAbout = (params) => {
  return service({
    url: '/clientAbout/findClientAbout',
    method: 'get',
    params
  })
}

// @Tags ClientAbout
// @Summary 分页获取关于我们列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取关于我们列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clientAbout/getClientAboutList [get]
export const getClientAboutList = (params) => {
  return service({
    url: '/clientAbout/getClientAboutList',
    method: 'get',
    params
  })
}

export const getClientAboutInfo = () => {
  return service({
    url: '/clientAbout/getClientAboutInfo',
    method: 'get'
  })
}


export const setClientAboutInfo = (data) => {
  return service({
    url: '/clientAbout/setClientAboutInfo',
    method: 'post',
    data
  })
}

