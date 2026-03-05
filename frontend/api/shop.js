import {request} from '~/utils/request.js'

// 商品分类
export const getCategoryTree = (params) => {
    return request('/shopCategory/getCategoryTree', {
        method: 'get',
        params
    })
}

// 品牌
export const getAllBrands = (params) => {
    return request('/shopBrand/getAllBrands', {
        method: 'get',
        params
    })
}

// 商品列表（客户端）
export const getClientProductList = (params) => {
    return request('/shopProduct/getClientProductList', {
        method: 'get',
        params
    })
}

// 商品详情（客户端）
export const getClientProductDetail = (params) => {
    return request('/shopProduct/getClientProductDetail', {
        method: 'get',
        params
    })
}

// 购物车
export const addToCart = (data) => {
    return request('/shopCart/addToCart', {
        method: 'post',
        body: data
    })
}

export const updateCartQuantity = (data) => {
    return request('/shopCart/updateCartQuantity', {
        method: 'put',
        body: data
    })
}

export const removeFromCart = (data) => {
    return request('/shopCart/removeFromCart', {
        method: 'post',
        body: data
    })
}

export const getCartList = (params) => {
    return request('/shopCart/getCartList', {
        method: 'get',
        params
    })
}

export const getCartCount = (params) => {
    return request('/shopCart/getCartCount', {
        method: 'get',
        params
    })
}

// 收货地址
export const createAddress = (data) => {
    return request('/shopAddress/createAddress', {
        method: 'post',
        body: data
    })
}

export const deleteAddress = (params) => {
    return request('/shopAddress/deleteAddress', {
        method: 'delete',
        params
    })
}

export const updateAddress = (data) => {
    return request('/shopAddress/updateAddress', {
        method: 'put',
        body: data
    })
}

export const getAddressList = (params) => {
    return request('/shopAddress/getAddressList', {
        method: 'get',
        params
    })
}

export const setDefaultAddress = (params) => {
    return request('/shopAddress/setDefaultAddress', {
        method: 'put',
        params
    })
}

// 订单
export const createOrder = (data) => {
    return request('/shopOrder/createOrder', {
        method: 'post',
        body: data
    })
}

export const cancelOrder = (params) => {
    return request('/shopOrder/cancelOrder', {
        method: 'put',
        params
    })
}

export const confirmReceive = (params) => {
    return request('/shopOrder/confirmReceive', {
        method: 'put',
        params
    })
}

export const getUserOrderList = (params) => {
    return request('/shopOrder/getUserOrderList', {
        method: 'get',
        params
    })
}

export const getUserOrder = (params) => {
    return request('/shopOrder/getUserOrder', {
        method: 'get',
        params
    })
}

// 支付
export const balancePay = (data) => {
    return request('/shopPayment/balancePay', {
        method: 'post',
        body: data
    })
}

export const getWallet = (params) => {
    return request('/shopPayment/getWallet', {
        method: 'get',
        params
    })
}
