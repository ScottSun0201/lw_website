module.exports = {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./app.vue",
    "./error.vue",
  ],
  theme: {
    extend: {},
  },
  important: true,
  corePlugins: {
    preflight: false
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}