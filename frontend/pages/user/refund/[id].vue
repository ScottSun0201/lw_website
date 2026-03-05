<template>
  <div class="bg-gray-50 min-h-screen">
    <div class="mx-auto max-w-4xl px-4 py-8 sm:px-6 lg:px-8">
      <NuxtLink to="/user/refunds" class="mb-6 inline-flex items-center gap-1 text-sm text-gray-500 hover:text-indigo-600">
        <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" /></svg>
        返回退款列表
      </NuxtLink>

      <div v-if="refund">
        <div class="mb-6 rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
          <div class="flex items-center justify-between mb-4">
            <el-tag size="large" :type="refundStatusTag(refund.status)">{{ refundStatusLabel(refund.status) }}</el-tag>
            <span class="text-2xl font-bold text-red-500">¥{{ refund.amount?.toFixed(2) }}</span>
          </div>
          <div class="grid grid-cols-2 gap-3 text-sm">
            <div><span class="text-gray-500">退款单号：</span>{{ refund.refundNo }}</div>
            <div><span class="text-gray-500">订单号：</span>{{ refund.orderNo }}</div>
            <div><span class="text-gray-500">退款类型：</span>{{ refund.type === 1 ? '仅退款' : '退货退款' }}</div>
            <div><span class="text-gray-500">退款原因：</span>{{ refund.reason }}</div>
            <div v-if="refund.description" class="col-span-2"><span class="text-gray-500">详细说明：</span>{{ refund.description }}</div>
            <div v-if="refund.adminRemark" class="col-span-2"><span class="text-gray-500">管理员备注：</span>{{ refund.adminRemark }}</div>
          </div>
          <div class="mt-4 flex gap-2">
            <el-button v-if="refund.status === 0 || refund.status === 1" @click="handleCancel">取消退款</el-button>
            <el-button v-if="refund.status === 1 && refund.type === 2" type="primary" @click="showShipDialog = true">填写退货物流</el-button>
          </div>
        </div>

        <!-- Timeline -->
        <div class="rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
          <h2 class="mb-4 text-lg font-semibold text-gray-900">退款进度</h2>
          <el-timeline v-if="logs.length">
            <el-timeline-item v-for="(log, idx) in logs" :key="idx" :timestamp="formatTime(log.CreatedAt)" placement="top">
              {{ log.remark }}
            </el-timeline-item>
          </el-timeline>
        </div>
      </div>

      <!-- Ship dialog -->
      <el-dialog v-model="showShipDialog" title="填写退货物流" width="450px">
        <el-form :model="shipForm" label-width="100px">
          <el-form-item label="物流公司"><el-input v-model="shipForm.shipCompany" placeholder="请输入物流公司" /></el-form-item>
          <el-form-item label="物流单号"><el-input v-model="shipForm.shipNo" placeholder="请输入物流单号" /></el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="showShipDialog = false">取消</el-button>
          <el-button type="primary" @click="handleShipReturn">提交</el-button>
        </template>
      </el-dialog>
    </div>
  </div>
</template>
<script setup>
import { ref, reactive } from 'vue'
import { useRoute } from 'vue-router'
import { getRefundDetail, cancelRefund, shipReturn } from '~/api/refund.js'
import { formatTimeToStr } from '~/utils/date.js'

definePageMeta({ middleware: 'auth' })

const route = useRoute()
const refund = ref(null)
const logs = ref([])
const showShipDialog = ref(false)
const shipForm = reactive({ shipCompany: '', shipNo: '' })

const refundStatusLabel = (s) => ({ 0: '待审核', 1: '审核通过', 2: '已拒绝', 3: '退款中', 4: '已退款', 5: '已取消' }[s] || '未知')
const refundStatusTag = (s) => ({ 0: 'warning', 1: 'primary', 2: 'danger', 3: '', 4: 'success', 5: 'info' }[s] || 'info')
const formatTime = (t) => t ? formatTimeToStr(t, 'yyyy-MM-dd hh:mm:ss') : ''

const fetchDetail = async () => {
  const res = await getRefundDetail({ refundNo: route.params.id })
  if (res && res.code === 0) { refund.value = res.data.data; logs.value = res.data.logs || [] }
}

const handleCancel = async () => {
  try {
    await ElMessageBox.confirm('确定要取消退款吗？', '取消退款', { type: 'warning' })
    const res = await cancelRefund({ refundNo: refund.value.refundNo })
    if (res && res.code === 0) { ElMessage.success('已取消'); fetchDetail() }
  } catch {}
}

const handleShipReturn = async () => {
  const res = await shipReturn({ refundNo: refund.value.refundNo, ...shipForm })
  if (res && res.code === 0) { ElMessage.success('提交成功'); showShipDialog.value = false; fetchDetail() }
}

await fetchDetail()
</script>
