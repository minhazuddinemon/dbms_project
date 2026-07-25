<!-- src/routes/helpdesk/+page.svelte -->
<script lang="ts">
	import { HelpCircle, Mail, MessageSquare, Send, CheckCircle2 } from 'lucide-svelte';

	let subject = $state('');
	let message = $state('');
	let isSubmitted = $state(false);

	function handleSubmit(e: Event) {
		e.preventDefault();
		if (!subject || !message) return;
		isSubmitted = true;
	}
</script>

<svelte:head>
	<title>Helpdesk & Support - UniApp</title>
</svelte:head>

<div class="py-10 bg-slate-50 min-h-screen">
	<div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
		<div>
			<h1 class="text-3xl font-extrabold text-slate-900 sm:text-4xl">Helpdesk & Support</h1>
			<p class="mt-2 text-slate-600">Have questions about university eligibility, payment status, or unit details? We are here to help.</p>
		</div>

		{#if isSubmitted}
			<div class="bg-white rounded-2xl border border-emerald-200 p-8 text-center space-y-4 shadow-sm">
				<CheckCircle2 class="w-12 h-12 text-emerald-500 mx-auto" />
				<h3 class="text-xl font-bold text-slate-900">Message Received</h3>
				<p class="text-slate-600 text-sm">Thank you! Our support team will respond to your query within 24 hours.</p>
				<button onclick={() => isSubmitted = false} class="px-6 py-2.5 rounded-xl font-semibold bg-indigo-600 text-white text-sm">
					Submit Another Ticket
				</button>
			</div>
		{:else}
			<form onsubmit={handleSubmit} class="bg-white rounded-2xl border border-slate-200 p-8 shadow-sm space-y-6">
				<h3 class="text-xl font-bold text-slate-900 border-b border-slate-100 pb-4 flex items-center gap-2">
					<MessageSquare class="w-5 h-5 text-indigo-600" />
					Submit a Support Ticket
				</h3>

				<div class="space-y-4">
					<div>
						<label for="subject" class="block text-sm font-semibold text-slate-700 mb-1">Topic / Subject</label>
						<input
							id="subject"
							type="text"
							bind:value={subject}
							required
							placeholder="e.g. Question regarding HSC physics mark requirement"
							class="w-full px-4 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 text-sm"
						/>
					</div>

					<div>
						<label for="message" class="block text-sm font-semibold text-slate-700 mb-1">Message</label>
						<textarea
							id="message"
							bind:value={message}
							required
							rows="5"
							placeholder="Describe your issue or query in detail..."
							class="w-full px-4 py-2.5 rounded-xl border border-slate-300 focus:ring-2 focus:ring-indigo-600 text-sm"
						></textarea>
					</div>
				</div>

				<button
					type="submit"
					class="w-full py-3 px-4 rounded-xl font-bold text-white bg-indigo-600 hover:bg-indigo-700 shadow-lg shadow-indigo-600/25 transition-all text-sm flex items-center justify-center gap-2"
				>
					<Send class="w-4 h-4" />
					Send Message
				</button>
			</form>
		{/if}
	</div>
</div>
