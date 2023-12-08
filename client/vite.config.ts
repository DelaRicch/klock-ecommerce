import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import * as path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
    server: {
      host: true,
      port: 5173,
      watch: {
        usePolling: true
      }  
    },
    build: {
        target: 'es2020',
    },
    resolve: {
        alias: [{find: '@', replacement: path.resolve(__dirname, "src")}]
    }
})
