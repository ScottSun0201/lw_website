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
        <el-form-item label="订单号" prop="orderNo">
          <el-input
            v-model="searchInfo.orderNo"
            placeholder="请输入订单号"
          />
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
      <!-- 状态快捷标签 -->
      <el-tabs v-model="activeStatusTab" @tab-change="onTabChange" style="margin-bottom: 10px;">
        <el-tab-pane label="全部" name="all" />
        <el-tab-pane label="待付款" name="0" />
        <el-tab-pane label="已付款" name="1" />
        <el-tab-pane label="已发货" name="2" />
        <el-tab-pane label="已收货" name="3" />
        <el-tab-pane label="已完成" name="4" />
        <el-tab-pane label="已取消" name="5" />
        <el-tab-pane label="退款中" name="6" />
        <el-tab-pane label="已退款" name="7" />
      </el-tabs>

      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
      >
        <el-table-column
          align="left"
          label="订单号"
          prop="orderNo"
          min-width="200"
        />
        <el-table-column
          align="center"
          label="用户ID"
          prop="userId"
          width="100"
        />
        <el-table-column
          align="center"
          label="状态"
          prop="status"
          width="100"
        >
          <template #default="scope">
            <el-tag :type="statusTagType(scope.row.status)">
              {{ statusLabel(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="right"
          label="支付金额"
          width="120"
        >
          <template #default="scope">
            {{ scope.row.payAmount != null ? '¥' + Number(scope.row.payAmount).toFixed(2) : '-' }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="收货人"
          prop="receiverName"
          width="120"
        />
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
          min-width="240"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              class="table-button"
              @click="viewDetail(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              查看详情
            </el-button>
            <el-button
              v-if="scope.row.status === 1"
              type="success"
              link
              class="table-button"
              @click="openShipDialog(scope.row)"
            >发货</el-button>
            <el-button
              v-if="scope.row.status === 0 || scope.row.status === 1"
              type="danger"
              link
              class="table-button"
              @click="cancelOrder(scope.row)"
            >取消</el-button>
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

    <!-- 订单详情对话框 -->
    <el-dialog
      v-model="detailVisible"
      title="订单详情"
      width="800px"
      destroy-on-close
    >
      <el-descriptions :column="2" border style="margin-bottom: 20px;">
        <el-descriptions-item label="订单号">{{ detailData.orderNo }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="statusTagType(detailData.status)">{{ statusLabel(detailData.status) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="用户ID">{{ detailData.userId }}</el-descriptions-item>
        <el-descriptions-item label="支付金额">
          {{ detailData.payAmount != null ? '¥' + Number(detailData.payAmount).toFixed(2) : '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="收货人">{{ detailData.receiverName }}</el-descriptions-item>
        <el-descriptions-item label="收货电话">{{ detailData.receiverPhone }}</el-descriptions-item>
        <el-descriptions-item label="收货地址" :span="2">{{ detailData.receiverAddress }}</el-descriptions-item>
        <el-descriptions-item label="物流公司">{{ detailData.shipCompany || '-' }}</el-descriptions-item>
        <el-descriptions-item label="物流单号">{{ detailData.shipNo || '-' }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ formatDate(detailData.CreatedAt) }}</el-descriptions-item>
        <el-descriptions-item label="支付时间">{{ detailData.payTime ? formatDate(detailData.payTime) : '-' }}</el-descriptions-item>
      </el-descriptions>

      <!-- 订单商品 -->
      <h4 style="margin-bottom: 10px;">商品明细</h4>
      <el-table :data="detailData.items || []" border style="margin-bottom: 20px;">
        <el-table-column label="商品名称" prop="productName" min-width="180" />
        <el-table-column label="SKU名称" prop="skuName" min-width="120" />
        <el-table-column label="单价" width="100" align="right">
          <template #default="scope">¥{{ Number(scope.row.price || 0).toFixed(2) }}</template>
        </el-table-column>
        <el-table-column label="数量" prop="quantity" width="80" align="center" />
        <el-table-column label="小计" width="100" align="right">
          <template #default="scope">¥{{ Number((scope.row.price || 0) * (scope.row.quantity || 0)).toFixed(2) }}</template>
        </el-table-column>
      </el-table>

      <!-- 订单日志 -->
      <h4 style="margin-bottom: 10px;">操作日志</h4>
      <el-timeline v-if="orderLogs.length > 0">
        <el-timeline-item
          v-for="(log, index) in orderLogs"
          :key="index"
          :timestamp="formatDate(log.CreatedAt)"
          placement="top"
        >
          <p>{{ log.action || log.content || log.remark }}</p>
          <p v-if="log.operator" style="color: #999; font-size: 12px;">操作人: {{ log.operator }}</p>
        </el-timeline-item>
      </el-timeline>
      <el-empty v-else description="暂无操作日志" :image-size="60" />

      <!-- 物流轨迹 -->
      <h4 style="margin-bottom: 10px;">物流轨迹</h4>
      <el-timeline v-if="logisticsData.length > 0">
        <el-timeline-item
          v-for="(item, index) in logisticsData"
          :key="index"
          :timestamp="formatDate(item.traceTime || item.CreatedAt)"
          placement="top"
        >
          <p>{{ item.detail }}</p>
        </el-timeline-item>
      </el-timeline>
      <el-empty v-else description="暂无物流信息" :image-size="60" />
      <el-button v-if="detailData.status >= 2 && detailData.status <= 3" type="primary" size="small" style="margin-top: 10px;" @click="openLogisticsDialog">添加物流轨迹</el-button>

      <template #footer>
        <el-button @click="detailVisible = false">关 闭</el-button>
      </template>
    </el-dialog>

    <!-- 物流轨迹添加对话框 -->
    <el-dialog v-model="logisticsDialogVisible" title="添加物流轨迹" width="450px" destroy-on-close>
      <el-form ref="logisticsFormRef" :model="logisticsForm" label-width="100px">
        <el-form-item label="物流详情">
          <el-input v-model="logisticsForm.detail" type="textarea" :rows="3" placeholder="请输入物流详情" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="logisticsDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="confirmAddTrace">确 定</el-button>
      </template>
    </el-dialog>

    <!-- 发货对话框 -->
    <el-dialog
      v-model="shipDialogVisible"
      title="订单发货"
      width="450px"
      destroy-on-close
    >
      <el-form
        ref="shipFormRef"
        :model="shipForm"
        :rules="shipRules"
        label-width="100px"
      >
        <el-form-item label="物流公司" prop="shipCompany">
          <el-input
            v-model="shipForm.shipCompany"
            placeholder="请输入物流公司"
            clearable
          />
        </el-form-item>
        <el-form-item label="物流单号" prop="shipNo">
          <el-input
            v-model="shipForm.shipNo"
            placeholder="请输入物流单号"
            clearable
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="shipDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="confirmShip">确认发货</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getOrderList,
  getOrderDetail,
  shipOrder,
  adminCancelOrder,
  getOrderLogs
} from '@/api/shopOrder'
import { addTrace, getTraces } from '@/api/shopLogistics'

import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { InfoFilled } from '@element-plus/icons-vue'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'ShopOrder'
})

// =========== 状态映射 ===========
const statusOptions = [
  { label: '待付款', value: 0 },
  { label: '已付款', value: 1 },
  { label: '已发货', value: 2 },
  { label: '已收货', value: 3 },
  { label: '已完成', value: 4 },
  { label: '已取消', value: 5 },
  { label: '退款中', value: 6 },
  { label: '已退款', value: 7 },
]

const statusTagType = (status) => {
  const map = { 0: 'warning', 1: 'primary', 2: '', 3: 'success', 4: 'success', 5: 'info', 6: 'danger', 7: 'info' }
  return map[status] || 'info'
}

const statusLabel = (status) => {
  const map = { 0: '待付款', 1: '已付款', 2: '已发货', 3: '已收货', 4: '已完成', 5: '已取消', 6: '退款中', 7: '已退款' }
  return map[status] !== undefined ? map[status] : '未知'
}

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const activeStatusTab = ref('all')

const onReset = () => {
  searchInfo.value = {}
  activeStatusTab.value = 'all'
  getTableData()
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const onTabChange = (tab) => {
  if (tab === 'all') {
    delete searchInfo.value.status
  } else {
    searchInfo.value.status = Number(tab)
  }
  page.value = 1
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
  const params = { page: page.value, pageSize: pageSize.value, ...searchInfo.value }
  const table = await getOrderList(params)
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 订单详情 ===============
const detailVisible = ref(false)
const detailData = ref({})
const orderLogs = ref([])

const viewDetail = async (row) => {
  const res = await getOrderDetail({ orderNo: row.orderNo })
  if (res.code === 0) {
    detailData.value = res.data.data
    orderLogs.value = res.data.logs || []
    logisticsData.value = []
    if (res.data.data.status >= 2) {
      loadLogistics(row.orderNo)
    }
    detailVisible.value = true
  }
}

const loadOrderLogs = async (orderNo) => {
  const res = await getOrderLogs({ orderNo })
  if (res.code === 0) {
    orderLogs.value = res.data.list || []
  }
}

// ============== 物流轨迹 ===============
const logisticsData = ref([])

const loadLogistics = async (orderNo) => {
  const res = await getTraces({ orderNo })
  if (res.code === 0) {
    logisticsData.value = res.data.list || []
  }
}

const logisticsDialogVisible = ref(false)
const logisticsFormRef = ref()
const logisticsForm = ref({ orderNo: '', detail: '' })

const openLogisticsDialog = () => {
  logisticsForm.value = { orderNo: detailData.value.orderNo, detail: '' }
  logisticsDialogVisible.value = true
}

const confirmAddTrace = async () => {
  const res = await addTrace({
    orderNo: logisticsForm.value.orderNo,
    detail: logisticsForm.value.detail,
    shipNo: detailData.value.shipNo || '',
    shipCompany: detailData.value.shipCompany || '',
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '添加成功' })
    logisticsDialogVisible.value = false
    loadLogistics(logisticsForm.value.orderNo)
  }
}

// ============== 发货 ===============
const shipDialogVisible = ref(false)
const shipFormRef = ref()
const shipForm = ref({
  orderNo: '',
  shipCompany: '',
  shipNo: '',
})

const shipRules = reactive({
  shipCompany: [
    { required: true, message: '请输入物流公司', trigger: ['input', 'blur'] }
  ],
  shipNo: [
    { required: true, message: '请输入物流单号', trigger: ['input', 'blur'] }
  ],
})

const openShipDialog = (row) => {
  shipForm.value = {
    orderNo: row.orderNo,
    shipCompany: '',
    shipNo: '',
  }
  shipDialogVisible.value = true
}

const confirmShip = async () => {
  shipFormRef.value?.validate(async (valid) => {
    if (!valid) return
    const res = await shipOrder(shipForm.value)
    if (res.code === 0) {
      ElMessage({ type: 'success', message: '发货成功' })
      shipDialogVisible.value = false
      getTableData()
    }
  })
}

// ============== 取消订单 ===============
const cancelOrder = (row) => {
  ElMessageBox.confirm('确定要取消该订单吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await adminCancelOrder({ orderNo: row.orderNo })
    if (res.code === 0) {
      ElMessage({ type: 'success', message: '订单已取消' })
      getTableData()
    }
  })
}
</script>

<style scoped>
</style>
