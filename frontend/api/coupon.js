import {request} from '~/utils/request.js'

export const claimCoupon = (data) => {
    return request('/shopCoupon/claimCoupon', { method: 'post', body: data })
}

export const getMyCoupons = (params) => {
    return request('/shopCoupon/getMyCoupons', { method: 'get', params })
}

export const getAvailableCoupons = () => {
    return request('/shopCoupon/getAvailableCoupons', { method: 'get' })
}
