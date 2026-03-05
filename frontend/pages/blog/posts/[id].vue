<template>
  <div class="relative bg-white min-h-full pb-14 px-8 sm:px-0">
    <div class="mx-auto max-w-7xl">
      <BlogPostsHeader :title="data?.title" :desc="data?.desc"></BlogPostsHeader>
      <div class="p-4 mb-8 rounded">
        <div v-html="data?.content"></div>
      </div>
      <BlogPostsCommit></BlogPostsCommit>
    </div>
  </div>

</template>

<script setup>
import { useRoute } from 'vue-router'
import {findCmsArticle} from "~/api/blog.js";

const route = useRoute()

const {data} = await useAsyncData("posts",async () => {
  const res = await findCmsArticle({ID:route.params.id})
  if(res.code === 0){
    return res.data.recmsArticle
  }
})

</script>
