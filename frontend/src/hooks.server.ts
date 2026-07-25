import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	if (event.url.pathname.startsWith('/api')) {
		const backendUrl = process.env.VITE_BACKEND_URL || 'http://localhost:8080';
		const targetPath = event.url.pathname.replace(/^\/api/, '');
		const targetUrl = `${backendUrl}${targetPath}${event.url.search}`;

		const reqHeaders = new Headers(event.request.headers);
		reqHeaders.delete('host');

		const options: RequestInit = {
			method: event.request.method,
			headers: reqHeaders
		};

		if (!['GET', 'HEAD'].includes(event.request.method)) {
			options.body = await event.request.arrayBuffer();
			// @ts-expect-error duplex is required for streaming request bodies in Node fetch
			options.duplex = 'half';
		}

		return fetch(targetUrl, options);
	}

	return resolve(event);
};
