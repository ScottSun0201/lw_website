<template>
  <div class="bg-gray-50 min-h-screen">
    <div class="mx-auto max-w-5xl px-4 py-8 sm:px-6 lg:px-8">
      <h1 class="mb-8 text-2xl font-bold text-gray-900">我的订单</h1>

      <!-- Status tabs -->
      <div class="mb-6">
        <div class="flex gap-1 overflow-x-auto rounded-lg bg-white p-1 shadow-sm ring-1 ring-gray-200">
          <button
            v-for="tab in statusTabs"
            :key="tab.value"
            :class="[
              'flex-shrink-0 rounded-md px-4 py-2 text-sm font-medium transition-colors',
              currentStatus === tab.value
                ? 'bg-indigo-600 text-white'
                : 'text-gray-600 hover:bg-gray-100'
            ]"
            @click="switchTab(tab.value)"
          >
            {{ tab.label }}
          </button>
        </div>
      </div>

      <!-- Order list -->
      <div v-if="orders.length > 0" class="space-y-4">
        <div
          v-for="order in orders"
          :key="order.ID"
          class="cursor-pointer rounded-lg bg-white p-5 shadow-sm ring-1 ring-gray-200 transition-shadow hover:shadow-md"
          @click="goOrderDetail(order.orderNo)"
        >
          <!-- Order header -->
          <div class="mb-4 flex items-center justify-between">
            <div class="flex items-center gap-3">
              <span class="text-sm text-gray-500">订单号：{{ order.orderNo }}</span>
              <span class="text-xs text-gray-400">{{ formatTime(order.CreatedAt) }}</span>
            </div>
            <el-tag :type="statusTagType(order.status)" size="small">
              {{ statusLabel(order.status) }}
            </el-tag>
          </div>

          <!-- Items preview (first 2) -->
          <div class="mb-4 flex gap-3 overflow-hidden">
            <div
              v-for="item in order.items?.slice(0, 2)"
              :key="item.ID"
              class="flex gap-3"
            >
              <div class="h-16 w-16 flex-shrink-0 overflow-hidden rounded-lg bg-gray-100">
                <img :src="getUrl(item.skuImage)" :alt="item.spuName" class="h-full w-full object-cover" />
              </div>
              <div class="flex flex-col justify-center min-w-0">
                <span class="text-sm text-gray-900 line-clamp-1">{{ item.spuName }}</span>
                <span v-if="item.specData" class="text-xs text-gray-400">{{ formatSpec(item.specData) }}</span>
                <span class="text-xs text-gray-500">x{{ item.quantity }}</span>
              </div>
            </div>
            <div v-if="order.items && order.items.length > 2" class="flex items-center text-sm text-gray-400">
              ... 共 {{ order.items.length }} 件商品
            </div>
          </div>

          <!-- Order footer -->
          <div class="flex items-center justify-between border-t border-gray-100 pt-3">
            <div class="text-sm text-gray-600">
              合计：<span class="text-lg font-bold text-red-500">&yen;{{ order.payAmount?.toFixed(2) }}</span>
            </div>
            <div class="flex gap-2" @click.stop>
              <el-button
                v-if="order.status === 0"
                type="danger"
                size="small"
                @click.stop="handlePay(order)"
              >
                立即支付
              </el-button>
              <el-button
                v-if="order.status === 0"
                size="small"
                @click.stop="handleCancel(order)"
              >
                取消订单
              </el-button>
              <el-button
                v-if="order.status === 2"
                type="primary"
                size="small"
                @click.stop="handleConfirmReceive(order)"
              >
                确认收货
              </el-button>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty state -->
      <div v-else class="flex flex-col items-center justify-center py-20 text-gray-400">
        <svg class="mb-4 h-16 w-16" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
        </svg>
        <p class="text-lg">暂无订单</p>
        <NuxtLink to="/shop" class="mt-4">
          <el-button type="primary">去逛逛</el-button>
        </NuxtLink>
      </div>

      <!-- Pagination -->
      <div v-if="total > 0" class="mt-8 flex justify-center">
        <el-pagination
          v-model:current-page="currentPage"
          :page-size="pageSize"
          :total="total"
          layout="prev, pager, next"
          background
          @current-change="handlePageChange"
        />
      </div>

      <!-- Pay dialog -->
      <el-dialog v-model="payDialogVisible" title="订单支付" width="420px">
        <div v-if="payingOrder" class="space-y-4">
          <div class="text-center">
            <p class="text-sm text-gray-500">订单号：{{ payingOrder.orderNo }}</p>
            <p class="mt-2 text-3xl font-bold text-red-500">&yen;{{ payingOrder.payAmount?.toFixed(2) }}</p>
          </div>
          <div class="rounded-lg bg-gray-50 p-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">钱包余额</span>
              <span class="font-medium text-gray-900">&yen;{{ walletBalance.toFixed(2) }}</span>
            </div>
          </div>
          <div v-if="walletBalance < (payingOrder.payAmount || 0)" class="text-center text-sm text-red-500">
            余额不足
          </div>
        </div>
        <template #footer>
          <div class="flex justify-end gap-3">
            <el-button @click="payDialogVisible = false">取消</el-button>
            <el-button
              type="danger"
              :loading="paying"
              :disabled="walletBalance < (payingOrder?.payAmount || 0)"
              @click="confirmPay"
            >
              立即支付
            </el-button>
          </div>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { getUserOrderList, cancelOrder, confirmReceive, balancePay, getWallet } from '~/api/shop.js'
import { getUrl } from '~/utils/image.js'
import { formatTimeToStr } from '~/utils/date.js'

definePageMeta({ middleware: 'auth' })

const router = useRouter()

const pageSize = 10
const currentPage = ref(1)
const currentStatus = ref(null)
const orders = ref([])
const total = ref(0)
const payDialogVisible = ref(false)
const payingOrder = ref(null)
const walletBalance = ref(0)
const paying = ref(false)

const statusTabs = [
  { label: '全部', value: null },
  { label: '待支付', value: 0 },
  { label: '待发货', value: 1 },
  { label: '已发货', value: 2 },
  { label: '已收货', value: 3 },
  { label: '已完成', value: 4 },
  { label: '已取消', value: 5 }
]

const statusMap = {
  0: '待支付',
  1: '待发货',
  2: '已发货',
  3: '已收货',
  4: '已完成',
  5: '已取消',
  6: '退款中',
  7: '已退款'
}

const statusTagTypeMap = {
  0: 'warning',
  1: 'primary',
  2: 'primary',
  3: 'success',
  4: 'success',
  5: 'info',
  6: 'danger',
  7: 'info'
}

const statusLabel = (status) => statusMap[status] || '未知'
const statusTagType = (status) => statusTagTypeMap[status] || 'info'

const formatTime = (time) => {
  if (!time) return ''
  return formatTimeToStr(time, 'yyyy-MM-dd hh:mm')
}

const formatSpec = (specData) => {
  try {
    const specs = JSON.parse(specData)
    return Object.entries(specs).map(([k, v]) => `${k}: ${v}`).join(', ')
  } catch (e) {
    return specData
  }
}

const fetchOrders = async () => {
  const params = {
    page: currentPage.value,
    pageSize
  }
  if (currentStatus.value !== null) {
    params.status = currentStatus.value
  }
  const res = await getUserOrderList(params)
  if (res && res.code === 0) {
    orders.value = res.data.list || []
    total.value = res.data.total || 0
  }
}

await fetchOrders()

const switchTab = (status) => {
  currentStatus.value = status
  currentPage.value = 1
  fetchOrders()
}

const handlePageChange = (page) => {
  currentPage.value = page
  fetchOrders()
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const goOrderDetail = (orderNo) => {
  router.push(`/shop/order/${orderNo}`)
}

const handlePay = async (order) => {
  payingOrder.value = order
  // Load wallet balance
  const res = await getWallet()
  if (res && res.code === 0) {
    walletBalance.value = res.data.balance || 0
  }
  payDialogVisible.value = true
}

const confirmPay = async () => {
  if (!payingOrder.value) return
  paying.value = true
  try {
    const res = await balancePay({ orderNo: payingOrder.value.orderNo })
    if (res && res.code === 0) {
      ElMessage.success('支付成功')
      payDialogVisible.value = false
      payingOrder.value = null
      await fetchOrders()
    }
  } finally {
    paying.value = false
  }
}

const handleCancel = async (order) => {
  try {
    await ElMessageBox.confirm('确定要取消该订单吗？', '取消订单', {
      confirmButtonText: '确定',
      cancelButtonText: '返回',
      type: 'warning'
    })
    const res = await cancelOrder({ orderNo: order.orderNo })
    if (res && res.code === 0) {
      ElMessage.success('订单已取消')
      await fetchOrders()
    }
  } catch (e) {
    // user cancelled dialog
  }
}

const handleConfirmReceive = async (order) => {
  try {
    await ElMessageBox.confirm('确认已收到商品？', '确认收货', {
      confirmButtonText: '确定',
      cancelButtonText: '返回',
      type: 'info'
    })
    const res = await confirmReceive({ orderNo: order.orderNo })
    if (res && res.code === 0) {
      ElMessage.success('已确认收货')
      await fetchOrders()
    }
  } catch (e) {
    // user cancelled dialog
  }
}
</script>
