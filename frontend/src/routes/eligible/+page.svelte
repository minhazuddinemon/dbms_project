<!-- src/routes/eligible/+page.svelte -->
<script lang="ts">
	import { fetchEligiblePrograms } from '$lib/api/programs';
	import type { EligibleProgram } from '$lib/types/models';
	import { authState } from '$lib/state/auth.svelte';
	import { Award, CheckCircle2, AlertCircle, ArrowRight, UserCheck, Sparkles, Building2, ShieldCheck, Filter, ArrowUpDown, TrendingUp, BookOpen } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let eligiblePrograms = $state<EligibleProgram[]>([]);
	let isLoading = $state(true);
	let error = $state<string | null>(null);
	let selectedFilter = $state<string>('ALL');

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
	<title>Eligible Universities - UniApp Portal</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		
		<!-- Banner Header & Title -->
		<div class="flex flex-col md:flex-row justify-between items-start md:items-end gap-6 bg-gradient-to-r from-tertiary via-tertiary-container to-primary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-tertiary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-3">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<Sparkles class="w-4 h-4" />
					Match Engine Evaluated
				</div>

				<h1 class="text-3xl sm:text-4xl font-black text-white">
					Eligible Universities
				</h1>

				<p class="text-slate-100 text-base leading-relaxed">
					Based on your academic profile and test scores, you meet the initial cutoff criteria for the following institutions.
				</p>
			</div>

			<!-- Filter & Sort Bar -->
			<div class="relative z-10 flex gap-3 w-full sm:w-auto">
				<button class="flex-1 sm:flex-initial px-4 py-2.5 rounded-xl bg-white/15 hover:bg-white/25 text-white font-bold text-xs uppercase tracking-wider backdrop-blur-md border border-white/20 transition-all flex items-center justify-center gap-2">
					<Filter class="w-4 h-4" />
					Filter
				</button>
				<button class="flex-1 sm:flex-initial px-4 py-2.5 rounded-xl bg-white/15 hover:bg-white/25 text-white font-bold text-xs uppercase tracking-wider backdrop-blur-md border border-white/20 transition-all flex items-center justify-center gap-2">
					<ArrowUpDown class="w-4 h-4" />
					Sort by Match
				</button>
			</div>
		</div>

		{#if !authState.isAuthenticated}
			<div class="glass-panel rounded-[2.5rem] border border-outline-variant/30 p-10 text-center max-w-md mx-auto space-y-5 bg-white/90 shadow-xl">
				<div class="w-16 h-16 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center mx-auto">
					<UserCheck class="w-8 h-8" />
				</div>
				<h3 class="text-2xl font-extrabold text-on-surface">Sign In to Evaluate Eligibility</h3>
				<p class="text-on-surface-variant text-sm leading-relaxed">Log in and save your HSC subject marks in your Academic Profile to run the match engine.</p>
				<div class="pt-2 flex gap-3 justify-center">
					<a href="/login" class="px-6 py-3 rounded-xl font-bold bg-primary text-white text-sm hover:bg-primary-container shadow-lg shadow-primary/20 transition-all">Sign In</a>
					<a href="/register" class="px-6 py-3 rounded-xl font-bold border border-outline-variant text-on-surface text-sm hover:bg-surface-container transition-all">Register</a>
				</div>
			</div>
		{:else if isLoading}
			<div class="py-20 text-center text-on-surface-variant">
				<div class="w-12 h-12 border-4 border-primary border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
				<p class="font-bold text-lg">Evaluating university eligibility rules...</p>
			</div>
		{:else if error}
			<div class="glass-panel rounded-[2.5rem] border border-error-container p-10 text-center max-w-lg mx-auto space-y-4 bg-white/95 shadow-xl">
				<div class="w-14 h-14 rounded-2xl bg-error-container text-error flex items-center justify-center mx-auto">
					<AlertCircle class="w-7 h-7" />
				</div>
				<h3 class="text-xl font-bold text-on-surface">Action Required</h3>
				<p class="text-on-surface-variant text-sm">{error}</p>
				<a href="/profile" class="inline-block px-6 py-3 rounded-xl font-bold bg-primary text-white text-sm hover:bg-primary-container shadow-lg shadow-primary/25 transition-all">
					Update Academic Marks
				</a>
			</div>
		{:else if eligiblePrograms.length === 0}
			<div class="glass-panel rounded-[2.5rem] border border-outline-variant/30 p-12 text-center max-w-lg mx-auto space-y-4 bg-white/90 shadow-xl">
				<div class="w-16 h-16 rounded-2xl bg-tertiary-fixed/40 text-tertiary flex items-center justify-center mx-auto">
					<Award class="w-8 h-8" />
				</div>
				<h3 class="text-2xl font-bold text-on-surface">No Matched Programs Yet</h3>
				<p class="text-on-surface-variant text-sm leading-relaxed">
					Ensure you have saved your SSC and HSC subject marks in your Academic Profile.
				</p>
				<a href="/profile" class="inline-block px-6 py-3 rounded-xl font-bold bg-primary text-white text-sm hover:bg-primary-container shadow-lg shadow-primary/25 transition-all">
					Fill Academic Marks &rarr;
				</a>
			</div>
		{:else}
			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
				{#each eligiblePrograms as prog}
					<div class="glass-panel rounded-[2.5rem] border border-outline-variant/40 p-7 shadow-lg hover:shadow-2xl hover:border-primary transition-all duration-300 flex flex-col justify-between space-y-6 bg-white/95 group relative overflow-hidden">
						<!-- Verified Partner Badge -->
						<div class="flex items-center justify-between">
							<span class="inline-flex items-center gap-1.5 text-xs font-bold text-tertiary bg-emerald-50 px-3 py-1 rounded-full border border-emerald-200">
								<ShieldCheck class="w-4 h-4 text-emerald-600" />
								Official Partner
							</span>
							<span class="text-xs font-mono font-bold text-outline">ID #{prog.program_id}</span>
						</div>

						<!-- University & Program Header -->
						<div class="flex items-start gap-4">
							<div class="w-14 h-14 rounded-2xl bg-primary-fixed/40 text-primary border border-primary/20 flex items-center justify-center shrink-0 group-hover:scale-110 transition-transform">
								<Building2 class="w-7 h-7" />
							</div>
							<div>
								<h3 class="text-2xl font-extrabold text-on-surface leading-tight">{prog.university_name || 'BUET'}</h3>
								<p class="text-xs font-semibold text-on-surface-variant">{prog.program_name}</p>
							</div>
						</div>

						<!-- Stats Bar: Seats & Match Score -->
						<div class="grid grid-cols-2 gap-4 p-4 bg-surface-container-low/70 rounded-2xl border border-outline-variant/30">
							<div>
								<p class="text-xs font-bold text-outline uppercase tracking-wider mb-0.5">Total Seats</p>
								<p class="text-lg font-black text-on-surface">1,060</p>
							</div>
							<div>
								<p class="text-xs font-bold text-outline uppercase tracking-wider mb-0.5">Match Score</p>
								<div class="flex items-center gap-1 text-tertiary">
									<TrendingUp class="w-4 h-4 text-emerald-600" />
									<p class="text-lg font-black text-emerald-600">98%</p>
								</div>
							</div>
						</div>

						<!-- Eligible Departments Tags -->
						<div class="space-y-2">
							<p class="text-xs font-bold text-outline uppercase tracking-wider">Eligible Departments</p>
							<div class="flex flex-wrap gap-2">
								<span class="px-2.5 py-1 rounded-lg bg-secondary-fixed/50 text-on-secondary-fixed text-xs font-bold">CSE</span>
								<span class="px-2.5 py-1 rounded-lg bg-primary-fixed/50 text-on-primary-fixed text-xs font-bold">EEE</span>
								<span class="px-2.5 py-1 rounded-lg bg-tertiary-fixed/40 text-on-tertiary-fixed-variant text-xs font-bold">Architecture</span>
							</div>
						</div>

						<!-- Apply Button -->
						<a
							href={`/apply/${prog.program_id}`}
							class="w-full py-3.5 px-4 rounded-xl text-center text-sm font-bold text-white bg-primary hover:bg-primary-container shadow-md shadow-primary/20 hover:shadow-lg transition-all flex items-center justify-center gap-2"
						>
							<span>Submit Application</span>
							<ArrowRight class="w-4 h-4" />
						</a>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</div>

