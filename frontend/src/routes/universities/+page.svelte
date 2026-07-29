<!-- src/routes/universities/+page.svelte -->
<script lang="ts">
	import { fetchUniversities } from '$lib/api/university';
	import type { University } from '$lib/types/models';
	import { Search, Building2, MapPin, Globe, ArrowRight, ExternalLink } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let universities = $state<University[]>([]);
	let search = $state('');
	let isLoading = $state(true);

	async function loadUniversities() {
		isLoading = true;
		try {
			universities = await fetchUniversities();
		} catch (err) {
			console.error('Failed to load universities:', err);
		} finally {
			isLoading = false;
		}
	}

	onMount(() => {
		loadUniversities();
	});

	let filteredUniversities = $derived(
		universities.filter((u) => {
			const term = search.toLowerCase();
			return (
				u.u_name.toLowerCase().includes(term) ||
				(u.location && u.location.toLowerCase().includes(term)) ||
				(u.university_description && u.university_description.toLowerCase().includes(term))
			);
		})
	);
</script>

<svelte:head>
	<title>Public Universities - UniApp</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		<!-- Header Banner -->
		<div class="bg-gradient-to-r from-primary via-primary-container to-secondary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-3">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<Building2 class="w-4 h-4" />
					University Directory
				</div>

				<h1 class="text-3xl sm:text-4xl font-black text-white">
					Partner Public Universities
				</h1>

				<p class="text-slate-100 text-base leading-relaxed">
					Explore top public universities in Bangladesh, view campus departments, picture galleries, and access transport route details.
				</p>
			</div>
		</div>

		<!-- Search Bar -->
		<div class="glass-panel p-4 rounded-[2rem] border border-outline-variant/40 bg-white/90 shadow-xl flex items-center">
			<div class="relative flex-1 w-full">
				<div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none text-outline">
					<Search class="w-5 h-5" />
				</div>
				<input
					type="text"
					bind:value={search}
					placeholder="Search university by name or location (e.g. BUET, Dhaka, Chittagong)..."
					class="w-full pl-11 pr-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
				/>
			</div>
		</div>

		<!-- Universities Grid -->
		{#if isLoading}
			<div class="py-20 text-center text-on-surface-variant">
				<div class="w-12 h-12 border-4 border-primary border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
				<p class="font-bold text-lg">Loading universities...</p>
			</div>
		{:else if filteredUniversities.length === 0}
			<div class="glass-panel rounded-[2.5rem] border border-outline-variant/30 p-12 text-center max-w-lg mx-auto space-y-4 bg-white/90 shadow-xl">
				<Building2 class="w-14 h-14 text-outline mx-auto" />
				<h3 class="text-2xl font-bold text-on-surface">No Universities Found</h3>
				<p class="text-on-surface-variant text-sm">Try adjusting your search keyword.</p>
			</div>
		{:else}
			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
				{#each filteredUniversities as uni}
					<div class="glass-panel rounded-[2.5rem] border border-outline-variant/40 p-7 shadow-sm hover:shadow-xl hover:border-primary/40 transition-all duration-300 flex flex-col justify-between space-y-6 bg-white/95 card-hover group">
						<div class="space-y-5">
							<!-- Header with Logo, Name, Location & Official Website Link -->
							<div class="flex flex-col sm:flex-row sm:items-center justify-between gap-3 border-b border-outline-variant/20 pb-4">
								<div class="flex items-center gap-3.5">
									{#if uni.logo_url}
										<img src={uni.logo_url} alt={uni.u_name} class="w-14 h-14 rounded-2xl object-contain p-1 border border-outline-variant/30 shadow-sm bg-white shrink-0" />
									{:else}
										<div class="w-14 h-14 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center font-black text-xl border border-primary/20 shrink-0">
											{uni.u_name.charAt(0)}
										</div>
									{/if}
									<div>
										<h3 class="text-2xl font-black text-on-surface leading-tight group-hover:text-primary transition-colors">
											{uni.u_name}
										</h3>
										{#if uni.location}
											<p class="flex items-center gap-1 text-xs font-bold text-on-surface-variant mt-0.5">
												<MapPin class="w-3.5 h-3.5 text-tertiary shrink-0" />
												<span>{uni.location}</span>
											</p>
										{/if}
									</div>
								</div>

								{#if uni.website}
									<a
										href={uni.website}
										target="_blank"
										rel="noopener noreferrer"
										class="inline-flex items-center gap-1.5 px-3.5 py-2 rounded-xl bg-primary-fixed/50 hover:bg-primary-fixed text-primary text-xs font-extrabold transition-all border border-primary/20 self-start sm:self-auto shadow-xs hover:shadow-sm"
										title="Visit Official University Website"
									>
										<Globe class="w-3.5 h-3.5 text-primary shrink-0" />
										<span>Official Website</span>
										<ExternalLink class="w-3 h-3 text-primary/80 shrink-0" />
									</a>
								{/if}
							</div>

							<!-- Both Description & History -->
							<div class="space-y-3">
								{#if uni.university_description}
									<div class="space-y-1 text-xs text-on-surface-variant line-clamp-3 leading-relaxed">
										{@html uni.university_description}
									</div>
								{/if}

								{#if uni.university_history}
									<div class="space-y-1">
										<span class="text-[10px] font-mono font-bold uppercase tracking-wider text-outline">History</span>
										<div class="text-xs text-on-surface-variant line-clamp-3 leading-relaxed">
											{@html uni.university_history}
										</div>
									</div>
								{/if}
							</div>
						</div>

						<div class="pt-2">
							<a
								href={`/universities/${uni.u_id}`}
								class="w-full py-3.5 px-4 rounded-2xl text-center text-xs font-black text-white bg-primary hover:bg-primary-container shadow-md shadow-primary/20 hover:shadow-lg transition-all flex items-center justify-center gap-2 group-hover:shadow-primary/25"
							>
								<span>View University Details</span>
								<ArrowRight class="w-4 h-4" />
							</a>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</div>
