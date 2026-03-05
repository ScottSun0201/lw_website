<template>
  <div class="bg-gray-50 min-h-screen">
    <div class="mx-auto max-w-4xl px-4 py-8 sm:px-6 lg:px-8">
      <h1 class="mb-6 text-2xl font-bold text-gray-900">退款记录</h1>
      <div class="space-y-4">
        <NuxtLink v-for="item in list" :key="item.ID" :to="`/user/refund/${item.refundNo}`" class="block rounded-lg bg-white p-4 shadow-sm ring-1 ring-gray-200 hover:shadow-md transition-shadow">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm text-gray-500">退款单号：{{ item.refundNo }}</span>
            <el-tag size="small" :type="refundStatusTag(item.status)">{{ refundStatusLabel(item.status) }}</el-tag>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">订单：{{ item.orderNo }}</span>
            <span class="text-base font-bold text-red-500">¥{{ item.amount?.toFixed(2) }}</span>
          </div>
          <p class="mt-1 text-xs text-gray-400">{{ item.reason }}</p>
        </NuxtLink>
      </div>
      <div v-if="!list.length" class="py-16 text-center text-gray-400">暂无退款记录</div>
      <div v-if="total > 0" class="mt-6 flex justify-center">
        <el-pagination v-model:current-page="page" :page-size="pageSize" :total="total" layout="prev, pager, next" background @current-change="fetchList" />
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { getUserRefunds } from '~/api/refund.js'

definePageMeta({ middleware: 'auth' })

const list = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = 20

const refundStatusLabel = (s) => ({ 0: '待审核', 1: '审核通过', 2: '已拒绝', 3: '退款中', 4: '已退款', 5: '已取消' }[s] || '未知')
const refundStatusTag = (s) => ({ 0: 'warning', 1: 'primary', 2: 'danger', 3: '', 4: 'success', 5: 'info' }[s] || 'info')

const fetchList = async () => {
  const res = await getUserRefunds({ page: page.value, pageSize })
  if (res && res.code === 0) { list.value = res.data.list || []; total.value = res.data.total || 0 }
}

await fetchList()
</script>
