<template>
  <div class="flex h-screen bg-white g">
    <div class="flex flex-1 flex-col justify-center py-12 px-4 sm:px-6 lg:px-20 xl:px-24">
      <div class="mx-auto w-full max-w-sm lg:w-96">
        <div>
          <img class="h-12 w-auto" :src="getUrl(seoStore.seo.logo)" alt="" />
          <h2 class="mt-6 text-3xl font-bold tracking-tight text-gray-900">登录您的账号</h2>
        </div>

        <div class="mt-8">
          <div class="mt-6 flex flex-col gap-4">
              <div>
                <label for="email" class="block text-sm font-medium text-gray-700">用户名</label>
                <div class="mt-1">
                  <el-input class="w-full" v-model="formInfo.username" ></el-input>
                </div>
              </div>

              <div class="space-y-1">
                <label for="password" class="block text-sm font-medium text-gray-700">密码</label>
                <div class="mt-1">
                  <el-input type="password" class="border-none" v-model="formInfo.password" ></el-input>
                </div>
              </div>
              <div>
                <div class="flex items-center justify-between">
                  <div class="flex items-center">

                  </div>

                  <div class="text-sm  mb-1">
                    <NuxtLink class="font-medium text-indigo-600 hover:text-indigo-500" to="/register">去注册</NuxtLink>
                  </div>
                </div>
                <button @click="loginFun" class="flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">登录</button>
              </div>
          </div>
        </div>
      </div>
    </div>
    <div class="relative hidden w-0 flex-1 lg:block">
      <img class="absolute inset-0 h-full w-full object-cover" :src="bg" alt="" />
    </div>
  </div>
</template>
<script setup>
import {useRouter, useRoute} from 'vue-router'
import { ref } from 'vue'
import {login} from "~/api/base.js";
import {useUserStore} from "~/stores/userStore.js";
import bg from "~/assets/image/bg.jpg";
import {useSeoStore} from "~/stores/seoStore.js";
import {getUrl} from "~/utils/image.js";
const formInfo = ref({
  username:"",
  password:""
})
const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const seoStore = useSeoStore()

const loginFun = async () => {
  const res = await login(formInfo.value)
  if (res.code === 0) {
    userStore.setToken(res.data.token)
    ElMessage.success('登录成功')
    const redirect = route.query.redirect
    if (redirect) {
      const decoded = decodeURIComponent(redirect)
      if (decoded.startsWith('/') && !decoded.startsWith('//')) {
        router.push(decoded)
      } else {
        router.push('/')
      }
    } else {
      router.push('/')
    }
  }
}
</script>
