import { ElMessage,ElMessageBox } from 'element-plus';
import cookie from 'js-cookie'
import {useRouter } from 'vue-router'

//请求体封装
function useGetFetchOptions(options={}) {
  let token = ""
  const config = useRuntimeConfig()
  if (process.client) {
    token = cookie.get('x-token') || ''
  }
  options.baseURL = options?.baseURL ?? config.public.baseURL
  options.headers = {
    ...options.headers,
    'x-token':token,
  }
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
          <p>无效的令牌</p>
          <p>错误码:<span style="color:red"> 401 </span>错误信息:${err}</p>
          `, '身份信息', {
                dangerouslyUseHTMLString: true,
                distinguishCancelAndClose: true,
                confirmButtonText: '重新登录',
                cancelButtonText: '取消'
            })
                .then(() => {
                    cookie.set("x-token", "")
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
