<script lang="ts">
	import { reveal } from '$lib/actions/reveal';
	import type { AboutData } from '$lib/data/portfolio';
	import TacticalCard from './TacticalCard.svelte';
	import TacticalSectionHeader from './TacticalSectionHeader.svelte';

	let { about } = $props<{ about: AboutData }>();
</script>

<section class="intel-section" id="about">
	<TacticalSectionHeader
		ghostText="INTEL"
		eyebrow="SRC_DIR: /ROOT/PERSONNEL/"
		title="OPERATOR_DOSSIER"
		description="Technical specifications and operational history. Decrypted for authorized review."
	>
		{#snippet headerRight()}
			<span class="tech-label classified-badge">CLASSIFIED//NOFORN</span>
		{/snippet}
	</TacticalSectionHeader>

	<div class="intel-grid">
		<!-- Bio Panel -->
		<TacticalCard
			hudBrackets
			headerIcon="person_search"
			headerLabel="BIOGRAPHICAL_DATA"
			class="intel-card--bio"
		>
			<div class="bio-content">
				{#each about.paragraphs as paragraph, i}
					<p class="bio-text">
						<span class="line-number">{(i + 1).toString().padStart(2, '0')}</span>
						{paragraph}
					</p>
				{/each}
			</div>
		</TacticalCard>

		<!-- Education Panel -->
		<TacticalCard
			headerIcon="school"
			headerLabel="EDUCATION_RECORDS"
			class="intel-card--education"
		>
			<div class="education-list">
				{#each about.education as item}
					<div class="education-item">
						<div class="edu-header">
							<strong class="edu-institution">{item.institution}</strong>
							<span class="tech-label edu-dates">{item.dates}</span>
						</div>
						<span class="edu-credential">{item.credential}</span>
						<span class="tech-label edu-location">{item.location}</span>
					</div>
				{/each}
			</div>
		</TacticalCard>

		<!-- Highlights Grid -->
		{#each about.highlights as item, i}
			<TacticalCard
				variant="highlight"
				class="intel-card--highlight"
			>
				{#snippet headerRight()}
					<span class="tech-label" style="color: var(--tertiary);">
						REF_{(i + 1).toString().padStart(2, '0')}_{item.eyebrow.toUpperCase().slice(0, 4)}
					</span>
				{/snippet}
				<h3 class="highlight-title">{item.title}</h3>
				<p class="highlight-desc">{item.description}</p>
				{#if item.href && item.linkLabel}
					<a href={item.href} class="highlight-link" target="_blank" rel="noreferrer">
						<span class="link-text">{item.linkLabel.toUpperCase()}</span>
						<span class="material-symbols-outlined">arrow_outward</span>
					</a>
				{/if}
			</TacticalCard>
		{/each}
	</div>
</section>

<style>
	.intel-section {
		padding: 4rem 0;
		position: relative;
		overflow: hidden;
	}

	.classified-badge {
		color: var(--outline);
	}

	:global(.intel-card--bio) {
		grid-column: span 7;
	}

	:global(.intel-card--education) {
		grid-column: span 5;
		background: var(--surface-container) !important;
	}

	:global(.intel-card--highlight) {
		grid-column: span 3;
	}

	.intel-grid {
		display: grid;
		grid-template-columns: repeat(12, 1fr);
		gap: 1rem;
	}

	.bio-content {
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.bio-text {
		color: var(--on-surface-variant);
		font-size: 0.9rem;
		line-height: 1.7;
		display: flex;
		gap: 1rem;
	}

	.line-number {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.65rem;
		color: var(--outline);
		flex-shrink: 0;
		min-width: 1.5rem;
	}

	.education-list {
		display: flex;
		flex-direction: column;
		gap: 1.25rem;
	}

	.education-item {
		display: flex;
		flex-direction: column;
		gap: 0.35rem;
		padding: 1rem;
		background: var(--surface-container-lowest);
		border-left: 3px solid var(--secondary);
	}

	.edu-header {
		display: flex;
		justify-content: space-between;
		align-items: flex-start;
		gap: 1rem;
	}

	.edu-institution {
		font-size: 0.95rem;
		color: var(--on-surface);
	}

	.edu-dates {
		color: var(--secondary);
		flex-shrink: 0;
	}

	.edu-credential {
		font-size: 0.85rem;
		color: var(--on-surface-variant);
	}

	.edu-location {
		color: var(--outline);
	}

	.highlight-title {
		font-size: 1.1rem;
		font-weight: 700;
		margin-bottom: 0.75rem;
		color: var(--on-surface);
	}

	.highlight-desc {
		font-size: 0.85rem;
		color: var(--on-surface-variant);
		line-height: 1.6;
		margin-bottom: 1rem;
	}

	.highlight-link {
		display: inline-flex;
		align-items: center;
		gap: 0.5rem;
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.65rem;
		letter-spacing: 0.05em;
		color: var(--primary);
		transition: color 150ms ease;
	}

	.highlight-link:hover {
		color: var(--primary-container);
	}

	.highlight-link .material-symbols-outlined {
		font-size: 14px;
	}

	/* Responsive */
	@media (max-width: 1024px) {
		:global(.intel-card--bio),
		:global(.intel-card--education) {
			grid-column: span 6;
		}

		:global(.intel-card--highlight) {
			grid-column: span 6;
		}
	}

	@media (max-width: 768px) {
		:global(.intel-card--bio),
		:global(.intel-card--education),
		:global(.intel-card--highlight) {
			grid-column: span 12;
		}
	}
</style>
