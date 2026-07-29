<!-- src/routes/programs/+page.svelte -->
<script lang="ts">
	import { fetchPrograms } from '$lib/api/programs';
	import { fetchUniversities } from '$lib/api/university';
	import ProgramCard from '$lib/components/ProgramCard.svelte';
	import type { Program, University } from '$lib/types/models';
	import { Search, Building2 } from 'lucide-svelte';
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
					placeholder="Search by program name (e.g. BSc, Unit A)..."
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
					<ProgramCard program={program} uniLogo={getUniLogo(program)} />
				{/each}
			</div>
		{/if}
	</div>
</div>
