<script lang="ts">
	let {
		href,
		label,
		image,
		imageAlt = ''
	} = $props<{
		href: string;
		label: string;
		image?: string;
		imageAlt?: string;
	}>();
</script>

<a {href} class="gothic-panel">
	<div class="panel-img-wrap">
		{#if image}
			<img src={image} alt={imageAlt} />
		{:else}
			<div class="panel-placeholder"></div>
		{/if}
	</div>
	<div class="hud-bracket-tl"></div>
	<div class="hud-bracket-tr"></div>
	<div class="hud-bracket-bl"></div>
	<div class="hud-bracket-br"></div>
	<div class="panel-footer">
		<span class="panel-label tech-label">{label}</span>
		<span class="panel-arrow">→</span>
	</div>
</a>

<style>
	.gothic-panel {
		display: block;
		position: relative;
		border: 1px solid rgba(157, 78, 221, 0.4);
		text-decoration: none;
		overflow: hidden;
		transition:
			border-color 200ms ease,
			box-shadow 200ms ease;
		background: #09090b;
	}

	.gothic-panel:hover {
		border-color: rgba(157, 78, 221, 0.9);
		box-shadow: 0 0 28px rgba(157, 78, 221, 0.35);
	}

	.gothic-panel:hover .panel-img-wrap img {
		opacity: 0.65;
	}

	.gothic-panel:hover .panel-label {
		color: #c084fc;
	}

	.panel-img-wrap {
		width: 100%;
		aspect-ratio: 4 / 3;
		background: #9d4edd;
		overflow: hidden;
	}

	.panel-img-wrap img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		filter: grayscale(1) contrast(1.15);
		mix-blend-mode: luminosity;
		opacity: 0.85;
		transition: opacity 200ms ease;
		display: block;
	}

	.panel-placeholder {
		width: 100%;
		height: 100%;
		background: repeating-linear-gradient(
			45deg,
			rgba(157, 78, 221, 0.08) 0px,
			rgba(157, 78, 221, 0.08) 1px,
			transparent 1px,
			transparent 8px
		);
	}

	.panel-footer {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 0.5rem 0.75rem;
		border-top: 1px solid rgba(157, 78, 221, 0.3);
		background: rgba(9, 9, 11, 0.9);
	}

	.panel-label {
		color: var(--on-surface-variant);
		transition: color 200ms ease;
	}

	.panel-arrow {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.75rem;
		color: rgba(157, 78, 221, 0.6);
	}

	/* HUD brackets override for correct positioning */
	:global(.gothic-panel .hud-bracket-tl),
	:global(.gothic-panel .hud-bracket-tr),
	:global(.gothic-panel .hud-bracket-bl),
	:global(.gothic-panel .hud-bracket-br) {
		border-color: rgba(157, 78, 221, 0.7);
		z-index: 2;
		width: 12px;
		height: 12px;
	}
</style>
