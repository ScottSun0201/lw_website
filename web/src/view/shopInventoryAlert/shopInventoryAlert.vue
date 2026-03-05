<template>
  <div>
    <!-- 全局预警设置 -->
    <el-card style="margin-bottom: 20px;">
      <template #header>
        <div style="display: flex; align-items: center; justify-content: space-between;">
          <span>库存预警设置</span>
          <el-button type="primary" size="small" @click="saveSetting">保存设置</el-button>
        </div>
      </template>
      <el-form :inline="true" :model="alertSetting" label-width="120px">
        <el-form-item label="低库存阈值">
          <el-input-number
            v-model="alertSetting.lowThreshold"
            :min="0"
            :max="99999"
            controls-position="right"
            style="width: 200px"
          />
          <span style="margin-left: 8px; color: #909399;">库存低于此值触发低库存预警</span>
        </el-form-item>
        <el-form-item label="缺货阈值">
          <el-input-number
            v-model="alertSetting.emptyThreshold"
            :min="0"
            :max="99999"
            controls-position="right"
            style="width: 200px"
          />
          <span style="margin-left: 8px; color: #909399;">库存低于此值触发缺货预警</span>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 预警列表 -->
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
            style="width: 140px"
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
          label="SKU ID"
          prop="skuId"
          width="100"
        />
        <el-table-column
          align="center"
          label="预警类型"
          prop="type"
          width="120"
        >
          <template #default="scope">
            <el-tag :type="scope.row.type === 1 ? 'warning' : 'danger'">
              {{ scope.row.type === 1 ? '低库存' : '缺货' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="阈值"
          prop="threshold"
          width="100"
        />
        <el-table-column
          align="center"
          label="当前库存"
          prop="currentStock"
          width="100"
        >
          <template #default="scope">
            <span :style="{ color: scope.row.currentStock <= 0 ? '#F56C6C' : '' }">
              {{ scope.row.currentStock }}
            </span>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="状态"
          prop="status"
          width="100"
        >
          <template #default="scope">
            <el-tag :type="scope.row.status === 0 ? 'danger' : 'success'">
              {{ scope.row.status === 0 ? '未处理' : '已处理' }}
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
              @click="handleAlertFunc(scope.row)"
            >标为已处理</el-button>
            <span v-else style="color: #999;">已处理</span>
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
  getAlertList,
  handleAlert,
  getSetting,
  updateSetting
} from '@/api/shopInventoryAlert'

import { formatDate } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { ref, onMounted } from 'vue'

defineOptions({
  name: 'ShopInventoryAlert'
})

// =========== 预警设置 ===========
const alertSetting = ref({
  lowThreshold: 10,
  emptyThreshold: 0,
})

const loadSetting = async () => {
  const res = await getSetting()
  if (res.code === 0) {
    alertSetting.value = {
      lowThreshold: res.data.data?.lowThreshold ?? 10,
      emptyThreshold: res.data.data?.emptyThreshold ?? 0,
    }
  }
}

const saveSetting = async () => {
  const res = await updateSetting(alertSetting.value)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '设置保存成功' })
  }
}

// =========== 类型映射 ===========
const typeOptions = [
  { label: '低库存', value: 1 },
  { label: '缺货', value: 2 },
]

const statusOptions = [
  { label: '未处理', value: 0 },
  { label: '已处理', value: 1 },
]

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
  const table = await getAlertList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

onMounted(() => {
  loadSetting()
  getTableData()
})

// ============== 标为已处理 ===============
const handleAlertFunc = async (row) => {
  const res = await handleAlert({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '已标为已处理' })
    getTableData()
  }
}
</script>

<style scoped>
</style>
