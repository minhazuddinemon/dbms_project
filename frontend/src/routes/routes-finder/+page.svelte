<!-- src/routes/routes-finder/+page.svelte -->
<script lang="ts">
	import { fetchUniversities, fetchUniversityTransport } from '$lib/api/university';
	import type { University, UniversityTransport } from '$lib/types/models';
	import { MapPin, Bus, Clock, Navigation, Building2, Search, AlertCircle, ChevronRight, RefreshCw } from 'lucide-svelte';
	import { onMount } from 'svelte';

	interface UniversityWithRoutes {
		university: University;
		routes: UniversityTransport[];
		isLoading: boolean;
	}

	let universities = $state<University[]>([]);
	let universityData = $state<UniversityWithRoutes[]>([]);
	let isLoading = $state(true);
	let error = $state<string | null>(null);
	let searchQuery = $state('');
	let selectedUniversityId = $state<number | null>(null);

	async function loadAll() {
		isLoading = true;
		error = null;
		try {
			const unis = await fetchUniversities();
			universities = unis || [];
			universityData = universities.map((u) => ({
				university: u,
				routes: [],
				isLoading: false
			}));

			// Load transport for all universities in parallel
			await Promise.all(
				universities.map(async (u, idx) => {
					universityData[idx].isLoading = true;
					try {
						const routes = await fetchUniversityTransport(u.u_id);
						universityData[idx].routes = routes || [];
					} catch {
						universityData[idx].routes = [];
					} finally {
						universityData[idx].isLoading = false;
					}
				})
			);
		} catch (err: any) {
			error = err?.message || 'Failed to load transport data.';
		} finally {
			isLoading = false;
		}
	}

	onMount(() => {
		loadAll();
	});

	let filtered = $derived(
		universityData.filter((d) => {
			if (!searchQuery.trim()) return true;
			const q = searchQuery.toLowerCase();
			return (
				d.university.u_name.toLowerCase().includes(q) ||
				(d.university.location && d.university.location.toLowerCase().includes(q))
			);
		})
	);

	let selectedData = $derived(
		selectedUniversityId !== null
			? universityData.find((d) => d.university.u_id === selectedUniversityId) || null
			: null
	);
</script>

<svelte:head>
	<title>University Transport Routes - UniApp</title>
	<meta name="description" content="Explore transport routes and travel times to universities. Real-time route tracker for Bangladesh public university campuses." />
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">

		<!-- Banner Header -->
		<div class="bg-gradient-to-r from-primary via-secondary to-tertiary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>
			<div class="absolute left-1/2 bottom-0 w-64 h-64 bg-primary-fixed/10 rounded-full blur-2xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-4">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<Navigation class="w-4 h-4" />
					Campus Transit & Directions
				</div>

				<h1 class="text-3xl sm:text-4xl font-black text-white leading-tight">
					University Transport Route Tracker
				</h1>

				<p class="text-slate-100 text-base leading-relaxed">
					View official transport routes and estimated travel times for each public university campus in Bangladesh.
				</p>
			</div>
		</div>

		<!-- Search and Filter -->
		<div class="glass-panel p-4 rounded-[2rem] border border-outline-variant/40 bg-white/90 shadow-xl flex flex-col md:flex-row gap-4 items-center">
			<div class="relative flex-1 w-full">
				<div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none text-outline">
					<Search class="w-5 h-5" />
				</div>
				<input
					type="text"
					bind:value={searchQuery}
					placeholder="Search university or location (e.g. BUET, Dhaka)..."
					class="w-full pl-11 pr-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
				/>
			</div>
			<button
				onclick={loadAll}
				class="flex items-center gap-2 px-5 py-3 rounded-xl font-bold text-white bg-primary text-sm shadow-md hover:bg-primary-container transition-all"
			>
				<RefreshCw class="w-4 h-4" />
				Refresh
			</button>
		</div>

		{#if isLoading}
			<div class="py-20 text-center text-on-surface-variant">
				<div class="w-12 h-12 border-4 border-primary border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
				<p class="font-bold text-lg">Loading transport routes...</p>
			</div>
		{:else if error}
			<div class="glass-panel rounded-[2.5rem] border border-error-container p-10 text-center max-w-lg mx-auto space-y-4 bg-white/95 shadow-xl">
				<AlertCircle class="w-14 h-14 text-error mx-auto" />
				<h3 class="text-xl font-bold text-on-surface">Failed to Load Routes</h3>
				<p class="text-on-surface-variant text-sm">{error}</p>
				<button onclick={loadAll} class="px-6 py-3 rounded-xl font-bold bg-primary text-white text-sm hover:bg-primary-container shadow-lg shadow-primary/25 transition-all">
					Try Again
				</button>
			</div>
		{:else if filtered.length === 0}
			<div class="glass-panel rounded-[2.5rem] border border-outline-variant/30 p-12 text-center max-w-lg mx-auto space-y-4 bg-white/90 shadow-xl">
				<Bus class="w-14 h-14 text-outline mx-auto" />
				<h3 class="text-2xl font-bold text-on-surface">No Universities Found</h3>
				<p class="text-on-surface-variant text-sm">Try a different search keyword.</p>
			</div>
		{:else}
			<div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
				<!-- Left: University List -->
				<div class="lg:col-span-1 space-y-3">
					<h2 class="text-base font-extrabold text-on-surface-variant uppercase tracking-wider px-2">
						Select University ({filtered.length})
					</h2>
					{#each filtered as d}
						<button
							onclick={() => selectedUniversityId = d.university.u_id}
							class="w-full text-left p-5 rounded-2xl border transition-all duration-200 flex items-center gap-4 group {selectedUniversityId === d.university.u_id ? 'border-primary bg-primary-fixed/30 shadow-md' : 'border-outline-variant/40 bg-white/90 hover:bg-surface-container-low hover:border-primary/40'}"
						>
							{#if d.university.logo_url}
								<img src={d.university.logo_url} alt={d.university.u_name} class="w-12 h-12 rounded-2xl object-cover border border-outline-variant/30 shrink-0" />
							{:else}
								<div class="w-12 h-12 rounded-2xl bg-primary-fixed/40 text-primary flex items-center justify-center font-black text-xl border border-primary/20 shrink-0">
									{d.university.u_name.charAt(0)}
								</div>
							{/if}
							<div class="flex-1 min-w-0">
								<h3 class="font-extrabold text-on-surface text-sm truncate">{d.university.u_name}</h3>
								{#if d.university.location}
									<p class="flex items-center gap-1 text-xs text-on-surface-variant mt-0.5">
										<MapPin class="w-3 h-3 text-primary shrink-0" />
										{d.university.location}
									</p>
								{/if}
								<span class="text-xs font-bold text-primary mt-1 block">
									{d.routes.length} {d.routes.length === 1 ? 'route' : 'routes'} listed
								</span>
							</div>
							<ChevronRight class="w-4 h-4 text-outline shrink-0 group-hover:text-primary transition-colors" />
						</button>
					{/each}
				</div>

				<!-- Right: Route Details -->
				<div class="lg:col-span-2">
					{#if selectedData}
						<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
							<!-- University Header -->
							<div class="flex items-center gap-5 border-b border-outline-variant/30 pb-6">
								{#if selectedData.university.logo_url}
									<img src={selectedData.university.logo_url} alt={selectedData.university.u_name} class="w-16 h-16 rounded-3xl object-cover border border-outline-variant/30 shadow-md" />
								{:else}
									<div class="w-16 h-16 rounded-3xl bg-primary-fixed/40 text-primary flex items-center justify-center font-black text-2xl border border-primary/20">
										{selectedData.university.u_name.charAt(0)}
									</div>
								{/if}
								<div>
									<h2 class="text-2xl font-extrabold text-on-surface">{selectedData.university.u_name}</h2>
									{#if selectedData.university.location}
										<p class="flex items-center gap-1.5 text-sm font-semibold text-on-surface-variant mt-1">
											<MapPin class="w-4 h-4 text-primary" />
											{selectedData.university.location}
										</p>
									{/if}
								</div>
							</div>

							<!-- Routes -->
							{#if selectedData.isLoading}
								<div class="py-8 text-center text-on-surface-variant">
									<div class="w-8 h-8 border-4 border-primary border-t-transparent rounded-full animate-spin mx-auto mb-3"></div>
									<p class="text-sm font-semibold">Loading routes...</p>
								</div>
							{:else if selectedData.routes.length === 0}
								<div class="p-8 rounded-2xl bg-surface-container-low/60 border border-outline-variant/30 text-center space-y-3">
									<Bus class="w-10 h-10 text-outline mx-auto" />
									<p class="font-bold text-on-surface-variant text-sm">No transport routes listed for this university yet.</p>
								</div>
							{:else}
								<div class="space-y-4">
									<h3 class="text-base font-extrabold text-on-surface flex items-center gap-2">
										<Bus class="w-5 h-5 text-primary" />
										Transport Routes ({selectedData.routes.length})
									</h3>
									{#each selectedData.routes as route, i}
										<div class="p-5 rounded-2xl border border-outline-variant/30 bg-surface-container-low/50 hover:border-primary/30 hover:bg-primary-fixed/10 transition-all duration-200 space-y-3">
											<div class="flex items-start justify-between gap-4">
												<div class="flex items-start gap-3">
													<div class="w-8 h-8 rounded-xl bg-primary-fixed/50 text-primary flex items-center justify-center font-black text-sm shrink-0 mt-0.5">
														{i + 1}
													</div>
													<div>
														<h4 class="font-extrabold text-on-surface text-sm leading-snug">{route.transport_route}</h4>
													</div>
												</div>
												<span class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full bg-tertiary-fixed/30 text-tertiary text-xs font-bold border border-tertiary/20 shrink-0">
													<Clock class="w-3.5 h-3.5" />
													{route.est_travel_time}
												</span>
											</div>
										</div>
									{/each}
								</div>
							{/if}
						</div>
					{:else}
						<!-- Placeholder state -->
						<div class="glass-panel p-12 rounded-[2.5rem] border border-outline-variant/30 bg-white/90 shadow-xl flex flex-col items-center justify-center text-center space-y-5 min-h-64">
							<div class="w-20 h-20 rounded-3xl bg-primary-fixed/40 text-primary flex items-center justify-center">
								<Bus class="w-10 h-10" />
							</div>
							<h3 class="text-xl font-bold text-on-surface">Select a University</h3>
							<p class="text-on-surface-variant text-sm leading-relaxed max-w-sm">
								Choose a university from the list on the left to view its official transport routes and estimated travel times.
							</p>
						</div>
					{/if}
				</div>
			</div>
		{/if}
	</div>
</div>
