<script lang="ts">
	let {
		title,
		children
	} = $props<{
		title: string;
		children: import('svelte').Snippet;
	}>();
</script>

<div class="gothic-shell">
	<header class="gothic-banner">
		<a href="/" class="banner-home">
			<span class="banner-arrow">←</span>
			<span class="banner-name">{title}</span>
		</a>
		<div class="banner-right">
			<svg class="barcode-mini" viewBox="0 0 60 16" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
				{#each Array(28) as _, i}
					<rect
						x={i * 2.1 + (i % 4 === 0 ? 0.3 : 0)}
						y="0"
						width={i % 5 === 0 ? 1.4 : 0.8}
						height="16"
						fill="rgba(157,78,221,0.55)"
					/>
				{/each}
			</svg>
			<span class="banner-year tech-label">2025</span>
		</div>
	</header>

	<main class="gothic-main">
		{@render children()}
	</main>
</div>

<style>
	.gothic-shell {
		min-height: 100vh;
		background: #09090b;
		display: flex;
		flex-direction: column;
	}

	.gothic-banner {
		position: sticky;
		top: 0;
		z-index: 50;
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 0 1.5rem;
		height: 52px;
		background: rgba(9, 9, 11, 0.92);
		backdrop-filter: blur(12px);
		border-bottom: 1px solid rgba(157, 78, 221, 0.35);
	}

	.banner-home {
		display: flex;
		align-items: center;
		gap: 0.6rem;
		text-decoration: none;
		transition: opacity 150ms ease;
	}

	.banner-home:hover {
		opacity: 0.7;
	}

	.banner-arrow {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.85rem;
		color: rgba(157, 78, 221, 0.8);
	}

	.banner-name {
		font-family: 'Pirata One', cursive;
		font-size: 1.4rem;
		color: #c084fc;
		letter-spacing: 0.04em;
		line-height: 1;
	}

	.banner-right {
		display: flex;
		align-items: center;
		gap: 0.75rem;
	}

	.barcode-mini {
		width: 48px;
		height: 14px;
	}

	.banner-year {
		color: rgba(157, 78, 221, 0.5);
		font-size: 0.55rem;
	}

	.gothic-main {
		flex: 1;
		padding: 2rem 0 4rem;
	}
</style>
