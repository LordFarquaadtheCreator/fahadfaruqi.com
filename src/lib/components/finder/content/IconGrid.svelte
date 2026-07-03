<script lang="ts">
	import type { EntryVM } from '$lib/finder/finderStore.svelte';

	interface Props {
		entries: EntryVM[];
		winId: string;
		selected: string;
	}

	let { entries, winId, selected }: Props = $props();
</script>

<div class="icon-grid">
	{#each entries as entry}
		<button
			class="icon-item"
			class:selected={selected === entry.name}
			onclick={() => window.os.selectEntry(winId, entry.name)}
			ondblclick={() => window.os.openEntry(winId, entry.path)}
		>
			<span class="icon">{entry.kind === 'dir' ? '📁' : '📄'}</span>
			<span class="name">{entry.name}</span>
		</button>
	{/each}
</div>

<style>
	.icon-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(80px, 1fr));
		gap: 12px;
		padding: 8px;
	}

	.icon-item {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 4px;
		padding: 8px;
		background: transparent;
		border: 2px solid transparent;
		border-radius: 6px;
		cursor: pointer;
		color: var(--gb-fg, #ebdbb2);
		font-family: inherit;
		font-size: 12px;
	}

	.icon-item:hover {
		background: var(--gb-border2, #3c3836);
	}

	.icon-item.selected {
		background: var(--gb-border, #504945);
		border-color: var(--gb-aqua, #83a598);
	}

	.icon {
		font-size: 32px;
	}

	.name {
		text-align: center;
		word-break: break-word;
	}
</style>
