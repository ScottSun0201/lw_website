<template>
  <div class="bg-white">
    <div class="mx-auto max-w-7xl py-16 px-6 sm:py-24 lg:px-8">
  <form>
    <div class="space-y-12 sm:space-y-16">
      <div>
        <h2 class="text-base font-semibold leading-7 text-gray-900">个人资料</h2>
        <p class="mt-1 max-w-2xl text-sm leading-6 text-gray-600">此信息将公开显示，请注意填写内容。</p>

        <div class="mt-10 space-y-8 border-b border-gray-900/10 pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0">
          <div class="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
            <label for="username" class="block text-sm font-medium leading-6 text-gray-900 sm:pt-1.5">昵称</label>
            <div class="mt-2 sm:col-span-2 sm:mt-0">
              <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md">
                <input v-model="userStore.userInfo.nickname" type="text" name="username" id="username" class="block flex-1 border-0 bg-transparent py-1.5 pl-1 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6"  />
              </div>
            </div>
          </div>

          <div class="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
            <label for="about" class="block text-sm font-medium leading-6 text-gray-900 sm:pt-1.5">个人简介</label>
            <div class="mt-2 sm:col-span-2 sm:mt-0">
              <textarea v-model="userStore.userInfo.about" id="about" name="about" rows="3" class="block w-full max-w-2xl rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
              <p class="mt-3 text-sm leading-6 text-gray-600">简单介绍一下自己。</p>
            </div>
          </div>

          <div class="sm:grid sm:grid-cols-3 sm:items-center sm:gap-4 sm:py-6">
            <label for="photo" class="block text-sm font-medium leading-6 text-gray-900">头像</label>
            <div class="mt-2 sm:col-span-2 sm:mt-0">
              <div class="flex items-center gap-x-3">
                <img v-if="userStore.userInfo.avatar" :src="getUrl(userStore.userInfo.avatar)" class="h-12 w-12 text-gray-300 rounded-full object-cover" alt="avatar"/>
                <UserCircleIcon v-else class="h-12 w-12 text-gray-300" aria-hidden="true" />
                <button @click="changeAvatar" type="button" class="rounded-md bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50">更换</button>
                <input type="file" class="hidden" ref="avatar" @change="uploadAvatar" />
              </div>
            </div>
          </div>

        </div>
      </div>

      <div>
        <h2 class="text-base font-semibold leading-7 text-gray-900">个人信息</h2>
        <p class="mt-1 max-w-2xl text-sm leading-6 text-gray-600">请填写您的真实信息。</p>

        <div class="mt-10 space-y-8 border-b border-gray-900/10 pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0">
          <div class="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
            <label for="first-name" class="block text-sm font-medium leading-6 text-gray-900 sm:pt-1.5">名</label>
            <div class="mt-2 sm:col-span-2 sm:mt-0">
              <input v-model="userStore.userInfo.firstName" type="text" name="first-name" id="first-name" autocomplete="given-name" class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6" />
            </div>
          </div>

          <div class="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
            <label for="last-name" class="block text-sm font-medium leading-6 text-gray-900 sm:pt-1.5">姓</label>
            <div class="mt-2 sm:col-span-2 sm:mt-0">
              <input  v-model="userStore.userInfo.lastName" type="text" name="last-name" id="last-name" autocomplete="family-name" class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6" />
            </div>
          </div>

          <div class="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
            <label for="email" class="block text-sm font-medium leading-6 text-gray-900 sm:pt-1.5">邮箱</label>
            <div class="mt-2 sm:col-span-2 sm:mt-0">
              <input  v-model="userStore.userInfo.email" id="email" name="email" type="email" autocomplete="email" class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-md sm:text-sm sm:leading-6" />
            </div>
          </div>

          <div class="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
            <label for="age" class="block text-sm font-medium leading-6 text-gray-900 sm:pt-1.5">年龄</label>
            <div class="mt-2 sm:col-span-2 sm:mt-0">
              <input  v-model.number="userStore.userInfo.age" id="age" name="age" type="text" autocomplete="age" class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-md sm:text-sm sm:leading-6" />
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-6 flex items-center justify-end gap-x-6">
      <button @click="logout" type="button" class="inline-flex justify-center rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">退出登录</button>
      <button @click="save" type="button" class="inline-flex justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">保存</button>
    </div>
  </form>
    </div>
  </div>
</template>

<script setup>
import { PhotoIcon, UserCircleIcon } from '@heroicons/vue/24/solid'
import {useUserStore} from "~/stores/userStore.js";

definePageMeta({ middleware: 'auth' })
import {setClientUser, upload} from "~/api/base.js";
import {getUrl} from "~/utils/image.js";

const userStore = useUserStore()
const avatar = ref(null)
const changeAvatar = () => {
  const input = unref(avatar)
  input.click()
}

const uploadAvatar = async (e) => {
  const file = e.target.files[0]
  const formData = new FormData()
  formData.append('file', file)
  const res = await upload(formData,{noSave:"1"})
  if (res.code === 0) {
   const setRes = await setClientUser({avatar:res.data.file.url})
    if (setRes.code === 0) {
      await userStore.getInfo()
    }
  }
}

const save = async () => {
    if (userStore.userInfo.nickname === "" || userStore.userInfo.about === "" || userStore.userInfo.firstName === "" || userStore.userInfo.lastName === "" || userStore.userInfo.email === "") {
      ElMessage.error("请填写完整信息")
      return
    }
  const res = await setClientUser(userStore.userInfo)
  if (res.code === 0) {
    await userStore.getInfo()
    ElMessage.success("更新成功")
  }
}

const logout = async () => {
  ElMessageBox.confirm('确定要退出登录吗？', '退出登录', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    userStore.logout()
  }).catch(() => {
  });
}

</script>
