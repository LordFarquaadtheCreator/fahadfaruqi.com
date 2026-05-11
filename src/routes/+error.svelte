<script lang="ts">
	import { goto } from '$app/navigation';
	import logo from '$lib/assets/favicon.svg';
	import { onMount } from 'svelte';

	onMount(() => {
		goto('/');
	});
</script>

<main class="redirect-container">
	<div class="redirect-content">
		<img src={logo} alt="" aria-hidden="true" />
		<p>Returning home...</p>
	</div>
</main>

<style>
	.redirect-container {
		min-height: 100vh;
		display: grid;
		place-items: center;
		background: var(--background);
	}

	.redirect-content {
		display: grid;
		justify-items: center;
		gap: 1rem;
		color: var(--on-surface-variant);
		font-family: var(--font-mono);
		font-size: 0.85rem;
		text-transform: uppercase;
		letter-spacing: 0.1em;
	}

	img {
		width: 56px;
		height: 56px;
		border: 1px solid var(--hairline);
		background: var(--surface-container-lowest);
		animation: mark-in 700ms ease both;
	}

	@keyframes mark-in {
		from {
			opacity: 0;
			transform: scale(0.9);
		}

		to {
			opacity: 1;
			transform: scale(1);
		}
	}
</style>
