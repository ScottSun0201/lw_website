<template>
  <div class="relative bg-white">
    <div class="mx-auto max-w-7xl">
      <div class="relative z-10 lg:w-full lg:max-w-2xl">
        <svg class="absolute inset-y-0 right-8 hidden h-full w-80 translate-x-1/2 transform fill-white lg:block" viewBox="0 0 100 100" preserveAspectRatio="none" aria-hidden="true">
          <polygon points="0,0 90,0 50,100 0,100" />
        </svg>

        <div class="relative py-32 px-6 sm:py-40 lg:py-56 lg:px-8 lg:pr-0">
          <div class="mx-auto max-w-2xl lg:mx-0 lg:max-w-xl">
            <div class="hidden sm:mb-10 sm:flex">
              <div class="relative rounded-full py-1 px-3 text-sm leading-6 text-gray-500 ring-1 ring-gray-900/10 hover:ring-gray-900/20">
                发现更多优质好物，尽在骊威商城。 <NuxtLink to="/shop" class="whitespace-nowrap font-semibold text-indigo-600"><span class="absolute inset-0" aria-hidden="true" />立即探索 <span aria-hidden="true">&rarr;</span></NuxtLink>
              </div>
            </div>
            <h1 class="text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl">品质生活，从这里开始</h1>
            <p class="mt-6 text-lg leading-8 text-gray-600">精选优质商品，为您提供一站式购物体验。品质保证，售后无忧。</p>
            <div class="mt-10 flex items-center gap-x-6">
              <NuxtLink to="/shop" class="rounded-md bg-indigo-600 px-3.5 py-1.5 text-base font-semibold leading-7 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">浏览商城</NuxtLink>
              <NuxtLink to="/about" class="text-base font-semibold leading-7 text-gray-900">了解更多 <span aria-hidden="true">→</span></NuxtLink>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="hidden lg:block bg-gray-50 lg:absolute lg:inset-y-0 lg:right-0 lg:w-1/2">
      <img class="aspect-[3/2] object-cover lg:aspect-auto lg:h-full lg:w-full" :src="bg" alt="" />
    </div>
  </div>
</template>

<script setup>
import bg from "~/assets/image/bg.jpg";
</script>