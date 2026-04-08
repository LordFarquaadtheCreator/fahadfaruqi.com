<script lang="ts">
	import { reveal } from '$lib/actions/reveal';

	let {
		ghostText = '',
		eyebrow = '',
		title,
		description = '',
		headerRight,
		variant = 'default'
	} = $props<{
		ghostText?: string;
		eyebrow?: string;
		title: string;
		description?: string;
		headerRight?: import('svelte').Snippet;
		variant?: 'default' | 'uplink';
	}>();
</script>

<div class="section-header-wrapper" class:uplink-variant={variant === 'uplink'}>
	{#if ghostText}
		<div class="ghost-text" aria-hidden="true">{ghostText}</div>
	{/if}

	<div class="section-header reveal" use:reveal>
		{#if variant === 'default'}
			<div class="header-top">
				{#if eyebrow}
					<span class="eyebrow">{eyebrow}</span>
				{/if}
				{#if headerRight}
					{@render headerRight()}
				{/if}
			</div>
		{:else if variant === 'uplink'}
			<div class="status-line">
				<span class="status-pip status-active"></span>
				<span class="tech-label" style="color: var(--tertiary);">SYSTEM_STATUS: OPEN_CHANNEL</span>
			</div>
		{/if}

		<h2 class="section-title chromatic">{title}</h2>

		{#if description}
			<p class="section-desc">{description}</p>
		{/if}
	</div>
</div>

<style>
	.section-header-wrapper {
		position: relative;
		overflow: hidden;
		margin-bottom: 3rem;
	}

	.ghost-text {
		position: absolute;
		top: 0;
		left: -2rem;
		font-size: clamp(8rem, 20vw, 16rem);
		font-weight: 900;
		color: var(--on-surface);
		opacity: 0.03;
		pointer-events: none;
		user-select: none;
		letter-spacing: -0.05em;
	}

	.section-header {
		position: relative;
		border-bottom: 1px solid rgba(175, 135, 134, 0.15);
		padding-bottom: 1.5rem;
	}

	.uplink-variant .section-header {
		border-bottom: none;
	}

	.uplink-variant .ghost-text {
		display: none;
	}

	.header-top {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 0.75rem;
		flex-wrap: wrap;
		gap: 0.5rem;
	}

	.status-line {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		margin-bottom: 1rem;
	}

	.section-title {
		font-size: clamp(2rem, 5vw, 3.5rem);
		font-weight: 900;
		letter-spacing: -0.02em;
		margin: 0;
		color: var(--on-surface);
	}

	.uplink-variant .section-title {
		font-size: clamp(4rem, 12vw, 10rem);
		letter-spacing: -0.05em;
		line-height: 0.85;
		color: var(--primary);
		mix-blend-mode: screen;
	}

	.section-desc {
		color: var(--on-surface-variant);
		font-size: 0.9rem;
		margin-top: 0.75rem;
		max-width: 36rem;
	}

	@media (max-width: 768px) {
		.ghost-text {
			display: none;
		}
	}
</style>
