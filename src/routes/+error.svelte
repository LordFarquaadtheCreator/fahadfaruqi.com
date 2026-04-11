<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	onMount(() => {
		// Immediate redirect to home
		goto('/');
	});
</script>

<!-- Fallback content while redirecting -->
<div class="redirect-container">
	<div class="redirect-content">
		<span class="material-symbols-outlined redirect-icon">radar</span>
		<p class="redirect-text">REROUTING...</p>
	</div>
</div>

<style>
	.redirect-container {
		display: flex;
		align-items: center;
		justify-content: center;
		min-height: 100vh;
		background: var(--surface-container-lowest);
	}

	.redirect-content {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1rem;
		color: var(--on-surface-variant);
	}

	.redirect-icon {
		font-size: 3rem;
		color: var(--primary);
		animation: pulse 1s ease-in-out infinite;
	}

	.redirect-text {
		font-family: 'JetBrains Mono', monospace;
		font-size: 1rem;
		letter-spacing: 0.1em;
	}

	@keyframes pulse {
		0%, 100% { opacity: 1; transform: scale(1); }
		50% { opacity: 0.5; transform: scale(0.95); }
	}
</style>
