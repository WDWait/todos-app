import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
    server: {
        proxy: {
            '/api': {
                target: 'http://localhost:8080', // Go 后端地址
                changeOrigin: true,
            }
        }
    }
})
