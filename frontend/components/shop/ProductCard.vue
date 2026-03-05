<template>
  <NuxtLink :to="`/shop/${product.ID}`" class="group block">
    <div class="overflow-hidden rounded-lg bg-white shadow-sm ring-1 ring-gray-200 transition-all duration-300 hover:shadow-md hover:-translate-y-1">
      <div class="relative aspect-square overflow-hidden bg-gray-100">
        <div class="absolute top-2 right-2 z-10" @click.prevent>
          <ShopFavoriteButton :spuId="product.ID" size="small" />
        </div>
        <img
          :src="getUrl(product.mainImage)"
          :alt="product.name"
          class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
        />
      </div>
      <div class="p-4">
        <h3 class="text-sm font-medium text-gray-900 line-clamp-2 min-h-[2.5rem]">{{ product.name }}</h3>
        <p v-if="product.subTitle" class="mt-1 text-xs text-gray-500 line-clamp-1">{{ product.subTitle }}</p>
        <div class="mt-2 flex items-center justify-between">
          <span class="text-base font-bold text-red-500">{{ product.priceRange }}</span>
          <span v-if="product.salesCount > 0" class="text-xs text-gray-400">已售 {{ product.salesCount }}+</span>
        </div>
      </div>
    </div>
  </NuxtLink>
</template>

<script setup>
import { getUrl } from '~/utils/image.js'

defineProps({
  product: {
    type: Object,
    required: true
  }
})
</script>
