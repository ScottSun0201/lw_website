<template>
  <div class="bg-gray-50 min-h-screen">
    <div class="mx-auto max-w-4xl px-4 py-8 sm:px-6 lg:px-8">
      <h1 class="mb-6 text-2xl font-bold text-gray-900">我的优惠券</h1>
      <el-tabs v-model="activeTab" @tab-change="onTabChange">
        <el-tab-pane label="可使用" name="0" />
        <el-tab-pane label="已使用" name="1" />
        <el-tab-pane label="已过期" name="2" />
      </el-tabs>
      <div class="space-y-3 mt-4">
        <ShopCouponCard v-for="uc in list" :key="uc.ID" :coupon="uc.coupon" :disabled="activeTab !== '0'" />
      </div>
      <div v-if="!list.length" class="py-16 text-center text-gray-400">暂无优惠券</div>
      <div v-if="total > 0" class="mt-6 flex justify-center">
        <el-pagination v-model:current-page="page" :page-size="pageSize" :total="total" layout="prev, pager, next" background @current-change="fetchList" />
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { getMyCoupons } from '~/api/coupon.js'

definePageMeta({ middleware: 'auth' })

const list = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = 20
const activeTab = ref('0')

const onTabChange = () => { page.value = 1; fetchList() }

const fetchList = async () => {
  const res = await getMyCoupons({ page: page.value, pageSize, status: Number(activeTab.value) })
  if (res && res.code === 0) { list.value = res.data.list || []; total.value = res.data.total || 0 }
}

await fetchList()
</script>
