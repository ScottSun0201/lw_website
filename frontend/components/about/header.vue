<template>
  <div class="relative bg-indigo-800">
    <div class="absolute inset-0">
      <img class="h-full w-full object-cover" :src="bg" alt="" />
      <div class="absolute inset-0 bg-indigo-800 mix-blend-multiply" aria-hidden="true" />
    </div>
    <div class="relative mx-auto max-w-7xl py-24 px-6 sm:py-32 lg:px-8">
      <h1 class="text-4xl font-bold tracking-tight text-white sm:text-5xl lg:text-6xl">{{ info?.title }}</h1>
      <p class="mt-6 max-w-3xl text-xl text-indigo-100">{{ info?.desc }}</p>
    </div>
  </div>
</template>
<script setup>
import bg from "~/assets/image/bg.jpg";
defineProps({
  info:{
    type:Object,
    default:{}
  },
})
</script>
