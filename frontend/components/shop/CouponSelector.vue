<template>
  <el-dialog v-model="visible" title="选择优惠券" width="500px">
    <div v-if="coupons.length" class="space-y-3">
      <div v-for="uc in coupons" :key="uc.ID" class="cursor-pointer" @click="selectCoupon(uc)">
        <ShopCouponCard :coupon="uc.coupon" :disabled="selectedId === uc.couponId" />
      </div>
    </div>
    <div v-else class="py-8 text-center text-gray-400">暂无可用优惠券</div>
    <template #footer>
      <el-button @click="visible = false">不使用优惠券</el-button>
    </template>
  </el-dialog>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { getAvailableCoupons } from '~/api/coupon.js'

const visible = defineModel({ type: Boolean, default: false })
const emit = defineEmits(['select'])

const coupons = ref([])
const selectedId = ref(null)

const selectCoupon = (uc) => {
  selectedId.value = uc.couponId
  emit('select', uc)
  visible.value = false
}

onMounted(async () => {
  const res = await getAvailableCoupons()
  if (res && res.code === 0) { coupons.value = res.data.list || [] }
})
</script>
