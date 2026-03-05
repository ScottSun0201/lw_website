<template>
  <div class="flex flex-col min-h-screen justify-between">
    <LayoutsHeader />
    <div class="flex-1"> <!-- pt-12 是为了防止头部遮挡内容 -->
      <slot />
    </div>
    <LayoutsFooter />
  </div>
</template>
<script setup>
import logo from '~/assets/image/logo.png'
import { useUserStore } from "~/stores/userStore.js";
import { logout } from "~/api/base.js";
const userStore = useUserStore()

const logoutFunc = async () => {
  const res = await logout()
  if (res.code === 0) {
    userStore.setToken("")
  }
}
</script>
