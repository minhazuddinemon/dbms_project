// src/lib/api/programs.ts

import { apiFetch } from './client';
import type { Program, EligibleProgram } from '$lib/types/models';

export async function fetchPrograms(search?: string, unit?: string): Promise<Program[]> {
	const params = new URLSearchParams();
	if (search) params.append('search', search);
	if (unit) params.append('unit', unit);

	const queryStr = params.toString() ? `?${params.toString()}` : '';
	return apiFetch<Program[]>(`/programs${queryStr}`, {
		method: 'GET'
	});
}

export async function fetchProgramByID(id: number): Promise<Program> {
	return apiFetch<Program>(`/programs/detail?id=${id}`, {
		method: 'GET'
	});
}

export async function fetchEligiblePrograms(): Promise<EligibleProgram[]> {
	return apiFetch<EligibleProgram[]>('/programs/eligible', {
		method: 'GET'
	});
}
