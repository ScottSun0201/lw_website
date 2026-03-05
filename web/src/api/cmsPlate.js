import service from '@/utils/request'

// @Tags CmsPlate
// @Summary 创建板块
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsPlate true "创建板块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsPlate/createCmsPlate [post]
export const createCmsPlate = (data) => {
  return service({
    url: '/cmsPlate/createCmsPlate',
    method: 'post',
    data
  })
}

// @Tags CmsPlate
// @Summary 删除板块
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsPlate true "删除板块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsPlate/deleteCmsPlate [delete]
export const deleteCmsPlate = (params) => {
  return service({
    url: '/cmsPlate/deleteCmsPlate',
    method: 'delete',
    params
  })
}

// @Tags CmsPlate
// @Summary 批量删除板块
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除板块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsPlate/deleteCmsPlate [delete]
export const deleteCmsPlateByIds = (params) => {
  return service({
    url: '/cmsPlate/deleteCmsPlateByIds',
    method: 'delete',
    params
  })
}

// @Tags CmsPlate
// @Summary 更新板块
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsPlate true "更新板块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsPlate/updateCmsPlate [put]
export const updateCmsPlate = (data) => {
  return service({
    url: '/cmsPlate/updateCmsPlate',
    method: 'put',
    data
  })
}

// @Tags CmsPlate
// @Summary 用id查询板块
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmsPlate true "用id查询板块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsPlate/findCmsPlate [get]
export const findCmsPlate = (params) => {
  return service({
    url: '/cmsPlate/findCmsPlate',
    method: 'get',
    params
  })
}

// @Tags CmsPlate
// @Summary 分页获取板块列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取板块列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsPlate/getCmsPlateList [get]
export const getCmsPlateList = (params) => {
  return service({
    url: '/cmsPlate/getCmsPlateList',
    method: 'get',
    params
  })
}
