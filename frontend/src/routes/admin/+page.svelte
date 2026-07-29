<!-- src/routes/admin/+page.svelte -->
<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import {
		adminLogin,
		fetchUniversityApplications,
		updateApplicationStatus,
		createUniversity,
		updateUniversity,
		deleteUniversity,
		createProgram,
		updateProgram,
		deleteProgram,
		createAdmissionTest,
		updateAdmissionTest,
		publishAdmissionTestResults,
		fetchProgramEligibilityRules,
		saveProgramEligibilityRule,
		deleteProgramEligibilityRule,
		createUniversityTransport,
		updateUniversityTransport,
		deleteUniversityTransport
	} from '$lib/api/admin';
	import { fetchUniversities, fetchUniversityByID, fetchUniversityTransport } from '$lib/api/university';
	import { fetchPrograms, fetchProgramByID } from '$lib/api/programs';
	import { authState } from '$lib/state/auth.svelte';
	import { getAdminToken } from '$lib/api/client';
	import { toastState } from '$lib/state/toast.svelte';
	import type {
		University,
		Program,
		StudentApplication,
		ProgramEligibilityRule,
		UniversityTransport,
		RequiredProfileField
	} from '$lib/types/models';
	import {
		Shield,
		Building2,
		BookOpen,
		FileCheck,
		LogOut,
		Plus,
		Trash2,
		Edit3,
		CheckCircle2,
		XCircle,
		AlertCircle,
		Bus,
		Award,
		Send,
		Clock,
		Image as ImageIcon,
		Users,
		Eye,
		EyeOff,
		X
	} from 'lucide-svelte';

	let isAdminLoggedIn = $derived(authState.isAdmin);
	let adminEmail = $state('');
	let adminPassword = $state('');
	let showAdminPassword = $state(false);

	let activeTab = $state<'universities' | 'programs' | 'applications' | 'eligibility' | 'transport' | 'publish' | 'admissiontests'>('universities');
	let isLoading = $state(false);
	let errorMessage = $state<string | null>(null);

	// Data
	let universities = $state<University[]>([]);
	let programs = $state<Program[]>([]);
	let applications = $state<StudentApplication[]>([]);
	let selectedUniversityID = $state<number>(1);

	// Eligibility Rules
	let selectedProgramID = $state<number>(1);
	let eligibilityRules = $state<ProgramEligibilityRule[]>([]);
	let newRuleType = $state('MIN_HSC_PHYSICS');
	let newRuleValue = $state('80.00');

	// Transport Routes
	let transportUniversityID = $state<number>(1);
	let transportRoutes = $state<UniversityTransport[]>([]);
	let newRouteName = $state('');
	let newRouteTime = $state('45 mins');
	let editingRouteName = $state<string | null>(null);

	// Admission Test Results Publish Batch
	let testIDToPublish = $state<number>(1);
	let publishResultsList = $state<{ student_id: number; marks: string; merit_position: number }[]>([
		{ student_id: 1, marks: '85.50', merit_position: 1 }
	]);

	function addPublishRow() {
		const nextPos = publishResultsList.length + 1;
		publishResultsList = [...publishResultsList, { student_id: nextPos, marks: '80.00', merit_position: nextPos }];
	}

	function removePublishRow(idx: number) {
		if (publishResultsList.length <= 1) return;
		publishResultsList = publishResultsList.filter((_, i) => i !== idx);
	}

	// Admission Test Management
	let atExamUnit = $state('A');
	let atExamCenter = $state('');
	let atExamDate = $state('2026-11-15');
	let atPrereqTestId = $state<number | null>(null);
	let atProgramId = $state<number>(1);
	let editingTestId = $state<number | null>(null);

	async function handleCreateAdmissionTest(e: Event) {
		e.preventDefault();
		isLoading = true;
		try {
			const payload = {
				exam_unit: atExamUnit,
				exam_center: atExamCenter,
				exam_date: atExamDate,
				prereq_test_id: atPrereqTestId || null,
				program_id: Number(atProgramId)
			};
			if (editingTestId) {
				await updateAdmissionTest(editingTestId, payload);
				editingTestId = null;
			} else {
				await createAdmissionTest(payload);
			}
			atExamUnit = 'A';
			atExamCenter = '';
			atExamDate = '2026-11-15';
			atPrereqTestId = null;
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to save admission test.');
		} finally {
			isLoading = false;
		}
	}

	// University Modal State (Supports Departments, History, Description & Campus Album)
	let showUniModal = $state(false);
	let editingUniId = $state<number | null>(null);
	let uniName = $state('');
	let uniWebsite = $state('');
	let uniLocation = $state('');
	let uniLogoUrl = $state('');
	let uniDescription = $state('');
	let uniHistory = $state('');
	let uniDepartments = $state<{ dept_name: string; dept_description: string; total_seats: number }[]>([]);
	let uniAlbum = $state<{ picture_title: string; picture_url: string }[]>([]);

	function addDepartmentRow() {
		uniDepartments = [...uniDepartments, { dept_name: '', dept_description: '', total_seats: 60 }];
	}

	function removeDepartmentRow(idx: number) {
		uniDepartments = uniDepartments.filter((_, i) => i !== idx);
	}

	function addAlbumRow() {
		uniAlbum = [...uniAlbum, { picture_title: '', picture_url: '' }];
	}

	function removeAlbumRow(idx: number) {
		uniAlbum = uniAlbum.filter((_, i) => i !== idx);
	}

	// Program Modal State
	let showProgModal = $state(false);
	let editingProgId = $state<number | null>(null);
	let progName = $state('');
	let progUnit = $state('A');
	let progSeats = $state<number>(100);
	let progCutmarks = $state<number>(80.0);
	let progDeadline = $state('2026-12-31');
	let progFee = $state<number>(500);
	let progUniId = $state<number>(1);
	let progRequiredFields = $state<RequiredProfileField[]>([
		'PRESENT_ADDRESS',
		'PERMANENT_ADDRESS',
		'FATHERS_NAME',
		'MOTHERS_NAME',
		'BLOOD_GROUP',
		'QUOTA',
		'PHOTO_URL',
		'SIGNATURE_URL'
	]);

	const availableRequiredFields: RequiredProfileField[] = [
		'PRESENT_ADDRESS',
		'PERMANENT_ADDRESS',
		'FATHERS_NAME',
		'MOTHERS_NAME',
		'BLOOD_GROUP',
		'QUOTA',
		'PHOTO_URL',
		'SIGNATURE_URL'
	];

	function toggleRequiredField(field: RequiredProfileField) {
		if (progRequiredFields.includes(field)) {
			progRequiredFields = progRequiredFields.filter(f => f !== field);
		} else {
			progRequiredFields = [...progRequiredFields, field];
		}
	}

	async function handleAdminLogin(e: Event) {
		e.preventDefault();
		isLoading = true;
		errorMessage = null;
		try {
			const res = await adminLogin({ email: adminEmail, password: adminPassword });
			if (res.token) {
				authState.setAdmin(res.token);
				isAdminLoggedIn = true;
				await loadAllData();
			}
		} catch (err: any) {
			errorMessage = err?.message || 'Admin login failed. Check credentials.';
		} finally {
			isLoading = false;
		}
	}

	function handleAdminLogout() {
		authState.logout();
		adminEmail = '';
		adminPassword = '';
		showAdminPassword = false;
		universities = [];
		programs = [];
		applications = [];
		goto('/admin');
	}

	onMount(() => {
		if (getAdminToken()) {
			isAdminLoggedIn = true;
			loadAllData();
		}
	});

	async function loadAllData() {
		isLoading = true;
		errorMessage = null;
		try {
			const [uList, pList] = await Promise.all([fetchUniversities(), fetchPrograms()]);
			universities = uList || [];
			programs = pList || [];
			if (universities.length > 0) {
				selectedUniversityID = universities[0].u_id;
				transportUniversityID = universities[0].u_id;
			}
			if (programs.length > 0) {
				selectedProgramID = programs[0].program_id;
			}
		} catch (err: any) {
			console.error(err);
		} finally {
			isLoading = false;
		}
	}

	function switchTab(tab: typeof activeTab) {
		activeTab = tab;
		if (tab === 'applications' && selectedUniversityID) {
			loadApplications(selectedUniversityID);
		} else if (tab === 'eligibility' && selectedProgramID) {
			loadEligibilityRules(selectedProgramID);
		} else if (tab === 'transport' && transportUniversityID) {
			loadTransport(transportUniversityID);
		}
	}

	async function loadApplications(uId: number) {
		try {
			applications = await fetchUniversityApplications(uId);
		} catch (err: any) {
			applications = [];
		}
	}

	async function loadEligibilityRules(pId: number) {
		try {
			eligibilityRules = await fetchProgramEligibilityRules(pId);
		} catch (err: any) {
			eligibilityRules = [];
		}
	}

	async function loadTransport(uId: number) {
		try {
			transportRoutes = await fetchUniversityTransport(uId);
		} catch (err: any) {
			transportRoutes = [];
		}
	}

	// University Actions
	function openCreateUniModal() {
		editingUniId = null;
		uniName = '';
		uniWebsite = '';
		uniLocation = '';
		uniLogoUrl = '';
		uniDescription = '';
		uniHistory = '';
		uniDepartments = [{ dept_name: 'CSE', dept_description: 'Computer Science and Engineering', total_seats: 120 }];
		uniAlbum = [];
		showUniModal = true;
	}

	async function openEditUniModal(u: University) {
		editingUniId = u.u_id;
		uniName = u.u_name;
		uniWebsite = u.website;
		uniLocation = u.location || '';
		uniLogoUrl = u.logo_url || '';
		uniDescription = u.university_description || '';
		uniHistory = u.university_history || '';
		uniDepartments = u.departments ? u.departments.map(d => ({ dept_name: d.dept_name, dept_description: d.dept_description, total_seats: d.total_seats })) : [];
		uniAlbum = u.album ? u.album.map(a => ({ picture_title: a.picture_title, picture_url: a.picture_url })) : [];
		showUniModal = true;

		try {
			const fullUni = await fetchUniversityByID(u.u_id);
			if (fullUni) {
				uniName = fullUni.u_name || uniName;
				uniWebsite = fullUni.website || uniWebsite;
				uniLocation = fullUni.location || uniLocation;
				uniLogoUrl = fullUni.logo_url || uniLogoUrl;
				uniDescription = fullUni.university_description || uniDescription;
				uniHistory = fullUni.university_history || uniHistory;
				if (fullUni.departments && fullUni.departments.length > 0) {
					uniDepartments = fullUni.departments.map(d => ({
						dept_name: d.dept_name,
						dept_description: d.dept_description,
						total_seats: d.total_seats
					}));
				}
				if (fullUni.album && fullUni.album.length > 0) {
					uniAlbum = fullUni.album.map(a => ({
						picture_title: a.picture_title,
						picture_url: a.picture_url
					}));
				}
			}
		} catch (err) {
			console.error('Failed to load full university details:', err);
		}
	}

	async function handleSaveUniversity(e: Event) {
		e.preventDefault();
		isLoading = true;
		try {
			const payload = {
				name: uniName,
				website: uniWebsite,
				location: uniLocation,
				logo_url: uniLogoUrl,
				university_description: uniDescription,
				university_history: uniHistory,
				departments: uniDepartments.filter(d => d.dept_name.trim() !== ''),
				album: uniAlbum.filter(a => a.picture_title.trim() !== '' && a.picture_url.trim() !== '')
			};
			if (editingUniId) {
				await updateUniversity(editingUniId, payload);
			} else {
				await createUniversity(payload);
			}
			showUniModal = false;
			await loadAllData();
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to save university.');
		} finally {
			isLoading = false;
		}
	}

	async function handleDeleteUniversity(uId: number) {
		if (!confirm('Are you sure you want to delete this university?')) return;
		isLoading = true;
		try {
			await deleteUniversity(uId);
			await loadAllData();
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to delete university.');
		} finally {
			isLoading = false;
		}
	}

	// Program Actions
	function openCreateProgModal() {
		editingProgId = null;
		progName = '';
		progUnit = 'A';
		progSeats = 100;
		progCutmarks = 80.0;
		progDeadline = '2026-12-31';
		progFee = 500;
		progRequiredFields = ['PRESENT_ADDRESS', 'PERMANENT_ADDRESS', 'FATHERS_NAME', 'MOTHERS_NAME', 'BLOOD_GROUP', 'QUOTA', 'PHOTO_URL', 'SIGNATURE_URL'];
		progUniId = universities.length > 0 ? universities[0].u_id : 1;
		showProgModal = true;
	}

	async function openEditProgModal(p: Program) {
		editingProgId = p.program_id;
		progName = p.p_name;
		progUnit = p.p_unit || 'A';
		progSeats = p.total_seats;
		progCutmarks = typeof p.prev_cutmarks === 'number' ? p.prev_cutmarks : parseFloat(String(p.prev_cutmarks) || '80');
		progDeadline = p.deadline ? p.deadline.split('T')[0] : '2026-12-31';

		const parseFee = (fee: any): number => {
			if (fee === undefined || fee === null || fee === '') return 500;
			if (typeof fee === 'number') return isNaN(fee) ? 500 : fee;
			if (typeof fee === 'string') {
				const n = parseFloat(fee);
				return isNaN(n) ? 500 : n;
			}
			if (typeof fee === 'object' && fee.String) {
				const n = parseFloat(fee.String);
				return isNaN(n) ? 500 : n;
			}
			return 500;
		};

		progFee = parseFee(p.application_fee);
		progRequiredFields = p.required_fields || [
			'PRESENT_ADDRESS',
			'PERMANENT_ADDRESS',
			'FATHERS_NAME',
			'MOTHERS_NAME',
			'BLOOD_GROUP',
			'QUOTA',
			'PHOTO_URL',
			'SIGNATURE_URL'
		];
		progUniId = p.u_id;
		showProgModal = true;

		try {
			const fullProg = await fetchProgramByID(p.program_id);
			if (fullProg) {
				progName = fullProg.p_name || progName;
				progUnit = fullProg.p_unit || progUnit;
				progSeats = fullProg.total_seats || progSeats;
				progCutmarks = typeof fullProg.prev_cutmarks === 'number' ? fullProg.prev_cutmarks : parseFloat(String(fullProg.prev_cutmarks) || '80');
				progDeadline = fullProg.deadline ? fullProg.deadline.split('T')[0] : progDeadline;
				progFee = parseFee(fullProg.application_fee);
				if (fullProg.required_fields && fullProg.required_fields.length > 0) {
					progRequiredFields = fullProg.required_fields;
				}
				progUniId = fullProg.u_id || progUniId;
			}
		} catch (err) {
			console.error('Failed to load full program details:', err);
		}
	}

	async function handleSaveProgram(e: Event) {
		e.preventDefault();
		isLoading = true;
		try {
			const payload = {
				p_name: progName,
				p_unit: progUnit,
				total_seats: Number(progSeats),
				prev_cutmarks: Number(progCutmarks),
				deadline: progDeadline,
				u_id: Number(progUniId),
				application_fee: Number(progFee),
				required_fields: progRequiredFields
			};
			if (editingProgId) {
				await updateProgram(editingProgId, payload);
			} else {
				await createProgram(payload);
			}
			showProgModal = false;
			await loadAllData();
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to save program.');
		} finally {
			isLoading = false;
		}
	}

	async function handleDeleteProgram(pId: number) {
		if (!confirm('Are you sure you want to delete this program?')) return;
		isLoading = true;
		try {
			await deleteProgram(pId);
			await loadAllData();
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to delete program.');
		} finally {
			isLoading = false;
		}
	}

	// Eligibility Rule Actions
	async function handleAddEligibilityRule(e: Event) {
		e.preventDefault();
		isLoading = true;
		try {
			await saveProgramEligibilityRule({
				program_id: selectedProgramID,
				rule_type: newRuleType,
				rule_value: newRuleValue
			});
			await loadEligibilityRules(selectedProgramID);
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to save eligibility rule.');
		} finally {
			isLoading = false;
		}
	}

	async function handleDeleteRule(ruleType: string) {
		try {
			await deleteProgramEligibilityRule(selectedProgramID, ruleType);
			await loadEligibilityRules(selectedProgramID);
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to delete eligibility rule.');
		}
	}

	// Transport Actions
	function startEditTransport(r: UniversityTransport) {
		editingRouteName = r.transport_route;
		newRouteName = r.transport_route;
		newRouteTime = r.est_travel_time;
	}

	async function handleSaveTransport(e: Event) {
		e.preventDefault();
		if (!newRouteName.trim()) return;
		isLoading = true;
		try {
			if (editingRouteName) {
				await updateUniversityTransport({
					u_id: transportUniversityID,
					transport_route: newRouteName,
					est_travel_time: newRouteTime
				});
				editingRouteName = null;
			} else {
				await createUniversityTransport({
					u_id: transportUniversityID,
					transport_route: newRouteName,
					est_travel_time: newRouteTime
				});
			}
			newRouteName = '';
			await loadTransport(transportUniversityID);
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to save transport route.');
		} finally {
			isLoading = false;
		}
	}

	async function handleDeleteTransportRoute(routeStr: string) {
		try {
			await deleteUniversityTransport(transportUniversityID, routeStr);
			await loadTransport(transportUniversityID);
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to delete transport route.');
		}
	}

	// Publish Exam Results Batch
	async function handlePublishResults(e: Event) {
		e.preventDefault();
		if (publishResultsList.length === 0) return;
		isLoading = true;
		try {
			await publishAdmissionTestResults({
				test_id: Number(testIDToPublish),
				results: publishResultsList.map((r) => ({
					student_id: Number(r.student_id),
					marks: String(r.marks),
					merit_position: Number(r.merit_position)
				}))
			});
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to publish exam results.');
		} finally {
			isLoading = false;
		}
	}

	// Application Status Update
	async function handleUpdateStatus(appId: number, status: string) {
		isLoading = true;
		try {
			await updateApplicationStatus(appId, status);
			await loadApplications(selectedUniversityID);
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to update application status.');
		} finally {
			isLoading = false;
		}
	}
</script>

<svelte:head>
	<title>Admin Portal - UniApp System</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		
		{#if !isAdminLoggedIn}
			<!-- Admin Login Screen -->
			<div class="max-w-md mx-auto glass-panel p-8 sm:p-10 rounded-[2.5rem] border border-outline-variant/40 shadow-2xl bg-white/95 space-y-6">
				<div class="text-center space-y-2">
					<div class="w-14 h-14 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center mx-auto shadow-md">
						<Shield class="w-8 h-8" />
					</div>
					<h2 class="text-2xl font-extrabold text-on-surface">System Admin Login</h2>
					<p class="text-xs text-on-surface-variant">Sign in with system administrator credentials</p>
				</div>

				{#if errorMessage}
					<div class="p-4 rounded-2xl bg-error-container/60 border border-error/30 flex items-center gap-3 text-on-error-container text-xs">
						<AlertCircle class="w-4 h-4 shrink-0 text-error" />
						<span>{errorMessage}</span>
					</div>
				{/if}

				<form onsubmit={handleAdminLogin} class="space-y-4">
					<div class="space-y-1">
						<label for="adminEmail" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Admin Email</label>
						<input
							id="adminEmail"
							type="email"
							bind:value={adminEmail}
							placeholder="admin@gmail.com"
							required
							class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
						/>
					</div>

					<div class="space-y-1">
						<label for="adminPass" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Admin Password</label>
						<div class="relative">
							<input
								id="adminPass"
								type={showAdminPassword ? 'text' : 'password'}
								bind:value={adminPassword}
								placeholder="Enter admin password"
								required
								class="w-full px-4 py-3 pr-11 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
							/>
							<button
								type="button"
								onclick={() => showAdminPassword = !showAdminPassword}
								class="absolute right-3.5 top-1/2 -translate-y-1/2 text-outline hover:text-on-surface transition-colors p-1"
								title={showAdminPassword ? 'Hide Password' : 'Show Password'}
							>
								{#if showAdminPassword}
									<EyeOff class="w-4 h-4" />
								{:else}
									<Eye class="w-4 h-4" />
								{/if}
							</button>
						</div>
					</div>

					<button
						type="submit"
						disabled={isLoading}
						class="w-full py-3.5 px-6 rounded-xl font-bold text-white bg-primary hover:bg-primary-container disabled:opacity-50 shadow-lg shadow-primary/25 transition-all text-sm mt-2"
					>
						{isLoading ? 'Authenticating...' : 'Sign In to Admin Dashboard'}
					</button>
				</form>
			</div>
		{:else}
			<!-- Admin Dashboard Banner Header -->
			<div class="flex flex-col md:flex-row justify-between items-start md:items-end gap-6 bg-gradient-to-r from-primary via-primary-container to-secondary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl relative overflow-hidden">
				<div class="relative z-10 max-w-2xl space-y-2">
					<div class="inline-flex items-center gap-2 px-3.5 py-1 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
						<Shield class="w-4 h-4" />
						Admin Control Panel
					</div>
					<h1 class="text-3xl sm:text-4xl font-black text-white">University & Academic System Admin</h1>
					<p class="text-slate-100 text-sm">Manage universities, departments, album photos, programs, cutoff rules, transport routes, and publish test results.</p>
				</div>

				<button
					onclick={handleAdminLogout}
					class="relative z-10 px-5 py-2.5 rounded-xl text-xs font-bold bg-white/15 hover:bg-white/25 text-white border border-white/20 backdrop-blur-md transition-all flex items-center gap-2"
				>
					<LogOut class="w-4 h-4" />
					Logout Admin
				</button>
			</div>

			<!-- Navigation Tabs -->
			<div class="flex items-center gap-2 bg-white/80 p-2 rounded-2xl border border-outline-variant/30 shadow-sm overflow-x-auto">
				<button
					onclick={() => switchTab('universities')}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'universities' ? 'bg-primary text-white shadow-md' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<Building2 class="w-4 h-4" /> Universities & Depts
				</button>
				<button
					onclick={() => switchTab('programs')}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'programs' ? 'bg-primary text-white shadow-md' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<BookOpen class="w-4 h-4" /> Programs
				</button>
				<button
					onclick={() => switchTab('eligibility')}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'eligibility' ? 'bg-primary text-white shadow-md' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<Award class="w-4 h-4" /> Cutoff Rules
				</button>
				<button
					onclick={() => switchTab('transport')}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'transport' ? 'bg-primary text-white shadow-md' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<Bus class="w-4 h-4" /> Transport Routes
				</button>
				<button
					onclick={() => switchTab('admissiontests')}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'admissiontests' ? 'bg-primary text-white shadow-md' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<Clock class="w-4 h-4" /> Admission Tests
				</button>
				<button
					onclick={() => switchTab('publish')}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'publish' ? 'bg-primary text-white shadow-md' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<Send class="w-4 h-4" /> Publish Results
				</button>
				<button
					onclick={() => switchTab('applications')}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'applications' ? 'bg-primary text-white shadow-md' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<FileCheck class="w-4 h-4" /> Applications
				</button>
			</div>

			<!-- Tab 1: Universities & Departments Management -->
			{#if activeTab === 'universities'}
				<div class="space-y-6">
					<div class="flex items-center justify-between">
						<div>
							<h3 class="text-xl font-extrabold text-on-surface">Registered Universities</h3>
							<p class="text-xs text-on-surface-variant">Manage university profiles, departments, total seat allocations, and campus photo albums.</p>
						</div>
						<button
							onclick={openCreateUniModal}
							class="px-5 py-2.5 rounded-xl text-xs font-bold text-white bg-primary hover:bg-primary-container shadow-md flex items-center gap-2"
						>
							<Plus class="w-4 h-4" /> Add University
						</button>
					</div>

					<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
						{#each universities as u}
							<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/40 bg-white/95 shadow-md flex flex-col justify-between space-y-4">
								<div class="space-y-3">
									<div class="flex items-start justify-between gap-3">
										<div class="flex items-center gap-3">
											{#if u.logo_url}
												<img src={u.logo_url} alt={u.u_name} class="w-12 h-12 rounded-2xl object-cover border shrink-0" />
											{:else}
												<div class="w-12 h-12 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center font-black text-xl shrink-0">
													{u.u_name.charAt(0)}
												</div>
											{/if}
											<div>
												<h4 class="text-lg font-extrabold text-on-surface leading-tight">{u.u_name}</h4>
												<p class="text-xs font-semibold text-on-surface-variant">{u.location || 'Location Not Specified'}</p>
											</div>
										</div>

										<div class="flex gap-1 shrink-0">
											<button onclick={() => openEditUniModal(u)} class="p-1.5 rounded-lg text-primary hover:bg-primary-fixed/40 transition-colors">
												<Edit3 class="w-4 h-4" />
											</button>
											<button onclick={() => handleDeleteUniversity(u.u_id)} class="p-1.5 rounded-lg text-error hover:bg-error-container/40 transition-colors">
												<Trash2 class="w-4 h-4" />
											</button>
										</div>
									</div>

									{#if u.website}
										<a href={u.website} target="_blank" class="text-xs font-semibold text-primary hover:underline block truncate">{u.website}</a>
									{/if}

									{#if u.university_description}
										<div class="space-y-0.5 text-xs text-on-surface-variant line-clamp-2 leading-relaxed">
											{@html u.university_description}
										</div>
									{/if}

									{#if u.university_history}
										<div class="space-y-0.5">
											<span class="text-[10px] font-bold text-outline uppercase tracking-wider block">History</span>
											<div class="text-xs text-on-surface-variant line-clamp-2 leading-relaxed">{@html u.university_history}</div>
										</div>
									{/if}
								</div>
							</div>
						{/each}
					</div>
				</div>
			{/if}

			<!-- Tab 2: Programs Management -->
			{#if activeTab === 'programs'}
				<div class="space-y-6">
					<div class="flex items-center justify-between">
						<div>
							<h3 class="text-xl font-extrabold text-on-surface">Admission Programs</h3>
							<p class="text-xs text-on-surface-variant">Manage academic degree programs, unit allocations, seat capacities, and cutmarks.</p>
						</div>
						<button
							onclick={openCreateProgModal}
							class="px-5 py-2.5 rounded-xl text-xs font-bold text-white bg-primary hover:bg-primary-container shadow-md flex items-center gap-2"
						>
							<Plus class="w-4 h-4" /> Add Program
						</button>
					</div>

					<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
						{#each programs as p}
							<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/40 bg-white/95 shadow-md flex flex-col justify-between space-y-4">
								<div class="space-y-3">
									<div class="flex items-center justify-between">
										<span class="text-xs font-mono font-bold text-outline">ID #{p.program_id}</span>
										<div class="flex gap-1.5">
											<button onclick={() => openEditProgModal(p)} class="p-1.5 rounded-lg text-primary hover:bg-primary-fixed/40 transition-colors">
												<Edit3 class="w-4 h-4" />
											</button>
											<button onclick={() => handleDeleteProgram(p.program_id)} class="p-1.5 rounded-lg text-error hover:bg-error-container/40 transition-colors">
												<Trash2 class="w-4 h-4" />
											</button>
										</div>
									</div>

									<div>
										<h4 class="text-xl font-black text-on-surface leading-tight">{p.p_name}</h4>
										<p class="text-xs font-bold text-primary mt-0.5">{p.university_name || p.u_name || 'University ID #' + p.u_id}</p>
									</div>

									<div class="grid grid-cols-2 gap-2 text-xs font-semibold bg-surface-container-low/60 p-3 rounded-xl border border-outline-variant/30">
										<div>
											<span class="text-[10px] text-outline font-bold uppercase block">Unit & Seats</span>
											<span class="text-on-surface">Unit {p.p_unit || 'A'} • {p.total_seats} seats</span>
										</div>
										<div>
											<span class="text-[10px] text-outline font-bold uppercase block">App Fee</span>
											<span class="text-emerald-700 font-extrabold">{p.application_fee ? `৳${p.application_fee}` : '৳500.00'}</span>
										</div>
										<div>
											<span class="text-[10px] text-outline font-bold uppercase block">Cutmarks</span>
											<span class="text-on-surface">{p.prev_cutmarks || '80.00'}</span>
										</div>
										<div>
											<span class="text-[10px] text-outline font-bold uppercase block">Deadline</span>
											<span class="text-on-surface">{p.deadline ? p.deadline.split('T')[0] : 'N/A'}</span>
										</div>
									</div>

									<!-- Required Attributes Badges -->
									{#if p.required_fields && p.required_fields.length > 0}
										<div class="space-y-1 pt-1">
											<span class="text-[10px] font-bold text-outline uppercase tracking-wider block">Required Documents & Info:</span>
											<div class="flex flex-wrap gap-1">
												{#each p.required_fields as rf}
													<span class="text-[10px] font-mono font-bold px-2 py-0.5 bg-primary-fixed/40 text-primary rounded-md border border-primary/20">
														{rf}
													</span>
												{/each}
											</div>
										</div>
									{/if}
								</div>
							</div>
						{/each}
					</div>
				</div>
			{/if}

			<!-- Tab 3: Eligibility Rules -->
			{#if activeTab === 'eligibility'}
				<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6 max-w-2xl mx-auto">
					<h3 class="text-xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-3 flex items-center gap-2">
						<Award class="w-5 h-5 text-primary" />
						Program Eligibility Rules
					</h3>

					<div class="space-y-1.5">
						<label for="selectProgRule" class="block text-xs font-bold uppercase text-on-surface-variant">Select Program</label>
						<select
							id="selectProgRule"
							bind:value={selectedProgramID}
							onchange={() => loadEligibilityRules(selectedProgramID)}
							class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 text-sm bg-white font-bold text-on-surface"
						>
							{#each programs as p}
								<option value={p.program_id}>{p.p_name} (Unit {p.p_unit || 'A'})</option>
							{/each}
						</select>
					</div>

					<form onsubmit={handleAddEligibilityRule} class="flex flex-col sm:flex-row gap-3 items-end bg-surface-container-low/50 p-4 rounded-2xl border border-outline-variant/30">
						<div class="space-y-1 flex-1 w-full">
							<label for="ruleType" class="block text-xs font-bold uppercase text-on-surface-variant">Rule Type</label>
							<select id="ruleType" bind:value={newRuleType} class="w-full px-3.5 py-2.5 rounded-xl border text-xs bg-white font-bold">
								<option value="MIN_HSC_PHYSICS">MIN_HSC_PHYSICS</option>
								<option value="MIN_HSC_MATH">MIN_HSC_MATH</option>
								<option value="MIN_HSC_CHEMISTRY">MIN_HSC_CHEMISTRY</option>
								<option value="MIN_HSC_GPA">MIN_HSC_GPA</option>
								<option value="MIN_SSC_GPA">MIN_SSC_GPA</option>
							</select>
						</div>

						<div class="space-y-1 w-full sm:w-36">
							<label for="ruleVal" class="block text-xs font-bold uppercase text-on-surface-variant">Value</label>
							<input id="ruleVal" type="text" bind:value={newRuleValue} placeholder="80.00" required class="w-full px-3.5 py-2.5 rounded-xl border text-xs bg-white font-bold" />
						</div>

						<button type="submit" class="px-5 py-2.5 rounded-xl font-bold text-white bg-primary text-xs shadow-md flex items-center gap-1 shrink-0">
							<Plus class="w-3.5 h-3.5" /> Save Rule
						</button>
					</form>

					<div class="space-y-3">
						<h4 class="text-xs font-extrabold uppercase tracking-wider text-on-surface-variant">Active Program Rules</h4>
						{#if eligibilityRules.length === 0}
							<p class="text-xs text-on-surface-variant text-center py-4">No eligibility rules configured for this program.</p>
						{:else}
							{#each eligibilityRules as r}
								<div class="flex items-center justify-between p-3.5 rounded-xl border border-outline-variant/30 bg-white">
									<span class="text-xs font-bold text-on-surface font-mono">{r.rule_type}</span>
									<div class="flex items-center gap-3">
										<span class="text-xs font-bold text-primary px-2.5 py-0.5 rounded-md bg-primary-fixed">{typeof r.rule_value === 'object' ? r.rule_value.String : r.rule_value}</span>
										<button onclick={() => handleDeleteRule(r.rule_type)} class="text-error hover:bg-error-container/40 p-1.5 rounded-lg transition-colors">
											<Trash2 class="w-4 h-4" />
										</button>
									</div>
								</div>
							{/each}
						{/if}
					</div>
				</div>
			{/if}

			<!-- Tab 4: Transport Routes -->
			{#if activeTab === 'transport'}
				<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6 max-w-2xl mx-auto">
					<h3 class="text-xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-3 flex items-center gap-2">
						<Bus class="w-5 h-5 text-primary" />
						University Transport Routes
					</h3>

					<div class="space-y-1.5">
						<label for="selectUniTrans" class="block text-xs font-bold uppercase text-on-surface-variant">Select University</label>
						<select
							id="selectUniTrans"
							bind:value={transportUniversityID}
							onchange={() => loadTransport(transportUniversityID)}
							class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 text-sm bg-white font-bold text-on-surface"
						>
							{#each universities as u}
								<option value={u.u_id}>{u.u_name}</option>
							{/each}
						</select>
					</div>

					<form onsubmit={handleSaveTransport} class="flex flex-col sm:flex-row gap-3 items-end bg-surface-container-low/50 p-4 rounded-2xl border border-outline-variant/30">
						<div class="space-y-1 flex-1 w-full">
							<label for="rName" class="block text-xs font-bold uppercase text-on-surface-variant">Transport Route</label>
							<input id="rName" type="text" bind:value={newRouteName} placeholder="e.g. Route A (Mirpur to Campus)" required class="w-full px-3.5 py-2.5 rounded-xl border text-xs bg-white font-bold" />
						</div>

						<div class="space-y-1 w-full sm:w-36">
							<label for="rTime" class="block text-xs font-bold uppercase text-on-surface-variant">Est. Travel Time</label>
							<input id="rTime" type="text" bind:value={newRouteTime} placeholder="45 mins" required class="w-full px-3.5 py-2.5 rounded-xl border text-xs bg-white font-bold" />
						</div>

						<div class="flex gap-2 shrink-0">
							{#if editingRouteName}
								<button type="button" onclick={() => { editingRouteName = null; newRouteName = ''; }} class="px-3 py-2.5 rounded-xl font-bold border text-xs">Cancel</button>
							{/if}
							<button type="submit" class="px-5 py-2.5 rounded-xl font-bold text-white bg-primary text-xs shadow-md flex items-center gap-1">
								<Plus class="w-3.5 h-3.5" /> {editingRouteName ? 'Save' : 'Add'}
							</button>
						</div>
					</form>

					<div class="space-y-3">
						<h4 class="text-xs font-extrabold uppercase tracking-wider text-on-surface-variant">Campus Routes List</h4>
						{#if transportRoutes.length === 0}
							<p class="text-xs text-on-surface-variant text-center py-4">No transport routes added for this university.</p>
						{:else}
							{#each transportRoutes as tr}
								<div class="flex items-center justify-between p-3.5 rounded-xl border border-outline-variant/30 bg-white">
									<div>
										<p class="text-xs font-bold text-on-surface">{tr.transport_route}</p>
										<p class="text-[10px] font-semibold text-primary">Time: {tr.est_travel_time}</p>
									</div>
									<div class="flex gap-1">
										<button onclick={() => startEditTransport(tr)} class="text-primary hover:bg-primary-fixed/40 p-1.5 rounded-lg transition-colors">
											<Edit3 class="w-4 h-4" />
										</button>
										<button onclick={() => handleDeleteTransportRoute(tr.transport_route)} class="text-error hover:bg-error-container/40 p-1.5 rounded-lg transition-colors">
											<Trash2 class="w-4 h-4" />
										</button>
									</div>
								</div>
							{/each}
						{/if}
					</div>
				</div>
			{/if}

			<!-- Tab 5: Admission Tests Management -->
			{#if activeTab === 'admissiontests'}
				<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6 max-w-2xl mx-auto">
					<h3 class="text-xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-3 flex items-center gap-2">
						<Clock class="w-5 h-5 text-primary" />
						{editingTestId ? `Edit Admission Test #${editingTestId}` : 'Create Admission Test'}
					</h3>

					<form onsubmit={handleCreateAdmissionTest} class="space-y-4">
						<div class="grid grid-cols-2 gap-4">
							<div>
								<label for="atUnit" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Exam Unit</label>
								<input id="atUnit" type="text" bind:value={atExamUnit} placeholder="A" class="w-full px-4 py-2.5 rounded-xl border border-outline-variant/50 text-sm bg-white" />
							</div>
							<div>
								<label for="atDate" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Exam Date (YYYY-MM-DD)</label>
								<input id="atDate" type="date" bind:value={atExamDate} required class="w-full px-4 py-2.5 rounded-xl border border-outline-variant/50 text-sm bg-white" />
							</div>
						</div>

						<div>
							<label for="atCenter" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Exam Center</label>
							<input id="atCenter" type="text" bind:value={atExamCenter} placeholder="Dhaka" class="w-full px-4 py-2.5 rounded-xl border border-outline-variant/50 text-sm bg-white" />
						</div>

						<div>
							<label for="atProgId" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Program</label>
							<select id="atProgId" bind:value={atProgramId} class="w-full px-4 py-2.5 rounded-xl border border-outline-variant/50 text-sm bg-white font-semibold">
								{#each programs as p}
									<option value={p.program_id}>{p.p_name} (ID #{p.program_id})</option>
								{/each}
							</select>
						</div>

						<div>
							<label for="atPrereq" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Prerequisite Test ID (optional)</label>
							<input id="atPrereq" type="number" bind:value={atPrereqTestId} placeholder="Leave blank if none" class="w-full px-4 py-2.5 rounded-xl border border-outline-variant/50 text-sm bg-white" />
						</div>

						<div class="flex gap-3 pt-2">
							{#if editingTestId}
								<button type="button" onclick={() => editingTestId = null} class="flex-1 py-2.5 px-4 rounded-xl font-semibold border border-outline-variant text-sm hover:bg-surface-container transition-all">Cancel Edit</button>
							{/if}
							<button type="submit" disabled={isLoading} class="flex-1 py-3 px-6 rounded-xl font-bold text-white bg-primary text-sm shadow-md hover:bg-primary-container disabled:opacity-50 transition-all flex items-center justify-center gap-2">
								<Clock class="w-4 h-4" />
								{editingTestId ? 'Update Admission Test' : 'Create Admission Test'}
							</button>
						</div>
					</form>
				</div>
			{/if}

			<!-- Tab 6: Publish Results (Batch Supported) -->
			{#if activeTab === 'publish'}
				<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6 max-w-2xl mx-auto">
					<h3 class="text-xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-3 flex items-center gap-2">
						<Send class="w-5 h-5 text-primary" />
						Publish Admission Test Results
					</h3>

					<form onsubmit={handlePublishResults} class="space-y-5">
						<div>
							<label for="pubTestId" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Admission Test ID</label>
							<input id="pubTestId" type="number" bind:value={testIDToPublish} required class="w-full px-4 py-2.5 rounded-xl border border-outline-variant/50 text-sm bg-white" />
						</div>

						<div class="space-y-3">
							<div class="flex items-center justify-between">
								<h4 class="text-xs font-extrabold uppercase tracking-wider text-on-surface">Batch Student Results ({publishResultsList.length})</h4>
								<button type="button" onclick={addPublishRow} class="text-xs font-bold text-primary flex items-center gap-1 hover:underline">
									<Plus class="w-3.5 h-3.5" /> Add Student Row
								</button>
							</div>

							{#each publishResultsList as row, idx}
								<div class="flex items-center gap-3 p-3.5 rounded-2xl border border-outline-variant/30 bg-surface-container-low/50">
									<div class="w-24 space-y-0.5">
										<label for={`sId_${idx}`} class="block text-[10px] font-bold uppercase text-on-surface-variant">Student ID</label>
										<input id={`sId_${idx}`} type="number" bind:value={row.student_id} required class="w-full px-3 py-1.5 rounded-lg border text-xs bg-white" />
									</div>

									<div class="flex-1 space-y-0.5">
										<label for={`sMarks_${idx}`} class="block text-[10px] font-bold uppercase text-on-surface-variant">Marks</label>
										<input id={`sMarks_${idx}`} type="text" bind:value={row.marks} required class="w-full px-3 py-1.5 rounded-lg border text-xs bg-white" />
									</div>

									<div class="w-28 space-y-0.5">
										<label for={`sMerit_${idx}`} class="block text-[10px] font-bold uppercase text-on-surface-variant">Merit Rank</label>
										<input id={`sMerit_${idx}`} type="number" bind:value={row.merit_position} required class="w-full px-3 py-1.5 rounded-lg border text-xs bg-white" />
									</div>

									<button type="button" onclick={() => removePublishRow(idx)} class="p-2 rounded-lg text-error hover:bg-error-container/40 shrink-0 mt-3">
										<Trash2 class="w-4 h-4" />
									</button>
								</div>
							{/each}
						</div>

						<button type="submit" disabled={isLoading} class="w-full py-3.5 px-6 rounded-xl font-bold text-white bg-primary text-sm shadow-md hover:bg-primary-container disabled:opacity-50 transition-all flex items-center justify-center gap-2">
							<Send class="w-4 h-4" /> Publish Results & Notify Students
						</button>
					</form>
				</div>
			{/if}

			<!-- Tab 7: Applications Review -->
			{#if activeTab === 'applications'}
				<div class="space-y-6">
					<div class="flex items-center justify-between">
						<h3 class="text-xl font-extrabold text-on-surface">Student Applications Review</h3>
						<div class="flex items-center gap-2">
							<label for="selectUni" class="text-xs font-bold text-on-surface-variant">Select University:</label>
							<select
								id="selectUni"
								bind:value={selectedUniversityID}
								onchange={() => loadApplications(selectedUniversityID)}
								class="px-3 py-2 rounded-xl border border-outline-variant/50 text-xs bg-white font-bold"
							>
								{#each universities as u}
									<option value={u.u_id}>{u.u_name}</option>
								{/each}
							</select>
						</div>
					</div>

					<div class="glass-panel p-6 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-4">
						{#if applications.length === 0}
							<p class="text-center py-10 text-on-surface-variant text-sm font-semibold">No applications submitted for this university yet.</p>
						{:else}
							<div class="divide-y divide-outline-variant/30 space-y-4">
								{#each applications as app}
									<div class="pt-4 first:pt-0 flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
										<div class="space-y-1">
											<div class="flex items-center gap-2">
												<h4 class="font-black text-on-surface text-base">App #{app.app_id}</h4>
												<span class="text-xs font-bold px-2.5 py-0.5 rounded-full {app.status === 'APPROVED' || app.status === 'Accepted' ? 'bg-emerald-100 text-emerald-700' : app.status === 'VERIFIED' || app.status === 'Verified' ? 'bg-blue-100 text-blue-700' : app.status === 'REJECTED' || app.status === 'Rejected' ? 'bg-red-100 text-red-700' : 'bg-amber-100 text-amber-700'}">
													{app.status}
												</span>
											</div>
											<p class="text-xs font-bold text-primary">
												{app.program_name ? app.program_name : 'Program #' + app.program_id}
											</p>
											<p class="text-xs text-on-surface-variant">
												Student: {app.first_name ? `${app.first_name} ${app.last_name || ''}` : `ID #${app.student_id}`} {app.email ? `(${app.email})` : ''}
											</p>
											{#if app.sub_date}
												<p class="text-[10px] text-outline font-medium">Submitted: {new Date(app.sub_date).toLocaleString()}</p>
											{/if}
										</div>

										<div class="flex flex-wrap gap-2 shrink-0">
											<button
												onclick={() => handleUpdateStatus(app.app_id, 'APPROVED')}
												class="px-3.5 py-2 rounded-xl text-xs font-bold bg-emerald-600 text-white hover:bg-emerald-700 transition-colors flex items-center gap-1 shadow-xs"
											>
												<CheckCircle2 class="w-3.5 h-3.5" /> Approve
											</button>
											<button
												onclick={() => handleUpdateStatus(app.app_id, 'VERIFIED')}
												class="px-3.5 py-2 rounded-xl text-xs font-bold bg-blue-600 text-white hover:bg-blue-700 transition-colors flex items-center gap-1 shadow-xs"
											>
												<ShieldCheck class="w-3.5 h-3.5" /> Verify
											</button>
											<button
												onclick={() => handleUpdateStatus(app.app_id, 'REJECTED')}
												class="px-3.5 py-2 rounded-xl text-xs font-bold bg-error text-white hover:bg-red-700 transition-colors flex items-center gap-1 shadow-xs"
											>
												<XCircle class="w-3.5 h-3.5" /> Reject
											</button>
										</div>
									</div>
								{/each}
							</div>
						{/if}
					</div>
				</div>
			{/if}
		{/if}
	</div>
</div>

<!-- Comprehensive University Modal (Includes Departments & Album Photos) -->
{#if showUniModal}
	<div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/40 backdrop-blur-sm overflow-y-auto">
		<div class="max-w-2xl w-full bg-white rounded-3xl p-6 sm:p-8 shadow-2xl border border-outline-variant/40 space-y-6 my-8">
			<div class="flex items-center justify-between border-b pb-3">
				<h3 class="text-xl font-extrabold text-on-surface">{editingUniId ? 'Edit University' : 'Create University'}</h3>
				<button onclick={() => showUniModal = false} class="p-2 rounded-xl hover:bg-surface-container transition-colors">
					<X class="w-5 h-5 text-outline" />
				</button>
			</div>

			<form onsubmit={handleSaveUniversity} class="space-y-5">
				<!-- Basic Details -->
				<div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
					<div>
						<label for="uName" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">University Name</label>
						<input id="uName" type="text" bind:value={uniName} required class="w-full px-3.5 py-2.5 rounded-xl border text-sm" />
					</div>
					<div>
						<label for="uWeb" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Website URL</label>
						<input id="uWeb" type="text" bind:value={uniWebsite} required class="w-full px-3.5 py-2.5 rounded-xl border text-sm" />
					</div>
					<div>
						<label for="uLoc" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Location</label>
						<input id="uLoc" type="text" bind:value={uniLocation} class="w-full px-3.5 py-2.5 rounded-xl border text-sm" />
					</div>
					<div>
						<label for="uLogo" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Logo URL</label>
						<input id="uLogo" type="text" bind:value={uniLogoUrl} class="w-full px-3.5 py-2.5 rounded-xl border text-sm" />
					</div>
				</div>

				<!-- Description & History -->
				<div class="space-y-4">
					<div>
						<label for="uDesc" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">University Description</label>
						<textarea id="uDesc" bind:value={uniDescription} rows="2" placeholder="Brief overview of the university..." class="w-full px-3.5 py-2.5 rounded-xl border text-sm"></textarea>
					</div>
					<div>
						<label for="uHist" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">University History</label>
						<textarea id="uHist" bind:value={uniHistory} rows="2" placeholder="Background and history..." class="w-full px-3.5 py-2.5 rounded-xl border text-sm"></textarea>
					</div>
				</div>

				<!-- Departments Builder -->
				<div class="space-y-3 p-4 rounded-2xl bg-surface-container-low/50 border border-outline-variant/30">
					<div class="flex items-center justify-between">
						<h4 class="text-xs font-extrabold uppercase tracking-wider text-on-surface flex items-center gap-2">
							<Users class="w-4 h-4 text-primary" />
							Departments ({uniDepartments.length})
						</h4>
						<button type="button" onclick={addDepartmentRow} class="text-xs font-bold text-primary flex items-center gap-1 hover:underline">
							<Plus class="w-3.5 h-3.5" /> Add Department
						</button>
					</div>

					{#each uniDepartments as d, idx}
						<div class="flex items-start gap-2 p-3 rounded-xl border border-outline-variant/30 bg-white">
							<div class="flex-1 grid grid-cols-1 sm:grid-cols-3 gap-2">
								<input type="text" bind:value={d.dept_name} placeholder="Dept Name (e.g. CSE)" class="px-3 py-1.5 rounded-lg border text-xs" />
								<input type="text" bind:value={d.dept_description} placeholder="Description" class="px-3 py-1.5 rounded-lg border text-xs" />
								<input type="number" bind:value={d.total_seats} placeholder="Seats" class="px-3 py-1.5 rounded-lg border text-xs" />
							</div>
							<button type="button" onclick={() => removeDepartmentRow(idx)} class="p-1.5 text-error hover:bg-error-container/40 rounded-lg">
								<Trash2 class="w-4 h-4" />
							</button>
						</div>
					{/each}
				</div>

				<!-- Campus Album Builder -->
				<div class="space-y-3 p-4 rounded-2xl bg-surface-container-low/50 border border-outline-variant/30">
					<div class="flex items-center justify-between">
						<h4 class="text-xs font-extrabold uppercase tracking-wider text-on-surface flex items-center gap-2">
							<ImageIcon class="w-4 h-4 text-primary" />
							Campus Album Photos ({uniAlbum.length})
						</h4>
						<button type="button" onclick={addAlbumRow} class="text-xs font-bold text-primary flex items-center gap-1 hover:underline">
							<Plus class="w-3.5 h-3.5" /> Add Photo
						</button>
					</div>

					{#each uniAlbum as a, idx}
						<div class="flex items-center gap-2 p-3 rounded-xl border border-outline-variant/30 bg-white">
							<div class="flex-1 grid grid-cols-1 sm:grid-cols-2 gap-2">
								<input type="text" bind:value={a.picture_title} placeholder="Picture Title" class="px-3 py-1.5 rounded-lg border text-xs" />
								<input type="text" bind:value={a.picture_url} placeholder="Photo URL (https://...)" class="px-3 py-1.5 rounded-lg border text-xs" />
							</div>
							<button type="button" onclick={() => removeAlbumRow(idx)} class="p-1.5 text-error hover:bg-error-container/40 rounded-lg">
								<Trash2 class="w-4 h-4" />
							</button>
						</div>
					{/each}
				</div>

				<div class="flex gap-3 pt-2">
					<button type="button" onclick={() => showUniModal = false} class="flex-1 py-3 px-4 rounded-xl font-semibold border text-sm">Cancel</button>
					<button type="submit" disabled={isLoading} class="flex-1 py-3 px-4 rounded-xl font-bold text-white bg-primary text-sm shadow-md">
						{isLoading ? 'Saving...' : 'Save University'}
					</button>
				</div>
			</form>
		</div>
	</div>
{/if}

<!-- Program Modal -->
{#if showProgModal}
	<div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/40 backdrop-blur-sm">
		<div class="max-w-md w-full bg-white rounded-3xl p-6 sm:p-8 shadow-2xl border border-outline-variant/40 space-y-4">
			<h3 class="text-xl font-bold text-on-surface">{editingProgId ? 'Edit Program' : 'Create Program'}</h3>
			<form onsubmit={handleSaveProgram} class="space-y-3">
				<div>
					<label for="pName" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Program Name</label>
					<input id="pName" type="text" bind:value={progName} required class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm" />
				</div>
				<div class="grid grid-cols-2 gap-3">
					<div>
						<label for="pUnit" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Exam Unit</label>
						<input id="pUnit" type="text" bind:value={progUnit} required class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm" />
					</div>
					<div>
						<label for="pSeats" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Total Seats</label>
						<input id="pSeats" type="number" bind:value={progSeats} required class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm" />
					</div>
				</div>
				<div class="grid grid-cols-2 gap-3">
					<div>
						<label for="pCut" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Cutmarks</label>
						<input id="pCut" type="number" step="0.01" bind:value={progCutmarks} required class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm" />
					</div>
					<div>
						<label for="pDead" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Deadline</label>
						<input id="pDead" type="date" bind:value={progDeadline} required class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm" />
					</div>
				</div>
				<div class="grid grid-cols-2 gap-3">
					<div>
						<label for="pFee" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Application Fee (BDT)</label>
						<input id="pFee" type="number" step="10" bind:value={progFee} required class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm" />
					</div>
					<div>
						<label for="pUni" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">University</label>
						<select id="pUni" bind:value={progUniId} class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm">
							{#each universities as u}
								<option value={u.u_id}>{u.u_name}</option>
							{/each}
						</select>
					</div>
				</div>

				<!-- Required Student Attributes Checklist -->
				<div class="space-y-2 pt-1">
					<label class="block text-xs font-extrabold uppercase tracking-wider text-on-surface-variant">
						Required Student Attributes & Docs
					</label>
					<div class="grid grid-cols-2 gap-1.5 p-3 rounded-2xl bg-surface-container-low/60 border border-outline-variant/30 text-xs">
						{#each availableRequiredFields as field}
							<label class="flex items-center gap-2 font-semibold text-on-surface select-none cursor-pointer hover:text-primary transition-colors">
								<input
									type="checkbox"
									checked={progRequiredFields.includes(field)}
									onchange={() => toggleRequiredField(field)}
									class="rounded text-primary focus:ring-primary/40 w-3.5 h-3.5"
								/>
								<span class="truncate">{field}</span>
							</label>
						{/each}
					</div>
				</div>
				<div class="flex gap-3 pt-2">
					<button type="button" onclick={() => showProgModal = false} class="flex-1 py-2.5 px-4 rounded-xl font-semibold border text-sm">Cancel</button>
					<button type="submit" disabled={isLoading} class="flex-1 py-2.5 px-4 rounded-xl font-bold text-white bg-primary text-sm shadow-md">
						{isLoading ? 'Saving...' : 'Save Program'}
					</button>
				</div>
			</form>
		</div>
	</div>
{/if}
