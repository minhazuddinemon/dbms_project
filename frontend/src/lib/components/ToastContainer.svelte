<!-- src/lib/components/ToastContainer.svelte -->
<script lang="ts">
	import { toastState } from '$lib/state/toast.svelte';
	import { CheckCircle2, AlertCircle, AlertTriangle, Info, X } from 'lucide-svelte';
</script>

<div
	class="fixed bottom-5 right-5 z-50 flex flex-col gap-2.5 max-w-sm w-full pointer-events-none px-4 sm:px-0"
	aria-live="polite"
>
	{#each toastState.toasts as toast (toast.id)}
		<div
			class="pointer-events-auto flex items-center justify-between gap-3 p-3.5 rounded-xl text-sm font-medium border shadow-lg backdrop-blur-md transition-all duration-300 transform translate-y-0
			{toast.type === 'success'
				? 'bg-white/95 dark:bg-neutral-900/95 border-emerald-500/30 text-neutral-900 dark:text-neutral-100'
				: ''}
			{toast.type === 'error'
				? 'bg-white/95 dark:bg-neutral-900/95 border-rose-500/30 text-neutral-900 dark:text-neutral-100'
				: ''}
			{toast.type === 'warning'
				? 'bg-white/95 dark:bg-neutral-900/95 border-amber-500/30 text-neutral-900 dark:text-neutral-100'
				: ''}
			{toast.type === 'info'
				? 'bg-white/95 dark:bg-neutral-900/95 border-sky-500/30 text-neutral-900 dark:text-neutral-100'
				: ''}"
		>
			<div class="flex items-center gap-3 min-w-0 flex-1">
				{#if toast.type === 'success'}
					<CheckCircle2 class="w-5 h-5 text-emerald-500 shrink-0" />
				{:else if toast.type === 'error'}
					<AlertCircle class="w-5 h-5 text-rose-500 shrink-0" />
				{:else if toast.type === 'warning'}
					<AlertTriangle class="w-5 h-5 text-amber-500 shrink-0" />
				{:else}
					<Info class="w-5 h-5 text-sky-500 shrink-0" />
				{/if}

				<span class="leading-snug font-medium text-slate-800 dark:text-slate-100">
					{toast.message}
				</span>
			</div>

			<button
				onclick={() => toastState.remove(toast.id)}
				class="text-neutral-400 hover:text-neutral-600 dark:hover:text-neutral-200 transition-colors p-1 rounded-md shrink-0 focus:outline-none"
				aria-label="Dismiss notification"
			>
				<X class="w-4 h-4" />
			</button>
		</div>
	{/each}
</div>