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

	let sysStatus = $state(72 + Math.random() * 31.21);

	// Stream data jitter state
	let bitrate = $state(0.5 + Math.random() * 1.567);
	let packets = $state(98 + Math.random() * 1.9873);
	let latency = $state(0.01 + Math.random() * 0.2456);
	let vizBars = $state([20, 60, 80, 45, 90, 70].map(h => h + Math.random() * 20 - 10));

	// Sys Status jitter
	$effect(() => {
		const interval = setInterval(() => {
			sysStatus = 72 + Math.random() * 31.21;
		}, 1000 + Math.random() * 3000);
		return () => clearInterval(interval);
	});

	// Stream data jitter
	$effect(() => {
		const interval = setInterval(() => {
			bitrate = 0.8 + Math.random() * 0.8;
			packets = 98 + Math.random() * 2;
			latency = 0.01 + Math.random() * 0.03;
			vizBars = vizBars.map(() => 30 + Math.random() * 70);
		}, 300 + Math.random() * 3000);
		return () => clearInterval(interval);
	});
</script>

<section class="hero" id="top">
	<!-- System Metadata -->
	<div class="sys-meta">
		<div class="sys-meta__item">
			<span class="tech-label">SYS_STATUS</span>
			<span class="sys-meta__value" style="color: var(--tertiary);">{sysStatus.toFixed(2)}%</span>
		</div>
		<div class="sys-meta__item">
			<span class="tech-label">ARCHIVE_VOL</span>
			<span class="sys-meta__value">6.7TB</span>
		</div>
		<div class="sys-meta__item">
			<span class="tech-label">LOC</span>
			<span class="sys-meta__value">P292+W9 New York</span>
		</div>
	</div>

	<!-- Main Hero Content -->
	<div class="hero__main reveal" use:reveal>
		<p class="eyebrow">{profile.location} // HOME_TOWN</p>
		<h1 class="hero__title chromatic">
			{profile.name.toUpperCase().replace(/ /g, '_')}
			<span class="hero__title-suffix">.MD</span>
		</h1>
		<p class="hero__role">{profile.title}</p>
		<p class="hero__intro">{profile.heroIntro}</p>
	</div>

	<!-- Bento Grid Dashboard -->
	<div class="hero__dashboard">
		<!-- Primary Signal Card -->
		<div class="dashboard-card dashboard-card--primary reveal" use:reveal>
			<div class="hud-bracket-tl"></div>
			<div class="hud-bracket-tr"></div>
			<div class="hud-bracket-bl"></div>
			<div class="hud-bracket-br"></div>

			<div class="card-header">
				<span class="tech-label" style="color: var(--primary);">REF_NO: 0882_X</span>
				<div class="status-badge">
					<span class="status-pip status-active"></span>
					<span class="tech-label" style="color: var(--tertiary);">DECRYPTED</span>
				</div>
			</div>

			<div class="card-body">
				<h2 class="card-title">NEURAL_INTERFACE_DUMP.LOG</h2>
				<p class="card-desc">
					High-density data extraction from core-processor. Contains volatile behavioral maps
					and interface protocols.
				</p>
			</div>

			<div class="card-metrics">
				{#each profile.metrics as metric, i}
					<div class="metric-item">
						<span class="metric-value">{metric.value}</span>
						<span class="metric-label">{metric.label}</span>
						<span class="metric-detail">{metric.detail}</span>
					</div>
				{/each}
			</div>

			<div class="card-actions">
				<a href="#resume" class="button">ACCESS_NODE</a>
				{#if resumeDownloadUrl}
					<a href={resumeDownloadUrl} class="button button--ghost" target="_blank" rel="noreferrer">
						DOWNLOAD_DAT
					</a>
				{/if}
			</div>
		</div>

		<!-- Data Stream Panel -->
		<div class="dashboard-card dashboard-card--stream reveal" use:reveal>
			<div class="card-header">
				<span class="tech-label" style="color: var(--secondary);">STREAM_ID_44</span>
				<span class="material-symbols-outlined" style="color: var(--secondary);">sensors</span>
			</div>

			<div class="stream-data">
				<div class="stream-row">
					<span class="tech-label">BITRATE</span>
					<span style="color: var(--on-surface);">{bitrate.toFixed(1)} GB/S</span>
				</div>
				<div class="stream-row">
					<span class="tech-label">PACKETS</span>
					<span style="color: var(--on-surface);">{packets.toFixed(1)}%</span>
				</div>
				<div class="stream-row">
					<span class="tech-label">LATENCY</span>
					<span style="color: var(--on-surface);">{latency.toFixed(2)}ms</span>
				</div>
			</div>

			<div class="stream-viz">
				<div class="viz-bar" style="height: {vizBars[0]}%;"></div>
				<div class="viz-bar" style="height: {vizBars[1]}%;"></div>
				<div class="viz-bar" style="height: {vizBars[2]}%;"></div>
				<div class="viz-bar" style="height: {vizBars[3]}%;"></div>
				<div class="viz-bar" style="height: {vizBars[4]}%;"></div>
			</div>
		</div>

		<!-- Quick Links Module -->
		<div class="dashboard-card dashboard-card--links reveal" use:reveal>
			<div class="card-header">
				<span class="tech-label">QUICK_UPLINK</span>
			</div>
			<div class="links-grid">
				{#each profile.socialLinks as link}
					<a href={link.href} class="link-item" target={link.href.startsWith('http') ? '_blank' : undefined} rel="noreferrer">
						<span class="link-label">{link.label.toUpperCase()}</span>
						<span class="material-symbols-outlined">arrow_outward</span>
					</a>
				{/each}
			</div>
		</div>
	</div>
</section>

<style>
	.hero {
		padding: 2rem 0 4rem;
		position: relative;
	}

	/* System Metadata - Top Right */
	.sys-meta {
		position: absolute;
		top: 0;
		right: 0;
		display: flex;
		flex-direction: column;
		gap: 0.25rem;
		text-align: right;
		opacity: 0.4;
	}

	.sys-meta__item {
		display: flex;
		flex-direction: column;
		align-items: flex-end;
		gap: 0.1rem;
	}

	.sys-meta__value {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.65rem;
		color: var(--on-surface);
	}

	/* Main Hero Content */
	.hero__main {
		margin-bottom: 3rem;
		position: relative;
	}

	.hero__title {
		font-size: clamp(2.5rem, 8vw, 5rem);
		font-weight: 900;
		font-style: italic;
		letter-spacing: -0.03em;
		line-height: 0.95;
		color: var(--on-surface);
		margin: 0.5rem 0 0;
	}

	.hero__title-suffix {
		color: var(--primary);
	}

	.hero__role {
		font-family: 'Space Grotesk', sans-serif;
		font-size: clamp(1rem, 2.5vw, 1.35rem);
		color: var(--primary);
		margin: 1rem 0 0;
		letter-spacing: 0.02em;
	}

	.hero__intro {
		max-width: 42rem;
		color: var(--on-surface-variant);
		margin: 1rem 0 0;
		font-size: 0.95rem;
		line-height: 1.7;
	}

	/* Dashboard Grid */
	.hero__dashboard {
		display: grid;
		grid-template-columns: repeat(12, 1fr);
		gap: 1rem;
	}

	.dashboard-card {
		background: var(--surface-container-low);
		border: 1px solid rgba(175, 135, 134, 0.15);
		padding: 1.5rem;
		position: relative;
		transition: all 200ms ease;
	}

	.dashboard-card:hover {
		border-color: rgba(255, 179, 178, 0.4);
		box-shadow: 0 0 30px rgba(255, 179, 178, 0.1);
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

	.dashboard-card--primary {
		grid-column: span 8;
	}

	.dashboard-card--stream {
		grid-column: span 4;
		background: var(--surface-container-high);
		border-left: 4px solid var(--secondary-container);
	}

	.dashboard-card--stream:hover {
		border-left-color: var(--secondary);
	}

	.dashboard-card--links {
		grid-column: span 12;
		padding: 1rem 1.5rem;
	}

	.dashboard-card--links:hover {
		animation: none;
		border-color: rgba(255, 179, 178, 0.4);
	}

	.card-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 1.5rem;
	}

	.status-badge {
		display: flex;
		align-items: center;
		gap: 0.5rem;
	}

	.card-title {
		font-size: clamp(1.25rem, 2.5vw, 1.75rem);
		font-weight: 800;
		font-style: italic;
		letter-spacing: -0.02em;
		margin-bottom: 0.75rem;
		color: var(--on-surface);
	}

	.card-desc {
		color: var(--on-surface-variant);
		font-size: 0.85rem;
		margin-bottom: 1.5rem;
		line-height: 1.6;
	}

	.card-metrics {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		gap: 1rem;
		margin-bottom: 1.5rem;
		padding: 1rem 0;
		border-top: 1px solid rgba(175, 135, 134, 0.1);
		border-bottom: 1px solid rgba(175, 135, 134, 0.1);
	}

	.metric-item {
		display: flex;
		flex-direction: column;
		gap: 0.25rem;
	}

	.metric-value {
		font-family: 'Space Grotesk', sans-serif;
		font-size: 1.75rem;
		font-weight: 700;
		color: var(--on-surface);
		line-height: 1;
	}

	.metric-label {
		font-size: 0.65rem;
		letter-spacing: 0.05em;
		text-transform: lowercase;
		color: var(--on-surface-variant);
	}

	.metric-detail {
		font-size: 0.7rem;
		color: var(--on-surface-variant);
		opacity: 0.7;
	}

	.card-actions {
		display: flex;
		gap: 0.75rem;
		flex-wrap: wrap;
	}

	/* Stream Panel */
	.stream-data {
		display: flex;
		flex-direction: column;
		gap: 0.75rem;
		margin-bottom: 1.5rem;
	}

	.stream-row {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding-bottom: 0.5rem;
		border-bottom: 1px solid var(--outline-variant);
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.7rem;
	}

	.stream-viz {
		height: 80px;
		background: var(--surface-container-lowest);
		display: flex;
		align-items: flex-end;
		justify-content: space-around;
		padding: 0.5rem;
		position: relative;
	}

	.viz-bar {
		width: 8px;
		background: var(--secondary);
		opacity: 0.6;
		transition: height 200ms ease;
	}

	/* Links Grid */
	.links-grid {
		display: flex;
		gap: 1rem;
		flex-wrap: wrap;
	}

	.link-item {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		padding: 0.5rem 1rem;
		background: var(--surface-container);
		border: 1px solid rgba(175, 135, 134, 0.1);
		transition: all 150ms ease;
	}

	.link-item:hover {
		background: var(--surface-container-high);
		border-color: var(--primary);
	}

	.link-item:hover .link-label {
		color: var(--primary);
	}

	.link-label {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.65rem;
		letter-spacing: 0.1em;
		color: var(--on-surface-variant);
		transition: color 150ms ease;
	}

	.link-item .material-symbols-outlined {
		font-size: 14px;
		color: var(--outline);
	}

	/* Responsive */
	@media (max-width: 900px) {
		.hero__dashboard {
			grid-template-columns: 1fr;
		}

		.dashboard-card--primary,
		.dashboard-card--stream,
		.dashboard-card--links {
			grid-column: span 1;
		}

		.card-metrics {
			grid-template-columns: 1fr;
		}
	}

	@media (max-width: 768px) {
		.sys-meta {
			display: none;
		}

		.hero {
			padding-top: 1rem;
		}
	}
</style>
