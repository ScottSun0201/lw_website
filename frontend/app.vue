<template>
  <NuxtLayout :name="layout">
    <NuxtPage />
  </NuxtLayout>
</template>
<script setup>
import { ref,watch } from 'vue'
import { useRoute } from 'vue-router'
import {findClientSEO} from "~/api/base.js";
import {useSeoStore} from "~/stores/seoStore.js";
const layout = ref('default')
const route = useRoute()
const seoStore = useSeoStore()
const userStore = useUserStore()

// HttpOnly cookie 无法通过 JS 读取，通过调用 getUserInfo 接口判断登录状态
// cookie 由浏览器自动携带，后端会自动验证
await userStore.checkAuth()

const withOutLayoutPaths = ['login','register']
// 这里创建白名单用于layout的变化，例如本示例中的登录页不需要layout
watch(route, () => {
  if (withOutLayoutPaths.includes(route.name)) {
    layout.value = 'empty'
  } else {
    layout.value = 'default'
  }
}, { immediate: true })

const {data} = await useAsyncData(async () => {
  const res = await findClientSEO()
  return res.data.reclientSEO
})

useServerSeoMeta({
  title: data.value.title,
  description: data.value.description,
  keywords: data.value.keywords,
  image: data.value.logo,
  site: data.value.name,
  url: route.fullPath
})

seoStore.setSeo(data.value)


</script>
