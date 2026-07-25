// src/lib/api/admin.ts

import { apiFetch, setAuthToken, setAdminToken } from './client';
import type {
	AdminLoginResponse,
	LoginRequest,
	StudentApplication,
	UniversityPayload,
	ProgramPayload
} from '$lib/types/models';

export async function adminLogin(payload: LoginRequest): Promise<AdminLoginResponse> {
	const res = await apiFetch<AdminLoginResponse>('/admin/login', {
		method: 'POST',
		body: JSON.stringify(payload)
	});
	if (res.token) {
		setAdminToken(res.token);
		setAuthToken(res.token);
	}
	return res;
}

export async function fetchUniversityApplications(uId: number): Promise<StudentApplication[]> {
	return apiFetch<StudentApplication[]>(`/admin/applications?u_id=${uId}`, {
		method: 'GET'
	});
}

export async function updateApplicationStatus(appId: number, status: string): Promise<{ status: string; message: string }> {
	return apiFetch<{ status: string; message: string }>('/admin/applications/status', {
		method: 'PUT',
		body: JSON.stringify({ app_id: appId, status })
	});
}

export async function createUniversity(payload: UniversityPayload): Promise<{ message: string; u_id: number }> {
	return apiFetch<{ message: string; u_id: number }>('/admin/university', {
		method: 'POST',
		body: JSON.stringify(payload)
	});
}

export async function updateUniversity(uId: number, payload: UniversityPayload): Promise<{ message: string }> {
	return apiFetch<{ message: string }>(`/admin/university?u_id=${uId}`, {
		method: 'PUT',
		body: JSON.stringify(payload)
	});
}

export async function deleteUniversity(uId: number): Promise<{ message: string }> {
	return apiFetch<{ message: string }>(`/admin/university?u_id=${uId}`, {
		method: 'DELETE'
	});
}

export async function createProgram(payload: ProgramPayload): Promise<{ message: string; program_id: number }> {
	return apiFetch<{ message: string; program_id: number }>('/admin/program', {
		method: 'POST',
		body: JSON.stringify(payload)
	});
}

export async function updateProgram(programId: number, payload: ProgramPayload): Promise<{ message: string }> {
	return apiFetch<{ message: string }>(`/admin/program?program_id=${programId}`, {
		method: 'PUT',
		body: JSON.stringify(payload)
	});
}

export async function deleteProgram(programId: number): Promise<{ message: string }> {
	return apiFetch<{ message: string }>(`/admin/program?program_id=${programId}`, {
		method: 'DELETE'
	});
}
