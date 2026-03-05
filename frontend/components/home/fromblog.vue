<template>
  <div class="relative bg-gray-50 px-6 pt-16 pb-20 lg:px-8 lg:pt-24 lg:pb-28">
    <div class="absolute inset-0">
      <div class="h-1/3 bg-white sm:h-2/3" />
    </div>
    <div class="relative mx-auto max-w-7xl">
      <div class="text-center">
        <h2 class="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">最新文章</h2>
        <p class="mx-auto mt-3 max-w-2xl text-xl text-gray-500 sm:mt-4">阅读文章，获取更多灵感与资讯。</p>
      </div>
      <div class="mx-auto mt-12 grid max-w-lg gap-5 lg:max-w-none lg:grid-cols-3">
        <div v-for="post in data.list" :key="post.title" class="flex flex-col overflow-hidden rounded-lg shadow-lg">
          <div class="flex-shrink-0">
            <img class="h-48 w-full object-cover" :src="getUrl(post.image)" alt="" />
          </div>
          <div class="flex flex-1 flex-col justify-between bg-white p-6">
            <div class="flex-1">
              <p class="text-sm font-medium text-indigo-600">
                <span class="flex gap-2">
                <a  class="hover:underline">{{ getDictLabel(data.flagOptions,post.flag)  }}</a>
                  <NuxtLink v-for="t in post.tag" :to="{name:'blog-tag-id',params:{id:t.ID}}" class="hover:underline" :key="t.ID">{{ t.name }}</NuxtLink>
                </span>
              </p>
              <NuxtLink :to="{name:'blog-posts-id',params:{id:post.ID}}" class="mt-2 block">
                <p class="text-xl font-semibold text-gray-900">{{ post.title }}</p>
                <p class="mt-3 text-base text-gray-500">{{ post.desc }}</p>
              </NuxtLink>
            </div>
            <div class="mt-6 flex items-center">
              <div class="flex-shrink-0">
                <NuxtLink :to="{name:'blog-posts-id',params:{id:post.ID}}">
                  <img class="h-10 w-10 rounded-full" :src="getUrl(post?.authorInfo?.headerImg) " alt="" />
                </NuxtLink>
              </div>
              <div class="ml-3">
                <p class="text-sm font-medium text-gray-900">
                  <NuxtLink class="hover:underline" :to="{name:'blog-posts-id',params:{id:post.ID}}">
                    {{ post?.authorInfo?.nickName }}
                  </NuxtLink>
                </p>
                <div class="flex space-x-1 text-sm text-gray-500">
                  <span >{{ formatTimeToStr(post.CreatedAt) }}</span>
                  <span aria-hidden="true">&middot;</span>
                  <span>{{ post.readingTime }} 分钟阅读</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {getCmsArticleList} from "~/api/blog.js";
import {findSysDictionary} from "~/api/dict.js";
import {formatTimeToStr} from "~/utils/date.js";
import {getDictLabel} from "~/utils/dict.js";


const {data} = await useAsyncData("home",async () => {
  const [listRes, dictRes] = await Promise.all([
    getCmsArticleList({page:1, pageSize:3}),
    findSysDictionary({type: 'flag'})
  ]);
  return {
    list: listRes.data.list,
    flagOptions: dictRes.data.resysDictionary.sysDictionaryDetails
  };
})

</script>
