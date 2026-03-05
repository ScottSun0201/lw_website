import service from '@/utils/request'

export const createCoupon = (data) => {
  return service({ url: '/shopCoupon/createCoupon', method: 'post', data })
}

export const updateCoupon = (data) => {
  return service({ url: '/shopCoupon/updateCoupon', method: 'put', data })
}

export const deleteCoupon = (data) => {
  return service({ url: '/shopCoupon/deleteCoupon', method: 'delete', data })
}

export const getCouponList = (params) => {
  return service({ url: '/shopCoupon/getCouponList', method: 'get', params })
}

export const getCouponDetail = (params) => {
  return service({ url: '/shopCoupon/getCouponDetail', method: 'get', params })
}
