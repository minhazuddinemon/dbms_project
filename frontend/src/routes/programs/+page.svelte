<!-- src/routes/programs/+page.svelte -->
<script lang="ts">
	import { fetchPrograms } from '$lib/api/programs';
	import { fetchUniversities } from '$lib/api/university';
	import type { Program, University } from '$lib/types/models';
	import { Search, Building2, Calendar, Users, Filter, ArrowRight, Award, Sparkles, MapPin } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { page } from '$app/state';

	let programs = $state<Program[]>([]);
	let universities = $state<University[]>([]);
	let search = $state(page.url.searchParams.get('search') || '');
	let selectedUnit = $state('');
	let isLoading = $state(true);

	async function loadPrograms() {
		isLoading = true;
		try {
			const [pList, uList] = await Promise.all([
				fetchPrograms(search, selectedUnit),
				universities.length > 0 ? Promise.resolve(universities) : fetchUniversities().catch(() => [])
			]);
			programs = pList || [];
			universities = uList || [];
		} catch (err) {
			console.error('Failed to load programs:', err);
		} finally {
			isLoading = false;
		}
	}

	function getUniLogo(prog: Program): string | null {
		if (prog.logo_url) return prog.logo_url;
		if (prog.university_logo) return prog.university_logo;
		let found = universities.find((u) => u.u_id === prog.u_id);
		if (found?.logo_url) return found.logo_url;
		const nameToMatch = prog.university_name || prog.u_name;
		if (nameToMatch) {
			found = universities.find((u) => u.u_name.toLowerCase() === nameToMatch.toLowerCase() || nameToMatch.toLowerCase().includes(u.u_name.toLowerCase()));
			if (found?.logo_url) return found.logo_url;
		}
		return null;
	}

	onMount(() => {
		loadPrograms();
	});

	function handleSearch(e: Event) {
		e.preventDefault();
		loadPrograms();
	}
</script>

<svelte:head>
	<title>Browse Programs - UniApp</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		<!-- Banner Header -->
		<div class="bg-gradient-to-r from-primary via-primary-container to-secondary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-3">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<Search class="w-4 h-4" />
					Academic Directory
				</div>

				<h1 class="text-3xl sm:text-4xl font-black text-white">
					Explore Academic Programs
				</h1>

				<p class="text-slate-100 text-base leading-relaxed">
					Filter public university programs in Bangladesh by Unit A, B, or C, search keywords, and compare seat availability.
				</p>
			</div>
		</div>

		<!-- Search & Filter Bar -->
		<form onsubmit={handleSearch} class="glass-panel p-4 rounded-[2rem] border border-outline-variant/40 bg-white/90 shadow-xl flex flex-col md:flex-row gap-4 items-center">
			<div class="relative flex-1 w-full">
				<div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none text-outline">
					<Search class="w-5 h-5" />
				</div>
				<input
					type="text"
					bind:value={search}
					placeholder="Search by program name (e.g. Computer Science, Electrical)..."
					class="w-full pl-11 pr-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
				/>
			</div>

			<div class="flex items-center gap-3 w-full md:w-auto">
				<div class="relative w-full md:w-48">
					<select
						bind:value={selectedUnit}
						onchange={loadPrograms}
						class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white font-semibold text-on-surface"
					>
						<option value="">All Units</option>
						<option value="A">Unit A (Engineering / Science)</option>
						<option value="B">Unit B (Humanities / Arts)</option>
						<option value="C">Unit C (Commerce / Business)</option>
					</select>
				</div>

				<button
					type="submit"
					class="px-8 py-3 rounded-xl font-bold text-white bg-primary hover:bg-primary-container shadow-md shadow-primary/25 hover:shadow-primary/40 transition-all text-sm whitespace-nowrap"
				>
					Search
				</button>
			</div>
		</form>

		<!-- Programs Grid -->
		{#if isLoading}
			<div class="py-20 text-center text-on-surface-variant">
				<div class="w-12 h-12 border-4 border-primary border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
				<p class="font-bold text-lg">Loading available programs...</p>
			</div>
		{:else if programs.length === 0}
			<div class="glass-panel rounded-[2.5rem] border border-outline-variant/30 p-12 text-center max-w-lg mx-auto space-y-4 bg-white/90 shadow-xl">
				<Building2 class="w-14 h-14 text-outline mx-auto" />
				<h3 class="text-2xl font-bold text-on-surface">No Programs Found</h3>
				<p class="text-on-surface-variant text-sm">Try adjusting your search keyword or unit filter.</p>
			</div>
		{:else}
			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
				{#each programs as program}
					<div class="glass-panel rounded-[2rem] border border-outline-variant/40 p-7 shadow-sm hover:shadow-xl hover:border-primary/40 transition-all duration-300 flex flex-col justify-between space-y-6 bg-white/90 card-hover">
						<div class="space-y-4">
							<!-- Header with Logo and Unit -->
							<div class="flex items-center justify-between gap-3 border-b border-outline-variant/20 pb-3">
								<div class="flex items-center gap-3">
									{#if getUniLogo(program)}
										<img
											src={getUniLogo(program)}
											alt={program.university_name || program.u_name}
											class="w-10 h-10 rounded-xl object-contain p-1 border border-outline-variant/30 bg-white shrink-0"
										/>
									{:else}
										<div class="w-10 h-10 rounded-xl bg-primary-fixed text-primary flex items-center justify-center font-black text-lg shrink-0">
											{(program.university_name || program.u_name || 'U').charAt(0)}
										</div>
									{/if}
									<div>
										<span class="block text-xs font-black text-primary uppercase leading-tight truncate">
											{program.university_name || program.u_name || 'Public University'}
										</span>
										{#if program.university_location || program.location}
											<span class="text-[11px] font-semibold text-on-surface-variant flex items-center gap-1">
												<MapPin class="w-3 h-3 text-tertiary shrink-0" />
												{program.university_location || program.location}
											</span>
										{/if}
									</div>
								</div>

								<span class="px-3 py-1 rounded-full bg-primary-fixed text-on-primary-fixed text-xs font-extrabold uppercase shrink-0">
									Unit {program.p_unit || 'A'}
								</span>
							</div>

							<h3 class="text-2xl font-extrabold text-on-surface leading-tight">
								{program.p_name}
							</h3>

							<div class="space-y-2 text-sm text-on-surface-variant pt-1">
								{#if program.prev_cutmarks}
									<p class="flex items-center gap-2">
										<Award class="w-4 h-4 text-tertiary shrink-0" />
										<span>Previous Cutmark: <strong class="text-primary font-bold">{program.prev_cutmarks}</strong></span>
									</p>
								{/if}
								<p class="flex items-center gap-2">
									<Users class="w-4 h-4 text-outline shrink-0" />
									<span>Total Seats: <strong class="text-on-surface font-bold">{program.total_seats}</strong></span>
								</p>
								<p class="flex items-center gap-2">
									<Calendar class="w-4 h-4 text-outline shrink-0" />
									<span>Deadline: <strong class="text-on-surface font-mono font-bold">{program.deadline ? program.deadline.split('T')[0] : 'N/A'}</strong></span>
								</p>
							</div>
						</div>

						<div class="pt-4 border-t border-outline-variant/30">
							<a
								href={`/apply/${program.program_id}`}
								class="w-full py-3 px-4 rounded-xl text-center text-sm font-bold text-white bg-primary hover:bg-primary-container shadow-md shadow-primary/20 hover:shadow-lg transition-all flex items-center justify-center gap-2"
							>
								<span>Apply Now</span>
								<ArrowRight class="w-4 h-4" />
							</a>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</div>
