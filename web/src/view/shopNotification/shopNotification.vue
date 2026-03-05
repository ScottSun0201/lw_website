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
        <el-form-item label="状态" prop="status">
          <el-select
            v-model="searchInfo.status"
            placeholder="请选择状态"
            clearable
            style="width: 140px"
          >
            <el-option
              v-for="item in statusOptions"
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
      <div class="gva-btn-list">
        <el-button
          type="primary"
          @click="markAllRead"
        >全部已读</el-button>
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
          align="left"
          label="标题"
          prop="title"
          min-width="180"
        />
        <el-table-column
          align="left"
          label="内容"
          prop="content"
          min-width="250"
          show-overflow-tooltip
        />
        <el-table-column
          align="center"
          label="状态"
          prop="status"
          width="100"
        >
          <template #default="scope">
            <el-tag :type="scope.row.status === 0 ? 'warning' : 'success'">
              {{ scope.row.status === 0 ? '未读' : '已读' }}
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
          min-width="120"
        >
          <template #default="scope">
            <el-button
              v-if="scope.row.status === 0"
              type="primary"
              link
              class="table-button"
              @click="markRead(scope.row)"
            >标为已读</el-button>
            <span v-else style="color: #999;">已读</span>
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
  </div>
</template>

<script setup>
import {
  getNotificationList,
  markAdminRead,
  markAllAdminRead
} from '@/api/shopNotification'

import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

defineOptions({
  name: 'ShopNotification'
})

// =========== 类型映射 ===========
const typeOptions = [
  { label: '库存预警', value: 1 },
  { label: '退款申请', value: 2 },
  { label: '新订单', value: 3 },
  { label: '新评价', value: 4 },
]

const statusOptions = [
  { label: '未读', value: 0 },
  { label: '已读', value: 1 },
]

const typeTagType = (type) => {
  const map = { 1: 'danger', 2: 'warning', 3: 'success', 4: '' }
  return map[type] || 'info'
}

const typeLabel = (type) => {
  const map = { 1: '库存预警', 2: '退款申请', 3: '新订单', 4: '新评价' }
  return map[type] || '未知'
}

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
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
  const table = await getNotificationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 标为已读 ===============
const markRead = async (row) => {
  const res = await markAdminRead({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '已标为已读' })
    getTableData()
  }
}

// ============== 全部已读 ===============
const markAllRead = () => {
  ElMessageBox.confirm('确定要将所有通知标为已读吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await markAllAdminRead()
    if (res.code === 0) {
      ElMessage({ type: 'success', message: '已全部标为已读' })
      getTableData()
    }
  })
}
</script>

<style scoped>
</style>
