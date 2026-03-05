<template>
  <div class="bg-white">
    <div class="mx-auto max-w-7xl py-16 px-6 sm:py-24 lg:px-8">
      <div class="text-center">
        <p class="mt-1 text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl lg:text-6xl">{{title}}</p>
        <p class="mx-auto mt-5 max-w-xl text-xl text-gray-500">{{desc}}</p>
      </div>
    </div>
  </div>
</template>
<script setup>

defineProps({
  title:{
    type:String,
    default:""
  },
  desc:{
    type:String,
    default:""
  }
})

</script>
