<!-- src/routes/admin/+page.svelte -->
<script lang="ts">
	import { onMount } from 'svelte';
	import {
		adminLogin,
		fetchUniversityApplications,
		updateApplicationStatus,
		createUniversity,
		updateUniversity,
		deleteUniversity,
		createProgram,
		updateProgram,
		deleteProgram
	} from '$lib/api/admin';
	import { fetchUniversities } from '$lib/api/university';
	import { fetchPrograms } from '$lib/api/programs';
	import type { University, Program, StudentApplication } from '$lib/types/models';
	import { Shield, Building2, BookOpen, FileCheck, LogOut, Plus, Trash2, Edit3, CheckCircle2, XCircle, AlertCircle } from 'lucide-svelte';

	let isAdminLoggedIn = $state(false);
	let adminEmail = $state('admin@system.com');
	let adminPassword = $state('admin123secret');

	let activeTab = $state<'universities' | 'programs' | 'applications'>('universities');
	let isLoading = $state(false);
	let statusMessage = $state<string | null>(null);
	let errorMessage = $state<string | null>(null);

	// Data
	let universities = $state<University[]>([]);
	let programs = $state<Program[]>([]);
	let applications = $state<StudentApplication[]>([]);
	let selectedUniversityID = $state<number>(1);

	// University Modal / Form State
	let showUniModal = $state(false);
	let editingUniId = $state<number | null>(null);
	let uniName = $state('');
	let uniWebsite = $state('');
	let uniLocation = $state('');
	let uniLogoUrl = $state('');

	// Program Modal / Form State
	let showProgModal = $state(false);
	let editingProgId = $state<number | null>(null);
	let progName = $state('');
	let progUnit = $state('A');
	let progSeats = $state<number>(100);
	let progCutmarks = $state<number>(80.0);
	let progDeadline = $state('2026-12-31');
	let progUniId = $state<number>(1);

	async function handleAdminLogin(e: Event) {
		e.preventDefault();
		isLoading = true;
		errorMessage = null;
		try {
			const res = await adminLogin({ email: adminEmail, password: adminPassword });
			if (res.token) {
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
		isAdminLoggedIn = false;
		universities = [];
		programs = [];
		applications = [];
	}

	async function loadAllData() {
		isLoading = true;
		errorMessage = null;
		try {
			const [uList, pList] = await Promise.all([fetchUniversities(), fetchPrograms()]);
			universities = uList || [];
			programs = pList || [];
			if (universities.length > 0) {
				selectedUniversityID = universities[0].u_id;
				await loadApplications(selectedUniversityID);
			}
		} catch (err: any) {
			console.error(err);
		} finally {
			isLoading = false;
		}
	}

	async function loadApplications(uId: number) {
		try {
			applications = await fetchUniversityApplications(uId);
		} catch (err: any) {
			applications = [];
		}
	}

	// University Actions
	function openCreateUniModal() {
		editingUniId = null;
		uniName = '';
		uniWebsite = '';
		uniLocation = '';
		uniLogoUrl = '';
		showUniModal = true;
	}

	function openEditUniModal(u: University) {
		editingUniId = u.u_id;
		uniName = u.u_name;
		uniWebsite = u.website;
		uniLocation = u.location || '';
		uniLogoUrl = u.logo_url || '';
		showUniModal = true;
	}

	async function handleSaveUniversity(e: Event) {
		e.preventDefault();
		isLoading = true;
		statusMessage = null;
		errorMessage = null;
		try {
			const payload = { name: uniName, website: uniWebsite, location: uniLocation, logo_url: uniLogoUrl };
			if (editingUniId) {
				await updateUniversity(editingUniId, payload);
				statusMessage = 'University updated successfully!';
			} else {
				await createUniversity(payload);
				statusMessage = 'University created successfully!';
			}
			showUniModal = false;
			await loadAllData();
		} catch (err: any) {
			errorMessage = err?.message || 'Failed to save university.';
		} finally {
			isLoading = false;
		}
	}

	async function handleDeleteUniversity(uId: number) {
		if (!confirm('Are you sure you want to delete this university?')) return;
		isLoading = true;
		try {
			await deleteUniversity(uId);
			statusMessage = 'University deleted successfully!';
			await loadAllData();
		} catch (err: any) {
			errorMessage = err?.message || 'Failed to delete university.';
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
		progUniId = universities.length > 0 ? universities[0].u_id : 1;
		showProgModal = true;
	}

	function openEditProgModal(p: Program) {
		editingProgId = p.program_id;
		progName = p.p_name;
		progUnit = p.p_unit || 'A';
		progSeats = p.total_seats;
		progCutmarks = typeof p.prev_cutmarks === 'number' ? p.prev_cutmarks : parseFloat(String(p.prev_cutmarks) || '80');
		progDeadline = p.deadline ? p.deadline.split('T')[0] : '2026-12-31';
		progUniId = p.u_id;
		showProgModal = true;
	}

	async function handleSaveProgram(e: Event) {
		e.preventDefault();
		isLoading = true;
		statusMessage = null;
		errorMessage = null;
		try {
			const payload = {
				p_name: progName,
				p_unit: progUnit,
				total_seats: Number(progSeats),
				prev_cutmarks: Number(progCutmarks),
				deadline: progDeadline,
				u_id: Number(progUniId)
			};
			if (editingProgId) {
				await updateProgram(editingProgId, payload);
				statusMessage = 'Program updated successfully!';
			} else {
				await createProgram(payload);
				statusMessage = 'Program created successfully!';
			}
			showProgModal = false;
			await loadAllData();
		} catch (err: any) {
			errorMessage = err?.message || 'Failed to save program.';
		} finally {
			isLoading = false;
		}
	}

	async function handleDeleteProgram(pId: number) {
		if (!confirm('Are you sure you want to delete this program?')) return;
		isLoading = true;
		try {
			await deleteProgram(pId);
			statusMessage = 'Program deleted successfully!';
			await loadAllData();
		} catch (err: any) {
			errorMessage = err?.message || 'Failed to delete program.';
		} finally {
			isLoading = false;
		}
	}

	// Application Status Update
	async function handleUpdateStatus(appId: number, status: string) {
		isLoading = true;
		statusMessage = null;
		errorMessage = null;
		try {
			await updateApplicationStatus(appId, status);
			statusMessage = `Application #${appId} updated to ${status}`;
			await loadApplications(selectedUniversityID);
		} catch (err: any) {
			errorMessage = err?.message || 'Failed to update application status.';
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
							required
							class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
						/>
					</div>

					<div class="space-y-1">
						<label for="adminPass" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Admin Password</label>
						<input
							id="adminPass"
							type="password"
							bind:value={adminPassword}
							required
							class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
						/>
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
						Admin Mode Active
					</div>
					<h1 class="text-3xl sm:text-4xl font-black text-white">University & Program Management</h1>
					<p class="text-slate-100 text-sm">Manage university records, admission programs, cutmarks, and student application approvals.</p>
				</div>

				<button
					onclick={handleAdminLogout}
					class="relative z-10 px-5 py-2.5 rounded-xl text-xs font-bold bg-white/15 hover:bg-white/25 text-white border border-white/20 backdrop-blur-md transition-all flex items-center gap-2"
				>
					<LogOut class="w-4 h-4" />
					Logout Admin
				</button>
			</div>

			{#if statusMessage}
				<div class="p-4 rounded-2xl bg-tertiary-fixed/30 border border-tertiary/40 flex items-center gap-3 text-on-tertiary-fixed-variant text-sm">
					<CheckCircle2 class="w-5 h-5 text-tertiary shrink-0" />
					<span class="font-bold">{statusMessage}</span>
				</div>
			{/if}

			{#if errorMessage}
				<div class="p-4 rounded-2xl bg-error-container/60 border border-error/30 flex items-center gap-3 text-on-error-container text-sm">
					<AlertCircle class="w-5 h-5 text-error shrink-0" />
					<span>{errorMessage}</span>
				</div>
			{/if}

			<!-- Navigation Tabs -->
			<div class="flex items-center gap-2 bg-white/80 p-2 rounded-2xl border border-outline-variant/30 shadow-sm">
				<button
					onclick={() => activeTab = 'universities'}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all flex items-center justify-center gap-2 {activeTab === 'universities' ? 'bg-primary text-white shadow-md' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<Building2 class="w-4 h-4" />
					Universities ({universities.length})
				</button>
				<button
					onclick={() => activeTab = 'programs'}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all flex items-center justify-center gap-2 {activeTab === 'programs' ? 'bg-primary text-white shadow-md' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<BookOpen class="w-4 h-4" />
					Programs ({programs.length})
				</button>
				<button
					onclick={() => activeTab = 'applications'}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all flex items-center justify-center gap-2 {activeTab === 'applications' ? 'bg-primary text-white shadow-md' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<FileCheck class="w-4 h-4" />
					Applications Review
				</button>
			</div>

			<!-- Tab 1: Universities Management -->
			{#if activeTab === 'universities'}
				<div class="space-y-6">
					<div class="flex items-center justify-between">
						<h3 class="text-xl font-extrabold text-on-surface">Registered Universities</h3>
						<button
							onclick={openCreateUniModal}
							class="px-5 py-2.5 rounded-xl text-xs font-bold text-white bg-primary hover:bg-primary-container shadow-md flex items-center gap-2"
						>
							<Plus class="w-4 h-4" />
							Add University
						</button>
					</div>

					<div class="grid grid-cols-1 md:grid-cols-3 gap-6">
						{#each universities as u}
							<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/40 bg-white/95 shadow-md flex flex-col justify-between space-y-4">
								<div class="space-y-2">
									<div class="flex items-center justify-between">
										<span class="text-xs font-mono font-bold text-outline">ID #{u.u_id}</span>
										<div class="flex gap-2">
											<button onclick={() => openEditUniModal(u)} class="p-1.5 rounded-lg text-primary hover:bg-primary-fixed/40 transition-colors">
												<Edit3 class="w-4 h-4" />
											</button>
											<button onclick={() => handleDeleteUniversity(u.u_id)} class="p-1.5 rounded-lg text-error hover:bg-error-container/40 transition-colors">
												<Trash2 class="w-4 h-4" />
											</button>
										</div>
									</div>
									<h4 class="text-xl font-bold text-on-surface">{u.u_name}</h4>
									<p class="text-xs text-on-surface-variant">{u.location || 'Location Not Specified'}</p>
									<a href={u.website} target="_blank" class="text-xs font-semibold text-primary hover:underline block truncate">{u.website}</a>
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
						<h3 class="text-xl font-extrabold text-on-surface">Admission Programs</h3>
						<button
							onclick={openCreateProgModal}
							class="px-5 py-2.5 rounded-xl text-xs font-bold text-white bg-primary hover:bg-primary-container shadow-md flex items-center gap-2"
						>
							<Plus class="w-4 h-4" />
							Add Program
						</button>
					</div>

					<div class="grid grid-cols-1 md:grid-cols-3 gap-6">
						{#each programs as p}
							<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/40 bg-white/95 shadow-md flex flex-col justify-between space-y-4">
								<div class="space-y-2">
									<div class="flex items-center justify-between">
										<span class="text-xs font-mono font-bold text-outline">ID #{p.program_id}</span>
										<div class="flex gap-2">
											<button onclick={() => openEditProgModal(p)} class="p-1.5 rounded-lg text-primary hover:bg-primary-fixed/40 transition-colors">
												<Edit3 class="w-4 h-4" />
											</button>
											<button onclick={() => handleDeleteProgram(p.program_id)} class="p-1.5 rounded-lg text-error hover:bg-error-container/40 transition-colors">
												<Trash2 class="w-4 h-4" />
											</button>
										</div>
									</div>
									<h4 class="text-lg font-bold text-on-surface">{p.p_name}</h4>
									<div class="flex items-center gap-2 text-xs font-semibold text-on-surface-variant">
										<span class="bg-primary-fixed/50 text-on-primary-fixed px-2 py-0.5 rounded">Unit {p.p_unit || 'A'}</span>
										<span>{p.total_seats} Seats</span>
										<span>Cutmark: {p.prev_cutmarks}</span>
									</div>
								</div>
							</div>
						{/each}
					</div>
				</div>
			{/if}

			<!-- Tab 3: Applications Review -->
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
										<div>
											<div class="flex items-center gap-2">
												<h4 class="font-bold text-on-surface text-base">App #{app.app_id}</h4>
												<span class="text-xs font-bold px-2.5 py-0.5 rounded-full {app.status === 'APPROVED' ? 'bg-emerald-100 text-emerald-700' : app.status === 'REJECTED' ? 'bg-red-100 text-red-700' : 'bg-amber-100 text-amber-700'}">
													{app.status}
												</span>
											</div>
											<p class="text-xs text-on-surface-variant mt-1">Student ID: {app.student_id} | Email: {app.email || 'N/A'}</p>
										</div>
										<div class="flex gap-2">
											<button
												onclick={() => handleUpdateStatus(app.app_id, 'APPROVED')}
												class="px-4 py-2 rounded-xl text-xs font-bold bg-emerald-600 text-white hover:bg-emerald-700 transition-colors flex items-center gap-1"
											>
												<CheckCircle2 class="w-3.5 h-3.5" /> Approve
											</button>
											<button
												onclick={() => handleUpdateStatus(app.app_id, 'REJECTED')}
												class="px-4 py-2 rounded-xl text-xs font-bold bg-error text-white hover:bg-red-700 transition-colors flex items-center gap-1"
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

<!-- University Modal -->
{#if showUniModal}
	<div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/40 backdrop-blur-sm">
		<div class="max-w-md w-full bg-white rounded-3xl p-6 sm:p-8 shadow-2xl border border-outline-variant/40 space-y-4">
			<h3 class="text-xl font-bold text-on-surface">{editingUniId ? 'Edit University' : 'Create University'}</h3>
			<form onsubmit={handleSaveUniversity} class="space-y-3">
				<div>
					<label for="uName" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">University Name</label>
					<input id="uName" type="text" bind:value={uniName} required class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm" />
				</div>
				<div>
					<label for="uWeb" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Website URL</label>
					<input id="uWeb" type="text" bind:value={uniWebsite} required class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm" />
				</div>
				<div>
					<label for="uLoc" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Location</label>
					<input id="uLoc" type="text" bind:value={uniLocation} class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm" />
				</div>
				<div>
					<label for="uLogo" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">Logo URL</label>
					<input id="uLogo" type="text" bind:value={uniLogoUrl} class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm" />
				</div>
				<div class="flex gap-3 pt-2">
					<button type="button" onclick={() => showUniModal = false} class="flex-1 py-2.5 px-4 rounded-xl font-semibold border text-sm">Cancel</button>
					<button type="submit" class="flex-1 py-2.5 px-4 rounded-xl font-bold text-white bg-primary text-sm">Save</button>
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
				<div>
					<label for="pUni" class="block text-xs font-bold uppercase text-on-surface-variant mb-1">University</label>
					<select id="pUni" bind:value={progUniId} class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm">
						{#each universities as u}
							<option value={u.u_id}>{u.u_name}</option>
						{/each}
					</select>
				</div>
				<div class="flex gap-3 pt-2">
					<button type="button" onclick={() => showProgModal = false} class="flex-1 py-2.5 px-4 rounded-xl font-semibold border text-sm">Cancel</button>
					<button type="submit" class="flex-1 py-2.5 px-4 rounded-xl font-bold text-white bg-primary text-sm">Save</button>
				</div>
			</form>
		</div>
	</div>
{/if}
