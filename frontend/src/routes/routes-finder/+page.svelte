<!-- src/routes/routes-finder/+page.svelte -->
<script lang="ts">
	import { MapPin, Bus, Clock, Navigation, Compass, Building2 } from 'lucide-svelte';

	const routes = [
		{
			university: 'Dhaka University (DU)',
			route: 'Farmgate -> Shahbagh -> TSC -> Curzon Hall',
			estTime: '25-35 mins',
			transport: 'Campus Bus, Metro Rail (TSC Station)'
		},
		{
			university: 'BUET',
			route: 'Nilkhet -> Palashi -> BUET Main Gate',
			estTime: '15-25 mins',
			transport: 'City Bus, CNG, Rickshaw'
		},
		{
			university: 'Jahangirnagar University (JU)',
			route: 'Gabatoli -> Savar -> JU Main Gate',
			estTime: '45-60 mins',
			transport: 'Intercity Bus, JU Special Shuttle'
		},
		{
			university: 'Chittagong University (CU)',
			route: 'Sholoshohor Station -> Fatehabad -> CU Station',
			estTime: '45 mins',
			transport: 'Shuttle Train'
		},
		{
			university: 'Rajshahi University (RU)',
			route: 'Kazla -> Binodpur -> RU Main Gate',
			estTime: '15-20 mins',
			transport: 'Auto Rickshaw, University Bus'
		}
	];
</script>

<svelte:head>
	<title>University Route Tracker - UniApp</title>
</svelte:head>

<div class="py-10 bg-slate-50 min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
		<!-- Header -->
		<div>
			<div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-indigo-100 text-indigo-700 text-xs font-semibold uppercase tracking-wider mb-2">
				<Navigation class="w-4 h-4" />
				Campus Connectivity
			</div>
			<h1 class="text-3xl font-extrabold text-slate-900 sm:text-4xl">University Transport & Route Finder</h1>
			<p class="mt-2 text-slate-600">Estimated travel times and recommended transit routes for major university exam centers</p>
		</div>

		<!-- Routes List -->
		<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
			{#each routes as item}
				<div class="bg-white rounded-2xl border border-slate-200 p-6 shadow-sm hover:shadow-md transition-all space-y-4">
					<div class="flex items-start justify-between">
						<h3 class="text-xl font-bold text-slate-900 flex items-center gap-2">
							<Building2 class="w-5 h-5 text-indigo-600" />
							{item.university}
						</h3>
						<span class="px-2.5 py-1 rounded-lg bg-slate-100 text-slate-700 text-xs font-semibold flex items-center gap-1">
							<Clock class="w-3.5 h-3.5 text-indigo-600" />
							{item.estTime}
						</span>
					</div>

					<div class="space-y-2 text-sm text-slate-600">
						<div class="flex items-start gap-2">
							<MapPin class="w-4 h-4 text-emerald-500 shrink-0 mt-0.5" />
							<div>
								<span class="font-semibold text-slate-800">Key Route:</span>
								<p class="text-slate-600">{item.route}</p>
							</div>
						</div>

						<div class="flex items-start gap-2 pt-1">
							<Bus class="w-4 h-4 text-indigo-500 shrink-0 mt-0.5" />
							<div>
								<span class="font-semibold text-slate-800">Recommended Transport:</span>
								<p class="text-slate-600">{item.transport}</p>
							</div>
						</div>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>
