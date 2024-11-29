import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173, // Frontend runs on port 5173
    proxy: {
      '/auth': {
        target: 'http://localhost:8000', // Backend API server
        changeOrigin: true, // Ensures the host header is updated to match the target
      },
    },
  },
})
