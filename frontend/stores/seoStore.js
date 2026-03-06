import { defineStore } from 'pinia'
import { ref } from 'vue'
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
