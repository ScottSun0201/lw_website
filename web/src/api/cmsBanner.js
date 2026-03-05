import service from '@/utils/request'

// @Tags CmsBanner
// @Summary 创建轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsBanner true "创建轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsBanner/createCmsBanner [post]
export const createCmsBanner = (data) => {
  return service({
    url: '/cmsBanner/createCmsBanner',
    method: 'post',
    data
  })
}

// @Tags CmsBanner
// @Summary 删除轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsBanner true "删除轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsBanner/deleteCmsBanner [delete]
export const deleteCmsBanner = (params) => {
  return service({
    url: '/cmsBanner/deleteCmsBanner',
    method: 'delete',
    params
  })
}

// @Tags CmsBanner
// @Summary 批量删除轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsBanner/deleteCmsBanner [delete]
export const deleteCmsBannerByIds = (params) => {
  return service({
    url: '/cmsBanner/deleteCmsBannerByIds',
    method: 'delete',
    params
  })
}

// @Tags CmsBanner
// @Summary 更新轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsBanner true "更新轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsBanner/updateCmsBanner [put]
export const updateCmsBanner = (data) => {
  return service({
    url: '/cmsBanner/updateCmsBanner',
    method: 'put',
    data
  })
}

// @Tags CmsBanner
// @Summary 用id查询轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmsBanner true "用id查询轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsBanner/findCmsBanner [get]
export const findCmsBanner = (params) => {
  return service({
    url: '/cmsBanner/findCmsBanner',
    method: 'get',
    params
  })
}

// @Tags CmsBanner
// @Summary 分页获取轮播图列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取轮播图列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsBanner/getCmsBannerList [get]
export const getCmsBannerList = (params) => {
  return service({
    url: '/cmsBanner/getCmsBannerList',
    method: 'get',
    params
  })
}
