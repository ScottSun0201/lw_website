import service from '@/utils/request'

// @Tags CmsFriend
// @Summary 创建友情链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsFriend true "创建友情链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsFriend/createCmsFriend [post]
export const createCmsFriend = (data) => {
  return service({
    url: '/cmsFriend/createCmsFriend',
    method: 'post',
    data
  })
}

// @Tags CmsFriend
// @Summary 删除友情链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsFriend true "删除友情链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsFriend/deleteCmsFriend [delete]
export const deleteCmsFriend = (params) => {
  return service({
    url: '/cmsFriend/deleteCmsFriend',
    method: 'delete',
    params
  })
}

// @Tags CmsFriend
// @Summary 批量删除友情链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除友情链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsFriend/deleteCmsFriend [delete]
export const deleteCmsFriendByIds = (params) => {
  return service({
    url: '/cmsFriend/deleteCmsFriendByIds',
    method: 'delete',
    params
  })
}

// @Tags CmsFriend
// @Summary 更新友情链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsFriend true "更新友情链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsFriend/updateCmsFriend [put]
export const updateCmsFriend = (data) => {
  return service({
    url: '/cmsFriend/updateCmsFriend',
    method: 'put',
    data
  })
}

// @Tags CmsFriend
// @Summary 用id查询友情链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmsFriend true "用id查询友情链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsFriend/findCmsFriend [get]
export const findCmsFriend = (params) => {
  return service({
    url: '/cmsFriend/findCmsFriend',
    method: 'get',
    params
  })
}

// @Tags CmsFriend
// @Summary 分页获取友情链接列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取友情链接列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsFriend/getCmsFriendList [get]
export const getCmsFriendList = (params) => {
  return service({
    url: '/cmsFriend/getCmsFriendList',
    method: 'get',
    params
  })
}

export const setCmsFriendInfo = (data) => {
  return service({
    url: '/cmsFriend/setCmsFriendInfo',
    method: 'post',
    data
  })
}

export const getCmsFriendInfo = () => {
  return service({
    url: '/cmsFriend/getCmsFriendInfo',
    method: 'get'
  })
}
