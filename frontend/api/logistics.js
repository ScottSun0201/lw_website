import {request} from '~/utils/request.js'

export const getTraces = (params) => {
    return request('/shopLogistics/getTraces', { method: 'get', params })
}
