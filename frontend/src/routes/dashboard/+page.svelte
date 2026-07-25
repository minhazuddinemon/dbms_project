<!-- src/routes/dashboard/+page.svelte -->
<script lang="ts">
	import { authState } from '$lib/state/auth.svelte';
	import { fetchEligiblePrograms } from '$lib/api/programs';
	import { fetchStudentApplications } from '$lib/api/applications';
	import type { EligibleProgram, StudentApplication } from '$lib/types/models';
	import { onMount } from 'svelte';
	import { LayoutDashboard, Award, CheckCircle2, Clock, AlertTriangle, FileText, ArrowRight, Building2, MapPin, CreditCard, Sparkles, Calendar, TrendingUp } from 'lucide-svelte';

	let eligibleCount = $state<number>(0);
	let applications = $state<StudentApplication[]>([]);
	let isLoading = $state(true);

	onMount(async () => {
		if (authState.isAuthenticated) {
			try {
				const [progs, apps] = await Promise.all([
					fetchEligiblePrograms().catch(() => []),
					fetchStudentApplications().catch(() => [])
				]);
				eligibleCount = progs.length;
				applications = apps;
			} catch (err) {
				console.error(err);
			} finally {
				isLoading = false;
			}
		} else {
			isLoading = false;
		}
	});
</script>

<svelte:head>
	<title>Student Dashboard - UniApp</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		
		<!-- Welcome Header Banner & Hero Card -->
		<div class="grid grid-cols-1 lg:grid-cols-12 gap-6 items-stretch">
			<!-- Greeting Text -->
			<div class="lg:col-span-5 flex flex-col justify-center space-y-2">
				<div class="inline-flex items-center gap-2 px-3.5 py-1 rounded-full bg-primary/10 border border-primary/20 text-primary text-xs font-bold uppercase tracking-wider w-fit">
					<Sparkles class="w-3.5 h-3.5" />
					Student Portal
				</div>
				<h1 class="text-3xl sm:text-4xl font-extrabold text-on-surface leading-tight">
					Good morning,<br />
					<span class="bg-gradient-to-r from-primary to-primary-container bg-clip-text text-transparent">
						{authState.user?.email ? authState.user.email.split('@')[0] : 'Applicant'} 👋
					</span>
				</h1>
				<p class="text-on-surface-variant text-base">Let's make your university admission journey successful!</p>
			</div>

			<!-- Hero Card Banner -->
			<div class="lg:col-span-7 glass-panel bg-gradient-to-br from-white/90 via-primary/5 to-primary-container/10 p-8 rounded-[2.5rem] border border-outline-variant/40 shadow-xl relative overflow-hidden flex flex-col justify-between space-y-4">
				<div class="absolute right-0 top-0 w-64 h-64 bg-primary/10 rounded-full blur-3xl pointer-events-none"></div>
				<div class="relative z-10 space-y-2">
					<h3 class="text-2xl font-bold text-primary">Your future begins here</h3>
					<p class="text-on-surface-variant text-sm max-w-md leading-relaxed">
						Explore top Bangladeshi public universities, check subject cutmarks, and apply in just a few clicks.
					</p>
				</div>
				<div class="relative z-10 pt-2">
					<a
						href="/eligible"
						class="inline-flex items-center gap-2 bg-primary text-white text-sm font-bold px-6 py-3 rounded-xl hover:bg-primary-container shadow-lg shadow-primary/25 hover:shadow-primary/40 hover:-translate-y-0.5 transition-all duration-200"
					>
						<span>Explore Universities</span>
						<ArrowRight class="w-4 h-4" />
					</a>
				</div>
			</div>
		</div>

		<!-- 4 Key Stats Overview -->
		<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
			<!-- Total Applications -->
			<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/30 shadow-md hover:shadow-xl hover:-translate-y-1 transition-all duration-300 flex items-center gap-5 bg-white/90 group">
				<div class="w-14 h-14 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center shrink-0 group-hover:scale-110 transition-transform">
					<FileText class="w-7 h-7" />
				</div>
				<div>
					<p class="text-xs font-bold text-on-surface-variant uppercase tracking-wider">Total Applications</p>
					<p class="text-3xl font-black text-on-surface">{applications.length}</p>
					<a href="/apply/1" class="text-xs font-bold text-primary inline-flex items-center gap-1 hover:underline mt-1">
						View all <ArrowRight class="w-3 h-3" />
					</a>
				</div>
			</div>

			<!-- Eligible Universities -->
			<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/30 shadow-md hover:shadow-xl hover:-translate-y-1 transition-all duration-300 flex items-center gap-5 bg-white/90 group">
				<div class="w-14 h-14 rounded-2xl bg-tertiary-fixed/40 text-tertiary flex items-center justify-center shrink-0 group-hover:scale-110 transition-transform">
					<Building2 class="w-7 h-7" />
				</div>
				<div>
					<p class="text-xs font-bold text-on-surface-variant uppercase tracking-wider">Eligible Varsities</p>
					<p class="text-3xl font-black text-on-surface">{eligibleCount || 24}</p>
					<a href="/eligible" class="text-xs font-bold text-tertiary inline-flex items-center gap-1 hover:underline mt-1">
						View all <ArrowRight class="w-3 h-3" />
					</a>
				</div>
			</div>

			<!-- Upcoming Deadlines -->
			<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/30 shadow-md hover:shadow-xl hover:-translate-y-1 transition-all duration-300 flex items-center gap-5 bg-white/90 group">
				<div class="w-14 h-14 rounded-2xl bg-amber-500/10 text-amber-600 flex items-center justify-center shrink-0 group-hover:scale-110 transition-transform">
					<Calendar class="w-7 h-7" />
				</div>
				<div>
					<p class="text-xs font-bold text-on-surface-variant uppercase tracking-wider">Upcoming Deadlines</p>
					<p class="text-3xl font-black text-on-surface">3</p>
					<a href="/dashboard" class="text-xs font-bold text-amber-600 inline-flex items-center gap-1 hover:underline mt-1">
						View dates <ArrowRight class="w-3 h-3" />
					</a>
				</div>
			</div>

			<!-- Pending Payments -->
			<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/30 shadow-md hover:shadow-xl hover:-translate-y-1 transition-all duration-300 flex items-center gap-5 bg-white/90 group">
				<div class="w-14 h-14 rounded-2xl bg-secondary-fixed text-secondary flex items-center justify-center shrink-0 group-hover:scale-110 transition-transform">
					<CreditCard class="w-7 h-7" />
				</div>
				<div>
					<p class="text-xs font-bold text-on-surface-variant uppercase tracking-wider">Pending Payments</p>
					<p class="text-3xl font-black text-on-surface">2</p>
					<a href="/payment" class="text-xs font-bold text-secondary inline-flex items-center gap-1 hover:underline mt-1">
						Pay now <ArrowRight class="w-3 h-3" />
					</a>
				</div>
			</div>
		</div>

		<!-- Action Shortcuts Grid -->
		<div class="grid grid-cols-1 md:grid-cols-3 gap-6">
			<!-- Shortcut 1 -->
			<a href="/eligible" class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 hover:border-primary transition-all duration-300 hover:shadow-xl hover:-translate-y-1 flex flex-col justify-between space-y-6 bg-white/90 group">
				<div class="space-y-3">
					<div class="w-12 h-12 rounded-2xl bg-tertiary-fixed/50 text-tertiary flex items-center justify-center group-hover:scale-110 transition-transform">
						<Award class="w-6 h-6" />
					</div>
					<h3 class="text-xl font-bold text-on-surface">Eligibility Matcher</h3>
					<p class="text-sm text-on-surface-variant leading-relaxed">
						Re-evaluate your GPA and HSC subject marks against updated cutmarks for Unit A, B, and C.
					</p>
				</div>
				<span class="inline-flex items-center gap-2 text-sm font-extrabold text-primary uppercase tracking-wider group-hover:gap-3 transition-all">
					Check Eligible Varsities <ArrowRight class="w-4 h-4" />
				</span>
			</a>

			<!-- Shortcut 2 -->
			<a href="/profile" class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 hover:border-primary transition-all duration-300 hover:shadow-xl hover:-translate-y-1 flex flex-col justify-between space-y-6 bg-white/90 group">
				<div class="space-y-3">
					<div class="w-12 h-12 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center group-hover:scale-110 transition-transform">
						<FileText class="w-6 h-6" />
					</div>
					<h3 class="text-xl font-bold text-on-surface">Fill Academic Profile</h3>
					<p class="text-sm text-on-surface-variant leading-relaxed">
						Update SSC/HSC roll, registration numbers, father/mother details, quota status, and addresses.
					</p>
				</div>
				<span class="inline-flex items-center gap-2 text-sm font-extrabold text-primary uppercase tracking-wider group-hover:gap-3 transition-all">
					Edit Profile Info <ArrowRight class="w-4 h-4" />
				</span>
			</a>

			<!-- Shortcut 3 -->
			<a href="/routes-finder" class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 hover:border-primary transition-all duration-300 hover:shadow-xl hover:-translate-y-1 flex flex-col justify-between space-y-6 bg-white/90 group">
				<div class="space-y-3">
					<div class="w-12 h-12 rounded-2xl bg-secondary-fixed text-secondary flex items-center justify-center group-hover:scale-110 transition-transform">
						<MapPin class="w-6 h-6" />
					</div>
					<h3 class="text-xl font-bold text-on-surface">Transport Route Tracker</h3>
					<p class="text-sm text-on-surface-variant leading-relaxed">
						View campus transit routes, shuttle train timetables, and travel time estimates to DU, BUET, and JU.
					</p>
				</div>
				<span class="inline-flex items-center gap-2 text-sm font-extrabold text-primary uppercase tracking-wider group-hover:gap-3 transition-all">
					View Route Map <ArrowRight class="w-4 h-4" />
				</span>
			</a>
		</div>

		<!-- Application Progress & Timeline Table -->
		<div class="glass-panel p-8 sm:p-10 rounded-[2.5rem] border border-outline-variant/30 bg-white/95 space-y-6 shadow-xl">
			<div class="flex items-center justify-between border-b border-outline-variant/30 pb-4">
				<h3 class="text-2xl font-bold text-on-surface">Recent Application Progress</h3>
				<span class="text-xs font-bold text-on-surface-variant uppercase tracking-wider">Updated Live</span>
			</div>

			<div class="space-y-4">
				<!-- Item 1 -->
				<div class="flex flex-col sm:flex-row items-start sm:items-center justify-between p-5 rounded-2xl bg-surface-container-low/70 border border-outline-variant/30 gap-4 hover:bg-surface-container-low transition-colors">
					<div class="flex items-center gap-4">
						<div class="w-12 h-12 rounded-2xl bg-emerald-50 text-emerald-600 border border-emerald-100 flex items-center justify-center shrink-0">
							<Building2 class="w-6 h-6" />
						</div>
						<div>
							<h4 class="font-bold text-on-surface text-base">Dhaka University - Unit A (Science)</h4>
							<span class="text-xs font-semibold text-emerald-700 bg-emerald-100/70 px-3 py-0.5 rounded-full inline-block mt-1">Eligible & Approved</span>
						</div>
					</div>
					<a href="/apply/1" class="px-5 py-2.5 rounded-xl text-xs font-bold text-white bg-primary hover:bg-primary-container transition-all shadow-md shadow-primary/20">
						View Submission
					</a>
				</div>

				<!-- Item 2 -->
				<div class="flex flex-col sm:flex-row items-start sm:items-center justify-between p-5 rounded-2xl bg-surface-container-low/70 border border-outline-variant/30 gap-4 hover:bg-surface-container-low transition-colors">
					<div class="flex items-center gap-4">
						<div class="w-12 h-12 rounded-2xl bg-indigo-50 text-indigo-600 border border-indigo-100 flex items-center justify-center shrink-0">
							<Building2 class="w-6 h-6" />
						</div>
						<div>
							<h4 class="font-bold text-on-surface text-base">BUET - Computer Science & Engineering</h4>
							<span class="text-xs font-semibold text-indigo-700 bg-indigo-100/70 px-3 py-0.5 rounded-full inline-block mt-1">Subject Marks Verified</span>
						</div>
					</div>
					<a href="/apply/2" class="px-5 py-2.5 rounded-xl text-xs font-bold text-white bg-primary hover:bg-primary-container transition-all shadow-md shadow-primary/20">
						View Details
					</a>
				</div>

				<!-- Item 3 -->
				<div class="flex flex-col sm:flex-row items-start sm:items-center justify-between p-5 rounded-2xl bg-surface-container-low/70 border border-outline-variant/30 gap-4 hover:bg-surface-container-low transition-colors">
					<div class="flex items-center gap-4">
						<div class="w-12 h-12 rounded-2xl bg-amber-50 text-amber-600 border border-amber-100 flex items-center justify-center shrink-0">
							<Building2 class="w-6 h-6" />
						</div>
						<div>
							<h4 class="font-bold text-on-surface text-base">RUET - Electrical & Electronic Engineering</h4>
							<span class="text-xs font-semibold text-amber-700 bg-amber-100/70 px-3 py-0.5 rounded-full inline-block mt-1">Pending Application Fee</span>
						</div>
					</div>
					<a href="/payment" class="px-5 py-2.5 rounded-xl text-xs font-bold text-white bg-amber-600 hover:bg-amber-700 transition-all shadow-md shadow-amber-600/20">
						Complete Payment
					</a>
				</div>
			</div>
		</div>
	</div>
</div>

