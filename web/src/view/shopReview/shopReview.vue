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
        <el-form-item label="评分" prop="rating">
          <el-select
            v-model="searchInfo.rating"
            placeholder="请选择评分"
            clearable
            style="width: 140px"
          >
            <el-option
              v-for="item in ratingOptions"
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
          label="用户ID"
          prop="userId"
          width="100"
        />
        <el-table-column
          align="center"
          label="商品ID"
          prop="spuId"
          width="100"
        />
        <el-table-column
          align="center"
          label="评分"
          width="160"
        >
          <template #default="scope">
            <el-rate
              :model-value="scope.row.rating"
              disabled
              show-score
              text-color="#ff9900"
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="评价内容"
          prop="content"
          min-width="200"
          show-overflow-tooltip
        />
        <el-table-column
          align="center"
          label="图片"
          width="120"
        >
          <template #default="scope">
            <template v-if="scope.row.images && scope.row.images.length > 0">
              <el-image
                style="width: 40px; height: 40px"
                :src="scope.row.images[0]"
                fit="cover"
                :preview-src-list="scope.row.images"
                preview-teleported
              />
              <span v-if="scope.row.images.length > 1" style="margin-left: 4px; color: #999;">
                +{{ scope.row.images.length - 1 }}
              </span>
            </template>
            <span v-else>-</span>
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
          label="回复"
          prop="reply"
          min-width="150"
          show-overflow-tooltip
        >
          <template #default="scope">
            {{ scope.row.reply || '-' }}
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
              v-if="scope.row.status === 0"
              type="primary"
              link
              class="table-button"
              @click="openAuditDialog(scope.row)"
            >审核</el-button>
            <el-button
              type="primary"
              link
              class="table-button"
              @click="openReplyDialog(scope.row)"
            >回复</el-button>
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

    <!-- 审核对话框 -->
    <el-dialog
      v-model="auditDialogVisible"
      title="评价审核"
      width="500px"
      destroy-on-close
    >
      <div style="margin-bottom: 16px;">
        <p><strong>评价内容：</strong>{{ auditRow.content }}</p>
        <p style="margin-top: 8px;">
          <strong>评分：</strong>
          <el-rate :model-value="auditRow.rating" disabled />
        </p>
      </div>
      <el-form label-width="80px">
        <el-form-item label="审核结果">
          <el-radio-group v-model="auditForm.status">
            <el-radio :value="1">通过</el-radio>
            <el-radio :value="2">拒绝</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="auditDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="confirmAudit">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 回复对话框 -->
    <el-dialog
      v-model="replyDialogVisible"
      title="回复评价"
      width="500px"
      destroy-on-close
    >
      <div style="margin-bottom: 16px;">
        <p><strong>评价内容：</strong>{{ replyRow.content }}</p>
      </div>
      <el-form label-width="80px">
        <el-form-item label="回复内容">
          <el-input
            v-model="replyForm.reply"
            type="textarea"
            :rows="4"
            placeholder="请输入回复内容"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="replyDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="confirmReply">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getReviewList,
  auditReview,
  replyReview
} from '@/api/shopReview'

import { formatDate } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'

defineOptions({
  name: 'ShopReview'
})

// =========== 状态映射 ===========
const statusOptions = [
  { label: '待审核', value: 0 },
  { label: '已通过', value: 1 },
  { label: '已拒绝', value: 2 },
]

const ratingOptions = [
  { label: '1星', value: 1 },
  { label: '2星', value: 2 },
  { label: '3星', value: 3 },
  { label: '4星', value: 4 },
  { label: '5星', value: 5 },
]

const statusTagType = (status) => {
  const map = { 0: 'warning', 1: 'success', 2: 'danger' }
  return map[status] || 'info'
}

const statusLabel = (status) => {
  const map = { 0: '待审核', 1: '已通过', 2: '已拒绝' }
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
  const table = await getReviewList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 审核 ===============
const auditDialogVisible = ref(false)
const auditRow = ref({})
const auditForm = ref({
  ID: 0,
  status: 1,
})

const openAuditDialog = (row) => {
  auditRow.value = row
  auditForm.value = {
    ID: row.ID,
    status: 1,
  }
  auditDialogVisible.value = true
}

const confirmAudit = async () => {
  const res = await auditReview(auditForm.value)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '审核成功' })
    auditDialogVisible.value = false
    getTableData()
  }
}

// ============== 回复 ===============
const replyDialogVisible = ref(false)
const replyRow = ref({})
const replyForm = ref({
  ID: 0,
  reply: '',
})

const openReplyDialog = (row) => {
  replyRow.value = row
  replyForm.value = {
    ID: row.ID,
    reply: row.reply || '',
  }
  replyDialogVisible.value = true
}

const confirmReply = async () => {
  if (!replyForm.value.reply.trim()) {
    ElMessage({ type: 'warning', message: '请输入回复内容' })
    return
  }
  const res = await replyReview(replyForm.value)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '回复成功' })
    replyDialogVisible.value = false
    getTableData()
  }
}
</script>

<style scoped>
</style>
