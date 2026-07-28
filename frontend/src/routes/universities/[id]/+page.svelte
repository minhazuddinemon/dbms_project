<!-- src/routes/universities/[id]/+page.svelte -->
<script lang="ts">
	import { fetchUniversityByID, fetchUniversityTransport } from '$lib/api/university';
	import type { University, UniversityTransport, AlbumResponse } from '$lib/types/models';
	import { page } from '$app/state';
	import {
		Building2,
		MapPin,
		Globe,
		ExternalLink,
		BookOpen,
		Image as ImageIcon,
		Bus,
		Clock,
		ArrowLeft,
		History,
		Users,
		X,
		Maximize2
	} from 'lucide-svelte';
	import { onMount } from 'svelte';

	let university = $state<University | null>(null);
	let transportRoutes = $state<UniversityTransport[]>([]);
	let isLoading = $state(true);
	let error = $state<string | null>(null);
	let selectedPhoto = $state<AlbumResponse | { picture_title: string; picture_url: string } | null>(null);

	const uId = parseInt(page.params.id, 10);

	async function loadUniversityDetail() {
		isLoading = true;
		error = null;
		try {
			const [uniData, transportData] = await Promise.all([
				fetchUniversityByID(uId),
				fetchUniversityTransport(uId).catch(() => [])
			]);
			university = uniData;
			transportRoutes = transportData;
		} catch (err: any) {
			error = err?.message || 'Failed to load university details';
		} finally {
			isLoading = false;
		}
	}

	onMount(() => {
		if (uId) {
			loadUniversityDetail();
		}
	});
</script>

<svelte:head>
	<title>{university ? university.u_name : 'University Details'} - UniApp</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		<!-- Back link -->
		<a href="/universities" class="inline-flex items-center gap-2 text-sm font-bold text-outline hover:text-primary transition-colors">
			<ArrowLeft class="w-4 h-4" />
			Back to All Universities
		</a>

		{#if isLoading}
			<div class="py-20 text-center text-on-surface-variant">
				<div class="w-12 h-12 border-4 border-primary border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
				<p class="font-bold text-lg">Loading university details...</p>
			</div>
		{:else if error || !university}
			<div class="glass-panel rounded-[2.5rem] border border-outline-variant/30 p-12 text-center max-w-lg mx-auto space-y-4 bg-white/90 shadow-xl">
				<Building2 class="w-14 h-14 text-error mx-auto" />
				<h3 class="text-2xl font-bold text-on-surface">University Not Found</h3>
				<p class="text-on-surface-variant text-sm">{error || 'The requested university could not be loaded.'}</p>
			</div>
		{:else}
			<!-- Header Card -->
			<div class="bg-gradient-to-r from-primary via-primary-container to-secondary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
				<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

				<div class="relative z-10 flex flex-col md:flex-row items-start md:items-center justify-between gap-6">
					<div class="flex items-center gap-6">
						{#if university.logo_url}
							<img src={university.logo_url} alt={university.u_name} class="w-24 h-24 rounded-3xl object-contain p-2 border-2 border-white/30 shadow-xl bg-white shrink-0" />
						{:else}
							<div class="w-24 h-24 rounded-3xl bg-white/10 text-white flex items-center justify-center font-black text-4xl border-2 border-white/30 backdrop-blur-md shrink-0">
								{university.u_name.charAt(0)}
							</div>
						{/if}

						<div class="space-y-2">
							<h1 class="text-3xl sm:text-5xl font-black text-white leading-tight">
								{university.u_name}
							</h1>
							{#if university.location}
								<p class="flex items-center gap-2 text-sm font-semibold text-slate-100">
									<MapPin class="w-4 h-4 text-tertiary-fixed" />
									{university.location}
								</p>
							{/if}
						</div>
					</div>

					{#if university.website}
						<a
							href={university.website}
							target="_blank"
							rel="noopener noreferrer"
							class="px-6 py-3 rounded-2xl bg-white text-primary hover:bg-slate-100 font-extrabold text-sm shadow-lg transition-all flex items-center gap-2 shrink-0"
						>
							<Globe class="w-4 h-4" />
							Official Website
							<ExternalLink class="w-4 h-4" />
						</a>
					{/if}
				</div>
			</div>

			<!-- Main Content Grid -->
			<div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
				<!-- Left Details (Description, History & Album) -->
				<div class="lg:col-span-2 space-y-8">
					{#if university.university_description}
						<div class="glass-panel p-8 rounded-[2rem] border border-outline-variant/40 bg-white/90 shadow-md space-y-4">
							<h3 class="text-xl font-bold text-on-surface flex items-center gap-2">
								<Building2 class="w-5 h-5 text-primary" />
								About {university.u_name}
							</h3>
							<p class="text-on-surface-variant leading-relaxed text-sm whitespace-pre-line">
								{university.university_description}
							</p>
						</div>
					{/if}

					{#if university.university_history}
						<div class="glass-panel p-8 rounded-[2rem] border border-outline-variant/40 bg-white/90 shadow-md space-y-4">
							<h3 class="text-xl font-bold text-on-surface flex items-center gap-2">
								<History class="w-5 h-5 text-tertiary" />
								University History
							</h3>
							<p class="text-on-surface-variant leading-relaxed text-sm whitespace-pre-line">
								{university.university_history}
							</p>
						</div>
					{/if}

					<!-- Departments List -->
					{#if university.departments && university.departments.length > 0}
						<div class="glass-panel p-8 rounded-[2rem] border border-outline-variant/40 bg-white/90 shadow-md space-y-6">
							<h3 class="text-xl font-bold text-on-surface flex items-center gap-2">
								<BookOpen class="w-5 h-5 text-primary" />
								Departments ({university.departments.length})
							</h3>

							<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
								{#each university.departments as dept}
									<div class="p-5 rounded-2xl border border-outline-variant/30 bg-surface-container-low/50 space-y-2">
										<h4 class="font-extrabold text-on-surface text-base">{dept.dept_name}</h4>
										{#if dept.dept_description}
											<p class="text-xs text-on-surface-variant">{dept.dept_description}</p>
										{/if}
										<div class="flex items-center gap-2 text-xs font-bold text-primary pt-1">
											<Users class="w-3.5 h-3.5" />
											Total Seats: {dept.total_seats}
										</div>
									</div>
								{/each}
							</div>
						</div>
					{/if}

					<!-- Campus Album Gallery (Aspect-Ratio Preserved with Fullscreen Lightbox) -->
					{#if university.album && university.album.length > 0}
						<div class="glass-panel p-8 rounded-[2rem] border border-outline-variant/40 bg-white/90 shadow-md space-y-6">
							<div class="flex items-center justify-between border-b border-outline-variant/30 pb-3">
								<h3 class="text-xl font-bold text-on-surface flex items-center gap-2">
									<ImageIcon class="w-5 h-5 text-primary" />
									Campus Album ({university.album.length})
								</h3>
								<span class="text-xs text-on-surface-variant font-semibold">Click image to enlarge</span>
							</div>

							<div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
								{#each university.album as pic}
									<button
										type="button"
										onclick={() => selectedPhoto = pic}
										class="group relative rounded-3xl overflow-hidden border border-outline-variant/40 shadow-sm bg-surface-container-low/70 hover:shadow-xl hover:border-primary/50 transition-all duration-300 text-left flex flex-col cursor-pointer"
									>
										<div class="w-full aspect-[4/3] bg-slate-900/5 p-3 flex items-center justify-center overflow-hidden relative">
											<img
												src={pic.picture_url}
												alt={pic.picture_title}
												class="w-full h-full object-contain group-hover:scale-105 transition-transform duration-300 rounded-2xl"
											/>
											<div class="absolute top-3 right-3 p-1.5 rounded-xl bg-black/40 text-white backdrop-blur-md opacity-0 group-hover:opacity-100 transition-opacity">
												<Maximize2 class="w-3.5 h-3.5" />
											</div>
										</div>
										<div class="p-3.5 bg-white/90 border-t border-outline-variant/20 flex items-center justify-between w-full">
											<p class="text-on-surface font-extrabold text-xs truncate">{pic.picture_title}</p>
											<span class="text-[10px] font-bold text-primary bg-primary-fixed px-2 py-0.5 rounded-full shrink-0">View</span>
										</div>
									</button>
								{/each}
							</div>
						</div>
					{/if}
				</div>

				<!-- Right Sidebar: Transport Route Tracker -->
				<div class="space-y-6">
					<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/40 bg-white/90 shadow-md space-y-5 sticky top-28">
						<div class="flex items-center justify-between border-b border-outline-variant/30 pb-4">
							<h3 class="text-lg font-extrabold text-on-surface flex items-center gap-2">
								<Bus class="w-5 h-5 text-primary" />
								Transport Route Tracker
							</h3>
						</div>

						<p class="text-xs text-on-surface-variant leading-relaxed">
							Check transport routes and estimated travel times to {university.u_name} campus.
						</p>

						{#if transportRoutes.length === 0}
							<div class="p-6 rounded-2xl bg-surface-container-low/60 border border-outline-variant/30 text-center space-y-2">
								<Bus class="w-8 h-8 text-outline mx-auto" />
								<p class="text-xs font-semibold text-on-surface-variant">No transport routes listed yet.</p>
							</div>
						{:else}
							<div class="space-y-3">
								{#each transportRoutes as route}
									<div class="p-4 rounded-2xl border border-outline-variant/30 bg-surface-container-low/40 space-y-1.5 hover:border-primary/30 transition-all">
										<h5 class="font-bold text-xs text-on-surface">{route.transport_route}</h5>
										<p class="flex items-center gap-1.5 text-xs text-primary font-semibold">
											<Clock class="w-3.5 h-3.5" />
											Est. Travel Time: {route.est_travel_time}
										</p>
									</div>
								{/each}
							</div>
						{/if}
					</div>
				</div>
			</div>
		{/if}
	</div>
</div>

<!-- Lightbox Photo Preview Modal -->
{#if selectedPhoto}
	<div
		class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/80 backdrop-blur-md animate-fade-in cursor-pointer"
		onclick={() => selectedPhoto = null}
		role="button"
		tabindex="0"
		onkeydown={(e) => e.key === 'Escape' && (selectedPhoto = null)}
	>
		<div
			class="relative max-w-4xl w-full bg-white rounded-3xl p-6 shadow-2xl space-y-4 max-h-[90vh] flex flex-col cursor-default"
			onclick={(e) => e.stopPropagation()}
			role="document"
		>
			<div class="flex items-center justify-between border-b pb-3">
				<h4 class="font-extrabold text-base text-on-surface">{selectedPhoto.picture_title}</h4>
				<button onclick={() => selectedPhoto = null} class="p-2 rounded-xl hover:bg-surface-container transition-colors text-on-surface-variant">
					<X class="w-5 h-5" />
				</button>
			</div>
			<div class="flex-1 flex items-center justify-center overflow-hidden p-2 min-h-[300px]">
				<img src={selectedPhoto.picture_url} alt={selectedPhoto.picture_title} class="max-w-full max-h-[70vh] object-contain rounded-2xl shadow-lg" />
			</div>
		</div>
	</div>
{/if}
