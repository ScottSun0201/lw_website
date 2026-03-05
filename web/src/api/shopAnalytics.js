import service from '@/utils/request'

export const getDashboardOverview = (params) => {
  return service({ url: '/shopAnalytics/getDashboardOverview', method: 'get', params })
}

export const getSalesReport = (params) => {
  return service({ url: '/shopAnalytics/getSalesReport', method: 'get', params })
}

export const getTopProducts = (params) => {
  return service({ url: '/shopAnalytics/getTopProducts', method: 'get', params })
}

export const getCategorySales = (params) => {
  return service({ url: '/shopAnalytics/getCategorySales', method: 'get', params })
}

export const getOrderStatusDistribution = (params) => {
  return service({ url: '/shopAnalytics/getOrderStatusDistribution', method: 'get', params })
}
