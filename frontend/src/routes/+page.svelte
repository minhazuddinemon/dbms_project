<!-- src/routes/+page.svelte -->
<script lang="ts">
	import { ArrowRight, CheckCircle2, Search, Award, Sparkles, Building2, ShieldCheck, Clock } from 'lucide-svelte';
	import { authState } from '$lib/state/auth.svelte';
</script>

<!-- Hero Section -->
<section class="relative overflow-hidden bg-gradient-to-b from-indigo-900 via-indigo-950 to-slate-950 text-white py-20 lg:py-28">
	<!-- Background Mesh & Glows -->
	<div class="absolute inset-0 bg-[radial-gradient(ellipse_at_top_right,_var(--tw-gradient-stops))] from-indigo-500/20 via-transparent to-transparent pointer-events-none"></div>
	<div class="absolute top-1/4 left-1/2 -translate-x-1/2 w-96 h-96 bg-emerald-500/10 rounded-full blur-3xl pointer-events-none"></div>

	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative z-10">
		<div class="max-w-3xl mx-auto text-center space-y-6">
			<div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-indigo-500/10 border border-indigo-400/20 text-indigo-300 text-xs font-semibold uppercase tracking-wider backdrop-blur-md">
				<Sparkles class="w-4 h-4 text-emerald-400" />
				Unified University Admission System
			</div>

			<h1 class="text-4xl sm:text-5xl lg:text-6xl font-extrabold tracking-tight text-white leading-tight">
				Find Your Dream <span class="bg-gradient-to-r from-emerald-400 via-teal-300 to-indigo-300 bg-clip-text text-transparent">University & Program</span>
			</h1>

			<p class="text-lg sm:text-xl text-slate-300 max-w-2xl mx-auto leading-relaxed">
				Check your academic eligibility instantly, apply to public university programs, track minimum cutmarks, and manage admission forms in one place.
			</p>

			<div class="pt-4 flex flex-col sm:flex-row items-center justify-center gap-4">
				<a
					href="/eligible"
					class="w-full sm:w-auto px-8 py-3.5 rounded-xl font-bold bg-gradient-to-r from-emerald-500 to-teal-500 text-white shadow-lg shadow-emerald-500/25 hover:shadow-emerald-500/40 hover:scale-105 transition-all flex items-center justify-center gap-2"
				>
					<Award class="w-5 h-5" />
					Check My Eligibility
					<ArrowRight class="w-4 h-4" />
				</a>

				<a
					href="/programs"
					class="w-full sm:w-auto px-8 py-3.5 rounded-xl font-bold bg-white/10 hover:bg-white/15 text-white border border-white/20 backdrop-blur-md transition-all flex items-center justify-center gap-2"
				>
					<Search class="w-5 h-5 text-indigo-300" />
					Browse Programs
				</a>
			</div>
		</div>

		<!-- Quick Stats Row -->
		<div class="mt-16 grid grid-cols-2 md:grid-cols-4 gap-4 max-w-4xl mx-auto">
			<div class="bg-white/5 border border-white/10 rounded-2xl p-5 text-center backdrop-blur-md">
				<span class="block text-3xl font-extrabold text-white">50+</span>
				<span class="text-xs font-medium text-slate-400 uppercase tracking-wider mt-1 block">Universities</span>
			</div>
			<div class="bg-white/5 border border-white/10 rounded-2xl p-5 text-center backdrop-blur-md">
				<span class="block text-3xl font-extrabold text-emerald-400">100%</span>
				<span class="text-xs font-medium text-slate-400 uppercase tracking-wider mt-1 block">Rule Accuracy</span>
			</div>
			<div class="bg-white/5 border border-white/10 rounded-2xl p-5 text-center backdrop-blur-md">
				<span class="block text-3xl font-extrabold text-indigo-300">Live</span>
				<span class="text-xs font-medium text-slate-400 uppercase tracking-wider mt-1 block">Cutmarks Tracker</span>
			</div>
			<div class="bg-white/5 border border-white/10 rounded-2xl p-5 text-center backdrop-blur-md">
				<span class="block text-3xl font-extrabold text-white">Instant</span>
				<span class="text-xs font-medium text-slate-400 uppercase tracking-wider mt-1 block">Application Form</span>
			</div>
		</div>
	</div>
</section>

<!-- Main Features Grid -->
<section class="py-20 bg-slate-50">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
		<div class="text-center max-w-2xl mx-auto mb-16">
			<h2 class="text-3xl font-extrabold text-slate-900 sm:text-4xl">Smart Admission Tools</h2>
			<p class="mt-3 text-lg text-slate-600">Everything you need to navigate public university admissions efficiently.</p>
		</div>

		<div class="grid grid-cols-1 md:grid-cols-3 gap-8">
			<!-- Feature 1 -->
			<div class="bg-white rounded-2xl p-8 border border-slate-200 shadow-sm hover:shadow-md transition-all space-y-4">
				<div class="w-12 h-12 rounded-xl bg-indigo-100 text-indigo-600 flex items-center justify-center">
					<Award class="w-6 h-6" />
				</div>
				<h3 class="text-xl font-bold text-slate-900">Eligibility Engine</h3>
				<p class="text-slate-600 text-sm leading-relaxed">
					Automatically verifies your SSC/HSC GPA and individual subject marks (Physics, Math, Chemistry) against program rules.
				</p>
				<a href="/eligible" class="inline-flex items-center text-sm font-semibold text-indigo-600 hover:text-indigo-700 gap-1 pt-2">
					Try Eligibility Engine &rarr;
				</a>
			</div>

			<!-- Feature 2 -->
			<div class="bg-white rounded-2xl p-8 border border-slate-200 shadow-sm hover:shadow-md transition-all space-y-4">
				<div class="w-12 h-12 rounded-xl bg-emerald-100 text-emerald-600 flex items-center justify-center">
					<Building2 class="w-6 h-6" />
				</div>
				<h3 class="text-xl font-bold text-slate-900">Program Search</h3>
				<p class="text-slate-600 text-sm leading-relaxed">
					Filter programs by Unit A, B, or C, search by keyword, view total seat count, cutmarks history, and application deadlines.
				</p>
				<a href="/programs" class="inline-flex items-center text-sm font-semibold text-indigo-600 hover:text-indigo-700 gap-1 pt-2">
					Search Programs &rarr;
				</a>
			</div>

			<!-- Feature 3 -->
			<div class="bg-white rounded-2xl p-8 border border-slate-200 shadow-sm hover:shadow-md transition-all space-y-4">
				<div class="w-12 h-12 rounded-xl bg-teal-100 text-teal-600 flex items-center justify-center">
					<ShieldCheck class="w-6 h-6" />
				</div>
				<h3 class="text-xl font-bold text-slate-900">One-Click Application</h3>
				<p class="text-slate-600 text-sm leading-relaxed">
					Apply directly to programs. The system automatically notifies you of any missing required profile parameters.
				</p>
				<a href={authState.isAuthenticated ? "/profile" : "/register"} class="inline-flex items-center text-sm font-semibold text-indigo-600 hover:text-indigo-700 gap-1 pt-2">
					{authState.isAuthenticated ? "Manage Profile" : "Register Account"} &rarr;
				</a>
			</div>
		</div>
	</div>
</section>
