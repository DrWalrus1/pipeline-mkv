import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import path from 'node:path'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
      '@/config': fileURLToPath(new URL('./src/config.ts', import.meta.url)),
    },
  },
  build: {
    outDir: path.resolve(__dirname, '../static/'),
    emptyOutDir: true,
  },
  server: {
    proxy: {
      '/api': {
        target: 'https://localhost:8080',
        secure: false,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
        ws: true
      },
    },
  },
})
