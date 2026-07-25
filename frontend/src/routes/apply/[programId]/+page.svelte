<!-- src/routes/apply/[programId]/+page.svelte -->
<script lang="ts">
	import { applyToProgram } from '$lib/api/applications';
	import { fetchProgramByID } from '$lib/api/programs';
	import { authState } from '$lib/state/auth.svelte';
	import type { Program } from '$lib/types/models';
	import { page } from '$app/state';
	import { Building2, CheckCircle2, AlertCircle, Calendar, Users, Award, ShieldAlert, ArrowLeft } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let programId = $derived(Number(page.params.programId));
	let program = $state<Program | null>(null);
	let isLoading = $state(true);
	let isSubmitting = $state(false);

	let successMessage = $state<string | null>(null);
	let errorMessage = $state<string | null>(null);
	let missingFields = $state<string[]>([]);

	async function loadProgram() {
		isLoading = true;
		try {
			program = await fetchProgramByID(programId);
		} catch (err) {
			console.error('Failed to load program:', err);
		} finally {
			isLoading = false;
		}
	}

	onMount(() => {
		loadProgram();
	});

	async function handleApply() {
		if (!authState.isAuthenticated) {
			errorMessage = 'Please sign in to submit your program application.';
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
</script>

<svelte:head>
	<title>Apply for Program - UniApp</title>
</svelte:head>

<div class="py-10 bg-slate-50 min-h-screen">
	<div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
		<a href="/programs" class="inline-flex items-center gap-1.5 text-sm font-semibold text-slate-600 hover:text-indigo-600">
			<ArrowLeft class="w-4 h-4" />
			Back to Programs
		</a>

		{#if isLoading}
			<div class="py-16 text-center text-slate-500">
				<div class="w-10 h-10 border-4 border-indigo-600 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
				<p>Loading program details...</p>
			</div>
		{:else if !program}
			<div class="bg-white rounded-2xl border border-slate-200 p-8 text-center space-y-4">
				<AlertCircle class="w-12 h-12 text-red-500 mx-auto" />
				<h3 class="text-xl font-bold text-slate-900">Program Not Found</h3>
				<p class="text-slate-600 text-sm">The program ID requested does not exist or has been removed.</p>
			</div>
		{:else}
			<div class="bg-white rounded-2xl border border-slate-200 p-8 shadow-sm space-y-6">
				<!-- Header info -->
				<div class="space-y-3 border-b border-slate-100 pb-6">
					<div class="flex items-center justify-between">
						<span class="px-3 py-1 bg-indigo-50 text-indigo-700 text-xs font-bold uppercase rounded-lg">
							Unit {program.p_unit || 'A'}
						</span>
						<span class="text-xs text-slate-400 font-mono">Program ID: #{program.program_id}</span>
					</div>

					<h1 class="text-2xl sm:text-3xl font-extrabold text-slate-900">{program.p_name}</h1>

					<div class="flex flex-wrap gap-4 text-sm text-slate-600 pt-2">
						<span class="flex items-center gap-1.5 font-medium text-slate-800">
							<Building2 class="w-4 h-4 text-slate-400" />
							{program.u_name || 'Public University'}
						</span>
						<span class="flex items-center gap-1.5">
							<Users class="w-4 h-4 text-slate-400" />
							{program.total_seats} Total Seats
						</span>
						<span class="flex items-center gap-1.5">
							<Calendar class="w-4 h-4 text-slate-400" />
							Deadline: {program.deadline}
						</span>
					</div>
				</div>

				<!-- Alerts -->
				{#if successMessage}
					<div class="p-6 rounded-2xl bg-emerald-50 border border-emerald-200 text-center space-y-3">
						<CheckCircle2 class="w-12 h-12 text-emerald-500 mx-auto" />
						<h3 class="text-xl font-bold text-emerald-900">Application Submitted!</h3>
						<p class="text-emerald-700 text-sm">{successMessage}</p>
					</div>
				{/if}

				{#if errorMessage}
					<div class="p-6 rounded-2xl bg-amber-50 border border-amber-200 space-y-3">
						<div class="flex items-start gap-3">
							<ShieldAlert class="w-6 h-6 text-amber-600 shrink-0 mt-0.5" />
							<div>
								<h3 class="text-lg font-bold text-amber-900">Incomplete Profile Information</h3>
								<p class="text-amber-700 text-sm mt-1">{errorMessage}</p>
							</div>
						</div>

						{#if missingFields.length > 0}
							<div class="bg-white/80 rounded-xl p-4 border border-amber-200/80 mt-3">
								<span class="block text-xs font-bold text-amber-900 uppercase tracking-wider mb-2">
									Missing Required Parameters:
								</span>
								<ul class="list-disc list-inside space-y-1 text-xs font-mono text-amber-800">
									{#each missingFields as field}
										<li>{field}</li>
									{/each}
								</ul>
								<div class="mt-4 pt-3 border-t border-amber-200/60">
									<a
										href="/profile"
										class="inline-block px-4 py-2 rounded-lg bg-amber-600 hover:bg-amber-700 text-white text-xs font-bold transition-colors"
									>
										Fill Missing Profile Fields &rarr;
									</a>
								</div>
							</div>
						{/if}
					</div>
				{/if}

				<!-- Apply Button -->
				{#if !successMessage}
					<div class="pt-4 border-t border-slate-100 flex items-center justify-between">
						<div>
							<span class="block text-xs text-slate-500">Submission Fee</span>
							<span class="text-lg font-extrabold text-slate-900">BDT 1,000</span>
						</div>

						<button
							onclick={handleApply}
							disabled={isSubmitting}
							class="px-8 py-3 rounded-xl font-bold text-white bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 shadow-lg shadow-indigo-600/25 transition-all text-sm"
						>
							{isSubmitting ? 'Submitting Application...' : 'Submit Application'}
						</button>
					</div>
				{/if}
			</div>
		{/if}
	</div>
</div>
