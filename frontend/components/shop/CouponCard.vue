<template>
  <div class="flex overflow-hidden rounded-lg border" :class="disabled ? 'border-gray-200 bg-gray-50' : 'border-red-200 bg-red-50'">
    <div class="flex w-24 flex-col items-center justify-center p-3" :class="disabled ? 'bg-gray-200 text-gray-500' : 'bg-red-500 text-white'">
      <span class="text-xl font-bold" v-if="coupon.type === 1 || coupon.type === 3">¥{{ coupon.value }}</span>
      <span class="text-xl font-bold" v-else>{{ coupon.value }}折</span>
      <span class="text-xs" v-if="coupon.minAmount > 0">满{{ coupon.minAmount }}可用</span>
      <span class="text-xs" v-else>无门槛</span>
    </div>
    <div class="flex flex-1 items-center justify-between p-3">
      <div>
        <div class="text-sm font-medium" :class="disabled ? 'text-gray-500' : 'text-gray-900'">{{ coupon.name }}</div>
        <div class="mt-1 text-xs text-gray-400" v-if="coupon.endTime">有效期至 {{ formatDate(coupon.endTime) }}</div>
      </div>
      <slot name="action">
        <el-button v-if="!disabled && showClaim" type="danger" size="small" @click="$emit('claim')">领取</el-button>
      </slot>
    </div>
  </div>
</template>
<script setup>
defineProps({
  coupon: { type: Object, required: true },
  disabled: { type: Boolean, default: false },
  showClaim: { type: Boolean, default: false }
})
defineEmits(['claim'])
const formatDate = (d) => d ? new Date(d).toLocaleDateString('zh-CN') : ''
</script>
