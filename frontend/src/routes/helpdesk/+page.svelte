<!-- src/routes/helpdesk/+page.svelte -->
<script lang="ts">
	import { HelpCircle, Mail, MessageSquare, Send, CheckCircle2, ChevronDown, ShieldCheck, PhoneCall } from 'lucide-svelte';

	let subject = $state('');
	let message = $state('');
	let isSubmitted = $state(false);

	let openFaq = $state<number | null>(0);

	const faqs = [
		{
			q: 'How does the Eligibility Engine calculate my qualification?',
			a: 'The rule engine evaluates your SSC and HSC overall GPA as well as individual HSC subject marks (Physics, Mathematics, Chemistry) against the minimum cutmark criteria published by each university unit.'
		},
		{
			q: 'What happens if a program application reports missing profile fields?',
			a: 'If a specific university program requires mandatory fields (such as Father Name, Quota Status, or Permanent Address), the API returns a list of missing parameters. Simply navigate to your Academic Profile and fill out those details.'
		},
		{
			q: 'How do I pay application fees via BKash or Nagad?',
			a: 'Go to the Apply / Payments page, select your preferred mobile banking method, send the fee to the merchant number provided, and enter your 8-digit Transaction ID (TrxID) for instant automated verification.'
		}
	];

	function handleSubmit(e: Event) {
		e.preventDefault();
		if (!subject || !message) return;
		isSubmitted = true;
	}
</script>

<svelte:head>
	<title>Support & Helpdesk - UniApp Portal</title>
</svelte:head>

<div class="py-10 bg-mesh min-h-screen">
	<div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8 animate-fade-in-up">
		
		<!-- Banner Header -->
		<div class="bg-gradient-to-r from-primary via-primary-container to-secondary text-white rounded-[2.5rem] p-8 sm:p-10 shadow-2xl shadow-primary/20 relative overflow-hidden">
			<div class="absolute right-0 top-0 w-96 h-96 bg-tertiary-fixed/20 rounded-full blur-3xl pointer-events-none"></div>

			<div class="relative z-10 max-w-2xl space-y-3">
				<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 border border-white/20 text-tertiary-fixed text-xs font-bold uppercase tracking-wider backdrop-blur-md">
					<HelpCircle class="w-4 h-4" />
					24/7 Applicant Assistance
				</div>

				<h1 class="text-3xl sm:text-4xl font-black text-white">
					Support & Helpdesk
				</h1>

				<p class="text-slate-100 text-base leading-relaxed">
					Have questions about university eligibility, payment status, or subject cutmarks? We are here to help.
				</p>
			</div>
		</div>

		<!-- Submit Ticket Form & FAQs -->
		<div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
			<!-- Submit Ticket -->
			<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
				<h3 class="text-2xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-4 flex items-center gap-2">
					<MessageSquare class="w-6 h-6 text-primary" />
					Submit Support Ticket
				</h3>

				{#if isSubmitted}
					<div class="p-6 rounded-2xl bg-tertiary-fixed/30 border border-tertiary/40 text-center space-y-3">
						<CheckCircle2 class="w-12 h-12 text-tertiary mx-auto" />
						<h4 class="text-xl font-bold text-on-surface">Ticket Received!</h4>
						<p class="text-xs text-on-surface-variant">Our support team will respond to your registered email within 24 hours.</p>
						<button onclick={() => isSubmitted = false} class="px-6 py-2.5 rounded-xl font-bold bg-primary text-white text-xs">
							Submit Another Query
						</button>
					</div>
				{:else}
					<form onsubmit={handleSubmit} class="space-y-4">
						<div class="space-y-1.5">
							<label for="subject" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Subject / Topic</label>
							<input
								id="subject"
								type="text"
								bind:value={subject}
								required
								placeholder="e.g. Question regarding HSC physics mark requirement"
								class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
							/>
						</div>

						<div class="space-y-1.5">
							<label for="message" class="block text-xs font-bold uppercase tracking-wider text-on-surface-variant">Detailed Query</label>
							<textarea
								id="message"
								bind:value={message}
								required
								rows="4"
								placeholder="Describe your issue or question in detail..."
								class="w-full px-4 py-3 rounded-xl border border-outline-variant/50 focus:ring-2 focus:ring-primary/40 text-sm bg-white"
							></textarea>
						</div>

						<button
							type="submit"
							class="w-full py-3.5 px-6 rounded-xl font-bold text-white bg-primary hover:bg-primary-container shadow-lg shadow-primary/25 hover:shadow-primary/40 transition-all flex items-center justify-center gap-2 text-sm"
						>
							<Send class="w-4 h-4" />
							Send Support Ticket
						</button>
					</form>
				{/if}
			</div>

			<!-- FAQ Accordion -->
			<div class="glass-panel p-8 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-xl space-y-6">
				<h3 class="text-2xl font-extrabold text-on-surface border-b border-outline-variant/30 pb-4">
					Frequently Asked Questions
				</h3>

				<div class="space-y-3">
					{#each faqs as item, index}
						<div class="rounded-2xl border border-outline-variant/30 overflow-hidden transition-all bg-surface-container-low/50">
							<button
								onclick={() => openFaq = openFaq === index ? null : index}
								class="w-full p-5 text-left font-bold text-on-surface flex items-center justify-between gap-4 text-sm"
							>
								<span>{item.q}</span>
								<ChevronDown class="w-4 h-4 shrink-0 text-primary transition-transform duration-200 {openFaq === index ? 'rotate-180' : ''}" />
							</button>

							{#if openFaq === index}
								<div class="px-5 pb-5 text-xs text-on-surface-variant leading-relaxed border-t border-outline-variant/20 pt-3 bg-white/80">
									{item.a}
								</div>
							{/if}
						</div>
					{/each}
				</div>
			</div>
		</div>
	</div>
</div>
