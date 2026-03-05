<template>
  <div class="bg-gray-50 min-h-screen">
    <div class="mx-auto max-w-4xl px-4 py-8 sm:px-6 lg:px-8">
      <!-- Back link -->
      <NuxtLink to="/shop/order" class="mb-6 inline-flex items-center gap-1 text-sm text-gray-500 hover:text-indigo-600">
        <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
        </svg>
        返回订单列表
      </NuxtLink>

      <div v-if="order">
        <!-- Order status banner -->
        <div :class="[
          'mb-6 rounded-lg p-6 shadow-sm',
          order.status === 0 ? 'bg-orange-50 ring-1 ring-orange-200' :
          order.status === 5 ? 'bg-gray-50 ring-1 ring-gray-200' :
          'bg-green-50 ring-1 ring-green-200'
        ]">
          <div class="flex items-center justify-between">
            <div>
              <el-tag :type="statusTagType(order.status)" size="large">
                {{ statusLabel(order.status) }}
              </el-tag>
              <p class="mt-2 text-sm text-gray-600">订单号：{{ order.orderNo }}</p>
              <p class="text-sm text-gray-500">创建时间：{{ formatTime(order.CreatedAt) }}</p>
            </div>
            <div class="flex gap-2">
              <el-button
                v-if="order.status === 0"
                type="danger"
                @click="handlePay"
              >
                立即支付
              </el-button>
              <el-button
                v-if="order.status === 0"
                @click="handleCancel"
              >
                取消订单
              </el-button>
              <el-button
                v-if="order.status === 2"
                type="primary"
                @click="handleConfirmReceive"
              >
                确认收货
              </el-button>
              <el-button
                v-if="order.status >= 1 && order.status <= 3"
                type="warning"
                @click="showRefundForm = true"
              >
                申请退款
              </el-button>
              <el-button
                v-if="order.status >= 3 && order.status <= 4"
                type="primary"
                @click="showReviewForm = true"
              >
                评价
              </el-button>
            </div>
          </div>
        </div>

        <!-- Receiver info -->
        <div class="mb-6 rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
          <h2 class="mb-3 text-lg font-semibold text-gray-900">收货信息</h2>
          <div class="grid grid-cols-1 gap-3 sm:grid-cols-2">
            <div>
              <span class="text-sm text-gray-500">收货人：</span>
              <span class="text-sm font-medium text-gray-900">{{ order.receiverName }}</span>
            </div>
            <div>
              <span class="text-sm text-gray-500">联系电话：</span>
              <span class="text-sm font-medium text-gray-900">{{ order.receiverPhone }}</span>
            </div>
            <div class="sm:col-span-2">
              <span class="text-sm text-gray-500">收货地址：</span>
              <span class="text-sm font-medium text-gray-900">{{ order.receiverAddress }}</span>
            </div>
            <div v-if="order.shipCompany">
              <span class="text-sm text-gray-500">物流公司：</span>
              <span class="text-sm font-medium text-gray-900">{{ order.shipCompany }}</span>
            </div>
            <div v-if="order.shipNo">
              <span class="text-sm text-gray-500">物流单号：</span>
              <span class="text-sm font-medium text-gray-900">{{ order.shipNo }}</span>
            </div>
          </div>
        </div>

        <!-- Logistics info -->
        <div v-if="order && order.status >= 2" class="mb-6 rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
          <h2 class="mb-4 text-lg font-semibold text-gray-900">物流信息</h2>
          <ShopLogisticsTimeline :orderNo="order.orderNo" />
        </div>

        <!-- Order items -->
        <div class="mb-6 rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
          <h2 class="mb-4 text-lg font-semibold text-gray-900">商品清单</h2>
          <div class="divide-y divide-gray-100">
            <div v-for="item in order.items" :key="item.ID" class="flex gap-4 py-4">
              <div class="h-20 w-20 flex-shrink-0 overflow-hidden rounded-lg bg-gray-100">
                <img :src="getUrl(item.skuImage)" :alt="item.spuName" class="h-full w-full object-cover" />
              </div>
              <div class="flex flex-1 items-center justify-between">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ item.spuName }}</div>
                  <div v-if="item.skuName" class="mt-0.5 text-xs text-gray-500">{{ item.skuName }}</div>
                  <div v-if="item.specData" class="mt-0.5 text-xs text-gray-400">{{ formatSpec(item.specData) }}</div>
                </div>
                <div class="text-right">
                  <div class="text-sm text-gray-900">&yen;{{ item.price?.toFixed(2) }} x {{ item.quantity }}</div>
                  <div class="mt-1 text-sm font-medium text-red-500">&yen;{{ item.totalPrice?.toFixed(2) }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Amount breakdown -->
        <div class="mb-6 rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
          <h2 class="mb-4 text-lg font-semibold text-gray-900">金额明细</h2>
          <div class="space-y-3">
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">商品合计</span>
              <span class="text-gray-900">&yen;{{ order.totalAmount?.toFixed(2) }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">运费</span>
              <span class="text-gray-900">&yen;{{ order.freightAmount?.toFixed(2) }}</span>
            </div>
            <div v-if="order.discountAmount > 0" class="flex justify-between text-sm">
              <span class="text-gray-500">优惠</span>
              <span class="text-green-600">-&yen;{{ order.discountAmount?.toFixed(2) }}</span>
            </div>
            <el-divider class="!my-3" />
            <div class="flex justify-between">
              <span class="text-base font-medium text-gray-900">实付金额</span>
              <span class="text-2xl font-bold text-red-500">&yen;{{ order.payAmount?.toFixed(2) }}</span>
            </div>
            <div v-if="order.payMethod" class="text-right text-xs text-gray-400">
              支付方式：{{ order.payMethod }}
            </div>
          </div>
        </div>

        <!-- Timeline info -->
        <div class="mb-6 rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
          <h2 class="mb-4 text-lg font-semibold text-gray-900">订单时间线</h2>
          <div class="space-y-3">
            <div class="flex items-center gap-3 text-sm">
              <span class="w-28 text-gray-500">创建时间</span>
              <span class="text-gray-900">{{ formatTime(order.CreatedAt) }}</span>
            </div>
            <div v-if="order.payTime" class="flex items-center gap-3 text-sm">
              <span class="w-28 text-gray-500">支付时间</span>
              <span class="text-gray-900">{{ formatTime(order.payTime) }}</span>
            </div>
            <div v-if="order.shipTime" class="flex items-center gap-3 text-sm">
              <span class="w-28 text-gray-500">发货时间</span>
              <span class="text-gray-900">{{ formatTime(order.shipTime) }}</span>
            </div>
            <div v-if="order.receiveTime" class="flex items-center gap-3 text-sm">
              <span class="w-28 text-gray-500">收货时间</span>
              <span class="text-gray-900">{{ formatTime(order.receiveTime) }}</span>
            </div>
            <div v-if="order.completeTime" class="flex items-center gap-3 text-sm">
              <span class="w-28 text-gray-500">完成时间</span>
              <span class="text-gray-900">{{ formatTime(order.completeTime) }}</span>
            </div>
            <div v-if="order.cancelTime" class="flex items-center gap-3 text-sm">
              <span class="w-28 text-gray-500">取消时间</span>
              <span class="text-gray-900">{{ formatTime(order.cancelTime) }}</span>
            </div>
          </div>
        </div>

        <!-- Remark -->
        <div v-if="order.remark" class="mb-6 rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
          <h2 class="mb-2 text-lg font-semibold text-gray-900">买家备注</h2>
          <p class="text-sm text-gray-600">{{ order.remark }}</p>
        </div>
      </div>

      <ShopRefundForm v-model="showRefundForm" :orderNo="order?.orderNo" :maxAmount="order?.payAmount" @submitted="fetchOrder" />
      <ShopReviewForm v-model="showReviewForm" :orderNo="order?.orderNo" @submitted="fetchOrder" />

      <!-- Pay dialog -->
      <el-dialog v-model="payDialogVisible" title="订单支付" width="420px">
        <div v-if="order" class="space-y-4">
          <div class="text-center">
            <p class="text-sm text-gray-500">订单号：{{ order.orderNo }}</p>
            <p class="mt-2 text-3xl font-bold text-red-500">&yen;{{ order.payAmount?.toFixed(2) }}</p>
          </div>
          <div class="rounded-lg bg-gray-50 p-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">钱包余额</span>
              <span class="font-medium text-gray-900">&yen;{{ walletBalance.toFixed(2) }}</span>
            </div>
          </div>
          <div v-if="walletBalance < (order.payAmount || 0)" class="text-center text-sm text-red-500">
            余额不足
          </div>
        </div>
        <template #footer>
          <div class="flex justify-end gap-3">
            <el-button @click="payDialogVisible = false">取消</el-button>
            <el-button
              type="danger"
              :loading="paying"
              :disabled="walletBalance < (order?.payAmount || 0)"
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
import { useRoute, useRouter } from 'vue-router'
import { getUserOrder, cancelOrder, confirmReceive, balancePay, getWallet } from '~/api/shop.js'
import { getTraces } from '~/api/logistics.js'
import { getUrl } from '~/utils/image.js'
import { formatTimeToStr } from '~/utils/date.js'

definePageMeta({ middleware: 'auth' })

const route = useRoute()
const router = useRouter()

const order = ref(null)
const payDialogVisible = ref(false)
const walletBalance = ref(0)
const paying = ref(false)
const showRefundForm = ref(false)
const showReviewForm = ref(false)
const logistics = ref([])

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
  return formatTimeToStr(time, 'yyyy-MM-dd hh:mm:ss')
}

const formatSpec = (specData) => {
  try {
    const specs = JSON.parse(specData)
    return Object.entries(specs).map(([k, v]) => `${k}: ${v}`).join(', ')
  } catch (e) {
    return specData
  }
}

const fetchOrder = async () => {
  const res = await getUserOrder({ orderNo: route.params.id })
  if (res && res.code === 0) {
    order.value = res.data.data
  }
}

await fetchOrder()

const handlePay = async () => {
  const res = await getWallet()
  if (res && res.code === 0) {
    walletBalance.value = res.data.balance || 0
  }
  payDialogVisible.value = true
}

const confirmPay = async () => {
  if (!order.value) return
  paying.value = true
  try {
    const res = await balancePay({ orderNo: order.value.orderNo })
    if (res && res.code === 0) {
      ElMessage.success('支付成功')
      payDialogVisible.value = false
      await fetchOrder()
    }
  } finally {
    paying.value = false
  }
}

const handleCancel = async () => {
  try {
    await ElMessageBox.confirm('确定要取消该订单吗？', '取消订单', {
      confirmButtonText: '确定',
      cancelButtonText: '返回',
      type: 'warning'
    })
    const res = await cancelOrder({ orderNo: order.value.orderNo })
    if (res && res.code === 0) {
      ElMessage.success('订单已取消')
      await fetchOrder()
    }
  } catch (e) {
    // user cancelled dialog
  }
}

const handleConfirmReceive = async () => {
  try {
    await ElMessageBox.confirm('确认已收到商品？', '确认收货', {
      confirmButtonText: '确定',
      cancelButtonText: '返回',
      type: 'info'
    })
    const res = await confirmReceive({ orderNo: order.value.orderNo })
    if (res && res.code === 0) {
      ElMessage.success('已确认收货')
      await fetchOrder()
    }
  } catch (e) {
    // user cancelled dialog
  }
}
</script>
