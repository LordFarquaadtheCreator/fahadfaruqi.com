<script lang="ts">
	import { page } from '$app/state';
	import logo from '$lib/assets/favicon.svg';
	import type { NavigationData, PortfolioLink } from '$lib/data/portfolio';

	let {
		name,
		navigation
	} = $props<{
		name: string;
		navigation: NavigationData;
	}>();

	const pathname = $derived(page.url.pathname);

	function isActive(href: string) {
		if (href === '/') return pathname === '/';
		return pathname === href || pathname.startsWith(`${href}/`);
	}
</script>

<header class="site-header" style={`--active-index: ${Math.max(0, navigation.links.findIndex((item: PortfolioLink) => isActive(item.href)))}`}>
	<a class="brand-mark" href="/" aria-label="{name} home">
		<img src={logo} alt="" aria-hidden="true" />
		<span class="brand-mark__serif">{name}</span>
	</a>

	<nav class="site-nav slider-nav" aria-label="Primary">
		<span class="slider-indicator" aria-hidden="true"></span>
		{#each navigation.links as item}
			<a class="nav-link" class:active={isActive(item.href)} href={item.href}>
				{item.label}
			</a>
		{/each}
	</nav>
</header>

<style>
	.site-header {
		position: sticky;
		top: 0;
		z-index: 50;
		display: grid;
		grid-template-columns: minmax(0, 1fr) auto;
		align-items: stretch;
		border-bottom: 1px solid var(--hairline);
		background: color-mix(in srgb, var(--surface) 90%, transparent);
		backdrop-filter: blur(18px);
	}

	.brand-mark {
		display: flex;
		align-items: center;
		gap: 1rem;
		padding: 1rem var(--safe-margin);
		border-right: 1px solid var(--hairline);
	}

	.brand-mark img {
		width: 28px;
		height: 28px;
		border: 1px solid var(--hairline);
		background: var(--surface-container-lowest);
	}

	.brand-mark__serif {
		font-family: var(--font-display);
		font-size: 1.35rem;
		font-style: italic;
		font-weight: 700;
		line-height: 1;
	}

	.brand-mark__mono {
		color: var(--on-surface-variant);
		font-family: var(--font-mono);
		font-size: 0.7rem;
		letter-spacing: 0.1em;
		text-transform: uppercase;
	}

	.site-nav {
		position: relative;
		display: flex;
		align-items: stretch;
		border-left: 1px solid var(--hairline);
	}

	.slider-indicator {
		position: absolute;
		inset: 0 auto 0 0;
		width: calc(100% / 3);
		background: color-mix(in srgb, var(--primary) 9%, transparent);
		border-bottom: 1px solid var(--secondary);
		transform: translateX(calc(var(--active-index) * 100%));
		transition: transform 360ms cubic-bezier(0.22, 1, 0.36, 1);
		pointer-events: none;
	}

	.nav-link {
		position: relative;
		z-index: 1;
		display: grid;
		place-items: center;
		min-width: 6.5rem;
		padding: 1rem 1.25rem;
		border-right: 1px solid var(--hairline);
		color: var(--on-surface-variant);
		font-family: var(--font-mono);
		font-size: 0.75rem;
		font-weight: 600;
		letter-spacing: 0.1em;
		text-transform: uppercase;
		text-decoration-thickness: 1px;
		text-underline-offset: 0.35em;
		transition:
			color 160ms ease,
			background-color 160ms ease,
			border-color 160ms ease;
	}

	.nav-link:hover,
	.nav-link:focus-visible,
	.nav-link.active {
		color: var(--primary);
		outline: none;
		text-decoration: underline;
		text-decoration-color: var(--secondary);
	}

	.nav-link::after {
		content: '';
		position: absolute;
		inset: 0;
		opacity: 0;
		pointer-events: none;
		background-image: var(--grain-texture);
		mix-blend-mode: multiply;
		transition: opacity 180ms ease;
	}

	.nav-link:hover::after,
	.nav-link:focus-visible::after {
		opacity: 0.18;
	}

	@media (max-width: 820px) {
		.site-header {
			grid-template-columns: 1fr;
		}

		.brand-mark {
			justify-content: space-between;
			border-right: 0;
		}

		.site-nav {
			display: grid;
			grid-template-columns: repeat(3, 1fr);
			border-top: 1px solid var(--hairline);
		}

		.slider-indicator {
			width: calc(100% / 3);
		}

		.nav-link {
			min-width: 0;
			padding: 0.8rem 0.35rem;
			border-right: 1px solid var(--hairline);
			font-size: 0.65rem;
		}

		.nav-link:last-child {
			border-right: 0;
		}
	}
</style>
