<!-- src/lib/components/UniversityCard.svelte -->
<script lang="ts">
	import type { University } from '$lib/types/models';
	import { Building2, MapPin, Globe, ExternalLink, ArrowRight, Edit3, Trash2 } from 'lucide-svelte';

	interface Props {
		university: University;
		isAdmin?: boolean;
		onedit?: (u: University) => void;
		ondelete?: (uId: number) => void;
	}

	let { university, isAdmin = false, onedit, ondelete }: Props = $props();
</script>

<div class="glass-panel rounded-[2.5rem] border border-outline-variant/40 p-7 shadow-sm hover:shadow-xl hover:border-primary/40 transition-all duration-300 flex flex-col justify-between space-y-6 bg-white/95 card-hover group">
	<div class="space-y-5">
		<!-- Header with Logo, Name, Location & Actions -->
		<div class="flex items-center justify-between gap-2.5 border-b border-outline-variant/20 pb-4">
			<div class="flex items-center gap-3 min-w-0 flex-1">
				{#if university.logo_url}
					<img src={university.logo_url} alt={university.u_name} class="w-12 h-12 rounded-2xl object-contain p-1 border border-outline-variant/30 shadow-sm bg-white shrink-0" />
				{:else}
					<div class="w-12 h-12 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center font-black text-lg border border-primary/20 shrink-0">
						{university.u_name.charAt(0)}
					</div>
				{/if}
				<div class="min-w-0 flex-1">
					<h3 class="text-base sm:text-lg font-black text-on-surface leading-snug group-hover:text-primary transition-colors line-clamp-2" title={university.u_name}>
						{university.u_name}
					</h3>
					{#if university.location}
						<p class="flex items-center gap-1 text-[11px] font-bold text-on-surface-variant mt-0.5 truncate">
							<MapPin class="w-3 h-3 text-tertiary shrink-0" />
							<span>{university.location}</span>
						</p>
					{/if}
				</div>
			</div>

			<!-- Controls: Admin Edit/Delete OR Client Website Link -->
			{#if isAdmin}
				<div class="flex items-center gap-1 shrink-0">
					{#if onedit}
						<button
							type="button"
							onclick={() => onedit?.(university)}
							class="p-1.5 rounded-xl text-primary hover:bg-primary-fixed/40 transition-colors cursor-pointer"
							title="Edit University"
						>
							<Edit3 class="w-4 h-4" />
						</button>
					{/if}
					{#if ondelete}
						<button
							type="button"
							onclick={() => ondelete?.(university.u_id)}
							class="p-1.5 rounded-lg text-error hover:bg-error-container/40 transition-colors cursor-pointer"
							title="Delete University"
						>
							<Trash2 class="w-4 h-4" />
						</button>
					{/if}
				</div>
			{:else if university.website}
				<a
					href={university.website}
					target="_blank"
					rel="noopener noreferrer"
					class="inline-flex items-center gap-1 px-2.5 py-1.5 rounded-xl bg-primary-fixed/50 hover:bg-primary-fixed text-primary text-[11px] font-extrabold transition-all border border-primary/20 shadow-xs hover:shadow-sm shrink-0"
					title="Visit Official University Website"
				>
					<Globe class="w-3.5 h-3.5 text-primary shrink-0" />
					<span class="hidden sm:inline">Website</span>
					<ExternalLink class="w-3 h-3 text-primary/80 shrink-0" />
				</a>
			{/if}
		</div>

		<!-- Description & History -->
		<div class="space-y-3">
			{#if university.university_description}
				<div class="space-y-1 text-xs text-on-surface-variant line-clamp-3 leading-relaxed">
					{@html university.university_description}
				</div>
			{/if}

			{#if university.university_history}
				<div class="space-y-1">
					<span class="text-[10px] font-mono font-bold uppercase tracking-wider text-outline">History</span>
					<div class="text-xs text-on-surface-variant line-clamp-3 leading-relaxed">
						{@html university.university_history}
					</div>
				</div>
			{/if}
		</div>
	</div>

	{#if !isAdmin}
		<div class="pt-2 border-t border-outline-variant/20">
			<a
				href={`/universities/${university.u_id}`}
				class="w-full py-3.5 px-4 rounded-2xl text-center text-xs font-black text-white bg-primary hover:bg-primary-container shadow-md shadow-primary/20 hover:shadow-lg transition-all flex items-center justify-center gap-2 group-hover:shadow-primary/25"
			>
				<span>View University Details</span>
				<ArrowRight class="w-4 h-4" />
			</a>
		</div>
	{/if}
</div>
