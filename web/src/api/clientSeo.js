import service from '@/utils/request'

// @Tags ClientSEO
// @Summary 创建客户端SEO
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClientSEO true "创建客户端SEO"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /clientSEO/createClientSEO [post]
export const createClientSEO = (data) => {
  return service({
    url: '/clientSEO/createClientSEO',
    method: 'post',
    data
  })
}

// @Tags ClientSEO
// @Summary 用id查询客户端SEO
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ClientSEO true "用id查询客户端SEO"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clientSEO/findClientSEO [get]
export const findClientSEO = () => {
  return service({
    url: '/clientSEO/findClientSEO',
    method: 'get'
  })
}
