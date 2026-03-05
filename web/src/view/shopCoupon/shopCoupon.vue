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
        <el-form-item label="名称" prop="name">
          <el-input
            v-model="searchInfo.name"
            placeholder="请输入优惠券名称"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select
            v-model="searchInfo.status"
            placeholder="请选择状态"
            clearable
            style="width: 140px"
          >
            <el-option label="禁用" :value="0" />
            <el-option label="启用" :value="1" />
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
          @click="openDialog"
        >新增</el-button>
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
          align="left"
          label="名称"
          prop="name"
          min-width="150"
        />
        <el-table-column
          align="left"
          label="兑换码"
          prop="code"
          width="140"
        />
        <el-table-column
          align="center"
          label="类型"
          prop="type"
          width="100"
        >
          <template #default="scope">
            <el-tag :type="couponTypeTagType(scope.row.type)">
              {{ couponTypeLabel(scope.row.type) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="right"
          label="面值"
          prop="value"
          width="100"
        >
          <template #default="scope">
            <span v-if="scope.row.type === 2">{{ scope.row.value }}折</span>
            <span v-else>¥{{ Number(scope.row.value || 0).toFixed(2) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="right"
          label="最低金额"
          width="100"
        >
          <template #default="scope">
            ¥{{ Number(scope.row.minAmount || 0).toFixed(2) }}
          </template>
        </el-table-column>
        <el-table-column
          align="right"
          label="最高抵扣"
          width="100"
        >
          <template #default="scope">
            {{ scope.row.maxDiscount ? '¥' + Number(scope.row.maxDiscount).toFixed(2) : '-' }}
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="总量"
          prop="totalCount"
          width="80"
        />
        <el-table-column
          align="center"
          label="已用"
          prop="usedCount"
          width="80"
        />
        <el-table-column
          align="center"
          label="每人限领"
          prop="perLimit"
          width="90"
        />
        <el-table-column
          align="left"
          label="开始时间"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.startTime) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="结束时间"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.endTime) }}</template>
        </el-table-column>
        <el-table-column
          align="center"
          label="状态"
          prop="status"
          width="80"
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
              @click="updateCouponFunc(scope.row)"
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
      :title="type === 'create' ? '新增优惠券' : '编辑优惠券'"
      width="600px"
      destroy-on-close
    >
      <el-form
        ref="elFormRef"
        :model="formData"
        :rules="rule"
        label-width="100px"
      >
        <el-form-item label="名称" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="请输入优惠券名称"
            clearable
          />
        </el-form-item>
        <el-form-item label="兑换码" prop="code">
          <el-input
            v-model="formData.code"
            placeholder="请输入兑换码"
            clearable
          />
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select
            v-model="formData.type"
            placeholder="请选择类型"
            style="width: 100%"
          >
            <el-option label="满减" :value="1" />
            <el-option label="折扣" :value="2" />
            <el-option label="固定" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="面值" prop="value">
          <el-input-number
            v-model="formData.value"
            :min="0"
            :precision="2"
            controls-position="right"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="最低金额" prop="minAmount">
          <el-input-number
            v-model="formData.minAmount"
            :min="0"
            :precision="2"
            controls-position="right"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="最高抵扣" prop="maxDiscount">
          <el-input-number
            v-model="formData.maxDiscount"
            :min="0"
            :precision="2"
            controls-position="right"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="发放总量" prop="totalCount">
          <el-input-number
            v-model="formData.totalCount"
            :min="0"
            controls-position="right"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="每人限领" prop="perLimit">
          <el-input-number
            v-model="formData.perLimit"
            :min="0"
            controls-position="right"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="开始时间" prop="startTime">
          <el-date-picker
            v-model="formData.startTime"
            type="datetime"
            placeholder="请选择开始时间"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="结束时间" prop="endTime">
          <el-date-picker
            v-model="formData.endTime"
            type="datetime"
            placeholder="请选择结束时间"
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
  createShopCoupon,
  deleteShopCoupon,
  updateShopCoupon,
  findShopCoupon,
  getShopCouponList
} from '@/api/shopCoupon'

import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'ShopCoupon'
})

// =========== 类型映射 ===========
const couponTypeTagType = (type) => {
  const map = { 1: 'danger', 2: 'warning', 3: '' }
  return map[type] || 'info'
}

const couponTypeLabel = (type) => {
  const map = { 1: '满减', 2: '折扣', 3: '固定' }
  return map[type] || '未知'
}

// 初始表单数据
const initFormData = () => ({
  name: '',
  code: '',
  type: 1,
  value: 0,
  minAmount: 0,
  maxDiscount: 0,
  totalCount: 0,
  perLimit: 1,
  startTime: '',
  endTime: '',
  status: 1,
})

const formData = ref(initFormData())

// 验证规则
const rule = reactive({
  name: [
    { required: true, message: '请输入优惠券名称', trigger: ['input', 'blur'] },
  ],
  type: [
    { required: true, message: '请选择类型', trigger: 'change' },
  ],
  value: [
    { required: true, message: '请输入面值', trigger: ['input', 'blur'] },
  ],
})

const elFormRef = ref()

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
  const table = await getShopCouponList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const type = ref('')

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
        res = await createShopCoupon(formData.value)
        break
      case 'update':
        res = await updateShopCoupon(formData.value)
        break
      default:
        res = await createShopCoupon(formData.value)
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

// 更新行
const updateCouponFunc = async (row) => {
  const res = await findShopCoupon({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.data
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除该优惠券吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteShopCoupon({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({ type: 'success', message: '删除成功' })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}
</script>

<style scoped>
</style>
