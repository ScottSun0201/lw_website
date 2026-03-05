import service from '@/utils/request'

export const getRefundList = (params) => {
  return service({ url: '/shopRefund/getRefundList', method: 'get', params })
}

export const getRefundDetail = (params) => {
  return service({ url: '/shopRefund/getRefundDetail', method: 'get', params })
}

export const auditRefund = (data) => {
  return service({ url: '/shopRefund/auditRefund', method: 'put', data })
}

export const confirmReturn = (data) => {
  return service({ url: '/shopRefund/confirmReturn', method: 'put', data })
}
