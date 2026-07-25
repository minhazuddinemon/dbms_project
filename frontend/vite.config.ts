import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig, loadEnv } from 'vite';

export default defineConfig(({ mode }) => {
	const env = loadEnv(mode, process.cwd(), '');
	const backendUrl = env.VITE_BACKEND_URL || process.env.VITE_BACKEND_URL || 'http://localhost:8080';

	return {
		server: {
			port: 3000,
			host: true,
			proxy: {
				'/api': {
					target: backendUrl,
					changeOrigin: true,
					rewrite: (path) => path.replace(/^\/api/, '')
				}
			}
		},
		preview: {
			port: 3000,
			host: true,
			proxy: {
				'/api': {
					target: backendUrl,
					changeOrigin: true,
					rewrite: (path) => path.replace(/^\/api/, '')
				}
			}
		},
		plugins: [
			tailwindcss(),
			sveltekit()
		]
	};
});
