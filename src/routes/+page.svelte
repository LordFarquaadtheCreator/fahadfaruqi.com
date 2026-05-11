<script lang="ts">
	import portfolio from '$lib/data/portfolio';
	import Icon from '$lib/components/Icon.svelte';

	let currentIndex = $state(0);

	function goToSlide(index: number) {
		currentIndex = index;
	}
</script>

<svelte:head>
	<title>{portfolio.profile.name}</title>
	<meta name="description" content={portfolio.home.description} />
</svelte:head>

<main class="home-page">
	<nav class="centered-nav">
		<h1 class="brand-title">{portfolio.home.title}</h1>

		<div class="nav-links">
			{#each portfolio.home.slides as slide, index}
				<a
					href={slide.href}
					class="nav-link"
					class:active={index === currentIndex}
					onclick={(e) => {
						if (slide.href === '/') {
							e.preventDefault();
							goToSlide(index);
						}
					}}
				>
					{slide.label}
				</a>
			{/each}
		</div>

		<div class="social-icons">
			<a href="mailto:{portfolio.contact.email}" aria-label="Email">
				<Icon name="mail" />
			</a>
			<a href="https://github.com/fahadfaruqi" target="_blank" rel="noopener" aria-label="GitHub">
				<Icon name="terminal" />
			</a>
		</div>
	</nav>
</main>

<style>
	.home-page {
		height: 100vh;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		position: relative;
		z-index: 1;
	}

	.centered-nav {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 3rem;
	}

	.brand-title {
		font-size: clamp(2rem, 5vw, 3rem);
		font-weight: 700;
		letter-spacing: 0.1em;
		text-transform: uppercase;
		color: var(--on-surface);
	}

	.nav-links {
		display: flex;
		align-items: center;
		gap: 2rem;
		border-bottom: 1px solid var(--hairline);
		padding-bottom: 0.5rem;
	}

	.nav-link {
		text-decoration: none;
		color: var(--on-surface-variant);
		font-family: var(--font-mono);
		font-size: 0.875rem;
		letter-spacing: -0.01em;
		padding-bottom: 0.25rem;
		border-bottom: 2px solid transparent;
		transition: color 200ms ease;
	}

	.nav-link:hover {
		color: var(--primary);
	}

	.nav-link.active {
		color: var(--primary);
		border-bottom-color: var(--primary);
	}

	.social-icons {
		display: flex;
		gap: 1.5rem;
		margin-top: 0.5rem;
	}

	.social-icons a {
		color: var(--on-surface-variant);
		transition: color 200ms ease;
	}

	.social-icons a:hover {
		color: var(--primary);
	}

	@media (min-width: 768px) {
		.nav-links {
			gap: 2.5rem;
		}
	}

	@media (max-width: 767px) {
		.centered-nav {
			gap: 2rem;
		}

		.brand-title {
			font-size: 1.5rem;
		}

		.nav-links {
			gap: 1.5rem;
		}
	}
</style>