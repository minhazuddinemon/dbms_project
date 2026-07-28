// src/lib/api/student.ts

import { apiFetch } from './client';
import type {
	StudentMobile,
	StudentMobileRequest,
	UpdateStudentMobileRequest,
	StudentAcademicRequest,
	StudentSubjectMarksRequest,
	StudentNotification
} from '$lib/types/models';

export async function updateStudentProfile(
	fields: Record<string, string>
): Promise<{ status: string; message: string }> {
	return apiFetch<{ status: string; message: string }>('/student/profile', {
		method: 'POST',
		body: JSON.stringify(fields),
		showSuccessToast: 'Profile updated successfully!'
	});
}

export async function fetchStudentMobiles(): Promise<StudentMobile[]> {
	return apiFetch<StudentMobile[]>('/student/mobile', {
		method: 'GET'
	});
}

export async function addStudentMobile(
	payload: StudentMobileRequest
): Promise<{ message: string; mobile: StudentMobile }> {
	return apiFetch<{ message: string; mobile: StudentMobile }>('/student/mobile', {
		method: 'POST',
		body: JSON.stringify(payload),
		showSuccessToast: 'Mobile number added successfully!'
	});
}

export async function updateStudentMobile(
	payload: UpdateStudentMobileRequest
): Promise<{ message: string; mobile: StudentMobile }> {
	return apiFetch<{ message: string; mobile: StudentMobile }>('/student/mobile', {
		method: 'PUT',
		body: JSON.stringify(payload),
		showSuccessToast: 'Mobile number updated successfully!'
	});
}

export async function deleteStudentMobile(
	mobileNo: string
): Promise<{ message: string }> {
	return apiFetch<{ message: string }>(`/student/mobile?mobile_no=${encodeURIComponent(mobileNo)}`, {
		method: 'DELETE',
		showSuccessToast: 'Mobile number deleted successfully!'
	});
}

export async function saveStudentAcademic(
	payload: StudentAcademicRequest
): Promise<{ message: string }> {
	return apiFetch<{ message: string }>('/student/academic', {
		method: 'POST',
		body: JSON.stringify(payload),
		showSuccessToast: 'Academic record saved successfully!'
	});
}

export async function saveStudentSubjectMarks(
	payload: StudentSubjectMarksRequest
): Promise<{ message: string }> {
	return apiFetch<{ message: string }>('/student/subject-marks', {
		method: 'POST',
		body: JSON.stringify(payload),
		showSuccessToast: 'Subject marks saved successfully!'
	});
}

export async function fetchStudentNotifications(): Promise<StudentNotification[]> {
	return apiFetch<StudentNotification[]>('/student/notifications', {
		method: 'GET'
	});
}
