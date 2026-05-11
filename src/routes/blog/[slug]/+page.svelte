<script lang="ts">
	import type { BlogPost } from '$lib/blogs';
	import portfolio from '$lib/data/portfolio';

	let { data } = $props<{
		data: {
			post: BlogPost;
			nextPost: BlogPost;
		};
	}>();
</script>

<svelte:head>
	<title>{data.post.title} | {portfolio.profile.name}</title>
	<meta name="description" content={data.post.excerpt} />
</svelte:head>

<main class="post-page">
	<article>
		<header class="post-header">
			<div class="post-meta">
				<span class="tech-label">{data.post.category}</span>
				<time>{data.post.date}</time>
			</div>
			<h1>{data.post.title}</h1>
			<p>{data.post.excerpt}</p>
		</header>

		<div class="hero-image" role="img" aria-label={data.post.imageAlt}></div>

		<div class="post-body">
			<pre aria-label={data.post.codeTitle}><code>{data.post.codeSnippet}</code></pre>
			<blockquote>{data.post.pullQuote}</blockquote>
			{@html data.post.html}
		</div>
	</article>

	<a class="next-article" href={`/blog/${data.nextPost.slug}`}>
		<span class="tech-label">Next Article</span>
		<strong>{data.nextPost.title}</strong>
	</a>
</main>

<style>
	.post-page {
		max-width: var(--container);
		margin: 0 auto;
		padding: clamp(4rem, 8vw, var(--section-gap)) var(--safe-margin);
	}

	article {
		display: grid;
		grid-template-columns: repeat(12, minmax(0, 1fr));
		gap: var(--gutter);
	}

	.post-header {
		grid-column: 3 / span 8;
	}

	.post-meta {
		display: flex;
		justify-content: space-between;
		gap: 1rem;
		color: var(--on-surface-variant);
		font-family: var(--font-mono);
		font-size: 0.8rem;
	}

	.post-meta .tech-label {
		color: var(--secondary);
	}

	h1 {
		margin-top: 1.5rem;
		font-size: clamp(3rem, 7vw, 6rem);
	}

	.post-header p {
		margin-top: 1.5rem;
		border-left: 1px solid var(--hairline-strong);
		padding-left: 1.5rem;
		color: var(--on-surface-variant);
		font-size: 1.12rem;
	}

	.hero-image {
		grid-column: 2 / span 10;
		aspect-ratio: 16 / 7;
		margin: 4rem 0;
		border: 1px solid var(--hairline);
		background:
			radial-gradient(circle at 18% 22%, color-mix(in srgb, var(--primary) 20%, transparent), transparent 32%),
			radial-gradient(circle at 84% 68%, color-mix(in srgb, var(--secondary) 16%, transparent), transparent 36%),
			linear-gradient(180deg, color-mix(in srgb, var(--surface-container-lowest) 20%, transparent), transparent),
			var(--surface-container-high);
		filter: grayscale(1);
	}

	.post-body {
		grid-column: 4 / span 6;
	}

	pre {
		margin: 0 0 3rem;
		overflow-x: auto;
		background: var(--surface-container-low);
		padding: 1.25rem;
		font-family: var(--font-mono);
		font-size: 0.86rem;
	}

	blockquote {
		margin: 0 0 3rem;
		border-left: 1px solid var(--secondary);
		padding-left: 1.5rem;
		color: var(--primary);
		font-family: var(--font-display);
		font-size: clamp(1.8rem, 4vw, 3rem);
		font-style: italic;
		line-height: 1.2;
	}

	.post-body :global(p) {
		margin-top: 1.5rem;
		color: var(--on-surface);
		font-size: 1.08rem;
	}

	.post-body :global(h2) {
		margin-top: 2.5rem;
		font-size: 2rem;
	}

	.post-body :global(blockquote) {
		margin: 2.5rem 0;
		border-left: 1px solid var(--secondary);
		padding-left: 1.5rem;
		color: var(--primary);
		font-family: var(--font-display);
		font-size: clamp(1.7rem, 3vw, 2.6rem);
		font-style: italic;
		line-height: 1.2;
	}

	.post-body :global(pre) {
		margin: 2.5rem 0;
		overflow-x: auto;
		background: var(--surface-container-low);
		padding: 1.25rem;
		font-family: var(--font-mono);
		font-size: 0.86rem;
	}

	.next-article {
		display: grid;
		gap: 0.7rem;
		max-width: 640px;
		margin: var(--section-gap) auto 0;
		border: 1px solid var(--hairline);
		padding: 1.5rem;
		background: color-mix(in srgb, var(--surface-container-lowest) 60%, transparent);
	}

	.next-article .tech-label {
		color: var(--secondary);
	}

	.next-article strong {
		font-family: var(--font-display);
		font-size: 2rem;
	}

	@media (max-width: 900px) {
		.post-header,
		.hero-image,
		.post-body {
			grid-column: 1 / -1;
		}
	}
</style>
