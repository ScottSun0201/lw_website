import {request} from '~/utils/request.js'

export const getMyNotifications = (params) => {
    return request('/shopNotification/getMyNotifications', { method: 'get', params })
}

export const markRead = (params) => {
    return request('/shopNotification/markRead', { method: 'put', params })
}

export const markAllRead = () => {
    return request('/shopNotification/markAllRead', { method: 'put' })
}

export const getUnreadCount = () => {
    return request('/shopNotification/getUnreadCount', { method: 'get' })
}
