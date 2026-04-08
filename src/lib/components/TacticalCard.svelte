<script lang="ts">
	import { reveal } from '$lib/actions/reveal';

	let {
		children,
		headerIcon = '',
		headerLabel = '',
		headerRight = '',
		hudBrackets = false,
		variant = 'default',
		class: className = ''
	} = $props<{
		children: import('svelte').Snippet;
		headerIcon?: string;
		headerLabel?: string;
		headerRight?: import('svelte').Snippet;
		hudBrackets?: boolean;
		variant?: 'default' | 'primary' | 'highlight' | 'stream';
		class?: string;
	}>();
</script>

<div
	class="tactical-card {className}"
	class:tactical-card--primary={variant === 'primary'}
	class:tactical-card--highlight={variant === 'highlight'}
	class:tactical-card--stream={variant === 'stream'}
	use:reveal
>
	{#if hudBrackets}
		<div class="hud-bracket-tl"></div>
		<div class="hud-bracket-tr"></div>
		<div class="hud-bracket-bl"></div>
		<div class="hud-bracket-br"></div>
	{/if}

	{#if headerLabel || headerIcon}
		<div class="card-header">
			<div class="header-left">
				{#if headerIcon}
					<span class="material-symbols-outlined" style="color: var(--primary);">{headerIcon}</span>
				{/if}
				<span class="tech-label">{headerLabel}</span>
			</div>
			{#if headerRight}
				{@render headerRight()}
			{/if}
		</div>
	{/if}

	{@render children()}
</div>

<style>
	.tactical-card {
		background: var(--surface-container-low);
		border: 1px solid rgba(175, 135, 134, 0.15);
		padding: 1.5rem;
		position: relative;
		transition: all 200ms ease;
	}

	.tactical-card:hover {
		border-color: rgba(255, 179, 178, 0.4);
		box-shadow: 0 0 30px rgba(255, 179, 178, 0.1);
		animation: glitch 0.3s ease-in-out;
	}

	@keyframes glitch {
		0% { transform: translate(0); }
		20% { transform: translate(-1px, 1px); }
		40% { transform: translate(1px, -1px); }
		60% { transform: translate(-1px, -1px); }
		80% { transform: translate(1px, 1px); }
		100% { transform: translate(0); }
	}

	.tactical-card--primary {
		background: var(--surface-container-low);
	}

	.tactical-card--highlight {
		background: var(--surface-container-high);
	}

	.tactical-card--stream {
		background: var(--surface-container-high);
		border-left: 4px solid var(--secondary-container);
	}

	.tactical-card--stream:hover {
		border-left-color: var(--secondary);
	}

	.card-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 1.25rem;
		padding-bottom: 0.75rem;
		border-bottom: 1px solid rgba(175, 135, 134, 0.1);
	}

	.header-left {
		display: flex;
		align-items: center;
		gap: 0.75rem;
	}

	.card-header .material-symbols-outlined {
		font-size: 20px;
	}
</style>
