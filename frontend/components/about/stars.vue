<template>
  <div class="relative bg-gray-900">
    <div class="absolute bottom-0 h-80 w-full xl:inset-0 xl:h-full">
      <div class="h-full w-full xl:grid xl:grid-cols-2">
        <div class="h-full xl:relative xl:col-start-2">
          <img class="h-full w-full object-cover opacity-25 xl:absolute xl:inset-0" :src="bg" alt="" />
          <div aria-hidden="true" class="absolute inset-x-0 top-0 h-32 bg-gradient-to-b from-gray-900 xl:inset-y-0 xl:left-0 xl:h-full xl:w-32 xl:bg-gradient-to-r" />
        </div>
      </div>
    </div>
    <div class="mx-auto max-w-4xl px-6 lg:max-w-7xl lg:px-8 xl:grid xl:grid-flow-col-dense xl:grid-cols-2 xl:gap-x-8">
      <div class="relative pt-12 pb-64 sm:pt-24 sm:pb-64 xl:col-start-1 xl:pb-24">
        <h2 class="text-base font-semibold text-indigo-300">数据亮点</h2>
        <p class="mt-3 text-3xl font-bold tracking-tight text-white">用数据说话，见证我们的成长</p>
        <p class="mt-5 text-lg text-gray-300">我们始终致力于为客户提供最优质的产品和服务。以下数据展示了我们的努力与成果，感谢每一位客户的信任与支持。</p>
        <div class="mt-12 grid grid-cols-1 gap-y-12 gap-x-6 sm:grid-cols-2">
          <p v-for="item in metrics" :key="item.id">
            <span class="block text-2xl font-bold text-white">{{ item.stat }}</span>
            <span class="mt-1 block text-base text-gray-300"
            ><span class="font-medium text-white">{{ item.emphasis }}</span> {{ item.rest }}</span
            >
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import bg from "~/assets/image/bg.jpg";
const metrics = [
  { id: 1, stat: '8K+', emphasis: '合作企业', rest: '信赖我们的产品与服务。' },
  { id: 2, stat: '25K+', emphasis: '覆盖地区', rest: '服务遍及全国各地。' },
  { id: 3, stat: '98%', emphasis: '客户满意度', rest: '持续提升服务品质。' },
  { id: 4, stat: '12M+', emphasis: '累计服务', rest: '为客户解决各类需求。' },
]
</script>
