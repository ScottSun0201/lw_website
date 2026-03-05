<template>
  <div class="bg-gray-50 min-h-screen">
    <div class="mx-auto max-w-3xl px-4 py-8 sm:px-6 lg:px-8">
      <div class="mb-8 flex items-center justify-between">
        <h1 class="text-2xl font-bold text-gray-900">我的地址</h1>
        <el-button type="primary" @click="openAddForm">
          + 新增地址
        </el-button>
      </div>

      <!-- Address list -->
      <div v-if="addresses.length > 0" class="space-y-4">
        <div
          v-for="addr in addresses"
          :key="addr.ID"
          class="rounded-lg bg-white p-5 shadow-sm ring-1 ring-gray-200"
        >
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center gap-3">
                <span class="text-base font-medium text-gray-900">{{ addr.receiverName }}</span>
                <span class="text-sm text-gray-500">{{ addr.phone }}</span>
                <el-tag v-if="addr.isDefault === 1" type="danger" size="small" effect="plain">默认</el-tag>
                <el-tag v-if="addr.label" size="small" effect="plain">{{ addr.label }}</el-tag>
              </div>
              <div class="mt-2 text-sm text-gray-600">
                {{ addr.province }} {{ addr.city }} {{ addr.district }} {{ addr.detailAddr }}
              </div>
            </div>
            <div class="ml-4 flex flex-shrink-0 items-center gap-2">
              <el-button
                v-if="addr.isDefault !== 1"
                type="primary"
                text
                size="small"
                @click="handleSetDefault(addr)"
              >
                设为默认
              </el-button>
              <el-button
                type="primary"
                text
                size="small"
                @click="openEditForm(addr)"
              >
                编辑
              </el-button>
              <el-button
                type="danger"
                text
                size="small"
                @click="handleDelete(addr)"
              >
                删除
              </el-button>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty state -->
      <div v-else class="flex flex-col items-center justify-center py-20 text-gray-400">
        <svg class="mb-4 h-16 w-16" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
        </svg>
        <p class="text-lg">暂无收货地址</p>
        <p class="mt-1 text-sm">添加一个收货地址开始购物吧</p>
      </div>

      <!-- Address form dialog -->
      <ShopAddressForm
        v-model="showForm"
        :address="editingAddress"
        @saved="onSaved"
      />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { getAddressList, deleteAddress, setDefaultAddress } from '~/api/shop.js'

definePageMeta({ middleware: 'auth' })

const addresses = ref([])
const showForm = ref(false)
const editingAddress = ref(null)

const loadAddresses = async () => {
  const res = await getAddressList()
  if (res && res.code === 0) {
    addresses.value = res.data.list || []
  }
}

await loadAddresses()

const openAddForm = () => {
  editingAddress.value = null
  showForm.value = true
}

const openEditForm = (addr) => {
  editingAddress.value = { ...addr }
  showForm.value = true
}

const onSaved = async () => {
  await loadAddresses()
}

const handleSetDefault = async (addr) => {
  const res = await setDefaultAddress({ ID: addr.ID })
  if (res && res.code === 0) {
    ElMessage.success('默认地址已更新')
    await loadAddresses()
  }
}

const handleDelete = async (addr) => {
  try {
    await ElMessageBox.confirm('确定要删除该地址吗？', '删除地址', {
      confirmButtonText: '删除',
      cancelButtonText: '取消',
      type: 'warning'
    })
    const res = await deleteAddress({ ID: addr.ID })
    if (res && res.code === 0) {
      ElMessage.success('地址已删除')
      await loadAddresses()
    }
  } catch (e) {
    // user cancelled
  }
}
</script>
