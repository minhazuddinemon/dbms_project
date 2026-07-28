// src/lib/api/auth.ts

import { apiFetch, setAuthToken } from './client';
import type { LoginRequest, RegisterRequest, AuthResponse } from '$lib/types/models';

export async function login(payload: LoginRequest): Promise<AuthResponse> {
	const res = await apiFetch<AuthResponse>('/login', {
		method: 'POST',
		body: JSON.stringify(payload),
		showSuccessToast: 'Login successful!'
	});

	if (res.token) {
		setAuthToken(res.token);
	}
	return res;
}

export async function register(payload: RegisterRequest): Promise<AuthResponse> {
	return apiFetch<AuthResponse>('/register', {
		method: 'POST',
		body: JSON.stringify(payload),
		showSuccessToast: 'Registration successful! You can now log in.'
	});
}

export async function getProfile(): Promise<{ message: string; student_id: number }> {
	return apiFetch<{ message: string; student_id: number }>('/profile', {
		method: 'GET',
		showErrorToast: false
	});
}

export function logout(): void {
	setAuthToken(null);
}
