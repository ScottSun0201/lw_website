import {request} from '~/utils/request.js'

export const getCmsFriendInfo = () => {
    return request('/cmsFriend/getCmsFriendInfo',{
        method: 'get'
    })
}

export const getCmsFriendList = () => {
    return request('/cmsFriend/getCmsFriendList',{
        method: 'get'
    })
}
