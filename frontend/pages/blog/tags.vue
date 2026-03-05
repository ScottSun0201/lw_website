<template>
  <div class="bg-white">
    <main class="mx-auto w-full max-w-7xl px-6 pb-16 pt-10 sm:pb-24 lg:px-8">
      <div class="mx-auto mt-20 max-w-2xl text-center sm:mt-24">
        <h1 class="mt-4 text-3xl font-bold tracking-tight text-gray-900 sm:text-5xl">标签</h1>
        <p class="mt-4 text-base leading-7 text-gray-600 sm:mt-6 sm:text-lg sm:leading-8">选择标签查看相关文章。</p>
      </div>
      <div class="mx-auto mt-16 flow-root max-w-lg sm:mt-20">
        <ul role="list" class="-mt-6 divide-y divide-gray-900/5 border-b border-gray-900/5">
          <li v-for="(link, linkIdx) in data?.list" :key="linkIdx" class="relative flex gap-x-6 py-6">
            <div class="flex-auto">
              <h3 class="text font-semibold leading-6 text-gray-900">
                <NuxtLink :to="{name:'blog-tag-id',params:{id:link.ID}}">
                  <span class="absolute inset-0" aria-hidden="true" />
                  {{ link.name }}
                </NuxtLink>
              </h3>
              <p class="mt-2 text leading-6 text-gray-600">{{ link.desc }}</p>
            </div>
            <div class="flex-none self-center">
              <ChevronRightIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
            </div>
          </li>
        </ul>
        <div class="mt-10 flex justify-center">
          <a href="/" class="text font-semibold leading-6 text-indigo-600">
            <span aria-hidden="true">&larr;</span>
            返回首页
          </a>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import {ChevronRightIcon} from '@heroicons/vue/20/solid'

import {getCmsTagList} from '~/api/tag.js'

const {data} = useAsyncData(async () => {
  const res = await getCmsTagList()
  if(res.code === 0){
    return res.data
  }
})


</script>