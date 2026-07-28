<!-- src/routes/merit/+page.svelte -->
<script lang="ts">
	import { fetchStudentNotifications } from '$lib/api/student';
	import { fetchStudentApplications } from '$lib/api/applications';
	import { authState } from '$lib/state/auth.svelte';
	import type { StudentNotification, StudentApplication } from '$lib/types/models';
	import { Award, CheckCircle2, AlertCircle, Building2, TrendingUp, Medal, Sparkles, Bell, UserCheck, ArrowRight, Calendar, CreditCard } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let notifications = $state<StudentNotification[]>([]);
	let applications = $state<StudentApplication[]>([]);
	let isLoading = $state(true);
	let error = $state<string | null>(null);

	async function loadData() {
		if (!authState.isAuthenticated) {
			isLoading = false;
			return;
		}
		isLoading = true;
		error = null;
		try {
			const [notifs, apps] = await Promise.all([
				fetchStudentNotifications().catch(() => []),
				fetchStudentApplications().catch(() => [])
			]);
			notifications = notifs;
			applications = apps;
		} catch (err: any) {
			error = err?.message || 'Failed to load results data.';
		} finally {
			isLoading = false;
		}
	}

	onMount(() => {
		loadData();
	});

	function formatDate(dateStr: string): string {
		try {
			return new Date(dateStr).toLocaleString('en-BD', {
				dateStyle: 'medium',
				timeStyle: 'short'
			});
		} catch {
			return dateStr;
		}
	}

	function getStatusClass(status: string): string {
		switch (status?.toUpperCase()) {
			case 'PAID':
			case 'APPROVED':
				return 'bg-emerald-100 text-emerald-700 border-emerald-200';
			case 'REJECTED':
				return 'bg-red-100 text-red-700 border-red-200';
			default:
				return 'bg-amber-100 text-amber-700 border-amber-200';
		}
	}
</script>

<svelte:head>
	<title>Admission Results & Notifications - UniApp</title>
	<meta name="description" content="View your admission test results, merit position, and system notifications from UniApp." />
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">

		<!-- Banner Header -->
		<div class="bg-gradient-to-r from-secondary via-primary to-primary-container text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>
			<div class="absolute bottom-0 left-1/3 w-64 h-64 bg-secondary-fixed/10 rounded-full blur-2xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-3">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<Medal class="w-4 h-4" />
					Admission Results Center
				</div>

				<h1 class="text-3xl sm:text-4xl font-black text-white">
					Results & Notifications
				</h1>

				<p class="text-slate-100 text-base leading-relaxed">
					Track your admission test results, merit positions, application statuses, and system notifications in one place.
				</p>
			</div>
		</div>

		{#if !authState.isAuthenticated}
			<!-- Not Logged In -->
			<div class="glass-panel rounded-[2.5rem] border border-outline-variant/30 p-10 text-center max-w-md mx-auto space-y-5 bg-white/90 shadow-xl">
				<div class="w-16 h-16 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center mx-auto">
					<UserCheck class="w-8 h-8" />
				</div>
				<h3 class="text-2xl font-extrabold text-on-surface">Sign In to View Results</h3>
				<p class="text-on-surface-variant text-sm leading-relaxed">Log in to check your admission test results, merit positions, and application status updates.</p>
				<div class="pt-2 flex gap-3 justify-center">
					<a href="/login" class="px-6 py-3 rounded-xl font-bold bg-primary text-white text-sm hover:bg-primary-container shadow-lg shadow-primary/20 transition-all">Sign In</a>
					<a href="/register" class="px-6 py-3 rounded-xl font-bold border border-outline-variant text-on-surface text-sm hover:bg-surface-container transition-all">Register</a>
				</div>
			</div>
		{:else if isLoading}
			<div class="py-20 text-center text-on-surface-variant">
				<div class="w-12 h-12 border-4 border-primary border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
				<p class="font-bold text-lg">Loading your results and notifications...</p>
			</div>
		{:else if error}
			<div class="glass-panel rounded-[2.5rem] border border-error-container p-10 text-center max-w-lg mx-auto space-y-4 bg-white/95 shadow-xl">
				<AlertCircle class="w-14 h-14 text-error mx-auto" />
				<h3 class="text-xl font-bold text-on-surface">Error Loading Data</h3>
				<p class="text-on-surface-variant text-sm">{error}</p>
			</div>
		{:else}
			<div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
				<!-- Main Content -->
				<div class="lg:col-span-2 space-y-8">

					<!-- Notifications Section -->
					<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
						<div class="flex items-center justify-between border-b border-outline-variant/30 pb-4">
							<h2 class="text-xl font-extrabold text-on-surface flex items-center gap-2">
								<Bell class="w-6 h-6 text-primary" />
								System Notifications
							</h2>
							<span class="px-3 py-1 rounded-full bg-primary-fixed text-on-primary-fixed text-xs font-bold">
								{notifications.length} Total
							</span>
						</div>

						{#if notifications.length === 0}
							<div class="py-10 text-center space-y-3">
								<div class="w-14 h-14 rounded-2xl bg-surface-container-low text-on-surface-variant flex items-center justify-center mx-auto">
									<Bell class="w-7 h-7" />
								</div>
								<p class="font-bold text-on-surface-variant">No notifications yet.</p>
								<p class="text-xs text-on-surface-variant">You will receive notifications here when admission test results are published.</p>
							</div>
						{:else}
							<div class="space-y-4">
								{#each notifications as notif}
									<div class="p-5 rounded-2xl border border-sky-200 bg-sky-50/60 space-y-2 hover:bg-sky-50 transition-colors">
										<div class="flex items-start gap-3">
											<div class="w-9 h-9 rounded-xl bg-sky-500/15 text-sky-600 flex items-center justify-center shrink-0 mt-0.5">
												<Sparkles class="w-4 h-4" />
											</div>
											<div class="flex-1">
												<p class="text-sm font-semibold text-slate-800 leading-relaxed">{notif.message}</p>
												<span class="text-[10px] text-slate-400 font-mono block mt-1.5">
													{formatDate(notif.created_at)}
												</span>
											</div>
										</div>
									</div>
								{/each}
							</div>
						{/if}
					</div>

					<!-- Applications & Status Section -->
					<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
						<div class="flex items-center justify-between border-b border-outline-variant/30 pb-4">
							<h2 class="text-xl font-extrabold text-on-surface flex items-center gap-2">
								<Award class="w-6 h-6 text-tertiary" />
								Application Status
							</h2>
							<span class="px-3 py-1 rounded-full bg-tertiary-fixed/50 text-tertiary text-xs font-bold">
								{applications.length} Applied
							</span>
						</div>

						{#if applications.length === 0}
							<div class="py-10 text-center space-y-4">
								<div class="w-14 h-14 rounded-2xl bg-surface-container-low text-on-surface-variant flex items-center justify-center mx-auto">
									<Building2 class="w-7 h-7" />
								</div>
								<p class="font-bold text-on-surface-variant">No applications submitted yet.</p>
								<a href="/eligible" class="inline-flex items-center gap-2 px-6 py-2.5 rounded-xl text-sm font-bold bg-primary text-white hover:bg-primary-container shadow-lg shadow-primary/25 transition-all">
									Find Eligible Programs
									<ArrowRight class="w-4 h-4" />
								</a>
							</div>
						{:else}
							<div class="space-y-4">
								{#each applications as app}
									<div class="flex flex-col sm:flex-row items-start sm:items-center justify-between p-5 rounded-2xl border border-outline-variant/30 bg-surface-container-low/50 hover:bg-surface-container-low gap-4 transition-colors">
										<div class="flex items-center gap-4">
											<div class="w-12 h-12 rounded-2xl bg-primary-fixed/40 text-primary border border-primary/20 flex items-center justify-center shrink-0">
												<Building2 class="w-6 h-6" />
											</div>
											<div>
												<h4 class="font-bold text-on-surface text-base">{app.program_name || 'Program #' + app.program_id}</h4>
												<p class="text-xs text-on-surface-variant">{app.university_name || 'Public University'}</p>
												{#if app.sub_date}
													<p class="text-xs text-on-surface-variant flex items-center gap-1 mt-0.5">
														<Calendar class="w-3 h-3" />
														{formatDate(app.sub_date)}
													</p>
												{/if}
											</div>
										</div>
										<div class="flex items-center gap-3 shrink-0">
											<span class="px-3 py-1 rounded-full text-xs font-bold border {getStatusClass(app.status)}">
												{app.status}
											</span>
											{#if app.status !== 'PAID' && app.status !== 'APPROVED'}
												<a href="/payment" class="px-4 py-2 rounded-xl text-xs font-bold text-white bg-amber-600 hover:bg-amber-700 transition-colors flex items-center gap-1 shadow-sm">
													<CreditCard class="w-3.5 h-3.5" />
													Pay Fee
												</a>
											{:else}
												<span class="flex items-center gap-1 text-xs font-bold text-emerald-700">
													<CheckCircle2 class="w-4 h-4" />
													Complete
												</span>
											{/if}
										</div>
									</div>
								{/each}
							</div>
						{/if}
					</div>
				</div>

				<!-- Sidebar -->
				<div class="space-y-6">
					<!-- Quick Stats -->
					<div class="glass-panel p-6 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-4">
						<h3 class="text-sm font-extrabold text-on-surface border-b border-outline-variant/30 pb-3">Overview</h3>
						<div class="grid grid-cols-2 gap-4">
							<div class="p-4 rounded-2xl bg-primary-fixed/30 text-center space-y-1">
								<p class="text-2xl font-black text-primary">{applications.length}</p>
								<p class="text-xs font-bold text-on-surface-variant">Applied</p>
							</div>
							<div class="p-4 rounded-2xl bg-tertiary-fixed/30 text-center space-y-1">
								<p class="text-2xl font-black text-tertiary">{applications.filter(a => a.status === 'PAID' || a.status === 'APPROVED').length}</p>
								<p class="text-xs font-bold text-on-surface-variant">Paid</p>
							</div>
							<div class="p-4 rounded-2xl bg-sky-100 text-center space-y-1">
								<p class="text-2xl font-black text-sky-600">{notifications.length}</p>
								<p class="text-xs font-bold text-on-surface-variant">Alerts</p>
							</div>
							<div class="p-4 rounded-2xl bg-amber-100 text-center space-y-1">
								<p class="text-2xl font-black text-amber-600">{applications.filter(a => a.status !== 'PAID' && a.status !== 'APPROVED').length}</p>
								<p class="text-xs font-bold text-on-surface-variant">Pending</p>
							</div>
						</div>
					</div>

					<!-- Tips Card -->
					<div class="glass-panel p-6 rounded-[2.5rem] border border-outline-variant/40 bg-white/90 shadow-md space-y-4">
						<h3 class="text-sm font-extrabold text-on-surface flex items-center gap-2">
							<TrendingUp class="w-4 h-4 text-primary" />
							Next Steps
						</h3>
						<div class="space-y-3 text-xs text-on-surface-variant">
							<div class="flex items-start gap-2">
								<CheckCircle2 class="w-3.5 h-3.5 text-emerald-500 shrink-0 mt-0.5" />
								<span>Check your notifications for published admission test results.</span>
							</div>
							<div class="flex items-start gap-2">
								<CheckCircle2 class="w-3.5 h-3.5 text-emerald-500 shrink-0 mt-0.5" />
								<span>Pay pending application fees to complete your submission.</span>
							</div>
							<div class="flex items-start gap-2">
								<CheckCircle2 class="w-3.5 h-3.5 text-emerald-500 shrink-0 mt-0.5" />
								<span>Keep your academic profile updated to maximize eligible matches.</span>
							</div>
						</div>
						<a href="/dashboard" class="block w-full py-2.5 px-4 rounded-xl text-center text-xs font-bold text-white bg-primary hover:bg-primary-container shadow-md transition-all">
							Go to Dashboard
						</a>
					</div>
				</div>
			</div>
		{/if}
	</div>
</div>
