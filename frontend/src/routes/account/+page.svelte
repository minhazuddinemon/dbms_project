<!-- src/routes/account/+page.svelte -->
<script lang="ts">
	import { updateStudentProfile, fetchStudentMobiles, addStudentMobile, updateStudentMobile, deleteStudentMobile } from '$lib/api/student';
	import { authState } from '$lib/state/auth.svelte';
	import { toastState } from '$lib/state/toast.svelte';
	import type { StudentMobile, StudentMobileOwnerType } from '$lib/types/models';
	import {
		User, Lock, Bell, Shield, Save, CheckCircle2, Phone, AlertTriangle,
		Smartphone, Plus, Trash2, Edit3, UserCheck, ArrowRight, FolderSync, Upload, FileCheck
	} from 'lucide-svelte';
	import { onMount } from 'svelte';

	// Profile Fields
	let presentAddress = $state('');
	let permanentAddress = $state('');
	let fathersName = $state('');
	let mothersName = $state('');
	let bloodGroup = $state('O+');
	let quota = $state('GENERAL');
	let photoUrl = $state('');
	let signatureUrl = $state('');

	// Mobiles
	let mobiles = $state<StudentMobile[]>([]);
	let newMobileNo = $state('');
	let newOwnerType = $state<StudentMobileOwnerType>('self');
	let editingMobile = $state<StudentMobile | null>(null);
	let editNewNo = $state('');
	let editNewOwnerType = $state<StudentMobileOwnerType>('self');

	// Security
	let enable2FA = $state(false);
	let appUpdates = $state(true);
	let paymentReceipts = $state(true);
	let promoOffers = $state(false);

	// State
	let activeTab = $state<'profile' | 'mobile' | 'security' | 'notifications'>('profile');
	let isLoading = $state(false);
	let isMobileLoading = $state(false);

	async function loadMobiles() {
		isMobileLoading = true;
		try {
			mobiles = await fetchStudentMobiles();
		} catch (err: any) {
			console.error('Failed to load mobiles:', err);
		} finally {
			isMobileLoading = false;
		}
	}

	onMount(() => {
		if (authState.isAuthenticated) {
			loadMobiles();
		}
	});

	async function handleSaveProfile(e: Event) {
		e.preventDefault();
		if (!authState.isAuthenticated) return;
		isLoading = true;
		try {
			await updateStudentProfile({
				PRESENT_ADDRESS: presentAddress,
				PERMANENT_ADDRESS: permanentAddress,
				FATHERS_NAME: fathersName,
				MOTHERS_NAME: mothersName,
				BLOOD_GROUP: bloodGroup,
				QUOTA: quota,
				PHOTO_URL: photoUrl,
				SIGNATURE_URL: signatureUrl
			});
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to save profile.');
		} finally {
			isLoading = false;
		}
	}

	async function handleAddMobile(e: Event) {
		e.preventDefault();
		if (!newMobileNo.trim()) return;
		try {
			await addStudentMobile({ mobile_no: newMobileNo, owner_type: newOwnerType });
			newMobileNo = '';
			newOwnerType = 'self';
			await loadMobiles();
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to add mobile.');
		}
	}

	async function handleDeleteMobile(mobileNo: string) {
		if (!confirm('Remove this mobile number?')) return;
		try {
			await deleteStudentMobile(mobileNo);
			await loadMobiles();
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to delete mobile.');
		}
	}

	function startEditMobile(m: StudentMobile) {
		editingMobile = m;
		editNewNo = m.mobile_no;
		editNewOwnerType = m.owner_type;
	}

	async function handleUpdateMobile(e: Event) {
		e.preventDefault();
		if (!editingMobile || !editNewNo.trim()) return;
		try {
			await updateStudentMobile({
				current_mobile_no: editingMobile.mobile_no,
				mobile_no: editNewNo,
				owner_type: editNewOwnerType
			});
			editingMobile = null;
			await loadMobiles();
		} catch (err: any) {
			toastState.error(err?.message || 'Failed to update mobile.');
		}
	}

	const ownerLabels: Record<StudentMobileOwnerType, string> = {
		self: 'Self',
		mother: 'Mother',
		father: 'Father'
	};
</script>

<svelte:head>
	<title>Account Settings - UniApp</title>
	<meta name="description" content="Manage your UniApp account profile, mobile numbers, security settings, and notification preferences." />
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">

		<!-- Banner Header -->
		<div class="flex flex-col md:flex-row justify-between items-start md:items-end gap-6 bg-gradient-to-r from-primary via-primary-container to-secondary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-3">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<Shield class="w-4 h-4" />
					Account Management
				</div>
				<h1 class="text-3xl sm:text-4xl font-black text-white">Account Settings</h1>
				<p class="text-slate-100 text-base leading-relaxed">
					Manage your personal information, mobile contacts, security settings, and notification preferences.
				</p>
			</div>

			{#if authState.isAuthenticated}
				<div class="relative z-10 flex items-center gap-3 bg-white/15 px-5 py-3 rounded-2xl border border-white/20 backdrop-blur-md">
					<div class="w-10 h-10 rounded-xl bg-white/20 flex items-center justify-center">
						<User class="w-5 h-5 text-white" />
					</div>
					<div>
						<p class="text-xs font-bold text-white/70 uppercase tracking-wider">Signed In As</p>
						<p class="text-sm font-extrabold text-white">{authState.user?.email || 'Student'}</p>
					</div>
				</div>
			{/if}
		</div>

		{#if !authState.isAuthenticated}
			<div class="glass-panel rounded-[2.5rem] border border-outline-variant/30 p-10 text-center max-w-md mx-auto space-y-5 bg-white/90 shadow-xl">
				<div class="w-16 h-16 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center mx-auto">
					<UserCheck class="w-8 h-8" />
				</div>
				<h3 class="text-2xl font-extrabold text-on-surface">Sign In Required</h3>
				<p class="text-on-surface-variant text-sm">Please log in to manage your account settings.</p>
				<div class="flex gap-3 justify-center">
					<a href="/login" class="px-6 py-3 rounded-xl font-bold bg-primary text-white text-sm hover:bg-primary-container shadow-lg shadow-primary/20 transition-all">Sign In</a>
					<a href="/register" class="px-6 py-3 rounded-xl font-bold border border-outline-variant text-on-surface text-sm hover:bg-surface-container transition-all">Register</a>
				</div>
			</div>
		{:else}
			<!-- Tabs -->
			<div class="flex items-center gap-2 bg-white/80 p-2 rounded-2xl border border-outline-variant/30 shadow-sm overflow-x-auto">
				{#each [['profile', User, 'Profile Fields'], ['mobile', Phone, 'Mobile Numbers'], ['security', Lock, 'Security'], ['notifications', Bell, 'Notifications']] as [tab, Icon, label]}
					<button
						onclick={() => activeTab = tab as any}
						class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all flex items-center justify-center gap-2 whitespace-nowrap {activeTab === tab ? 'bg-primary text-white shadow-md' : 'text-on-surface-variant hover:text-on-surface'}"
					>
						<svelte:component this={Icon} class="w-4 h-4" />
						{label}
					</button>
				{/each}
			</div>

			<!-- Tab: Profile Fields -->
			{#if activeTab === 'profile'}
				<div class="glass-panel p-8 sm:p-10 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
					<div class="flex items-center justify-between border-b border-outline-variant/30 pb-4">
						<h2 class="text-xl font-extrabold text-on-surface flex items-center gap-2">
							<User class="w-5 h-5 text-primary" />
							Personal & Contact Information
						</h2>
						<a href="/profile" class="text-xs font-bold text-primary hover:underline flex items-center gap-1">
							Full Academic Profile <ArrowRight class="w-3.5 h-3.5" />
						</a>
					</div>

					<form onsubmit={handleSaveProfile} class="space-y-6">
						<div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
							<div class="space-y-1.5">
								<label for="fathersName" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Father's Name</label>
								<input id="fathersName" type="text" bind:value={fathersName} placeholder="Md. Anowar Hossain" class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white" />
							</div>
							<div class="space-y-1.5">
								<label for="mothersName" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Mother's Name</label>
								<input id="mothersName" type="text" bind:value={mothersName} placeholder="Sultana Begum" class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white" />
							</div>
							<div class="space-y-1.5">
								<label for="bloodGroup" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Blood Group</label>
								<select id="bloodGroup" bind:value={bloodGroup} class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white font-semibold">
									{#each ['A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-'] as bg}
										<option value={bg}>{bg}</option>
									{/each}
								</select>
							</div>
							<div class="space-y-1.5">
								<label for="quota" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Quota</label>
								<select id="quota" bind:value={quota} class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white font-semibold">
									{#each ['GENERAL', 'FREEDOM_FIGHTER', 'DISTRICT', 'TRIBAL'] as q}
										<option value={q}>{q}</option>
									{/each}
								</select>
							</div>
							<div class="space-y-1.5 sm:col-span-2">
								<label for="presentAddr" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Present Address</label>
								<input id="presentAddr" type="text" bind:value={presentAddress} placeholder="Dhaka, Bangladesh" class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white" />
							</div>
							<div class="space-y-1.5 sm:col-span-2">
								<label for="permanentAddr" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Permanent Address</label>
								<input id="permanentAddr" type="text" bind:value={permanentAddress} placeholder="Tangail, Bangladesh" class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white" />
							</div>
						</div>

						<!-- Document URLs -->
						<div class="space-y-4 p-5 rounded-2xl bg-surface-container-low/50 border border-outline-variant/30">
							<h3 class="text-sm font-extrabold text-on-surface flex items-center gap-2">
								<FolderSync class="w-4 h-4 text-primary" />
								Document URLs (To be Added)
							</h3>
							<div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
								<div class="space-y-1.5">
									<label for="photoUrl" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Photo URL</label>
									<input id="photoUrl" type="url" bind:value={photoUrl} placeholder="https://example.com/photo.jpg" class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white font-mono text-xs" />
								</div>
								<div class="space-y-1.5">
									<label for="signatureUrl" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Signature URL</label>
									<input id="signatureUrl" type="url" bind:value={signatureUrl} placeholder="https://example.com/signature.png" class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white font-mono text-xs" />
								</div>
							</div>
							<p class="text-xs text-on-surface-variant">
								Upload your photo and signature to an image hosting service and paste the URL here. These are required for application submission.
							</p>
						</div>

						<button
							type="submit"
							disabled={isLoading}
							class="w-full sm:w-auto px-8 py-3.5 rounded-xl font-bold text-white bg-primary hover:bg-primary-container disabled:opacity-50 shadow-lg shadow-primary/25 transition-all flex items-center gap-2 text-sm"
						>
							<Save class="w-4 h-4" />
							{isLoading ? 'Saving...' : 'Save Profile Fields'}
						</button>
					</form>
				</div>
			{/if}

			<!-- Tab: Mobile Numbers -->
			{#if activeTab === 'mobile'}
				<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
					<div class="flex items-center justify-between border-b border-outline-variant/30 pb-4">
						<h2 class="text-xl font-extrabold text-on-surface flex items-center gap-2">
							<Smartphone class="w-5 h-5 text-primary" />
							Mobile Numbers
						</h2>
						<span class="text-xs font-bold text-on-surface-variant">{mobiles.length} saved</span>
					</div>

					<!-- Add New Mobile Form -->
					{#if !editingMobile}
						<form onsubmit={handleAddMobile} class="flex flex-col sm:flex-row gap-4 items-end bg-surface-container-low/50 p-5 rounded-2xl border border-outline-variant/30">
							<div class="space-y-1.5 flex-1 w-full">
								<label for="newMobile" class="block text-xs font-bold uppercase text-on-surface-variant">Mobile Number</label>
								<input id="newMobile" type="tel" bind:value={newMobileNo} placeholder="e.g. 01711111111" required class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm bg-white" />
							</div>
							<div class="space-y-1.5 w-full sm:w-40">
								<label for="ownerType" class="block text-xs font-bold uppercase text-on-surface-variant">Owner Type</label>
								<select id="ownerType" bind:value={newOwnerType} class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm bg-white font-semibold">
									<option value="self">Self</option>
									<option value="mother">Mother</option>
									<option value="father">Father</option>
								</select>
							</div>
							<button type="submit" class="px-6 py-2.5 rounded-xl font-bold text-white bg-primary text-sm shadow-md hover:bg-primary-container transition-all flex items-center gap-2 shrink-0">
								<Plus class="w-4 h-4" />
								Add
							</button>
						</form>
					{:else}
						<!-- Edit Form -->
						<form onsubmit={handleUpdateMobile} class="p-5 rounded-2xl border border-primary/30 bg-primary-fixed/20 space-y-4">
							<h3 class="text-sm font-bold text-primary">Edit Mobile Number</h3>
							<div class="flex flex-col sm:flex-row gap-4 items-end">
								<div class="space-y-1.5 flex-1 w-full">
									<label for="editMobile" class="block text-xs font-bold uppercase text-on-surface-variant">New Number</label>
									<input id="editMobile" type="tel" bind:value={editNewNo} required class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm bg-white" />
								</div>
								<div class="space-y-1.5 w-full sm:w-40">
									<label for="editOwnerType" class="block text-xs font-bold uppercase text-on-surface-variant">Owner Type</label>
									<select id="editOwnerType" bind:value={editNewOwnerType} class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm bg-white font-semibold">
										<option value="self">Self</option>
										<option value="mother">Mother</option>
										<option value="father">Father</option>
									</select>
								</div>
								<div class="flex gap-2 shrink-0">
									<button type="submit" class="px-5 py-2.5 rounded-xl font-bold text-white bg-primary text-sm shadow-md hover:bg-primary-container transition-all">Save</button>
									<button type="button" onclick={() => editingMobile = null} class="px-5 py-2.5 rounded-xl font-bold text-on-surface border border-outline-variant/50 text-sm hover:bg-surface-container transition-all">Cancel</button>
								</div>
							</div>
						</form>
					{/if}

					<!-- Mobiles List -->
					{#if isMobileLoading}
						<div class="py-8 text-center text-on-surface-variant">
							<div class="w-8 h-8 border-4 border-primary border-t-transparent rounded-full animate-spin mx-auto mb-2"></div>
							<p class="text-sm font-semibold">Loading contacts...</p>
						</div>
					{:else if mobiles.length === 0}
						<div class="py-8 text-center space-y-2">
							<Phone class="w-10 h-10 text-outline mx-auto" />
							<p class="text-sm font-semibold text-on-surface-variant">No mobile numbers saved yet.</p>
							<p class="text-xs text-on-surface-variant">Add your own number and your parents' numbers for notifications.</p>
						</div>
					{:else}
						<div class="space-y-3">
							{#each mobiles as m}
								<div class="flex items-center justify-between p-4 rounded-2xl border border-outline-variant/30 bg-white hover:border-primary/30 transition-all group">
									<div class="flex items-center gap-3">
										<div class="w-10 h-10 rounded-xl bg-primary-fixed/40 text-primary flex items-center justify-center">
											<Phone class="w-5 h-5" />
										</div>
										<div>
											<p class="font-bold text-on-surface font-mono">{m.mobile_no}</p>
											<p class="text-xs text-on-surface-variant">{ownerLabels[m.owner_type]}</p>
										</div>
									</div>
									<div class="flex gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
										<button onclick={() => startEditMobile(m)} class="p-2 rounded-lg text-primary hover:bg-primary-fixed/40 transition-colors">
											<Edit3 class="w-4 h-4" />
										</button>
										<button onclick={() => handleDeleteMobile(m.mobile_no)} class="p-2 rounded-lg text-error hover:bg-error-container/40 transition-colors">
											<Trash2 class="w-4 h-4" />
										</button>
									</div>
								</div>
							{/each}
						</div>
					{/if}
				</div>
			{/if}

			<!-- Tab: Security -->
			{#if activeTab === 'security'}
				<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
					<div class="flex items-center border-b border-outline-variant/30 pb-4">
						<h2 class="text-xl font-extrabold text-on-surface flex items-center gap-2">
							<Lock class="w-5 h-5 text-primary" />
							Security Settings
						</h2>
					</div>

					<div class="p-5 rounded-2xl bg-amber-50 border border-amber-200 space-y-3">
						<div class="flex items-start gap-3">
							<AlertTriangle class="w-5 h-5 text-amber-600 shrink-0 mt-0.5" />
							<div>
								<h3 class="font-bold text-amber-900 text-sm">Password Management</h3>
								<p class="text-amber-700 text-xs mt-1">Password change is managed through the backend. Contact support or re-register if you need to reset your password.</p>
							</div>
						</div>
					</div>

					<div class="space-y-4">
						<div class="flex items-center justify-between p-5 rounded-2xl border border-outline-variant/30 bg-white">
							<div>
								<h4 class="font-bold text-on-surface text-sm">Signed In Account</h4>
								<p class="text-xs text-on-surface-variant mt-0.5">{authState.user?.email || 'N/A'}</p>
							</div>
							<CheckCircle2 class="w-5 h-5 text-emerald-500" />
						</div>

						<div class="flex items-start justify-between p-5 rounded-2xl border border-outline-variant/30 bg-white gap-4">
							<div>
								<h4 class="font-bold text-on-surface text-sm">Two-Factor Authentication</h4>
								<p class="text-xs text-on-surface-variant mt-0.5">Add extra account security via SMS verification.</p>
							</div>
							<label class="flex items-center cursor-pointer shrink-0">
								<input type="checkbox" bind:checked={enable2FA} class="w-5 h-5 rounded text-primary focus:ring-primary" />
							</label>
						</div>
					</div>
				</div>
			{/if}

			<!-- Tab: Notifications -->
			{#if activeTab === 'notifications'}
				<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
					<div class="flex items-center border-b border-outline-variant/30 pb-4">
						<h2 class="text-xl font-extrabold text-on-surface flex items-center gap-2">
							<Bell class="w-5 h-5 text-primary" />
							Notification Preferences
						</h2>
					</div>

					<div class="space-y-4">
						{#each [['appUpdates', 'Application Updates', 'Receive alerts when your application status changes.', appUpdates], ['paymentReceipts', 'Payment Receipts', 'Get receipts for every transaction made.', paymentReceipts], ['promoOffers', 'Promotional Offers', 'Optional announcements from partner universities.', promoOffers]] as [id, label, desc, checked]}
							<label class="flex items-start justify-between p-5 rounded-2xl border border-outline-variant/30 bg-white cursor-pointer hover:border-primary/30 transition-all gap-4 group">
								<div>
									<h4 class="font-bold text-on-surface text-sm group-hover:text-primary transition-colors">{label}</h4>
									<p class="text-xs text-on-surface-variant mt-0.5">{desc}</p>
								</div>
								<input
									type="checkbox"
									checked={checked as boolean}
									class="w-5 h-5 rounded text-primary focus:ring-primary cursor-pointer shrink-0 mt-1"
									onchange={(e) => {
										if (id === 'appUpdates') appUpdates = (e.target as HTMLInputElement).checked;
										else if (id === 'paymentReceipts') paymentReceipts = (e.target as HTMLInputElement).checked;
										else promoOffers = (e.target as HTMLInputElement).checked;
									}}
								/>
							</label>
						{/each}
					</div>

					<button class="px-8 py-3.5 rounded-xl font-bold text-white bg-primary hover:bg-primary-container shadow-lg shadow-primary/25 transition-all flex items-center gap-2 text-sm">
						<Save class="w-4 h-4" />
						Save Preferences
					</button>
				</div>
			{/if}
		{/if}
	</div>
</div>
