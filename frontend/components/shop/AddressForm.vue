<template>
  <el-dialog
    :model-value="modelValue"
    :title="address ? '编辑地址' : '新增地址'"
    width="520px"
    @update:model-value="emit('update:modelValue', $event)"
    @close="resetForm"
  >
    <el-form ref="formRef" :model="form" :rules="rules" label-width="100px" label-position="top">
      <el-form-item label="收货人" prop="receiverName">
        <el-input v-model="form.receiverName" placeholder="请输入收货人姓名" />
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input v-model="form.phone" placeholder="请输入手机号" />
      </el-form-item>
      <div class="grid grid-cols-3 gap-3">
        <el-form-item label="省份" prop="province">
          <el-input v-model="form.province" placeholder="省份" />
        </el-form-item>
        <el-form-item label="城市" prop="city">
          <el-input v-model="form.city" placeholder="城市" />
        </el-form-item>
        <el-form-item label="区/县" prop="district">
          <el-input v-model="form.district" placeholder="区/县" />
        </el-form-item>
      </div>
      <el-form-item label="详细地址" prop="detailAddr">
        <el-input v-model="form.detailAddr" type="textarea" :rows="2" placeholder="请输入详细地址" />
      </el-form-item>
      <div class="flex items-center gap-4">
        <el-form-item label="标签" class="flex-1 mb-0">
          <el-radio-group v-model="form.label">
            <el-radio-button value="家">家</el-radio-button>
            <el-radio-button value="公司">公司</el-radio-button>
            <el-radio-button value="其他">其他</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="设为默认" class="mb-0">
          <el-switch v-model="isDefault" />
        </el-form-item>
      </div>
    </el-form>
    <template #footer>
      <div class="flex justify-end gap-3">
        <el-button @click="emit('update:modelValue', false)">取消</el-button>
        <el-button type="primary" :loading="submitting" @click="handleSubmit">保存</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, watch, computed } from 'vue'
import { createAddress, updateAddress } from '~/api/shop.js'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  address: {
    type: Object,
    default: null
  }
})

const emit = defineEmits(['update:modelValue', 'saved'])

const formRef = ref(null)
const submitting = ref(false)

const defaultForm = () => ({
  receiverName: '',
  phone: '',
  province: '',
  city: '',
  district: '',
  detailAddr: '',
  label: '家',
  isDefault: 0
})

const form = ref(defaultForm())
const isDefault = computed({
  get: () => form.value.isDefault === 1,
  set: (val) => { form.value.isDefault = val ? 1 : 0 }
})

const rules = {
  receiverName: [{ required: true, message: '请输入收货人姓名', trigger: 'blur' }],
  phone: [{ required: true, message: '请输入手机号', trigger: 'blur' }],
  province: [{ required: true, message: '请输入省份', trigger: 'blur' }],
  city: [{ required: true, message: '请输入城市', trigger: 'blur' }],
  district: [{ required: true, message: '请输入区/县', trigger: 'blur' }],
  detailAddr: [{ required: true, message: '请输入详细地址', trigger: 'blur' }]
}

watch(() => props.address, (addr) => {
  if (addr) {
    form.value = {
      receiverName: addr.receiverName || '',
      phone: addr.phone || '',
      province: addr.province || '',
      city: addr.city || '',
      district: addr.district || '',
      detailAddr: addr.detailAddr || '',
      label: addr.label || '家',
      isDefault: addr.isDefault || 0
    }
  } else {
    form.value = defaultForm()
  }
}, { immediate: true })

const resetForm = () => {
  form.value = defaultForm()
  formRef.value?.resetFields()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    submitting.value = true
    try {
      let res
      if (props.address && props.address.ID) {
        res = await updateAddress({ ...form.value, ID: props.address.ID })
      } else {
        res = await createAddress(form.value)
      }
      if (res && res.code === 0) {
        ElMessage.success(props.address ? '地址已更新' : '地址已创建')
        emit('update:modelValue', false)
        emit('saved')
        resetForm()
      }
    } finally {
      submitting.value = false
    }
  })
}
</script>
