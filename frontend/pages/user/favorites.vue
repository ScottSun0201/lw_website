<template>
  <div class="bg-gray-50 min-h-screen">
    <div class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
      <h1 class="mb-6 text-2xl font-bold text-gray-900">我的收藏</h1>
      <div v-if="list.length" class="grid grid-cols-2 gap-4 sm:gap-6 lg:grid-cols-4">
        <ShopProductCard v-for="product in list" :key="product.ID" :product="product" />
      </div>
      <div v-else class="py-16 text-center text-gray-400">暂无收藏</div>
      <div v-if="total > 0" class="mt-6 flex justify-center">
        <el-pagination v-model:current-page="page" :page-size="pageSize" :total="total" layout="prev, pager, next" background @current-change="fetchList" />
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { getFavoriteList } from '~/api/favorite.js'

definePageMeta({ middleware: 'auth' })

const list = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = 12

const fetchList = async () => {
  const res = await getFavoriteList({ page: page.value, pageSize })
  if (res && res.code === 0) { list.value = res.data.list || []; total.value = res.data.total || 0 }
}

await fetchList()
</script>
