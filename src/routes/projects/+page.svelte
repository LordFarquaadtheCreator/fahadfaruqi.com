<script lang="ts">
	import GothicShell from '$lib/components/GothicShell.svelte';
	import portfolio from '$lib/data/portfolio';
</script>

<svelte:head>
	<title>Projects | Fahad Faruqi</title>
</svelte:head>

<GothicShell title="Fahad Faruqi">
	{#snippet children()}
		<div class="projects-container">
			<div class="projects-header">
				<span class="eyebrow">WORK_LOG</span>
				<h2 class="projects-title chromatic">Experience</h2>
				<p class="projects-desc">Production work across fintech, civic tech, startups, and fellowships.</p>
			</div>

			<div class="projects-grid">
				{#each portfolio.resume.experiences as exp}
					<article class="project-card" data-testid="project-card">
						<div class="card-top">
							<div class="card-company-row">
								{#if exp.logo}
									<img
										class="company-logo"
										src="/{exp.logo}"
										alt="{exp.company} logo"
										loading="lazy"
									/>
								{/if}
								<div class="card-company-info">
									{#if exp.companyUrl}
										<a
											href={exp.companyUrl}
											class="company-name"
											target="_blank"
											rel="noopener noreferrer"
										>{exp.company}</a>
									{:else}
										<span class="company-name">{exp.company}</span>
									{/if}
									<span class="card-dates tech-label">{exp.start} – {exp.end}</span>
								</div>
							</div>
							<h3 class="card-role">{exp.role}</h3>
							{#if exp.summary}
								<p class="card-summary">{exp.summary}</p>
							{/if}
						</div>

						<div class="card-divider"></div>

						<div class="card-bottom">
							<div class="tools-row">
								{#each exp.tools as tool}
									<span class="chip">{tool}</span>
								{/each}
							</div>
						</div>
					</article>
				{/each}
			</div>
		</div>
	{/snippet}
</GothicShell>

<style>
	.projects-container {
		width: min(1200px, calc(100% - 2rem));
		margin: 0 auto;
		padding: 0 1rem;
	}

	.projects-header {
		padding: 2rem 0 2.5rem;
		border-bottom: 1px solid rgba(157, 78, 221, 0.2);
		margin-bottom: 2.5rem;
	}

	.projects-title {
		font-family: 'Pirata One', cursive;
		font-size: clamp(2.5rem, 6vw, 4.5rem);
		font-weight: 400;
		color: #c084fc;
		margin: 0.25rem 0 0.75rem;
		line-height: 1;
	}

	.projects-desc {
		font-family: 'Work Sans', sans-serif;
		font-size: 0.9rem;
		color: var(--on-surface-variant);
		opacity: 0.6;
		margin: 0;
	}

	.projects-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
		gap: 1px;
		background: rgba(157, 78, 221, 0.15);
		border: 1px solid rgba(157, 78, 221, 0.15);
	}

	.project-card {
		background: #09090b;
		padding: 1.5rem;
		display: flex;
		flex-direction: column;
		gap: 1rem;
		transition: background 150ms ease;
	}

	.project-card:hover {
		background: rgba(157, 78, 221, 0.04);
	}

	.card-top {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
		flex: 1;
	}

	.card-company-row {
		display: flex;
		align-items: center;
		gap: 0.75rem;
	}

	.company-logo {
		width: 28px;
		height: 28px;
		object-fit: contain;
		filter: grayscale(1) brightness(1.5);
		opacity: 0.7;
	}

	.card-company-info {
		display: flex;
		flex-direction: column;
		gap: 0.1rem;
	}

	.company-name {
		font-family: 'Space Grotesk', sans-serif;
		font-size: 0.8rem;
		font-weight: 600;
		color: #c084fc;
		letter-spacing: 0.05em;
		text-transform: uppercase;
		text-decoration: none;
	}

	a.company-name:hover {
		text-decoration: underline;
		text-decoration-color: rgba(157, 78, 221, 0.5);
	}

	.card-dates {
		color: rgba(157, 78, 221, 0.5);
		font-size: 0.5rem;
	}

	.card-role {
		font-family: 'Work Sans', sans-serif;
		font-size: 0.95rem;
		font-weight: 500;
		color: var(--on-surface);
		margin: 0;
	}

	.card-summary {
		font-family: 'Work Sans', sans-serif;
		font-size: 0.78rem;
		color: var(--on-surface-variant);
		opacity: 0.55;
		line-height: 1.5;
		margin: 0;
	}

	.card-divider {
		height: 1px;
		background: rgba(157, 78, 221, 0.15);
	}

	.card-bottom {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}

	.tools-row {
		display: flex;
		flex-wrap: wrap;
		gap: 0.35rem;
	}
</style>
