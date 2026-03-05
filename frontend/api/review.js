import {request} from '~/utils/request.js'

export const createReview = (data) => {
    return request('/shopReview/createReview', { method: 'post', body: data })
}

export const getProductReviews = (params) => {
    return request('/shopReview/getProductReviews', { method: 'get', params })
}

export const getReviewStat = (params) => {
    return request('/shopReview/getReviewStat', { method: 'get', params })
}

export const getMyReviews = (params) => {
    return request('/shopReview/getMyReviews', { method: 'get', params })
}
