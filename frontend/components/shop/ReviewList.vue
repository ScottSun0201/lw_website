<template>
  <div>
    <!-- Review stats -->
    <div v-if="stat" class="mb-6 flex items-center gap-6 rounded-lg bg-gray-50 p-4">
      <div class="text-center">
        <div class="text-3xl font-bold text-orange-500">{{ stat.avgRating?.toFixed(1) || '0.0' }}</div>
        <div class="text-sm text-gray-500">平均评分</div>
      </div>
      <div class="flex-1 space-y-1">
        <div v-for="star in [5,4,3,2,1]" :key="star" class="flex items-center gap-2 text-sm">
          <span class="w-8 text-gray-500">{{ star }}星</span>
          <div class="h-2 flex-1 rounded-full bg-gray-200">
            <div class="h-full rounded-full bg-orange-400" :style="{ width: barWidth(star) + '%' }" />
          </div>
          <span class="w-8 text-right text-gray-400">{{ starCount(star) }}</span>
        </div>
      </div>
    </div>

    <!-- Review list -->
    <div v-if="reviews.length" class="space-y-4">
      <div v-for="review in reviews" :key="review.ID" class="rounded-lg border border-gray-100 p-4">
        <div class="mb-2 flex items-center gap-3">
          <img v-if="review.avatar" :src="getUrl(review.avatar)" class="h-8 w-8 rounded-full object-cover" />
          <div v-else class="flex h-8 w-8 items-center justify-center rounded-full bg-gray-200 text-sm text-gray-500">{{ (review.nickname || '匿名')[0] }}</div>
          <span class="text-sm font-medium text-gray-700">{{ review.nickname || '匿名用户' }}</span>
          <div class="flex text-orange-400">
            <span v-for="i in 5" :key="i" class="text-sm">{{ i <= review.rating ? '★' : '☆' }}</span>
          </div>
          <span class="ml-auto text-xs text-gray-400">{{ formatTime(review.CreatedAt) }}</span>
        </div>
        <p class="text-sm text-gray-600">{{ review.content }}</p>
        <div v-if="review.images" class="mt-2 flex gap-2">
          <img v-for="(img, idx) in parseImages(review.images)" :key="idx" :src="getUrl(img)" class="h-16 w-16 rounded object-cover" />
        </div>
        <div v-if="review.reply" class="mt-3 rounded bg-gray-50 p-3">
          <span class="text-xs font-medium text-indigo-600">商家回复：</span>
          <span class="text-xs text-gray-600">{{ review.reply }}</span>
        </div>
      </div>
    </div>
    <div v-else class="py-8 text-center text-gray-400">暂无评价</div>

    <div v-if="total > pageSize" class="mt-4 flex justify-center">
      <el-pagination v-model:current-page="page" :page-size="pageSize" :total="total" layout="prev, pager, next" background @current-change="fetchReviews" />
    </div>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { getProductReviews, getReviewStat } from '~/api/review.js'
import { getUrl } from '~/utils/image.js'
import { formatTimeToStr } from '~/utils/date.js'

const props = defineProps({ spuId: { type: Number, required: true } })

const reviews = ref([])
const stat = ref(null)
const total = ref(0)
const page = ref(1)
const pageSize = 10

const formatTime = (t) => t ? formatTimeToStr(t, 'yyyy-MM-dd hh:mm') : ''
const parseImages = (s) => { try { return JSON.parse(s) } catch { return [] } }
const starCount = (star) => stat.value ? stat.value[`star${star}Count`] || 0 : 0
const barWidth = (star) => {
  if (!stat.value || !stat.value.totalCount) return 0
  return (starCount(star) / stat.value.totalCount) * 100
}

const fetchReviews = async () => {
  const res = await getProductReviews({ spuId: props.spuId, page: page.value, pageSize })
  if (res && res.code === 0) { reviews.value = res.data.list || []; total.value = res.data.total || 0 }
}

const fetchStat = async () => {
  const res = await getReviewStat({ spuId: props.spuId })
  if (res && res.code === 0) { stat.value = res.data.data }
}

onMounted(() => { fetchStat(); fetchReviews() })
</script>
