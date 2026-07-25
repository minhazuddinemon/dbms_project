<!-- src/routes/merit/+page.svelte -->
<script lang="ts">
	import { Award, CheckCircle2, Search, Building2, TrendingUp, Medal, Sparkles } from 'lucide-svelte';

	let searchUnit = $state('A');

	const meritData = [
		{
			university: 'Dhaka University (DU)',
			unit: 'Unit A (Science)',
			meritPosition: 450,
			status: 'Selected for Admission',
			subjectAllocated: 'Computer Science & Engineering',
			cutmark: '82.50 / 100'
		},
		{
			university: 'BUET',
			unit: 'Engineering',
			meritPosition: 820,
			status: 'Eligible for Viva',
			subjectAllocated: 'Electrical & Electronic Engineering',
			cutmark: '78.00 / 100'
		},
		{
			university: 'Jahangirnagar University (JU)',
			unit: 'Unit A',
			meritPosition: 140,
			status: 'Selected for Admission',
			subjectAllocated: 'Institute of Information Technology',
			cutmark: '85.20 / 100'
		},
		{
			university: 'Rajshahi University (RU)',
			unit: 'Unit C',
			meritPosition: 310,
			status: 'Waiting List',
			subjectAllocated: 'Software Engineering',
			cutmark: '76.40 / 100'
		}
	];
</script>

<svelte:head>
	<title>Merit List & Cutoff Analysis - UniApp</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		
		<!-- Banner Header -->
		<div class="bg-gradient-to-r from-secondary via-primary to-primary-container text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-3">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<Medal class="w-4 h-4" />
					Merit Position Tracker
				</div>

				<h1 class="text-3xl sm:text-4xl font-black text-white">
					Results & Cutoff Analysis
				</h1>

				<p class="text-slate-100 text-base leading-relaxed">
					Track your merit rank across public universities, check historical cutoff scores, and view subject allocations.
				</p>
			</div>
		</div>

		<!-- Merit Cards Grid -->
		<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
			{#each meritData as item}
				<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/90 shadow-sm hover:shadow-xl transition-all duration-300 space-y-6 card-hover">
					<div class="flex items-start justify-between">
						<div>
							<span class="px-3 py-1 bg-primary-fixed text-on-primary-fixed text-xs font-bold uppercase rounded-lg">
								{item.unit}
							</span>
							<h3 class="text-2xl font-extrabold text-on-surface mt-2">{item.university}</h3>
						</div>
						<div class="text-right">
							<span class="block text-xs font-bold text-outline uppercase tracking-wider">Merit Rank</span>
							<span class="text-3xl font-black text-primary">#{item.meritPosition}</span>
						</div>
					</div>

					<div class="p-5 rounded-2xl bg-surface-container-low/70 border border-outline-variant/30 space-y-3">
						<div class="flex items-center justify-between text-sm">
							<span class="font-semibold text-on-surface-variant">Allocated Subject:</span>
							<strong class="text-on-surface font-extrabold">{item.subjectAllocated}</strong>
						</div>
						<div class="flex items-center justify-between text-sm">
							<span class="font-semibold text-on-surface-variant">Previous Cutmark:</span>
							<span class="text-tertiary font-bold">{item.cutmark}</span>
						</div>
					</div>

					<div class="pt-2 flex items-center justify-between">
						<span class="inline-flex items-center gap-1.5 text-xs font-bold text-tertiary bg-tertiary-fixed/40 px-3 py-1 rounded-full border border-tertiary/20">
							<CheckCircle2 class="w-4 h-4" />
							{item.status}
						</span>
						<a href="/eligible" class="text-xs font-extrabold text-primary uppercase tracking-wider hover:underline">
							View Subject Preference List &rarr;
						</a>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>
