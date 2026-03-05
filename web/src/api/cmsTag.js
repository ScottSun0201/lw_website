import service from '@/utils/request'

// @Tags CmsTag
// @Summary 创建标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsTag true "创建标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsTag/createCmsTag [post]
export const createCmsTag = (data) => {
  return service({
    url: '/cmsTag/createCmsTag',
    method: 'post',
    data
  })
}

// @Tags CmsTag
// @Summary 删除标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsTag true "删除标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsTag/deleteCmsTag [delete]
export const deleteCmsTag = (params) => {
  return service({
    url: '/cmsTag/deleteCmsTag',
    method: 'delete',
    params
  })
}

// @Tags CmsTag
// @Summary 批量删除标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsTag/deleteCmsTag [delete]
export const deleteCmsTagByIds = (params) => {
  return service({
    url: '/cmsTag/deleteCmsTagByIds',
    method: 'delete',
    params
  })
}

// @Tags CmsTag
// @Summary 更新标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsTag true "更新标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsTag/updateCmsTag [put]
export const updateCmsTag = (data) => {
  return service({
    url: '/cmsTag/updateCmsTag',
    method: 'put',
    data
  })
}

// @Tags CmsTag
// @Summary 用id查询标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmsTag true "用id查询标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsTag/findCmsTag [get]
export const findCmsTag = (params) => {
  return service({
    url: '/cmsTag/findCmsTag',
    method: 'get',
    params
  })
}

// @Tags CmsTag
// @Summary 分页获取标签列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取标签列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsTag/getCmsTagList [get]
export const getCmsTagList = (params) => {
  return service({
    url: '/cmsTag/getCmsTagList',
    method: 'get',
    params
  })
}
