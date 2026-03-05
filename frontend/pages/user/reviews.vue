<template>
  <div class="bg-gray-50 min-h-screen">
    <div class="mx-auto max-w-4xl px-4 py-8 sm:px-6 lg:px-8">
      <h1 class="mb-6 text-2xl font-bold text-gray-900">我的评价</h1>
      <div class="space-y-4">
        <div v-for="review in list" :key="review.ID" class="rounded-lg bg-white p-4 shadow-sm ring-1 ring-gray-200">
          <div class="flex items-center gap-2 mb-2">
            <div class="flex text-orange-400">
              <span v-for="i in 5" :key="i" class="text-sm">{{ i <= review.rating ? '★' : '☆' }}</span>
            </div>
            <el-tag size="small" :type="review.status === 0 ? 'warning' : review.status === 1 ? 'success' : 'danger'">
              {{ review.status === 0 ? '待审核' : review.status === 1 ? '已通过' : '已拒绝' }}
            </el-tag>
            <span class="ml-auto text-xs text-gray-400">{{ formatTime(review.CreatedAt) }}</span>
          </div>
          <p class="text-sm text-gray-600">{{ review.content || '(无文字评价)' }}</p>
          <p class="mt-1 text-xs text-gray-400">订单号：{{ review.orderNo }}</p>
          <div v-if="review.reply" class="mt-3 rounded bg-gray-50 p-3">
            <span class="text-xs font-medium text-indigo-600">商家回复：</span>
            <span class="text-xs text-gray-600">{{ review.reply }}</span>
          </div>
        </div>
      </div>
      <div v-if="!list.length" class="py-16 text-center text-gray-400">暂无评价</div>
      <div v-if="total > 0" class="mt-6 flex justify-center">
        <el-pagination v-model:current-page="page" :page-size="pageSize" :total="total" layout="prev, pager, next" background @current-change="fetchList" />
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { getMyReviews } from '~/api/review.js'
import { formatTimeToStr } from '~/utils/date.js'

definePageMeta({ middleware: 'auth' })

const list = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = 20

const formatTime = (t) => t ? formatTimeToStr(t, 'yyyy-MM-dd hh:mm') : ''

const fetchList = async () => {
  const res = await getMyReviews({ page: page.value, pageSize })
  if (res && res.code === 0) { list.value = res.data.list || []; total.value = res.data.total || 0 }
}

await fetchList()
</script>
