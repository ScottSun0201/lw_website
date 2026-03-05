import service from '@/utils/request'

// @Tags CmsCommit
// @Summary 创建评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsCommit true "创建评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsCommit/createCmsCommit [post]
export const createCmsCommit = (data) => {
  return service({
    url: '/cmsCommit/createCmsCommit',
    method: 'post',
    data
  })
}

// @Tags CmsCommit
// @Summary 删除评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsCommit true "删除评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsCommit/deleteCmsCommit [delete]
export const deleteCmsCommit = (params) => {
  return service({
    url: '/cmsCommit/deleteCmsCommit',
    method: 'delete',
    params
  })
}

// @Tags CmsCommit
// @Summary 批量删除评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsCommit/deleteCmsCommit [delete]
export const deleteCmsCommitByIds = (params) => {
  return service({
    url: '/cmsCommit/deleteCmsCommitByIds',
    method: 'delete',
    params
  })
}

// @Tags CmsCommit
// @Summary 更新评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsCommit true "更新评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsCommit/updateCmsCommit [put]
export const updateCmsCommit = (data) => {
  return service({
    url: '/cmsCommit/updateCmsCommit',
    method: 'put',
    data
  })
}

// @Tags CmsCommit
// @Summary 用id查询评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmsCommit true "用id查询评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsCommit/findCmsCommit [get]
export const findCmsCommit = (params) => {
  return service({
    url: '/cmsCommit/findCmsCommit',
    method: 'get',
    params
  })
}

// @Tags CmsCommit
// @Summary 分页获取评论列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取评论列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCommit/getCmsCommitList [get]
export const getCmsCommitList = (params) => {
  return service({
    url: '/cmsCommit/getCmsCommitList',
    method: 'get',
    params
  })
}
