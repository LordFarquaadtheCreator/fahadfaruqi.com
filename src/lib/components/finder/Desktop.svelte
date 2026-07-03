<script lang="ts">
	import { onMount } from 'svelte';
	import { finderState, initFinder } from '$lib/finder/finderStore.svelte';
	import FinderWindow from './FinderWindow.svelte';

	onMount(() => {
		return initFinder();
	});
</script>

<div class="desktop" role="presentation">
	{#each finderState.vm?.windows ?? [] as win (win.id)}
		<FinderWindow {win} />
	{/each}
</div>

<style>
	.desktop {
		position: fixed;
		inset: 0;
		pointer-events: none;
		z-index: 100;
	}
	.desktop > :global(*) {
		pointer-events: auto;
	}
</style>
