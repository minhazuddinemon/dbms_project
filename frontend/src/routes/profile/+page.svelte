<!-- src/routes/profile/+page.svelte -->
<script lang="ts">
	import { updateStudentProfile } from '$lib/api/student';
	import { authState } from '$lib/state/auth.svelte';
	import { User, MapPin, Award, CheckCircle2, AlertCircle, Save, FileText, BookOpen, Shield, Sparkles, FolderSync, Upload, Image, FileCheck, AlertTriangle } from 'lucide-svelte';

	let presentAddress = $state('Dhaka, Bangladesh');
	let permanentAddress = $state('Dhaka, Bangladesh');
	let fathersName = $state('');
	let mothersName = $state('');
	let bloodGroup = $state('O+');
	let quota = $state('GENERAL');
	let sscGpa = $state('5.00');
	let hscGpa = $state('5.00');
	let hscGroup = $state('Science');
	let physicsMarks = $state('85');
	let mathMarks = $state('90');
	let chemistryMarks = $state('88');

	let activeTab = $state<'personal' | 'academics' | 'address' | 'documents'>('personal');
	let isLoading = $state(false);
	let successMessage = $state<string | null>(null);
	let errorMessage = $state<string | null>(null);

	async function handleSaveProfile(e: Event) {
		e.preventDefault();
		isLoading = true;
		successMessage = null;
		errorMessage = null;

		try {
			await updateStudentProfile({
				PRESENT_ADDRESS: presentAddress,
				PERMANENT_ADDRESS: permanentAddress,
				FATHERS_NAME: fathersName,
				MOTHERS_NAME: mothersName,
				BLOOD_GROUP: bloodGroup,
				QUOTA: quota,
				SSC_GPA: sscGpa,
				HSC_GPA: hscGpa,
				HSC_GROUP: hscGroup,
				PHYSICS_MARKS: physicsMarks,
				MATH_MARKS: mathMarks,
				CHEMISTRY_MARKS: chemistryMarks
			});
			successMessage = 'Academic Profile and Subject Marks saved successfully!';
		} catch (err: any) {
			errorMessage = err?.message || 'Failed to update academic profile fields.';
		} finally {
			isLoading = false;
		}
	}
</script>

<svelte:head>
	<title>Academic Profile & Marks - UniApp</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		
		<!-- Banner Header -->
		<div class="bg-gradient-to-r from-primary via-primary-container to-secondary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-3">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<BookOpen class="w-4 h-4" />
					Academic Data Collection
				</div>

				<h1 class="text-3xl sm:text-4xl font-black text-white">
					Student Academic Profile
				</h1>

				<p class="text-slate-100 text-base leading-relaxed">
					Provide your guardian details, SSC/HSC GPA, and individual subject marks for automated eligibility verification.
				</p>
			</div>
		</div>

		{#if !authState.isAuthenticated}
			<div class="glass-panel rounded-[2.5rem] border border-outline-variant/30 p-10 text-center space-y-4 bg-white/90 shadow-xl">
				<User class="w-14 h-14 text-outline mx-auto" />
				<h3 class="text-2xl font-bold text-on-surface">Sign In Required</h3>
				<p class="text-on-surface-variant text-sm">Please sign in to update your academic information.</p>
				<a href="/login" class="inline-block px-8 py-3 rounded-xl font-bold bg-primary text-white text-sm hover:bg-primary-container shadow-lg shadow-primary/25 transition-all">Sign In</a>
			</div>
		{:else}
			{#if successMessage}
				<div class="p-4 rounded-2xl bg-tertiary-fixed/30 border border-tertiary/40 flex items-center gap-3 text-on-tertiary-fixed-variant text-sm">
					<CheckCircle2 class="w-5 h-5 text-tertiary shrink-0" />
					<span class="font-bold">{successMessage}</span>
				</div>
			{/if}

			{#if errorMessage}
				<div class="p-4 rounded-2xl bg-error-container/60 border border-error/30 flex items-center gap-3 text-on-error-container text-sm">
					<AlertCircle class="w-5 h-5 text-error shrink-0" />
					<span>{errorMessage}</span>
				</div>
			{/if}

			<!-- Step Tabs Navigation -->
			<div class="flex items-center gap-2 bg-surface-container-low/80 p-2 rounded-2xl border border-outline-variant/30 overflow-x-auto">
				<button
					type="button"
					onclick={() => activeTab = 'personal'}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all duration-200 flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'personal' ? 'bg-white text-primary shadow-sm border border-outline-variant/30' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<User class="w-4 h-4" />
					Guardian & Personal
				</button>
				<button
					type="button"
					onclick={() => activeTab = 'academics'}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all duration-200 flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'academics' ? 'bg-white text-primary shadow-sm border border-outline-variant/30' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<BookOpen class="w-4 h-4" />
					HSC Marks & GPA
				</button>
				<button
					type="button"
					onclick={() => activeTab = 'address'}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all duration-200 flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'address' ? 'bg-white text-primary shadow-sm border border-outline-variant/30' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<MapPin class="w-4 h-4" />
					Address & Quota
				</button>
				<button
					type="button"
					onclick={() => activeTab = 'documents'}
					class="flex-1 py-3 px-4 rounded-xl font-bold text-sm transition-all duration-200 flex items-center justify-center gap-2 whitespace-nowrap {activeTab === 'documents' ? 'bg-white text-primary shadow-sm border border-outline-variant/30' : 'text-on-surface-variant hover:text-on-surface'}"
				>
					<FolderSync class="w-4 h-4" />
					Documents
				</button>
			</div>

			<!-- Main Form Panel -->
			<form onsubmit={handleSaveProfile} class="glass-panel rounded-[2.5rem] border border-outline-variant/40 p-8 sm:p-10 shadow-xl bg-white/95 space-y-8">
				
				{#if activeTab === 'personal'}
					<div class="space-y-6 animate-fade-in-up">
						<h3 class="text-xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-3 flex items-center gap-2">
							<User class="w-5 h-5 text-primary" />
							Guardian Information
						</h3>

						<div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
							<div class="space-y-1.5">
								<label for="fathersName" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Father's Name</label>
								<input
									id="fathersName"
									type="text"
									bind:value={fathersName}
									placeholder="e.g. Md. Anowar Hossain"
									class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all bg-white"
								/>
							</div>

							<div class="space-y-1.5">
								<label for="mothersName" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Mother's Name</label>
								<input
									id="mothersName"
									type="text"
									bind:value={mothersName}
									placeholder="e.g. Sultana Begum"
									class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all bg-white"
								/>
							</div>

							<div class="space-y-1.5">
								<label for="bloodGroup" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Blood Group</label>
								<select
									id="bloodGroup"
									bind:value={bloodGroup}
									class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
								>
									<option value="A+">A+</option>
									<option value="A-">A-</option>
									<option value="B+">B+</option>
									<option value="B-">B-</option>
									<option value="O+">O+</option>
									<option value="O-">O-</option>
									<option value="AB+">AB+</option>
									<option value="AB-">AB-</option>
								</select>
							</div>
						</div>
					</div>
				{/if}

				{#if activeTab === 'academics'}
					<div class="space-y-6 animate-fade-in-up">
						<h3 class="text-xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-3 flex items-center gap-2">
							<BookOpen class="w-5 h-5 text-primary" />
							GPA & Subject Marks Verification
						</h3>

						<div class="grid grid-cols-1 sm:grid-cols-3 gap-6">
							<div class="space-y-1.5">
								<label for="sscGpa" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">SSC GPA</label>
								<input
									id="sscGpa"
									type="text"
									bind:value={sscGpa}
									placeholder="5.00"
									class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
								/>
							</div>

							<div class="space-y-1.5">
								<label for="hscGpa" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">HSC GPA</label>
								<input
									id="hscGpa"
									type="text"
									bind:value={hscGpa}
									placeholder="5.00"
									class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
								/>
							</div>

							<div class="space-y-1.5">
								<label for="hscGroup" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">HSC Group</label>
								<select
									id="hscGroup"
									bind:value={hscGroup}
									class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
								>
									<option value="Science">Science</option>
									<option value="Humanities">Humanities</option>
									<option value="Commerce">Commerce / Business Studies</option>
								</select>
							</div>
						</div>

						<div class="pt-4 space-y-4">
							<span class="block text-sm font-extrabold text-on-surface uppercase tracking-wider">HSC Subject Marks (Out of 100)</span>
							<div class="grid grid-cols-1 sm:grid-cols-3 gap-6">
								<div class="space-y-1.5">
									<label for="physics" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Physics</label>
									<input
										id="physics"
										type="number"
										bind:value={physicsMarks}
										placeholder="85"
										class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
									/>
								</div>
								<div class="space-y-1.5">
									<label for="math" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Mathematics</label>
									<input
										id="math"
										type="number"
										bind:value={mathMarks}
										placeholder="90"
										class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
									/>
								</div>
								<div class="space-y-1.5">
									<label for="chemistry" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Chemistry</label>
									<input
										id="chemistry"
										type="number"
										bind:value={chemistryMarks}
										placeholder="88"
										class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
									/>
								</div>
							</div>
						</div>
					</div>
				{/if}

				{#if activeTab === 'address'}
					<div class="space-y-6 animate-fade-in-up">
						<h3 class="text-xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-3 flex items-center gap-2">
							<MapPin class="w-5 h-5 text-primary" />
							Address & Special Quotas
						</h3>

						<div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
							<div class="space-y-1.5">
								<label for="quota" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Admission Quota</label>
								<select
									id="quota"
									bind:value={quota}
									class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
								>
									<option value="GENERAL">General</option>
									<option value="FREEDOM_FIGHTER">Freedom Fighter Quota</option>
									<option value="TRIBAL">Tribal / Ethnic Quota</option>
									<option value="PHYSICALLY_CHALLENGED">Physically Challenged</option>
								</select>
							</div>

							<div class="space-y-1.5 sm:col-span-2">
								<label for="presentAddress" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Present Address</label>
								<textarea
									id="presentAddress"
									bind:value={presentAddress}
									rows="3"
									placeholder="House, Road, Thana, District"
									class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
								></textarea>
							</div>

							<div class="space-y-1.5 sm:col-span-2">
								<label for="permanentAddress" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Permanent Address</label>
								<textarea
									id="permanentAddress"
									bind:value={permanentAddress}
									rows="3"
									placeholder="House, Village, Upazila, District"
									class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
								></textarea>
							</div>
						</div>
					</div>
				{/if}

				{#if activeTab === 'documents'}
					<div class="space-y-6 animate-fade-in-up">
						<h3 class="text-xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-3 flex items-center gap-2">
							<FolderSync class="w-5 h-5 text-primary" />
							Uploaded Documents
						</h3>

						<div class="grid grid-cols-1 sm:grid-cols-3 gap-6">
							<!-- Doc 1 -->
							<div class="border border-outline-variant/40 rounded-2xl p-5 flex flex-col items-center text-center bg-white hover:border-primary/50 transition-all duration-200 group cursor-pointer shadow-sm">
								<div class="w-16 h-16 rounded-2xl bg-primary-fixed/40 text-primary flex items-center justify-center mb-3 group-hover:scale-110 transition-transform">
									<User class="w-8 h-8" />
								</div>
								<h4 class="font-bold text-on-surface text-sm">Passport Photo</h4>
								<span class="inline-flex items-center gap-1 bg-emerald-100/70 text-emerald-700 text-xs font-bold px-2.5 py-0.5 rounded-full mt-2">
									<CheckCircle2 class="w-3.5 h-3.5" /> Verified
								</span>
								<div class="w-full mt-4 pt-4 border-t border-outline-variant/20 flex gap-2">
									<button type="button" class="flex-1 text-xs font-bold text-on-surface-variant hover:text-primary transition-colors">Preview</button>
									<button type="button" class="flex-1 text-xs font-bold text-on-surface-variant hover:text-primary transition-colors">Replace</button>
								</div>
							</div>

							<!-- Doc 2 -->
							<div class="border border-outline-variant/40 rounded-2xl p-5 flex flex-col items-center text-center bg-white hover:border-primary/50 transition-all duration-200 group cursor-pointer shadow-sm">
								<div class="w-16 h-16 rounded-2xl bg-primary-fixed/40 text-primary flex items-center justify-center mb-3 group-hover:scale-110 transition-transform">
									<FileCheck class="w-8 h-8" />
								</div>
								<h4 class="font-bold text-on-surface text-sm">Digital Signature</h4>
								<span class="inline-flex items-center gap-1 bg-emerald-100/70 text-emerald-700 text-xs font-bold px-2.5 py-0.5 rounded-full mt-2">
									<CheckCircle2 class="w-3.5 h-3.5" /> Verified
								</span>
								<div class="w-full mt-4 pt-4 border-t border-outline-variant/20 flex gap-2">
									<button type="button" class="flex-1 text-xs font-bold text-on-surface-variant hover:text-primary transition-colors">Preview</button>
									<button type="button" class="flex-1 text-xs font-bold text-on-surface-variant hover:text-primary transition-colors">Replace</button>
								</div>
							</div>

							<!-- Doc 3 (Missing / Action Required) -->
							<div class="border border-amber-300 bg-amber-50/50 rounded-2xl p-5 flex flex-col items-center text-center hover:border-amber-400 transition-all duration-200 group cursor-pointer shadow-sm">
								<div class="w-16 h-16 rounded-2xl bg-white flex items-center justify-center mb-3 border border-amber-300 border-dashed">
									<Upload class="w-8 h-8 text-amber-600" />
								</div>
								<h4 class="font-bold text-on-surface text-sm">HSC Transcript</h4>
								<span class="inline-flex items-center gap-1 text-amber-700 text-xs font-bold px-2.5 py-0.5 rounded-full mt-2 bg-amber-100">
									<AlertTriangle class="w-3.5 h-3.5" /> Action Required
								</span>
								<div class="w-full mt-4 pt-4 border-t border-amber-200 flex gap-2">
									<button type="button" class="flex-1 text-xs font-bold text-amber-700 hover:underline transition-colors">Upload Now</button>
								</div>
							</div>
						</div>
					</div>
				{/if}

				<!-- Save Action Bar -->
				<div class="pt-6 border-t border-outline-variant/30 flex items-center justify-between">
					<span class="text-xs text-on-surface-variant font-medium">All changes automatically validated against REST API</span>
					<button
						type="submit"
						disabled={isLoading}
						class="px-8 py-3.5 rounded-xl font-bold text-white bg-primary hover:bg-primary-container disabled:opacity-50 shadow-lg shadow-primary/25 hover:shadow-primary/40 transition-all flex items-center gap-2 text-sm"
					>
						<Save class="w-4 h-4" />
						{isLoading ? 'Saving...' : 'Save Academic Profile'}
					</button>
				</div>
			</form>
		{/if}
	</div>
</div>

