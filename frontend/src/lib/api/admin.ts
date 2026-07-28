// src/lib/api/admin.ts

import { apiFetch, setAuthToken, setAdminToken } from './client';
import type {
	AdminLoginResponse,
	LoginRequest,
	StudentApplication,
	UniversityPayload,
	ProgramPayload,
	AdmissionTestPayload,
	ProgramEligibilityRule,
	ProgramEligibilityRuleRequest,
	UniversityTransportRequest,
	PublishTestResultsRequest
} from '$lib/types/models';

export async function adminLogin(payload: LoginRequest): Promise<AdminLoginResponse> {
	const res = await apiFetch<AdminLoginResponse>('/admin/login', {
		method: 'POST',
		body: JSON.stringify(payload),
		showSuccessToast: 'Admin login successful!'
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
		body: JSON.stringify({ app_id: appId, status }),
		showSuccessToast: `Application status updated to ${status}`
	});
}

export async function createUniversity(payload: UniversityPayload): Promise<{ message: string; u_id: number }> {
	return apiFetch<{ message: string; u_id: number }>('/admin/university', {
		method: 'POST',
		body: JSON.stringify(payload),
		showSuccessToast: 'University created successfully!'
	});
}

export async function updateUniversity(uId: number, payload: UniversityPayload): Promise<{ message: string }> {
	return apiFetch<{ message: string }>(`/admin/university?u_id=${uId}`, {
		method: 'PUT',
		body: JSON.stringify(payload),
		showSuccessToast: 'University updated successfully!'
	});
}

export async function deleteUniversity(uId: number): Promise<{ message: string }> {
	return apiFetch<{ message: string }>(`/admin/university?u_id=${uId}`, {
		method: 'DELETE',
		showSuccessToast: 'University deleted successfully!'
	});
}

export async function createProgram(payload: ProgramPayload): Promise<{ message: string; program_id: number }> {
	return apiFetch<{ message: string; program_id: number }>('/admin/program', {
		method: 'POST',
		body: JSON.stringify(payload),
		showSuccessToast: 'Program created successfully!'
	});
}

export async function updateProgram(programId: number, payload: ProgramPayload): Promise<{ message: string }> {
	return apiFetch<{ message: string }>(`/admin/program?program_id=${programId}`, {
		method: 'PUT',
		body: JSON.stringify(payload),
		showSuccessToast: 'Program updated successfully!'
	});
}

export async function deleteProgram(programId: number): Promise<{ message: string }> {
	return apiFetch<{ message: string }>(`/admin/program?program_id=${programId}`, {
		method: 'DELETE',
		showSuccessToast: 'Program deleted successfully!'
	});
}

export async function createAdmissionTest(payload: AdmissionTestPayload): Promise<{ message: string; test_id: number }> {
	return apiFetch<{ message: string; test_id: number }>('/admin/admission-test', {
		method: 'POST',
		body: JSON.stringify(payload),
		showSuccessToast: 'Admission test created successfully!'
	});
}

export async function updateAdmissionTest(testId: number, payload: AdmissionTestPayload): Promise<{ message: string }> {
	return apiFetch<{ message: string }>(`/admin/admission-test?test_id=${testId}`, {
		method: 'PUT',
		body: JSON.stringify(payload),
		showSuccessToast: 'Admission test updated successfully!'
	});
}

export async function fetchProgramEligibilityRules(programId: number): Promise<ProgramEligibilityRule[]> {
	return apiFetch<ProgramEligibilityRule[]>(`/admin/program/eligibility-rules?program_id=${programId}`, {
		method: 'GET'
	});
}

export async function saveProgramEligibilityRule(
	payload: ProgramEligibilityRuleRequest
): Promise<{ message: string }> {
	return apiFetch<{ message: string }>('/admin/program/eligibility-rules', {
		method: 'POST',
		body: JSON.stringify(payload),
		showSuccessToast: 'Eligibility rule saved successfully!'
	});
}

export async function deleteProgramEligibilityRule(
	programId: number,
	ruleType: string
): Promise<{ message: string }> {
	return apiFetch<{ message: string }>(
		`/admin/program/eligibility-rules?program_id=${programId}&rule_type=${encodeURIComponent(ruleType)}`,
		{
			method: 'DELETE',
			showSuccessToast: 'Eligibility rule deleted successfully!'
		}
	);
}

export async function createUniversityTransport(
	payload: UniversityTransportRequest
): Promise<{ message: string }> {
	return apiFetch<{ message: string }>('/admin/universities/transport', {
		method: 'POST',
		body: JSON.stringify(payload),
		showSuccessToast: 'Transport route created successfully!'
	});
}

export async function updateUniversityTransport(
	payload: UniversityTransportRequest
): Promise<{ message: string }> {
	return apiFetch<{ message: string }>('/admin/universities/transport', {
		method: 'PUT',
		body: JSON.stringify(payload),
		showSuccessToast: 'Transport route updated successfully!'
	});
}

export async function deleteUniversityTransport(
	uId: number,
	transportRoute: string
): Promise<{ message: string }> {
	return apiFetch<{ message: string }>(
		`/admin/universities/transport?u_id=${uId}&transport_route=${encodeURIComponent(transportRoute)}`,
		{
			method: 'DELETE',
			showSuccessToast: 'Transport route deleted successfully!'
		}
	);
}

export async function publishAdmissionTestResults(
	payload: PublishTestResultsRequest
): Promise<{ message: string }> {
	return apiFetch<{ message: string }>('/admin/admission-test/publish', {
		method: 'POST',
		body: JSON.stringify(payload),
		showSuccessToast: 'Admission test results published and students notified!'
	});
}
