// src/lib/api/programs.ts

import { apiFetch } from './client';
import type { Program, EligibleProgram } from '$lib/types/models';

export function normalizeProgram(raw: any): Program {
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

	return {
		...raw,
		p_unit: extractString(raw.p_unit),
		prev_cutmarks: extractString(raw.prev_cutmarks),
		university_location: extractString(raw.university_location || raw.location),
		location: extractString(raw.location || raw.university_location),
		university_name: extractString(raw.university_name || raw.u_name),
		u_name: extractString(raw.u_name || raw.university_name)
	};
}

export async function fetchPrograms(search?: string, unit?: string): Promise<Program[]> {
	const params = new URLSearchParams();
	if (search) params.append('search', search);
	if (unit) params.append('unit', unit);

	const queryStr = params.toString() ? `?${params.toString()}` : '';
	const data = await apiFetch<any[]>(`/programs${queryStr}`, {
		method: 'GET'
	});
	return (data || []).map(normalizeProgram);
}

export async function fetchProgramByID(id: number): Promise<Program> {
	const data = await apiFetch<any>(`/programs/detail?id=${id}`, {
		method: 'GET'
	});
	return normalizeProgram(data);
}

export async function fetchEligiblePrograms(): Promise<EligibleProgram[]> {
	return apiFetch<EligibleProgram[]>('/programs/eligible', {
		method: 'GET'
	});
}
