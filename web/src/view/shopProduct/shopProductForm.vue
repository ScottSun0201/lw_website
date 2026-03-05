<template>
  <div>
    <div class="gva-form-box" style="padding: 20px;">
      <el-tabs v-model="activeTab" type="border-card">
        <!-- Tab1: 基本信息 -->
        <el-tab-pane label="基本信息" name="basic">
          <el-form
            ref="basicFormRef"
            :model="formData"
            :rules="basicRules"
            label-width="100px"
            style="max-width: 700px; margin-top: 20px;"
          >
            <el-form-item label="商品名称" prop="name">
              <el-input
                v-model="formData.name"
                placeholder="请输入商品名称"
                clearable
              />
            </el-form-item>
            <el-form-item label="副标题" prop="subTitle">
              <el-input
                v-model="formData.subTitle"
                placeholder="请输入副标题"
                clearable
              />
            </el-form-item>
            <el-form-item label="商品分类" prop="categoryId">
              <el-cascader
                v-model="formData.categoryId"
                :options="categoryTree"
                :props="{
                  value: 'ID',
                  label: 'name',
                  children: 'children',
                  checkStrictly: true,
                  emitPath: false
                }"
                placeholder="请选择商品分类"
                clearable
                style="width: 100%"
              />
            </el-form-item>
            <el-form-item label="品牌" prop="brandId">
              <el-select
                v-model="formData.brandId"
                placeholder="请选择品牌"
                clearable
                style="width: 100%"
              >
                <el-option
                  v-for="item in brandList"
                  :key="item.ID"
                  :label="item.name"
                  :value="item.ID"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="排序" prop="sort">
              <el-input-number
                v-model="formData.sort"
                :min="0"
                :max="9999"
                controls-position="right"
              />
            </el-form-item>
          </el-form>
        </el-tab-pane>

        <!-- Tab2: 商品图片 -->
        <el-tab-pane label="商品图片" name="images">
          <div style="max-width: 700px; margin-top: 20px;">
            <el-form label-width="100px">
              <el-form-item label="主图">
                <SelectImage
                  v-model="formData.mainImage"
                  file-type="image"
                />
              </el-form-item>
              <el-form-item label="商品图片集">
                <div class="image-list-box">
                  <div
                    v-for="(img, index) in imageList"
                    :key="index"
                    class="image-item"
                  >
                    <el-image
                      style="width: 100px; height: 100px"
                      :src="getUrl(img)"
                      fit="cover"
                    />
                    <div class="image-actions">
                      <el-button
                        v-if="index > 0"
                        type="primary"
                        link
                        size="small"
                        @click="moveImage(index, -1)"
                      >上移</el-button>
                      <el-button
                        v-if="index < imageList.length - 1"
                        type="primary"
                        link
                        size="small"
                        @click="moveImage(index, 1)"
                      >下移</el-button>
                      <el-button
                        type="danger"
                        link
                        size="small"
                        @click="removeImage(index)"
                      >删除</el-button>
                    </div>
                  </div>
                  <div class="image-add">
                    <SelectImage
                      v-model="tempImage"
                      file-type="image"
                      @update:model-value="onAddImage"
                    />
                  </div>
                </div>
              </el-form-item>
            </el-form>
          </div>
        </el-tab-pane>

        <!-- Tab3: SKU规格 -->
        <el-tab-pane label="SKU规格" name="sku">
          <div style="margin-top: 20px;">
            <div style="margin-bottom: 16px;">
              <el-button type="primary" icon="plus" @click="addSkuRow">添加SKU</el-button>
            </div>
            <el-table :data="formData.skus" style="width: 100%" border>
              <el-table-column label="SKU编码" min-width="120">
                <template #default="scope">
                  <el-input
                    v-model="scope.row.skuCode"
                    placeholder="SKU编码"
                    size="small"
                  />
                </template>
              </el-table-column>
              <el-table-column label="名称" min-width="140">
                <template #default="scope">
                  <el-input
                    v-model="scope.row.name"
                    placeholder="SKU名称"
                    size="small"
                  />
                </template>
              </el-table-column>
              <el-table-column label="规格数据(JSON)" min-width="180">
                <template #default="scope">
                  <el-input
                    v-model="scope.row.specData"
                    placeholder='例: {"颜色":"红","尺寸":"L"}'
                    size="small"
                  />
                </template>
              </el-table-column>
              <el-table-column label="售价" width="120">
                <template #default="scope">
                  <el-input-number
                    v-model="scope.row.price"
                    :min="0"
                    :precision="2"
                    :controls="false"
                    size="small"
                    style="width: 100%"
                  />
                </template>
              </el-table-column>
              <el-table-column label="市场价" width="120">
                <template #default="scope">
                  <el-input-number
                    v-model="scope.row.marketPrice"
                    :min="0"
                    :precision="2"
                    :controls="false"
                    size="small"
                    style="width: 100%"
                  />
                </template>
              </el-table-column>
              <el-table-column label="成本价" width="120">
                <template #default="scope">
                  <el-input-number
                    v-model="scope.row.costPrice"
                    :min="0"
                    :precision="2"
                    :controls="false"
                    size="small"
                    style="width: 100%"
                  />
                </template>
              </el-table-column>
              <el-table-column label="图片" width="100">
                <template #default="scope">
                  <SelectImage
                    v-model="scope.row.image"
                    file-type="image"
                  />
                </template>
              </el-table-column>
              <el-table-column label="重量(g)" width="100">
                <template #default="scope">
                  <el-input-number
                    v-model="scope.row.weight"
                    :min="0"
                    :controls="false"
                    size="small"
                    style="width: 100%"
                  />
                </template>
              </el-table-column>
              <el-table-column label="状态" width="80" align="center">
                <template #default="scope">
                  <el-switch
                    v-model="scope.row.status"
                    :active-value="1"
                    :inactive-value="0"
                    size="small"
                  />
                </template>
              </el-table-column>
              <el-table-column label="操作" width="80" align="center" fixed="right">
                <template #default="scope">
                  <el-button
                    type="danger"
                    link
                    icon="delete"
                    size="small"
                    @click="removeSkuRow(scope.$index)"
                  >删除</el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-tab-pane>

        <!-- Tab4: 商品详情 -->
        <el-tab-pane label="商品详情" name="detail">
          <div style="margin-top: 20px;">
            <RichEdit v-model="formData.detail" />
          </div>
        </el-tab-pane>
      </el-tabs>

      <div style="margin-top: 20px; text-align: center;">
        <el-button type="primary" size="large" @click="save">保 存</el-button>
        <el-button size="large" @click="back">返 回</el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  createSpu,
  updateSpu,
  getSpu
} from '@/api/shopProduct'
import { getCategoryTree } from '@/api/shopCategory'
import { getAllBrands } from '@/api/shopBrand'
import { getUrl } from '@/utils/image'
import SelectImage from '@/components/selectImage/selectImage.vue'
import RichEdit from '@/components/richtext/rich-edit.vue'
import { ElMessage } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

defineOptions({
  name: 'ShopProductForm'
})

const route = useRoute()
const router = useRouter()

const activeTab = ref('basic')

// =========== 辅助数据 ===========
const categoryTree = ref([])
const brandList = ref([])

const loadCategoryTree = async () => {
  const res = await getCategoryTree()
  if (res.code === 0) {
    categoryTree.value = res.data.list || []
  }
}

const loadBrandList = async () => {
  const res = await getAllBrands()
  if (res.code === 0) {
    brandList.value = res.data.list || []
  }
}

// =========== 表单数据 ===========
const type = ref('create')
const basicFormRef = ref()

const formData = ref({
  name: '',
  subTitle: '',
  categoryId: undefined,
  brandId: undefined,
  sort: 0,
  mainImage: '',
  images: '',
  detail: '',
  skus: [],
})

const imageList = ref([])
const tempImage = ref('')

const basicRules = reactive({
  name: [
    { required: true, message: '请输入商品名称', trigger: ['input', 'blur'] },
    { whitespace: true, message: '不能只输入空格', trigger: ['input', 'blur'] }
  ],
  categoryId: [
    { required: true, message: '请选择商品分类', trigger: 'change' }
  ],
})

// =========== 图片操作 ===========
const onAddImage = (val) => {
  if (val) {
    imageList.value.push(val)
    tempImage.value = ''
  }
}

const removeImage = (index) => {
  imageList.value.splice(index, 1)
}

const moveImage = (index, direction) => {
  const newIndex = index + direction
  if (newIndex < 0 || newIndex >= imageList.value.length) return
  const temp = imageList.value[index]
  imageList.value[index] = imageList.value[newIndex]
  imageList.value[newIndex] = temp
}

// =========== SKU操作 ===========
const initSkuRow = () => ({
  skuCode: '',
  name: '',
  specData: '',
  price: 0,
  marketPrice: 0,
  costPrice: 0,
  image: '',
  weight: 0,
  status: 1,
})

const addSkuRow = () => {
  formData.value.skus.push(initSkuRow())
}

const removeSkuRow = (index) => {
  formData.value.skus.splice(index, 1)
}

// =========== 初始化 ===========
const init = async () => {
  await loadCategoryTree()
  await loadBrandList()

  if (route.query.id) {
    type.value = 'update'
    const res = await getSpu({ ID: route.query.id })
    if (res.code === 0) {
      const data = res.data.data
      formData.value = {
        ...data,
        skus: data.skus || [],
      }
      // 解析图片列表
      if (data.images) {
        try {
          imageList.value = typeof data.images === 'string' ? JSON.parse(data.images) : data.images
        } catch (e) {
          imageList.value = data.images ? data.images.split(',') : []
        }
      }
    }
  } else {
    type.value = 'create'
  }
}

onMounted(() => {
  init()
})

// =========== 保存 ===========
const save = async () => {
  // 校验基本信息
  const basicValid = await new Promise((resolve) => {
    basicFormRef.value?.validate((valid) => {
      resolve(valid)
    })
  })
  if (!basicValid) {
    activeTab.value = 'basic'
    return
  }

  // 组装图片数据
  formData.value.images = JSON.stringify(imageList.value)

  let res
  switch (type.value) {
    case 'create':
      res = await createSpu(formData.value)
      break
    case 'update':
      res = await updateSpu(formData.value)
      break
    default:
      res = await createSpu(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '保存成功'
    })
    router.back()
  }
}

// 返回
const back = () => {
  router.back()
}
</script>

<style scoped>
.image-list-box {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
}
.image-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
  border: 1px dashed #dcdfe6;
  border-radius: 4px;
  padding: 8px;
}
.image-actions {
  display: flex;
  gap: 4px;
}
.image-add {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
