import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getCartCount } from '~/api/shop.js'

export const useCartStore = defineStore('cart', () => {
    const count = ref(0)

    const refreshCount = async () => {
        const res = await getCartCount()
        if (res && res.code === 0) {
            count.value = res.data.count || 0
        }
    }

    const increment = () => {
        count.value++
    }

    const setCount = (val) => {
        count.value = val
    }

    return {
        count,
        refreshCount,
        increment,
        setCount
    }
})
