<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="品牌名称" prop="name">
          <el-input
            v-model="searchInfo.name"
            placeholder="请输入品牌名称"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
          >查询</el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog"
        >新增</el-button>
        <el-popover
          v-model:visible="deleteVisible"
          :disabled="!multipleSelection.length"
          placement="top"
          width="160"
        >
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button
              type="primary"
              link
              @click="deleteVisible = false"
            >取消</el-button>
            <el-button
              type="primary"
              @click="onDelete"
            >确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="center"
          label="Logo"
          width="100"
        >
          <template #default="scope">
            <el-image
              v-if="scope.row.logo"
              style="width: 50px; height: 50px"
              :src="getUrl(scope.row.logo)"
              fit="cover"
              :preview-src-list="[getUrl(scope.row.logo)]"
              preview-teleported
            />
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="品牌名称"
          prop="name"
          min-width="150"
        />
        <el-table-column
          align="left"
          label="描述"
          prop="desc"
          min-width="200"
          show-overflow-tooltip
        />
        <el-table-column
          align="center"
          label="排序"
          prop="sort"
          width="100"
        />
        <el-table-column
          align="center"
          label="状态"
          prop="status"
          width="100"
        >
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="200"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateBrandFunc(scope.row)"
            >编辑</el-button>
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- 新增/编辑对话框 -->
    <el-dialog
      v-model="dialogFormVisible"
      :title="type === 'create' ? '新增品牌' : '编辑品牌'"
      width="500px"
      destroy-on-close
    >
      <el-form
        ref="elFormRef"
        :model="formData"
        :rules="rule"
        label-width="100px"
      >
        <el-form-item label="品牌名称" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="请输入品牌名称"
            clearable
          />
        </el-form-item>
        <el-form-item label="Logo" prop="logo">
          <SelectImage
            v-model="formData.logo"
            file-type="image"
          />
        </el-form-item>
        <el-form-item label="描述" prop="desc">
          <el-input
            v-model="formData.desc"
            type="textarea"
            :rows="3"
            placeholder="请输入品牌描述"
            clearable
          />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number
            v-model="formData.sort"
            :min="0"
            :max="9999"
            controls-position="right"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-switch
            v-model="formData.status"
            :active-value="1"
            :inactive-value="0"
            active-text="启用"
            inactive-text="禁用"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createShopBrand,
  deleteShopBrand,
  updateShopBrand,
  findShopBrand,
  getShopBrandList
} from '@/api/shopBrand'

import { getUrl } from '@/utils/image'
import SelectImage from '@/components/selectImage/selectImage.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'ShopBrand'
})

// 初始表单数据
const initFormData = () => ({
  name: '',
  logo: '',
  desc: '',
  sort: 0,
  status: 1,
})

const formData = ref(initFormData())

// 验证规则
const rule = reactive({
  name: [
    { required: true, message: '请输入品牌名称', trigger: ['input', 'blur'] },
    { whitespace: true, message: '不能只输入空格', trigger: ['input', 'blur'] }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getShopBrandList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 多选数据
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteBrandFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
  const ids = multipleSelection.value.map(item => item.ID)
  if (ids.length === 0) {
    ElMessage({ type: 'warning', message: '请选择要删除的数据' })
    return
  }
  for (const id of ids) {
    await deleteShopBrand({ ID: id })
  }
  ElMessage({ type: 'success', message: '删除成功' })
  if (tableData.value.length === ids.length && page.value > 1) {
    page.value--
  }
  deleteVisible.value = false
  getTableData()
}

// 行为控制标记
const type = ref('')

// 更新行
const updateBrandFunc = async (row) => {
  const res = await findShopBrand({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.data
    dialogFormVisible.value = true
  }
}

// 删除
const deleteBrandFunc = async (row) => {
  const res = await deleteShopBrand({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '删除成功' })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  formData.value = initFormData()
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = initFormData()
}

// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createShopBrand(formData.value)
        break
      case 'update':
        res = await updateShopBrand(formData.value)
        break
      default:
        res = await createShopBrand(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}
</script>

<style scoped>
</style>
