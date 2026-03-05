<template>
  <div class="bg-white">
    <div class="mx-auto max-w-7xl py-12 px-6 lg:px-8 lg:py-24">
      <div class="space-y-12 lg:grid lg:grid-cols-3 lg:gap-8 lg:space-y-0">
        <div class="space-y-5 sm:space-y-4">
          <h2 class="text-3xl font-bold tracking-tight sm:text-4xl">{{ data?.info?.teamTitle}}</h2>
          <p class="text-xl text-gray-500">{{ data?.info?.teamDesc }} </p>
        </div>
        <div class="lg:col-span-2">
          <ul role="list" class="space-y-12 sm:grid sm:grid-cols-2 sm:gap-x-6 sm:gap-y-12 sm:space-y-0 lg:gap-x-8">
            <li v-for="person in data?.list" :key="person?.ID">
              <div class="space-y-4">
                <div class="aspect-w-3 aspect-h-2">
                  <img class="w-full rounded-lg object-cover shadow-lg" :src="getUrl(person?.avatar)" alt="" />
                </div>
                <div class="space-y-1 text-lg font-medium leading-6">
                  <h3>{{ person?.name }}</h3>
                  <p class="text-indigo-600">{{ person?.title }}</p>
                </div>
                <div class="text-lg">
                  <p class="text-gray-500">{{ person?.desc }}</p>
                </div>
              </div>
            </li>
          </ul>
        </div>
      </div>
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
