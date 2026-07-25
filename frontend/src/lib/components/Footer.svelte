<!-- src/lib/components/Footer.svelte -->
<script lang="ts">
	import { GraduationCap, Heart } from 'lucide-svelte';
</script>

<footer class="bg-slate-900 text-slate-300 border-t border-slate-800 mt-auto">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
		<div class="grid grid-cols-1 md:grid-cols-4 gap-8">
			<!-- Brand -->
			<div class="space-y-4 md:col-span-1">
				<div class="flex items-center gap-3">
					<div class="w-9 h-9 rounded-lg bg-indigo-600 flex items-center justify-center text-white">
						<GraduationCap class="w-5 h-5" />
					</div>
					<span class="text-xl font-extrabold text-white">UniApp</span>
				</div>
				<p class="text-sm text-slate-400 leading-relaxed">
					Unified public university admission portal in Bangladesh. Simplified eligibility tracking & application submissions.
				</p>
			</div>

			<!-- Quick Links -->
			<div>
				<h4 class="text-white font-semibold text-sm mb-4 uppercase tracking-wider">Navigation</h4>
				<ul class="space-y-2.5 text-sm text-slate-400">
					<li><a href="/programs" class="hover:text-white transition-colors">Programs List</a></li>
					<li><a href="/eligible" class="hover:text-white transition-colors">Eligibility Calculator</a></li>
					<li><a href="/routes-finder" class="hover:text-white transition-colors">University Transport Routes</a></li>
				</ul>
			</div>

			<!-- Student Portal -->
			<div>
				<h4 class="text-white font-semibold text-sm mb-4 uppercase tracking-wider">Student Account</h4>
				<ul class="space-y-2.5 text-sm text-slate-400">
					<li><a href="/login" class="hover:text-white transition-colors">Sign In</a></li>
					<li><a href="/register" class="hover:text-white transition-colors">Create Account</a></li>
					<li><a href="/profile" class="hover:text-white transition-colors">Academic Profile</a></li>
				</ul>
			</div>

			<!-- Contact / Support -->
			<div>
				<h4 class="text-white font-semibold text-sm mb-4 uppercase tracking-wider">Helpdesk</h4>
				<p class="text-sm text-slate-400 mb-2">Need assistance with your applications?</p>
				<a href="/helpdesk" class="inline-block px-4 py-2 bg-slate-800 hover:bg-slate-700 text-xs font-semibold text-indigo-400 rounded-lg border border-slate-700 transition-colors">
					Contact Support
				</a>
			</div>
		</div>

		<div class="border-t border-slate-800 mt-10 pt-6 flex flex-col sm:flex-row items-center justify-between text-xs text-slate-500">
			<p>&copy; {new Date().getFullYear()} UniApp. DBMS Project Management System.</p>
			<p class="flex items-center gap-1 mt-2 sm:mt-0">
				Built with <Heart class="w-3.5 h-3.5 text-red-500 fill-red-500" /> using Svelte 5 & Go REST API
			</p>
		</div>
	</div>
</footer>
