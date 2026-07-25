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

export interface University {
	u_id: number;
	u_name: string;
	website: string;
	location?: string;
	logo_url?: string;
}

export interface UniversityPayload {
	name: string;
	website: string;
	location: string;
	logo_url: string;
}

export interface ProgramPayload {
	p_name: string;
	p_unit: string;
	total_seats: number;
	prev_cutmarks: number;
	deadline: string;
	u_id: number;
}

export interface StudentApplication {
	app_id: number;
	sub_date: string;
	status: string;
	program_id: number;
	program_name?: string;
	university_name?: string;
	student_id?: number;
	first_name?: string;
	last_name?: string;
	email?: string;
}

export interface RequiredFieldStatus {
	field_name: string;
	value: string | null;
	is_provided: boolean;
}

export interface ProgramRequirementsResponse {
	program_id: number;
	is_ready_to_apply: boolean;
	required_fields: RequiredFieldStatus[];
	missing_fields: string[];
}

export interface PaymentRequest {
	application_id: number;
	amount: string;
	payment_method: string;
	transaction_id: string;
}

export interface PaymentResponse {
	status: string;
	message: string;
	application_id: number;
	transaction_id: string;
}

export interface AdminLoginResponse {
	token: string;
	role: string;
}
