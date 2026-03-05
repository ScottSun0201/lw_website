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
        <el-form-item label="类型" prop="type">
          <el-select
            v-model="searchInfo.type"
            placeholder="请选择类型"
            clearable
            style="width: 160px"
          >
            <el-option
              v-for="item in typeOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
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
      <!-- 类型标签页 -->
      <el-tabs v-model="activeTypeTab" @tab-change="onTabChange" style="margin-bottom: 10px;">
        <el-tab-pane label="全部" name="all" />
        <el-tab-pane label="热销" name="1" />
        <el-tab-pane label="新品" name="2" />
        <el-tab-pane label="猜你喜欢" name="3" />
        <el-tab-pane label="相关推荐" name="4" />
        <el-tab-pane label="人工推荐" name="5" />
      </el-tabs>

      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog"
        >新增推荐</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
      >
        <el-table-column
          align="center"
          label="ID"
          prop="ID"
          width="80"
        />
        <el-table-column
          align="center"
          label="类型"
          prop="type"
          width="120"
        >
          <template #default="scope">
            <el-tag :type="typeTagType(scope.row.type)">
              {{ typeLabel(scope.row.type) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="商品ID"
          prop="spuId"
          width="100"
        />
        <el-table-column
          align="center"
          label="排序"
          width="150"
        >
          <template #default="scope">
            <el-input-number
              v-model="scope.row.sort"
              :min="0"
              :max="9999"
              size="small"
              controls-position="right"
              @change="onSortChange(scope.row)"
            />
          </template>
        </el-table-column>
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
          label="创建时间"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
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
              @click="updateFunc(scope.row)"
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
      :title="dialogType === 'create' ? '新增推荐' : '编辑推荐'"
      width="500px"
      destroy-on-close
    >
      <el-form
        ref="elFormRef"
        :model="formData"
        :rules="rule"
        label-width="100px"
      >
        <el-form-item label="类型" prop="type">
          <el-select
            v-model="formData.type"
            placeholder="请选择类型"
            style="width: 100%"
          >
            <el-option
              v-for="item in typeOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="商品ID" prop="spuId">
          <el-input-number
            v-model="formData.spuId"
            :min="1"
            controls-position="right"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number
            v-model="formData.sort"
            :min="0"
            :max="9999"
            controls-position="right"
            style="width: 100%"
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
  getRecommendationList,
  setRecommendation,
  removeRecommendation,
  updateSort
} from '@/api/shopRecommendation'

import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'ShopRecommendation'
})

// =========== 类型映射 ===========
const typeOptions = [
  { label: '热销', value: 1 },
  { label: '新品', value: 2 },
  { label: '猜你喜欢', value: 3 },
  { label: '相关推荐', value: 4 },
  { label: '人工推荐', value: 5 },
]

const typeTagType = (type) => {
  const map = { 1: 'danger', 2: 'success', 3: '', 4: 'warning', 5: 'primary' }
  return map[type] || 'info'
}

const typeLabel = (type) => {
  const map = { 1: '热销', 2: '新品', 3: '猜你喜欢', 4: '相关推荐', 5: '人工推荐' }
  return map[type] || '未知'
}

// 初始表单数据
const initFormData = () => ({
  type: 1,
  spuId: undefined,
  sort: 0,
  status: 1,
})

const formData = ref(initFormData())

// 验证规则
const rule = reactive({
  type: [
    { required: true, message: '请选择类型', trigger: 'change' },
  ],
  spuId: [
    { required: true, message: '请输入商品ID', trigger: ['input', 'blur'] },
  ],
})

const elFormRef = ref()

// =========== Tab 控制 ===========
const activeTypeTab = ref('all')

const onTabChange = (tab) => {
  if (tab === 'all') {
    delete searchInfo.value.type
  } else {
    searchInfo.value.type = Number(tab)
  }
  page.value = 1
  getTableData()
}

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
  activeTypeTab.value = 'all'
  getTableData()
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const getTableData = async () => {
  const table = await getRecommendationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== CRUD ===============

// 行为控制标记
const dialogType = ref('')

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  dialogType.value = 'create'
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
    const res = await setRecommendation(formData.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: dialogType.value === 'create' ? '新增成功' : '更新成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

// 编辑行
const updateFunc = (row) => {
  dialogType.value = 'update'
  formData.value = {
    ID: row.ID,
    type: row.type,
    spuId: row.spuId,
    sort: row.sort,
    status: row.status,
  }
  dialogFormVisible.value = true
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要移除该推荐吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await removeRecommendation({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({ type: 'success', message: '删除成功' })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 排序变更
const onSortChange = async (row) => {
  const res = await updateSort({ ID: row.ID, sort: row.sort })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '排序更新成功' })
  }
}
</script>

<style scoped>
</style>
