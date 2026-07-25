<!-- src/routes/register/+page.svelte -->
<script lang="ts">
	import { register } from '$lib/api/auth';
	import { goto } from '$app/navigation';
	import { UserPlus, User, Mail, Lock, Calendar, AlertCircle, CheckCircle } from 'lucide-svelte';

	let firstName = $state('');
	let lastName = $state('');
	let email = $state('');
	let password = $state('');
	let dob = $state('');

	let isLoading = $state(false);
	let errorMessage = $state<string | null>(null);
	let successMessage = $state<string | null>(null);

	async function handleSubmit(e: Event) {
		e.preventDefault();
		if (!firstName || !lastName || !email || !password || !dob) {
			errorMessage = 'All fields are required.';
			return;
		}

		isLoading = true;
		errorMessage = null;
		successMessage = null;

		try {
			await register({
				first_name: firstName,
				last_name: lastName,
				email,
				password,
				dob
			});
			successMessage = 'Account registered successfully! Redirecting to sign in...';
			setTimeout(() => {
				goto('/login');
			}, 1500);
		} catch (err: any) {
			errorMessage = err?.message || 'Registration failed. Email might already exist.';
		} finally {
			isLoading = false;
		}
	}
</script>

<svelte:head>
	<title>Register Account - UniApp</title>
</svelte:head>

<div class="min-h-[80vh] flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8 bg-slate-50">
	<div class="max-w-lg w-full space-y-8 bg-white p-8 rounded-2xl border border-slate-200 shadow-xl shadow-slate-200/50">
		<div class="text-center">
			<div class="w-12 h-12 rounded-2xl bg-indigo-600 text-white flex items-center justify-center mx-auto mb-4 shadow-lg shadow-indigo-600/30">
				<UserPlus class="w-6 h-6" />
			</div>
			<h2 class="text-2xl font-extrabold text-slate-900">Create Student Account</h2>
			<p class="mt-2 text-sm text-slate-600">Register to apply for public university programs in Bangladesh</p>
		</div>

		{#if errorMessage}
			<div class="p-4 rounded-xl bg-red-50 border border-red-200 flex items-start gap-3 text-red-700 text-sm">
				<AlertCircle class="w-5 h-5 shrink-0 text-red-500 mt-0.5" />
				<span>{errorMessage}</span>
			</div>
		{/if}

		{#if successMessage}
			<div class="p-4 rounded-xl bg-emerald-50 border border-emerald-200 flex items-start gap-3 text-emerald-700 text-sm">
				<CheckCircle class="w-5 h-5 shrink-0 text-emerald-500 mt-0.5" />
				<span>{successMessage}</span>
			</div>
		{/if}

		<form onsubmit={handleSubmit} class="space-y-4">
			<div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
				<div>
					<label for="firstName" class="block text-sm font-semibold text-slate-700 mb-1">First Name</label>
					<div class="relative">
						<div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
							<User class="w-4 h-4" />
						</div>
						<input
							id="firstName"
							type="text"
							bind:value={firstName}
							required
							placeholder="Rahim"
							class="w-full pl-9 pr-3 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 focus:border-indigo-600 text-sm"
						/>
					</div>
				</div>

				<div>
					<label for="lastName" class="block text-sm font-semibold text-slate-700 mb-1">Last Name</label>
					<div class="relative">
						<div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
							<User class="w-4 h-4" />
						</div>
						<input
							id="lastName"
							type="text"
							bind:value={lastName}
							required
							placeholder="Uddin"
							class="w-full pl-9 pr-3 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 focus:border-indigo-600 text-sm"
						/>
					</div>
				</div>
			</div>

			<div>
				<label for="email" class="block text-sm font-semibold text-slate-700 mb-1">Email Address</label>
				<div class="relative">
					<div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
						<Mail class="w-4 h-4" />
					</div>
					<input
						id="email"
						type="email"
						bind:value={email}
						required
						placeholder="student@example.com"
						class="w-full pl-9 pr-3 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 focus:border-indigo-600 text-sm"
					/>
				</div>
			</div>

			<div>
				<label for="dob" class="block text-sm font-semibold text-slate-700 mb-1">Date of Birth</label>
				<div class="relative">
					<div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
						<Calendar class="w-4 h-4" />
					</div>
					<input
						id="dob"
						type="date"
						bind:value={dob}
						required
						class="w-full pl-9 pr-3 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 focus:border-indigo-600 text-sm"
					/>
				</div>
			</div>

			<div>
				<label for="password" class="block text-sm font-semibold text-slate-700 mb-1">Password</label>
				<div class="relative">
					<div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
						<Lock class="w-4 h-4" />
					</div>
					<input
						id="password"
						type="password"
						bind:value={password}
						required
						placeholder="••••••••"
						class="w-full pl-9 pr-3 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 focus:border-indigo-600 text-sm"
					/>
				</div>
			</div>

			<button
				type="submit"
				disabled={isLoading}
				class="w-full py-3 px-4 rounded-xl font-bold text-white bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 shadow-lg shadow-indigo-600/25 transition-all mt-4"
			>
				{isLoading ? 'Creating account...' : 'Create Account'}
			</button>
		</form>

		<div class="text-center pt-4 border-t border-slate-100">
			<p class="text-sm text-slate-600">
				Already registered?
				<a href="/login" class="font-semibold text-indigo-600 hover:text-indigo-700 ml-1">Sign In</a>
			</p>
		</div>
	</div>
</div>
