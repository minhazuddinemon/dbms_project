<!-- src/routes/apply/[programId]/+page.svelte -->
<script lang="ts">
	import { applyToProgram, fetchProgramRequirements, fetchStudentApplications } from '$lib/api/applications';
	import { fetchProgramByID } from '$lib/api/programs';
	import { fetchUniversities } from '$lib/api/university';
	import { authState } from '$lib/state/auth.svelte';
	import { toastState } from '$lib/state/toast.svelte';
	import type { Program, ProgramRequirementsResponse, StudentApplication } from '$lib/types/models';
	import { page } from '$app/state';
	import {
		Building2,
		CheckCircle2,
		AlertCircle,
		Calendar,
		Users,
		Award,
		ShieldAlert,
		ArrowLeft,
		ArrowRight,
		CreditCard,
		Sparkles,
		ClipboardCheck,
		XCircle,
		Info,
		MapPin
	} from 'lucide-svelte';
	import { onMount } from 'svelte';

	let programId = $derived(Number(page.params.programId));
	let program = $state<Program | null>(null);
	let requirements = $state<ProgramRequirementsResponse | null>(null);
	let isLoading = $state(true);
	let isSubmitting = $state(false);

	let successMessage = $state<string | null>(null);
	let errorMessage = $state<string | null>(null);
	let missingFields = $state<string[]>([]);

	let existingApplication = $state<StudentApplication | null>(null);
	let universityLogo = $state<string | null>(null);

	async function loadData() {
		isLoading = true;
		try {
			const [prog, req, userApps, uList] = await Promise.allSettled([
				fetchProgramByID(programId),
				authState.isAuthenticated ? fetchProgramRequirements(programId) : Promise.resolve(null),
				authState.isAuthenticated ? fetchStudentApplications() : Promise.resolve([]),
				fetchUniversities().catch(() => [])
			]);
			if (prog.status === 'fulfilled' && prog.value) {
				program = prog.value;
				if (program.logo_url) {
					universityLogo = program.logo_url;
				} else if (program.university_logo) {
					universityLogo = program.university_logo;
				} else if (uList.status === 'fulfilled' && Array.isArray(uList.value)) {
					const targetName = (program?.university_name || program?.u_name || '').toLowerCase();
					const found = uList.value.find((u) => u && (u.u_id === program?.u_id || (u.u_name && u.u_name.toLowerCase() === targetName)));
					if (found?.logo_url) universityLogo = found.logo_url;
				}
			}
			if (req.status === 'fulfilled') requirements = req.value;
			if (userApps.status === 'fulfilled' && Array.isArray(userApps.value)) {
				const match = userApps.value.find((app) => app.program_id === programId);
				if (match) existingApplication = match;
			}
		} catch (err) {
			console.error('Failed to load data:', err);
		} finally {
			isLoading = false;
		}
	}

	onMount(() => {
		loadData();
	});

	async function handleApply() {
		if (!authState.isAuthenticated) {
			toastState.warning('Please sign in or register to continue applying.');
			errorMessage = 'Please sign in or register to submit your program application.';
			return;
		}

		if (existingApplication) {
			errorMessage = 'You have already submitted an application for this program.';
			return;
		}

		isSubmitting = true;
		successMessage = null;
		errorMessage = null;
		missingFields = [];

		try {
			const res = await applyToProgram(programId);
			if (res.status === 'success') {
				successMessage = res.message || 'Application submitted successfully!';
			} else if (res.status === 'incomplete_profile') {
				errorMessage = res.message;
				missingFields = res.missing_fields || [];
			}
		} catch (err: any) {
			errorMessage = err?.message || 'Failed to submit application.';
		} finally {
			isSubmitting = false;
		}
	}

	let readinessPercent = $derived.by(() => {
		if (!requirements || !requirements.required_fields) return 100;
		const total = requirements.required_fields.length;
		if (total === 0) return 100;
		const provided = requirements.required_fields.filter((f) => f.is_provided).length;
		return Math.round((provided / total) * 100);
	});
</script>

<svelte:head>
	<title>{program ? `Apply: ${program.p_name}` : 'Apply for Program'} - UniApp</title>
	<meta name="description" content="Submit your application for a public university program. Check eligibility requirements and proceed to payment." />
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">

		<!-- Back Link -->
		<a href="/programs" class="inline-flex items-center gap-2 text-sm font-bold text-outline hover:text-primary transition-colors group">
			<ArrowLeft class="w-4 h-4 group-hover:-translate-x-1 transition-transform" />
			Back to Programs
		</a>

		{#if isLoading}
			<div class="py-20 text-center text-on-surface-variant">
				<div class="w-12 h-12 border-4 border-primary border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
				<p class="font-bold text-lg">Loading program details...</p>
			</div>
		{:else if !program}
			<div class="glass-panel rounded-[2.5rem] border border-outline-variant/30 p-12 text-center max-w-lg mx-auto space-y-4 bg-white/90 shadow-xl">
				<AlertCircle class="w-14 h-14 text-error mx-auto" />
				<h3 class="text-2xl font-bold text-on-surface">Program Not Found</h3>
				<p class="text-on-surface-variant text-sm">The program ID requested does not exist or has been removed.</p>
				<a href="/programs" class="inline-block px-6 py-3 rounded-xl font-bold bg-primary text-white text-sm hover:bg-primary-container shadow-lg shadow-primary/25 transition-all">
					Browse All Programs
				</a>
			</div>
		{:else}
			<!-- Program Header Banner -->
			<div class="bg-gradient-to-r from-primary via-primary-container to-secondary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
				<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

				<div class="relative z-10 flex flex-col md:flex-row items-start md:items-center justify-between gap-6">
					<div class="flex items-center gap-5">
						{#if universityLogo || program.logo_url || program.university_logo}
							<img
								src={universityLogo || program.logo_url || program.university_logo}
								alt={program.university_name || program.u_name}
								class="w-16 h-16 rounded-2xl object-contain p-1.5 border border-white/30 bg-white shadow-lg shrink-0"
							/>
						{:else}
							<div class="w-16 h-16 rounded-2xl bg-white/10 text-white flex items-center justify-center font-black text-2xl border border-white/30 backdrop-blur-md shrink-0">
								{(program.university_name || program.u_name || 'U').charAt(0)}
							</div>
						{/if}

						<div class="space-y-2">
							<div class="flex flex-wrap items-center gap-3">
								<span class="px-3.5 py-1 rounded-full bg-white/15 border border-white/20 text-white text-xs font-bold uppercase tracking-wider backdrop-blur-md">
									Unit {program.p_unit || 'A'}
								</span>
								<span class="text-xs font-mono text-white/70">Program ID: #{program.program_id}</span>
							</div>

							<h1 class="text-2xl sm:text-3xl font-black text-white leading-tight">{program.p_name}</h1>

							<div class="flex flex-wrap gap-5 text-sm text-white/80 pt-0.5">
								<span class="flex items-center gap-2 font-extrabold text-white">
									<Building2 class="w-4 h-4 text-white/80" />
									{program.university_name || program.u_name || 'Public University'}
								</span>
								{#if program.university_location || program.location}
									<span class="flex items-center gap-2 font-semibold">
										<MapPin class="w-4 h-4 text-tertiary-fixed" />
										{program.university_location || program.location}
									</span>
								{/if}
								<span class="flex items-center gap-2">
									<Users class="w-4 h-4 text-white/60" />
									{program.total_seats} Total Seats
								</span>
								<span class="flex items-center gap-2">
									<Calendar class="w-4 h-4 text-white/60" />
									Deadline: {program.deadline ? program.deadline.split('T')[0] : 'TBA'}
								</span>
								{#if program.prev_cutmarks}
									<span class="flex items-center gap-2">
										<Award class="w-4 h-4 text-white/60" />
										Prev. Cutmark: {program.prev_cutmarks}
									</span>
								{/if}
							</div>
						</div>
					</div>
				</div>
			</div>

			<div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
				<!-- Main Application Panel -->
				<div class="lg:col-span-2 space-y-6">

					<!-- Auth Gate -->
					{#if !authState.isAuthenticated}
						<div class="glass-panel p-8 sm:p-10 rounded-[2.5rem] border border-amber-200/80 bg-amber-50/60 shadow-xl text-center space-y-6">
							<div class="w-16 h-16 rounded-2xl bg-amber-500 text-white flex items-center justify-center mx-auto shadow-md">
								<ShieldAlert class="w-8 h-8" />
							</div>
							<div class="space-y-2">
								<h3 class="text-2xl font-extrabold text-amber-950">Please Sign In or Register to Continue</h3>
								<p class="text-amber-900/80 text-sm leading-relaxed max-w-md mx-auto font-medium">
									You must be logged in as a student to submit your application for <strong class="text-amber-950 font-bold">{program?.p_name}</strong>.
								</p>
							</div>
							<div class="flex flex-wrap gap-4 justify-center pt-2">
								<a href="/login" class="px-8 py-3.5 rounded-xl font-black bg-primary text-white text-sm hover:bg-primary-container shadow-lg shadow-primary/25 transition-all">
									Sign In to Apply
								</a>
								<a href="/register" class="px-8 py-3.5 rounded-xl font-black bg-white border border-outline-variant/60 text-on-surface text-sm hover:bg-surface-container shadow-sm transition-all">
									Create Free Account
								</a>
							</div>
						</div>
					{:else if existingApplication}
						<!-- Already Applied State -->
						<div class="glass-panel p-8 rounded-[2.5rem] border border-primary/30 bg-primary-fixed/20 shadow-xl text-center space-y-5">
							<div class="w-16 h-16 rounded-2xl bg-primary text-white flex items-center justify-center mx-auto shadow-md">
								<CheckCircle2 class="w-8 h-8" />
							</div>
							<div class="space-y-1">
								<h3 class="text-2xl font-extrabold text-on-surface">Application Already Submitted</h3>
								<p class="text-xs text-on-surface-variant font-mono">
									Application ID: <strong>#{existingApplication.app_id}</strong> | Status: <strong class="text-primary uppercase">{existingApplication.status}</strong>
								</p>
							</div>
							<p class="text-on-surface-variant text-sm leading-relaxed max-w-md mx-auto">
								You have already submitted an application for <strong>{program?.p_name}</strong>. Re-applying to the same program is disabled.
							</p>
							<div class="flex flex-wrap gap-3 justify-center pt-2">
								{#if existingApplication.status === 'PAID' || existingApplication.status === 'APPROVED'}
									<div class="inline-flex items-center gap-2 px-6 py-3 rounded-xl font-bold bg-emerald-600 text-white text-sm shadow-md">
										<CheckCircle2 class="w-4 h-4 text-white" />
										Payment Completed
									</div>
								{:else}
									<a href="/payment" class="inline-flex items-center gap-2 px-6 py-3 rounded-xl font-bold bg-primary text-white text-sm hover:bg-primary-container shadow-lg shadow-primary/25 transition-all">
										<CreditCard class="w-4 h-4" />
										Proceed to Payment
									</a>
								{/if}
								<a href="/dashboard" class="px-6 py-3 rounded-xl font-bold border border-outline-variant text-on-surface text-sm hover:bg-surface-container transition-all">
									View Dashboard
								</a>
							</div>
						</div>
					{:else}
						<!-- Success State -->
						{#if successMessage}
							<div class="glass-panel p-8 rounded-[2.5rem] border border-tertiary/30 bg-tertiary-fixed/20 shadow-xl text-center space-y-5">
								<CheckCircle2 class="w-16 h-16 text-tertiary mx-auto" />
								<h3 class="text-2xl font-extrabold text-on-surface">Application Submitted!</h3>
								<p class="text-on-surface-variant text-sm leading-relaxed">{successMessage}</p>
								<div class="flex gap-3 justify-center pt-2">
									<a href="/payment" class="inline-flex items-center gap-2 px-6 py-3 rounded-xl font-bold bg-primary text-white text-sm hover:bg-primary-container shadow-lg shadow-primary/25 transition-all">
										<CreditCard class="w-4 h-4" />
										Proceed to Payment
									</a>
									<a href="/dashboard" class="px-6 py-3 rounded-xl font-bold border border-outline-variant text-on-surface text-sm hover:bg-surface-container transition-all">
										View Dashboard
									</a>
								</div>
							</div>
						{:else}
							<!-- Application Form -->
							<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
								<div class="flex items-center gap-3 border-b border-outline-variant/30 pb-5">
									<div class="w-10 h-10 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center">
										<ClipboardCheck class="w-5 h-5" />
									</div>
									<div>
										<h2 class="text-xl font-extrabold text-on-surface">Submit Application</h2>
										<p class="text-xs text-on-surface-variant">Your application will be reviewed after payment.</p>
									</div>
								</div>

								<!-- Error / Missing Fields Alert -->
								{#if errorMessage}
									<div class="p-5 rounded-2xl bg-amber-50 border border-amber-200 space-y-4">
										<div class="flex items-start gap-3">
											<ShieldAlert class="w-6 h-6 text-amber-600 shrink-0 mt-0.5" />
											<div>
												<h3 class="text-base font-bold text-amber-900">Incomplete Profile</h3>
												<p class="text-amber-700 text-sm mt-1">{errorMessage}</p>
											</div>
										</div>
										{#if missingFields.length > 0}
											<div class="bg-white/80 rounded-xl p-4 border border-amber-200/80">
												<span class="block text-xs font-bold text-amber-900 uppercase tracking-wider mb-2">Missing Required Fields:</span>
												<div class="flex flex-wrap gap-2">
													{#each missingFields as field}
														<span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-lg bg-amber-100 text-amber-800 text-xs font-bold border border-amber-200">
															<XCircle class="w-3 h-3" />
															{field}
														</span>
													{/each}
												</div>
												<div class="mt-4 pt-3 border-t border-amber-200/60">
													<a
														href="/profile"
														class="inline-flex items-center gap-2 px-5 py-2.5 rounded-xl bg-amber-600 hover:bg-amber-700 text-white text-xs font-bold transition-colors shadow-md"
													>
														Fill Missing Profile Fields
														<ArrowRight class="w-3.5 h-3.5" />
													</a>
												</div>
											</div>
										{/if}
									</div>
								{/if}

								<!-- Fee & Submit -->
								<div class="p-5 rounded-2xl bg-surface-container-low/60 border border-outline-variant/30 space-y-4">
									<div class="flex items-center justify-between">
										<div>
											<span class="block text-xs font-bold text-on-surface-variant uppercase tracking-wider">Application Fee</span>
											<span class="text-2xl font-black text-on-surface">BDT 500.00</span>
										</div>
										<div class="w-12 h-12 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center">
											<CreditCard class="w-6 h-6" />
										</div>
									</div>
									<p class="text-xs text-on-surface-variant leading-relaxed">
										Fee is charged once per program application. Payment can be completed via bKash, Nagad, or Card after submission.
									</p>
								</div>

								<button
									onclick={handleApply}
									disabled={isSubmitting}
									class="w-full py-4 px-6 rounded-2xl font-bold text-white bg-primary hover:bg-primary-container disabled:opacity-50 shadow-lg shadow-primary/25 hover:shadow-primary/40 transition-all flex items-center justify-center gap-2 text-sm"
								>
									{#if isSubmitting}
										<div class="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
										<span>Submitting Application...</span>
									{:else}
										<ClipboardCheck class="w-5 h-5" />
										<span>Submit Application</span>
									{/if}
								</button>
							</div>
						{/if}
					{/if}
				</div>

				<!-- Sidebar: Requirements Checklist -->
				<div class="space-y-6">
					{#if requirements && authState.isAuthenticated}
						<div class="glass-panel p-6 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-5 sticky top-28">
							<h3 class="text-base font-extrabold text-on-surface flex items-center gap-2 border-b border-outline-variant/30 pb-4">
								<Info class="w-5 h-5 text-primary" />
								Profile Readiness
							</h3>

							<!-- Progress Bar -->
							<div class="space-y-2">
								<div class="flex items-center justify-between text-xs font-bold">
									<span class="text-on-surface-variant">Completion</span>
									<span class="{readinessPercent === 100 ? 'text-tertiary' : 'text-amber-600'}">{readinessPercent}%</span>
								</div>
								<div class="w-full bg-surface-container-low rounded-full h-2.5 overflow-hidden">
									<div
										class="h-full rounded-full transition-all duration-500 {readinessPercent === 100 ? 'bg-tertiary' : 'bg-amber-500'}"
										style="width: {readinessPercent}%"
									></div>
								</div>
							</div>

							<!-- Fields List -->
							{#if requirements.required_fields && requirements.required_fields.length > 0}
								<div class="space-y-2.5">
									{#each requirements.required_fields as field}
										<div class="flex items-center justify-between p-3 rounded-xl border {field.is_provided ? 'border-emerald-200 bg-emerald-50/50' : 'border-amber-200 bg-amber-50/50'}">
											<span class="text-xs font-semibold text-on-surface">{field.field_name}</span>
											{#if field.is_provided}
												<CheckCircle2 class="w-4 h-4 text-emerald-600 shrink-0" />
											{:else}
												<XCircle class="w-4 h-4 text-amber-600 shrink-0" />
											{/if}
										</div>
									{/each}
								</div>
							{:else}
								<div class="p-3.5 rounded-xl border border-emerald-200 bg-emerald-50/50 flex items-center justify-between text-xs font-bold text-emerald-900">
									<span>All Mandatory Fields Provided</span>
									<CheckCircle2 class="w-4 h-4 text-emerald-600 shrink-0" />
								</div>
							{/if}

							{#if requirements && !requirements.is_ready_to_apply}
								<a href="/profile" class="w-full py-3 px-4 rounded-xl text-center text-xs font-bold text-white bg-amber-600 hover:bg-amber-700 shadow-md flex items-center justify-center gap-2 transition-colors">
									<ArrowRight class="w-4 h-4" />
									Complete Profile First
								</a>
							{/if}
						</div>
					{:else if authState.isAuthenticated}
						<div class="glass-panel p-6 rounded-[2.5rem] border border-outline-variant/40 bg-white/90 shadow-md space-y-3">
							<h3 class="text-sm font-extrabold text-on-surface flex items-center gap-2">
								<Info class="w-4 h-4 text-primary" />
								Important Note
							</h3>
							<p class="text-xs text-on-surface-variant leading-relaxed">
								Make sure your academic profile, subject marks, photo URL, and signature URL are filled in before applying. Incomplete profiles cannot submit applications.
							</p>
							<a href="/profile" class="block w-full py-2.5 px-4 rounded-xl text-center text-xs font-bold text-white bg-primary hover:bg-primary-container shadow-md transition-all">
								Go to Profile
							</a>
						</div>
					{/if}

					<!-- Program Info Card -->
					<div class="glass-panel p-6 rounded-[2.5rem] border border-outline-variant/40 bg-white/90 shadow-md space-y-3">
						<h3 class="text-sm font-extrabold text-on-surface">Program Summary</h3>
						<div class="space-y-2 text-xs text-on-surface-variant">
							<div class="flex justify-between font-medium">
								<span>University</span>
								<span class="text-on-surface font-extrabold text-right">{program.university_name || program.u_name || 'Public University'}</span>
							</div>
							{#if program.university_location || program.location}
								<div class="flex justify-between font-medium">
									<span>Campus Location</span>
									<span class="text-on-surface font-semibold text-right">{program.university_location || program.location}</span>
								</div>
							{/if}
							<div class="flex justify-between font-medium">
								<span>Unit</span>
								<span class="text-on-surface font-bold">{program.p_unit || 'N/A'}</span>
							</div>
							<div class="flex justify-between font-medium">
								<span>Total Seats</span>
								<span class="text-on-surface font-bold">{program.total_seats}</span>
							</div>
							{#if program.prev_cutmarks}
								<div class="flex justify-between font-medium">
									<span>Previous Cutmark</span>
									<span class="text-primary font-bold">{program.prev_cutmarks}</span>
								</div>
							{/if}
							<div class="flex justify-between font-medium">
								<span>Deadline</span>
								<span class="text-on-surface font-mono font-bold">{program.deadline ? program.deadline.split('T')[0] : 'TBA'}</span>
							</div>
						</div>
					</div>
				</div>
			</div>
		{/if}
	</div>
</div>
