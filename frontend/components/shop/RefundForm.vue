<template>
  <el-dialog v-model="visible" title="申请退款" width="500px" destroy-on-close>
    <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
      <el-form-item label="退款类型" prop="type">
        <el-radio-group v-model="form.type">
          <el-radio :value="1">仅退款</el-radio>
          <el-radio :value="2">退货退款</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="退款金额" prop="amount">
        <el-input-number v-model="form.amount" :min="0.01" :max="maxAmount" :precision="2" :step="0.01" />
      </el-form-item>
      <el-form-item label="退款原因" prop="reason">
        <el-select v-model="form.reason" placeholder="请选择退款原因" class="w-full">
          <el-option label="不想要了" value="不想要了" />
          <el-option label="商品与描述不符" value="商品与描述不符" />
          <el-option label="质量问题" value="质量问题" />
          <el-option label="收到商品损坏" value="收到商品损坏" />
          <el-option label="其他" value="其他" />
        </el-select>
      </el-form-item>
      <el-form-item label="详细说明">
        <el-input v-model="form.description" type="textarea" :rows="3" placeholder="请描述具体情况" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" :loading="submitting" @click="handleSubmit">提交申请</el-button>
    </template>
  </el-dialog>
</template>
<script setup>
import { ref, reactive } from 'vue'
import { createRefund } from '~/api/refund.js'

const props = defineProps({ orderNo: String, maxAmount: { type: Number, default: 0 } })
const emit = defineEmits(['success'])
const visible = defineModel({ type: Boolean, default: false })

const formRef = ref()
const submitting = ref(false)
const form = reactive({ type: 1, amount: props.maxAmount, reason: '', description: '', images: '' })
const rules = {
  type: [{ required: true, message: '请选择退款类型', trigger: 'change' }],
  amount: [{ required: true, message: '请输入退款金额', trigger: 'change' }],
  reason: [{ required: true, message: '请选择退款原因', trigger: 'change' }]
}

const handleSubmit = async () => {
  await formRef.value?.validate()
  submitting.value = true
  try {
    const res = await createRefund({ orderNo: props.orderNo, ...form })
    if (res && res.code === 0) { ElMessage.success('申请成功'); visible.value = false; emit('success') }
  } finally { submitting.value = false }
}
</script>
