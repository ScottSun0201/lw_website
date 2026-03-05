<template>
  <ul role="list" class="space-y-6">
    <li v-for="activityItem in commits" :key="activityItem.ID" class="relative flex gap-x-4">
        <img :src="getUrl(activityItem.user.avatar)" alt="" class="relative mt-2 h-8 w-8 flex-none rounded-full bg-gray-50" />
        <div class="flex-auto rounded-md p-3 ring-1 ring-inset ring-gray-200">
          <div class="flex justify-between gap-x-4">
            <div class="py-0.5 text-xs leading-5 text-gray-500">
              <span class="font-medium text-gray-900">{{ activityItem.user.nickname }}</span> 评论了
            </div>
            <time :datetime="activityItem.dateTime" class="flex-none py-0.5 text-xs leading-5 text-gray-500">{{ formatTimeToStr(activityItem.CreatedAt)  }}</time>
          </div>
          <p class="text-sm leading-6 text-gray-500">{{ activityItem.info }}</p>
        </div>
    </li>
  </ul>

  <!-- New comment form -->
  <div class="mt-6 flex gap-x-3" v-if="userStore.token">
    <img :src="getUrl(userStore.userInfo.avatar)" alt="" class="h-8 w-8 flex-none rounded-full bg-gray-50" />
    <div class="relative flex-auto">
      <div class="overflow-hidden rounded-lg pb-12 shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-indigo-600">
        <label for="comment" class="sr-only">添加评论</label>
        <textarea v-model="info" rows="2" name="comment" id="comment" class="block w-full resize-none border-0 bg-transparent py-1.5 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" placeholder="写下你的评论..." />
      </div>

      <div class="absolute inset-x-0 bottom-0 flex justify-end py-2 pl-3 pr-2">
        <button @click="commit" class="rounded-md cursor-pointer bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50">评论</button>
      </div>
    </div>
  </div>
</template>

<script setup>

import {useUserStore} from "~/stores/userStore";
import {createCmsCommit, getCmsCommitList} from "~/api/blog.js";

import { ref } from 'vue'
import {useRoute} from "vue-router";

const route = useRoute()

const userStore = useUserStore()

const commits = ref([])
const info = ref("")

const getCommit = async () => {
  const res = await getCmsCommitList({
    cmsArticleID: Number(route.params.id)
  })
  commits.value = res.data.list
}

getCommit()


const commit = async () => {
  const res = await createCmsCommit({
    cmsArticleID: Number(route.params.id),
    info: info.value
  })
  if(res.code === 0){
    getCommit()
    info.value = ""
  }
}


</script>
