<template>
  <el-dialog v-model="visible" title="评价商品" width="500px" destroy-on-close>
    <el-form ref="formRef" :model="form" :rules="rules" label-width="80px">
      <el-form-item label="评分" prop="rating">
        <div class="flex gap-1">
          <span v-for="i in 5" :key="i" class="cursor-pointer text-2xl" :class="i <= form.rating ? 'text-orange-400' : 'text-gray-300'" @click="form.rating = i">★</span>
        </div>
      </el-form-item>
      <el-form-item label="评价内容" prop="content">
        <el-input v-model="form.content" type="textarea" :rows="4" placeholder="分享您的使用体验..." maxlength="500" show-word-limit />
      </el-form-item>
      <el-form-item label="匿名评价">
        <el-switch v-model="isAnon" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" :loading="submitting" @click="handleSubmit">提交评价</el-button>
    </template>
  </el-dialog>
</template>
<script setup>
import { ref, reactive } from 'vue'
import { createReview } from '~/api/review.js'

const props = defineProps({ orderNo: String })
const emit = defineEmits(['success'])
const visible = defineModel({ type: Boolean, default: false })

const formRef = ref()
const submitting = ref(false)
const isAnon = ref(false)
const form = reactive({ rating: 5, content: '', images: '' })
const rules = { rating: [{ required: true, message: '请评分', trigger: 'change' }] }

const handleSubmit = async () => {
  await formRef.value?.validate()
  submitting.value = true
  try {
    const res = await createReview({ orderNo: props.orderNo, rating: form.rating, content: form.content, images: form.images, isAnon: isAnon.value ? 1 : 0 })
    if (res && res.code === 0) { ElMessage.success('评价成功'); visible.value = false; emit('success') }
  } finally { submitting.value = false }
}
</script>
