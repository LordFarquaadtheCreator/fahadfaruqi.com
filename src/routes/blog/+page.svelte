<script lang="ts">
	import { reveal } from '$lib/actions/reveal';
	import { blogPosts } from '$lib/blogs';
	import portfolio from '$lib/data/portfolio';
</script>

<svelte:head>
	<title>{portfolio.blog.title} | {portfolio.profile.name}</title>
	<meta name="description" content={portfolio.blog.description} />
</svelte:head>

<main class="blog-page">
	<header class="page-header reveal" use:reveal>
		<p class="tech-label">Blog</p>
		<h1>{portfolio.blog.title}</h1>
		<p>{portfolio.blog.description}</p>
	</header>

	<section class="dispatch-grid" aria-label="Dispatches">
		{#each blogPosts as post, index}
			<article class:featured={index === 0} class="reveal" use:reveal>
				<a href={`/blog/${post.slug}`} aria-label={post.title}>
					<div class="image-panel" role="img" aria-label={post.imageAlt}></div>
					<div class="post-copy">
						<div class="post-meta">
							<span class="tech-label">{post.category}</span>
							<time>{post.date}</time>
						</div>
						<h2>{post.title}</h2>
						<p>{post.excerpt}</p>
						<pre aria-label={post.codeTitle}><code>{post.codeSnippet}</code></pre>
					</div>
				</a>
			</article>
		{/each}
	</section>
</main>

<style>
	.blog-page {
		max-width: var(--container);
		margin: 0 auto;
		padding: clamp(4rem, 8vw, var(--section-gap)) var(--safe-margin);
	}

	.page-header {
		max-width: 760px;
		margin-bottom: var(--section-gap);
	}

	.page-header .tech-label {
		color: var(--secondary);
	}

	h1 {
		margin-top: 1rem;
		font-size: clamp(3.25rem, 8vw, 6.5rem);
	}

	.page-header p:last-child {
		margin-top: 1.5rem;
		border-left: 1px solid var(--hairline-strong);
		padding-left: 1.5rem;
		color: var(--on-surface-variant);
		font-size: 1.08rem;
	}

	.dispatch-grid {
		display: grid;
		grid-template-columns: repeat(12, minmax(0, 1fr));
		gap: var(--gutter);
	}

	article {
		position: relative;
		grid-column: span 4;
		border: 1px solid var(--hairline);
		background: color-mix(in srgb, var(--surface-container-lowest) 66%, transparent);
		transition:
			transform 220ms ease,
			border-color 220ms ease,
			background-color 220ms ease;
	}

	article::after {
		content: '';
		position: absolute;
		inset: 0;
		opacity: 0;
		pointer-events: none;
		background-image: var(--grain-texture);
		transition: opacity 180ms ease;
	}

	article:hover,
	article:focus-within {
		transform: translateY(-6px);
		border-color: var(--primary);
		background: color-mix(in srgb, var(--surface-container-lowest) 80%, var(--primary) 4%);
	}

	article:hover::after,
	article:focus-within::after {
		opacity: 0.14;
	}

	article.featured {
		grid-column: span 8;
	}

	article:nth-child(3) {
		transform: translateY(3rem);
	}

	a {
		display: flex;
		flex-direction: column;
		min-height: 100%;
	}

	.image-panel {
		aspect-ratio: 16 / 10;
		border-bottom: 1px solid var(--hairline);
		background:
			radial-gradient(circle at 22% 18%, color-mix(in srgb, var(--primary) 24%, transparent), transparent 30%),
			radial-gradient(circle at 78% 70%, color-mix(in srgb, var(--secondary) 18%, transparent), transparent 34%),
			linear-gradient(180deg, color-mix(in srgb, var(--surface-container-lowest) 20%, transparent), transparent),
			var(--surface-container-high);
		filter: grayscale(1);
		transition: filter 240ms ease;
	}

	article:hover .image-panel,
	article:focus-within .image-panel {
		filter: grayscale(0);
	}

	.post-copy {
		display: flex;
		flex: 1;
		flex-direction: column;
		padding: clamp(1.25rem, 3vw, 2rem);
	}

	.post-meta {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 1rem;
		color: var(--on-surface-variant);
		font-family: var(--font-mono);
		font-size: 0.78rem;
	}

	.post-meta .tech-label {
		color: var(--secondary);
	}

	h2 {
		margin-top: 1.5rem;
		font-size: clamp(1.7rem, 3vw, 2.5rem);
	}

	.post-copy p {
		margin-top: 1rem;
		color: var(--on-surface-variant);
	}

	pre {
		margin: 2rem 0 0;
		overflow-x: auto;
		background: var(--surface-container-low);
		padding: 1rem;
		color: var(--on-surface);
		font-family: var(--font-mono);
		font-size: 0.82rem;
	}

	@media (max-width: 960px) {
		article,
		article.featured {
			grid-column: 1 / -1;
		}

		article:nth-child(3) {
			transform: none;
		}
	}
</style>
