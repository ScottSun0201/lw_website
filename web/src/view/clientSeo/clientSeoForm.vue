<template>
  <div>
    <div class="gva-form-box">
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item
          label="title:"
          prop="title"
        >
          <el-input
            v-model="formData.title"
            :clearable="true"
            placeholder="请输入title"
          />
        </el-form-item>
        <el-form-item
          label="description:"
          prop="description"
        >
          <el-input
            v-model="formData.description"
            :clearable="true"
            placeholder="请输入description"
          />
        </el-form-item>
        <el-form-item
          label="keywords:"
          prop="keywords"
        >
          <el-input
            v-model="formData.keywords"
            :clearable="true"
            placeholder="请输入keywords"
          />
        </el-form-item>
        <el-form-item
          label="name:"
          prop="name"
        >
          <el-input
            v-model="formData.name"
            :clearable="true"
            placeholder="请输入name"
          />
        </el-form-item>
        <el-form-item
          label="Logo:"
          prop="logo"
        >
          <SelectImage
            v-model="formData.logo"
            file-type="image"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            @click="save"
          >保存</el-button>
          <el-button
            type="primary"
            @click="back"
          >返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createClientSEO,
  findClientSEO
} from '@/api/clientSeo'

defineOptions({
  name: 'ClientSEOForm'
})

// 自动获取字典
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from '@/components/selectImage/selectImage.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  title: '',
  description: '',
  keywords: '',
  name: '',
  logo: '',
})
// 验证规则
const rule = reactive({
  title: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  description: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  keywords: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  logo: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  const res = await findClientSEO()
  if (res.code === 0) {
    formData.value = res.data.reclientSEO
  }
}

init()
// 保存按钮
const save = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    const res = await createClientSEO(formData.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
