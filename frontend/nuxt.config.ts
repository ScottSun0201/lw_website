// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
    runtimeConfig: {
        // public 命名空间中定义的，在服务器端和客户端都可以普遍访问
        public: {
            baseURL: process.env.NUXT_BASE_URL,
            basePROXY: process.env.NUXT_BASE_PROXY
        }
    },
    css: ['~/assets/css/main.scss'],
    modules: [
      '@element-plus/nuxt',
      '@pinia/nuxt',
  ],
    postcss: {
        plugins: {
            tailwindcss: {},
            autoprefixer: {},
        },
    },
})
