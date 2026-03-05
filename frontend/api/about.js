import {request} from "~/utils/request.js";

export const getClientAboutInfo  = () => {
    return request('/clientAbout/getClientAboutInfo',{
        method: 'get'
    })
}


export const getClientAboutList  = () => {
    return request('/clientAbout/getClientAboutList',{
        method: 'get'
    })
}
