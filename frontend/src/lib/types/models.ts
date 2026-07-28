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
	university_name?: string;
	location?: string;
	university_location?: string;
	website?: string;
	logo_url?: string;
	university_logo?: string;
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

export interface DepartmentPayload {
	dept_name: string;
	dept_description?: string;
	total_seats: number;
}

export interface AlbumPayload {
	picture_title: string;
	picture_url: string;
}

export interface DepartmentResponse {
	dept_id: number;
	dept_name: string;
	dept_description: string;
	total_seats: number;
}

export interface AlbumResponse {
	album_id: number;
	picture_title: string;
	picture_url: string;
}

export interface University {
	u_id: number;
	u_name: string;
	website: string;
	location?: string;
	logo_url?: string;
	university_description?: string;
	university_history?: string;
	departments?: DepartmentResponse[];
	album?: AlbumResponse[];
}

export interface UniversityPayload {
	name: string;
	website: string;
	location?: string;
	logo_url?: string;
	university_description?: string;
	university_history?: string;
	departments?: DepartmentPayload[];
	album?: AlbumPayload[];
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

export type StudentMobileOwnerType = 'self' | 'mother' | 'father';

export interface StudentMobile {
	student_id: number;
	mobile_no: string;
	owner_type: StudentMobileOwnerType;
}

export interface StudentMobileRequest {
	mobile_no: string;
	owner_type: StudentMobileOwnerType;
}

export interface UpdateStudentMobileRequest {
	current_mobile_no: string;
	mobile_no: string;
	owner_type: StudentMobileOwnerType;
}

export interface AdmissionTestPayload {
	exam_unit?: string;
	exam_center?: string;
	exam_date: string; // YYYY-MM-DD
	prereq_test_id?: number | null;
	program_id: number;
}

export interface StudentAcademicRequest {
	exam_level: string; // 'SSC', 'HSC'
	year: number;
	roll_no: string;
	reg_no: string;
	gpa: string;
	board: string;
	edu_group: string; // 'Science', 'Humanities', 'Business'
}

export interface SubjectMarkItem {
	subject_name: string;
	marks: string;
	grade?: string;
}

export interface StudentSubjectMarksRequest {
	exam_level: string;
	subjects: SubjectMarkItem[];
}

export interface StudentNotification {
	notif_id: number;
	student_id: number;
	message: string;
	created_at: string;
}

export interface UniversityTransport {
	u_id: number;
	transport_route: string;
	est_travel_time: string;
}

export interface UniversityTransportRequest {
	u_id: number;
	transport_route: string;
	est_travel_time: string;
}

export interface ProgramEligibilityRule {
	program_id: number;
	rule_type: string;
	rule_value: string;
}

export interface ProgramEligibilityRuleRequest {
	program_id: number;
	rule_type: string;
	rule_value: string;
}

export interface StudentTestResultItem {
	student_id: number;
	marks: string;
	merit_position: number;
}

export interface PublishTestResultsRequest {
	test_id: number;
	results: StudentTestResultItem[];
}
