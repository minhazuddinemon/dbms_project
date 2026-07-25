// src/lib/api/university.ts

import { apiFetch } from './client';
import type { University } from '$lib/types/models';

export async function fetchUniversities(): Promise<University[]> {
	return apiFetch<University[]>('/universities', {
		method: 'GET'
	});
}

export async function fetchUniversityByID(uId: number): Promise<University> {
	return apiFetch<University>(`/universities/detail?u_id=${uId}`, {
		method: 'GET'
	});
}
