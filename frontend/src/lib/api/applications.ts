// src/lib/api/applications.ts

import { apiFetch } from './client';
import type { ApplyRequest, ApplyResponse } from '$lib/types/models';

export async function applyToProgram(programId: number): Promise<ApplyResponse> {
	const payload: ApplyRequest = { program_id: programId };
	return apiFetch<ApplyResponse>('/applications/apply', {
		method: 'POST',
		body: JSON.stringify(payload)
	});
}
