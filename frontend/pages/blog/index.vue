<template>
  <div class="bg-white px-6 pt-16 pb-20 lg:px-8 lg:pt-24 lg:pb-28">
    <div class="relative mx-auto max-w-lg divide-y-2 divide-gray-200 lg:max-w-7xl">
      <div>
        <h2 class="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">博客</h2>
        <div class="mt-3 sm:mt-4 lg:grid lg:grid-cols-2 lg:items-center lg:gap-5">
          <p class="text-xl text-gray-500">阅读帮助您更好地了解世界。</p>
        </div>
      </div>
      <div class="mt-6 grid gap-16 pt-10 lg:grid-cols-1 lg:gap-x-5 lg:gap-y-12">
        <div v-for="post in data" :key="post?.title" class="flex flex-col lg:flex-row gap-2 overflow-hidden">
          <div class="h-40">
            <img :src="getUrl(post?.image)" class="h-full lg:w-auto w-full object-cover rounded-xl"/>
          </div>
          <div>
          <p class="text-sm text-gray-500">
            <time :datetime="post?.datetime">{{ formatTimeToStr(post?.CreatedAt) }}</time>
          </p>
          <NuxtLink :to="{name:'blog-posts-id',params:{id:post?.ID}}" class="mt-2 block">
            <p class="text-xl font-semibold text-gray-900">{{ post?.title }}</p>
            <p class="mt-3 text-base text-gray-500">{{ post?.desc }}</p>
          </NuxtLink>
          <div class="mt-3">
            <NuxtLink :to="{name:'blog-posts-id',params:{id:post?.ID}}" class="text-base font-semibold text-indigo-600 hover:text-indigo-500">阅读全文</NuxtLink>
          </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {getCmsArticleList} from "~/api/blog.js";


const {data} = await useAsyncData("posts",async () => {
  const res = await getCmsArticleList({page:1,pageSize:10})
  if(res.code === 0){
    return res.data.list
  }
})


</script>
