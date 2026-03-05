import service from '@/utils/request'

export const getAdminNotifications = (params) => {
  return service({ url: '/shopNotification/getAdminNotifications', method: 'get', params })
}

export const markAdminRead = (data) => {
  return service({ url: '/shopNotification/markAdminRead', method: 'put', data })
}

export const markAllAdminRead = (data) => {
  return service({ url: '/shopNotification/markAllAdminRead', method: 'put', data })
}

export const getAdminUnreadCount = (params) => {
  return service({ url: '/shopNotification/getAdminUnreadCount', method: 'get', params })
}
