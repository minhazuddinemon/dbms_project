// src/lib/api/applications.ts

import { apiFetch } from './client';
import type { ApplyRequest, ApplyResponse, StudentApplication, ProgramRequirementsResponse } from '$lib/types/models';

export async function applyToProgram(programId: number): Promise<ApplyResponse> {
	const payload: ApplyRequest = { program_id: programId };
	return apiFetch<ApplyResponse>('/applications/apply', {
		method: 'POST',
		body: JSON.stringify(payload)
	});
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
