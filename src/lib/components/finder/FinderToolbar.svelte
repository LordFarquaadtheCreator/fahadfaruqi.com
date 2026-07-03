<script lang="ts">
	import type { WindowVM } from '$lib/finder/finderStore.svelte';

	interface Props {
		win: WindowVM;
	}

	let { win }: Props = $props();

	const modes = [
		{ key: 'icon', label: 'Icon' },
		{ key: 'list', label: 'List' },
		{ key: 'column', label: 'Column' },
		{ key: 'gallery', label: 'Gallery' }
	];
</script>

<div class="toolbar">
	<div class="nav-buttons">
		<button disabled={!win.canBack} onclick={() => window.os.back(win.id)} aria-label="Back">&#8592;</button>
		<button disabled={!win.canForward} onclick={() => window.os.forward(win.id)} aria-label="Forward">&#8594;</button>
	</div>

	<span class="breadcrumb">{win.title}</span>

	<div class="view-modes">
		{#each modes as mode}
			<button
				class:active={win.viewMode === mode.key}
				onclick={() => window.os.setViewMode(win.id, mode.key)}
			>
				{mode.label}
			</button>
		{/each}
	</div>
</div>

<style>
	.toolbar {
		display: flex;
		align-items: center;
		gap: 8px;
		padding: 6px 12px;
		border-bottom: 1px solid var(--gb-border, #504945);
		background: var(--gb-bg, #282828);
	}

	.nav-buttons {
		display: flex;
		gap: 4px;
	}

	.nav-buttons button {
		background: var(--gb-border2, #3c3836);
		border: 1px solid var(--gb-border, #504945);
		color: var(--gb-fg, #ebdbb2);
		padding: 2px 8px;
		cursor: pointer;
		font-family: inherit;
		font-size: 12px;
	}

	.nav-buttons button:disabled {
		opacity: 0.3;
		cursor: not-allowed;
	}

	.breadcrumb {
		flex: 1;
		font-size: 12px;
		color: var(--gb-fg2, #a89984);
	}

	.view-modes {
		display: flex;
		gap: 2px;
	}

	.view-modes button {
		background: transparent;
		border: 1px solid var(--gb-border, #504945);
		color: var(--gb-fg2, #a89984);
		padding: 2px 8px;
		cursor: pointer;
		font-family: inherit;
		font-size: 11px;
	}

	.view-modes button.active {
		background: var(--gb-border2, #3c3836);
		color: var(--gb-fg, #ebdbb2);
	}
</style>
