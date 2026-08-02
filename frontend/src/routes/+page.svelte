<!-- src/routes/+page.svelte -->
<script lang="ts">
	import { fetchStats, fetchFeatures, fetchTestimonials } from '$lib/api/landing';
	import { fetchPrograms } from '$lib/api/programs';
	import { fetchUniversities } from '$lib/api/university';
	import ProgramCard from '$lib/components/ProgramCard.svelte';
	import UniversityCard from '$lib/components/UniversityCard.svelte';
	import type { Stat, Feature, Testimonial } from '$lib/types/landing';
	import type { Program, University } from '$lib/types/models';
	import { authState } from '$lib/state/auth.svelte';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import {
		Sparkles,
		ArrowRight,
		Search,
		Award,
		MapPin,
		ShieldCheck,
		CheckCircle2,
		BookOpen,
		Building2,
		GraduationCap,
		Users,
		Target,
		Clock,
		Quote,
		Globe,
		ExternalLink
	} from 'lucide-svelte';

	let stats = $state<Stat[]>([]);
	let features = $state<Feature[]>([]);
	let testimonials = $state<Testimonial[]>([]);
	let searchKeyword = $state('');

	// Showcase Tab State
	let activeTab = $state<'programs' | 'universities'>('programs');
	let programsList = $state<Program[]>([]);
	let universitiesList = $state<University[]>([]);
	let isShowcaseLoading = $state(true);

	function getUniLogo(prog: Program): string | null {
		if (prog.logo_url) return prog.logo_url;
		if (prog.university_logo) return prog.university_logo;
		let found = universitiesList.find((u) => u.u_id === prog.u_id);
		if (found?.logo_url) return found.logo_url;
		const nameToMatch = prog.university_name || prog.u_name;
		if (nameToMatch) {
			found = universitiesList.find((u) => u.u_name.toLowerCase() === nameToMatch.toLowerCase() || nameToMatch.toLowerCase().includes(u.u_name.toLowerCase()));
			if (found?.logo_url) return found.logo_url;
		}
		return null;
	}

	onMount(async () => {
		try {
			const [s, f, t, pData, uData] = await Promise.all([
				fetchStats(),
				fetchFeatures(),
				fetchTestimonials(),
				fetchPrograms().catch(() => []),
				fetchUniversities().catch(() => [])
			]);
			stats = s;
			features = f;
			testimonials = t;
			programsList = pData || [];
			universitiesList = uData || [];
		} catch (err) {
			console.error('Error fetching landing data:', err);
		} finally {
			isShowcaseLoading = false;
		}

		// Scroll reveal observer
		const observer = new IntersectionObserver(
			(entries) => {
				entries.forEach((entry) => {
					if (entry.isIntersecting) {
						entry.target.classList.add('active');
					}
				});
			},
			{ threshold: 0.1 }
		);

		document.querySelectorAll('.reveal').forEach((el) => observer.observe(el));

		return () => observer.disconnect();
	});

	function handleSearch(e: Event) {
		e.preventDefault();
		if (searchKeyword.trim()) {
			goto(`/programs?search=${encodeURIComponent(searchKeyword.trim())}`);
		} else {
			goto('/programs');
		}
	}
</script>

<svelte:head>
	<title>UniApp - Unified Public University Admission System</title>
</svelte:head>

<!-- Hero Section with Background Parallax Effect -->
<section class="relative w-full min-h-[85vh] flex flex-col items-center justify-center px-6 overflow-hidden hero-parallax-container py-16">
	<!-- High quality background hero gradient mesh -->
	<div
		class="absolute inset-0 w-full h-full z-0 opacity-20 transition-transform duration-500 ease-out bg-cover bg-center"
		style='background-image: url("https://lh3.googleusercontent.com/aida-public/AB6AXuD1R5Z3nJp6JEaw2UxWiGNNHQKr2GeT2M7g_7wd6MlmdFjSWiOx5YryrDrtsxZr2kdqKC4HNLp9p2cIY7OKeysYiAbrmFs37kJfEs3kGXsugLAiPCGTeMd2-TXv0kxhE2fSxX66VFiU8w4UyRwNNP9tOOyYcMBZ57YNIllJfkS1_DeVtJyymmc1X3E8Z1FgXE_D48Nkb6CGgvG2sHCUIB4Af0mdAcxlLRtDdibQcUO2QGnhtHyeCHJ4FTTZOrC5_ZEzkSlS3wWrRN4o");'
	></div>
	
	<!-- Glow accent -->
	<div class="absolute top-1/3 left-1/2 -translate-x-1/2 w-[600px] h-[600px] bg-gradient-to-tr from-primary/20 via-primary-container/10 to-tertiary-fixed/15 rounded-full blur-3xl pointer-events-none"></div>

	<div class="relative z-10 max-w-4xl mx-auto text-center flex flex-col items-center gap-6 animate-fade-in-up">
		<div class="glass-panel px-5 py-2 rounded-full inline-flex items-center gap-2 shadow-sm border border-outline-variant/40 hover:shadow-md transition-shadow">
			<Sparkles class="w-4 h-4 text-primary" />
			<span class="text-xs font-bold text-primary uppercase tracking-wider">One Profile. Unlimited Opportunities.</span>
		</div>

		<h1 class="text-4xl sm:text-5xl lg:text-6xl font-black text-on-surface leading-tight tracking-tight">
			Your Future Begins <span class="bg-gradient-to-r from-primary via-secondary to-tertiary bg-clip-text text-transparent">Here.</span>
		</h1>

		<p class="text-lg sm:text-xl text-on-surface-variant max-w-2xl leading-relaxed">
			Streamline your Bangladesh public university application process. Explore institutions, track travel routes, and match minimum cutmarks with a single profile.
		</p>

		<div class="flex flex-col sm:flex-row gap-4 mt-2 w-full sm:w-auto">
			<button
				type="button"
				onclick={() => { activeTab = 'programs'; document.getElementById('showcase')?.scrollIntoView({ behavior: 'smooth' }); }}
				class="bg-primary text-white text-base font-bold px-8 py-3.5 rounded-2xl hover:bg-primary-container hover:shadow-xl hover:shadow-primary/30 transition-all duration-300 transform hover:-translate-y-1 flex items-center justify-center gap-2.5 shadow-lg shadow-primary/25 cursor-pointer"
			>
				<Search class="w-5 h-5" />
				<span>Browse Programs</span>
			</button>
			<button
				type="button"
				onclick={() => { activeTab = 'universities'; document.getElementById('showcase')?.scrollIntoView({ behavior: 'smooth' }); }}
				class="glass-panel text-primary text-base font-bold px-8 py-3.5 rounded-2xl hover:bg-surface-container-low transition-all duration-300 transform hover:-translate-y-1 flex items-center justify-center gap-2 border border-outline-variant/40 cursor-pointer"
			>
				<Building2 class="w-5 h-5 text-tertiary" />
				<span>Explore Universities</span>
			</button>
		</div>

		<!-- Hero Search Input Box -->
		<form onsubmit={handleSearch} class="w-full max-w-2xl mt-8 glass-panel rounded-2xl p-2.5 flex items-center gap-2 border border-outline-variant/40 focus-within:ring-2 focus-within:ring-primary/40 transition-all duration-300 hover:shadow-lg">
			<Search class="w-5 h-5 text-outline ml-3 shrink-0" />
			<input
				type="text"
				bind:value={searchKeyword}
				placeholder="Search universities, programs, or units (e.g. BUET, Unit A)..."
				class="w-full bg-transparent border-none text-base text-on-surface placeholder:text-outline p-3 focus:outline-none focus:ring-0"
			/>
			<button type="submit" class="bg-primary text-white font-bold px-6 py-3 rounded-xl hover:bg-primary-container transition-all shrink-0 cursor-pointer">
				Search
			</button>
		</form>
	</div>
</section>

<!-- Mesh Background Wrapper for Stats, Showcase & Features -->
<div class="bg-mesh relative z-10 w-full min-h-screen">
	<!-- Interactive Showcase Section: Active Programs & Explore Universities -->
	<section id="showcase" class="py-12 px-6 max-w-7xl mx-auto reveal">
		<div class="glass-panel p-8 sm:p-12 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-2xl space-y-8">
			<!-- Section Header & Tab Controls -->
			<div class="flex flex-col md:flex-row items-center justify-between gap-6 border-b border-outline-variant/30 pb-6">
				<div class="text-center md:text-left space-y-1">
					<h2 class="text-3xl sm:text-4xl font-black text-on-surface">Explore Admissions</h2>
					<p class="text-on-surface-variant text-sm font-semibold">Switch between active degree programs and top public universities across Bangladesh.</p>
				</div>

				<!-- Tab Switcher -->
				<div class="inline-flex p-1.5 rounded-2xl bg-surface-container-low border border-outline-variant/30 shadow-inner shrink-0">
					<button
						type="button"
						onclick={() => activeTab = 'programs'}
						class="px-6 py-3 rounded-xl font-bold text-sm transition-all duration-300 flex items-center gap-2.5 cursor-pointer {activeTab === 'programs' ? 'bg-primary text-white shadow-lg shadow-primary/25 scale-[1.02]' : 'text-on-surface-variant hover:text-on-surface'}"
					>
						<BookOpen class="w-4.5 h-4.5" />
						<span>Active Programs</span>
						<span class="px-2 py-0.5 rounded-full text-xs font-black {activeTab === 'programs' ? 'bg-white/20 text-white' : 'bg-surface-container text-primary'}">{programsList.length}</span>
					</button>
					<button
						type="button"
						onclick={() => activeTab = 'universities'}
						class="px-6 py-3 rounded-xl font-bold text-sm transition-all duration-300 flex items-center gap-2.5 cursor-pointer {activeTab === 'universities' ? 'bg-primary text-white shadow-lg shadow-primary/25 scale-[1.02]' : 'text-on-surface-variant hover:text-on-surface'}"
					>
						<Building2 class="w-4.5 h-4.5" />
						<span>Explore Universities</span>
						<span class="px-2 py-0.5 rounded-full text-xs font-black {activeTab === 'universities' ? 'bg-white/20 text-white' : 'bg-surface-container text-primary'}">{universitiesList.length}</span>
					</button>
				</div>
			</div>

			<!-- Tab 1: Active Programs View -->
			{#if activeTab === 'programs'}
				{#if isShowcaseLoading}
					<div class="py-12 text-center text-on-surface-variant">
						<div class="w-10 h-10 border-4 border-primary border-t-transparent rounded-full animate-spin mx-auto mb-3"></div>
						<p class="font-bold text-sm">Loading active admission circulars...</p>
					</div>
				{:else if programsList.length === 0}
					<div class="text-center py-12 text-on-surface-variant space-y-2">
						<BookOpen class="w-12 h-12 text-outline mx-auto" />
						<p class="font-extrabold text-base text-on-surface">No active programs available.</p>
					</div>
				{:else}
					<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
						{#each programsList.slice(0, 6) as prog}
							<ProgramCard program={prog} uniLogo={getUniLogo(prog)} />
						{/each}
					</div>

					{#if programsList.length > 6}
						<div class="text-center pt-4">
							<a href="/programs" class="inline-flex items-center gap-2 text-sm font-extrabold text-primary hover:underline">
								View All {programsList.length} Active Programs
								<ArrowRight class="w-4 h-4" />
							</a>
						</div>
					{/if}
				{/if}
			{/if}

			<!-- Tab 2: Explore Universities View -->
			{#if activeTab === 'universities'}
				{#if isShowcaseLoading}
					<div class="py-12 text-center text-on-surface-variant">
						<div class="w-10 h-10 border-4 border-primary border-t-transparent rounded-full animate-spin mx-auto mb-3"></div>
						<p class="font-bold text-sm">Loading public universities...</p>
					</div>
				{:else if universitiesList.length === 0}
					<div class="text-center py-12 text-on-surface-variant space-y-2">
						<Building2 class="w-12 h-12 text-outline mx-auto" />
						<p class="font-extrabold text-base text-on-surface">No universities listed.</p>
					</div>
				{:else}
					<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
						{#each universitiesList.slice(0, 6) as uni}
							<UniversityCard university={uni} />
						{/each}
					</div>

					{#if universitiesList.length > 6}
						<div class="text-center pt-4">
							<a href="/universities" class="inline-flex items-center gap-2 text-sm font-extrabold text-primary hover:underline">
								Explore All {universitiesList.length} Public Universities
								<ArrowRight class="w-4 h-4" />
							</a>
						</div>
					{/if}
				{/if}
			{/if}
		</div>
	</section>

	<!-- Stats Section -->
	<section class="py-12 px-6 md:px-10 max-w-7xl mx-auto reveal">
		<div class="bg-white/90 backdrop-blur-2xl rounded-[2.5rem] p-8 sm:p-12 shadow-2xl shadow-primary/5 border border-outline-variant/30">
			<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8 lg:gap-0 lg:divide-x divide-outline-variant/30">
				<!-- Stat 1 -->
				<div class="flex flex-col items-center text-center px-4 group hover:-translate-y-1 transition-all duration-300">
					<div class="w-12 h-12 rounded-2xl bg-indigo-50 text-indigo-600 border border-indigo-100 flex items-center justify-center mb-4 shadow-sm group-hover:scale-110 group-hover:bg-indigo-600 group-hover:text-white transition-all duration-300">
						<GraduationCap class="w-6 h-6" />
					</div>
					<span class="text-4xl sm:text-5xl font-black text-indigo-600 tracking-tight mb-1">50+</span>
					<span class="text-xs sm:text-sm font-extrabold text-on-surface-variant uppercase tracking-wider">Partner Universities</span>
				</div>

				<!-- Stat 2 -->
				<div class="flex flex-col items-center text-center px-4 group hover:-translate-y-1 transition-all duration-300">
					<div class="w-12 h-12 rounded-2xl bg-violet-50 text-violet-600 border border-violet-100 flex items-center justify-center mb-4 shadow-sm group-hover:scale-110 group-hover:bg-violet-600 group-hover:text-white transition-all duration-300">
						<Users class="w-6 h-6" />
					</div>
					<span class="text-4xl sm:text-5xl font-black text-violet-600 tracking-tight mb-1">10k</span>
					<span class="text-xs sm:text-sm font-extrabold text-on-surface-variant uppercase tracking-wider">Active Students</span>
				</div>

				<!-- Stat 3 -->
				<div class="flex flex-col items-center text-center px-4 group hover:-translate-y-1 transition-all duration-300">
					<div class="w-12 h-12 rounded-2xl bg-emerald-50 text-emerald-600 border border-emerald-100 flex items-center justify-center mb-4 shadow-sm group-hover:scale-110 group-hover:bg-emerald-600 group-hover:text-white transition-all duration-300">
						<Target class="w-6 h-6" />
					</div>
					<span class="text-4xl sm:text-5xl font-black text-emerald-600 tracking-tight mb-1">95%</span>
					<span class="text-xs sm:text-sm font-extrabold text-on-surface-variant uppercase tracking-wider">Match Accuracy</span>
				</div>

				<!-- Stat 4 -->
				<div class="flex flex-col items-center text-center px-4 group hover:-translate-y-1 transition-all duration-300">
					<div class="w-12 h-12 rounded-2xl bg-amber-50 text-amber-600 border border-amber-100 flex items-center justify-center mb-4 shadow-sm group-hover:scale-110 group-hover:bg-amber-600 group-hover:text-white transition-all duration-300">
						<Clock class="w-6 h-6" />
					</div>
					<span class="text-4xl sm:text-5xl font-black text-amber-600 tracking-tight mb-1">24h</span>
					<span class="text-xs sm:text-sm font-extrabold text-on-surface-variant uppercase tracking-wider">Support Response</span>
				</div>
			</div>
		</div>
	</section>

	<!-- Features Section -->
	<section class="py-20 px-6 max-w-7xl mx-auto">
		<div class="mb-14 text-center reveal">
			<h2 class="text-3xl sm:text-4xl font-extrabold text-on-surface mb-3">Powerful Features for Your Journey</h2>
			<p class="text-lg text-on-surface-variant max-w-2xl mx-auto">Everything you need to manage your applications in one seamless platform.</p>
		</div>

		<div class="grid grid-cols-1 md:grid-cols-3 gap-6">
			<!-- Feature 1: Wide Card -->
			<div class="md:col-span-2 glass-panel bg-white/80 backdrop-blur-2xl rounded-[2rem] p-8 sm:p-10 flex flex-col justify-between relative overflow-hidden group hover:shadow-xl hover:shadow-primary/10 hover:-translate-y-1 transition-all duration-300 reveal delay-100 border border-outline-variant/30">
				<div class="absolute -right-16 -top-16 w-72 h-72 bg-primary/10 rounded-full blur-3xl group-hover:bg-primary/20 transition-colors duration-500"></div>
				<div>
					<div class="w-14 h-14 bg-primary-fixed text-primary rounded-2xl flex items-center justify-center mb-6 shadow-md transform group-hover:scale-110 transition-transform duration-300">
						<Award class="w-7 h-7" />
					</div>
					<h3 class="text-2xl font-bold text-on-surface mb-3">Smart Eligibility Engine</h3>
					<p class="text-on-surface-variant text-base leading-relaxed max-w-lg">
						Instantly see which universities you qualify for based on your GPA and HSC subject marks (Physics, Mathematics, Chemistry). Stop guessing and start applying with confidence.
					</p>
				</div>
				<a href="/eligible" class="inline-flex items-center gap-2 text-primary font-extrabold text-sm uppercase tracking-wider hover:gap-3 transition-all pt-6">
					<span>Try Engine</span>
					<ArrowRight class="w-4 h-4" />
				</a>
			</div>

			<!-- Feature 2 -->
			<div class="glass-panel bg-white/80 backdrop-blur-2xl rounded-[2rem] p-8 sm:p-10 flex flex-col justify-between group hover:shadow-xl hover:shadow-primary/10 hover:-translate-y-1 transition-all duration-300 reveal delay-200 border border-outline-variant/30">
				<div>
					<div class="w-14 h-14 bg-tertiary-fixed text-tertiary rounded-2xl flex items-center justify-center mb-6 shadow-md transform group-hover:scale-110 transition-transform duration-300">
						<MapPin class="w-7 h-7" />
					</div>
					<h3 class="text-2xl font-bold text-on-surface mb-3">Route Tracker</h3>
					<p class="text-on-surface-variant text-base leading-relaxed">
						Visualize travel times, transport routes, and shuttle train timings for major university exam centers across Bangladesh.
					</p>
				</div>
				<a href="/routes-finder" class="inline-flex items-center gap-2 text-tertiary font-extrabold text-sm uppercase tracking-wider hover:gap-3 transition-all pt-6">
					<span>View Routes</span>
					<ArrowRight class="w-4 h-4" />
				</a>
			</div>

			<!-- Feature 3 -->
			<div class="glass-panel bg-white/80 backdrop-blur-2xl rounded-[2rem] p-8 sm:p-10 flex flex-col justify-between group hover:shadow-xl hover:shadow-primary/10 hover:-translate-y-1 transition-all duration-300 reveal delay-100 border border-outline-variant/30">
				<div>
					<div class="w-14 h-14 bg-secondary-fixed text-secondary rounded-2xl flex items-center justify-center mb-6 shadow-md transform group-hover:scale-110 transition-transform duration-300">
						<Sparkles class="w-7 h-7" />
					</div>
					<h3 class="text-2xl font-bold text-on-surface mb-3">Live Alerts</h3>
					<p class="text-on-surface-variant text-base leading-relaxed">
						Never miss a deadline or exam date with synchronized notifications for Unit A, B, and C admissions.
					</p>
				</div>
				<a href="/dashboard" class="inline-flex items-center gap-2 text-secondary font-extrabold text-sm uppercase tracking-wider hover:gap-3 transition-all pt-6">
					<span>View Alerts</span>
					<ArrowRight class="w-4 h-4" />
				</a>
			</div>

			<!-- Feature 4: Wide Card -->
			<div class="md:col-span-2 glass-panel bg-white/80 backdrop-blur-2xl rounded-[2rem] p-8 sm:p-10 flex flex-col justify-between relative overflow-hidden group hover:shadow-xl hover:shadow-primary/10 hover:-translate-y-1 transition-all duration-300 reveal delay-300 border border-outline-variant/30">
				<div>
					<div class="w-14 h-14 bg-error-container text-error rounded-2xl flex items-center justify-center mb-6 shadow-md transform group-hover:scale-110 transition-transform duration-300">
						<ShieldCheck class="w-7 h-7" />
					</div>
					<h3 class="text-2xl font-bold text-on-surface mb-3">Unified Applications</h3>
					<p class="text-on-surface-variant text-base leading-relaxed max-w-lg">
						Fill out your academic profile once. Apply to multiple programs with automatic verification of missing required profile parameters.
					</p>
				</div>
				<a href="/profile" class="inline-flex items-center gap-2 text-primary font-extrabold text-sm uppercase tracking-wider hover:gap-3 transition-all pt-6">
					<span>Manage Profile</span>
					<ArrowRight class="w-4 h-4" />
				</a>
			</div>
		</div>
	</section>

	<!-- Testimonials Section -->
	<section class="py-20 px-6 max-w-7xl mx-auto reveal">
		<div class="mb-14 text-center">
			<h2 class="text-3xl sm:text-4xl font-extrabold text-on-surface mb-3">Loved by Applicants Nationwide</h2>
			<p class="text-lg text-on-surface-variant max-w-2xl mx-auto">See how UniApp helped thousands of Bangladeshi students secure university admission.</p>
		</div>

		<div class="grid grid-cols-1 md:grid-cols-3 gap-8">
			<!-- Testimonial 1 -->
			<div class="bg-white/90 backdrop-blur-2xl rounded-[2.5rem] p-8 sm:p-10 border border-outline-variant/30 shadow-lg shadow-primary/5 hover:shadow-2xl hover:shadow-primary/10 hover:-translate-y-1.5 transition-all duration-300 flex flex-col justify-between space-y-6 group">
				<div class="space-y-4">
					<div class="w-12 h-12 rounded-2xl bg-indigo-50 text-indigo-500 border border-indigo-100 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
						<Quote class="w-6 h-6 rotate-180" />
					</div>
					<p class="text-on-surface/90 text-base leading-relaxed font-normal">
						"UniApp made finding eligible universities incredibly easy. The interface is clean, and the recommendations were spot on. Highly recommended for stressed students!"
					</p>
				</div>
				<div class="flex items-center gap-4 pt-6 border-t border-outline-variant/20">
					<img
						src="https://images.unsplash.com/photo-1539571696357-5a69c17a67c6?w=150&auto=format&fit=crop&q=80"
						alt="Rahim Uddin"
						class="w-12 h-12 rounded-full object-cover border-2 border-primary/20 shadow-sm"
					/>
					<div>
						<h4 class="font-bold text-on-surface text-base">Rahim Uddin</h4>
						<p class="text-xs font-semibold text-primary">Accepted to BUET</p>
					</div>
				</div>
			</div>

			<!-- Testimonial 2 -->
			<div class="bg-white/90 backdrop-blur-2xl rounded-[2.5rem] p-8 sm:p-10 border border-outline-variant/30 shadow-lg shadow-primary/5 hover:shadow-2xl hover:shadow-primary/10 hover:-translate-y-1.5 transition-all duration-300 flex flex-col justify-between space-y-6 group">
				<div class="space-y-4">
					<div class="w-12 h-12 rounded-2xl bg-violet-50 text-violet-500 border border-violet-100 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
						<Quote class="w-6 h-6 rotate-180" />
					</div>
					<p class="text-on-surface/90 text-base leading-relaxed font-normal">
						"The unified application feature saved me weeks of repetitive typing. I could focus on my exams instead of filling out the same forms twenty times."
					</p>
				</div>
				<div class="flex items-center gap-4 pt-6 border-t border-outline-variant/20">
					<img
						src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=150&auto=format&fit=crop&q=80"
						alt="Sarah Ahmed"
						class="w-12 h-12 rounded-full object-cover border-2 border-primary/20 shadow-sm"
					/>
					<div>
						<h4 class="font-bold text-on-surface text-base">Sarah Ahmed</h4>
						<p class="text-xs font-semibold text-violet-600">Accepted to DU</p>
					</div>
				</div>
			</div>

			<!-- Testimonial 3 -->
			<div class="bg-white/90 backdrop-blur-2xl rounded-[2.5rem] p-8 sm:p-10 border border-outline-variant/30 shadow-lg shadow-primary/5 hover:shadow-2xl hover:shadow-primary/10 hover:-translate-y-1.5 transition-all duration-300 flex flex-col justify-between space-y-6 group">
				<div class="space-y-4">
					<div class="w-12 h-12 rounded-2xl bg-emerald-50 text-emerald-500 border border-emerald-100 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
						<Quote class="w-6 h-6 rotate-180" />
					</div>
					<p class="text-on-surface/90 text-base leading-relaxed font-normal">
						"Tracking deadlines across different varsity routes used to be a nightmare. The calendar and alerts on this platform are lifesavers."
					</p>
				</div>
				<div class="flex items-center gap-4 pt-6 border-t border-outline-variant/20">
					<img
						src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=150&auto=format&fit=crop&q=80"
						alt="Fahad Hossain"
						class="w-12 h-12 rounded-full object-cover border-2 border-primary/20 shadow-sm"
					/>
					<div>
						<h4 class="font-bold text-on-surface text-base">Fahad Hossain</h4>
						<p class="text-xs font-semibold text-emerald-600">Accepted to RUET</p>
					</div>
				</div>
			</div>
		</div>
	</section>
</div>
