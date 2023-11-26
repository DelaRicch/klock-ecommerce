import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import * as path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
    build: {
        target: 'es2020',
    },
    resolve: {
        alias: [{find: '@', replacement: path.resolve(__dirname, "src")}]
    }
})
