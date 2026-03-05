import {request} from '~/utils/request.js'

export const getCmsTagList = () => {
    return request('/cmsTag/getCmsTagList',{
        method: 'get'
    })
}
