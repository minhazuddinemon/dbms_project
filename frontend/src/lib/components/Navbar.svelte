<!-- src/lib/components/Navbar.svelte -->
<script lang="ts">
	import { authState } from '$lib/state/auth.svelte';
	import { GraduationCap, LogIn, UserPlus, LogOut, User, Compass, Award, MapPin, Search } from 'lucide-svelte';

	let mobileMenuOpen = $state(false);

	function toggleMenu() {
		mobileMenuOpen = !mobileMenuOpen;
	}
</script>

<nav class="sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-slate-200/80 shadow-sm transition-all duration-300">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
		<div class="flex items-center justify-between h-16">
			<!-- Logo -->
			<a href="/" class="flex items-center gap-3 group">
				<div class="w-10 h-10 rounded-xl bg-gradient-to-tr from-indigo-600 via-indigo-500 to-emerald-400 flex items-center justify-center text-white shadow-md shadow-indigo-500/20 group-hover:scale-105 transition-transform">
					<GraduationCap class="w-6 h-6" />
				</div>
				<div>
					<span class="text-xl font-extrabold bg-gradient-to-r from-indigo-900 via-indigo-700 to-indigo-900 bg-clip-text text-transparent">UniApp</span>
					<span class="block text-[10px] uppercase tracking-widest text-slate-400 font-semibold -mt-1">Admission Portal</span>
				</div>
			</a>

			<!-- Desktop Nav Links -->
			<div class="hidden md:flex items-center gap-1 lg:gap-2">
				<a href="/programs" class="px-3 py-2 rounded-lg text-sm font-medium text-slate-700 hover:text-indigo-600 hover:bg-slate-50 transition-colors flex items-center gap-1.5">
					<Search class="w-4 h-4 text-slate-400" />
					Programs
				</a>
				<a href="/eligible" class="px-3 py-2 rounded-lg text-sm font-medium text-slate-700 hover:text-indigo-600 hover:bg-slate-50 transition-colors flex items-center gap-1.5">
					<Award class="w-4 h-4 text-emerald-500" />
					Eligible Varsity
				</a>
				<a href="/routes-finder" class="px-3 py-2 rounded-lg text-sm font-medium text-slate-700 hover:text-indigo-600 hover:bg-slate-50 transition-colors flex items-center gap-1.5">
					<MapPin class="w-4 h-4 text-slate-400" />
					Route Tracker
				</a>
			</div>

			<!-- User Auth Actions -->
			<div class="hidden md:flex items-center gap-3">
				{#if authState.isAuthenticated}
					<a href="/profile" class="px-4 py-2 rounded-lg text-sm font-semibold text-slate-700 hover:bg-slate-100 transition-all flex items-center gap-2 border border-slate-200">
						<User class="w-4 h-4 text-indigo-600" />
						Profile & Marks
					</a>
					<button
						onclick={() => authState.logout()}
						class="px-4 py-2 rounded-lg text-sm font-semibold text-slate-600 hover:text-red-600 hover:bg-red-50 transition-all flex items-center gap-1.5"
					>
						<LogOut class="w-4 h-4" />
						Logout
					</button>
				{:else}
					<a href="/login" class="px-4 py-2 rounded-lg text-sm font-semibold text-slate-700 hover:text-indigo-600 hover:bg-slate-50 transition-all flex items-center gap-1.5">
						<LogIn class="w-4 h-4 text-slate-400" />
						Sign In
					</a>
					<a href="/register" class="px-4 py-2 rounded-lg text-sm font-semibold text-white bg-indigo-600 hover:bg-indigo-700 shadow-md shadow-indigo-600/20 hover:shadow-indigo-600/30 transition-all flex items-center gap-1.5">
						<UserPlus class="w-4 h-4" />
						Register
					</a>
				{/if}
			</div>

			<!-- Mobile Menu Button -->
			<div class="md:hidden flex items-center">
				<button
					onclick={toggleMenu}
					class="p-2 rounded-lg text-slate-600 hover:text-slate-900 hover:bg-slate-100 transition-colors"
					aria-label="Toggle menu"
				>
					<Compass class="w-6 h-6" />
				</button>
			</div>
		</div>
	</div>

	<!-- Mobile Dropdown -->
	{#if mobileMenuOpen}
		<div class="md:hidden border-t border-slate-200 bg-white px-4 pt-2 pb-4 space-y-2">
			<a href="/programs" onclick={() => mobileMenuOpen = false} class="block px-3 py-2 rounded-lg text-base font-medium text-slate-700 hover:bg-slate-50">Explore Programs</a>
			<a href="/eligible" onclick={() => mobileMenuOpen = false} class="block px-3 py-2 rounded-lg text-base font-medium text-slate-700 hover:bg-slate-50">Check Eligibility</a>
			<a href="/routes-finder" onclick={() => mobileMenuOpen = false} class="block px-3 py-2 rounded-lg text-base font-medium text-slate-700 hover:bg-slate-50">Route Finder</a>
			<div class="pt-3 border-t border-slate-100 flex flex-col gap-2">
				{#if authState.isAuthenticated}
					<a href="/profile" onclick={() => mobileMenuOpen = false} class="w-full text-center px-4 py-2 rounded-lg font-semibold bg-slate-100 text-slate-800">My Profile</a>
					<button onclick={() => { authState.logout(); mobileMenuOpen = false; }} class="w-full text-center px-4 py-2 rounded-lg font-semibold bg-red-50 text-red-600">Logout</button>
				{:else}
					<a href="/login" onclick={() => mobileMenuOpen = false} class="w-full text-center px-4 py-2 rounded-lg font-semibold border border-slate-300 text-slate-700">Sign In</a>
					<a href="/register" onclick={() => mobileMenuOpen = false} class="w-full text-center px-4 py-2 rounded-lg font-semibold bg-indigo-600 text-white">Register Account</a>
				{/if}
			</div>
		</div>
	{/if}
</nav>
