import {request} from '~/utils/request.js'
export const getCmsArticleList  = (params) => {
    return request('/cmsArticle/getCmsArticleList',{
        method: 'get',
        params
    })
}


export const findCmsArticle  = (params) => {
    return request('/cmsArticle/findCmsArticle',{
        method: 'get',
        params
    })
}


export const createCmsCommit  = (data) => {
    return request('/cmsCommit/createCmsCommit',{
        method: 'post',
        body:data
    })
}

export const getCmsCommitList  = (params) => {
    return request('/cmsCommit/getCmsCommitList',{
        method: 'get',
        params
    })
}

