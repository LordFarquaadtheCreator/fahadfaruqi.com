<script lang="ts">
	import type { EntryVM } from '$lib/finder/finderStore.svelte';

	interface Props {
		entries: EntryVM[];
		winId: string;
		selected: string;
	}

	let { entries, winId, selected }: Props = $props();
</script>

<div class="column-view">
	{#each entries as entry}
		<button
			class="column-item"
			class:selected={selected === entry.name}
			onclick={() => window.os.selectEntry(winId, entry.name)}
			ondblclick={() => window.os.openEntry(winId, entry.path)}
		>
			<span class="col-icon">{entry.kind === 'dir' ? '📁' : '📄'}</span>
			<span class="col-name">{entry.name}</span>
		</button>
	{/each}
</div>

<style>
	.column-view {
		display: flex;
		flex-direction: column;
		gap: 2px;
		padding: 4px;
	}

	.column-item {
		display: flex;
		align-items: center;
		gap: 8px;
		padding: 4px 8px;
		background: transparent;
		border: none;
		border-left: 3px solid transparent;
		color: var(--gb-fg, #ebdbb2);
		cursor: pointer;
		font-family: inherit;
		font-size: 12px;
		text-align: left;
	}

	.column-item:hover {
		background: var(--gb-border2, #3c3836);
	}

	.column-item.selected {
		background: var(--gb-border, #504945);
		border-left-color: var(--gb-aqua, #83a598);
	}
</style>
