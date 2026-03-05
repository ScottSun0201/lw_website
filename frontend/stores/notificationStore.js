import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getUnreadCount } from '~/api/notification.js'

export const useNotificationStore = defineStore('notification', () => {
    const unreadCount = ref(0)

    const refreshCount = async () => {
        const res = await getUnreadCount()
        if (res && res.code === 0) {
            unreadCount.value = res.data.count || 0
        }
    }

    return { unreadCount, refreshCount }
})
