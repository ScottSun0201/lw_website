import { ElMessage,ElMessageBox } from 'element-plus';
import {useRouter } from 'vue-router'
import {useUserStore} from "~/stores/userStore.js";

//请求体封装
function useGetFetchOptions(options={}) {
  const config = useRuntimeConfig()
  options.baseURL = options?.baseURL ?? config.public.baseURL
  options.headers = {
    ...options.headers,
  }
  // SSR 时从请求头转发 cookie，客户端由浏览器自动携带
  if (process.server) {
    const nuxtApp = useNuxtApp()
    const cookieHeader = nuxtApp.ssrContext?.event?.node?.req?.headers?.cookie
    if (cookieHeader) {
      options.headers['cookie'] = cookieHeader
    }
  }
  options.credentials = 'include' // 浏览器自动携带 HttpOnly cookie
  options.initialCache = options.initialCache ?? false
  options.lazy = options.lazy ?? false

  return options
}

//http请求封装
export async function request(url,options) {
    const router = useRouter()
  options = useGetFetchOptions(options)
  if (options) {
    try {
      const res = await $fetch(url, options)
      if (res.code !== 0) {
        ElMessage({
          showClose: true,
          message: res.msg,
          type: 'error'
        })
      }
      return res
    } catch (err) {

        if (err.status === 401){
            ElMessageBox.confirm(`
          <p>登录已过期，请重新登录</p>
          `, '身份信息', {
                dangerouslyUseHTMLString: true,
                distinguishCancelAndClose: true,
                confirmButtonText: '重新登录',
                cancelButtonText: '取消'
            })
                .then(() => {
                    const userStore = useUserStore()
                    userStore.clearLocal()
                    router.push({ name: 'login', replace: true })
                })
            return
        }


        ElMessage({
            showClose: true,
            message: err?.message || '请求失败',
            type: 'error'
        })
        return { code: err?.status || -1, msg: err?.message || '请求失败', data: null }
    }
  }
}
