// src/lib/api/university.ts

import { apiFetch } from './client';
import type { University, UniversityTransport } from '$lib/types/models';

export function normalizeUniversity(raw: any): University {
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
		location: extractString(raw.location),
		logo_url: extractString(raw.logo_url),
		university_description: extractString(raw.university_description),
		university_history: extractString(raw.university_history),
		u_name: extractString(raw.u_name || raw.name),
		website: extractString(raw.website)
	};
}

export async function fetchUniversities(): Promise<University[]> {
	const data = await apiFetch<any[]>('/universities', {
		method: 'GET'
	});
	return (data || []).map(normalizeUniversity);
}

export async function fetchUniversityByID(uId: number): Promise<University> {
	const data = await apiFetch<any>(`/universities/detail?u_id=${uId}`, {
		method: 'GET'
	});
	return normalizeUniversity(data);
}

export async function fetchUniversityTransport(uId: number): Promise<UniversityTransport[]> {
	return apiFetch<UniversityTransport[]>(`/universities/transport?u_id=${uId}`, {
		method: 'GET'
	});
}
