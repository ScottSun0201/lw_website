<template>
  <div class="bg-gray-50 min-h-screen">
    <div class="mx-auto max-w-5xl px-4 py-8 sm:px-6 lg:px-8">
      <h1 class="mb-8 text-2xl font-bold text-gray-900">购物车</h1>

      <div v-if="cartItems.length > 0">
        <!-- Cart header -->
        <div class="hidden sm:grid sm:grid-cols-12 sm:gap-4 mb-4 px-4 text-sm font-medium text-gray-500">
          <div class="col-span-1">
            <el-checkbox v-model="selectAll" @change="handleSelectAll" />
          </div>
          <div class="col-span-5">商品</div>
          <div class="col-span-2 text-center">单价</div>
          <div class="col-span-2 text-center">数量</div>
          <div class="col-span-1 text-center">小计</div>
          <div class="col-span-1 text-center">操作</div>
        </div>

        <!-- Cart items -->
        <div class="space-y-3">
          <div
            v-for="item in cartItems"
            :key="item.ID"
            :class="[
              'rounded-lg bg-white p-4 shadow-sm ring-1 ring-gray-200',
              item.invalid ? 'opacity-60' : ''
            ]"
          >
            <div class="grid grid-cols-12 gap-4 items-center">
              <!-- Checkbox -->
              <div class="col-span-1">
                <el-checkbox
                  v-if="!item.invalid"
                  v-model="item._selected"
                  @change="updateSelectAll"
                />
                <el-tag v-else type="info" size="small" effect="dark">已失效</el-tag>
              </div>

              <!-- Product info -->
              <div class="col-span-12 sm:col-span-5 flex gap-3">
                <div class="h-20 w-20 flex-shrink-0 overflow-hidden rounded-lg bg-gray-100">
                  <img :src="getUrl(item.skuImage || item.spuImage)" :alt="item.spuName" class="h-full w-full object-cover" />
                </div>
                <div class="flex flex-col justify-center min-w-0">
                  <NuxtLink :to="`/shop/${item.sku?.spuId}`" class="text-sm font-medium text-gray-900 hover:text-indigo-600 line-clamp-2">
                    {{ item.spuName }}
                  </NuxtLink>
                  <span v-if="item.specData" class="mt-1 text-xs text-gray-400">{{ formatSpec(item.specData) }}</span>
                </div>
              </div>

              <!-- Price -->
              <div class="col-span-4 sm:col-span-2 text-center">
                <span class="text-sm font-medium text-red-500">&yen;{{ item.skuPrice?.toFixed(2) }}</span>
              </div>

              <!-- Quantity -->
              <div class="col-span-4 sm:col-span-2 flex justify-center">
                <el-input-number
                  v-model="item.quantity"
                  :min="1"
                  :max="item.stock || 99"
                  size="small"
                  :disabled="item.invalid"
                  @change="(val) => handleQuantityChange(item, val)"
                />
              </div>

              <!-- Subtotal -->
              <div class="col-span-2 sm:col-span-1 text-center">
                <span class="text-sm font-bold text-red-500">&yen;{{ (item.skuPrice * item.quantity).toFixed(2) }}</span>
              </div>

              <!-- Delete -->
              <div class="col-span-2 sm:col-span-1 text-center">
                <el-button type="danger" text size="small" @click="handleRemove(item)">
                  删除
                </el-button>
              </div>
            </div>
          </div>
        </div>

        <!-- Bottom bar -->
        <div class="mt-6 flex flex-col gap-4 rounded-lg bg-white p-4 shadow-sm ring-1 ring-gray-200 sm:flex-row sm:items-center sm:justify-between">
          <div class="flex items-center gap-4">
            <el-checkbox v-model="selectAll" @change="handleSelectAll">
              全选 ({{ selectedCount }}/{{ validItems.length }})
            </el-checkbox>
            <el-button type="danger" text :disabled="selectedCount === 0" @click="handleBatchRemove">
              删除所选
            </el-button>
          </div>
          <div class="flex items-center gap-6">
            <div class="text-sm text-gray-600">
              已选 <span class="font-bold text-gray-900">{{ selectedCount }}</span> 件商品，
              合计：<span class="text-xl font-bold text-red-500">&yen;{{ totalAmount.toFixed(2) }}</span>
            </div>
            <el-button type="danger" size="large" :disabled="selectedCount === 0" @click="goCheckout">
              去结算
            </el-button>
          </div>
        </div>
      </div>

      <!-- Empty cart -->
      <div v-else class="flex flex-col items-center justify-center py-20">
        <svg class="mb-6 h-24 w-24 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 100 4 2 2 0 000-4z" />
        </svg>
        <p class="mb-4 text-lg text-gray-500">购物车是空的</p>
        <NuxtLink to="/shop">
          <el-button type="primary">去逛逛</el-button>
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { getCartList, updateCartQuantity, removeFromCart } from '~/api/shop.js'
import { useCartStore } from '~/stores/cartStore.js'
import { getUrl } from '~/utils/image.js'

definePageMeta({ middleware: 'auth' })

const router = useRouter()
const cartStore = useCartStore()

const cartItems = ref([])
const selectAll = ref(false)

const validItems = computed(() => cartItems.value.filter(i => !i.invalid))
const selectedItems = computed(() => cartItems.value.filter(i => i._selected && !i.invalid))
const selectedCount = computed(() => selectedItems.value.length)

const totalAmount = computed(() => {
  return selectedItems.value.reduce((sum, item) => {
    return sum + (item.skuPrice || 0) * (item.quantity || 0)
  }, 0)
})

const updateSelectAll = () => {
  if (validItems.value.length === 0) {
    selectAll.value = false
    return
  }
  selectAll.value = validItems.value.every(i => i._selected)
}

const loadCart = async () => {
  const res = await getCartList()
  if (res && res.code === 0) {
    cartItems.value = (res.data.list || []).map(item => ({
      ...item,
      _selected: !item.invalid
    }))
    updateSelectAll()
  }
}

await loadCart()

const handleSelectAll = (val) => {
  cartItems.value.forEach(item => {
    if (!item.invalid) {
      item._selected = val
    }
  })
}

const formatSpec = (specData) => {
  try {
    const specs = JSON.parse(specData)
    return Object.entries(specs).map(([k, v]) => `${k}: ${v}`).join(', ')
  } catch (e) {
    return specData
  }
}

let quantityTimer = null
const handleQuantityChange = (item, val) => {
  clearTimeout(quantityTimer)
  quantityTimer = setTimeout(async () => {
    const res = await updateCartQuantity({ id: item.ID, quantity: val })
    if (res && res.code !== 0) {
      await loadCart()
    }
  }, 500)
}

const handleRemove = async (item) => {
  try {
    await ElMessageBox.confirm('确定要删除该商品吗？', '提示', {
      confirmButtonText: '删除',
      cancelButtonText: '取消',
      type: 'warning'
    })
    const res = await removeFromCart({ ids: [item.ID] })
    if (res && res.code === 0) {
      await loadCart()
      cartStore.refreshCount()
    }
  } catch (e) {
    // user cancelled
  }
}

const handleBatchRemove = async () => {
  const ids = selectedItems.value.map(i => i.ID)
  if (ids.length === 0) return
  try {
    await ElMessageBox.confirm(`确定要删除选中的 ${ids.length} 件商品吗？`, '提示', {
      confirmButtonText: '删除',
      cancelButtonText: '取消',
      type: 'warning'
    })
    const res = await removeFromCart({ ids })
    if (res && res.code === 0) {
      await loadCart()
      cartStore.refreshCount()
    }
  } catch (e) {
    // user cancelled
  }
}

const goCheckout = () => {
  const cartIds = selectedItems.value.map(i => i.ID)
  if (cartIds.length === 0) {
    ElMessage.warning('请至少选择一件商品')
    return
  }
  router.push('/shop/checkout?cartIds=' + cartIds.join(','))
}
</script>
