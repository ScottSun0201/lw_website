import service from '@/utils/request'

export const getTraces = (params) => {
  return service({ url: '/shopLogistics/getTraces', method: 'get', params })
}

export const addTrace = (data) => {
  return service({ url: '/shopLogistics/addTrace', method: 'post', data })
}
