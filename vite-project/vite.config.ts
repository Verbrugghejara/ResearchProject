import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { VitePWA } from 'vite-plugin-pwa'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(),
    VitePWA(),],
  server:{
    host: '0.0.0.0',
    port:80,
    headers:{
      'Access-Control-Allow-Origin': '*',
    },
    cors:true,

  }
})
