import {request} from '~/utils/request.js'

export const getHotProducts = (params) => {
    return request('/shopRecommendation/getHotProducts', { method: 'get', params })
}

export const getNewProducts = (params) => {
    return request('/shopRecommendation/getNewProducts', { method: 'get', params })
}

export const getRelatedProducts = (params) => {
    return request('/shopRecommendation/getRelatedProducts', { method: 'get', params })
}
