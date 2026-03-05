import {useUserStore} from "~/stores/userStore.js";

export default defineNuxtRouteMiddleware((to, from) => {
 const store = useUserStore()
 if(process.client) {
   if(to.name === "setting" && !store.token) {
    return navigateTo('/')
   }
 }
})
