<template>
  <div class="space-y-3">
    <div
      v-for="addr in addresses"
      :key="addr.ID"
      :class="[
        'relative flex cursor-pointer items-start gap-3 rounded-lg border-2 p-4 transition-colors',
        modelValue === addr.ID
          ? 'border-indigo-500 bg-indigo-50'
          : 'border-gray-200 hover:border-gray-300'
      ]"
      @click="emit('update:modelValue', addr.ID)"
    >
      <div class="mt-0.5">
        <div
          :class="[
            'flex h-5 w-5 items-center justify-center rounded-full border-2',
            modelValue === addr.ID
              ? 'border-indigo-600 bg-indigo-600'
              : 'border-gray-300'
          ]"
        >
          <div v-if="modelValue === addr.ID" class="h-2 w-2 rounded-full bg-white"></div>
        </div>
      </div>
      <div class="flex-1">
        <div class="flex items-center gap-2">
          <span class="font-medium text-gray-900">{{ addr.receiverName }}</span>
          <span class="text-sm text-gray-500">{{ addr.phone }}</span>
          <el-tag v-if="addr.isDefault === 1" type="danger" size="small" effect="plain">默认</el-tag>
          <el-tag v-if="addr.label" size="small" effect="plain">{{ addr.label }}</el-tag>
        </div>
        <div class="mt-1 text-sm text-gray-600">
          {{ addr.province }} {{ addr.city }} {{ addr.district }} {{ addr.detailAddr }}
        </div>
      </div>
    </div>

    <div v-if="!addresses || addresses.length === 0" class="rounded-lg border-2 border-dashed border-gray-300 p-8 text-center text-gray-400">
      暂无收货地址，请添加一个
    </div>
  </div>
</template>

<script setup>
defineProps({
  modelValue: {
    type: [Number, null],
    default: null
  },
  addresses: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['update:modelValue'])
</script>
