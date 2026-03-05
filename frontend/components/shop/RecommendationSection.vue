<template>
  <div v-if="products.length" class="my-8">
    <h2 class="mb-4 text-xl font-bold text-gray-900">{{ title }}</h2>
    <div class="grid grid-cols-2 gap-4 sm:gap-6 lg:grid-cols-4">
      <ShopProductCard v-for="product in products" :key="product.ID" :product="product" />
    </div>
  </div>
</template>
<script setup>
defineProps({
  title: { type: String, default: '推荐商品' },
  products: { type: Array, default: () => [] }
})
</script>
