import service from '@/utils/request'

// @Tags CmsArticle
// @Summary 创建文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsArticle true "创建文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsArticle/createCmsArticle [post]
export const createCmsArticle = (data) => {
  return service({
    url: '/cmsArticle/createCmsArticle',
    method: 'post',
    data
  })
}

// @Tags CmsArticle
// @Summary 删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsArticle true "删除文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsArticle/deleteCmsArticle [delete]
export const deleteCmsArticle = (params) => {
  return service({
    url: '/cmsArticle/deleteCmsArticle',
    method: 'delete',
    params
  })
}

// @Tags CmsArticle
// @Summary 批量删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsArticle/deleteCmsArticle [delete]
export const deleteCmsArticleByIds = (params) => {
  return service({
    url: '/cmsArticle/deleteCmsArticleByIds',
    method: 'delete',
    params
  })
}

// @Tags CmsArticle
// @Summary 更新文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsArticle true "更新文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsArticle/updateCmsArticle [put]
export const updateCmsArticle = (data) => {
  return service({
    url: '/cmsArticle/updateCmsArticle',
    method: 'put',
    data
  })
}

// @Tags CmsArticle
// @Summary 用id查询文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmsArticle true "用id查询文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsArticle/findCmsArticle [get]
export const findCmsArticle = (params) => {
  return service({
    url: '/cmsArticle/findCmsArticle',
    method: 'get',
    params
  })
}

// @Tags CmsArticle
// @Summary 分页获取文章列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取文章列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsArticle/getCmsArticleList [get]
export const getCmsArticleList = (params) => {
  return service({
    url: '/cmsArticle/getCmsArticleList',
    method: 'get',
    params
  })
}
