<!-- src/lib/components/Footer.svelte -->
<script lang="ts">
	import { GraduationCap, Heart, ArrowUpRight, ShieldCheck, Mail, Phone } from 'lucide-svelte';
</script>

<footer class="bg-on-surface text-surface-container-lowest border-t border-outline/20 mt-auto relative overflow-hidden">
	<!-- Glow accent -->
	<div class="absolute bottom-0 right-0 w-96 h-96 bg-primary/10 rounded-full blur-3xl pointer-events-none"></div>

	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16 relative z-10">
		<div class="grid grid-cols-1 md:grid-cols-4 gap-10">
			<!-- Brand Info -->
			<div class="space-y-5 md:col-span-1">
				<div class="flex items-center gap-3">
					<div class="w-10 h-10 rounded-2xl bg-linear-to-tr from-primary to-tertiary-fixed flex items-center justify-center text-white shadow-lg shadow-primary/20">
						<GraduationCap class="w-6 h-6" />
					</div>
					<span class="text-2xl font-black tracking-tight text-white">UniApp</span>
				</div>
				<p class="text-sm text-outline-variant leading-relaxed">
					Unified public university admission portal in Bangladesh. Engineered for clarity, route tracking, and automated eligibility matching.
				</p>
				<div class="flex items-center gap-2 text-xs font-semibold text-tertiary-fixed">
					<ShieldCheck class="w-4 h-4" />
					<span>Verified Academic Rule Engine</span>
				</div>
			</div>

			<!-- Navigation -->
			<div>
				<h4 class="text-white font-extrabold text-sm mb-5 uppercase tracking-wider">Navigation</h4>
				<ul class="space-y-3 text-sm text-surface-variant">
					<li>
						<a href="/programs" class="hover:text-tertiary-fixed transition-colors flex items-center gap-1 group">
							<span>Browse Programs</span>
							<ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition-opacity" />
						</a>
					</li>
					<li>
						<a href="/eligible" class="hover:text-tertiary-fixed transition-colors flex items-center gap-1 group">
							<span>Eligibility Calculator</span>
							<ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition-opacity" />
						</a>
					</li>
					<li>
						<a href="/routes-finder" class="hover:text-tertiary-fixed transition-colors flex items-center gap-1 group">
							<span>Transport Route Tracker</span>
							<ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition-opacity" />
						</a>
					</li>
					<li>
						<a href="/dashboard" class="hover:text-tertiary-fixed transition-colors flex items-center gap-1 group">
							<span>Student Dashboard</span>
							<ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition-opacity" />
						</a>
					</li>
				</ul>
			</div>

			<!-- Account & Services -->
			<div>
				<h4 class="text-white font-extrabold text-sm mb-5 uppercase tracking-wider">Student Portal</h4>
				<ul class="space-y-3 text-sm text-surface-variant">
					<li><a href="/login" class="hover:text-tertiary-fixed transition-colors">Sign In</a></li>
					<li><a href="/register" class="hover:text-tertiary-fixed transition-colors">Register Account</a></li>
					<li><a href="/profile" class="hover:text-tertiary-fixed transition-colors">Academic Profile & Marks</a></li>
					<li><a href="/helpdesk" class="hover:text-tertiary-fixed transition-colors">Helpdesk Ticket Support</a></li>
				</ul>
			</div>

			<!-- Helpdesk Box -->
			<div>
				<h4 class="text-white font-extrabold text-sm mb-5 uppercase tracking-wider">Support & Help</h4>
				<p class="text-sm text-surface-variant mb-4 leading-relaxed">
					Need help with minimum cutmarks, eligibility rules, or application fee payment?
				</p>
				<a href="/helpdesk" class="inline-flex items-center gap-2 px-5 py-2.5 bg-surface-container/10 hover:bg-surface-container/20 text-xs font-bold text-tertiary-fixed rounded-xl border border-tertiary-fixed/30 transition-all">
					<Mail class="w-4 h-4" />
					Open Support Ticket
				</a>
			</div>
		</div>

		<div class="border-t border-outline/20 mt-12 pt-8 flex flex-col sm:flex-row items-center justify-between text-xs text-outline">
			<p>&copy; {new Date().getFullYear()} UniApp Portal. Database Management System Project.</p>
			<p class="flex items-center gap-1.5 mt-3 sm:mt-0 font-medium">
				Built with Svelte 5 + Tailwind CSS v4 + Go REST API
			</p>
		</div>
	</div>
</footer>
