<!-- src/routes/account/+page.svelte -->
<script lang="ts">
	import { User, Lock, Bell, Shield, Key, Save, CheckCircle2, Phone, Mail, Calendar, FolderShared, FileCheck, Upload, AlertTriangle, Smartphone } from 'lucide-svelte';
	import { authState } from '$lib/state/auth.svelte';

	let firstName = $state('Fahad');
	let lastName = $state('Hossain');
	let email = $state('fahad.hossain@example.com');
	let phone = $state('+880 1712 345678');
	let dob = $state('2002-05-15');

	let currentPassword = $state('');
	let newPassword = $state('');
	let enable2FA = $state(false);

	let appUpdates = $state(true);
	let paymentReceipts = $state(true);
	let promoOffers = $state(false);

	let isSaved = $state(false);

	function handleSave(e: Event) {
		e.preventDefault();
		isSaved = true;
		setTimeout(() => isSaved = false, 2500);
	}
</script>

<svelte:head>
	<title>Profile & Settings - UniApp</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		
		<!-- Banner Header -->
		<div class="flex flex-col md:flex-row justify-between items-start md:items-end gap-6 bg-gradient-to-r from-primary via-primary-container to-secondary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-3">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<Shield class="w-4 h-4" />
					Security & Preferences
				</div>

				<h1 class="text-3xl sm:text-4xl font-black text-white">
					Profile & Settings
				</h1>

				<p class="text-slate-100 text-base leading-relaxed">
					Manage your personal information, security options, uploaded documents, and communication preferences.
				</p>
			</div>

			<button
				onclick={handleSave}
				class="relative z-10 px-7 py-3 rounded-xl text-sm font-bold bg-white text-primary hover:bg-slate-100 shadow-lg transition-all flex items-center gap-2"
			>
				<Save class="w-4 h-4" />
				Save Changes
			</button>
		</div>

		{#if isSaved}
			<div class="p-4 rounded-2xl bg-tertiary-fixed/30 border border-tertiary/40 flex items-center gap-3 text-on-tertiary-fixed-variant text-sm">
				<CheckCircle2 class="w-5 h-5 text-tertiary shrink-0" />
				<span class="font-bold">Settings and profile preferences updated successfully!</span>
			</div>
		{/if}

		<div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
			<!-- Left Column: Personal Info & Documents -->
			<div class="lg:col-span-8 space-y-8">
				
				<!-- Personal Information Card -->
				<form onsubmit={handleSave} class="glass-panel p-8 sm:p-10 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
					<div class="flex items-center justify-between border-b border-outline-variant/30 pb-4">
						<h3 class="text-xl font-extrabold text-on-surface flex items-center gap-2">
							<User class="w-5 h-5 text-primary" />
							Personal Information
						</h3>
						<span class="text-xs font-bold text-primary uppercase tracking-wider">Editable</span>
					</div>

					<div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
						<div class="space-y-1.5">
							<label for="firstName" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">First Name</label>
							<input
								id="firstName"
								type="text"
								bind:value={firstName}
								class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
							/>
						</div>

						<div class="space-y-1.5">
							<label for="lastName" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Last Name</label>
							<input
								id="lastName"
								type="text"
								bind:value={lastName}
								class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
							/>
						</div>

						<div class="space-y-1.5">
							<label for="accEmail" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Email Address</label>
							<input
								id="accEmail"
								type="email"
								bind:value={email}
								class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
							/>
						</div>

						<div class="space-y-1.5">
							<label for="accPhone" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Phone Number</label>
							<input
								id="accPhone"
								type="tel"
								bind:value={phone}
								class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
							/>
						</div>

						<div class="space-y-1.5 sm:col-span-2">
							<label for="accDob" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Date of Birth</label>
							<input
								id="accDob"
								type="date"
								bind:value={dob}
								class="w-full sm:w-1/2 px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
							/>
						</div>
					</div>
				</form>

				<!-- Uploaded Documents Preview Card -->
				<div class="glass-panel p-8 sm:p-10 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
					<div class="flex items-center justify-between border-b border-outline-variant/30 pb-4">
						<h3 class="text-xl font-extrabold text-on-surface flex items-center gap-2">
							<FolderShared class="w-5 h-5 text-primary" />
							Uploaded Verification Documents
						</h3>
						<a href="/profile" class="text-xs font-bold text-primary hover:underline">Manage All &rarr;</a>
					</div>

					<div class="grid grid-cols-1 sm:grid-cols-3 gap-6">
						<!-- Doc 1 -->
						<div class="border border-outline-variant/40 rounded-2xl p-5 flex flex-col items-center text-center bg-white hover:border-primary/50 transition-all duration-200 group cursor-pointer shadow-sm">
							<div class="w-14 h-14 rounded-2xl bg-primary-fixed/40 text-primary flex items-center justify-center mb-3 group-hover:scale-110 transition-transform">
								<User class="w-7 h-7" />
							</div>
							<h4 class="font-bold text-on-surface text-sm">Passport Photo</h4>
							<span class="inline-flex items-center gap-1 bg-emerald-100/70 text-emerald-700 text-xs font-bold px-2.5 py-0.5 rounded-full mt-2">
								<CheckCircle2 class="w-3.5 h-3.5" /> Verified
							</span>
						</div>

						<!-- Doc 2 -->
						<div class="border border-outline-variant/40 rounded-2xl p-5 flex flex-col items-center text-center bg-white hover:border-primary/50 transition-all duration-200 group cursor-pointer shadow-sm">
							<div class="w-14 h-14 rounded-2xl bg-primary-fixed/40 text-primary flex items-center justify-center mb-3 group-hover:scale-110 transition-transform">
								<FileCheck class="w-7 h-7" />
							</div>
							<h4 class="font-bold text-on-surface text-sm">Digital Signature</h4>
							<span class="inline-flex items-center gap-1 bg-emerald-100/70 text-emerald-700 text-xs font-bold px-2.5 py-0.5 rounded-full mt-2">
								<CheckCircle2 class="w-3.5 h-3.5" /> Verified
							</span>
						</div>

						<!-- Doc 3 -->
						<div class="border border-amber-300 bg-amber-50/50 rounded-2xl p-5 flex flex-col items-center text-center hover:border-amber-400 transition-all duration-200 group cursor-pointer shadow-sm">
							<div class="w-14 h-14 rounded-2xl bg-white flex items-center justify-center mb-3 border border-amber-300 border-dashed">
								<Upload class="w-7 h-7 text-amber-600" />
							</div>
							<h4 class="font-bold text-on-surface text-sm">HSC Transcript</h4>
							<span class="inline-flex items-center gap-1 text-amber-700 text-xs font-bold px-2.5 py-0.5 rounded-full mt-2 bg-amber-100">
								<AlertTriangle class="w-3.5 h-3.5" /> Missing
							</span>
						</div>
					</div>
				</div>
			</div>

			<!-- Right Column: Security & Preferences -->
			<div class="lg:col-span-4 space-y-8">
				
				<!-- Security Card -->
				<div class="glass-panel p-7 sm:p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
					<div class="flex items-center justify-between border-b border-outline-variant/30 pb-4">
						<h3 class="text-lg font-extrabold text-on-surface flex items-center gap-2">
							<Shield class="w-5 h-5 text-primary" />
							Security
						</h3>
					</div>

					<!-- Password Change -->
					<div class="space-y-4">
						<div>
							<h4 class="font-bold text-on-surface text-sm">Password</h4>
							<p class="text-xs text-on-surface-variant mt-0.5">Last changed 3 months ago</p>
						</div>
						<div class="space-y-2">
							<input
								type="password"
								bind:value={currentPassword}
								placeholder="Current password"
								class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm focus:ring-2 focus:ring-primary/40 bg-white"
							/>
							<input
								type="password"
								bind:value={newPassword}
								placeholder="New password"
								class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm focus:ring-2 focus:ring-primary/40 bg-white"
							/>
						</div>
					</div>

					<hr class="border-outline-variant/30" />

					<!-- 2FA Toggle -->
					<div class="flex items-start justify-between gap-4">
						<div>
							<h4 class="font-bold text-on-surface text-sm">Two-Factor Auth</h4>
							<p class="text-xs text-on-surface-variant mt-0.5">Add extra account security layer via SMS.</p>
						</div>
						<input
							type="checkbox"
							bind:checked={enable2FA}
							class="w-5 h-5 rounded text-primary focus:ring-primary cursor-pointer mt-1"
						/>
					</div>
				</div>

				<!-- Notification Preferences -->
				<div class="glass-panel p-7 sm:p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
					<div class="flex items-center justify-between border-b border-outline-variant/30 pb-4">
						<h3 class="text-lg font-extrabold text-on-surface flex items-center gap-2">
							<Bell class="w-5 h-5 text-primary" />
							Notifications
						</h3>
					</div>

					<div class="space-y-4">
						<label class="flex items-center justify-between cursor-pointer">
							<span class="text-sm font-medium text-on-surface">Application Updates</span>
							<input type="checkbox" bind:checked={appUpdates} class="w-5 h-5 rounded text-primary focus:ring-primary" />
						</label>

						<label class="flex items-center justify-between cursor-pointer">
							<span class="text-sm font-medium text-on-surface">Payment Receipts</span>
							<input type="checkbox" bind:checked={paymentReceipts} class="w-5 h-5 rounded text-primary focus:ring-primary" />
						</label>

						<label class="flex items-center justify-between cursor-pointer">
							<span class="text-sm font-medium text-on-surface">Promotional Offers</span>
							<input type="checkbox" bind:checked={promoOffers} class="w-5 h-5 rounded text-primary focus:ring-primary" />
						</label>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

