<template>
  <div class="flex h-screen bg-white">
    <div class="flex flex-1 flex-col justify-center py-12 px-4 sm:px-6 lg:px-20 xl:px-24">
      <div class="mx-auto w-full max-w-sm lg:w-96">
        <div>
          <img class="h-12 w-auto" :src="getUrl(seoStore.seo.logo)" alt="" />
          <h2 class="mt-6 text-3xl font-bold tracking-tight text-gray-900">注册新账号</h2>
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
                <el-input type="password" show-password v-model="formInfo.password" ></el-input>
              </div>
            </div>
            <div class="space-y-1">
              <label for="rePassword" class="block text-sm font-medium text-gray-700">确认密码</label>
              <div class="mt-1">
                <el-input type="password" show-password v-model="formInfo.rePassword" ></el-input>
              </div>
            </div>
            <div class="mt-4">
              <el-button @click="registerFunc" class="flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">注册</el-button>
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
import {useRouter} from 'vue-router'
import { ref } from 'vue'
import {register} from "~/api/base.js";
import {useUserStore} from "~/stores/userStore.js";
import {useSeoStore} from "~/stores/seoStore.js";
import {getUrl} from "~/utils/image.js";
import bg from "~/assets/image/bg.jpg";
const userStore = useUserStore()
const seoStore = useSeoStore()

const formInfo = ref({
  username:"",
  password:"",
  rePassword:""
})
const router = useRouter()

const registerFunc = async ()=>{
  if (formInfo.value.username === "" || formInfo.value.password === "" || formInfo.value.rePassword === "") {
    ElMessage.error("请填写完整信息")
    return
  }

  if (formInfo.value.username.length < 6 || formInfo.value.username.length > 20 ) {
    ElMessage.error("用户名长度为6-20位")
    return
  }

  if (formInfo.value.password.length < 6 || formInfo.value.password.length > 20 ) {
    ElMessage.error("密码长度为6-20位")
    return
  }

  const passwordReg = /^(?=.*[a-zA-Z])(?=.*\d).{6,20}$/
  if (!passwordReg.test(formInfo.value.password)) {
    ElMessage.error("密码需包含字母和数字")
    return
  }

  if (formInfo.value.password !== formInfo.value.rePassword) {
    ElMessage.error("两次密码不一致")
    return
  }
  const res = await register(formInfo.value)
  if(res.code === 0){
    userStore.setToken(res.data.token)
    ElMessage.success("注册成功")
    router.push({name:"index"})
  }
}

</script>
