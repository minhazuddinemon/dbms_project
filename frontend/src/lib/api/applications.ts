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

export async function fetchStudentApplications(): Promise<StudentApplication[]> {
	return apiFetch<StudentApplication[]>('/applications', {
		method: 'GET'
	});
}

export async function fetchProgramRequirements(programId: number): Promise<ProgramRequirementsResponse> {
	return apiFetch<ProgramRequirementsResponse>(`/program/requirements?program_id=${programId}`, {
		method: 'GET'
	});
}
