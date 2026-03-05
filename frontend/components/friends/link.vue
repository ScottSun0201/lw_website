<template>
  <div class="bg-white py-24 sm:py-32">
    <div class="mx-auto grid max-w-7xl gap-x-8 gap-y-20 px-6 lg:px-8 xl:grid-cols-3">
      <div class="max-w-2xl">
        <h2 class="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">{{ data?.info?.friendTitle }}</h2>
        <p class="mt-6 text-lg leading-8 text-gray-600">{{ data?.info?.friendDesc }}</p>
      </div>
      <ul role="list" class="grid gap-x-8 gap-y-12 sm:grid-cols-2 sm:gap-y-16 xl:col-span-2">
        <li v-for="person in data?.list" :key="person?.name">
          <div class="flex items-center gap-x-6">
            <img class="h-16 w-16 rounded-full" :src="getUrl(person?.image)" alt="" />
            <div>
              <h3 class="text-base font-semibold leading-7 tracking-tight text-gray-900">{{ person?.title }}</h3>
              <a :href="person?.link" :target="person?.newWindow?'_blank':'_self'" class="text-sm font-semibold leading-6 text-indigo-600">查看详情</a>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>

defineProps({
  data:{
    type:Object,
    default:{}
  },
})

</script>
