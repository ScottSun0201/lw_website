import service from '@/utils/request'

export const setRecommendation = (data) => {
  return service({ url: '/shopRecommendation/setRecommendation', method: 'post', data })
}

export const removeRecommendation = (data) => {
  return service({ url: '/shopRecommendation/removeRecommendation', method: 'delete', data })
}

export const getRecommendationList = (params) => {
  return service({ url: '/shopRecommendation/getRecommendationList', method: 'get', params })
}

export const updateSort = (data) => {
  return service({ url: '/shopRecommendation/updateSort', method: 'put', data })
}
