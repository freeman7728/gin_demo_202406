import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { fileURLToPath, URL } from 'url'; 

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      // 在此处编写代理规则
      '/api': {
          target: 'http://localhost:4000',
          changeOrigin: true,
          rewrite: (path) => {
              return path.replace(/\/api/, '')
          }
      }
  }
  
}
});
