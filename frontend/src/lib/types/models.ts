// src/lib/types/models.ts

export interface StudentUser {
	student_id: number;
	first_name?: string;
	last_name?: string;
	email?: string;
}

export interface RegisterRequest {
	first_name: string;
	last_name: string;
	email: string;
	password: string;
	dob: string; // YYYY-MM-DD
}

export interface LoginRequest {
	email: string;
	password: string;
}

export interface AuthResponse {
	message: string;
	token?: string;
}

export interface Program {
	program_id: number;
	p_name: string;
	p_unit: string;
	total_seats: number;
	prev_cutmarks: string | number;
	deadline: string;
	u_id: number;
	u_name?: string;
	location?: string;
	website?: string;
}

export interface EligibleProgram {
	program_id: number;
	program_name: string;
	university_name: string;
}

export interface ApplyRequest {
	program_id: number;
}

export interface ApplyResponse {
	status: 'success' | 'incomplete_profile';
	message: string;
	missing_fields?: string[];
}

export type RequiredProfileField =
	| 'PRESENT_ADDRESS'
	| 'PERMANENT_ADDRESS'
	| 'FATHERS_NAME'
	| 'MOTHERS_NAME'
	| 'BLOOD_GROUP'
	| 'QUOTA'
	| 'PHOTO_URL'
	| 'SIGNATURE_URL';
