<template>
  <button @click.prevent.stop="toggleFavorite" class="transition-colors" :class="isFav ? 'text-red-500' : 'text-gray-300 hover:text-red-400'">
    <svg class="h-6 w-6" :fill="isFav ? 'currentColor' : 'none'" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
    </svg>
  </button>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { addFavorite, removeFavorite, isFavorite as checkFavorite } from '~/api/favorite.js'
import { useUserStore } from '~/stores/userStore.js'

const props = defineProps({ spuId: { type: Number, required: true } })
const userStore = useUserStore()
const isFav = ref(false)

const toggleFavorite = async () => {
  if (!userStore.token) { ElMessage.warning('请先登录'); return }
  if (isFav.value) {
    const res = await removeFavorite({ spuId: props.spuId })
    if (res && res.code === 0) { isFav.value = false; ElMessage.success('已取消收藏') }
  } else {
    const res = await addFavorite({ spuId: props.spuId })
    if (res && res.code === 0) { isFav.value = true; ElMessage.success('已收藏') }
  }
}

onMounted(async () => {
  if (!userStore.token) return
  const res = await checkFavorite({ spuId: props.spuId })
  if (res && res.code === 0) { isFav.value = res.data.isFavorite }
})
</script>
