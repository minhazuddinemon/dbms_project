// src/lib/api/client.ts

const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

export function getAuthToken(): string | null {
	if (typeof window !== 'undefined') {
		return localStorage.getItem('uniapp_token');
	}
	return null;
}

export function setAuthToken(token: string | null): void {
	if (typeof window !== 'undefined') {
		if (token) {
			localStorage.setItem('uniapp_token', token);
		} else {
			localStorage.removeItem('uniapp_token');
		}
	}
}

export async function apiFetch<T>(
	endpoint: string,
	options: RequestInit = {}
): Promise<T> {
	const token = getAuthToken();
	const headers: Record<string, string> = {
		'Content-Type': 'application/json',
		...(options.headers as Record<string, string>)
	};

	if (token) {
		headers['Authorization'] = `Bearer ${token}`;
	}

	const url = endpoint.startsWith('http') ? endpoint : `${BASE_URL}${endpoint}`;

	const res = await fetch(url, {
		...options,
		headers
	});

	if (!res.ok) {
		let errorMessage = `HTTP Error ${res.status}: ${res.statusText}`;
		try {
			const errorData = await res.json();
			if (errorData && typeof errorData === 'object') {
				if (errorData.message) errorMessage = errorData.message;
				else if (typeof errorData === 'string') errorMessage = errorData;
				return Promise.reject({ status: res.status, ...errorData });
			}
		} catch {
			const text = await res.text();
			if (text) errorMessage = text;
		}
		throw new Error(errorMessage);
	}

	// For empty response bodies (e.g. 204 or no content)
	const text = await res.text();
	if (!text) return {} as T;

	try {
		return JSON.parse(text) as T;
	} catch {
		return text as unknown as T;
	}
}
