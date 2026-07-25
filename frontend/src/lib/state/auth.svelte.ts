// src/lib/state/auth.svelte.ts

import { getAuthToken, setAuthToken } from '$lib/api/client';
import { getProfile, login as apiLogin, logout as apiLogout } from '$lib/api/auth';
import type { LoginRequest } from '$lib/types/models';

class AuthState {
	token = $state<string | null>(getAuthToken());
	studentId = $state<number | null>(null);
	isAuthenticated = $derived(!!this.token);
	isLoading = $state<boolean>(false);
	error = $state<string | null>(null);

	async checkAuth() {
		if (!this.token) return;
		this.isLoading = true;
		try {
			const res = await getProfile();
			this.studentId = res.student_id;
			this.error = null;
		} catch (err: any) {
			this.logout();
		} finally {
			this.isLoading = false;
		}
	}

	async login(credentials: LoginRequest) {
		this.isLoading = true;
		this.error = null;
		try {
			const res = await apiLogin(credentials);
			if (res.token) {
				this.token = res.token;
				await this.checkAuth();
			}
			return true;
		} catch (err: any) {
			this.error = err?.message || 'Login failed. Please check credentials.';
			return false;
		} finally {
			this.isLoading = false;
		}
	}

	logout() {
		apiLogout();
		this.token = null;
		this.studentId = null;
		this.error = null;
	}
}

export const authState = new AuthState();
