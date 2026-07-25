<!-- src/routes/profile/+page.svelte -->
<script lang="ts">
	import { updateStudentProfile } from '$lib/api/student';
	import { authState } from '$lib/state/auth.svelte';
	import { User, MapPin, Award, CheckCircle, AlertCircle, Save, FileText, Camera, Shield } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let presentAddress = $state('Dhaka, Bangladesh');
	let permanentAddress = $state('Dhaka, Bangladesh');
	let fathersName = $state('');
	let mothersName = $state('');
	let bloodGroup = $state('O+');
	let quota = $state('GENERAL');
	let photoUrl = $state('https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=150');

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
				PHOTO_URL: photoUrl
			});
			successMessage = 'Profile updated successfully!';
		} catch (err: any) {
			errorMessage = err?.message || 'Failed to update profile fields.';
		} finally {
			isLoading = false;
		}
	}
</script>

<svelte:head>
	<title>Student Academic Profile - UniApp</title>
</svelte:head>

<div class="py-10 bg-slate-50 min-h-screen">
	<div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
		<!-- Header -->
		<div>
			<h1 class="text-3xl font-extrabold text-slate-900 sm:text-4xl">Student Profile & Academic Information</h1>
			<p class="mt-2 text-slate-600">Provide required personal and academic details to complete program applications</p>
		</div>

		{#if !authState.isAuthenticated}
			<div class="bg-white rounded-2xl border border-slate-200 p-8 text-center space-y-4">
				<User class="w-12 h-12 text-slate-400 mx-auto" />
				<h3 class="text-lg font-bold text-slate-800">Sign In Required</h3>
				<p class="text-slate-600 text-sm">Please log in to manage your profile.</p>
				<a href="/login" class="inline-block px-6 py-2.5 rounded-xl font-semibold bg-indigo-600 text-white text-sm">Sign In</a>
			</div>
		{:else}
			{#if successMessage}
				<div class="p-4 rounded-xl bg-emerald-50 border border-emerald-200 flex items-center gap-3 text-emerald-700 text-sm">
					<CheckCircle class="w-5 h-5 text-emerald-500 shrink-0" />
					<span>{successMessage}</span>
				</div>
			{/if}

			{#if errorMessage}
				<div class="p-4 rounded-xl bg-red-50 border border-red-200 flex items-center gap-3 text-red-700 text-sm">
					<AlertCircle class="w-5 h-5 text-red-500 shrink-0" />
					<span>{errorMessage}</span>
				</div>
			{/if}

			<form onsubmit={handleSaveProfile} class="bg-white rounded-2xl border border-slate-200 p-8 shadow-sm space-y-6">
				<h3 class="text-xl font-bold text-slate-900 border-b border-slate-100 pb-4 flex items-center gap-2">
					<User class="w-5 h-5 text-indigo-600" />
					Personal & Guardian Information
				</h3>

				<div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
					<div>
						<label for="fathersName" class="block text-sm font-semibold text-slate-700 mb-1">Father's Name</label>
						<input
							id="fathersName"
							type="text"
							bind:value={fathersName}
							placeholder="Father's full name"
							class="w-full px-4 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 text-sm"
						/>
					</div>

					<div>
						<label for="mothersName" class="block text-sm font-semibold text-slate-700 mb-1">Mother's Name</label>
						<input
							id="mothersName"
							type="text"
							bind:value={mothersName}
							placeholder="Mother's full name"
							class="w-full px-4 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 text-sm"
						/>
					</div>

					<div>
						<label for="bloodGroup" class="block text-sm font-semibold text-slate-700 mb-1">Blood Group</label>
						<select
							id="bloodGroup"
							bind:value={bloodGroup}
							class="w-full px-4 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 text-sm bg-white"
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

					<div>
						<label for="quota" class="block text-sm font-semibold text-slate-700 mb-1">Admission Quota</label>
						<select
							id="quota"
							bind:value={quota}
							class="w-full px-4 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 text-sm bg-white"
						>
							<option value="GENERAL">General</option>
							<option value="FREEDOM_FIGHTER">Freedom Fighter Quota</option>
							<option value="TRIBAL">Tribal / Ethnic Quota</option>
							<option value="PHYSICALLY_CHALLENGED">Physically Challenged</option>
						</select>
					</div>
				</div>

				<h3 class="text-xl font-bold text-slate-900 border-b border-slate-100 pb-4 pt-4 flex items-center gap-2">
					<MapPin class="w-5 h-5 text-indigo-600" />
					Address Details
				</h3>

				<div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
					<div>
						<label for="presentAddress" class="block text-sm font-semibold text-slate-700 mb-1">Present Address</label>
						<textarea
							id="presentAddress"
							bind:value={presentAddress}
							rows="3"
							placeholder="House, Road, Thana, District"
							class="w-full px-4 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 text-sm"
						></textarea>
					</div>

					<div>
						<label for="permanentAddress" class="block text-sm font-semibold text-slate-700 mb-1">Permanent Address</label>
						<textarea
							id="permanentAddress"
							bind:value={permanentAddress}
							rows="3"
							placeholder="House, Village, Upazila, District"
							class="w-full px-4 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 text-sm"
						></textarea>
					</div>
				</div>

				<div class="pt-4 border-t border-slate-100 flex justify-end">
					<button
						type="submit"
						disabled={isLoading}
						class="px-8 py-3 rounded-xl font-bold text-white bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 shadow-lg shadow-indigo-600/25 transition-all flex items-center gap-2 text-sm"
					>
						<Save class="w-4 h-4" />
						{isLoading ? 'Saving Changes...' : 'Save Profile Info'}
					</button>
				</div>
			</form>
		{/if}
	</div>
</div>
