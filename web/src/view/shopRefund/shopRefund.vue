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
        <el-form-item label="退款单号" prop="refundNo">
          <el-input
            v-model="searchInfo.refundNo"
            placeholder="请输入退款单号"
          />
        </el-form-item>
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
          align="left"
          label="退款单号"
          prop="refundNo"
          min-width="200"
        />
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
          label="类型"
          prop="type"
          width="110"
        >
          <template #default="scope">
            <el-tag :type="scope.row.type === 1 ? 'warning' : 'danger'">
              {{ scope.row.type === 1 ? '仅退款' : '退货退款' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="right"
          label="退款金额"
          width="120"
        >
          <template #default="scope">
            ¥{{ Number(scope.row.amount || 0).toFixed(2) }}
          </template>
        </el-table-column>
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
          align="left"
          label="原因"
          prop="reason"
          min-width="150"
          show-overflow-tooltip
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
          min-width="260"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              class="table-button"
              @click="viewDetail(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              详情
            </el-button>
            <el-button
              v-if="scope.row.status === 0"
              type="primary"
              link
              class="table-button"
              @click="openAuditDialog(scope.row)"
            >审核</el-button>
            <el-button
              v-if="scope.row.status === 1 && scope.row.type === 2"
              type="success"
              link
              class="table-button"
              @click="confirmReturnFunc(scope.row)"
            >确认收货</el-button>
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

    <!-- 退款详情对话框 -->
    <el-dialog
      v-model="detailVisible"
      title="退款详情"
      width="600px"
      destroy-on-close
    >
      <el-descriptions :column="2" border>
        <el-descriptions-item label="退款单号">{{ detailData.refundNo }}</el-descriptions-item>
        <el-descriptions-item label="订单号">{{ detailData.orderNo }}</el-descriptions-item>
        <el-descriptions-item label="用户ID">{{ detailData.userId }}</el-descriptions-item>
        <el-descriptions-item label="类型">
          <el-tag :type="detailData.type === 1 ? 'warning' : 'danger'">
            {{ detailData.type === 1 ? '仅退款' : '退货退款' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="退款金额">
          ¥{{ Number(detailData.amount || 0).toFixed(2) }}
        </el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="statusTagType(detailData.status)">{{ statusLabel(detailData.status) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="退款原因" :span="2">{{ detailData.reason || '-' }}</el-descriptions-item>
        <el-descriptions-item label="管理员备注" :span="2">{{ detailData.adminRemark || '-' }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ formatDate(detailData.CreatedAt) }}</el-descriptions-item>
        <el-descriptions-item label="更新时间">{{ formatDate(detailData.UpdatedAt) }}</el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="detailVisible = false">关 闭</el-button>
      </template>
    </el-dialog>

    <!-- 审核对话框 -->
    <el-dialog
      v-model="auditDialogVisible"
      title="退款审核"
      width="500px"
      destroy-on-close
    >
      <el-form
        ref="auditFormRef"
        :model="auditForm"
        :rules="auditRules"
        label-width="100px"
      >
        <el-form-item label="退款单号">
          <el-input :model-value="auditForm.refundNo" disabled />
        </el-form-item>
        <el-form-item label="退款金额">
          <el-input :model-value="'¥' + Number(auditForm.amount || 0).toFixed(2)" disabled />
        </el-form-item>
        <el-form-item label="审核结果" prop="status">
          <el-radio-group v-model="auditForm.status">
            <el-radio :value="1">同意</el-radio>
            <el-radio :value="2">拒绝</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="管理员备注" prop="adminRemark">
          <el-input
            v-model="auditForm.adminRemark"
            type="textarea"
            :rows="3"
            placeholder="请输入备注"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="auditDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="confirmAudit">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getRefundList,
  getRefundDetail,
  auditRefund,
  confirmReturn
} from '@/api/shopRefund'

import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { InfoFilled } from '@element-plus/icons-vue'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'ShopRefund'
})

// =========== 状态映射 ===========
const statusOptions = [
  { label: '待审核', value: 0 },
  { label: '已同意', value: 1 },
  { label: '已拒绝', value: 2 },
  { label: '退款中', value: 3 },
  { label: '已退款', value: 4 },
  { label: '已取消', value: 5 },
]

const statusTagType = (status) => {
  const map = { 0: 'warning', 1: 'primary', 2: 'danger', 3: '', 4: 'success', 5: 'info' }
  return map[status] || 'info'
}

const statusLabel = (status) => {
  const map = { 0: '待审核', 1: '已同意', 2: '已拒绝', 3: '退款中', 4: '已退款', 5: '已取消' }
  return map[status] !== undefined ? map[status] : '未知'
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
  const table = await getRefundList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 详情 ===============
const detailVisible = ref(false)
const detailData = ref({})

const viewDetail = async (row) => {
  const res = await getRefundDetail({ ID: row.ID })
  if (res.code === 0) {
    detailData.value = res.data.data
    detailVisible.value = true
  }
}

// ============== 审核 ===============
const auditDialogVisible = ref(false)
const auditFormRef = ref()
const auditForm = ref({
  ID: 0,
  refundNo: '',
  amount: 0,
  status: 1,
  adminRemark: '',
})

const auditRules = reactive({
  status: [
    { required: true, message: '请选择审核结果', trigger: 'change' },
  ],
})

const openAuditDialog = (row) => {
  auditForm.value = {
    ID: row.ID,
    refundNo: row.refundNo,
    amount: row.amount,
    status: 1,
    adminRemark: '',
  }
  auditDialogVisible.value = true
}

const confirmAudit = async () => {
  auditFormRef.value?.validate(async (valid) => {
    if (!valid) return
    const res = await auditRefund({
      ID: auditForm.value.ID,
      status: auditForm.value.status,
      adminRemark: auditForm.value.adminRemark,
    })
    if (res.code === 0) {
      ElMessage({ type: 'success', message: '审核成功' })
      auditDialogVisible.value = false
      getTableData()
    }
  })
}

// ============== 确认收货 ===============
const confirmReturnFunc = (row) => {
  ElMessageBox.confirm('确定已收到退货商品吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await confirmReturn({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({ type: 'success', message: '已确认收货' })
      getTableData()
    }
  })
}
</script>

<style scoped>
</style>
