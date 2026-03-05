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
        <el-form-item label="SKU编码" prop="skuCode">
          <el-input
            v-model="searchInfo.skuCode"
            placeholder="请输入SKU编码"
          />
        </el-form-item>
        <el-form-item label="商品名称" prop="spuName">
          <el-input
            v-model="searchInfo.spuName"
            placeholder="请输入商品名称"
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
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
      >
        <el-table-column
          align="center"
          label="SKU ID"
          prop="skuId"
          width="100"
        />
        <el-table-column
          align="left"
          label="SKU编码"
          width="160"
        >
          <template #default="scope">
            {{ scope.row.sku?.skuCode || scope.row.skuCode || '-' }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="SKU名称"
          min-width="200"
        >
          <template #default="scope">
            {{ scope.row.sku?.name || scope.row.skuName || '-' }}
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="总库存"
          prop="totalStock"
          width="120"
        />
        <el-table-column
          align="center"
          label="锁定库存"
          prop="lockedStock"
          width="120"
        >
          <template #default="scope">
            <el-tag v-if="scope.row.lockedStock > 0" type="warning">{{ scope.row.lockedStock }}</el-tag>
            <span v-else>{{ scope.row.lockedStock || 0 }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="可用库存"
          width="120"
        >
          <template #default="scope">
            <span :style="{ color: availableStock(scope.row) <= 0 ? '#F56C6C' : '' }">
              {{ availableStock(scope.row) }}
            </span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="预警" width="100">
          <template #default="scope">
            <el-tag v-if="availableStock(scope.row) <= 0" type="danger" size="small">缺货</el-tag>
            <el-tag v-else-if="availableStock(scope.row) <= 10" type="warning" size="small">低库存</el-tag>
            <el-tag v-else type="success" size="small">正常</el-tag>
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
              @click="openStockDialog(scope.row)"
            >设置库存</el-button>
            <el-button
              type="primary"
              link
              class="table-button"
              @click="openLogDialog(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              库存日志
            </el-button>
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

    <!-- 设置库存对话框 -->
    <el-dialog
      v-model="stockDialogVisible"
      title="设置库存"
      width="450px"
      destroy-on-close
    >
      <el-form
        ref="stockFormRef"
        :model="stockForm"
        :rules="stockRules"
        label-width="100px"
      >
        <el-form-item label="SKU编码">
          <el-input :model-value="stockForm.skuCode" disabled />
        </el-form-item>
        <el-form-item label="SKU名称">
          <el-input :model-value="stockForm.skuName" disabled />
        </el-form-item>
        <el-form-item label="当前库存">
          <el-input :model-value="stockForm.currentStock" disabled />
        </el-form-item>
        <el-form-item label="新库存值" prop="totalStock">
          <el-input-number
            v-model="stockForm.totalStock"
            :min="0"
            :max="999999"
            controls-position="right"
            style="width: 100%"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="stockDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="confirmSetStock">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 库存日志对话框 -->
    <el-dialog
      v-model="logDialogVisible"
      title="库存日志"
      width="700px"
      destroy-on-close
    >
      <div style="margin-bottom: 10px;">
        <span style="font-weight: bold;">SKU: </span>
        <span>{{ logSkuInfo.skuCode }} - {{ logSkuInfo.skuName }}</span>
      </div>
      <el-table :data="logTableData" border style="width: 100%">
        <el-table-column
          align="left"
          label="时间"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column
          align="center"
          label="变动类型"
          prop="changeType"
          width="120"
        >
          <template #default="scope">
            <el-tag :type="logTypeTag(scope.row.changeType)">
              {{ scope.row.changeType || '-' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="变动数量"
          prop="changeAmount"
          width="100"
        >
          <template #default="scope">
            <span :style="{ color: scope.row.changeAmount > 0 ? '#67C23A' : '#F56C6C' }">
              {{ scope.row.changeAmount > 0 ? '+' : '' }}{{ scope.row.changeAmount }}
            </span>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="变动后库存"
          prop="afterStock"
          width="100"
        />
        <el-table-column
          align="left"
          label="备注"
          prop="remark"
          min-width="150"
          show-overflow-tooltip
        />
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="logPage"
          :page-size="logPageSize"
          :page-sizes="[10, 30, 50]"
          :total="logTotal"
          @current-change="handleLogCurrentChange"
          @size-change="handleLogSizeChange"
        />
      </div>
      <template #footer>
        <el-button @click="logDialogVisible = false">关 闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  setStock,
  getInventoryList,
  getInventoryLogList
} from '@/api/shopInventory'

import { formatDate } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { InfoFilled } from '@element-plus/icons-vue'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'ShopInventory'
})

// =========== 辅助方法 ===========
const availableStock = (row) => {
  return (row.totalStock || 0) - (row.lockedStock || 0)
}

const logTypeTag = (type) => {
  if (!type) return 'info'
  const map = {
    'add': 'success',
    'set': 'primary',
    'lock': 'warning',
    'unlock': '',
    'deduct': 'danger',
  }
  return map[type] || 'info'
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
  const table = await getInventoryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 设置库存 ===============
const stockDialogVisible = ref(false)
const stockFormRef = ref()
const stockForm = ref({
  ID: 0,
  skuId: 0,
  skuCode: '',
  skuName: '',
  currentStock: 0,
  totalStock: 0,
})

const stockRules = reactive({
  totalStock: [
    { required: true, message: '请输入新库存值', trigger: ['input', 'blur'] }
  ],
})

const openStockDialog = (row) => {
  stockForm.value = {
    ID: row.ID,
    skuId: row.skuId,
    skuCode: row.sku?.skuCode || row.skuCode || '',
    skuName: row.sku?.name || row.skuName || '',
    currentStock: row.totalStock || 0,
    totalStock: row.totalStock || 0,
  }
  stockDialogVisible.value = true
}

const confirmSetStock = async () => {
  stockFormRef.value?.validate(async (valid) => {
    if (!valid) return
    const res = await setStock({
      ID: stockForm.value.ID,
      skuId: stockForm.value.skuId,
      totalStock: stockForm.value.totalStock,
    })
    if (res.code === 0) {
      ElMessage({ type: 'success', message: '库存设置成功' })
      stockDialogVisible.value = false
      getTableData()
    }
  })
}

// ============== 库存日志 ===============
const logDialogVisible = ref(false)
const logTableData = ref([])
const logPage = ref(1)
const logTotal = ref(0)
const logPageSize = ref(10)
const logSkuInfo = ref({
  skuId: 0,
  skuCode: '',
  skuName: '',
})

const openLogDialog = (row) => {
  logSkuInfo.value = {
    skuId: row.skuId,
    skuCode: row.sku?.skuCode || row.skuCode || '',
    skuName: row.sku?.name || row.skuName || '',
  }
  logPage.value = 1
  logDialogVisible.value = true
  loadLogData()
}

const handleLogSizeChange = (val) => {
  logPageSize.value = val
  loadLogData()
}

const handleLogCurrentChange = (val) => {
  logPage.value = val
  loadLogData()
}

const loadLogData = async () => {
  const res = await getInventoryLogList({
    page: logPage.value,
    pageSize: logPageSize.value,
    skuId: logSkuInfo.value.skuId,
  })
  if (res.code === 0) {
    logTableData.value = res.data.list || []
    logTotal.value = res.data.total || 0
    logPage.value = res.data.page || 1
    logPageSize.value = res.data.pageSize || 10
  }
}
</script>

<style scoped>
</style>
