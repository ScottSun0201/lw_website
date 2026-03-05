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
        <el-form-item label="商品名称" prop="name">
          <el-input
            v-model="searchInfo.name"
            placeholder="请输入商品名称"
          />
        </el-form-item>
        <el-form-item label="分类" prop="categoryId">
          <el-cascader
            v-model="searchInfo.categoryId"
            :options="categoryTree"
            :props="{
              value: 'ID',
              label: 'name',
              children: 'children',
              checkStrictly: true,
              emitPath: false
            }"
            placeholder="请选择分类"
            clearable
            style="width: 200px"
          />
        </el-form-item>
        <el-form-item label="品牌" prop="brandId">
          <el-select
            v-model="searchInfo.brandId"
            placeholder="请选择品牌"
            clearable
            style="width: 160px"
          >
            <el-option
              v-for="item in brandList"
              :key="item.ID"
              :label="item.name"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select
            v-model="searchInfo.status"
            placeholder="请选择状态"
            clearable
            style="width: 120px"
          >
            <el-option label="草稿" :value="0" />
            <el-option label="上架" :value="1" />
            <el-option label="下架" :value="2" />
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
          @click="toCreateProduct"
        >新增商品</el-button>
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
          label="主图"
          width="100"
        >
          <template #default="scope">
            <el-image
              v-if="scope.row.mainImage"
              style="width: 60px; height: 60px"
              :src="getUrl(scope.row.mainImage)"
              fit="cover"
              :preview-src-list="[getUrl(scope.row.mainImage)]"
              preview-teleported
            />
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="商品名称"
          prop="name"
          min-width="200"
          show-overflow-tooltip
        />
        <el-table-column
          align="center"
          label="分类"
          width="120"
        >
          <template #default="scope">
            {{ scope.row.categoryName || scope.row.categoryId || '-' }}
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="品牌"
          width="120"
        >
          <template #default="scope">
            {{ scope.row.brandName || scope.row.brandId || '-' }}
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="价格区间"
          width="140"
        >
          <template #default="scope">
            {{ scope.row.priceRange || '-' }}
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
          align="center"
          label="销量"
          prop="salesCount"
          width="100"
        />
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
              icon="edit"
              class="table-button"
              @click="toEditProduct(scope.row)"
            >编辑</el-button>
            <el-button
              v-if="scope.row.status !== 1"
              type="success"
              link
              class="table-button"
              @click="changeStatus(scope.row, 1)"
            >上架</el-button>
            <el-button
              v-if="scope.row.status === 1"
              type="warning"
              link
              class="table-button"
              @click="changeStatus(scope.row, 2)"
            >下架</el-button>
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
  </div>
</template>

<script setup>
import {
  getSpuList,
  deleteSpu,
  updateSpuStatus
} from '@/api/shopProduct'
import { getCategoryTree } from '@/api/shopCategory'
import { getAllBrands } from '@/api/shopBrand'
import { getUrl } from '@/utils/image'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

defineOptions({
  name: 'ShopProduct'
})

const router = useRouter()

// =========== 辅助数据 ===========
const categoryTree = ref([])
const brandList = ref([])

const loadCategoryTree = async () => {
  const res = await getCategoryTree()
  if (res.code === 0) {
    categoryTree.value = res.data.list || []
  }
}

const loadBrandList = async () => {
  const res = await getAllBrands()
  if (res.code === 0) {
    brandList.value = res.data.list || []
  }
}

loadCategoryTree()
loadBrandList()

// =========== 状态映射 ===========
const statusTagType = (status) => {
  const map = { 0: 'info', 1: 'success', 2: 'warning' }
  return map[status] || 'info'
}

const statusLabel = (status) => {
  const map = { 0: '草稿', 1: '上架', 2: '下架' }
  return map[status] || '未知'
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
  const table = await getSpuList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 跳转新增商品页面
const toCreateProduct = () => {
  router.push({ name: 'shopProductForm' })
}

// 跳转编辑商品页面
const toEditProduct = (row) => {
  router.push({ name: 'shopProductForm', query: { id: row.ID } })
}

// 上架/下架
const changeStatus = (row, status) => {
  const actionText = status === 1 ? '上架' : '下架'
  ElMessageBox.confirm(`确定要${actionText}该商品吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await updateSpuStatus({ ID: row.ID, status })
    if (res.code === 0) {
      ElMessage({ type: 'success', message: `${actionText}成功` })
      getTableData()
    }
  })
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除该商品吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteSpu({ ID: row.ID })
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
