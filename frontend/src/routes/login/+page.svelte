<!-- src/routes/login/+page.svelte -->
<script lang="ts">
	import { authState } from '$lib/state/auth.svelte';
	import { goto } from '$app/navigation';
	import { LogIn, Mail, Lock, Eye, EyeOff, AlertCircle, ArrowRight, GraduationCap, LifeBuoy, X, Send } from 'lucide-svelte';

	let email = $state('');
	let password = $state('');
	let showPassword = $state(false);
	let errorMessage = $state<string | null>(null);
	let showSupportModal = $state(false);
	let supportSubject = $state('');
	let supportMessage = $state('');
	let supportSent = $state(false);

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
			errorMessage = authState.error || 'Invalid email or password. Please check your credentials.';
		}
	}

	function handleSupportSubmit(e: Event) {
		e.preventDefault();
		if (supportSubject && supportMessage) {
			supportSent = true;
			setTimeout(() => {
				supportSent = false;
				showSupportModal = false;
				supportSubject = '';
				supportMessage = '';
			}, 2000);
		}
	}
</script>

<svelte:head>
	<title>Sign In - UniApp Portal</title>
</svelte:head>

<div class="min-h-[85vh] flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8 bg-mesh">
	<div class="max-w-md w-full space-y-6">
		<div class="glass-panel p-8 sm:p-10 rounded-[2.5rem] border border-outline-variant/40 shadow-2xl shadow-primary/10 animate-fade-in-up">
			
			<!-- Tab Selector -->
			<div class="flex items-center border-b border-outline-variant/30 pb-2">
				<a href="/login" class="flex-1 text-center py-3 font-bold text-primary border-b-2 border-primary text-base transition-all">
					Sign In
				</a>
				<a href="/register" class="flex-1 text-center py-3 font-semibold text-on-surface-variant hover:text-primary text-base transition-all">
					Sign Up
				</a>
			</div>

			<div class="text-center pt-2">
				<div class="w-12 h-12 rounded-2xl bg-gradient-to-tr from-primary to-primary-container text-white flex items-center justify-center mx-auto mb-3 shadow-lg shadow-primary/30">
					<GraduationCap class="w-6 h-6" />
				</div>
				<h2 class="text-2xl font-extrabold text-on-surface">Welcome Back</h2>
				<p class="mt-1 text-sm text-on-surface-variant">Sign in to manage your university applications</p>
			</div>

			{#if errorMessage}
				<div class="p-4 rounded-2xl bg-error-container/60 border border-error/30 flex items-start gap-3 text-on-error-container text-sm">
					<AlertCircle class="w-5 h-5 shrink-0 text-error mt-0.5" />
					<span>{errorMessage}</span>
				</div>
			{/if}

			<form onsubmit={handleSubmit} class="space-y-5">
				<div class="space-y-1.5">
					<label for="email" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Email Address</label>
					<div class="relative">
						<div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-outline">
							<Mail class="w-5 h-5" />
						</div>
						<input
							id="email"
							type="email"
							bind:value={email}
							required
							placeholder="student@example.com"
							class="w-full pl-11 pr-4 py-3 rounded-xl border border-outline-variant/50 bg-white/90 text-on-surface focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all shadow-sm"
						/>
					</div>
				</div>

				<div class="space-y-1.5">
					<div class="flex justify-between items-center">
						<label for="password" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Password</label>
					</div>
					<div class="relative">
						<div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-outline">
							<Lock class="w-5 h-5" />
						</div>
						<input
							id="password"
							type={showPassword ? 'text' : 'password'}
							bind:value={password}
							required
							placeholder="••••••••"
							class="w-full pl-11 pr-11 py-3 rounded-xl border border-outline-variant/50 bg-white/90 text-on-surface focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all shadow-sm"
						/>
						<button
							type="button"
							onclick={() => (showPassword = !showPassword)}
							class="absolute inset-y-0 right-0 pr-3.5 flex items-center text-outline hover:text-on-surface transition-colors"
							aria-label="Toggle password visibility"
						>
							{#if showPassword}
								<EyeOff class="w-5 h-5" />
							{:else}
								<Eye class="w-5 h-5" />
							{/if}
						</button>
					</div>
				</div>

				<button
					type="submit"
					disabled={authState.isLoading}
					class="w-full py-3.5 px-6 rounded-xl font-bold text-white bg-primary hover:bg-primary-container disabled:opacity-50 shadow-lg shadow-primary/25 hover:shadow-primary/40 hover:-translate-y-0.5 transition-all duration-200 flex items-center justify-center gap-2 text-base mt-2"
				>
					{#if authState.isLoading}
						<span>Signing in...</span>
					{:else}
						<span>Sign In</span>
						<ArrowRight class="w-5 h-5" />
					{/if}
				</button>
			</form>
		</div>

		<!-- Footer Contact Support Link -->
		<p class="text-center text-sm text-on-surface-variant pt-1">
			Having trouble?
			<button
				type="button"
				onclick={() => (showSupportModal = true)}
				class="font-bold text-primary hover:underline transition-colors inline-flex items-center gap-1"
			>
				<LifeBuoy class="w-4 h-4 inline" />
				Contact Support
			</button>
		</p>
	</div>
</div>

<!-- Support Modal -->
{#if showSupportModal}
	<div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/40 backdrop-blur-sm animate-fade-in">
		<div class="max-w-md w-full bg-white rounded-3xl p-6 sm:p-8 shadow-2xl border border-outline-variant/40 relative">
			<button
				type="button"
				onclick={() => (showSupportModal = false)}
				class="absolute top-5 right-5 text-outline hover:text-on-surface p-1 rounded-full hover:bg-surface-container transition-colors"
			>
				<X class="w-5 h-5" />
			</button>

			<div class="flex items-center gap-3 mb-4">
				<div class="w-10 h-10 rounded-xl bg-primary/10 text-primary flex items-center justify-center">
					<LifeBuoy class="w-5 h-5" />
				</div>
				<div>
					<h3 class="text-lg font-bold text-on-surface">Contact Support</h3>
					<p class="text-xs text-on-surface-variant">We're here to help with your admission account</p>
				</div>
			</div>

			{#if supportSent}
				<div class="p-4 rounded-xl bg-tertiary-fixed/40 text-on-tertiary-fixed-variant text-sm font-medium text-center py-6">
					Support request submitted successfully! We will contact you via email shortly.
				</div>
			{:else}
				<form onsubmit={handleSupportSubmit} class="space-y-4">
					<div>
						<label for="supSubject" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant mb-1">Subject</label>
						<input
							id="supSubject"
							type="text"
							bind:value={supportSubject}
							required
							placeholder="e.g. Cannot log in to account"
							class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm focus:ring-2 focus:ring-primary/40 outline-none"
						/>
					</div>

					<div>
						<label for="supMessage" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant mb-1">Message</label>
						<textarea
							id="supMessage"
							rows="4"
							bind:value={supportMessage}
							required
							placeholder="Describe the issue you are experiencing..."
							class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm focus:ring-2 focus:ring-primary/40 outline-none resize-none"
						></textarea>
					</div>

					<div class="flex gap-3 pt-2">
						<button
							type="button"
							onclick={() => (showSupportModal = false)}
							class="flex-1 py-2.5 px-4 rounded-xl font-semibold border border-outline-variant text-on-surface-variant hover:bg-surface-container text-sm transition-colors"
						>
							Cancel
						</button>
						<button
							type="submit"
							class="flex-1 py-2.5 px-4 rounded-xl font-bold text-white bg-primary hover:bg-primary-container text-sm flex items-center justify-center gap-2 shadow-md transition-colors"
						>
							<span>Send Request</span>
							<Send class="w-4 h-4" />
						</button>
					</div>
				</form>
			{/if}
		</div>
	</div>
{/if}

