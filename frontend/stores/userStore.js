import { defineStore } from 'pinia'
import { ref, onMounted } from 'vue'
import {getUserInfo, jsonInBlacklist} from "~/api/base.js";
import cookie from "js-cookie";
import {useRouter} from "vue-router";

export const useUserStore = defineStore('user', () => {
    const router = useRouter()
    const userInfo = ref({
        username: '',
        nickname: '',
        about: '',
        avatar: '',
        firstName: '',
        lastName: '',
        email: '',
    })

    const token = ref('')

    const getInfo = async () =>{
        const res =  await getUserInfo()
        if(res.code === 0){
            userInfo.value = res.data
        }
    }

    const setToken = (newToken) => {
        if(newToken){
            getInfo()
        }
        token.value = newToken
    }

    const logout = async () => {
        const res = await jsonInBlacklist()
        if(res.code === 0){
            router.push({name: 'index'})
            cookie.remove('x-token')
            setToken('')
        }
    }

    return {
        userInfo,
        token,
        setToken,
        getInfo,
        logout
    }
})
