<template>
  <div class="bg-gray-50 min-h-screen">
    <div class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
      <!-- Page header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">商城</h1>
        <p class="mt-2 text-gray-500">浏览我们的商品</p>
      </div>

      <!-- Recommendations -->
      <div class="mb-8">
        <ShopRecommendationSection title="热销商品" type="hot" />
      </div>
      <div class="mb-8">
        <ShopRecommendationSection title="新品上架" type="new" />
      </div>

      <div class="flex gap-8">
        <!-- Sidebar: Category tree -->
        <aside class="hidden w-64 flex-shrink-0 lg:block">
          <div class="rounded-lg bg-white p-4 shadow-sm ring-1 ring-gray-200">
            <h3 class="mb-3 text-sm font-semibold uppercase tracking-wide text-gray-500">商品分类</h3>
            <div
              class="mb-1 cursor-pointer rounded-md px-3 py-2 text-sm transition-colors"
              :class="!currentCategoryId ? 'bg-indigo-50 text-indigo-700 font-medium' : 'text-gray-700 hover:bg-gray-50'"
              @click="selectCategory(null)"
            >
              全部分类
            </div>
            <template v-if="categories">
              <div v-for="cat in categories" :key="cat.ID">
                <div
                  class="flex cursor-pointer items-center justify-between rounded-md px-3 py-2 text-sm transition-colors"
                  :class="currentCategoryId === cat.ID ? 'bg-indigo-50 text-indigo-700 font-medium' : 'text-gray-700 hover:bg-gray-50'"
                  @click="selectCategory(cat.ID)"
                >
                  <span>{{ cat.name }}</span>
                  <svg
                    v-if="cat.children && cat.children.length"
                    class="h-4 w-4 transition-transform"
                    :class="expandedCategories.has(cat.ID) ? 'rotate-90' : ''"
                    fill="none" stroke="currentColor" viewBox="0 0 24 24"
                    @click.stop="toggleCategory(cat.ID)"
                  >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                  </svg>
                </div>
                <div v-if="cat.children && cat.children.length && expandedCategories.has(cat.ID)" class="ml-4">
                  <div
                    v-for="child in cat.children"
                    :key="child.ID"
                    class="cursor-pointer rounded-md px-3 py-1.5 text-sm transition-colors"
                    :class="currentCategoryId === child.ID ? 'bg-indigo-50 text-indigo-700 font-medium' : 'text-gray-600 hover:bg-gray-50'"
                    @click="selectCategory(child.ID)"
                  >
                    {{ child.name }}
                  </div>
                </div>
              </div>
            </template>
          </div>
        </aside>

        <!-- Main content -->
        <div class="flex-1">
          <!-- Brand filter chips -->
          <div v-if="brands && brands.length" class="mb-4 flex flex-wrap items-center gap-2">
            <span class="text-sm text-gray-500 mr-1">品牌：</span>
            <button
              :class="[
                'rounded-full px-3 py-1 text-sm transition-colors',
                !currentBrandId
                  ? 'bg-indigo-600 text-white'
                  : 'bg-white text-gray-700 ring-1 ring-gray-300 hover:bg-gray-50'
              ]"
              @click="selectBrand(null)"
            >
              全部
            </button>
            <button
              v-for="brand in brands"
              :key="brand.ID"
              :class="[
                'rounded-full px-3 py-1 text-sm transition-colors',
                currentBrandId === brand.ID
                  ? 'bg-indigo-600 text-white'
                  : 'bg-white text-gray-700 ring-1 ring-gray-300 hover:bg-gray-50'
              ]"
              @click="selectBrand(brand.ID)"
            >
              {{ brand.name }}
            </button>
          </div>

          <!-- Sort buttons and search -->
          <div class="mb-6 flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
            <div class="flex items-center gap-2">
              <span class="text-sm text-gray-500">排序：</span>
              <button
                v-for="s in sortOptions"
                :key="s.value"
                :class="[
                  'rounded-md px-3 py-1.5 text-sm font-medium transition-colors',
                  currentSort === s.value
                    ? 'bg-indigo-600 text-white'
                    : 'bg-white text-gray-700 ring-1 ring-gray-300 hover:bg-gray-50'
                ]"
                @click="selectSort(s.value)"
              >
                {{ s.label }}
              </button>
            </div>
            <div class="flex items-center gap-2">
              <el-input
                v-model="searchKeyword"
                placeholder="搜索商品..."
                clearable
                class="w-60"
                @keyup.enter="doSearch"
                @clear="doSearch"
              >
                <template #prefix>
                  <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                  </svg>
                </template>
              </el-input>
              <el-button type="primary" @click="doSearch">搜索</el-button>
            </div>
          </div>

          <!-- Mobile category selector -->
          <div class="mb-4 lg:hidden">
            <el-select v-model="mobileCategoryId" placeholder="选择分类" clearable class="w-full" @change="selectCategory($event || null)">
              <template v-if="categories">
                <template v-for="cat in categories" :key="cat.ID">
                  <el-option :label="cat.name" :value="cat.ID" />
                  <el-option
                    v-for="child in (cat.children || [])"
                    :key="child.ID"
                    :label="'  ' + child.name"
                    :value="child.ID"
                  />
                </template>
              </template>
            </el-select>
          </div>

          <!-- Product grid -->
          <div v-if="productList && productList.length" class="grid grid-cols-2 gap-4 sm:gap-6 lg:grid-cols-4">
            <ShopProductCard
              v-for="product in productList"
              :key="product.ID"
              :product="product"
            />
          </div>

          <!-- Empty state -->
          <div v-else class="flex flex-col items-center justify-center py-20 text-gray-400">
            <svg class="mb-4 h-16 w-16" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
            <p class="text-lg">暂无商品</p>
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
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getCategoryTree, getAllBrands, getClientProductList } from '~/api/shop.js'
import { getHotProducts, getNewProducts } from '~/api/recommendation.js'

const route = useRoute()
const router = useRouter()

const pageSize = 12
const currentPage = ref(Number(route.query.page) || 1)
const currentCategoryId = ref(route.query.categoryId ? Number(route.query.categoryId) : null)
const currentBrandId = ref(route.query.brandId ? Number(route.query.brandId) : null)
const currentSort = ref(route.query.sort || 'newest')
const searchKeyword = ref(route.query.keyword || '')
const mobileCategoryId = ref(currentCategoryId.value)
const expandedCategories = ref(new Set())
const total = ref(0)
const productList = ref([])

const sortOptions = [
  { label: '最新', value: 'newest' },
  { label: '价格低到高', value: 'price_asc' },
  { label: '价格高到低', value: 'price_desc' },
  { label: '销量优先', value: 'sales' }
]

// Fetch categories and brands via SSR
const { data: categories } = await useAsyncData('shopCategories', async () => {
  const res = await getCategoryTree()
  if (res && res.code === 0) {
    return res.data.list || []
  }
  return []
})

const { data: brands } = await useAsyncData('shopBrands', async () => {
  const res = await getAllBrands()
  if (res && res.code === 0) {
    return res.data.list || []
  }
  return []
})

// Fetch products
const fetchProducts = async () => {
  const params = {
    page: currentPage.value,
    pageSize
  }
  if (currentCategoryId.value) params.categoryId = currentCategoryId.value
  if (currentBrandId.value) params.brandId = currentBrandId.value
  if (searchKeyword.value) params.name = searchKeyword.value

  const res = await getClientProductList(params)
  if (res && res.code === 0) {
    productList.value = res.data.list || []
    total.value = res.data.total || 0
  }
}

// SSR initial load
const { data: initialProducts } = await useAsyncData('shopProducts', async () => {
  const params = {
    page: currentPage.value,
    pageSize
  }
  if (currentCategoryId.value) params.categoryId = currentCategoryId.value
  if (currentBrandId.value) params.brandId = currentBrandId.value
  if (searchKeyword.value) params.name = searchKeyword.value

  const res = await getClientProductList(params)
  if (res && res.code === 0) {
    return res.data
  }
  return { list: [], total: 0 }
})

if (initialProducts.value) {
  productList.value = initialProducts.value.list || []
  total.value = initialProducts.value.total || 0
}

const updateQuery = () => {
  const query = {}
  if (currentCategoryId.value) query.categoryId = currentCategoryId.value
  if (currentBrandId.value) query.brandId = currentBrandId.value
  if (searchKeyword.value) query.keyword = searchKeyword.value
  if (currentPage.value > 1) query.page = currentPage.value
  if (currentSort.value !== 'newest') query.sort = currentSort.value
  router.replace({ query })
}

const selectCategory = (id) => {
  currentCategoryId.value = id
  mobileCategoryId.value = id
  currentPage.value = 1
  updateQuery()
  fetchProducts()
}

const selectBrand = (id) => {
  currentBrandId.value = id
  currentPage.value = 1
  updateQuery()
  fetchProducts()
}

const selectSort = (sort) => {
  currentSort.value = sort
  currentPage.value = 1
  updateQuery()
  fetchProducts()
}

const doSearch = () => {
  currentPage.value = 1
  updateQuery()
  fetchProducts()
}

const handlePageChange = (page) => {
  currentPage.value = page
  updateQuery()
  fetchProducts()
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const toggleCategory = (id) => {
  const newSet = new Set(expandedCategories.value)
  if (newSet.has(id)) {
    newSet.delete(id)
  } else {
    newSet.add(id)
  }
  expandedCategories.value = newSet
}
</script>
