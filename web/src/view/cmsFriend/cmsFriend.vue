<template>
  <div>
    <div class="gva-search-box">
      <h3>
        友情链接顶部内容配置
      </h3>
      <el-form
        ref="elSearchFormRef"
        :model="friendInfo"
        class="demo-form-inline"
        :rules="friendInfoRule"
        label-position="top"
      >
        <el-form-item
          label="友情链接Title"
        >
          <el-input
            v-model="friendInfo.title"
            placeholder="请输入友情链接Title"
          />
        </el-form-item>

        <el-form-item
          label="友情链接Desc"
        >
          <el-input
            v-model="friendInfo.desc"
            placeholder="友情链接Desc"
          />
        </el-form-item>

        <el-form-item
          label="友情链接二级Title"
        >
          <el-input
            v-model="friendInfo.friendTitle"
            placeholder="请输入友情链接Title"
          />
        </el-form-item>

        <el-form-item
          label="友情链接二级Desc"
        >
          <el-input
            v-model="friendInfo.friendDesc"
            placeholder="友情链接Desc"
          />
        </el-form-item>

        <el-form-item
          label="友情链接跳转"
        >
          <div>
            <div class="mb-2">
              <el-button
                type="primary"
                @click="addLinks"
              >添加</el-button>
            </div>
            <el-row
              v-for="(item, index) in friendInfo.links"
              :key="index"
              class="flex items-center"
            >
              <el-col
                :span="10"
              >
                <el-input
                  v-model="item.name"
                  placeholder="请输入名称"
                />
              </el-col>
              <el-col
                :span="10"
              >
                <el-input
                  v-model="item.href"
                  placeholder="请输入链接"
                />
              </el-col>
              <el-col
                :span="4"
              >
                <el-button
                  type="danger"
                  class="ml-2"
                  @click="cutLinks(index)"
                >删除</el-button>
              </el-col>
            </el-row>
          </div>
        </el-form-item>

        <el-form-item
          label="友情链接信息条"
        >
          <div>
            <div class="mb-2">
              <el-button
                type="primary"
                @click="addStats"
              >添加</el-button>
            </div>
            <el-row
              v-for="(item, index) in friendInfo.stats"
              :key="index"
              class="flex items-center"
            >
              <el-col
                :span="10"
              >
                <el-input
                  v-model="item.name"
                  placeholder="请输入名称"
                />
              </el-col>
              <el-col
                :span="10"
              >
                <el-input
                  v-model="item.value"
                  placeholder="请输入值"
                />
              </el-col>
              <el-col
                :span="4"
              >
                <el-button
                  type="danger"
                  class="ml-2"
                  @click="cutStats(index)"
                >删除</el-button>
              </el-col>
            </el-row>
          </div>
        </el-form-item>
        <el-form-item>
          <div
            class="flex justify-end w-full"
          >
            <el-button
              type="primary"
              @click="saveInfo"
            >保存</el-button>
          </div>
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
        <el-popover
          v-model:visible="deleteVisible"
          :disabled="!multipleSelection.length"
          placement="top"
          width="160"
        >
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button
              type="primary"
              link
              @click="deleteVisible = false"
            >取消</el-button>
            <el-button
              type="primary"
              @click="onDelete"
            >确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />

        <el-table-column
          align="left"
          label="日期"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          label="标题"
          prop="title"
          width="120"
        />
        <el-table-column
          label="友链图片"
          width="200"
        >
          <template #default="scope">
            <el-image
              style="width: 100px; height: 100px"
              :src="getUrl(scope.row.image)"
              fit="cover"
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="链接"
          prop="link"
          width="120"
        />
        <el-table-column
          align="left"
          label="新窗口打开"
          prop="newWindow"
          width="120"
        >
          <template #default="scope">{{ formatBoolean(scope.row.newWindow) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="启用"
          prop="open"
          width="120"
        >
          <template #default="scope">{{ formatBoolean(scope.row.open) }}</template>
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
              @click="getDetails(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              查看详情
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateCmsFriendFunc(scope.row)"
            >变更</el-button>
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
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="type==='create'?'添加':'修改'"
      destroy-on-close
    >
      <el-scrollbar height="500px">
        <el-form
          ref="elFormRef"
          :model="formData"
          label-position="right"
          :rules="rule"
          label-width="120px"
        >
          <el-form-item
            label="标题:"
            prop="title"
          >
            <el-input
              v-model="formData.title"
              :clearable="true"
              placeholder="请输入标题"
            />
          </el-form-item>
          <el-form-item
            label="友链图片:"
            prop="image"
          >
            <SelectImage
              v-model="formData.image"
              file-type="image"
            />
          </el-form-item>
          <el-form-item
            label="链接:"
            prop="link"
          >
            <el-input
              v-model="formData.link"
              :clearable="true"
              placeholder="请输入链接"
            />
          </el-form-item>
          <el-form-item
            label="新窗口打开:"
            prop="newWindow"
          >
            <el-switch
              v-model="formData.newWindow"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
              clearable
            />
          </el-form-item>
          <el-form-item
            label="启用:"
            prop="open"
          >
            <el-switch
              v-model="formData.open"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
              clearable
            />
          </el-form-item>
        </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button
            type="primary"
            @click="enterDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="detailShow"
      style="width: 800px"
      lock-scroll
      :before-close="closeDetailShow"
      title="查看详情"
      destroy-on-close
    >
      <el-scrollbar height="550px">
        <el-descriptions
          column="1"
          border
        >
          <el-descriptions-item label="标题">
            {{ formData.title }}
          </el-descriptions-item>
          <el-descriptions-item label="友链图片">
            <el-image
              style="width: 50px; height: 50px"
              :preview-src-list="ReturnArrImg(formData.image)"
              :src="getUrl(formData.image)"
              fit="cover"
            />
          </el-descriptions-item>
          <el-descriptions-item label="链接">
            {{ formData.link }}
          </el-descriptions-item>
          <el-descriptions-item label="新窗口打开">
            {{ formatBoolean(formData.newWindow) }}
          </el-descriptions-item>
          <el-descriptions-item label="启用">
            {{ formatBoolean(formData.open) }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createCmsFriend,
  deleteCmsFriend,
  deleteCmsFriendByIds,
  updateCmsFriend,
  findCmsFriend,
  getCmsFriendList, setCmsFriendInfo, getCmsFriendInfo
} from '@/api/cmsFriend'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'CmsFriend'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  title: '',
  image: '',
  link: '',
  newWindow: false,
  open: false,
})

// 验证规则
const rule = reactive({
})

const friendInfo = ref({
  title: '',
  desc: '',
  friendTitle: '',
  friendDesc: '',
  links: [
    { name: 'Open roles', href: '#' },
    { name: 'Internship program', href: '#' },
    { name: 'Our values', href: '#' },
    { name: 'Meet our leadership', href: '#' },
  ],
  stats: [
    { name: 'Offices worldwide', value: '12' },
    { name: 'Full-time colleagues', value: '300+' },
    { name: 'Hours per week', value: '40' },
    { name: 'Paid time off', value: 'Unlimited' },
  ]
})

const addLinks = () => {
  friendInfo.value.links.push({ name: '', href: '' })
}
const cutLinks = (index) => {
  friendInfo.value.links.splice(index, 1)
}

const addStats = () => {
  friendInfo.value.stats.push({ name: '', value: '' })
}

const cutStats = (index) => {
  friendInfo.value.stats.splice(index, 1)
}

const saveInfo = async() => {
  const res = await setCmsFriendInfo({ friendInfo: friendInfo.value, ID: 1 })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '保存成功'
    })
  }
}

const getInfo = async() => {
  const res = await getCmsFriendInfo()
  if (res.code === 0) {
    friendInfo.value = res.data.friendInfo
  }
}

getInfo()

const friendInfoRule = reactive({
  title: [
    { required: false, message: '请输入友情链接Title', trigger: 'blur' }
  ],
  desc: [
    { required: false, message: '请输入友情链接Desc', trigger: 'blur' }
  ]
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getCmsFriendList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteCmsFriendFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const IDs = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
  const res = await deleteCmsFriendByIds({ IDs })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === IDs.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCmsFriendFunc = async(row) => {
  const res = await findCmsFriend({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.recmsFriend
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteCmsFriendFunc = async(row) => {
  const res = await deleteCmsFriend({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async(row) => {
  // 打开弹窗
  const res = await findCmsFriend({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recmsFriend
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    title: '',
    link: '',
    newWindow: false,
    open: false,
  }
}

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    title: '',
    link: '',
    newWindow: false,
    open: false,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createCmsFriend(formData.value)
        break
      case 'update':
        res = await updateCmsFriend(formData.value)
        break
      default:
        res = await createCmsFriend(formData.value)
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

</script>

<style>

</style>
