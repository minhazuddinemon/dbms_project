// src/lib/api/student.ts

import { apiFetch } from './client';

export async function updateStudentProfile(
	fields: Record<string, string>
): Promise<{ status: string; message: string }> {
	return apiFetch<{ status: string; message: string }>('/student/profile', {
		method: 'POST',
		body: JSON.stringify(fields)
	});
}
