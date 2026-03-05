<template>
  <header class="bg-white z-50 sticky top-0">
    <nav class="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8" aria-label="导航">
      <div class="flex lg:flex-1">
        <a href="/" class="-m-1.5 p-1.5 flex items-center gap-4">
          <img class="h-8 w-auto rounded-full" :src="getUrl(seoStore.seo.logo)" alt="" />
          <span class="text-lg font-bold">{{seoStore.seo.name}}</span>
        </a>
      </div>
      <div class="flex lg:hidden">
        <button type="button" class="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700" @click="mobileMenuOpen = true">
          <span class="sr-only">打开菜单</span>
          <Bars3Icon class="h-6 w-6" aria-hidden="true" />
        </button>
      </div>
      <PopoverGroup class="hidden lg:flex lg:gap-x-12 cursor-pointer">
        <template v-for="nav in data" >
          <Popover v-if="nav.children?.length" class="relative">
            <PopoverButton class="flex items-center gap-x-1 font-semibold leading-6 text-gray-900">
              {{ nav.name }}
              <ChevronDownIcon class="h-5 w-5 flex-none text-gray-400" aria-hidden="true" />
            </PopoverButton>

            <transition enter-active-class="transition ease-out duration-200" enter-from-class="opacity-0 translate-y-1" enter-to-class="opacity-100 translate-y-0" leave-active-class="transition ease-in duration-150" leave-from-class="opacity-100 translate-y-0" leave-to-class="opacity-0 translate-y-1">
              <PopoverPanel class="absolute -left-8 top-full z-10 mt-3 w-screen max-w-md overflow-hidden rounded-3xl bg-white shadow-lg ring-1 ring-gray-900/5">
                <div class="p-4">
                  <div v-for="item in nav.children" :key="item.name" class="group relative flex items-center gap-x-6 rounded-lg p-4 leading-6 hover:bg-gray-50">
                    <div class="flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white">
                      <component :is="icons[item.icon]" class="h-6 w-6 text-gray-600 group-hover:text-indigo-600" aria-hidden="true" />
                    </div>
                    <div class="flex-auto">
                      <a :href="item.router" class="block font-semibold text-gray-900">
                        {{ item.name }}
                        <span class="absolute inset-0" />
                      </a>
                      <p class="mt-1 text-gray-600">{{ item.desc }}</p>
                    </div>
                  </div>
                </div>
              </PopoverPanel>
            </transition>
          </Popover>
          <NuxtLink v-else :to="nav.router"  class="font-semibold leading-6 text-gray-900">{{ nav.name }}</NuxtLink>
        </template>
        <NuxtLink to="/shop" class="font-semibold leading-6 text-gray-900">商城</NuxtLink>
      </PopoverGroup>
      <div class="hidden lg:flex lg:flex-1 lg:justify-end items-center gap-4">
        <!-- Notification bell -->
        <ShopNotificationBell />
        <!-- Cart icon -->
        <NuxtLink to="/shop/cart" class="relative text-gray-600 hover:text-indigo-600">
          <ShoppingCartIcon class="h-6 w-6" />
          <span
            v-if="cartStore.count > 0"
            class="absolute -top-2 -right-2 flex h-5 w-5 items-center justify-center rounded-full bg-red-500 text-xs font-bold text-white"
          >
            {{ cartStore.count > 99 ? '99+' : cartStore.count }}
          </span>
        </NuxtLink>
        <NuxtLink v-if="!userStore.token" :to="{name:'login'}" class="font-semibold leading-6 text-gray-900">登录</NuxtLink>
        <NuxtLink :to="{name:'setting'}" v-else>
          <div  class="flex items-center gap-2 cursor-pointer">
            <img class="w-12 h-12 object-cover rounded-full" :src="userStore.userInfo.avatar ? getUrl(userStore.userInfo.avatar) : ''" alt="avatar" @error="(e) => e.target.style.display='none'" />
            <div class="flex flex-col justify-around text-center">
              <div class="text text-gray-700 font-bold">{{userStore.userInfo.nickname}}</div>
              <div class="text-sm text-gray-500">{{userStore.userInfo.firstName}} {{userStore.userInfo.lastName}}</div>
            </div>
          </div>
        </NuxtLink>
      </div>
    </nav>
    <Dialog as="div" class="lg:hidden" @close="mobileMenuOpen = false" :open="mobileMenuOpen">
      <div class="fixed inset-0 z-10" />
      <DialogPanel class="box-border fixed inset-y-0 right-0 z-10 w-full overflow-y-auto bg-white px-6 py-6 sm:max-w-sm sm:ring-1 sm:ring-gray-900/10">
        <div class="flex items-center justify-between">
          <a href="/" class="-m-1.5 p-1.5">
            <span class="sr-only">{{seoStore.seo.name}}</span>
            <img class="h-8 w-auto rounded-full" :src="getUrl(seoStore.seo.logo)" alt="" />
          </a>
          <button type="button" class="-m-2.5 rounded-md p-2.5 text-gray-700" @click="mobileMenuOpen = false">
            <span class="sr-only">关闭菜单</span>
            <XMarkIcon class="h-6 w-6" aria-hidden="true" />
          </button>
        </div>
        <div class="mt-6 flow-root">
          <div class="-my-6 divide-y divide-gray-500/10">
            <div class="space-y-2 py-6">
              <NuxtLink v-for="nav in getPhoneNavs(data)" :to="nav.router" class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50" @click="mobileMenuOpen = false">{{ nav.name }}</NuxtLink>
              <NuxtLink to="/shop" class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50" @click="mobileMenuOpen = false">商城</NuxtLink>
              <div class="-mx-3 flex items-center gap-2 rounded-lg px-3 py-2">
                <ShopNotificationBell />
              </div>
              <NuxtLink to="/shop/cart" class="-mx-3 flex items-center gap-2 rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50" @click="mobileMenuOpen = false">
                购物车
                <span v-if="cartStore.count > 0" class="inline-flex h-5 w-5 items-center justify-center rounded-full bg-red-500 text-xs font-bold text-white">
                  {{ cartStore.count > 99 ? '99+' : cartStore.count }}
                </span>
              </NuxtLink>
            </div>
            <div class="py-6">
              <NuxtLink v-if="!userStore.token" :to="{name:'login'}" class="-mx-3 block rounded-lg px-3 py-2.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50" @click="mobileMenuOpen = false">登录</NuxtLink>
              <NuxtLink v-else :to="{name:'setting'}" class="-mx-3 block rounded-lg px-3 py-2.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50" @click="mobileMenuOpen = false">个人中心</NuxtLink>
            </div>
          </div>
        </div>
      </DialogPanel>
    </Dialog>


  </header>
</template>
<script setup>
import {
  Dialog,
  DialogPanel,
  Popover,
  PopoverButton,
  PopoverGroup,
  PopoverPanel,
} from '@headlessui/vue'

import { Menu, MenuButton, MenuItems, MenuItem } from '@headlessui/vue'

import {
  Bars3Icon,
  ChartPieIcon,
  CursorArrowRaysIcon,
  XMarkIcon,
  ShoppingCartIcon,
} from '@heroicons/vue/24/outline'
import { ChevronDownIcon } from '@heroicons/vue/20/solid'

import {useUserStore} from "~/stores/userStore.js";
import {useSeoStore} from "~/stores/seoStore.js";
import {useCartStore} from "~/stores/cartStore.js";
import {useNotificationStore} from "~/stores/notificationStore.js";
import {getCmsPlateList} from "~/api/base.js";

const mobileMenuOpen = ref(false)

const userStore = useUserStore()
const seoStore = useSeoStore()
const cartStore = useCartStore()
const notificationStore = useNotificationStore()

const icons = {ChartPieIcon}


const {data} = useAsyncData("header",async () => {
  const res = await getCmsPlateList()
  return res.data.list
})

// 已登录时自动刷新购物车数量和通知数量
if (userStore.token) {
  cartStore.refreshCount()
  notificationStore.refreshCount()
}

const getPhoneNavs = (navs) =>{
  const newNavs = []
  navs.forEach(item=>{
    if(item.children){
      newNavs.push(...item.children)
    }else{
      newNavs.push(item)
    }
  })
  return newNavs
}


</script>
