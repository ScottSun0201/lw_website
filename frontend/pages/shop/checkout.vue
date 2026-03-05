<template>
  <div class="bg-gray-50 min-h-screen">
    <div class="mx-auto max-w-4xl px-4 py-8 sm:px-6 lg:px-8">
      <h1 class="mb-8 text-2xl font-bold text-gray-900">确认订单</h1>

      <!-- Address section -->
      <div class="mb-6 rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
        <div class="mb-4 flex items-center justify-between">
          <h2 class="text-lg font-semibold text-gray-900">收货地址</h2>
          <el-button type="primary" text @click="showAddressForm = true; editingAddress = null">
            + 新增地址
          </el-button>
        </div>
        <ShopAddressSelector
          v-model="selectedAddressId"
          :addresses="addressList"
        />
      </div>

      <!-- Order items -->
      <div class="mb-6 rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
        <h2 class="mb-4 text-lg font-semibold text-gray-900">商品清单</h2>
        <div class="divide-y divide-gray-100">
          <div v-for="item in orderItems" :key="item.ID" class="flex gap-4 py-4">
            <div class="h-20 w-20 flex-shrink-0 overflow-hidden rounded-lg bg-gray-100">
              <img :src="getUrl(item.skuImage || item.spuImage)" :alt="item.spuName" class="h-full w-full object-cover" />
            </div>
            <div class="flex flex-1 items-center justify-between">
              <div>
                <div class="text-sm font-medium text-gray-900">{{ item.spuName }}</div>
                <div v-if="item.specData" class="mt-1 text-xs text-gray-400">{{ formatSpec(item.specData) }}</div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-red-500">&yen;{{ item.skuPrice?.toFixed(2) }}</div>
                <div class="text-xs text-gray-400">x {{ item.quantity }}</div>
              </div>
            </div>
          </div>
        </div>
        <div v-if="!orderItems.length" class="py-8 text-center text-gray-400">
          暂无商品
        </div>
      </div>

      <!-- Coupon section -->
      <div class="mb-6 rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
        <div class="flex items-center justify-between">
          <h2 class="text-lg font-semibold text-gray-900">优惠券</h2>
          <el-button type="primary" text @click="loadCoupons">选择优惠券</el-button>
        </div>
        <div v-if="selectedCoupon" class="mt-3 flex items-center justify-between rounded-md bg-orange-50 p-3">
          <span class="text-sm text-orange-700">{{ selectedCoupon.name }} - 优惠 ¥{{ selectedCoupon.discountAmount?.toFixed(2) }}</span>
          <el-button type="danger" text size="small" @click="selectedCoupon = null">取消</el-button>
        </div>
        <div v-else class="mt-2 text-sm text-gray-400">暂未选择优惠券</div>
      </div>
      <ShopCouponSelector v-model="showCouponSelector" :coupons="availableCoupons" @select="onCouponSelect" />

      <!-- Remark -->
      <div class="mb-6 rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
        <h2 class="mb-3 text-lg font-semibold text-gray-900">订单备注</h2>
        <el-input
          v-model="remark"
          type="textarea"
          :rows="2"
          placeholder="给卖家留言（选填）"
          maxlength="200"
          show-word-limit
        />
      </div>

      <!-- Amount summary -->
      <div class="mb-6 rounded-lg bg-white p-6 shadow-sm ring-1 ring-gray-200">
        <h2 class="mb-4 text-lg font-semibold text-gray-900">订单金额</h2>
        <div class="space-y-3">
          <div class="flex justify-between text-sm">
            <span class="text-gray-500">商品合计</span>
            <span class="text-gray-900">&yen;{{ itemsTotal.toFixed(2) }}</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-500">运费</span>
            <span class="text-green-600">免运费</span>
          </div>
          <div v-if="selectedCoupon" class="flex justify-between text-sm">
            <span class="text-gray-500">优惠券抵扣</span>
            <span class="text-green-600">-&yen;{{ selectedCoupon.discountAmount?.toFixed(2) }}</span>
          </div>
          <el-divider class="!my-3" />
          <div class="flex justify-between">
            <span class="text-base font-medium text-gray-900">应付金额</span>
            <span class="text-2xl font-bold text-red-500">&yen;{{ payTotal.toFixed(2) }}</span>
          </div>
        </div>
      </div>

      <!-- Submit button -->
      <div class="flex justify-end">
        <el-button
          type="danger"
          size="large"
          :disabled="!selectedAddressId || orderItems.length === 0"
          :loading="submitting"
          @click="handleSubmitOrder"
          class="px-12"
        >
          提交订单
        </el-button>
      </div>

      <!-- Address form dialog -->
      <ShopAddressForm
        v-model="showAddressForm"
        :address="editingAddress"
        @saved="onAddressSaved"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getAddressList, getCartList, createOrder } from '~/api/shop.js'
import { getAvailableCoupons } from '~/api/coupon.js'
import { useCartStore } from '~/stores/cartStore.js'
import { getUrl } from '~/utils/image.js'

definePageMeta({ middleware: 'auth' })

const route = useRoute()
const router = useRouter()
const cartStore = useCartStore()

const addressList = ref([])
const selectedAddressId = ref(null)
const orderItems = ref([])
const remark = ref('')
const submitting = ref(false)
const showAddressForm = ref(false)
const editingAddress = ref(null)
const selectedCoupon = ref(null)
const showCouponSelector = ref(false)
const availableCoupons = ref([])

// Cart IDs from query (from cart page checkout)
const cartIdsParam = route.query.cartIds
// Direct buy params (from product detail buy now)
const directSkuId = route.query.skuId ? Number(route.query.skuId) : null
const directQuantity = route.query.quantity ? Number(route.query.quantity) : 1

const loadAddresses = async () => {
  const res = await getAddressList()
  if (res && res.code === 0) {
    addressList.value = res.data.list || []
    // Select default address or first one
    const defaultAddr = addressList.value.find(a => a.isDefault === 1)
    if (defaultAddr) {
      selectedAddressId.value = defaultAddr.ID
    } else if (addressList.value.length > 0) {
      selectedAddressId.value = addressList.value[0].ID
    }
  }
}

const loadCartItems = async () => {
  if (cartIdsParam) {
    const cartIds = cartIdsParam.split(',').map(Number)
    // Load full cart and filter by IDs
    const res = await getCartList()
    if (res && res.code === 0) {
      const allItems = res.data.list || []
      orderItems.value = allItems.filter(item => cartIds.includes(item.ID))
    }
  } else if (directSkuId) {
    // Direct buy - load from cart (the item was already added)
    const res = await getCartList()
    if (res && res.code === 0) {
      const allItems = res.data.list || []
      const matched = allItems.find(item => item.skuId === directSkuId)
      if (matched) {
        matched.quantity = directQuantity
        orderItems.value = [matched]
      }
    }
  }
}

await loadAddresses()
await loadCartItems()

const itemsTotal = computed(() => {
  return orderItems.value.reduce((sum, item) => {
    return sum + (item.skuPrice || 0) * (item.quantity || 0)
  }, 0)
})

const payTotal = computed(() => {
  let total = itemsTotal.value
  if (selectedCoupon.value && selectedCoupon.value.discountAmount) {
    total -= selectedCoupon.value.discountAmount
  }
  return total > 0 ? total : 0
})

const loadCoupons = async () => {
  const res = await getAvailableCoupons({ amount: itemsTotal.value })
  if (res && res.code === 0) {
    availableCoupons.value = res.data.list || []
  }
  showCouponSelector.value = true
}

const onCouponSelect = (coupon) => {
  selectedCoupon.value = coupon
  showCouponSelector.value = false
}

const formatSpec = (specData) => {
  try {
    const specs = JSON.parse(specData)
    return Object.entries(specs).map(([k, v]) => `${k}: ${v}`).join(', ')
  } catch (e) {
    return specData
  }
}

const onAddressSaved = async () => {
  await loadAddresses()
}

const handleSubmitOrder = async () => {
  if (!selectedAddressId.value) {
    ElMessage.warning('请选择收货地址')
    return
  }
  if (orderItems.value.length === 0) {
    ElMessage.warning('暂无商品')
    return
  }

  submitting.value = true
  try {
    const payload = {
      addressId: selectedAddressId.value,
      remark: remark.value,
      couponCode: selectedCoupon.value?.code || ''
    }

    if (directSkuId && !cartIdsParam) {
      // Direct buy
      payload.skuId = directSkuId
      payload.quantity = directQuantity
    } else {
      // Cart checkout
      payload.cartIds = orderItems.value.map(item => item.ID)
    }

    const res = await createOrder(payload)
    if (res && res.code === 0) {
      ElMessage.success('订单提交成功')
      cartStore.refreshCount()
      const order = res.data.order
      router.replace(`/shop/order/${order.orderNo}`)
    }
  } finally {
    submitting.value = false
  }
}
</script>
