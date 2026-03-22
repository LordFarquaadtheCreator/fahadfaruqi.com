<script lang="ts">
	import { reveal } from '$lib/actions/reveal';
	import type { ResumeData } from '$lib/data/portfolio';
	import MediaFrame from './MediaFrame.svelte';
	import SectionHeading from './SectionHeading.svelte';

	let { resume } = $props<{ resume: ResumeData }>();
</script>

<section class="resume section-shell" id="resume">
	<div class="resume__header reveal" use:reveal>
		<SectionHeading
			eyebrow="Resume"
			title="Experience across startups, product teams, civic software, and AI fellowships."
			copy="Every entry below is driven from the JSON source of truth so you can revise copy, imagery, and skills without touching the Svelte components."
		/>

		{#if resume.downloadUrl}
			<a class="button button--ghost" href={resume.downloadUrl} target="_blank" rel="noreferrer">
				Download resume
			</a>
		{/if}
	</div>

	<ol class="resume__timeline">
		{#each resume.experiences as experience}
			<li class="resume__item">
				<article class="surface resume__card reveal" use:reveal data-testid="experience-card">
					<div class="resume__identity">
						<div class="resume__logo">
							<MediaFrame
								src={experience.logo}
								alt={`${experience.company} logo`}
								label={experience.company}
							/>
						</div>

						<div class="resume__meta">
							<p class="resume__dates">{experience.start} - {experience.end}</p>
							<h3>{experience.role}</h3>
							<p class="resume__company">
								{#if experience.companyUrl}
									<a href={experience.companyUrl} target="_blank" rel="noreferrer">
										{experience.company}
									</a>
								{:else}
									<span>{experience.company}</span>
								{/if}
								<span>· {experience.location}</span>
							</p>
							{#if experience.summary}
								<p class="resume__summary">{experience.summary}</p>
							{/if}
						</div>
					</div>

					<ul class="resume__highlights">
						{#each experience.highlights as highlight}
							<li>{highlight}</li>
						{/each}
					</ul>

					<div class="resume__tags">
						<div>
							<h4>Tools</h4>
							<div class="resume__chips">
								{#each experience.tools as tool}
									<span class="chip">{tool}</span>
								{/each}
							</div>
						</div>

						<div>
							<h4>Skills</h4>
							<div class="resume__chips">
								{#each experience.skills as skill}
									<span class="chip">{skill}</span>
								{/each}
							</div>
						</div>
					</div>
				</article>
			</li>
		{/each}
	</ol>
</section>

<style>
	.resume {
		padding: 2.5rem 0;
	}

	.resume__header {
		display: flex;
		flex-wrap: wrap;
		justify-content: space-between;
		gap: 1rem;
		align-items: end;
	}

	.resume__timeline {
		position: relative;
		padding: 2rem 0 0 1.1rem;
		margin: 0;
		list-style: none;
	}

	.resume__timeline::before {
		content: '';
		position: absolute;
		left: 0.22rem;
		top: 2.6rem;
		bottom: 1rem;
		width: 1px;
		background: rgba(45, 32, 23, 0.16);
	}

	.resume__item {
		position: relative;
		padding-left: 1.5rem;
	}

	.resume__item + .resume__item {
		margin-top: 1.3rem;
	}

	.resume__item::before {
		content: '';
		position: absolute;
		top: 2rem;
		left: -0.16rem;
		width: 0.8rem;
		height: 0.8rem;
		border: 2px solid rgba(179, 92, 46, 0.5);
		border-radius: 50%;
		background: var(--bg-soft);
		box-shadow: 0 0 0 6px rgba(246, 239, 230, 0.92);
	}

	.resume__card {
		padding: clamp(1.2rem, 2vw, 1.7rem);
	}

	.resume__identity {
		display: grid;
		grid-template-columns: minmax(10rem, 13rem) 1fr;
		gap: 1.15rem;
		align-items: start;
	}

	.resume__logo {
		max-width: 13rem;
	}

	.resume__dates {
		margin: 0;
		color: var(--accent-deep);
		font-size: 0.83rem;
		font-weight: 700;
		letter-spacing: 0.14em;
		text-transform: uppercase;
	}

	.resume__meta h3 {
		margin: 0.65rem 0 0;
		font-size: clamp(1.45rem, 2vw, 1.8rem);
		line-height: 1.05;
	}

	.resume__company {
		display: flex;
		flex-wrap: wrap;
		gap: 0.4rem;
		margin: 0.55rem 0 0;
		font-weight: 600;
		color: var(--text-soft);
	}

	.resume__company a {
		color: var(--text);
	}

	.resume__summary {
		margin: 0.85rem 0 0;
		color: var(--text-soft);
	}

	.resume__highlights {
		display: grid;
		gap: 0.75rem;
		padding-left: 1rem;
		margin: 1.25rem 0 0;
		color: var(--text-soft);
	}

	.resume__tags {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 1rem;
		margin-top: 1.35rem;
	}

	.resume__tags h4 {
		margin: 0 0 0.7rem;
		font-size: 0.86rem;
		letter-spacing: 0.12em;
		text-transform: uppercase;
		color: var(--accent-deep);
	}

	.resume__chips {
		display: flex;
		flex-wrap: wrap;
		gap: 0.6rem;
	}

	@media (max-width: 860px) {
		.resume__identity,
		.resume__tags {
			grid-template-columns: 1fr;
		}

		.resume__logo {
			max-width: 100%;
		}
	}

	@media (max-width: 640px) {
		.resume__timeline {
			padding-left: 0.9rem;
		}

		.resume__item {
			padding-left: 1.1rem;
		}
	}
</style>
