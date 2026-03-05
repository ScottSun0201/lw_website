<template>
  <div class="bg-white py-24 sm:py-32">
    <div class="mx-auto max-w-7xl px-6 lg:px-8">
      <div class="mx-auto max-w-2xl lg:mx-0">
        <h2 class="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">我们的优势</h2>
        <p class="mt-6 text-lg leading-8 text-gray-600">为您提供优质的产品和贴心的服务，让购物更简单、更放心。</p>
      </div>
      <div class="mx-auto mt-16 max-w-2xl sm:mt-20 lg:mt-24 lg:max-w-none">
        <dl class="grid max-w-xl grid-cols-1 gap-x-8 gap-y-16 lg:max-w-none lg:grid-cols-3">
          <div v-for="feature in features" :key="feature.name" class="flex flex-col">
            <dt class="text-base font-semibold leading-7 text-gray-900">
              <div class="mb-6 flex h-10 w-10 items-center justify-center rounded-lg bg-indigo-600">
                <component :is="feature.icon" class="h-6 w-6 text-white" aria-hidden="true" />
              </div>
              {{ feature.name }}
            </dt>
            <dd class="mt-1 flex flex-auto flex-col text-base leading-7 text-gray-600">
              <p class="flex-auto">{{ feature.description }}</p>
              <p class="mt-6">
                <a :href="feature.href" class="text-base font-semibold leading-7 text-indigo-600">了解更多 <span aria-hidden="true">→</span></a>
              </p>
            </dd>
          </div>
        </dl>
      </div>
    </div>
  </div>
</template>

<script setup>
import { InboxIcon, TrashIcon, UsersIcon } from '@heroicons/vue/24/outline'

const features = [
  {
    name: '品质保证',
    description: '严格筛选每一件商品，确保品质可靠。从源头把控质量，让您买得放心、用得安心。',
    href: '/shop',
    icon: InboxIcon,
  },
  {
    name: '贴心服务',
    description: '专业的客服团队随时为您解答疑问。售前咨询、售后保障，全程为您提供优质服务体验。',
    href: '/shop',
    icon: UsersIcon,
  },
  {
    name: '快速配送',
    description: '高效的物流配送体系，让您的订单快速送达。实时追踪物流信息，随时掌握包裹动态。',
    href: '/shop',
    icon: TrashIcon,
  },
]
</script>
