import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [svelte()],
	optimizeDeps: { exclude: ['svelte-navigator'] },
	resolve: {
		alias: {
			src: path.resolve(__dirname, 'src'),
			'@': path.resolve(__dirname, 'src'),
		},
	},
})
