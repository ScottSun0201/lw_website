<template>
  <div class="relative cursor-pointer" @click="navigateTo('/user/notifications')">
    <svg class="h-6 w-6 text-gray-600 hover:text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
    </svg>
    <span
      v-if="notificationStore.unreadCount > 0"
      class="absolute -top-2 -right-2 flex h-5 w-5 items-center justify-center rounded-full bg-red-500 text-xs font-bold text-white"
    >
      {{ notificationStore.unreadCount > 99 ? '99+' : notificationStore.unreadCount }}
    </span>
  </div>
</template>
<script setup>
import { useNotificationStore } from '~/stores/notificationStore.js'
const notificationStore = useNotificationStore()
const router = useRouter()
const navigateTo = (path) => { router.push(path) }
</script>
