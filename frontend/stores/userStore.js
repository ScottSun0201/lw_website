import { defineStore } from 'pinia'
import { ref } from 'vue'
import {getUserInfo, jsonInBlacklist} from "~/api/base.js";
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
        if(res && res.code === 0){
            userInfo.value = res.data
        }
    }

    // 通过调用 getUserInfo 检查登录状态（HttpOnly cookie 由浏览器自动携带）
    const checkAuth = async () => {
        try {
            const res = await getUserInfo()
            if (res && res.code === 0) {
                token.value = 'authenticated'
                userInfo.value = res.data
            }
        } catch (e) {
            // 未登录或 token 无效
        }
    }

    const setToken = (newToken) => {
        token.value = newToken
        if (newToken) {
            getInfo()
        }
    }

    const clearLocal = () => {
        token.value = ''
        userInfo.value = { username: '', nickname: '', about: '', avatar: '', firstName: '', lastName: '', email: '' }
    }

    const logout = async () => {
        const res = await jsonInBlacklist()
        if(res && res.code === 0){
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
        checkAuth,
        clearLocal,
        logout
    }
})
