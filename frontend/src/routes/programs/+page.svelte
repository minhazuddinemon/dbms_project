<!-- src/routes/programs/+page.svelte -->
<script lang="ts">
	import { fetchPrograms } from '$lib/api/programs';
	import type { Program } from '$lib/types/models';
	import { Search, Building2, Calendar, Users, Filter, ArrowRight, Award } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let programs = $state<Program[]>([]);
	let search = $state('');
	let selectedUnit = $state('');
	let isLoading = $state(true);

	async function loadPrograms() {
		isLoading = true;
		try {
			programs = await fetchPrograms(search, selectedUnit);
		} catch (err) {
			console.error('Failed to load programs:', err);
		} finally {
			isLoading = false;
		}
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

<div class="py-10 bg-slate-50 min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
		<!-- Header -->
		<div>
			<h1 class="text-3xl font-extrabold text-slate-900 sm:text-4xl">Explore Academic Programs</h1>
			<p class="mt-2 text-slate-600">Search programs across Bangladesh public universities and view admission details</p>
		</div>

		<!-- Search & Filter Bar -->
		<form onsubmit={handleSearch} class="bg-white p-4 rounded-2xl border border-slate-200 shadow-sm flex flex-col md:flex-row gap-4 items-center">
			<div class="relative flex-1 w-full">
				<div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-400">
					<Search class="w-5 h-5" />
				</div>
				<input
					type="text"
					bind:value={search}
					placeholder="Search by program name (e.g. Computer Science, Electrical)..."
					class="w-full pl-10 pr-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-indigo-600 focus:border-indigo-600 text-sm"
				/>
			</div>

			<div class="flex items-center gap-3 w-full md:w-auto">
				<div class="relative w-full md:w-48">
					<select
						bind:value={selectedUnit}
						onchange={loadPrograms}
						class="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-indigo-600 focus:border-indigo-600 text-sm bg-white"
					>
						<option value="">All Units</option>
						<option value="A">Unit A (Engineering / Science)</option>
						<option value="B">Unit B (Humanities / Arts)</option>
						<option value="C">Unit C (Commerce / Business)</option>
					</select>
				</div>

				<button
					type="submit"
					class="px-6 py-3 rounded-xl font-semibold text-white bg-indigo-600 hover:bg-indigo-700 shadow-md shadow-indigo-600/20 text-sm transition-all whitespace-nowrap"
				>
					Search
				</button>
			</div>
		</form>

		<!-- Programs Grid -->
		{#if isLoading}
			<div class="py-16 text-center text-slate-500">
				<div class="w-10 h-10 border-4 border-indigo-600 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
				<p>Loading programs...</p>
			</div>
		{:else if programs.length === 0}
			<div class="py-16 text-center bg-white rounded-2xl border border-slate-200 p-8 space-y-4">
				<Building2 class="w-12 h-12 text-slate-400 mx-auto" />
				<h3 class="text-lg font-bold text-slate-800">No programs found</h3>
				<p class="text-slate-500 text-sm">Try adjusting your search keyword or unit filter.</p>
			</div>
		{:else}
			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
				{#each programs as program}
					<div class="bg-white rounded-2xl border border-slate-200 p-6 hover:border-indigo-300 hover:shadow-lg hover:shadow-indigo-500/5 transition-all flex flex-col justify-between space-y-4">
						<div class="space-y-3">
							<div class="flex items-start justify-between gap-2">
								<span class="px-2.5 py-1 rounded-lg bg-indigo-50 text-indigo-700 text-xs font-bold uppercase tracking-wider">
									Unit {program.p_unit || 'A'}
								</span>
								{#if program.prev_cutmarks}
									<span class="px-2.5 py-1 rounded-lg bg-emerald-50 text-emerald-700 text-xs font-semibold flex items-center gap-1">
										<Award class="w-3.5 h-3.5" />
										Cutmark: {program.prev_cutmarks}
									</span>
								{/if}
							</div>

							<h3 class="text-xl font-bold text-slate-900 group-hover:text-indigo-600 transition-colors">
								{program.p_name}
							</h3>

							<div class="space-y-2 text-sm text-slate-600 pt-1">
								<p class="flex items-center gap-2">
									<Building2 class="w-4 h-4 text-slate-400 shrink-0" />
									<span class="font-medium text-slate-800">{program.u_name || 'Public University'}</span>
								</p>
								<p class="flex items-center gap-2">
									<Users class="w-4 h-4 text-slate-400 shrink-0" />
									<span>Total Seats: <strong class="text-slate-800">{program.total_seats}</strong></span>
								</p>
								<p class="flex items-center gap-2">
									<Calendar class="w-4 h-4 text-slate-400 shrink-0" />
									<span>Deadline: <strong class="text-slate-800">{program.deadline}</strong></span>
								</p>
							</div>
						</div>

						<div class="pt-4 border-t border-slate-100 flex items-center justify-between">
							<a
								href={`/apply/${program.program_id}`}
								class="w-full py-2.5 px-4 rounded-xl text-center text-sm font-semibold text-white bg-indigo-600 hover:bg-indigo-700 transition-colors flex items-center justify-center gap-1.5"
							>
								Apply Now
								<ArrowRight class="w-4 h-4" />
							</a>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</div>
