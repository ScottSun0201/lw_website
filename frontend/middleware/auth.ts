export default defineNuxtRouteMiddleware((to) => {
  const cookie = useCookie('x-token')
  if (!cookie.value) {
    return navigateTo(`/login?redirect=${encodeURIComponent(to.fullPath)}`)
  }
})
