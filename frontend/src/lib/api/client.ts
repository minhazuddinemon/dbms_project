// src/lib/api/client.ts

import { toastState } from '$lib/state/toast.svelte';

const BASE_URL = import.meta.env.VITE_API_BASE_URL || '/api';

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

export function getAdminToken(): string | null {
	if (typeof window !== 'undefined') {
		return localStorage.getItem('uniapp_admin_token');
	}
	return null;
}

export function setAdminToken(token: string | null): void {
	if (typeof window !== 'undefined') {
		if (token) {
			localStorage.setItem('uniapp_admin_token', token);
		} else {
			localStorage.removeItem('uniapp_admin_token');
		}
	}
}

export interface ApiFetchOptions extends RequestInit {
	showErrorToast?: boolean;
	showSuccessToast?: string;
}

/**
 * Recursively unwraps Go sql.NullString / sql.NullTime objects
 * e.g. { String: "Dhaka", Valid: true } => "Dhaka"
 * e.g. { String: "", Valid: false } => null
 */
export function sanitizeNullFields<T>(data: any): T {
	if (data === null || data === undefined) return data;

	if (Array.isArray(data)) {
		return data.map((item) => sanitizeNullFields(item)) as unknown as T;
	}

	if (typeof data === 'object') {
		// Check for sql.NullString shape
		if (typeof data.Valid === 'boolean' && Object.prototype.hasOwnProperty.call(data, 'String')) {
			return (data.Valid ? data.String : null) as unknown as T;
		}

		// Check for sql.NullTime shape
		if (typeof data.Valid === 'boolean' && Object.prototype.hasOwnProperty.call(data, 'Time')) {
			return (data.Valid ? data.Time : null) as unknown as T;
		}

		const result: Record<string, any> = {};
		for (const key of Object.keys(data)) {
			result[key] = sanitizeNullFields(data[key]);
		}
		return result as T;
	}

	return data as T;
}

export async function apiFetch<T>(
	endpoint: string,
	options: ApiFetchOptions = {}
): Promise<T> {
	const { showErrorToast = true, showSuccessToast, ...fetchOptions } = options;
	const isAdminEndpoint = endpoint.startsWith('/admin') && !endpoint.startsWith('/admin/login');
	const token = isAdminEndpoint ? getAdminToken() : getAuthToken();

	const customHeaders = (fetchOptions.headers as Record<string, string>) || {};
	const headers: Record<string, string> = {
		'Content-Type': 'application/json',
		...customHeaders
	};

	const hasAuthHeader = Object.keys(headers).some((k) => k.toLowerCase() === 'authorization');

	if (token && !hasAuthHeader) {
		headers['Authorization'] = `Bearer ${token}`;
	}

	const url = endpoint.startsWith('http') ? endpoint : `${BASE_URL}${endpoint}`;

	let res: Response;
	try {
		res = await fetch(url, {
			...fetchOptions,
			headers
		});
	} catch (err: any) {
		const networkError = 'Network error: Unable to connect to server';
		if (showErrorToast) {
			toastState.error(networkError);
		}
		throw new Error(networkError);
	}

	if (!res.ok) {
		let errorMessage = `HTTP Error ${res.status}: ${res.statusText}`;
		let errorPayload: any = null;

		try {
			const text = await res.text();
			if (text) {
				try {
					errorPayload = JSON.parse(text);
					if (errorPayload && typeof errorPayload === 'object') {
						if (errorPayload.message) errorMessage = errorPayload.message;
						else if (errorPayload.error) errorMessage = errorPayload.error;
					}
				} catch {
					// Backend http.Error(...) plain text response
					errorMessage = text.trim();
				}
			}
		} catch {
			// fallback to default status text
		}

		if (showErrorToast) {
			toastState.error(errorMessage);
		}

		const errObj = new Error(errorMessage);
		if (errorPayload && typeof errorPayload === 'object') {
			Object.assign(errObj, errorPayload);
		}
		throw errObj;
	}

	if (showSuccessToast) {
		toastState.success(showSuccessToast);
	}

	const text = await res.text();
	if (!text) return {} as T;

	let parsedData: any;
	try {
		parsedData = JSON.parse(text);
	} catch {
		parsedData = text;
	}

	return sanitizeNullFields<T>(parsedData);
}
