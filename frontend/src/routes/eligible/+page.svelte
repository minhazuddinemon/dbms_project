<!-- src/routes/eligible/+page.svelte -->
<script lang="ts">
	import { fetchEligiblePrograms } from '$lib/api/programs';
	import type { EligibleProgram } from '$lib/types/models';
	import { authState } from '$lib/state/auth.svelte';
	import { Award, CheckCircle2, AlertCircle, ArrowRight, UserCheck, Sparkles, Building2 } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let eligiblePrograms = $state<EligibleProgram[]>([]);
	let isLoading = $state(true);
	let error = $state<string | null>(null);

	async function loadEligible() {
		if (!authState.isAuthenticated) {
			isLoading = false;
			return;
		}
		isLoading = true;
		error = null;
		try {
			eligiblePrograms = await fetchEligiblePrograms();
		} catch (err: any) {
			error = err?.message || 'Failed to fetch eligible programs. Ensure your profile and marks are updated.';
		} finally {
			isLoading = false;
		}
	}

	onMount(() => {
		loadEligible();
	});
</script>

<svelte:head>
	<title>Eligible Universities - UniApp</title>
</svelte:head>

<div class="py-10 bg-slate-50 min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
		<!-- Header -->
		<div class="bg-gradient-to-r from-indigo-900 to-indigo-950 text-white rounded-3xl p-8 shadow-xl relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-emerald-500/10 rounded-full blur-3xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-4">
				<div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold uppercase tracking-wider">
					<Sparkles class="w-4 h-4 text-emerald-400" />
					Smart Match Engine
				</div>
				<h1 class="text-3xl font-extrabold sm:text-4xl text-white">Your Eligible Programs</h1>
				<p class="text-slate-300 text-sm sm:text-base leading-relaxed">
					Based on your SSC & HSC GPA and individual subject marks (Physics, Mathematics, Chemistry), here are the university programs you qualify for right now.
				</p>
			</div>
		</div>

		{#if !authState.isAuthenticated}
			<div class="bg-white rounded-2xl border border-slate-200 p-8 text-center max-w-md mx-auto space-y-4 shadow-sm">
				<UserCheck class="w-12 h-12 text-indigo-600 mx-auto" />
				<h3 class="text-xl font-bold text-slate-900">Sign In to View Eligible Programs</h3>
				<p class="text-slate-600 text-sm">Please log in and update your SSC/HSC academic records to run the rule matching engine.</p>
				<div class="pt-2 flex gap-3 justify-center">
					<a href="/login" class="px-6 py-2.5 rounded-xl font-semibold bg-indigo-600 text-white text-sm hover:bg-indigo-700 transition-colors">Sign In</a>
					<a href="/register" class="px-6 py-2.5 rounded-xl font-semibold border border-slate-300 text-slate-700 text-sm hover:bg-slate-50 transition-colors">Register</a>
				</div>
			</div>
		{:else if isLoading}
			<div class="py-16 text-center text-slate-500">
				<div class="w-10 h-10 border-4 border-indigo-600 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
				<p>Evaluating admission rules against your profile...</p>
			</div>
		{:else if error}
			<div class="bg-white rounded-2xl border border-red-200 p-8 text-center max-w-lg mx-auto space-y-4 shadow-sm">
				<AlertCircle class="w-12 h-12 text-red-500 mx-auto" />
				<h3 class="text-lg font-bold text-slate-900">Action Required</h3>
				<p class="text-slate-600 text-sm">{error}</p>
				<a href="/profile" class="inline-block px-6 py-2.5 rounded-xl font-semibold bg-indigo-600 text-white text-sm hover:bg-indigo-700 transition-colors">
					Update Academic Marks
				</a>
			</div>
		{:else if eligiblePrograms.length === 0}
			<div class="bg-white rounded-2xl border border-slate-200 p-12 text-center max-w-lg mx-auto space-y-4">
				<Award class="w-12 h-12 text-amber-500 mx-auto" />
				<h3 class="text-xl font-bold text-slate-900">No Matched Programs Yet</h3>
				<p class="text-slate-600 text-sm">
					Make sure you have filled out your SSC and HSC subject marks in your Academic Profile.
				</p>
				<a href="/profile" class="inline-block px-6 py-2.5 rounded-xl font-semibold bg-indigo-600 text-white text-sm hover:bg-indigo-700 transition-colors">
					Fill Academic Profile
				</a>
			</div>
		{:else}
			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
				{#each eligiblePrograms as prog}
					<div class="bg-white rounded-2xl border border-emerald-200 p-6 shadow-sm hover:shadow-md transition-all space-y-4 flex flex-col justify-between">
						<div class="space-y-3">
							<div class="flex items-center justify-between">
								<span class="inline-flex items-center gap-1 text-xs font-bold text-emerald-700 bg-emerald-50 px-2.5 py-1 rounded-lg">
									<CheckCircle2 class="w-3.5 h-3.5" />
									Eligible
								</span>
								<span class="text-xs text-slate-400 font-mono">ID: {prog.program_id}</span>
							</div>

							<h3 class="text-xl font-bold text-slate-900">{prog.program_name}</h3>

							<p class="flex items-center gap-2 text-sm text-slate-600">
								<Building2 class="w-4 h-4 text-slate-400" />
								<span>{prog.university_name}</span>
							</p>
						</div>

						<a
							href={`/apply/${prog.program_id}`}
							class="w-full py-2.5 px-4 rounded-xl text-center text-sm font-semibold text-white bg-indigo-600 hover:bg-indigo-700 transition-colors flex items-center justify-center gap-1.5"
						>
							Apply Now
							<ArrowRight class="w-4 h-4" />
						</a>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</div>
