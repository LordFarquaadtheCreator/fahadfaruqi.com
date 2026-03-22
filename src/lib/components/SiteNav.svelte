<script lang="ts">
	import type { PortfolioLink } from '$lib/data/portfolio';

	interface NavLink {
		label: string;
		href: string;
	}

	let {
		name,
		links,
		socialLinks
	} = $props<{
		name: string;
		links: NavLink[];
		socialLinks: PortfolioLink[];
	}>();
</script>

<div class="site-nav-wrap">
	<nav class="section-shell surface site-nav" aria-label="Primary">
		<a class="site-nav__brand" href="#top">{name}</a>

		<div class="site-nav__links">
			{#each links as link}
				<a href={link.href}>{link.label}</a>
			{/each}
		</div>

		<div class="site-nav__social">
			{#each socialLinks as link}
				<a href={link.href} target={link.href.startsWith('http') ? '_blank' : undefined} rel="noreferrer">
					{link.label}
				</a>
			{/each}
		</div>
	</nav>
</div>

<style>
	.site-nav-wrap {
		position: sticky;
		top: 0;
		z-index: 10;
		padding: 1rem 0 0;
	}

	.site-nav {
		display: grid;
		grid-template-columns: auto 1fr auto;
		align-items: center;
		gap: 1rem;
		padding: 0.9rem 1.1rem;
		border-radius: 999px;
	}

	.site-nav__brand {
		font-size: 0.95rem;
		font-weight: 700;
		letter-spacing: 0.06em;
		text-transform: uppercase;
	}

	.site-nav__links,
	.site-nav__social {
		display: flex;
		flex-wrap: wrap;
		gap: 0.9rem;
	}

	.site-nav__links {
		justify-content: center;
	}

	.site-nav__social {
		justify-content: flex-end;
	}

	.site-nav a {
		color: var(--text-soft);
		font-size: 0.92rem;
		font-weight: 500;
		transition: color 180ms ease;
	}

	.site-nav a:hover,
	.site-nav__brand {
		color: var(--text);
	}

	@media (max-width: 900px) {
		.site-nav {
			grid-template-columns: 1fr;
			justify-items: center;
			border-radius: 2rem;
		}

		.site-nav__social {
			justify-content: center;
		}
	}
</style>
