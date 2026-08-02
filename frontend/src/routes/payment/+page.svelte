<!-- src/routes/payment/+page.svelte -->
<script lang="ts">
	import { onMount } from 'svelte';
	import { fetchStudentApplications } from '$lib/api/applications';
	import { processPayment } from '$lib/api/payments';
	import { authState } from '$lib/state/auth.svelte';
	import type { StudentApplication } from '$lib/types/models';
	import { CreditCard, CheckCircle2, AlertCircle, ShieldCheck, ArrowRight, DollarSign, Wallet, RefreshCw } from 'lucide-svelte';

	import { page } from '$app/state';

	let selectedMethod = $state<'bKash' | 'Nagad' | 'Card'>('bKash');
	let txId = $state('');
	let selectedAppId = $state<number | null>(null);
	let isProcessing = $state(false);
	let paymentSuccess = $state(false);
	let errorMessage = $state<string | null>(null);

	let applications = $state<StudentApplication[]>([]);
	let pendingApplications = $derived(applications.filter(a => a.status !== 'PAID' && a.status !== 'APPROVED'));
	let selectedApp = $derived(pendingApplications.find(a => a.app_id === selectedAppId) || null);
	let selectedAppFee = $derived(selectedApp?.program_fee || selectedApp?.application_fee);

	async function reloadApplications() {
		if (!authState.isAuthenticated) return;
		try {
			applications = await fetchStudentApplications();
			if (pendingApplications.length > 0) {
				const urlAppId = Number(page.url.searchParams.get('app_id'));
				if (urlAppId && pendingApplications.some(a => a.app_id === urlAppId)) {
					selectedAppId = urlAppId;
				} else if (!selectedAppId || !pendingApplications.some(a => a.app_id === selectedAppId)) {
					selectedAppId = pendingApplications[0].app_id;
				}
			} else {
				selectedAppId = null;
			}
		} catch (err) {
			console.error(err);
		}
	}

	onMount(() => {
		reloadApplications();
	});

	async function handlePayment(e: Event) {
		e.preventDefault();
		if (!txId.trim() || !selectedAppId) return;

		isProcessing = true;
		errorMessage = null;

		try {
			await processPayment({
				application_id: selectedAppId,
				amount: String(selectedAppFee),
				payment_method: selectedMethod,
				transaction_id: txId
			});
			paymentSuccess = true;
			// Immediately reload student applications so pending fee state updates live
			await reloadApplications();
		} catch (err: any) {
			errorMessage = err?.message || 'Payment processing failed.';
		} finally {
			isProcessing = false;
		}
	}
</script>

<svelte:head>
	<title>Apply & Payments - UniApp</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		
		<!-- Header -->
		<div class="bg-gradient-to-r from-primary via-secondary to-primary-container text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-3">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<CreditCard class="w-4 h-4" />
					Unified Payment Gateway
				</div>

				<h1 class="text-3xl sm:text-4xl font-black text-white">
					Application Fees & Payment
				</h1>

				<p class="text-slate-100 text-base leading-relaxed">
					Pay application fees securely using bKash, Nagad, or Debit/Credit card with instant transaction verification.
				</p>
			</div>
		</div>

		<div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
			<!-- Applications Summary -->
			<div class="lg:col-span-2 space-y-6">
				<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
					<div class="flex items-center justify-between border-b border-outline-variant/30 pb-4">
						<h3 class="text-2xl font-extrabold text-on-surface">
							Pending Applications Fee ({pendingApplications.length})
						</h3>
						<button
							type="button"
							onclick={reloadApplications}
							class="p-2 rounded-xl text-primary hover:bg-primary-fixed/40 transition-colors flex items-center gap-1.5 text-xs font-bold cursor-pointer"
						>
							<RefreshCw class="w-4 h-4" />
							Refresh List
						</button>
					</div>

					{#if pendingApplications.length === 0}
						<div class="text-center py-10 text-on-surface-variant space-y-2">
							<CheckCircle2 class="w-14 h-14 text-emerald-500 mx-auto" />
							<p class="font-extrabold text-xl text-on-surface">No Pending Application Fees!</p>
							<p class="text-xs max-w-md mx-auto">All your submitted applications have been paid and confirmed.</p>
						</div>
					{:else}
						<div class="space-y-4">
							{#each pendingApplications as app}
								<button
									type="button"
									onclick={() => selectedAppId = app.app_id}
									class="w-full text-left p-6 rounded-2xl border transition-all flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4 cursor-pointer {selectedAppId === app.app_id ? 'border-primary bg-primary-fixed/30 shadow-md' : 'bg-surface-container-low/70 border-outline-variant/30 hover:bg-surface-container-low'}"
								>
									<div class="space-y-1">
										<span class="px-2.5 py-0.5 bg-primary-fixed text-on-primary-fixed text-xs font-bold rounded-lg uppercase">App #{app.app_id}</span>
										<h4 class="text-xl font-bold text-on-surface">{app.program_name || 'Degree Program'}</h4>
										<p class="text-sm font-semibold text-on-surface-variant">{app.university_name || 'University'}</p>
									</div>
									<div class="text-right">
										<span class="block text-2xl font-black text-primary">BDT {app.program_fee || app.application_fee || '500.00'}</span>
										<span class="text-xs font-bold text-amber-600 bg-amber-50 px-2.5 py-0.5 rounded-full uppercase tracking-wider">{app.status}</span>
									</div>
								</button>
							{/each}
						</div>
					{/if}
				</div>
			</div>

			<!-- Payment Options Box -->
			<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
				<h3 class="text-xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-4">Select Payment Method</h3>

				{#if errorMessage}
					<div class="p-4 rounded-2xl bg-error-container/60 border border-error/30 text-on-error-container text-xs font-bold">
						{errorMessage}
					</div>
				{/if}

				{#if paymentSuccess}
					<div class="p-6 rounded-2xl bg-tertiary-fixed/30 border border-tertiary/40 text-center space-y-4 animate-fade-in">
						<CheckCircle2 class="w-12 h-12 text-tertiary mx-auto" />
						<div class="space-y-1">
							<h4 class="text-xl font-bold text-on-surface">Payment Verified & Updated!</h4>
							<p class="text-xs text-on-surface-variant">Transaction TrxID <strong class="font-mono uppercase text-on-surface">{txId}</strong> confirmed.</p>
						</div>

						<div class="pt-2 flex flex-col gap-2">
							{#if pendingApplications.length > 0}
								<button
									type="button"
									onclick={() => { paymentSuccess = false; txId = ''; }}
									class="w-full py-2.5 px-4 rounded-xl font-bold text-white bg-primary text-xs shadow-md hover:bg-primary-container transition-all cursor-pointer"
								>
									Pay Another Pending Fee ({pendingApplications.length} remaining)
								</button>
							{/if}
							<a
								href="/dashboard"
								class="w-full py-2.5 px-4 rounded-xl font-bold border border-outline-variant text-on-surface text-xs hover:bg-surface-container text-center transition-all"
							>
								View Dashboard & Applications
							</a>
						</div>
					</div>
				{:else}
					<form onsubmit={handlePayment} class="space-y-5">
						<div class="grid grid-cols-3 gap-3">
							<button
								type="button"
								onclick={() => selectedMethod = 'bKash'}
								class="p-4 rounded-2xl border text-center font-bold text-sm transition-all cursor-pointer {selectedMethod === 'bKash' ? 'border-primary bg-primary-fixed/40 text-primary shadow-md' : 'border-outline-variant/40 text-on-surface-variant hover:bg-surface-container'}"
							>
								bKash
							</button>
							<button
								type="button"
								onclick={() => selectedMethod = 'Nagad'}
								class="p-4 rounded-2xl border text-center font-bold text-sm transition-all cursor-pointer {selectedMethod === 'Nagad' ? 'border-primary bg-primary-fixed/40 text-primary shadow-md' : 'border-outline-variant/40 text-on-surface-variant hover:bg-surface-container'}"
							>
								Nagad
							</button>
							<button
								type="button"
								onclick={() => selectedMethod = 'Card'}
								class="p-4 rounded-2xl border text-center font-bold text-sm transition-all cursor-pointer {selectedMethod === 'Card' ? 'border-primary bg-primary-fixed/40 text-primary shadow-md' : 'border-outline-variant/40 text-on-surface-variant hover:bg-surface-container'}"
							>
								Card
							</button>
						</div>

						<div class="space-y-1.5">
							<label for="txId" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Transaction ID (TrxID)</label>
							<input
								id="txId"
								type="text"
								bind:value={txId}
								required
								placeholder="e.g. TRX987654321"
								class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm font-mono uppercase bg-white"
							/>
						</div>

						<button
							type="submit"
							disabled={isProcessing || !selectedAppId}
							class="w-full py-3.5 px-6 rounded-xl font-bold text-white bg-primary hover:bg-primary-container disabled:opacity-50 shadow-lg shadow-primary/25 hover:shadow-primary/40 transition-all flex items-center justify-center gap-2 text-sm cursor-pointer"
						>
							{#if isProcessing}
								<span>Verifying Payment...</span>
							{:else}
								<span>Submit Payment BDT {selectedAppFee}</span>
								<ArrowRight class="w-4 h-4" />
							{/if}
						</button>
					</form>
				{/if}
			</div>
		</div>
	</div>
</div>
