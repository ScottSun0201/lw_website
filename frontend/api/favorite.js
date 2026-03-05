import {request} from '~/utils/request.js'

export const addFavorite = (data) => {
    return request('/shopFavorite/addFavorite', { method: 'post', body: data })
}

export const removeFavorite = (data) => {
    return request('/shopFavorite/removeFavorite', { method: 'post', body: data })
}

export const getFavoriteList = (params) => {
    return request('/shopFavorite/getFavoriteList', { method: 'get', params })
}

export const isFavorite = (params) => {
    return request('/shopFavorite/isFavorite', { method: 'get', params })
}

export const batchCheckFavorite = (data) => {
    return request('/shopFavorite/batchCheckFavorite', { method: 'post', body: data })
}
