<script lang="ts">
	import type { WindowVM } from '$lib/finder/finderStore.svelte';

	interface Props {
		win: WindowVM;
	}

	let { win }: Props = $props();
</script>

<nav class="sidebar">
	{#each win.sidebar as item}
		<button
			class="sidebar-item"
			class:active={item.active}
			style:padding-left="{8 + item.level * 12}px"
			onclick={() => window.os.navigate(win.id, item.path)}
		>
			{item.name}
		</button>
	{/each}
</nav>

<style>
	.sidebar {
		width: 160px;
		flex-shrink: 0;
		border-right: 1px solid var(--gb-border, #504945);
		overflow-y: auto;
		padding: 4px 0;
		background: var(--gb-bg, #282828);
	}

	.sidebar-item {
		display: block;
		width: 100%;
		text-align: left;
		background: transparent;
		border: none;
		color: var(--gb-fg2, #a89984);
		padding: 4px 8px;
		cursor: pointer;
		font-family: inherit;
		font-size: 12px;
	}

	.sidebar-item:hover {
		background: var(--gb-border2, #3c3836);
		color: var(--gb-fg, #ebdbb2);
	}

	.sidebar-item.active {
		background: var(--gb-border, #504945);
		color: var(--gb-fg, #ebdbb2);
	}
</style>
