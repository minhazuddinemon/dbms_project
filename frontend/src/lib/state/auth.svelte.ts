// src/lib/state/auth.svelte.ts

import { getAuthToken, setAuthToken, getAdminToken, setAdminToken } from '$lib/api/client';
import { getProfile, login as apiLogin, logout as apiLogout } from '$lib/api/auth';
import { toastState } from '$lib/state/toast.svelte';
import type { LoginRequest } from '$lib/types/models';

class AuthState {
	token = $state<string | null>(getAuthToken());
	adminToken = $state<string | null>(getAdminToken());
	studentId = $state<number | null>(null);
	isAuthenticated = $derived(!!this.token || !!this.adminToken);
	isAdmin = $derived(!!this.adminToken);
	isLoading = $state<boolean>(false);
	error = $state<string | null>(null);

	async checkAuth() {
		if (this.adminToken) return;
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
				setAuthToken(res.token);
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

	setAdmin(token: string) {
		this.adminToken = token;
		this.token = token;
		setAdminToken(token);
		setAuthToken(token);
	}

	logout() {
		apiLogout();
		setAdminToken(null);
		setAuthToken(null);
		this.token = null;
		this.adminToken = null;
		this.studentId = null;
		this.error = null;

		if (typeof window !== 'undefined') {
			localStorage.removeItem('uniapp_profile_created');
			localStorage.removeItem('uniapp_student_academic_profile');
			localStorage.removeItem('uniapp_token');
			localStorage.removeItem('uniapp_admin_token');
		}

		toastState.info('Logged out successfully');
	}
}

export const authState = new AuthState();
