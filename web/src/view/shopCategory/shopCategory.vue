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
        <el-form-item label="分类名称" prop="name">
          <el-input
            v-model="searchInfo.name"
            placeholder="请输入分类名称"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select
            v-model="searchInfo.status"
            placeholder="请选择状态"
            clearable
          >
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
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
          @click="openDialog(null)"
        >新增顶级分类</el-button>
      </div>
      <el-table
        ref="tableRef"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        default-expand-all
      >
        <el-table-column
          align="left"
          label="分类名称"
          prop="name"
          min-width="200"
        />
        <el-table-column
          align="center"
          label="层级"
          prop="level"
          width="100"
        >
          <template #default="scope">
            <el-tag v-if="scope.row.level === 1" type="primary">一级</el-tag>
            <el-tag v-else-if="scope.row.level === 2" type="success">二级</el-tag>
            <el-tag v-else-if="scope.row.level === 3" type="warning">三级</el-tag>
            <el-tag v-else type="info">{{ scope.row.level }}</el-tag>
          </template>
        </el-table-column>
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
          align="center"
          label="操作"
          fixed="right"
          min-width="280"
        >
          <template #default="scope">
            <el-button
              v-if="scope.row.level < 3"
              type="primary"
              link
              icon="plus"
              class="table-button"
              @click="openDialog(scope.row)"
            >添加子分类</el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateCategoryFunc(scope.row)"
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
    </div>

    <!-- 新增/编辑对话框 -->
    <el-dialog
      v-model="dialogFormVisible"
      :title="type === 'create' ? '新增分类' : '编辑分类'"
      width="500px"
      destroy-on-close
    >
      <el-form
        ref="elFormRef"
        :model="formData"
        :rules="rule"
        label-width="100px"
      >
        <el-form-item label="分类名称" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="请输入分类名称"
            clearable
          />
        </el-form-item>
        <el-form-item label="上级分类" prop="parentId">
          <el-cascader
            v-model="formData.parentId"
            :options="categoryTreeOptions"
            :props="{
              value: 'ID',
              label: 'name',
              children: 'children',
              checkStrictly: true,
              emitPath: false
            }"
            placeholder="请选择上级分类（留空为顶级）"
            clearable
            style="width: 100%"
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
        <el-form-item label="图标" prop="icon">
          <el-input
            v-model="formData.icon"
            placeholder="请输入图标名称或URL"
            clearable
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
  createShopCategory,
  deleteShopCategory,
  updateShopCategory,
  findShopCategory,
  getCategoryTree
} from '@/api/shopCategory'

import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'ShopCategory'
})

// =========== 表格控制部分 ===========
const tableData = ref([])
const searchInfo = ref({})
const categoryTreeOptions = ref([])

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  getTableData()
}

// 查询分类树
const getTableData = async () => {
  const res = await getCategoryTree(searchInfo.value)
  if (res.code === 0) {
    tableData.value = res.data.list || []
  }
}

// 获取分类树选项（用于cascader）
const getCategoryOptions = async () => {
  const res = await getCategoryTree()
  if (res.code === 0) {
    categoryTreeOptions.value = res.data.list || []
  }
}

getTableData()
getCategoryOptions()

// ============== 表格控制部分结束 ===============

// 表单数据
const formData = ref({
  name: '',
  parentId: 0,
  sort: 0,
  icon: '',
  status: 1,
})

// 验证规则
const rule = reactive({
  name: [
    { required: true, message: '请输入分类名称', trigger: ['input', 'blur'] },
    { whitespace: true, message: '不能只输入空格', trigger: ['input', 'blur'] }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// 行为控制标记
const type = ref('')
const dialogFormVisible = ref(false)

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除该分类吗？删除后子分类也会被删除！', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteCategoryFunc(row)
  })
}

// 删除
const deleteCategoryFunc = async (row) => {
  const res = await deleteShopCategory({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    getTableData()
    getCategoryOptions()
  }
}

// 打开弹窗 - parentRow 为 null 时为顶级分类
const openDialog = (parentRow) => {
  type.value = 'create'
  formData.value = {
    name: '',
    parentId: parentRow ? parentRow.ID : 0,
    sort: 0,
    icon: '',
    status: 1,
  }
  dialogFormVisible.value = true
}

// 更新行
const updateCategoryFunc = async (row) => {
  const res = await findShopCategory({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.data
    dialogFormVisible.value = true
  }
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    parentId: 0,
    sort: 0,
    icon: '',
    status: 1,
  }
}

// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createShopCategory(formData.value)
        break
      case 'update':
        res = await updateShopCategory(formData.value)
        break
      default:
        res = await createShopCategory(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
      getCategoryOptions()
    }
  })
}
</script>

<style scoped>
</style>
