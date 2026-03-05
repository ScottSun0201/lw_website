<template>
  <div class="bg-white">
    <div class="mx-auto max-w-7xl py-12 px-6 lg:px-8">
      <p class="text-center text-base font-semibold text-gray-500">值得信赖的合作伙伴</p>
      <div class="mt-6 grid grid-cols-2 gap-8 md:grid-cols-5">
        <div v-for="i in 5" :key="i" class="col-span-1 flex justify-center items-center">
          <div class="h-12 w-28 rounded bg-gray-100 flex items-center justify-center text-gray-400 text-sm">Logo {{ i }}</div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
</script>
