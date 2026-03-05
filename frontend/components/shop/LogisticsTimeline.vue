<template>
  <div v-if="traces.length">
    <el-timeline>
      <el-timeline-item v-for="(trace, idx) in traces" :key="idx" :timestamp="formatTime(trace.traceTime)" placement="top" :type="idx === 0 ? 'primary' : ''">
        {{ trace.detail }}
      </el-timeline-item>
    </el-timeline>
  </div>
  <div v-else class="py-4 text-center text-sm text-gray-400">暂无物流信息</div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { getTraces } from '~/api/logistics.js'
import { formatTimeToStr } from '~/utils/date.js'

const props = defineProps({ orderNo: { type: String, required: true } })
const traces = ref([])
const formatTime = (t) => t ? formatTimeToStr(t, 'yyyy-MM-dd hh:mm:ss') : ''

onMounted(async () => {
  const res = await getTraces({ orderNo: props.orderNo })
  if (res && res.code === 0) { traces.value = res.data.list || [] }
})
</script>
