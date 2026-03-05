<template>
  <div>
    <!-- 概览统计卡片 -->
    <el-row :gutter="20" style="margin-bottom: 20px;">
      <el-col :span="4">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-title">今日订单</div>
          <div class="stat-value">{{ overview.todayOrders ?? '-' }}</div>
        </el-card>
      </el-col>
      <el-col :span="4">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-title">今日收入</div>
          <div class="stat-value">{{ overview.todayRevenue != null ? '¥' + Number(overview.todayRevenue).toFixed(2) : '-' }}</div>
        </el-card>
      </el-col>
      <el-col :span="4">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-title">待处理订单</div>
          <div class="stat-value" style="color: #E6A23C;">{{ overview.pendingOrders ?? '-' }}</div>
        </el-card>
      </el-col>
      <el-col :span="4">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-title">待处理退款</div>
          <div class="stat-value" style="color: #F56C6C;">{{ overview.pendingRefunds ?? '-' }}</div>
        </el-card>
      </el-col>
      <el-col :span="4">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-title">商品总数</div>
          <div class="stat-value">{{ overview.totalProducts ?? '-' }}</div>
        </el-card>
      </el-col>
      <el-col :span="4">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-title">低库存数</div>
          <div class="stat-value" style="color: #F56C6C;">{{ overview.lowStockCount ?? '-' }}</div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 日期范围筛选 -->
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="日期范围">
          <el-date-picker
            v-model="searchInfo.dateRange"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD"
            style="width: 360px"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="loadAllReports">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 数据报表区域 -->
    <el-row :gutter="20">
      <!-- 销售趋势 -->
      <el-col :span="12">
        <el-card style="margin-bottom: 20px;">
          <template #header>
            <span>销售趋势</span>
          </template>
          <!-- 图表占位，后续接入 echarts -->
          <div class="chart-placeholder" ref="salesChartRef">
            <el-table :data="salesReport" border size="small" style="width: 100%">
              <el-table-column align="left" label="日期" prop="date" min-width="120" />
              <el-table-column align="right" label="订单数" prop="orderCount" width="100" />
              <el-table-column align="right" label="销售额" width="140">
                <template #default="scope">
                  ¥{{ Number(scope.row.salesAmount || 0).toFixed(2) }}
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-card>
      </el-col>

      <!-- 热销商品 -->
      <el-col :span="12">
        <el-card style="margin-bottom: 20px;">
          <template #header>
            <span>热销商品 TOP10</span>
          </template>
          <div class="chart-placeholder" ref="topProductsChartRef">
            <el-table :data="topProducts" border size="small" style="width: 100%">
              <el-table-column align="center" label="排名" width="70">
                <template #default="scope">
                  <el-tag
                    :type="scope.$index < 3 ? 'danger' : 'info'"
                    size="small"
                  >{{ scope.$index + 1 }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column align="left" label="商品名称" prop="productName" min-width="180" show-overflow-tooltip />
              <el-table-column align="right" label="销量" prop="salesCount" width="100" />
              <el-table-column align="right" label="销售额" width="140">
                <template #default="scope">
                  ¥{{ Number(scope.row.salesAmount || 0).toFixed(2) }}
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20">
      <!-- 分类销售分布 -->
      <el-col :span="12">
        <el-card style="margin-bottom: 20px;">
          <template #header>
            <span>分类销售分布</span>
          </template>
          <div class="chart-placeholder" ref="categoryChartRef">
            <el-table :data="categorySales" border size="small" style="width: 100%">
              <el-table-column align="left" label="分类名称" prop="categoryName" min-width="150" />
              <el-table-column align="right" label="订单数" prop="orderCount" width="100" />
              <el-table-column align="right" label="销售额" width="140">
                <template #default="scope">
                  ¥{{ Number(scope.row.salesAmount || 0).toFixed(2) }}
                </template>
              </el-table-column>
              <el-table-column align="right" label="占比" width="100">
                <template #default="scope">
                  {{ scope.row.percentage ? scope.row.percentage + '%' : '-' }}
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-card>
      </el-col>

      <!-- 订单状态分布 -->
      <el-col :span="12">
        <el-card style="margin-bottom: 20px;">
          <template #header>
            <span>订单状态分布</span>
          </template>
          <div class="chart-placeholder" ref="orderStatusChartRef">
            <el-table :data="orderStatusData" border size="small" style="width: 100%">
              <el-table-column align="left" label="状态" width="120">
                <template #default="scope">
                  <el-tag :type="orderStatusTagType(scope.row.status)" size="small">
                    {{ orderStatusLabel(scope.row.status) }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column align="right" label="数量" prop="count" min-width="100" />
              <el-table-column align="right" label="占比" width="100">
                <template #default="scope">
                  {{ scope.row.percentage ? scope.row.percentage + '%' : '-' }}
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import {
  getDashboardOverview,
  getSalesReport,
  getTopProducts,
  getCategorySales,
  getOrderStatusDistribution
} from '@/api/shopAnalytics'

import { formatDate } from '@/utils/format'
import { ref, onMounted } from 'vue'

defineOptions({
  name: 'ShopAnalytics'
})

// =========== 订单状态映射 ===========
const orderStatusTagType = (status) => {
  const map = { 0: 'warning', 1: 'primary', 2: '', 3: 'success', 4: 'success', 5: 'info' }
  return map[status] || 'info'
}

const orderStatusLabel = (status) => {
  const map = { 0: '待付款', 1: '已付款', 2: '已发货', 3: '已收货', 4: '已完成', 5: '已取消' }
  return map[status] !== undefined ? map[status] : '未知'
}

// =========== 概览数据 ===========
const overview = ref({
  todayOrders: 0,
  todayRevenue: 0,
  pendingOrders: 0,
  pendingRefunds: 0,
  totalProducts: 0,
  lowStockCount: 0,
})

const loadOverview = async () => {
  const res = await getDashboardOverview()
  if (res.code === 0) {
    overview.value = res.data.data || {}
  }
}

// =========== 筛选 ===========
const searchInfo = ref({
  dateRange: [],
})

const getDateParams = () => {
  const params = {}
  if (searchInfo.value.dateRange && searchInfo.value.dateRange.length === 2) {
    params.startDate = searchInfo.value.dateRange[0]
    params.endDate = searchInfo.value.dateRange[1]
  }
  return params
}

const onReset = () => {
  searchInfo.value = { dateRange: [] }
  loadAllReports()
}

// =========== 报表数据 ===========
const salesReport = ref([])
const topProducts = ref([])
const categorySales = ref([])
const orderStatusData = ref([])

const loadSalesReport = async () => {
  const res = await getSalesReport(getDateParams())
  if (res.code === 0) {
    salesReport.value = res.data.list || []
  }
}

const loadTopProducts = async () => {
  const res = await getTopProducts({ ...getDateParams(), limit: 10 })
  if (res.code === 0) {
    topProducts.value = res.data.list || []
  }
}

const loadCategorySales = async () => {
  const res = await getCategorySales(getDateParams())
  if (res.code === 0) {
    categorySales.value = res.data.list || []
  }
}

const loadOrderStatus = async () => {
  const res = await getOrderStatusDistribution(getDateParams())
  if (res.code === 0) {
    orderStatusData.value = res.data.list || []
  }
}

const loadAllReports = () => {
  loadSalesReport()
  loadTopProducts()
  loadCategorySales()
  loadOrderStatus()
}

// =========== 图表 refs (预留 echarts) ===========
const salesChartRef = ref()
const topProductsChartRef = ref()
const categoryChartRef = ref()
const orderStatusChartRef = ref()

onMounted(() => {
  loadOverview()
  loadAllReports()
})
</script>

<style scoped>
.stat-card {
  text-align: center;
}

.stat-title {
  font-size: 14px;
  color: #909399;
  margin-bottom: 8px;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
}

.chart-placeholder {
  min-height: 200px;
}
</style>
