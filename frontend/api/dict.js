import {request} from '~/utils/request.js'
export const findSysDictionary  = (params) => {
    return request('/sysDictionary/findSysDictionary',{
        method: 'get',
        params
    })
}
