import {request} from '~/utils/request.js'
export const getCaptcha  = () => {
    return request('/base/captcha',{
        method: 'post'
    })
}

export const register  = (data) => {
    return request('/clientUser/register',{
        method: 'post',
        body:data
    })
}

export const login  = (data) => {
    return request('/clientUser/login',{
        method: 'post',
        body:data
    })
}

export const getUserInfo  = () => {
    return request('/clientUser/getUserInfo',{
        method: 'get'
    })
}


export const logout  = () => {
    return request('/jwt/jsonInBlacklist',{
        method: 'post'
    })
}


export const setClientUser = (data) => {
    return request('/clientUser/setClientUser',{
        method: 'put',
        body:data
    })
}

export const upload = (data,params) => {
    return request('/fileUploadAndDownload/upload',{
        method: 'post',
        body:data,
        params
    })
}


export const findClientSEO = () => {
    return request('/clientSEO/findClientSEO',{
        method: 'get',
    })
}


export const getCmsPlateList = () => {
    return request('/cmsPlate/getCmsPlateList',{
        method: 'get',
    })
}

export const jsonInBlacklist = () => {
    return request('/jwt/jsonInBlacklist',{
        method: 'post',
    })
}
