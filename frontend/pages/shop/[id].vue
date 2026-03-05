<template>
  <div class="bg-white min-h-screen">
    <div class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
      <!-- Breadcrumb -->
      <nav class="mb-6 flex items-center gap-2 text-sm text-gray-500">
        <NuxtLink to="/shop" class="hover:text-indigo-600">商城</NuxtLink>
        <span>/</span>
        <span v-if="product?.category" class="hover:text-indigo-600">{{ product.category.name }}</span>
        <span v-if="product?.category">/</span>
        <span class="text-gray-900">{{ product?.name }}</span>
      </nav>

      <div v-if="product" class="lg:grid lg:grid-cols-2 lg:gap-12">
        <!-- Left: Image gallery -->
        <div>
          <!-- Main image -->
          <div class="aspect-square overflow-hidden rounded-xl bg-gray-100">
            <img
              :src="getUrl(currentImage)"
              :alt="product.name"
              class="h-full w-full object-cover"
            />
          </div>
          <!-- Thumbnails -->
          <div v-if="imageList.length > 1" class="mt-4 flex gap-3 overflow-x-auto pb-2">
            <div
              v-for="(img, idx) in imageList"
              :key="img || idx"
              class="h-20 w-20 flex-shrink-0 cursor-pointer overflow-hidden rounded-lg border-2 transition-colors"
              :class="currentImage === img ? 'border-indigo-500' : 'border-gray-200 hover:border-gray-300'"
              @click="currentImage = img"
            >
              <img :src="getUrl(img)" :alt="`图片 ${idx + 1}`" class="h-full w-full object-cover" />
            </div>
          </div>
        </div>

        <!-- Right: Product info -->
        <div class="mt-8 lg:mt-0">
          <h1 class="text-2xl font-bold text-gray-900 sm:text-3xl">{{ product.name }}</h1>
          <p v-if="product.subTitle" class="mt-2 text-base text-gray-500">{{ product.subTitle }}</p>

          <!-- Price -->
          <div class="mt-4 flex items-baseline gap-3">
            <span class="text-3xl font-bold text-red-500">
              {{ selectedSku ? `¥${selectedSku.price.toFixed(2)}` : product.priceRange }}
            </span>
            <span v-if="selectedSku && selectedSku.marketPrice > selectedSku.price" class="text-lg text-gray-400 line-through">
              &yen;{{ selectedSku.marketPrice.toFixed(2) }}
            </span>
          </div>

          <!-- Sales info -->
          <div class="mt-3 flex items-center gap-4 text-sm text-gray-500">
            <span>销量：{{ product.salesCount || 0 }}</span>
            <span v-if="selectedSku">库存：{{ selectedSku.stock }}</span>
          </div>

          <el-divider />

          <!-- SKU Selector -->
          <div v-if="product.skus && product.skus.length" class="mt-6">
            <ShopSkuSelector
              v-model="selectedSkuId"
              :skus="product.skus"
              @select="onSkuSelect"
            />
          </div>

          <!-- Quantity -->
          <div class="mt-6 flex items-center gap-4">
            <span class="text-sm font-medium text-gray-700">数量</span>
            <el-input-number
              v-model="quantity"
              :min="1"
              :max="selectedSku ? selectedSku.stock : 99"
              size="default"
            />
          </div>

          <!-- Action buttons -->
          <div class="mt-8 flex gap-4">
            <el-button
              type="warning"
              size="large"
              class="flex-1"
              :disabled="!selectedSkuId"
              :loading="addingToCart"
              @click="handleAddToCart"
            >
              加入购物车
            </el-button>
            <el-button
              type="danger"
              size="large"
              class="flex-1"
              :disabled="!selectedSkuId"
              @click="handleBuyNow"
            >
              立即购买
            </el-button>
          </div>

          <div v-if="!selectedSkuId && product.skus && product.skus.length > 1" class="mt-3 text-sm text-orange-500">
            请选择商品规格
          </div>

          <!-- Favorite button -->
          <div class="mt-4">
            <ShopFavoriteButton :spuId="Number(route.params.id)" />
          </div>
        </div>
      </div>

      <!-- Product detail HTML -->
      <div v-if="product" class="mt-12">
        <el-divider>
          <span class="text-lg font-semibold text-gray-700">商品详情</span>
        </el-divider>
        <div class="prose mx-auto max-w-4xl mt-6" v-html="sanitizeHtml(product.detail)"></div>
      </div>

      <!-- Reviews section -->
      <div v-if="product" class="mt-12">
        <el-divider>
          <span class="text-lg font-semibold text-gray-700">商品评价</span>
        </el-divider>
        <ShopReviewList :spuId="Number(route.params.id)" />
      </div>

      <!-- Related recommendations -->
      <div v-if="product" class="mt-12">
        <ShopRecommendationSection title="相关推荐" type="related" :spuId="Number(route.params.id)" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getClientProductDetail, addToCart } from '~/api/shop.js'
import { getProductReviews, getReviewStat } from '~/api/review.js'
import { getRelatedProducts } from '~/api/recommendation.js'
import { useCartStore } from '~/stores/cartStore.js'
import { useUserStore } from '~/stores/userStore.js'
import { getUrl } from '~/utils/image.js'
import { sanitizeHtml } from '~/utils/sanitize.js'

const route = useRoute()
const router = useRouter()
const cartStore = useCartStore()
const userStore = useUserStore()

const selectedSkuId = ref(null)
const selectedSku = ref(null)
const quantity = ref(1)
const addingToCart = ref(false)

const { data: product } = await useAsyncData(`product-${route.params.id}`, async () => {
  const res = await getClientProductDetail({ ID: route.params.id })
  if (res && res.code === 0) {
    return res.data.data
  }
  return null
})

// Build image list from mainImage + images JSON
const imageList = computed(() => {
  if (!product.value) return []
  const list = []
  if (product.value.mainImage) {
    list.push(product.value.mainImage)
  }
  if (product.value.images) {
    try {
      const parsed = JSON.parse(product.value.images)
      if (Array.isArray(parsed)) {
        parsed.forEach(img => {
          if (img && !list.includes(img)) {
            list.push(img)
          }
        })
      }
    } catch (e) {
      // ignore
    }
  }
  return list
})

const currentImage = ref(imageList.value.length > 0 ? imageList.value[0] : '')

const onSkuSelect = (sku) => {
  selectedSku.value = sku
  if (sku && sku.image) {
    currentImage.value = sku.image
  }
  quantity.value = 1
}

const handleAddToCart = async () => {
  if (!userStore.token) {
    router.push('/login?redirect=' + encodeURIComponent(route.fullPath))
    return
  }
  if (!selectedSkuId.value) {
    ElMessage.warning('请选择商品规格')
    return
  }
  addingToCart.value = true
  try {
    const res = await addToCart({
      skuId: selectedSkuId.value,
      quantity: quantity.value
    })
    if (res && res.code === 0) {
      ElMessage.success('已加入购物车')
      cartStore.increment()
    }
  } finally {
    addingToCart.value = false
  }
}

const handleBuyNow = async () => {
  if (!userStore.token) {
    router.push('/login?redirect=' + encodeURIComponent(route.fullPath))
    return
  }
  if (!selectedSkuId.value) {
    ElMessage.warning('请选择商品规格')
    return
  }
  addingToCart.value = true
  try {
    const res = await addToCart({
      skuId: selectedSkuId.value,
      quantity: quantity.value
    })
    if (res && res.code === 0) {
      cartStore.increment()
      router.push('/shop/checkout?skuId=' + selectedSkuId.value + '&quantity=' + quantity.value)
    }
  } finally {
    addingToCart.value = false
  }
}
</script>
