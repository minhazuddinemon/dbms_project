<!-- src/lib/components/ProgramCard.svelte -->
<script lang="ts">
	import type { Program } from '$lib/types/models';
	import { MapPin, ArrowRight, Edit3, Trash2 } from 'lucide-svelte';

	interface Props {
		program: Program;
		uniLogo?: string | null;
		isAdmin?: boolean;
		onedit?: (p: Program) => void;
		ondelete?: (pId: number) => void;
	}

	let { program, uniLogo = null, isAdmin = false, onedit, ondelete }: Props = $props();

	const uniNameRaw = $derived(program.university_name || program.u_name || 'Public University');
	const locationStr = $derived(program.university_location || program.location);

	const nameFontSizeClass = $derived.by(() => {
		const len = uniNameRaw.length;
		if (len > 36) return 'text-[10px] sm:text-[11px] leading-tight';
		if (len > 22) return 'text-[11px] sm:text-xs leading-tight';
		return 'text-xs sm:text-sm leading-snug';
	});
</script>

<div class="glass-panel p-6 rounded-[2.5rem] border border-outline-variant/40 bg-white/95 shadow-md hover:shadow-xl hover:border-primary/50 transition-all duration-300 flex flex-col justify-between space-y-5 group">
	<div class="space-y-4">
		<!-- Header with Logo, Uni Name & Controls -->
		<div class="flex items-center gap-3.5 border-b border-outline-variant/20 pb-3.5">
			{#if uniLogo || program.logo_url || program.university_logo}
				<img
					src={uniLogo || program.logo_url || program.university_logo || ''}
					alt={uniNameRaw}
					class="w-12 h-12 rounded-2xl object-contain p-1 border border-outline-variant/30 bg-white shadow-xs shrink-0"
				/>
			{:else}
				<div class="w-12 h-12 rounded-2xl bg-primary-fixed text-primary flex items-center justify-center font-black text-xl shrink-0 shadow-xs">
					{uniNameRaw.charAt(0)}
				</div>
			{/if}

			<div class="flex-1 min-w-0">
				<h5 class="font-black uppercase text-primary tracking-wider line-clamp-2 break-words {nameFontSizeClass}" title={uniNameRaw}>
					{uniNameRaw}
				</h5>
				{#if locationStr}
					<p class="text-[11px] font-semibold text-on-surface-variant flex items-center gap-1 truncate mt-0.5">
						<MapPin class="w-3 h-3 text-tertiary shrink-0" />
						{locationStr}
					</p>
				{/if}
			</div>

			<span class="px-2.5 py-1 rounded-full bg-primary-fixed text-on-primary-fixed text-[11px] font-extrabold uppercase shrink-0">
				Unit {program.p_unit || 'A'}
			</span>
		</div>

		<!-- Circular Badge or Admin Controls -->
		<div class="space-y-1">
			<div class="flex items-center justify-between">
				<span class="text-[10px] font-mono font-bold text-outline uppercase tracking-wider">Circular #{program.program_id}</span>
				{#if isAdmin}
					<div class="flex items-center gap-1">
						{#if onedit}
							<button
								type="button"
								onclick={() => onedit?.(program)}
								class="p-1 rounded-lg text-primary hover:bg-primary-fixed/40 transition-colors cursor-pointer"
								title="Edit Program"
							>
								<Edit3 class="w-4 h-4" />
							</button>
						{/if}
						{#if ondelete}
							<button
								type="button"
								onclick={() => ondelete?.(program.program_id)}
								class="p-1 rounded-lg text-error hover:bg-error-container/40 transition-colors cursor-pointer"
								title="Delete Program"
							>
								<Trash2 class="w-4 h-4" />
							</button>
						{/if}
					</div>
				{:else}
					<span class="text-[10px] font-bold text-emerald-700 bg-emerald-50 px-2 py-0.5 rounded-md border border-emerald-200">Verified Program</span>
				{/if}
			</div>

			<h4 class="text-xl font-black text-on-surface group-hover:text-primary transition-colors leading-tight">{program.p_name}</h4>
		</div>

		<!-- Parameters Grid -->
		<div class="grid grid-cols-3 gap-2 text-xs font-semibold pt-1">
			<div class="bg-surface-container-low/80 p-2.5 rounded-2xl border border-outline-variant/20 text-center">
				<span class="block text-[10px] uppercase font-bold text-outline">Total Seats</span>
				<span class="font-black text-on-surface text-sm">{program.total_seats}</span>
			</div>
			<div class="bg-surface-container-low/80 p-2.5 rounded-2xl border border-outline-variant/20 text-center">
				<span class="block text-[10px] uppercase font-bold text-outline">Cutmark</span>
				<span class="font-black text-primary text-sm">{program.prev_cutmarks || 'N/A'}</span>
			</div>
			<div class="bg-surface-container-low/80 p-2.5 rounded-2xl border border-outline-variant/20 text-center">
				<span class="block text-[10px] uppercase font-bold text-outline">Deadline</span>
				<span class="font-extrabold text-on-surface font-mono text-[11px] block mt-0.5">{program.deadline ? program.deadline.split('T')[0] : 'TBA'}</span>
			</div>
		</div>

		<!-- Admin Required Fields Badges -->
		{#if isAdmin && program.required_fields && program.required_fields.length > 0}
			<div class="space-y-1 pt-1">
				<span class="text-[10px] font-bold text-outline uppercase tracking-wider block">Required Info:</span>
				<div class="flex flex-wrap gap-1">
					{#each program.required_fields as rf}
						<span class="text-[10px] font-mono font-bold px-2 py-0.5 bg-primary-fixed/40 text-primary rounded-md border border-primary/20">
							{rf}
						</span>
					{/each}
				</div>
			</div>
		{/if}
	</div>

	{#if !isAdmin}
		<a
			href={`/apply/${program.program_id}`}
			class="w-full py-3.5 px-4 rounded-2xl font-bold text-xs text-white bg-primary hover:bg-primary-container shadow-md transition-all flex items-center justify-center gap-2 group-hover:shadow-lg group-hover:shadow-primary/25 cursor-pointer"
		>
			<span>Apply Now</span>
			<ArrowRight class="w-4 h-4" />
		</a>
	{/if}
</div>
