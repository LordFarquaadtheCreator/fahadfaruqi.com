<script lang="ts">
	import { reveal } from '$lib/actions/reveal';
	import type { ResumeData } from '$lib/data/portfolio';
	import TacticalSectionHeader from './TacticalSectionHeader.svelte';

	let { resume } = $props<{ resume: ResumeData }>();
</script>

<section class="ops-log" id="resume">
	<TacticalSectionHeader
		ghostText="Resume"
		eyebrow=""
		title="Resume"
		description=""
	>
		{#snippet headerRight()}
			<div class="header-meta-group">
				<span class="tech-label" style="color: var(--tertiary);">ALL_SYSTEMS_NOMINAL</span>
				{#if resume.downloadUrl}
					<a href={resume.downloadUrl} class="download-btn" target="_blank" rel="noreferrer">
						<span class="material-symbols-outlined">download</span>
						<span class="tech-label">DOWNLOAD_DAT</span>
					</a>
				{/if}
			</div>
		{/snippet}
	</TacticalSectionHeader>

	<!-- Operations Grid -->
	<div class="ops-grid">
		{#each resume.experiences as experience, i}
			<article class="op-card reveal" use:reveal data-testid="experience-card">
				<div class="op-header">
					<div class="op-meta">
						<span class="tech-label op-index">OP_{(i + 1).toString().padStart(3, '0')}</span>
						<span class="op-date">{experience.start} — {experience.end}</span>
					</div>
					<div class="status-indicator">
						{#if experience.end === 'Present'}
							<span class="status-pip status-active"></span>
							<span class="tech-label" style="color: var(--tertiary);">ACTIVE</span>
						{:else}
							<span class="status-pip status-idle"></span>
							<span class="tech-label" style="color: var(--outline);">ARCHIVED</span>
						{/if}
					</div>
				</div>

				<div class="op-body">
					<div class="op-identity">
						{#if experience.logo}
							<div class="op-logo">
								<img src={experience.logo} alt="{experience.company} logo" />
							</div>
						{/if}
						<div class="op-titles">
							<h3 class="op-role">{experience.role}</h3>
							<div class="op-org">
								{#if experience.companyUrl}
									<a href={experience.companyUrl} target="_blank" rel="noreferrer">
										{experience.company}
									</a>
								{:else}
									<span>{experience.company}</span>
								{/if}
								<span class="op-location">· {experience.location}</span>
							</div>
						</div>
					</div>

					{#if experience.summary}
						<p class="op-summary">{experience.summary}</p>
					{/if}

					<div class="op-highlights">
						{#each experience.highlights as highlight}
							<div class="highlight-item">
								<span class="material-symbols-outlined bullet">chevron_right</span>
								<span>{highlight}</span>
							</div>
						{/each}
					</div>

					<div class="op-tags">
						<div class="tag-group">
							<span class="tech-label tag-label">TOOLS</span>
							<div class="tag-list">
								{#each experience.tools as tool}
									<span class="chip">{tool}</span>
								{/each}
							</div>
						</div>
						<div class="tag-group">
							<span class="tech-label tag-label">SKILLS</span>
							<div class="tag-list">
								{#each experience.skills as skill}
									<span class="chip">{skill}</span>
								{/each}
							</div>
						</div>
					</div>
				</div>

				<div class="hud-bracket-bl"></div>
				<div class="hud-bracket-br"></div>
			</article>
		{/each}
	</div>
</section>

<style>
	.ops-log {
		padding: 4rem 0;
		position: relative;
		overflow: hidden;
	}

	.header-meta-group {
		display: flex;
		align-items: center;
		gap: 1rem;
	}

	.download-btn {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		padding: 0.5rem 1rem;
		background: var(--surface-container-high);
		border: 1px solid var(--outline-variant);
		color: var(--on-surface);
		transition: all 150ms ease;
	}

	.download-btn:hover {
		border-color: var(--primary);
		color: var(--primary);
	}

	.download-btn .material-symbols-outlined {
		font-size: 18px;
	}

	.ops-grid {
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.op-card {
		background: var(--surface-container-low);
		border: 1px solid rgba(175, 135, 134, 0.15);
		position: relative;
		overflow: hidden;
		transition: all 200ms ease;
	}

	.op-card:hover {
		border-color: rgba(211, 251, 255, 0.4);
		box-shadow: 0 0 30px rgba(211, 251, 255, 0.08);
		animation: glitch 0.3s ease-in-out;
	}

	@keyframes glitch {
		0% { transform: translate(0); }
		20% { transform: translate(-1px, 1px); }
		40% { transform: translate(1px, -1px); }
		60% { transform: translate(-1px, -1px); }
		80% { transform: translate(1px, 1px); }
		100% { transform: translate(0); }
	}

	.op-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 1rem 1.5rem;
		background: var(--surface-container);
		border-bottom: 1px solid rgba(175, 135, 134, 0.1);
	}

	.op-meta {
		display: flex;
		align-items: center;
		gap: 1rem;
	}

	.op-index {
		color: var(--secondary);
		font-weight: 700;
	}

	.op-date {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.75rem;
		color: var(--on-surface);
	}

	.status-indicator {
		display: flex;
		align-items: center;
		gap: 0.5rem;
	}

	.op-body {
		padding: 1.5rem;
	}

	.op-identity {
		display: flex;
		gap: 1rem;
		align-items: flex-start;
		margin-bottom: 1rem;
	}

	.op-logo {
		width: 48px;
		height: 48px;
		background: var(--surface-container-high);
		display: flex;
		align-items: center;
		justify-content: center;
		flex-shrink: 0;
	}

	.op-logo img {
		max-width: 32px;
		max-height: 32px;
		filter: grayscale(100%) brightness(1.2);
	}

	.op-titles {
		flex: 1;
	}

	.op-role {
		font-size: 1.35rem;
		font-weight: 700;
		color: var(--on-surface);
		margin: 0 0 0.25rem;
	}

	.op-org {
		font-size: 0.9rem;
		color: var(--on-surface-variant);
		display: flex;
		gap: 0.5rem;
		flex-wrap: wrap;
	}

	.op-org a {
		color: var(--primary);
		transition: color 150ms ease;
	}

	.op-org a:hover {
		color: var(--primary-container);
	}

	.op-location {
		color: var(--outline);
	}

	.op-summary {
		color: var(--on-surface-variant);
		font-size: 0.9rem;
		margin-bottom: 1rem;
		padding-bottom: 1rem;
		border-bottom: 1px solid rgba(175, 135, 134, 0.1);
	}

	.op-highlights {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
		margin-bottom: 1.5rem;
	}

	.highlight-item {
		display: flex;
		align-items: flex-start;
		gap: 0.5rem;
		color: var(--on-surface-variant);
		font-size: 0.85rem;
		line-height: 1.6;
	}

	.highlight-item .bullet {
		font-size: 16px;
		color: var(--secondary);
		flex-shrink: 0;
		margin-top: 0.1rem;
	}

	.op-tags {
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		gap: 1rem;
		padding-top: 1rem;
		border-top: 1px solid rgba(175, 135, 134, 0.1);
	}

	.tag-group {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}

	.tag-label {
		color: var(--primary);
	}

	.tag-list {
		display: flex;
		flex-wrap: wrap;
		gap: 0.5rem;
	}

	/* Responsive */
	@media (max-width: 768px) {
		.op-header {
			flex-direction: column;
			gap: 0.75rem;
			align-items: flex-start;
		}

		.op-meta {
			flex-direction: column;
			align-items: flex-start;
			gap: 0.25rem;
		}

		.op-identity {
			flex-direction: column;
		}

		.op-tags {
			grid-template-columns: 1fr;
		}
	}
</style>
