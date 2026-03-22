<script lang="ts">
	import { reveal } from '$lib/actions/reveal';
	import type { AboutData } from '$lib/data/portfolio';
	import SectionHeading from './SectionHeading.svelte';

	let { about } = $props<{ about: AboutData }>();
</script>

<section class="about section-shell" id="about">
	<div class="surface about__panel reveal" use:reveal>
		<SectionHeading
			eyebrow="About"
			title="A builder who likes the full arc from idea to launch."
			copy="The portfolio is deliberately content-first: structured enough for recruiters, but still personal and expressive."
		/>

		<div class="about__top">
			<div class="about__story">
				{#each about.paragraphs as paragraph}
					<p>{paragraph}</p>
				{/each}
			</div>

			<div class="about__education">
				<h3>Education</h3>
				{#each about.education as item}
					<article>
						<strong>{item.institution}</strong>
						<span>{item.credential}</span>
						<small>{item.dates} · {item.location}</small>
					</article>
				{/each}
			</div>
		</div>

		<div class="about__highlights">
			{#each about.highlights as item}
				<article class="about__highlight">
					<p class="eyebrow">{item.eyebrow}</p>
					<h3>{item.title}</h3>
					<p>{item.description}</p>
					{#if item.href && item.linkLabel}
						<a href={item.href} target="_blank" rel="noreferrer">{item.linkLabel}</a>
					{/if}
				</article>
			{/each}
		</div>
	</div>
</section>

<style>
	.about {
		padding: 2rem 0 1rem;
	}

	.about__panel {
		padding: clamp(1.4rem, 3vw, 2.4rem);
	}

	.about__top {
		display: grid;
		grid-template-columns: minmax(0, 1.5fr) minmax(18rem, 0.95fr);
		gap: 1.5rem;
		margin-top: 2rem;
	}

	.about__story p {
		margin: 0 0 1rem;
		color: var(--text-soft);
	}

	.about__education {
		padding: 1.3rem;
		border: 1px solid var(--line);
		border-radius: var(--radius-md);
		background: rgba(255, 255, 255, 0.55);
	}

	.about__education h3,
	.about__highlight h3 {
		margin: 0;
		font-size: 1.15rem;
	}

	.about__education article + article {
		margin-top: 1.15rem;
		padding-top: 1.15rem;
		border-top: 1px solid rgba(45, 32, 23, 0.08);
	}

	.about__education strong,
	.about__education span,
	.about__education small {
		display: block;
	}

	.about__education span {
		margin-top: 0.2rem;
		color: var(--text-soft);
	}

	.about__education small {
		margin-top: 0.35rem;
		color: var(--accent-deep);
		font-weight: 600;
	}

	.about__highlights {
		display: grid;
		grid-template-columns: repeat(4, minmax(0, 1fr));
		gap: 1rem;
		margin-top: 1.6rem;
	}

	.about__highlight {
		padding: 1.25rem;
		border: 1px solid rgba(45, 32, 23, 0.08);
		border-radius: var(--radius-md);
		background: rgba(255, 255, 255, 0.66);
	}

	.about__highlight p:not(.eyebrow) {
		margin: 0.7rem 0 0;
		color: var(--text-soft);
	}

	.about__highlight a {
		display: inline-flex;
		margin-top: 0.85rem;
		color: var(--accent-deep);
		font-weight: 700;
	}

	@media (max-width: 980px) {
		.about__top,
		.about__highlights {
			grid-template-columns: 1fr;
		}
	}
</style>
