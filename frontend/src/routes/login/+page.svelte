<!-- src/routes/login/+page.svelte -->
<script lang="ts">
	import { authState } from '$lib/state/auth.svelte';
	import { goto } from '$app/navigation';
	import { LogIn, Mail, Lock, AlertCircle, ArrowRight } from 'lucide-svelte';

	let email = $state('');
	let password = $state('');
	let errorMessage = $state<string | null>(null);

	async function handleSubmit(e: Event) {
		e.preventDefault();
		if (!email || !password) {
			errorMessage = 'Please enter both email and password.';
			return;
		}

		errorMessage = null;
		const success = await authState.login({ email, password });
		if (success) {
			goto('/profile');
		} else {
			errorMessage = authState.error || 'Invalid credentials. Please check your email and password.';
		}
	}
</script>

<svelte:head>
	<title>Sign In - UniApp</title>
</svelte:head>

<div class="min-h-[80vh] flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8 bg-slate-50">
	<div class="max-w-md w-full space-y-8 bg-white p-8 rounded-2xl border border-slate-200 shadow-xl shadow-slate-200/50">
		<div class="text-center">
			<div class="w-12 h-12 rounded-2xl bg-indigo-600 text-white flex items-center justify-center mx-auto mb-4 shadow-lg shadow-indigo-600/30">
				<LogIn class="w-6 h-6" />
			</div>
			<h2 class="text-2xl font-extrabold text-slate-900">Sign in to your account</h2>
			<p class="mt-2 text-sm text-slate-600">Access your university applications and academic profile</p>
		</div>

		{#if errorMessage}
			<div class="p-4 rounded-xl bg-red-50 border border-red-200 flex items-start gap-3 text-red-700 text-sm">
				<AlertCircle class="w-5 h-5 shrink-0 text-red-500 mt-0.5" />
				<span>{errorMessage}</span>
			</div>
		{/if}

		<form onsubmit={handleSubmit} class="space-y-5">
			<div>
				<label for="email" class="block text-sm font-semibold text-slate-700 mb-1">Email Address</label>
				<div class="relative">
					<div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
						<Mail class="w-5 h-5" />
					</div>
					<input
						id="email"
						type="email"
						bind:value={email}
						required
						placeholder="student@example.com"
						class="w-full pl-10 pr-4 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 focus:border-indigo-600 text-sm transition-all"
					/>
				</div>
			</div>

			<div>
				<label for="password" class="block text-sm font-semibold text-slate-700 mb-1">Password</label>
				<div class="relative">
					<div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
						<Lock class="w-5 h-5" />
					</div>
					<input
						id="password"
						type="password"
						bind:value={password}
						required
						placeholder="••••••••"
						class="w-full pl-10 pr-4 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 focus:border-indigo-600 text-sm transition-all"
					/>
				</div>
			</div>

			<button
				type="submit"
				disabled={authState.isLoading}
				class="w-full py-3 px-4 rounded-xl font-bold text-white bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 shadow-lg shadow-indigo-600/25 hover:shadow-indigo-600/35 transition-all flex items-center justify-center gap-2"
			>
				{#if authState.isLoading}
					<span>Signing in...</span>
				{:else}
					<span>Sign In</span>
					<ArrowRight class="w-4 h-4" />
				{/if}
			</button>
		</form>

		<div class="text-center pt-4 border-t border-slate-100">
			<p class="text-sm text-slate-600">
				Don't have an account?
				<a href="/register" class="font-semibold text-indigo-600 hover:text-indigo-700 ml-1">Register now</a>
			</p>
		</div>
	</div>
</div>
