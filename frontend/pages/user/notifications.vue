<template>
  <div class="bg-gray-50 min-h-screen">
    <div class="mx-auto max-w-4xl px-4 py-8 sm:px-6 lg:px-8">
      <div class="mb-6 flex items-center justify-between">
        <h1 class="text-2xl font-bold text-gray-900">我的消息</h1>
        <el-button type="primary" text @click="handleMarkAllRead">全部标记已读</el-button>
      </div>
      <el-tabs v-model="activeTab" @tab-change="onTabChange">
        <el-tab-pane label="全部" name="all" />
        <el-tab-pane label="订单" name="1" />
        <el-tab-pane label="物流" name="2" />
        <el-tab-pane label="系统" name="4" />
      </el-tabs>
      <div class="space-y-3 mt-4">
        <div v-for="item in list" :key="item.ID"
          class="flex items-start gap-3 rounded-lg bg-white p-4 shadow-sm ring-1 ring-gray-200"
          :class="item.status === 0 ? 'border-l-4 border-indigo-500' : ''"
        >
          <div class="flex-1">
            <div class="flex items-center gap-2">
              <span class="text-sm font-medium text-gray-900">{{ item.title }}</span>
              <span v-if="item.status === 0" class="h-2 w-2 rounded-full bg-red-500"></span>
            </div>
            <p class="mt-1 text-sm text-gray-500">{{ item.content }}</p>
            <p class="mt-1 text-xs text-gray-400">{{ formatTime(item.CreatedAt) }}</p>
          </div>
          <el-button v-if="item.status === 0" type="primary" text size="small" @click="handleMarkRead(item.ID)">标记已读</el-button>
        </div>
      </div>
      <div v-if="!list.length" class="py-16 text-center text-gray-400">暂无消息</div>
      <div v-if="total > 0" class="mt-6 flex justify-center">
        <el-pagination v-model:current-page="page" :page-size="pageSize" :total="total" layout="prev, pager, next" background @current-change="fetchList" />
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { getMyNotifications, markRead, markAllRead } from '~/api/notification.js'
import { useNotificationStore } from '~/stores/notificationStore.js'
import { formatTimeToStr } from '~/utils/date.js'

definePageMeta({ middleware: 'auth' })

const notificationStore = useNotificationStore()
const list = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = 20
const activeTab = ref('all')

const formatTime = (t) => t ? formatTimeToStr(t, 'yyyy-MM-dd hh:mm') : ''

const fetchList = async () => {
  const params = { page: page.value, pageSize }
  if (activeTab.value !== 'all') params.type = Number(activeTab.value)
  const res = await getMyNotifications(params)
  if (res && res.code === 0) { list.value = res.data.list || []; total.value = res.data.total || 0 }
}

const onTabChange = () => { page.value = 1; fetchList() }

const handleMarkRead = async (id) => {
  const res = await markRead({ id })
  if (res && res.code === 0) { fetchList(); notificationStore.refreshCount() }
}

const handleMarkAllRead = async () => {
  const res = await markAllRead()
  if (res && res.code === 0) { fetchList(); notificationStore.refreshCount(); ElMessage.success('已全部标记已读') }
}

await fetchList()
</script>
