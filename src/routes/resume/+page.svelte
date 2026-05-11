<script lang="ts">
	import { reveal } from '$lib/actions/reveal';
	import portfolio from '$lib/data/portfolio';

	let activeFilters = $state<Set<string>>(new Set());

	const filterGroups = [
		{
			label: 'Languages',
			items: ['TypeScript', 'JavaScript', 'Python', 'C#', 'Visual Basic .NET']
		},
		{
			label: 'Mobile',
			items: ['SwiftUI', 'React Native', 'iOS development', 'mobile development']
		},
		{
			label: 'Web',
			items: ['Next.js', 'Express', 'React Native']
		},
		{
			label: 'Backend',
			items: [
				'Node.js',
				'.NET',
				'Go Fiber',
				'REST APIs',
				'BetterAuth',
				'Auth0',
				'Docker',
				'AWS EC2',
				'CI/CD'
			]
		},
		{
			label: 'Data',
			items: ['PostgreSQL', 'MySQL', 'MongoDB', 'Entity Framework', 'pandas']
		},
		{
			label: 'AI/ML',
			items: [
				'NLP',
				'Neural networks',
				'LLM embeddings',
				'RAG iteration',
				'Computer vision',
				'OpenCV',
				'YOLOv8',
				'scikit-learn',
				'Hugging Face Spaces',
				'Kaggle',
				'Bayesian optimization',
				'Recommendation modeling',
				'classification',
				'model evaluation',
				'model tuning'
			]
		},
		{
			label: 'Engineering',
			items: [
				'API design',
				'backend engineering',
				'auth systems',
				'developer tooling',
				'performance optimization',
				'UX refinement',
				'full-stack debugging',
				'full-stack development',
				'cost optimization',
				'feature delivery',
				'enterprise support',
				'automation',
				'database tuning',
				'experimentation',
				'rapid prototyping',
				'algorithms',
				'problem solving',
				'application hardening'
			]
		},
		{
			label: 'Leadership',
			items: [
				'technical communication',
				'technical leadership',
				'startup execution',
				'product translation',
				'interview readiness',
				'mentorship sessions',
				'fellowship collaboration'
			]
		}
	];

	const filteredExperiences = $derived(
		activeFilters.size === 0
			? portfolio.resume.experiences
			: portfolio.resume.experiences.filter((exp) => {
					const expItems = [...exp.tools, ...exp.skills];
					return [...activeFilters].some((filter) => expItems.includes(filter));
				})
	);

	function toggleFilter(filter: string) {
		if (activeFilters.has(filter)) {
			activeFilters.delete(filter);
		} else {
			activeFilters.add(filter);
		}
		activeFilters = new Set(activeFilters);
	}

	function clearFilters() {
		activeFilters = new Set();
	}
</script>

<svelte:head>
	<title>{portfolio.resume.title} | {portfolio.profile.name}</title>
	<meta name="description" content={portfolio.resume.description} />
</svelte:head>

<main class="cv-page">
	<header class="page-header reveal" use:reveal>
		<p class="tech-label">Resume</p>
		<h1>{portfolio.resume.title}</h1>
		<p>{portfolio.resume.description}</p>
	</header>

	<section class="cv-layout" aria-label="Curriculum vitae">
		<aside class="filters">
			<div class="filter-header">
				<span class="tech-label">FILTERS</span>
				{#if activeFilters.size > 0}
					<button class="clear-btn" onclick={clearFilters}>Clear</button>
				{/if}
			</div>
			{#each filterGroups as group}
				<div class="filter-group">
					<span class="filter-group-label">{group.label}</span>
					<div class="filter-list">
						{#each group.items as item}
							<button
								class="filter-chip {activeFilters.has(item) ? 'active' : ''}"
								onclick={() => toggleFilter(item)}
							>
								{item}
							</button>
						{/each}
					</div>
				</div>
			{/each}
		</aside>
		<div class="timeline">
			{#key JSON.stringify([...activeFilters].sort())}
				{#each filteredExperiences as experience, index (experience.company + experience.role)}
				<article class="experience-card reveal" use:reveal data-testid="experience-card">
					<div class="experience-date tech-label">{experience.start} - {experience.end}</div>
					<div class="experience-content">
						<div class="experience-heading">
							{#if experience.logo}
								<div class="company-logo">
									<img src={`/${experience.logo}`} alt="{experience.company} logo" />
								</div>
							{/if}
							<div>
								<span class="experience-index tech-label">No. {(index + 1).toString().padStart(2, '0')}</span>
								<h2>{experience.role}</h2>
								<p>
									{#if experience.companyUrl}
										<a href={experience.companyUrl} target="_blank" rel="noreferrer">{experience.company}</a>
									{:else}
										<span>{experience.company}</span>
									{/if}
									<span>{experience.location}</span>
								</p>
							</div>
						</div>

						{#if experience.summary}
							<p class="summary">{experience.summary}</p>
						{/if}

						<ul>
							{#each experience.highlights as highlight}
								<li>{highlight}</li>
							{/each}
						</ul>
					</div>
				</article>
			{/each}
			{/key}

			{#if filteredExperiences.length === 0}
				<p class="empty-state">No experience matches selected filters.</p>
			{/if}
		</div>
	</section>
</main>

<style>
	.cv-page {
		max-width: var(--container);
		margin: 0 auto;
		padding: clamp(4rem, 8vw, var(--section-gap)) var(--safe-margin);
	}

	.page-header {
		display: grid;
		grid-template-columns: repeat(12, minmax(0, 1fr));
		gap: var(--gutter);
		padding-bottom: var(--section-gap);
	}

	.page-header .tech-label {
		grid-column: 2 / span 2;
		color: var(--secondary);
	}

	.page-header h1 {
		grid-column: 4 / span 5;
		font-size: clamp(3rem, 6vw, 5.25rem);
	}

	.page-header > p:last-child {
		grid-column: 4 / span 6;
		color: var(--on-surface-variant);
		font-size: 1.08rem;
		border-left: 1px solid var(--hairline-strong);
		padding-left: 1.5rem;
	}

	.cv-layout {
		display: grid;
		grid-template-columns: repeat(12, minmax(0, 1fr));
		gap: var(--gutter);
	}

	.filters {
		grid-column: 9 / span 4;
		grid-row: 1;
		align-self: start;
		position: sticky;
		top: 0;
		max-height: calc(100vh - 2rem);
		overflow-y: auto;
		padding-right: 0.5rem;
	}

	.filters::-webkit-scrollbar {
		width: 4px;
	}

	.filters::-webkit-scrollbar-track {
		background: transparent;
	}

	.filters::-webkit-scrollbar-thumb {
		background: var(--hairline);
		border-radius: 2px;
	}

	.filters::-webkit-scrollbar-thumb:hover {
		background: var(--secondary);
	}

	.filter-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 1rem;
		padding-bottom: 0.75rem;
		border-bottom: 1px solid var(--hairline);
	}

	.clear-btn {
		background: none;
		border: none;
		color: var(--primary);
		cursor: pointer;
		font-family: inherit;
		font-size: 0.85rem;
		padding: 0;
	}

	.clear-btn:hover {
		text-decoration: underline;
	}

	.filter-group {
		margin-bottom: 1.25rem;
	}

	.filter-group-label {
		display: block;
		font-size: 0.7rem;
		color: var(--secondary);
		margin-bottom: 0.5rem;
		text-transform: uppercase;
		letter-spacing: 0.05em;
	}

	.filter-list {
		display: flex;
		flex-wrap: wrap;
		gap: 0.35rem;
	}

	.filter-chip {
		background: var(--surface-container-low);
		border: 1px solid var(--hairline);
		color: var(--on-surface);
		padding: 0.25rem 0.5rem;
		font-size: 0.7rem;
		cursor: pointer;
		transition: all 150ms ease;
		white-space: nowrap;
	}

	.filter-chip:hover {
		border-color: var(--primary);
	}

	.filter-chip.active {
		background: var(--primary);
		color: var(--on-primary);
		border-color: var(--primary);
	}

	.timeline {
		grid-column: 1 / span 8;
		grid-row: 1;
		align-self: start;
		border-left: 1px solid var(--hairline-strong);
	}

	.experience-card {
		display: grid;
		grid-template-columns: 180px minmax(0, 1fr);
		gap: var(--gutter);
		position: relative;
		padding: 0 0 3.5rem 2rem;
		transition:
			opacity 240ms ease,
			transform 240ms ease;
	}

	.experience-card::before {
		content: '';
		position: absolute;
		top: 0.45rem;
		left: -5px;
		width: 9px;
		height: 9px;
		background: var(--primary);
	}

	.experience-date,
	.experience-index {
		color: var(--secondary);
	}

	.experience-content {
		padding-bottom: 2rem;
		border-bottom: 1px solid var(--hairline);
	}

	.experience-card:last-child,
	.experience-card:last-child .experience-content {
		padding-bottom: 0;
		border-bottom: 0;
	}

	.experience-heading {
		display: flex;
		gap: 1rem;
	}

	.company-logo {
		width: 48px;
		height: 48px;
		display: grid;
		place-items: center;
		flex: 0 0 48px;
		border: 1px solid var(--hairline);
		background: var(--surface-container-lowest);
	}

	.company-logo img {
		max-width: 32px;
		max-height: 32px;
		object-fit: contain;
		filter: grayscale(100%);
	}

	:global(:root[data-theme='dark']) .company-logo img {
		filter: grayscale(100%) brightness(1.35);
	}

	@media (prefers-color-scheme: dark) {
		:global(:root:not([data-theme])) .company-logo img {
			filter: grayscale(100%) brightness(1.35);
		}
	}

	.experience-content h2 {
		margin-top: 0.6rem;
		font-size: clamp(1.55rem, 3vw, 2rem);
	}

	.experience-content p {
		color: var(--on-surface-variant);
	}

	.experience-content a {
		color: var(--primary);
	}

	.experience-heading p {
		display: flex;
		flex-wrap: wrap;
		gap: 0.8rem;
		margin-top: 0.4rem;
	}

	.summary {
		margin-top: 1.25rem;
	}

	ul {
		display: grid;
		gap: 0.7rem;
		margin: 1.3rem 0 0;
		padding: 0;
		list-style: none;
	}

	li {
		position: relative;
		padding-left: 1.25rem;
	}

	li::before {
		content: '';
		position: absolute;
		left: 0;
		top: 0.75em;
		width: 0.45rem;
		height: 1px;
		background: var(--secondary);
	}

	.empty-state {
		padding-left: 2rem;
		color: var(--on-surface-variant);
	}

	@media (max-width: 960px) {
		.page-header .tech-label,
		.page-header h1,
		.page-header > p:last-child,
		.filters,
		.timeline {
			grid-column: 1 / -1;
		}

		.filters {
			position: static;
			max-height: none;
			margin-bottom: 2rem;
		}
	}

	@media (max-width: 680px) {
		.experience-card {
			grid-template-columns: 1fr;
			gap: 0.8rem;
			padding-left: 1.25rem;
		}
	}
</style>