<template>
  <div class="space-y-4">
    <div v-for="(values, specKey) in specGroups" :key="specKey">
      <div class="mb-2 text-sm font-medium text-gray-700">{{ specKey }}</div>
      <div class="flex flex-wrap gap-2">
        <button
          v-for="val in values"
          :key="val"
          type="button"
          :class="[
            'rounded-md border px-4 py-2 text-sm transition-colors',
            selectedSpecs[specKey] === val
              ? 'border-indigo-600 bg-indigo-50 text-indigo-700 font-medium'
              : 'border-gray-300 bg-white text-gray-700 hover:border-gray-400'
          ]"
          @click="selectSpec(specKey, val)"
        >
          {{ val }}
        </button>
      </div>
    </div>

    <div v-if="matchedSku" class="mt-4 rounded-lg bg-gray-50 p-3">
      <div class="flex items-center justify-between">
        <span class="text-lg font-bold text-red-500">&yen;{{ matchedSku.price.toFixed(2) }}</span>
        <span class="text-sm text-gray-500">
          {{ matchedSku.stock > 0 ? `库存：${matchedSku.stock}` : '暂无库存' }}
        </span>
      </div>
      <div v-if="matchedSku.marketPrice > matchedSku.price" class="mt-1">
        <span class="text-xs text-gray-400 line-through">&yen;{{ matchedSku.marketPrice.toFixed(2) }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

const props = defineProps({
  skus: {
    type: Array,
    default: () => []
  },
  modelValue: {
    type: [Number, null],
    default: null
  }
})

const emit = defineEmits(['update:modelValue', 'select'])

const selectedSpecs = ref({})

// Parse specData JSON for all SKUs and group by spec key
const parsedSkus = computed(() => {
  return props.skus.map(sku => {
    let specs = {}
    try {
      if (sku.specData) {
        specs = JSON.parse(sku.specData)
      }
    } catch (e) {
      // ignore parse error
    }
    return { ...sku, specs }
  })
})

const specGroups = computed(() => {
  const groups = {}
  parsedSkus.value.forEach(sku => {
    Object.entries(sku.specs).forEach(([key, value]) => {
      if (!groups[key]) {
        groups[key] = new Set()
      }
      groups[key].add(value)
    })
  })
  // Convert sets to arrays
  const result = {}
  Object.entries(groups).forEach(([key, valueSet]) => {
    result[key] = Array.from(valueSet)
  })
  return result
})

const specKeys = computed(() => Object.keys(specGroups.value))

const matchedSku = computed(() => {
  // Only match when all spec keys are selected
  if (specKeys.value.length === 0) return null
  const allSelected = specKeys.value.every(key => selectedSpecs.value[key] !== undefined)
  if (!allSelected) return null

  return parsedSkus.value.find(sku => {
    return specKeys.value.every(key => sku.specs[key] === selectedSpecs.value[key])
  }) || null
})

const selectSpec = (key, value) => {
  if (selectedSpecs.value[key] === value) {
    // Deselect
    const newSpecs = { ...selectedSpecs.value }
    delete newSpecs[key]
    selectedSpecs.value = newSpecs
  } else {
    selectedSpecs.value = { ...selectedSpecs.value, [key]: value }
  }
}

watch(matchedSku, (sku) => {
  if (sku) {
    emit('update:modelValue', sku.ID)
    emit('select', sku)
  } else {
    emit('update:modelValue', null)
    emit('select', null)
  }
})

// If there is only one SKU, auto-select it
watch(() => props.skus, (skus) => {
  if (skus && skus.length === 1) {
    const sku = parsedSkus.value[0]
    if (sku && sku.specs) {
      selectedSpecs.value = { ...sku.specs }
    }
  }
}, { immediate: true })
</script>
