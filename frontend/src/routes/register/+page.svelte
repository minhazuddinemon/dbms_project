<!-- src/routes/register/+page.svelte -->
<script lang="ts">
	import { register } from '$lib/api/auth';
	import { goto } from '$app/navigation';
	import { UserPlus, User, Mail, Lock, Calendar, AlertCircle, CheckCircle2, ArrowRight, Eye, EyeOff, Phone, LifeBuoy, X, Send } from 'lucide-svelte';

	let firstName = $state('');
	let lastName = $state('');
	let email = $state('');
	let mobile = $state('');
	let password = $state('');
	let confirmPassword = $state('');
	let dob = $state('');

	let showPassword = $state(false);
	let showConfirmPassword = $state(false);
	let agreedTerms = $state(false);

	let isLoading = $state(false);
	let errorMessage = $state<string | null>(null);
	let successMessage = $state<string | null>(null);

	let showSupportModal = $state(false);
	let supportSubject = $state('');
	let supportMessage = $state('');
	let supportSent = $state(false);

	// Password strength calculation
	let strengthScore = $derived.by(() => {
		if (!password) return 0;
		let score = 0;
		if (password.length >= 8) score++;
		if (/[a-z]/.test(password) && /[A-Z]/.test(password)) score++;
		if (/\d/.test(password)) score++;
		if (/[^a-zA-Z0-9]/.test(password)) score++;
		return score;
	});

	let strengthText = $derived.by(() => {
		if (!password) return '';
		switch (strengthScore) {
			case 1:
				return 'Weak: Add numbers and symbols';
			case 2:
				return 'Moderate: Add uppercase and special characters';
			case 3:
				return 'Strong: Solid password';
			case 4:
				return 'Very Strong: Excellent password!';
			default:
				return 'Too short';
		}
	});

	let strengthColor = $derived.by(() => {
		switch (strengthScore) {
			case 1:
				return 'bg-red-500';
			case 2:
				return 'bg-amber-500';
			case 3:
				return 'bg-blue-500';
			case 4:
				return 'bg-emerald-500';
			default:
				return 'bg-outline-variant/30';
		}
	});

	let strengthTextColor = $derived.by(() => {
		switch (strengthScore) {
			case 1:
				return 'text-red-600';
			case 2:
				return 'text-amber-600';
			case 3:
				return 'text-blue-600';
			case 4:
				return 'text-emerald-600';
			default:
				return 'text-on-surface-variant';
		}
	});

	let passwordsMatch = $derived.by(() => {
		if (!confirmPassword) return true;
		return password === confirmPassword;
	});

	async function handleSubmit(e: Event) {
		e.preventDefault();
		if (!firstName || !lastName || !email || !password || !dob) {
			errorMessage = 'All required fields must be filled out.';
			return;
		}

		if (password !== confirmPassword) {
			errorMessage = 'Passwords do not match.';
			return;
		}

		if (!agreedTerms) {
			errorMessage = 'You must agree to the Terms of Service and Privacy Policy.';
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
	<title>Sign Up - UniApp Portal</title>
</svelte:head>

<div class="min-h-[85vh] flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8 bg-mesh">
	<div class="max-w-lg w-full space-y-6">
		<div class="glass-panel p-8 sm:p-10 rounded-[2.5rem] border border-outline-variant/40 shadow-2xl shadow-primary/10 animate-fade-in-up">
			
			<!-- Tab Selector -->
			<div class="flex items-center border-b border-outline-variant/30 pb-2">
				<a href="/login" class="flex-1 text-center py-3 font-semibold text-on-surface-variant hover:text-primary text-base transition-all">
					Sign In
				</a>
				<a href="/register" class="flex-1 text-center py-3 font-bold text-primary border-b-2 border-primary text-base transition-all">
					Sign Up
				</a>
			</div>

			<div class="text-center pt-2">
				<div class="w-12 h-12 rounded-2xl bg-gradient-to-tr from-primary to-primary-container text-white flex items-center justify-center mx-auto mb-3 shadow-lg shadow-primary/30">
					<UserPlus class="w-6 h-6" />
				</div>
				<h2 class="text-2xl font-extrabold text-on-surface">Join UniApp Portal</h2>
				<p class="mt-1 text-sm text-on-surface-variant">One account to check eligibility and apply to public universities</p>
			</div>

			{#if errorMessage}
				<div class="p-4 rounded-2xl bg-error-container/60 border border-error/30 flex items-start gap-3 text-on-error-container text-sm">
					<AlertCircle class="w-5 h-5 shrink-0 text-error mt-0.5" />
					<span>{errorMessage}</span>
				</div>
			{/if}

			{#if successMessage}
				<div class="p-4 rounded-2xl bg-tertiary-fixed/40 border border-tertiary/30 flex items-start gap-3 text-on-tertiary-fixed-variant text-sm">
					<CheckCircle2 class="w-5 h-5 shrink-0 text-tertiary mt-0.5" />
					<span>{successMessage}</span>
				</div>
			{/if}

			<form onsubmit={handleSubmit} class="space-y-4">
				<!-- Full Name Fields -->
				<div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
					<div class="space-y-1.5">
						<label for="firstName" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">First Name</label>
						<div class="relative">
							<div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-outline">
								<User class="w-4 h-4" />
							</div>
							<input
								id="firstName"
								type="text"
								bind:value={firstName}
								required
								placeholder="John"
								class="w-full pl-10 pr-3 py-2.5 rounded-xl border border-outline-variant/50 bg-white/90 text-on-surface focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all shadow-sm"
							/>
						</div>
					</div>

					<div class="space-y-1.5">
						<label for="lastName" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Last Name</label>
						<div class="relative">
							<div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-outline">
								<User class="w-4 h-4" />
							</div>
							<input
								id="lastName"
								type="text"
								bind:value={lastName}
								required
								placeholder="Doe"
								class="w-full pl-10 pr-3 py-2.5 rounded-xl border border-outline-variant/50 bg-white/90 text-on-surface focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all shadow-sm"
							/>
						</div>
					</div>
				</div>

				<!-- Email & Mobile -->
				<div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
					<div class="space-y-1.5">
						<label for="email" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Email</label>
						<div class="relative">
							<div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-outline">
								<Mail class="w-4 h-4" />
							</div>
							<input
								id="email"
								type="email"
								bind:value={email}
								required
								placeholder="john@example.com"
								class="w-full pl-10 pr-3 py-2.5 rounded-xl border border-outline-variant/50 bg-white/90 text-on-surface focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all shadow-sm"
							/>
						</div>
					</div>

					<div class="space-y-1.5">
						<label for="mobile" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Mobile</label>
						<div class="relative">
							<div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-outline">
								<Phone class="w-4 h-4" />
							</div>
							<input
								id="mobile"
								type="tel"
								bind:value={mobile}
								placeholder="+880 1xxx"
								class="w-full pl-10 pr-3 py-2.5 rounded-xl border border-outline-variant/50 bg-white/90 text-on-surface focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all shadow-sm"
							/>
						</div>
					</div>
				</div>

				<!-- Date of Birth -->
				<div class="space-y-1.5">
					<label for="dob" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Date of Birth</label>
					<div class="relative">
						<div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-outline">
							<Calendar class="w-4 h-4" />
						</div>
						<input
							id="dob"
							type="date"
							bind:value={dob}
							required
							class="w-full pl-10 pr-3 py-2.5 rounded-xl border border-outline-variant/50 bg-white/90 text-on-surface focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all shadow-sm"
						/>
					</div>
				</div>

				<!-- Password & Confirm Fields -->
				<div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
					<div class="space-y-1.5">
						<label for="password" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Password</label>
						<div class="relative">
							<div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-outline">
								<Lock class="w-4 h-4" />
							</div>
							<input
								id="password"
								type={showPassword ? 'text' : 'password'}
								bind:value={password}
								required
								placeholder="••••••••"
								class="w-full pl-10 pr-10 py-2.5 rounded-xl border border-outline-variant/50 bg-white/90 text-on-surface focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all shadow-sm"
							/>
							<button
								type="button"
								onclick={() => (showPassword = !showPassword)}
								class="absolute inset-y-0 right-0 pr-3 flex items-center text-outline hover:text-on-surface transition-colors"
								aria-label="Toggle password visibility"
							>
								{#if showPassword}
									<EyeOff class="w-4 h-4" />
								{:else}
									<Eye class="w-4 h-4" />
								{/if}
							</button>
						</div>
					</div>

					<div class="space-y-1.5">
						<label for="confirmPassword" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Confirm</label>
						<div class="relative">
							<div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-outline">
								<Lock class="w-4 h-4" />
							</div>
							<input
								id="confirmPassword"
								type={showConfirmPassword ? 'text' : 'password'}
								bind:value={confirmPassword}
								required
								placeholder="Confirm"
								class="w-full pl-10 pr-10 py-2.5 rounded-xl border border-outline-variant/50 bg-white/90 text-on-surface focus:ring-2 focus:ring-primary/40 focus:border-primary text-sm transition-all shadow-sm"
							/>
							<button
								type="button"
								onclick={() => (showConfirmPassword = !showConfirmPassword)}
								class="absolute inset-y-0 right-0 pr-3 flex items-center text-outline hover:text-on-surface transition-colors"
								aria-label="Toggle confirm password visibility"
							>
								{#if showConfirmPassword}
									<EyeOff class="w-4 h-4" />
								{:else}
									<Eye class="w-4 h-4" />
								{/if}
							</button>
						</div>
					</div>
				</div>

				<!-- Password Strength Indicator -->
				{#if password}
					<div class="space-y-1.5 pt-1">
						<div class="grid grid-cols-4 gap-1.5 h-1.5 rounded-full overflow-hidden bg-outline-variant/30">
							<div class="h-full rounded-full transition-all duration-300 {strengthScore >= 1 ? strengthColor : 'bg-transparent'}"></div>
							<div class="h-full rounded-full transition-all duration-300 {strengthScore >= 2 ? strengthColor : 'bg-transparent'}"></div>
							<div class="h-full rounded-full transition-all duration-300 {strengthScore >= 3 ? strengthColor : 'bg-transparent'}"></div>
							<div class="h-full rounded-full transition-all duration-300 {strengthScore >= 4 ? strengthColor : 'bg-transparent'}"></div>
						</div>
						<div class="flex items-center justify-between text-xs">
							<span class="font-medium {strengthTextColor}">{strengthText}</span>
							{#if confirmPassword && !passwordsMatch}
								<span class="text-red-600 font-medium">Passwords do not match</span>
							{/if}
						</div>
					</div>
				{/if}

				<!-- Terms and Conditions Checkbox -->
				<div class="pt-2">
					<label class="flex items-center gap-2.5 text-xs text-on-surface-variant cursor-pointer select-none">
						<input
							type="checkbox"
							bind:checked={agreedTerms}
							required
							class="rounded border-outline-variant text-primary focus:ring-primary/40 w-4 h-4"
						/>
						<span>
							I agree to the <a href="/helpdesk" class="text-primary hover:underline font-semibold">Terms of Service</a> and <a href="/helpdesk" class="text-primary hover:underline font-semibold">Privacy Policy</a>
						</span>
					</label>
				</div>

				<button
					type="submit"
					disabled={isLoading || !agreedTerms || (confirmPassword !== '' && !passwordsMatch)}
					class="w-full py-3.5 px-6 rounded-xl font-bold text-white bg-primary hover:bg-primary-container disabled:opacity-50 shadow-lg shadow-primary/25 hover:shadow-primary/40 hover:-translate-y-0.5 transition-all duration-200 flex items-center justify-center gap-2 text-base mt-4"
				>
					{#if isLoading}
						<span>Creating Account...</span>
					{:else}
						<span>Create Account</span>
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
					<p class="text-xs text-on-surface-variant">We're here to help with your registration</p>
				</div>
			</div>

			{#if supportSent}
				<div class="p-4 rounded-xl bg-tertiary-fixed/40 text-on-tertiary-fixed-variant text-sm font-medium text-center py-6">
					Support request submitted successfully! We will contact you via email shortly.
				</div>
			{:else}
				<form onsubmit={handleSupportSubmit} class="space-y-4">
					<div>
						<label for="supSubjectReg" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant mb-1">Subject</label>
						<input
							id="supSubjectReg"
							type="text"
							bind:value={supportSubject}
							required
							placeholder="e.g. Issue during account creation"
							class="w-full px-3.5 py-2.5 rounded-xl border border-outline-variant/50 text-sm focus:ring-2 focus:ring-primary/40 outline-none"
						/>
					</div>

					<div>
						<label for="supMessageReg" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant mb-1">Message</label>
						<textarea
							id="supMessageReg"
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

