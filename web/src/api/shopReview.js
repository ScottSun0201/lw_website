import service from '@/utils/request'

export const getReviewList = (params) => {
  return service({ url: '/shopReview/getReviewList', method: 'get', params })
}

export const auditReview = (data) => {
  return service({ url: '/shopReview/auditReview', method: 'put', data })
}

export const replyReview = (data) => {
  return service({ url: '/shopReview/replyReview', method: 'put', data })
}
