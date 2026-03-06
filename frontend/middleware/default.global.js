import {useUserStore} from "~/stores/userStore.js";

const protectedRoutes = ['setting']

export default defineNuxtRouteMiddleware((to, from) => {
 const store = useUserStore()
 if (protectedRoutes.includes(to.name) && !store.token) {
   if (process.server) {
     const cookie = useCookie('x-token')
     if (!cookie.value) {
       return navigateTo(`/login?redirect=${encodeURIComponent(to.fullPath)}`)
     }
   } else {
     return navigateTo(`/login?redirect=${encodeURIComponent(to.fullPath)}`)
   }
 }
})
