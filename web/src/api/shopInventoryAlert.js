import service from '@/utils/request'

export const getAlertList = (params) => {
  return service({ url: '/shopInventoryAlert/getAlertList', method: 'get', params })
}

export const handleAlert = (data) => {
  return service({ url: '/shopInventoryAlert/handleAlert', method: 'put', data })
}

export const getSetting = (params) => {
  return service({ url: '/shopInventoryAlert/getSetting', method: 'get', params })
}

export const updateSetting = (data) => {
  return service({ url: '/shopInventoryAlert/updateSetting', method: 'put', data })
}
