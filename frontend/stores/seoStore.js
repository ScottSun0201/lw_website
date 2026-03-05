import { defineStore } from 'pinia'
import { ref, watch } from 'vue'
import {getUserInfo} from "~/api/base.js";
import cookie from "js-cookie";
export const useSeoStore = defineStore('seo', () => {
    const seo = ref({
        title: '',
        keywords: '',
        description: '',
        name: '',
        logo: '',
        favicon: '',
    })

    const setSeo = (newSeo) => {
        seo.value = newSeo
    }

    return {
        seo,
        setSeo
    }
})
