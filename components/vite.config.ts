import { defineConfig } from 'vite'

export default defineConfig({
  plugins: [],
  build: {
    rollupOptions: {
      output: {
        entryFileNames: `assets/main.js`,
        assetFileNames: `assets/style.css`,
      }
    }
  }
})