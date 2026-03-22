<script lang="ts">
	import { reveal } from '$lib/actions/reveal';
	import type { ProfileData } from '$lib/data/portfolio';

	let {
		profile,
		resumeDownloadUrl
	} = $props<{
		profile: ProfileData;
		resumeDownloadUrl?: string;
	}>();
</script>

<section class="hero section-shell reveal" id="top" use:reveal>
	<div class="hero__copy">
		<p class="eyebrow">{profile.location}</p>
		<h1>{profile.name}</h1>
		<p class="hero__title">{profile.title}</p>
		<p class="hero__lede">{profile.heroIntro}</p>
		<p class="hero__supporting">{profile.supportingIntro}</p>

		<div class="hero__actions">
			<a class="button" href="#resume">Explore the resume</a>
			<a class="button button--ghost" href="#contact">Start a conversation</a>
			{#if resumeDownloadUrl}
				<a class="button button--ghost" href={resumeDownloadUrl} target="_blank" rel="noreferrer">
					Download resume
				</a>
			{/if}
		</div>

		<ul class="hero__socials">
			{#each profile.socialLinks as link}
				<li><a href={link.href} target={link.href.startsWith('http') ? '_blank' : undefined} rel="noreferrer">{link.label}</a></li>
			{/each}
		</ul>
	</div>

	<div class="hero__panel surface">
		<div class="hero__panel-copy">
			<p class="hero__panel-label">Selected signals</p>
			<h2>Range without losing craft.</h2>
			<p>
				I like shipping the visible parts of a product and the invisible systems that make them
				feel effortless.
			</p>
		</div>

		<ul class="hero__metrics">
			{#each profile.metrics as metric}
				<li>
					<strong>{metric.value}</strong>
					<span>{metric.label}</span>
					<small>{metric.detail}</small>
				</li>
			{/each}
		</ul>
	</div>
</section>

<style>
	.hero {
		display: grid;
		grid-template-columns: minmax(0, 1.25fr) minmax(19rem, 0.92fr);
		gap: 1.5rem;
		padding: 7rem 0 4rem;
		align-items: end;
	}

	.hero__copy {
		padding: 2rem 0 0.75rem;
	}

	.hero h1 {
		margin: 0.6rem 0 0;
		font-size: clamp(3.8rem, 11vw, 7.3rem);
		line-height: 0.92;
		letter-spacing: -0.07em;
	}

	.hero__title {
		margin: 1rem 0 0;
		font-family: 'Source Serif 4', serif;
		font-size: clamp(1.45rem, 3vw, 2.35rem);
		color: var(--accent-deep);
	}

	.hero__lede,
	.hero__supporting {
		max-width: 40rem;
		color: var(--text-soft);
	}

	.hero__lede {
		margin: 1.2rem 0 0;
		font-size: 1.08rem;
	}

	.hero__supporting {
		margin: 0.7rem 0 0;
	}

	.hero__actions {
		display: flex;
		flex-wrap: wrap;
		gap: 0.85rem;
		margin-top: 2rem;
	}

	.hero__socials {
		display: flex;
		flex-wrap: wrap;
		gap: 1rem;
		padding: 0;
		margin: 1.8rem 0 0;
		list-style: none;
	}

	.hero__socials a {
		font-size: 0.95rem;
		font-weight: 600;
		color: var(--text-soft);
	}

	.hero__panel {
		padding: 1.5rem;
	}

	.hero__panel-copy {
		padding-bottom: 1rem;
		border-bottom: 1px solid var(--line);
	}

	.hero__panel-label {
		margin: 0;
		font-size: 0.8rem;
		font-weight: 700;
		letter-spacing: 0.16em;
		text-transform: uppercase;
		color: var(--accent-deep);
	}

	.hero__panel h2 {
		margin: 0.75rem 0 0;
		font-size: clamp(1.8rem, 4vw, 2.6rem);
		line-height: 1;
		letter-spacing: -0.05em;
	}

	.hero__panel p {
		margin: 0.9rem 0 0;
		color: var(--text-soft);
	}

	.hero__metrics {
		display: grid;
		gap: 1rem;
		padding: 1.25rem 0 0;
		margin: 0;
		list-style: none;
	}

	.hero__metrics li {
		padding: 1rem 0;
		border-bottom: 1px solid rgba(45, 32, 23, 0.08);
	}

	.hero__metrics li:last-child {
		border-bottom: 0;
		padding-bottom: 0;
	}

	.hero__metrics strong {
		display: block;
		font-size: 2rem;
		line-height: 1;
	}

	.hero__metrics span {
		display: block;
		margin-top: 0.3rem;
		font-weight: 700;
		text-transform: lowercase;
	}

	.hero__metrics small {
		display: block;
		margin-top: 0.35rem;
		color: var(--text-soft);
		font-size: 0.92rem;
	}

	@media (max-width: 900px) {
		.hero {
			grid-template-columns: 1fr;
			padding-top: 5rem;
		}

		.hero__copy {
			padding-top: 0.5rem;
		}
	}
</style>
