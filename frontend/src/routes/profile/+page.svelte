<!-- src/routes/profile/+page.svelte -->
<script lang="ts">
	import {
		fetchStudentProfile,
		updateStudentProfile,
		fetchStudentAcademic,
		saveStudentAcademic,
		saveStudentSubjectMarks,
		fetchStudentMobiles,
		addStudentMobile,
		updateStudentMobile,
		deleteStudentMobile
	} from '$lib/api/student';
	import { authState } from '$lib/state/auth.svelte';
	import { toastState } from '$lib/state/toast.svelte';
	import type { StudentMobile, StudentMobileOwnerType } from '$lib/types/models';
	import {
		User,
		MapPin,
		CheckCircle2,
		AlertCircle,
		Save,
		BookOpen,
		Phone,
		Plus,
		Trash2,
		Edit3,
		ArrowLeft,
		ArrowRight,
		Award,
		ShieldCheck,
		Image as ImageIcon,
		FileText,
		Camera,
		Upload
	} from 'lucide-svelte';
	import { onMount } from 'svelte';

	let presentAddress = $state('');
	let permanentAddress = $state('');
	let fathersName = $state('');
	let mothersName = $state('');
	let bloodGroup = $state('O+');
	let quota = $state('GENERAL');
	let photoUrl = $state('');
	let signatureUrl = $state('');

	// Academics (Compulsory fields)
	let regNo = $state('');
	let sscYear = $state('');
	let sscRollNo = $state('');
	let sscGpa = $state('');

	let hscYear = $state('');
	let hscRollNo = $state('');
	let hscGpa = $state('');
	let board = $state('Dhaka');
	let hscGroup = $state('Science');

	// Subject Marks (Compulsory fields)
	let physicsMarks = $state('');
	let mathMarks = $state('');
	let chemistryMarks = $state('');

	// Saved marks indicator & summary
	let hasSavedMarks = $state(false);
	let savedMarksSummary = $state<{ hscGpa: string; sscGpa: string; physics: string; math: string; chemistry: string } | null>(null);

	// Student Mobiles
	let mobiles = $state<StudentMobile[]>([]);
	let newMobileNo = $state('');
	let newOwnerType = $state<StudentMobileOwnerType>('self');
	let editingMobileNo = $state<string | null>(null);

	let activeTab = $state<'personal' | 'academics' | 'documents' | 'mobiles' | 'address'>('personal');
	let isLoading = $state(false);
	let successMessage = $state<string | null>(null);
	let errorMessage = $state<string | null>(null);

	async function loadMobiles() {
		try {
			mobiles = await fetchStudentMobiles();
		} catch (err) {
			console.error('Failed to load mobiles:', err);
		}
	}

	function loadSavedAcademicProfile() {
		const raw = localStorage.getItem('uniapp_student_academic_profile');
		if (raw) {
			try {
				const data = JSON.parse(raw);
				if (data.presentAddress) presentAddress = data.presentAddress;
				if (data.permanentAddress) permanentAddress = data.permanentAddress;
				if (data.fathersName) fathersName = data.fathersName;
				if (data.mothersName) mothersName = data.mothersName;
				if (data.bloodGroup) bloodGroup = data.bloodGroup;
				if (data.quota) quota = data.quota;
				if (data.photoUrl) photoUrl = data.photoUrl;
				if (data.signatureUrl) signatureUrl = data.signatureUrl;
				if (data.examLevel) examLevel = data.examLevel;
				if (data.passingYear) passingYear = data.passingYear;
				if (data.rollNo) rollNo = data.rollNo;
				if (data.regNo) regNo = data.regNo;
				if (data.sscYear) sscYear = data.sscYear;
				if (data.sscRollNo) sscRollNo = data.sscRollNo;
				if (data.sscGpa) sscGpa = data.sscGpa;
				if (data.hscYear) hscYear = data.hscYear;
				if (data.hscRollNo) hscRollNo = data.hscRollNo;
				if (data.hscGpa) hscGpa = data.hscGpa;
				if (data.board) board = data.board;
				if (data.hscGroup) hscGroup = data.hscGroup;
				if (data.physicsMarks) physicsMarks = data.physicsMarks;
				if (data.mathMarks) mathMarks = data.mathMarks;
				if (data.chemistryMarks) chemistryMarks = data.chemistryMarks;

				if (data.hscGpa && data.physicsMarks && data.mathMarks && data.chemistryMarks) {
					hasSavedMarks = true;
					savedMarksSummary = {
						hscGpa: String(data.hscGpa),
						sscGpa:String(data.sscGpa),
						physics: String(data.physicsMarks),
						math: String(data.mathMarks),
						chemistry: String(data.chemistryMarks)
					};
				}

				if (data.hscGpa && data.physicsMarks && data.mathMarks && data.chemistryMarks) {
					hasSavedMarks = true;
					savedMarksSummary = {
						hscGpa: data.hscGpa,
						sscGpa: data.sscGpa,
						physics: data.physicsMarks,
						math: data.mathMarks,
						chemistry: data.chemistryMarks
					};
				}
			} catch (e) {
				console.error('Failed to parse saved academic profile:', e);
			}
		}
	}

	async function loadProfileDataFromBackend() {
		try {
			// 1. Fetch Profile Fields (Addresses, Guardian Info, Documents)
			const fields = await fetchStudentProfile();
			if (fields) {
				if (fields.PRESENT_ADDRESS) presentAddress = fields.PRESENT_ADDRESS;
				if (fields.PERMANENT_ADDRESS) permanentAddress = fields.PERMANENT_ADDRESS;
				if (fields.FATHERS_NAME) fathersName = fields.FATHERS_NAME;
				if (fields.MOTHERS_NAME) mothersName = fields.MOTHERS_NAME;
				if (fields.BLOOD_GROUP) bloodGroup = fields.BLOOD_GROUP;
				if (fields.QUOTA) quota = fields.QUOTA;
				if (fields.PHOTO_URL) photoUrl = fields.PHOTO_URL;
				if (fields.SIGNATURE_URL) signatureUrl = fields.SIGNATURE_URL;
			}
		} catch (e) {
			console.error('Failed to fetch profile fields from backend:', e);
		}

		try {
			// 2. Fetch Academics & Subject Marks
			const res = await fetchStudentAcademic();
			if (res && res.academics && res.academics.length > 0) {
				for (const a of res.academics) {
					if (a.reg_no) regNo = a.reg_no;
					if (a.board) board = a.board;
					if (a.edu_group) hscGroup = a.edu_group;

					if (a.exam_level === 'SSC') {
						if (a.year) sscYear = a.year;
						if (a.roll_no) sscRollNo = a.roll_no;
						if (a.gpa) sscGpa = a.gpa;
					} else if (a.exam_level === 'HSC') {
						if (a.year) hscYear = a.year;
						if (a.roll_no) hscRollNo = a.roll_no;
						if (a.gpa) hscGpa = a.gpa;
					}
				}
			}

			if (res && res.subject_marks && res.subject_marks.length > 0) {
				for (const sm of res.subject_marks) {
					const nameLower = sm.subject_name.toLowerCase();
					if (nameLower.includes('physics')) physicsMarks = sm.marks;
					else if (nameLower.includes('math')) mathMarks = sm.marks;
					else if (nameLower.includes('chem')) chemistryMarks = sm.marks;
				}
			}
		} catch (e) {
			console.error('Failed to fetch academics from backend:', e);
		}

		if (hscGpa && physicsMarks && mathMarks && chemistryMarks) {
			hasSavedMarks = true;
			savedMarksSummary = {
				hscGpa: String(hscGpa),
				sscGpa: String(sscGpa),
				physics: String(physicsMarks),
				math: String(mathMarks),
				chemistry: String(chemistryMarks)
			};
		}
	}

	onMount(async () => {
		if (authState.isAuthenticated) {
			loadMobiles();
			loadSavedAcademicProfile();
			await loadProfileDataFromBackend();
		}
	});

	async function handleSaveProfile(e: Event) {
		e.preventDefault();
		isLoading = true;
		successMessage = null;
		errorMessage = null;

		const isFilled = (val: any): boolean => {
			if (val === null || val === undefined) return false;
			return String(val).trim().length > 0;
		};

		// Compulsory Academic Marks Validation
		if (
			!isFilled(regNo) ||
			!isFilled(sscRollNo) ||
			!isFilled(sscGpa) ||
			!isFilled(hscRollNo) ||
			!isFilled(hscGpa) ||
			!isFilled(physicsMarks) ||
			!isFilled(mathMarks) ||
			!isFilled(chemistryMarks)
		) {
			errorMessage = 'Academic marks (Registration No, SSC Roll, SSC GPA, HSC Roll, HSC GPA, Physics, Mathematics, and Chemistry marks) are compulsory to save your profile!';
			toastState.error(errorMessage);
			isLoading = false;
			return;
		}

		try {
			// 1. Update basic profile info & required documents
			await updateStudentProfile({
				PRESENT_ADDRESS: String(presentAddress || ''),
				PERMANENT_ADDRESS: String(permanentAddress || ''),
				FATHERS_NAME: String(fathersName || ''),
				MOTHERS_NAME: String(mothersName || ''),
				BLOOD_GROUP: String(bloodGroup || 'O+'),
				QUOTA: String(quota || 'GENERAL'),
				PHOTO_URL: String(photoUrl || ''),
				SIGNATURE_URL: String(signatureUrl || '')
			});

			// 2. Save SSC Academic Record
			try {
				await saveStudentAcademic({
					exam_level: 'SSC',
					year: Number(sscYear),
					roll_no: String(sscRollNo),
					reg_no: String(regNo),
					gpa: String(sscGpa),
					board: String(board),
					edu_group: String(hscGroup)
				});
			} catch (e) {
				// Ignore duplicate key if already saved
			}

			// 3. Save HSC Academic Record
			await saveStudentAcademic({
				exam_level: 'HSC',
				year: Number(hscYear),
				roll_no: String(hscRollNo),
				reg_no: String(regNo),
				gpa: String(hscGpa),
				board: String(board),
				edu_group: String(hscGroup)
			});

			// 4. Save HSC Subject Marks
			await saveStudentSubjectMarks({
				exam_level: 'HSC',
				subjects: [
					{ subject_name: 'Physics', marks: String(physicsMarks), grade: 'A+' },
					{ subject_name: 'Mathematics', marks: String(mathMarks), grade: 'A+' },
					{ subject_name: 'Chemistry', marks: String(chemistryMarks), grade: 'A+' }
				]
			});

			// Save locally so frontend persists and displays filled-up marks
			const academicProfileData = {
				presentAddress,
				permanentAddress,
				fathersName,
				mothersName,
				bloodGroup,
				quota,
				photoUrl,
				signatureUrl,
				regNo: String(regNo),
				sscYear: Number(sscYear),
				sscRollNo: String(sscRollNo),
				sscGpa: String(sscGpa),
				hscYear: Number(hscYear),
				hscRollNo: String(hscRollNo),
				hscGpa: String(hscGpa),
				board,
				hscGroup,
				physicsMarks: String(physicsMarks),
				mathMarks: String(mathMarks),
				chemistryMarks: String(chemistryMarks),
				savedAt: new Date().toISOString()
			};
			localStorage.setItem('uniapp_student_academic_profile', JSON.stringify(academicProfileData));
			localStorage.setItem('uniapp_profile_created', 'true');
			hasSavedMarks = true;
			savedMarksSummary = {
				hscGpa: String(hscGpa),
				sscGpa: String(sscGpa),
				physics: String(physicsMarks),
				math: String(mathMarks),
				chemistry: String(chemistryMarks)
			};

			successMessage = 'Profile, Academic Records & Subject Marks saved successfully!';
			// SINGLE UNIFIED TOAST
			toastState.success('Academic profile, exam records & subject marks saved successfully!');
		} catch (err: any) {
			console.error('Save profile error:', err);
			errorMessage = err?.message || 'Failed to update academic profile fields.';
			toastState.error(errorMessage);
		} finally {
			isLoading = false;
		}
	}

	// Mobile Number Handlers
	async function handleAddMobile(e: Event) {
		e.preventDefault();
		if (!newMobileNo.trim()) return;
		isLoading = true;
		try {
			if (editingMobileNo) {
				await updateStudentMobile({
					current_mobile_no: editingMobileNo,
					mobile_no: newMobileNo,
					owner_type: newOwnerType
				});
				editingMobileNo = null;
			} else {
				await addStudentMobile({
					mobile_no: newMobileNo,
					owner_type: newOwnerType
				});
			}
			newMobileNo = '';
			await loadMobiles();
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to save mobile number');
		} finally {
			isLoading = false;
		}
	}

	function startEditMobile(m: StudentMobile) {
		editingMobileNo = m.mobile_no;
		newMobileNo = m.mobile_no;
		newOwnerType = m.owner_type;
	}

	async function handleDeleteMobile(mobileNo: string) {
		if (!confirm(`Delete mobile number ${mobileNo}?`)) return;
		try {
			await deleteStudentMobile(mobileNo);
			await loadMobiles();
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to delete mobile number');
		}
	}
</script>

<svelte:head>
	<title>Academic Profile & Marks - UniApp</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		
		<!-- Banner Header -->
		<div class="bg-gradient-to-r from-primary via-primary-container to-secondary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-3">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<BookOpen class="w-4 h-4" />
					Academic Data Collection
				</div>

				<h1 class="text-3xl sm:text-4xl font-black text-white">
					Student Academic Profile
				</h1>

				<p class="text-slate-100 text-base leading-relaxed">
					Provide your guardian details, SSC/HSC GPA, subject marks, and contact mobile numbers for automated eligibility verification.
				</p>
			</div>
		</div>

		{#if !authState.isAuthenticated}
			<div class="glass-panel rounded-[2.5rem] border border-outline-variant/30 p-10 text-center space-y-4 bg-white/90 shadow-xl">
				<User class="w-14 h-14 text-outline mx-auto" />
				<h3 class="text-2xl font-bold text-on-surface">Sign In Required</h3>
				<p class="text-on-surface-variant text-sm">Please sign in to update your academic information.</p>
				<a href="/login" class="inline-block px-8 py-3 rounded-xl font-bold bg-primary text-white text-sm hover:bg-primary-container shadow-lg shadow-primary/25 transition-all">Sign In</a>
			</div>
		{:else}
			{#if hasSavedMarks && savedMarksSummary}
				<!-- Submitted Marks Banner Indicator -->
				<div class="p-5 rounded-2xl bg-emerald-50 border border-emerald-200 flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4 text-emerald-900 shadow-sm">
					<div class="flex items-center gap-3.5">
						<div class="w-10 h-10 rounded-xl bg-emerald-500/20 text-emerald-700 flex items-center justify-center shrink-0">
							<ShieldCheck class="w-6 h-6" />
						</div>
						<div>
							<h4 class="font-extrabold text-sm text-emerald-950">Submitted Academic Marks</h4>
							<p class="text-xs text-emerald-700 font-semibold mt-0.5">
								HSC GPA: <strong>{savedMarksSummary.hscGpa}</strong> | SSC GPA: <strong>{savedMarksSummary.sscGpa}</strong> | Physics: <strong>{savedMarksSummary.physics}</strong> | Math: <strong>{savedMarksSummary.math}</strong> | Chemistry: <strong>{savedMarksSummary.chemistry}</strong>
							</p>
						</div>
					</div>
					<span class="px-3.5 py-1 bg-emerald-600 text-white text-xs font-bold rounded-full uppercase tracking-wider shrink-0 shadow-xs">Submitted</span>
				</div>
			{/if}

			{#if successMessage}
				<div class="p-4 rounded-2xl bg-tertiary-fixed/30 border border-tertiary/40 flex items-center gap-3 text-on-tertiary-fixed-variant text-sm">
					<CheckCircle2 class="w-5 h-5 text-tertiary shrink-0" />
					<span class="font-bold">{successMessage}</span>
				</div>
			{/if}

			{#if errorMessage}
				<div class="p-4 rounded-2xl bg-error-container/60 border border-error/30 flex items-center gap-3 text-on-error-container text-sm">
					<AlertCircle class="w-5 h-5 text-error shrink-0" />
					<span>{errorMessage}</span>
				</div>
			{/if}

			<!-- Step Tabs Navigation -->
			<div class="flex items-center gap-2 bg-surface-container-low/80 p-2 rounded-2xl border border-outline-variant/30 overflow-x-auto">
				<button
					type="button"
					onclick={() => activeTab = 'personal'}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all duration-200 flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'personal' ? 'bg-white text-primary shadow-sm border border-outline-variant/30' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<User class="w-4 h-4" />
					Guardian Info
				</button>
				<button
					type="button"
					onclick={() => activeTab = 'academics'}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all duration-200 flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'academics' ? 'bg-white text-primary shadow-sm border border-outline-variant/30' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<BookOpen class="w-4 h-4" />
					Marks & GPA <span class="text-xs text-error font-extrabold">*</span>
				</button>
				<button
					type="button"
					onclick={() => activeTab = 'documents'}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all duration-200 flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'documents' ? 'bg-white text-primary shadow-sm border border-outline-variant/30' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<Camera class="w-4 h-4" />
					Photos & Docs
				</button>
				<button
					type="button"
					onclick={() => activeTab = 'mobiles'}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all duration-200 flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'mobiles' ? 'bg-white text-primary shadow-sm border border-outline-variant/30' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<Phone class="w-4 h-4" />
					Mobiles ({mobiles.length})
				</button>
				<button
					type="button"
					onclick={() => activeTab = 'address'}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all duration-200 flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'address' ? 'bg-white text-primary shadow-sm border border-outline-variant/30' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<MapPin class="w-4 h-4" />
					Address & Quota
				</button>
			</div>

			<!-- Tab: Mobile Numbers -->
			{#if activeTab === 'mobiles'}
				<div class="glass-panel rounded-[2.5rem] border border-outline-variant/40 p-8 sm:p-10 shadow-xl bg-white/95 space-y-6 animate-fade-in-up">
					<h3 class="text-xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-3 flex items-center gap-2">
						<Phone class="w-5 h-5 text-primary" />
						Student Contact Mobile Numbers
					</h3>

					<!-- Form to Add/Edit Mobile -->
					<form onsubmit={handleAddMobile} class="flex flex-col sm:flex-row gap-4 items-end bg-surface-container-low/50 p-5 rounded-2xl border border-outline-variant/30">
						<div class="space-y-1.5 flex-1 w-full">
							<label for="mobileNo" class="block text-xs font-bold uppercase text-on-surface-variant">Mobile Number</label>
							<input
								id="mobileNo"
								type="text"
								bind:value={newMobileNo}
								placeholder="01711111111"
								required
								class="w-full px-4 py-2.5 rounded-xl border border-outline-variant/50 text-sm bg-white"
							/>
						</div>

						<div class="space-y-1.5 w-full sm:w-48">
							<label for="ownerType" class="block text-xs font-bold uppercase text-on-surface-variant">Owner Type</label>
							<select
								id="ownerType"
								bind:value={newOwnerType}
								class="w-full px-4 py-2.5 rounded-xl border border-outline-variant/50 text-sm bg-white"
							>
								<option value="self">Self</option>
								<option value="father">Father</option>
								<option value="mother">Mother</option>
							</select>
						</div>

						<button
							type="submit"
							class="px-6 py-2.5 rounded-xl font-bold text-white bg-primary hover:bg-primary-container text-sm shadow-md flex items-center gap-2 whitespace-nowrap w-full sm:w-auto justify-center"
						>
							<Plus class="w-4 h-4" />
							{editingMobileNo ? 'Update' : 'Add Mobile'}
						</button>
					</form>

					<!-- Mobiles List -->
					<div class="space-y-3">
						{#if mobiles.length === 0}
							<p class="text-center py-6 text-xs font-semibold text-on-surface-variant">No mobile numbers saved yet.</p>
						{:else}
							{#each mobiles as m}
								<div class="flex items-center justify-between p-4 rounded-2xl border border-outline-variant/30 bg-white shadow-xs">
									<div class="flex items-center gap-3">
										<div class="w-10 h-10 rounded-xl bg-primary-fixed/40 text-primary flex items-center justify-center font-bold">
											<Phone class="w-5 h-5" />
										</div>
										<div>
											<p class="font-bold text-on-surface text-sm">{m.mobile_no}</p>
											<span class="text-xs uppercase font-extrabold text-tertiary bg-tertiary-fixed/30 px-2 py-0.5 rounded">Owner: {m.owner_type}</span>
										</div>
									</div>

									<div class="flex items-center gap-2">
										<button onclick={() => startEditMobile(m)} class="p-2 rounded-lg text-primary hover:bg-primary-fixed/40 transition-colors">
											<Edit3 class="w-4 h-4" />
										</button>
										<button onclick={() => handleDeleteMobile(m.mobile_no)} class="p-2 rounded-lg text-error hover:bg-error-container/40 transition-colors">
											<Trash2 class="w-4 h-4" />
										</button>
									</div>
								</div>
							{/each}
						{/if}
					</div>

					<!-- Navigation Buttons for Mobiles Tab -->
					<div class="pt-6 border-t border-outline-variant/30 flex items-center justify-between">
						<button
							type="button"
							onclick={() => activeTab = 'academics'}
							class="px-6 py-3 rounded-xl font-bold text-on-surface-variant hover:text-on-surface border border-outline-variant/50 hover:bg-surface-container transition-all flex items-center gap-2 text-sm"
						>
							<ArrowLeft class="w-4 h-4" />
							Previous: Marks & GPA
						</button>
						<button
							type="button"
							onclick={() => activeTab = 'address'}
							class="px-6 py-3 rounded-xl font-bold text-white bg-primary hover:bg-primary-container shadow-md transition-all flex items-center gap-2 text-sm"
						>
							Next: Address & Quota
							<ArrowRight class="w-4 h-4" />
						</button>
					</div>
				</div>
			{/if}

			<!-- Main Form Panel (Personal, Academics, Address) -->
			{#if activeTab !== 'mobiles'}
				<form onsubmit={handleSaveProfile} class="glass-panel rounded-[2.5rem] border border-outline-variant/40 p-8 sm:p-10 shadow-xl bg-white/95 space-y-8">
					
					{#if activeTab === 'personal'}
						<div class="space-y-6 animate-fade-in-up">
							<h3 class="text-xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-3 flex items-center gap-2">
								<User class="w-5 h-5 text-primary" />
								Guardian Information
							</h3>

							<div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
								<div class="space-y-1.5">
									<label for="fathersName" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Father's Name</label>
									<input
										id="fathersName"
										type="text"
										bind:value={fathersName}
										placeholder="e.g. Md. Anowar Hossain"
										class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all bg-white"
									/>
								</div>

								<div class="space-y-1.5">
									<label for="mothersName" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Mother's Name</label>
									<input
										id="mothersName"
										type="text"
										bind:value={mothersName}
										placeholder="e.g. Sultana Begum"
										class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all bg-white"
									/>
								</div>

								<div class="space-y-1.5">
									<label for="bloodGroup" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Blood Group</label>
									<select
										id="bloodGroup"
										bind:value={bloodGroup}
										class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
									>
										<option value="A+">A+</option>
										<option value="A-">A-</option>
										<option value="B+">B+</option>
										<option value="B-">B-</option>
										<option value="O+">O+</option>
										<option value="O-">O-</option>
										<option value="AB+">AB+</option>
										<option value="AB-">AB-</option>
									</select>
								</div>
							</div>

							<!-- Navigation Buttons for Guardian Info -->
							<div class="pt-6 border-t border-outline-variant/30 flex items-center justify-end">
								<button
									type="button"
									onclick={() => activeTab = 'academics'}
									class="px-6 py-3 rounded-xl font-bold text-white bg-primary hover:bg-primary-container shadow-md transition-all flex items-center gap-2 text-sm"
								>
									Next: Marks & GPA
									<ArrowRight class="w-4 h-4" />
								</button>
							</div>
						</div>
					{/if}

					{#if activeTab === 'academics'}
						<div class="space-y-6 animate-fade-in-up">
							<div class="flex items-center justify-between border-b border-outline-variant/30 pb-3">
								<h3 class="text-xl font-extrabold text-on-surface flex items-center gap-2">
									<BookOpen class="w-5 h-5 text-primary" />
									Academic History & Subject Marks
								</h3>
								<span class="text-xs font-extrabold text-error bg-error-container/60 px-3 py-1 rounded-full uppercase tracking-wider">Compulsory</span>
							</div>

							<!-- Shared Identification & Board -->
							<div class="grid grid-cols-1 sm:grid-cols-3 gap-6 p-4 rounded-2xl bg-surface-container-low/50 border border-outline-variant/30">
								<div class="space-y-1.5">
									<label for="regNo" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Registration No <span class="text-error">*</span></label>
									<input
										id="regNo"
										type="text"
										bind:value={regNo}
										placeholder="789012"
										required
										class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
									/>
								</div>

								<div class="space-y-1.5">
									<label for="board" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Education Board <span class="text-error">*</span></label>
									<select
										id="board"
										bind:value={board}
										class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
									>
										<option value="Dhaka">Dhaka</option>
										<option value="Chattogram">Chattogram</option>
										<option value="Rajshahi">Rajshahi</option>
										<option value="Sylhet">Sylhet</option>
										<option value="Barishal">Barishal</option>
										<option value="Dinajpur">Dinajpur</option>
										<option value="Comilla">Comilla</option>
										<option value="Jessore">Jessore</option>
										<option value="Mymensingh">Mymensingh</option>
									</select>
								</div>

								<div class="space-y-1.5">
									<label for="hscGroup" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">HSC Group <span class="text-error">*</span></label>
									<select
										id="hscGroup"
										bind:value={hscGroup}
										class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
									>
										<option value="Science">Science</option>
										<option value="Humanities">Humanities</option>
										<option value="Business">Business Studies / Commerce</option>
									</select>
								</div>
							</div>

							<!-- SSC Academic Details -->
							<div class="space-y-3">
								<span class="block text-xs font-extrabold text-on-surface uppercase tracking-wider">Secondary School Certificate (SSC) Details</span>
								<div class="grid grid-cols-1 sm:grid-cols-3 gap-6 p-4 rounded-2xl bg-white border border-outline-variant/30 shadow-xs">
									<div class="space-y-1.5">
										<label for="sscYear" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">SSC Exam Year <span class="text-error">*</span></label>
										<input
											id="sscYear"
											type="number"
											bind:value={sscYear}
											placeholder="2022"
											class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
										/>
									</div>
									<div class="space-y-1.5">
										<label for="sscRollNo" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">SSC Roll No <span class="text-error">*</span></label>
										<input
											id="sscRollNo"
											type="text"
											bind:value={sscRollNo}
											placeholder="123456"
											class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
										/>
									</div>
									<div class="space-y-1.5">
										<label for="sscGpa" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">SSC GPA <span class="text-error">*</span></label>
										<input
											id="sscGpa"
											type="text"
											bind:value={sscGpa}
											placeholder="5.00"
										required
											class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
										/>
									</div>
								</div>
							</div>

							<!-- HSC Academic Details -->
							<div class="space-y-3">
								<span class="block text-xs font-extrabold text-on-surface uppercase tracking-wider">Higher Secondary Certificate (HSC) Details</span>
								<div class="grid grid-cols-1 sm:grid-cols-3 gap-6 p-4 rounded-2xl bg-white border border-outline-variant/30 shadow-xs">
									<div class="space-y-1.5">
										<label for="hscYear" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">HSC Exam Year <span class="text-error">*</span></label>
										<input
											id="hscYear"
											type="number"
											bind:value={hscYear}
											placeholder="2024"
											class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
										/>
									</div>
									<div class="space-y-1.5">
										<label for="hscRollNo" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">HSC Roll No <span class="text-error">*</span></label>
										<input
											id="hscRollNo"
											type="text"
											bind:value={hscRollNo}
											placeholder="654321"
											class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
										/>
									</div>
									<div class="space-y-1.5">
										<label for="hscGpa" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">HSC GPA <span class="text-error">*</span></label>
										<input
											id="hscGpa"
											type="text"
											bind:value={hscGpa}
											placeholder="5.00"
										required
											class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
										/>
									</div>
								</div>
							</div>

							<!-- HSC Subject Marks -->
							<div class="pt-2 space-y-3">
								<span class="block text-xs font-extrabold text-on-surface uppercase tracking-wider">HSC Subject Marks (Out of 100) <span class="text-error">*</span></span>
								<div class="grid grid-cols-1 sm:grid-cols-3 gap-6 p-4 rounded-2xl bg-white border border-outline-variant/30 shadow-xs">
									<div class="space-y-1.5">
										<label for="physics" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Physics <span class="text-error">*</span></label>
										<input
											id="physics"
											type="number"
											bind:value={physicsMarks}
											placeholder="85"
											required
											class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
										/>
									</div>
									<div class="space-y-1.5">
										<label for="math" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Mathematics <span class="text-error">*</span></label>
										<input
											id="math"
											type="number"
											bind:value={mathMarks}
											placeholder="90"
											required
											class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
										/>
									</div>
									<div class="space-y-1.5">
										<label for="chemistry" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Chemistry <span class="text-error">*</span></label>
										<input
											id="chemistry"
											type="number"
											bind:value={chemistryMarks}
											placeholder="88"
											required
											class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
										/>
									</div>
								</div>
							</div>

							<!-- Navigation Buttons for Marks & GPA Tab -->
							<div class="pt-6 border-t border-outline-variant/30 flex items-center justify-between">
								<button
									type="button"
									onclick={() => activeTab = 'personal'}
									class="px-6 py-3 rounded-xl font-bold text-on-surface-variant hover:text-on-surface border border-outline-variant/50 hover:bg-surface-container transition-all flex items-center gap-2 text-sm"
								>
									<ArrowLeft class="w-4 h-4" />
									Previous: Guardian Info
								</button>
								<button
									type="button"
									onclick={() => activeTab = 'documents'}
									class="px-6 py-3 rounded-xl font-bold text-white bg-primary hover:bg-primary-container shadow-md transition-all flex items-center gap-2 text-sm"
								>
									Next: Photos & Docs
									<ArrowRight class="w-4 h-4" />
								</button>
							</div>
						</div>
					{/if}

					<!-- Tab: Photos & Required Documents (PHOTO_URL & SIGNATURE_URL) -->
					{#if activeTab === 'documents'}
						<div class="space-y-6 animate-fade-in-up">
							<div class="flex items-center justify-between border-b border-outline-variant/30 pb-3">
								<h3 class="text-xl font-extrabold text-on-surface flex items-center gap-2">
									<Camera class="w-5 h-5 text-primary" />
									Required Student Documents & Photos
								</h3>
							</div>

							<div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
								<!-- PHOTO_URL Input & Live Preview -->
								<div class="space-y-3 bg-surface-container-low/60 p-5 rounded-2xl border border-outline-variant/30">
									<div class="space-y-1.5">
										<label for="photoUrl" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant flex items-center gap-1.5">
											<Camera class="w-4 h-4 text-primary" />
											Student Photo URL (PHOTO_URL)
										</label>
										<input
											id="photoUrl"
											type="url"
											bind:value={photoUrl}
											placeholder="https://example.com/student-photo.jpg"
											class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
										/>
									</div>

									<!-- Live Photo Preview Box -->
									<div class="space-y-1">
										<span class="text-[11px] font-bold text-outline uppercase tracking-wider">Live Photo Preview</span>
										<div class="w-32 h-36 rounded-2xl border-2 border-dashed border-outline-variant/50 bg-white flex flex-col items-center justify-center overflow-hidden shadow-inner">
											{#if photoUrl && photoUrl.trim()}
												<img
													src={photoUrl}
													alt="Student Photo Preview"
													class="w-full h-full object-cover"
													onerror={() => toastState.error('Failed to load photo from specified URL')}
												/>
											{:else}
												<div class="text-center p-3 space-y-1">
													<User class="w-8 h-8 text-outline mx-auto" />
													<span class="text-[10px] font-bold text-outline block">No Photo URL</span>
												</div>
											{/if}
										</div>
									</div>
								</div>

								<!-- SIGNATURE_URL Input & Live Preview -->
								<div class="space-y-3 bg-surface-container-low/60 p-5 rounded-2xl border border-outline-variant/30">
									<div class="space-y-1.5">
										<label for="signatureUrl" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant flex items-center gap-1.5">
											<FileText class="w-4 h-4 text-primary" />
											Digital Signature URL (SIGNATURE_URL)
										</label>
										<input
											id="signatureUrl"
											type="url"
											bind:value={signatureUrl}
											placeholder="https://example.com/student-signature.png"
											class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
										/>
									</div>

									<!-- Live Signature Preview Box -->
									<div class="space-y-1">
										<span class="text-[11px] font-bold text-outline uppercase tracking-wider">Live Signature Preview</span>
										<div class="w-full h-24 rounded-2xl border-2 border-dashed border-outline-variant/50 bg-white flex flex-col items-center justify-center overflow-hidden p-2 shadow-inner">
											{#if signatureUrl && signatureUrl.trim()}
												<img
													src={signatureUrl}
													alt="Student Signature Preview"
													class="max-h-full object-contain"
													onerror={() => toastState.error('Failed to load signature from specified URL')}
												/>
											{:else}
												<div class="text-center p-2 space-y-1">
													<FileText class="w-6 h-6 text-outline mx-auto" />
													<span class="text-[10px] font-bold text-outline block">No Signature URL</span>
												</div>
											{/if}
										</div>
									</div>
								</div>
							</div>

							<!-- Navigation & Save Buttons for Documents Tab -->
							<div class="pt-6 border-t border-outline-variant/30 flex flex-wrap items-center justify-between gap-3">
								<button
									type="button"
									onclick={() => activeTab = 'academics'}
									class="px-6 py-3 rounded-xl font-bold text-on-surface-variant hover:text-on-surface border border-outline-variant/50 hover:bg-surface-container transition-all flex items-center gap-2 text-sm"
								>
									<ArrowLeft class="w-4 h-4" />
									Previous: Marks & GPA
								</button>
								<div class="flex gap-2">
									<button
										type="button"
										onclick={() => activeTab = 'mobiles'}
										class="px-6 py-3 rounded-xl font-bold text-on-surface-variant hover:text-on-surface border border-outline-variant/50 hover:bg-surface-container transition-all flex items-center gap-2 text-sm"
									>
										Next: Mobile Numbers
										<ArrowRight class="w-4 h-4" />
									</button>
								</div>
							</div>
						</div>
					{/if}

					{#if activeTab === 'address'}
						<div class="space-y-6 animate-fade-in-up">
							<h3 class="text-xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-3 flex items-center gap-2">
								<MapPin class="w-5 h-5 text-primary" />
								Address & Special Quotas
							</h3>

							<div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
								<div class="space-y-1.5">
									<label for="quota" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Admission Quota</label>
									<select
										id="quota"
										bind:value={quota}
										class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
									>
										<option value="GENERAL">General</option>
										<option value="FREEDOM_FIGHTER">Freedom Fighter Quota</option>
										<option value="TRIBAL">Tribal / Ethnic Quota</option>
										<option value="PHYSICALLY_CHALLENGED">Physically Challenged</option>
									</select>
								</div>

								<div class="space-y-1.5 sm:col-span-2">
									<label for="presentAddress" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Present Address</label>
									<textarea
										id="presentAddress"
										bind:value={presentAddress}
										rows="3"
										placeholder="House, Road, Thana, District"
										class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
									></textarea>
								</div>

								<div class="space-y-1.5 sm:col-span-2">
									<label for="permanentAddress" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Permanent Address</label>
									<textarea
										id="permanentAddress"
										bind:value={permanentAddress}
										rows="3"
										placeholder="House, Village, Upazila, District"
										class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
									></textarea>
								</div>
							</div>

							<!-- Navigation & Save Action Bar for Address & Quota (Last Tab) -->
							<div class="pt-6 border-t border-outline-variant/30 flex items-center justify-between">
								<button
									type="button"
									onclick={() => activeTab = 'mobiles'}
									class="px-6 py-3 rounded-xl font-bold text-on-surface-variant hover:text-on-surface border border-outline-variant/50 hover:bg-surface-container transition-all flex items-center gap-2 text-sm"
								>
									<ArrowLeft class="w-4 h-4" />
									Previous: Mobile Numbers
								</button>
								<button
									type="submit"
									disabled={isLoading}
									class="px-8 py-3.5 rounded-xl font-bold text-white bg-primary hover:bg-primary-container disabled:opacity-50 shadow-lg shadow-primary/25 hover:shadow-primary/40 transition-all flex items-center gap-2 text-sm"
								>
									<Save class="w-4 h-4" />
									{isLoading ? 'Saving...' : 'Save Academic Profile'}
								</button>
							</div>
						</div>
					{/if}
				</form>
			{/if}
		{/if}
	</div>
</div>
