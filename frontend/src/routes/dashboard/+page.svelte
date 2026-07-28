<!-- src/routes/dashboard/+page.svelte -->
<script lang="ts">
	import { authState } from '$lib/state/auth.svelte';
	import { fetchEligiblePrograms } from '$lib/api/programs';
	import { fetchStudentApplications } from '$lib/api/applications';
	import { fetchStudentNotifications } from '$lib/api/student';
	import type { EligibleProgram, StudentApplication, StudentNotification } from '$lib/types/models';
	import { onMount } from 'svelte';
	import {
		LayoutDashboard,
		Award,
		CheckCircle2,
		Clock,
		AlertTriangle,
		FileText,
		ArrowRight,
		Building2,
		MapPin,
		CreditCard,
		Sparkles,
		Calendar,
		Bell
	} from 'lucide-svelte';

	let eligibleCount = $state<number>(0);
	let applications = $state<StudentApplication[]>([]);
	let notifications = $state<StudentNotification[]>([]);
	let isLoading = $state(true);

	onMount(async () => {
		if (authState.isAuthenticated) {
			try {
				const [progs, apps, notifs] = await Promise.all([
					fetchEligiblePrograms().catch(() => []),
					fetchStudentApplications().catch(() => []),
					fetchStudentNotifications().catch(() => [])
				]);
				eligibleCount = progs.length;
				applications = apps;
				notifications = notifs;
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
					Welcome back,<br />
					<span class="bg-gradient-to-r from-primary to-primary-container bg-clip-text text-transparent">
						{authState.user?.email ? authState.user.email.split('@')[0] : 'Applicant'}
					</span>
				</h1>
				<p class="text-on-surface-variant text-base">Track your admission applications and notifications in real-time.</p>
			</div>

			<!-- Hero Card Banner -->
			<div class="lg:col-span-7 glass-panel bg-gradient-to-br from-white/90 via-primary/5 to-primary-container/10 p-8 rounded-[2.5rem] border border-outline-variant/40 shadow-xl relative overflow-hidden flex flex-col justify-between space-y-4">
				<div class="absolute right-0 top-0 w-64 h-64 bg-primary/10 rounded-full blur-3xl pointer-events-none"></div>
				<div class="relative z-10 space-y-2">
					<h3 class="text-2xl font-bold text-primary">Explore Eligible Varsities</h3>
					<p class="text-on-surface-variant text-sm max-w-md leading-relaxed">
						View public universities matching your SSC/HSC GPA and individual subject cutmarks.
					</p>
				</div>
				<div class="relative z-10 pt-2">
					<a
						href="/eligible"
						class="inline-flex items-center gap-2 bg-primary text-white text-sm font-bold px-6 py-3 rounded-xl hover:bg-primary-container shadow-lg shadow-primary/25 transition-all duration-200"
					>
						<span>Explore Eligible Programs</span>
						<ArrowRight class="w-4 h-4" />
					</a>
				</div>
			</div>
		</div>

		<!-- Notifications Alert Widget -->
		{#if notifications.length > 0}
			<div class="glass-panel p-6 rounded-[2.5rem] border border-sky-200 bg-sky-50/70 shadow-md space-y-3">
				<div class="flex items-center gap-2 text-sky-900 font-extrabold text-base border-b border-sky-200/60 pb-2">
					<Bell class="w-5 h-5 text-sky-600" />
					System Notifications & Announcements ({notifications.length})
				</div>
				<div class="space-y-2">
					{#each notifications as n}
						<div class="p-3.5 rounded-xl bg-white border border-sky-100 shadow-xs flex items-start gap-3">
							<Sparkles class="w-4 h-4 text-sky-500 shrink-0 mt-0.5" />
							<div class="flex-1">
								<p class="text-xs font-semibold text-slate-800 leading-relaxed">{n.message}</p>
								<span class="text-[10px] text-slate-400 font-mono block mt-1">{new Date(n.created_at).toLocaleString()}</span>
							</div>
						</div>
					{/each}
				</div>
			</div>
		{/if}

		<!-- 4 Key Stats Overview -->
		<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
			<!-- Total Applications -->
			<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/30 shadow-md hover:shadow-xl transition-all duration-300 flex items-center gap-5 bg-white/90 group">
				<div class="w-14 h-14 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center shrink-0 group-hover:scale-110 transition-transform">
					<FileText class="w-7 h-7" />
				</div>
				<div>
					<p class="text-xs font-bold text-on-surface-variant uppercase tracking-wider">Total Applications</p>
					<p class="text-3xl font-black text-on-surface">{applications.length}</p>
				</div>
			</div>

			<!-- Eligible Universities -->
			<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/30 shadow-md hover:shadow-xl transition-all duration-300 flex items-center gap-5 bg-white/90 group">
				<div class="w-14 h-14 rounded-2xl bg-tertiary-fixed/40 text-tertiary flex items-center justify-center shrink-0 group-hover:scale-110 transition-transform">
					<Building2 class="w-7 h-7" />
				</div>
				<div>
					<p class="text-xs font-bold text-on-surface-variant uppercase tracking-wider">Eligible Varsities</p>
					<p class="text-3xl font-black text-on-surface">{eligibleCount}</p>
					<a href="/eligible" class="text-xs font-bold text-tertiary inline-flex items-center gap-1 hover:underline mt-1">
						View all <ArrowRight class="w-3 h-3" />
					</a>
				</div>
			</div>

			<!-- System Notifications -->
			<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/30 shadow-md hover:shadow-xl transition-all duration-300 flex items-center gap-5 bg-white/90 group">
				<div class="w-14 h-14 rounded-2xl bg-sky-500/10 text-sky-600 flex items-center justify-center shrink-0 group-hover:scale-110 transition-transform">
					<Bell class="w-7 h-7" />
				</div>
				<div>
					<p class="text-xs font-bold text-on-surface-variant uppercase tracking-wider">Live Alerts</p>
					<p class="text-3xl font-black text-on-surface">{notifications.length}</p>
				</div>
			</div>

			<!-- Payments -->
			<div class="glass-panel p-6 rounded-[2rem] border border-outline-variant/30 shadow-md hover:shadow-xl transition-all duration-300 flex items-center gap-5 bg-white/90 group">
				<div class="w-14 h-14 rounded-2xl bg-secondary-fixed text-secondary flex items-center justify-center shrink-0 group-hover:scale-110 transition-transform">
					<CreditCard class="w-7 h-7" />
				</div>
				<div>
					<p class="text-xs font-bold text-on-surface-variant uppercase tracking-wider">Payments</p>
					<a href="/payment" class="text-xs font-bold text-secondary inline-flex items-center gap-1 hover:underline mt-2">
						Process Payment <ArrowRight class="w-3 h-3" />
					</a>
				</div>
			</div>
		</div>

		<!-- Action Shortcuts Grid -->
		<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
			<!-- Shortcut 1 -->
			<a href="/eligible" class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 hover:border-primary transition-all duration-300 hover:shadow-xl flex flex-col justify-between space-y-6 bg-white/90 group">
				<div class="space-y-3">
					<div class="w-12 h-12 rounded-2xl bg-tertiary-fixed/50 text-tertiary flex items-center justify-center group-hover:scale-110 transition-transform">
						<Award class="w-6 h-6" />
					</div>
					<h3 class="text-xl font-bold text-on-surface">Eligibility Engine</h3>
					<p class="text-sm text-on-surface-variant leading-relaxed">
						Re-evaluate your GPA and HSC subject marks against updated cutmarks for Unit A, B, and C.
					</p>
				</div>
				<span class="inline-flex items-center gap-2 text-sm font-extrabold text-primary uppercase tracking-wider group-hover:gap-3 transition-all">
					Check Eligible Varsities <ArrowRight class="w-4 h-4" />
				</span>
			</a>

			<!-- Shortcut 2 -->
			<a href="/profile" class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 hover:border-primary transition-all duration-300 hover:shadow-xl flex flex-col justify-between space-y-6 bg-white/90 group">
				<div class="space-y-3">
					<div class="w-12 h-12 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center group-hover:scale-110 transition-transform">
						<FileText class="w-6 h-6" />
					</div>
					<h3 class="text-xl font-bold text-on-surface">Fill Academic Profile</h3>
					<p class="text-sm text-on-surface-variant leading-relaxed">
						Update SSC/HSC roll, registration numbers, father/mother details, quota status, and contact mobile numbers.
					</p>
				</div>
				<span class="inline-flex items-center gap-2 text-sm font-extrabold text-primary uppercase tracking-wider group-hover:gap-3 transition-all">
					Edit Profile Info <ArrowRight class="w-4 h-4" />
				</span>
			</a>
		</div>

		<!-- Application Progress & Timeline Table -->
		<div class="glass-panel p-8 sm:p-10 rounded-[2.5rem] border border-outline-variant/30 bg-white/95 space-y-6 shadow-xl">
			<div class="flex items-center justify-between border-b border-outline-variant/30 pb-4">
				<h3 class="text-2xl font-bold text-on-surface">Submitted Applications ({applications.length})</h3>
				<span class="text-xs font-bold text-on-surface-variant uppercase tracking-wider">Live Status</span>
			</div>

			<div class="space-y-4">
				{#if applications.length === 0}
					<p class="text-center py-8 text-on-surface-variant text-sm font-semibold">You have not submitted any applications yet.</p>
				{:else}
					{#each applications as app}
						<div class="flex flex-col sm:flex-row items-start sm:items-center justify-between p-5 rounded-2xl bg-surface-container-low/70 border border-outline-variant/30 gap-4 hover:bg-surface-container-low transition-colors">
							<div class="flex items-center gap-4">
								<div class="w-12 h-12 rounded-2xl bg-primary-fixed/40 text-primary border border-primary/20 flex items-center justify-center shrink-0">
									<Building2 class="w-6 h-6" />
								</div>
								<div>
									<h4 class="font-bold text-on-surface text-base">{app.program_name || 'Program #' + app.program_id}</h4>
									<p class="text-xs text-on-surface-variant">{app.university_name || 'Public University'}</p>
									<span class="text-xs font-semibold px-3 py-0.5 rounded-full inline-block mt-1.5 {app.status === 'PAID' || app.status === 'APPROVED' ? 'bg-emerald-100 text-emerald-700' : 'bg-amber-100 text-amber-700'}">
										Status: {app.status}
									</span>
								</div>
							</div>
							{#if app.status !== 'PAID' && app.status !== 'APPROVED'}
								<a href="/payment" class="px-5 py-2.5 rounded-xl text-xs font-bold text-white bg-amber-600 hover:bg-amber-700 transition-all shadow-md shadow-amber-600/20">
									Pay Application Fee
								</a>
							{/if}
						</div>
					{/each}
				{/if}
			</div>
		</div>
	</div>
</div>
