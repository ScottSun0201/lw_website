<template>
  <footer class="bg-gray-900 box-border" aria-labelledby="footer-heading">
    <h2 id="footer-heading" class="sr-only">页脚</h2>
    <div class="mx-auto max-w-7xl px-6 pb-8 pt-20 sm:pt-24 lg:px-8 lg:pt-32">
      <div class="xl:grid xl:grid-cols-3 xl:gap-8">
        <div class="grid grid-cols-2 gap-8 xl:col-span-2">
          <div class="md:grid md:grid-cols-2 md:gap-8">
            <div>
              <h3 class="font-semibold leading-6 text-white">产品服务</h3>
              <ul role="list" class="mt-6 space-y-4">
                <li v-for="item in navigation.products" :key="item.name">
                  <a :href="item.href" class="leading-6 text-gray-300 hover:text-white">{{ item.name }}</a>
                </li>
              </ul>
            </div>
            <div class="mt-10 md:mt-0">
              <h3 class="font-semibold leading-6 text-white">帮助支持</h3>
              <ul role="list" class="mt-6 space-y-4">
                <li v-for="item in navigation.support" :key="item.name">
                  <a :href="item.href" class="leading-6 text-gray-300 hover:text-white">{{ item.name }}</a>
                </li>
              </ul>
            </div>
          </div>
          <div class="md:grid md:grid-cols-2 md:gap-8">
            <div>
              <h3 class="font-semibold leading-6 text-white">关于我们</h3>
              <ul role="list" class="mt-6 space-y-4">
                <li v-for="item in navigation.company" :key="item.name">
                  <a :href="item.href" class="leading-6 text-gray-300 hover:text-white">{{ item.name }}</a>
                </li>
              </ul>
            </div>
            <div class="mt-10 md:mt-0">
              <h3 class="font-semibold leading-6 text-white">法律条款</h3>
              <ul role="list" class="mt-6 space-y-4">
                <li v-for="item in navigation.legal" :key="item.name">
                  <a :href="item.href" class="leading-6 text-gray-300 hover:text-white">{{ item.name }}</a>
                </li>
              </ul>
            </div>
          </div>
        </div>
        <div class="mt-10 xl:mt-0">
          <h3 class="font-semibold leading-6 text-white">订阅我们的资讯</h3>
          <p class="mt-2 leading-6 text-gray-300">获取最新的新闻、文章和资源，每周发送到您的邮箱。</p>
          <form class="mt-6 sm:flex sm:max-w-md">
            <label for="email-address" class="sr-only">邮箱地址</label>
            <input type="email" name="email-address" id="email-address" autocomplete="email" required="" class="w-full min-w-0 appearance-none rounded-md border-white/10 bg-gray-400/10 px-[calc(theme(spacing.3)-1px)] py-[calc(theme(spacing[1.5])-1px)] text-base leading-7 text-white placeholder-gray-500 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:w-64 sm:sm:leading-6 xl:w-full" placeholder="输入您的邮箱" />
            <div class="mt-4 rounded-md sm:mt-0 sm:ml-4 sm:flex-shrink-0">
              <button type="submit" class="flex w-full items-center justify-center rounded-md bg-indigo-500 py-1.5 px-3 text-base font-semibold leading-7 text-white hover:bg-indigo-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-400 sm:sm:leading-6">订阅</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup>

const navigation = {
  products: [
    { name: '商城', href: '/shop' },
    { name: '博客', href: '/blog' },
  ],
  support: [
    { name: '帮助中心', href: '#' },
    { name: '联系我们', href: '#' },
  ],
  company: [
    { name: '关于', href: '/about' },
    { name: '博客', href: '/blog' },
  ],
  legal: [
    { name: '隐私政策', href: '#' },
    { name: '服务条款', href: '#' },
  ],
}

</script>
