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

    let infoAbort = null
    const setToken = (newToken) => {
        if (infoAbort) {
            infoAbort = null
        }
        token.value = newToken
        if (newToken) {
            infoAbort = true
            getInfo().finally(() => { infoAbort = null })
        }
    }

    const logout = async () => {
        const res = await jsonInBlacklist()
        if(res.code === 0){
            cookie.remove('x-token')
            token.value = ''
            userInfo.value = { username: '', nickname: '', about: '', avatar: '', firstName: '', lastName: '', email: '' }
            router.push({name: 'index'})
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
