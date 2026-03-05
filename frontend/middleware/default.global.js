import {useUserStore} from "~/stores/userStore.js";

const protectedRoutes = ['setting']

export default defineNuxtRouteMiddleware((to, from) => {
 const store = useUserStore()
 if (protectedRoutes.includes(to.name) && !store.token) {
   return navigateTo(`/login?redirect=${encodeURIComponent(to.fullPath)}`)
 }
})
