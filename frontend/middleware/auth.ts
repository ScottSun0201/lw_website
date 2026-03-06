export default defineNuxtRouteMiddleware((to) => {
  // SSR 时通过 cookie 检查（服务端可以读取 HttpOnly cookie）
  if (process.server) {
    const cookie = useCookie('x-token')
    if (!cookie.value) {
      return navigateTo(`/login?redirect=${encodeURIComponent(to.fullPath)}`)
    }
  } else {
    // 客户端通过 store 检查登录状态
    const userStore = useUserStore()
    if (!userStore.token) {
      return navigateTo(`/login?redirect=${encodeURIComponent(to.fullPath)}`)
    }
  }
})
