import { joinURL } from 'ufo'

// 请求代理 无论是从前端还是node端发起的请求都会被代理
 export default defineEventHandler(async event => {
    // 获取环境变量内容
     const config = useRuntimeConfig()
     const basePROXY = config.public.basePROXY
     const baseURL = config.public.baseURL
     if (event.path.indexOf(baseURL) !== 0) {
            return
     }

     // 以baseURL开头的请求都会被代理这里决定是否重写请求取消前端的代理关键key
     const path = event.path.replace(baseURL, '')

     // 代理的目标地址拼接
     const target = joinURL(basePROXY, path)

     // 代理操作
     return proxyRequest(event, target)
 })
