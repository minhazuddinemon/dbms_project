<!-- src/routes/routes-finder/+page.svelte -->
<script lang="ts">
	import { MapPin, Bus, Clock, Navigation, Compass, Building2, Train, ChevronRight, ShieldCheck } from 'lucide-svelte';

	const routes = [
		{
			university: 'Dhaka University (DU)',
			route: 'Farmgate -> Shahbagh -> TSC -> Curzon Hall',
			estTime: '25-35 mins',
			transport: 'Campus Bus (Khyber, Lalon), Metro Rail (TSC Station)',
			color: 'primary'
		},
		{
			university: 'BUET',
			route: 'Nilkhet -> Palashi -> BUET Main Gate',
			estTime: '15-25 mins',
			transport: 'City Bus, Leguna, Rickshaw',
			color: 'tertiary'
		},
		{
			university: 'Jahangirnagar University (JU)',
			route: 'Gabatoli -> Savar -> JU Dairy Gate / Main Gate',
			estTime: '45-60 mins',
			transport: 'Intercity Bus, JU Tarzan Special Shuttle',
			color: 'secondary'
		},
		{
			university: 'Chittagong University (CU)',
			route: 'Sholoshohor Station -> Fatehabad -> CU Station',
			estTime: '45 mins',
			transport: 'Shuttle Train (Demu)',
			color: 'primary'
		},
		{
			university: 'Rajshahi University (RU)',
			route: 'Kazla -> Binodpur -> RU Main Gate',
			estTime: '15-20 mins',
			transport: 'Auto Rickshaw, University Bus',
			color: 'tertiary'
		}
	];
</script>

<svelte:head>
	<title>University Transport Routes - UniApp</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		
		<!-- Banner Header -->
		<div class="bg-gradient-to-r from-primary via-secondary to-tertiary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-3">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<Navigation class="w-4 h-4" />
					Campus Transit & Directions
				</div>

				<h1 class="text-3xl sm:text-4xl font-black text-white">
					University Transport Route Tracker
				</h1>

				<p class="text-slate-100 text-base leading-relaxed">
					Estimated travel times, transit routes, and shuttle train schedules for major exam centers in Bangladesh.
				</p>
			</div>
		</div>

		<!-- Routes Grid -->
		<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
			{#each routes as item}
				<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/90 shadow-sm hover:shadow-xl transition-all duration-300 space-y-5 card-hover">
					<div class="flex items-start justify-between">
						<div class="space-y-1">
							<h3 class="text-2xl font-black text-on-surface flex items-center gap-2">
								<Building2 class="w-6 h-6 text-primary" />
								{item.university}
							</h3>
						</div>
						<span class="px-3 py-1 rounded-full bg-tertiary-fixed/30 text-tertiary text-xs font-bold flex items-center gap-1.5 border border-tertiary/20 shrink-0">
							<Clock class="w-3.5 h-3.5" />
							{item.estTime}
						</span>
					</div>

					<div class="space-y-3 text-sm">
						<div class="p-4 rounded-2xl bg-surface-container-low/70 border border-outline-variant/30 space-y-1">
							<span class="text-xs font-extrabold uppercase tracking-wider text-outline block">Transit Route:</span>
							<p class="font-bold text-on-surface leading-relaxed flex items-center gap-2">
								<MapPin class="w-4 h-4 text-emerald-500 shrink-0" />
								{item.route}
							</p>
						</div>

						<div class="p-4 rounded-2xl bg-surface-container-low/70 border border-outline-variant/30 space-y-1">
							<span class="text-xs font-extrabold uppercase tracking-wider text-outline block">Recommended Mode:</span>
							<p class="font-bold text-on-surface leading-relaxed flex items-center gap-2">
								<Bus class="w-4 h-4 text-primary shrink-0" />
								{item.transport}
							</p>
						</div>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>
