<template>
    <div>
      <HomeHero></HomeHero>
      <HomeFeature></HomeFeature>
      <HomeLogocloud></HomeLogocloud>
      <HomeFromblog></HomeFromblog>
      <div class="mx-auto max-w-7xl px-4 py-12 sm:px-6 lg:px-8">
        <ShopRecommendationSection title="热销推荐" type="hot" />
      </div>
      <div class="mx-auto max-w-7xl px-4 py-12 sm:px-6 lg:px-8">
        <ShopRecommendationSection title="新品上架" type="new" />
      </div>
    </div>
  </template>
<script setup>

useHead({
  titleTemplate: (titleChunk) => {
    return `${titleChunk} - 首页`
  },
})

</script>
