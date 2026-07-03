<script lang="ts">
	import type { EntryVM } from '$lib/finder/finderStore.svelte';

	interface Props {
		entries: EntryVM[];
		winId: string;
		selected: string;
	}

	let { entries, winId, selected }: Props = $props();
</script>

<div class="gallery-view">
	{#each entries as entry}
		<button
			class="gallery-item"
			class:selected={selected === entry.name}
			onclick={() => window.os.selectEntry(winId, entry.name)}
			ondblclick={() => window.os.openEntry(winId, entry.path)}
		>
			<span class="g-icon">{entry.kind === 'dir' ? '📁' : '📄'}</span>
			<span class="g-name">{entry.name}</span>
		</button>
	{/each}
</div>

<style>
	.gallery-view {
		display: flex;
		flex-wrap: wrap;
		gap: 16px;
		padding: 12px;
	}

	.gallery-item {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 6px;
		width: 120px;
		padding: 12px;
		background: transparent;
		border: 2px solid transparent;
		border-radius: 8px;
		cursor: pointer;
		color: var(--gb-fg, #ebdbb2);
		font-family: inherit;
		font-size: 12px;
	}

	.gallery-item:hover {
		background: var(--gb-border2, #3c3836);
	}

	.gallery-item.selected {
		background: var(--gb-border, #504945);
		border-color: var(--gb-aqua, #83a598);
	}

	.g-icon {
		font-size: 48px;
	}
</style>
