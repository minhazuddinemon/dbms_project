<!-- src/lib/components/Navbar.svelte -->
<script lang="ts">
	import { authState } from '$lib/state/auth.svelte';
	import { fetchStudentNotifications } from '$lib/api/student';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { onMount } from 'svelte';
	import {
		GraduationCap,
		Building2,
		LogIn,
		UserPlus,
		LogOut,
		User,
		Search,
		Award,
		LayoutDashboard,
		Menu,
		X,
		Shield,
		Bus
	} from 'lucide-svelte';

	let mobileMenuOpen = $state(false);
	let notificationCount = $state<number>(0);

	async function loadNotifications() {
		if (authState.isAuthenticated && !authState.isAdmin) {
			try {
				const notifs = await fetchStudentNotifications();
				notificationCount = notifs ? notifs.length : 0;
			} catch (err) {
				notificationCount = 0;
			}
		} else {
			notificationCount = 0;
		}
	}

	$effect(() => {
		if (authState.isAuthenticated && !authState.isAdmin) {
			loadNotifications();
		}
	});

	onMount(() => {
		loadNotifications();
	});

	function toggleMenu() {
		mobileMenuOpen = !mobileMenuOpen;
	}

	function handleLogout() {
		const isAdminPage = page.url.pathname.startsWith('/admin');
		authState.logout();
		if (isAdminPage) {
			goto('/admin');
		} else {
			goto('/login');
		}
	}
</script>

<nav class="sticky top-0 z-50 bg-white/85 backdrop-blur-xl border-b border-outline-variant/30 shadow-xs transition-all duration-300">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
		<div class="flex items-center justify-between h-20">
			<!-- Logo -->
			<a href="/" class="flex items-center gap-3.5 group">
				<div class="w-11 h-11 rounded-2xl bg-linear-to-tr from-primary via-primary-container to-tertiary-fixed-dim flex items-center justify-center text-white shadow-lg shadow-primary/25 group-hover:scale-105 group-hover:rotate-3 transition-all duration-300">
					<GraduationCap class="w-6 h-6" />
				</div>
				<div>
					<span class="text-2xl font-black bg-linear-to-r from-on-surface via-primary to-primary-container bg-clip-text text-transparent tracking-tight">UniApp</span>
					<span class="block text-[11px] font-bold uppercase tracking-widest text-outline -mt-1">Admission Portal</span>
				</div>
			</a>

			<!-- Desktop Nav Links -->
			<div class="hidden lg:flex items-center gap-1 xl:gap-2 bg-surface-container-low/70 p-1.5 rounded-2xl border border-outline-variant/30">
				<!-- 1. Universities (Public) -->
				<a href="/universities" class="px-4 py-2 rounded-xl text-sm font-semibold text-on-surface-variant hover:text-primary hover:bg-white transition-all duration-200 flex items-center gap-2">
					<Building2 class="w-4 h-4 text-primary" />
					Universities
				</a>

				<!-- 2. Programs (Public) -->
				<a href="/programs" class="px-4 py-2 rounded-xl text-sm font-semibold text-on-surface-variant hover:text-primary hover:bg-white transition-all duration-200 flex items-center gap-2">
					<Search class="w-4 h-4 text-outline" />
					Programs
				</a>

				<!-- 3. Route Tracker (Public for Everyone) -->
				<a href="/routes-finder" class="px-4 py-2 rounded-xl text-sm font-semibold text-on-surface-variant hover:text-primary hover:bg-white transition-all duration-200 flex items-center gap-2">
					<Bus class="w-4 h-4 text-primary" />
					Route Tracker
				</a>

				<!-- 4. Eligible Uni (Only if Logged In as student) -->
				{#if authState.isAuthenticated && !authState.isAdmin}
					<a href="/eligible" class="px-4 py-2 rounded-xl text-sm font-semibold text-on-surface-variant hover:text-tertiary hover:bg-white transition-all duration-200 flex items-center gap-2">
						<Award class="w-4 h-4 text-tertiary" />
						Eligible Uni
					</a>

					<!-- 5. Dashboard (Only if Logged In as student) -->
					<a href="/dashboard" class="relative px-4 py-2 rounded-xl text-sm font-semibold text-on-surface-variant hover:text-primary hover:bg-white transition-all duration-200 flex items-center gap-2">
						<LayoutDashboard class="w-4 h-4 text-primary" />
						<span>Dashboard</span>
						{#if notificationCount > 0}
							<span class="inline-flex items-center justify-center h-5 min-w-5 px-1.5 text-[10px] font-black text-white rounded-full bg-red-600 animate-blink-red shadow-sm ml-0.5">
								{notificationCount}
							</span>
						{/if}
					</a>
				{/if}

				<!-- Admin Panel Link if Admin Logged In -->
				{#if authState.isAdmin}
					<a href="/admin" class="px-4 py-2 rounded-xl text-sm font-semibold text-primary hover:bg-white transition-all duration-200 flex items-center gap-2">
						<Shield class="w-4 h-4 text-primary" />
						Admin Dashboard
					</a>
				{/if}
			</div>

			<!-- User Actions -->
			<div class="hidden lg:flex items-center gap-3">
				{#if authState.isAdmin}
					<!-- Admin Logged In: Only Admin Badge & Logout -->
					<a href="/admin" class="px-5 py-2.5 rounded-xl text-sm font-bold text-white bg-primary hover:bg-primary-container shadow-md transition-all duration-200 flex items-center gap-2">
						<Shield class="w-4 h-4" />
						Admin Panel
					</a>
					<button
						onclick={handleLogout}
						class="px-4 py-2.5 rounded-xl text-sm font-bold text-error hover:bg-error-container/40 transition-all duration-200 flex items-center gap-1.5 cursor-pointer"
					>
						<LogOut class="w-4 h-4" />
						Logout
					</button>
				{:else if authState.isAuthenticated}
					<!-- Student Logged In: Profile & Logout -->
					<a href="/profile" class="px-5 py-2.5 rounded-xl text-sm font-bold text-on-surface bg-surface-container hover:bg-surface-container-high transition-all duration-200 flex items-center gap-2 border border-outline-variant/40">
						<User class="w-4 h-4 text-primary" />
						My Profile
					</a>
					<button
						onclick={handleLogout}
						class="px-4 py-2.5 rounded-xl text-sm font-bold text-error hover:bg-error-container/40 transition-all duration-200 flex items-center gap-1.5 cursor-pointer"
					>
						<LogOut class="w-4 h-4" />
						Logout
					</button>
				{:else}
					<!-- Not Logged In: Show Sign In & Register -->
					<a href="/login" class="px-5 py-2.5 rounded-xl text-sm font-bold text-on-surface-variant hover:text-primary hover:bg-surface-container-low transition-all duration-200 flex items-center gap-2">
						<LogIn class="w-4 h-4 text-outline" />
						Sign In
					</a>
					<a href="/register" class="px-6 py-2.5 rounded-xl text-sm font-bold text-white bg-primary hover:bg-primary-container shadow-lg shadow-primary/25 hover:shadow-primary/40 hover:-translate-y-0.5 transition-all duration-200 flex items-center gap-2">
						<UserPlus class="w-4 h-4" />
						Register
					</a>
				{/if}
			</div>

			<!-- Mobile Menu Button -->
			<div class="lg:hidden flex items-center">
				<button
					onclick={toggleMenu}
					class="p-2.5 rounded-xl text-on-surface-variant hover:text-on-surface hover:bg-surface-container transition-colors"
					aria-label="Toggle navigation menu"
				>
					{#if mobileMenuOpen}
						<X class="w-6 h-6" />
					{:else}
						<Menu class="w-6 h-6" />
					{/if}
				</button>
			</div>
		</div>
	</div>

	<!-- Mobile Menu Drawer -->
	{#if mobileMenuOpen}
		<div class="lg:hidden border-t border-outline-variant/30 bg-white/95 backdrop-blur-2xl px-6 pt-4 pb-6 space-y-3 animate-fade-in-up">
			<a href="/universities" onclick={() => mobileMenuOpen = false} class="flex items-center gap-3 px-4 py-3 rounded-xl font-semibold text-on-surface hover:bg-surface-container">
				<Building2 class="w-5 h-5 text-primary" /> Universities
			</a>
			<a href="/programs" onclick={() => mobileMenuOpen = false} class="flex items-center gap-3 px-4 py-3 rounded-xl font-semibold text-on-surface hover:bg-surface-container">
				<Search class="w-5 h-5 text-outline" /> Browse Programs
			</a>
			<a href="/routes-finder" onclick={() => mobileMenuOpen = false} class="flex items-center gap-3 px-4 py-3 rounded-xl font-semibold text-on-surface hover:bg-surface-container">
				<Bus class="w-5 h-5 text-primary" /> Route Tracker
			</a>
			{#if authState.isAdmin}
				<a href="/admin" onclick={() => mobileMenuOpen = false} class="flex items-center gap-3 px-4 py-3 rounded-xl font-semibold text-primary hover:bg-surface-container">
					<Shield class="w-5 h-5 text-primary" /> Admin Panel
				</a>
			{:else if authState.isAuthenticated}
				<a href="/eligible" onclick={() => mobileMenuOpen = false} class="flex items-center gap-3 px-4 py-3 rounded-xl font-semibold text-on-surface hover:bg-surface-container">
					<Award class="w-5 h-5 text-tertiary" /> Eligible Uni
				</a>
				<a href="/dashboard" onclick={() => mobileMenuOpen = false} class="flex items-center justify-between px-4 py-3 rounded-xl font-semibold text-on-surface hover:bg-surface-container">
					<div class="flex items-center gap-3">
						<LayoutDashboard class="w-5 h-5 text-primary" /> Dashboard
					</div>
					{#if notificationCount > 0}
						<span class="inline-flex items-center justify-center px-2 py-0.5 text-xs font-black text-white rounded-full bg-red-600 animate-blink-red">
							{notificationCount}
						</span>
					{/if}
				</a>
			{/if}

			<div class="pt-4 border-t border-outline-variant/30 flex flex-col gap-2">
				{#if authState.isAdmin}
					<a href="/admin" onclick={() => mobileMenuOpen = false} class="w-full text-center px-4 py-3 rounded-xl font-bold bg-primary text-white shadow-md">Admin Panel</a>
					<button onclick={() => { handleLogout(); mobileMenuOpen = false; }} class="w-full text-center px-4 py-3 rounded-xl font-bold bg-error-container text-on-error-container cursor-pointer">Logout</button>
				{:else if authState.isAuthenticated}
					<a href="/profile" onclick={() => mobileMenuOpen = false} class="w-full text-center px-4 py-3 rounded-xl font-bold bg-primary-fixed text-on-primary-fixed">My Profile</a>
					<button onclick={() => { handleLogout(); mobileMenuOpen = false; }} class="w-full text-center px-4 py-3 rounded-xl font-bold bg-error-container text-on-error-container cursor-pointer">Logout</button>
				{:else}
					<a href="/login" onclick={() => mobileMenuOpen = false} class="w-full text-center px-4 py-3 rounded-xl font-bold border border-outline-variant text-on-surface">Sign In</a>
					<a href="/register" onclick={() => mobileMenuOpen = false} class="w-full text-center px-4 py-3 rounded-xl font-bold bg-primary text-white shadow-lg shadow-primary/20">Register Account</a>
				{/if}
			</div>
		</div>
	{/if}
</nav>
