import {request} from '~/utils/request.js'

export const createRefund = (data) => {
    return request('/shopRefund/createRefund', { method: 'post', body: data })
}

export const shipReturn = (data) => {
    return request('/shopRefund/shipReturn', { method: 'put', body: data })
}

export const cancelRefund = (params) => {
    return request('/shopRefund/cancelRefund', { method: 'put', params })
}

export const getUserRefunds = (params) => {
    return request('/shopRefund/getUserRefunds', { method: 'get', params })
}

export const getRefundDetail = (params) => {
    return request('/shopRefund/getRefundDetail', { method: 'get', params })
}
