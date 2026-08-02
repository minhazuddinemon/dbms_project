// src/lib/api/applications.ts

import { apiFetch } from './client';
import { toastState } from '$lib/state/toast.svelte';
import type { ApplyRequest, ApplyResponse, StudentApplication, ProgramRequirementsResponse } from '$lib/types/models';

export async function applyToProgram(programId: number): Promise<ApplyResponse> {
	const payload: ApplyRequest = { program_id: programId };
	try {
		const res = await apiFetch<ApplyResponse>('/applications/apply', {
			method: 'POST',
			body: JSON.stringify(payload),
			showErrorToast: false
		});

		if (res.status === 'success') {
			toastState.success(res.message || 'Application submitted successfully! Please proceed to payment.');
		} else if (res.status === 'incomplete_profile') {
			const missing = res.missing_fields?.join(', ');
			const msg = missing ? `${res.message} Missing: ${missing}` : res.message;
			toastState.warning(msg);
		}

		return res;
	} catch (err: any) {
		if (err?.status === 'incomplete_profile' || err?.missing_fields) {
			const missing = err.missing_fields?.join(', ');
			const msg = missing ? `${err.message || 'Incomplete profile.'} Missing: ${missing}` : (err.message || 'Incomplete profile.');
			toastState.warning(msg);
			return {
				status: 'incomplete_profile',
				message: err.message || 'You must provide additional information to apply to this program.',
				missing_fields: err.missing_fields
			};
		}

		const errMsg = err?.message || 'Failed to submit application';
		toastState.error(errMsg);
		throw err;
	}
}

export function normalizeApplication(raw: any): StudentApplication {
	if (!raw) return raw;

	const extractString = (val: any): string => {
		if (val === null || val === undefined) return '';
		if (typeof val === 'string') return val;
		if (typeof val === 'object') {
			if (val.Valid && val.String !== undefined) return val.String;
			if (val.String !== undefined) return val.String;
		}
		return String(val);
	};

	const extractDateString = (val: any): string => {
		if (!val) return '';
		if (typeof val === 'string') return val;
		if (typeof val === 'object') {
			if (val.Valid && val.Time) return val.Time;
			if (val.Time) return val.Time;
		}
		return String(val);
	};

	const extractNumber = (val: any, defaultVal = 0): number => {
		if (val === null || val === undefined || val === '') return defaultVal;
		if (typeof val === 'number') return isNaN(val) ? defaultVal : val;
		if (typeof val === 'string') {
			const parsed = parseFloat(val);
			return isNaN(parsed) ? defaultVal : parsed;
		}
		if (typeof val === 'object') {
			const str = val.Valid && val.String !== undefined ? val.String : val.String;
			if (str) {
				const parsed = parseFloat(str);
				return isNaN(parsed) ? defaultVal : parsed;
			}
		}
		return defaultVal;
	};

	return {
		...raw,
		app_id: extractNumber(raw.app_id),
		program_id: extractNumber(raw.program_id),
		sub_date: extractDateString(raw.sub_date),
		status: extractString(raw.status || 'PENDING'),
		program_name: extractString(raw.program_name || raw.p_name || `Program #${raw.program_id}`),
		university_name: extractString(raw.university_name || raw.u_name || 'Public University'),
		program_fee: extractString(raw.program_fee || raw.application_fee || '500.00'),
		student_id: extractNumber(raw.student_id),
		first_name: extractString(raw.first_name),
		last_name: extractString(raw.last_name),
		email: extractString(raw.email)
	};
}

export async function fetchStudentApplications(): Promise<StudentApplication[]> {
	const data = await apiFetch<any[]>('/applications', {
		method: 'GET'
	});
	return (data || []).map(normalizeApplication);
}

export async function fetchProgramRequirements(programId: number): Promise<ProgramRequirementsResponse> {
	return apiFetch<ProgramRequirementsResponse>(`/program/requirements?program_id=${programId}`, {
		method: 'GET'
	});
}
