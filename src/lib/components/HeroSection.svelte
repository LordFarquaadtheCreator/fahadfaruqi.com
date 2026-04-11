<script lang="ts">
	import { reveal } from '$lib/actions/reveal';
	import type { ProfileData } from '$lib/data/portfolio';
	import Icon from './Icon.svelte';

	let {
		profile,
		resumeDownloadUrl
	} = $props<{
		profile: ProfileData;
		resumeDownloadUrl?: string;
	}>();

	let uptimeProgress = $state(33);

	// Uptime animation
	$effect(() => {
		const interval = setInterval(() => {
			uptimeProgress = 30 + Math.random() * 10;
		}, 2000 + Math.random() * 2000);
		return () => clearInterval(interval);
	});
</script>

<section class="hero-hub" id="top">
	<!-- Left Section: Hub Navigation Cards (Asymmetric Bento) -->
	<div class="hub-grid">
		<!-- ABOUT ME: Large Hero Card -->
		<a href="#about" class="hub-card hub-card--projects">
			<div class="ref-label">REF_NO: BIO_882</div>
			<div class="halftone-overlay"></div>
			<div class="card-content">
				<div class="card-header-content">
					<span class="status-indicator-pip"></span>
					<h2 class="card-title chromatic">About Me</h2>
				</div>
				<p class="card-desc">Technical specifications and operational history. Decrypted for authorized review.</p>
				<div class="card-footer">
					<span class="enter-link">View_Dossier</span>
				</div>
			</div>
			<!-- HUD Corner Brackets -->
			<div class="hud-corner hud-corner-tl"></div>
			<div class="hud-corner hud-corner-br"></div>
		</a>

		<!-- RESUME: Medium Card -->
		<a href="#resume" class="hub-card hub-card--resume">
			<div class="ref-label">TYPE: DOC_X</div>
			<div class="card-content-vertical">
				<Icon name="description" class="resume-icon" />
				<h3 class="card-title-sm chromatic">Resume</h3>
			</div>
			<div class="card-meta">LOGGED_EXP: 09_YEARS</div>
		</a>

		<!-- CONTACT: Small Square -->
		<a href="#contact" class="hub-card hub-card--small">
			<div class="card-content-center">
				<Icon name="alternate_email" class="contact-icon" />
				<span class="card-label chromatic">Contact</span>
			</div>
		</a>

		<!-- STATUS/DECORATIVE: Wide Minimalist -->
		<div class="hub-card hub-card--status">
			<div class="status-info">
				<span class="status-label">Uptime_Sequence</span>
				<span class="status-value">99.9%_STABLE</span>
			</div>
			<div class="progress-bar">
				<div class="progress-fill" style="width: {uptimeProgress}%"></div>
			</div>
		</div>
	</div>

	<!-- Right Section: Immersive Graphic Area -->
	<div class="graphic-area">
		<div class="graphic-container">
			<!-- Abstract Graphic Background -->
			<div class="graphic-bg">
				<div class="gradient-overlay"></div>
			</div>
			
			<!-- HUD Overlays -->
			<div class="hud-overlay">
				<div class="hud-top">
					<div class="targeting-box">
						<div class="target-label">Target_Selected</div>
						<div class="target-value chromatic">FAHAD FARUQI // DEFAULT_PHOTO</div>
					</div>
					<div class="coordinates">COORDINATES: 34.0522° N, 118.2437° W</div>
				</div>
				<p>my image image placeholder</p>
				<div class="hud-bottom">
					<div class="manifesto-box">
						<div class="manifesto-label">Core_Manifesto</div>
						<p class="manifesto-text">In conjunction with generative tools, anything is possible as long as the human is centered in development and mind.</p>
					</div>
					<div class="ghost-text chromatic">SIGNAL</div>
				</div>
			</div>
			
			<!-- Targeting Reticle (HUD-Gaze) -->
			<div class="targeting-reticle">
				<div class="reticle-inner">
					<div class="corner corner-tl"></div>
					<div class="corner corner-tr"></div>
					<div class="corner corner-bl"></div>
					<div class="corner corner-br"></div>
				</div>
				<div class="crosshair-h"></div>
				<div class="crosshair-v"></div>
			</div>
		</div>
	</div>
</section>

<!-- Mobile Hero Layout -->
<section class="hero-mobile" id="top-mobile">
	<!-- Ghost Background Text -->
	<div class="ghost-bg">
		<span class="ghost-text-mobile">OPERATOR_HUB</span>
	</div>

	<div class="mobile-content">
		<!-- Hero Data Module: ABOUT ME -->
		<a href="#about" class="mobile-card mobile-card--hero">
			<div class="hud-bracket-mobile hud-tl"></div>
			<div class="hud-bracket-mobile hud-tr"></div>
			<div class="hud-bracket-mobile hud-bl"></div>
			<div class="hud-bracket-mobile hud-br"></div>
			<div class="ref-label-mobile">REF_NO: BIO_0029-A</div>
			
			<div class="mobile-card-content">
				<span class="active-badge">ACTIVE_NODE</span>
				<h2 class="mobile-title chromatic">ABOUT_ME</h2>
				<p class="mobile-desc">Accessing operator dossier: technical specifications and deployment history.</p>
			</div>
			
			<div class="mobile-card-footer">
				<div class="network-status">
					<div class="status-row">
						<span class="status-dot network-dot"></span>
						<span class="status-text">Network: Linked</span>
					</div>
					<div class="status-row">
						<span class="status-dot status-dot-blue"></span>
						<span class="status-text">Status: Optimal</span>
					</div>
				</div>
				<span class="init-button">
					INITIALIZE <Icon name="arrow_forward_ios" class="init-button__icon" />
				</span>
			</div>
			
			<div class="card-glow"></div>
		</a>

		<!-- Resume Card -->
		<a href="#resume" class="mobile-card mobile-card--resume">
			<div class="ref-label-mobile">REF_ID: HX-04</div>
			<div class="mobile-card-header">
				<h3 class="mobile-title-sm">RESUME</h3>
				<div class="title-underline"></div>
			</div>
			<div class="skill-bar">
				<div class="skill-fill"></div>
			</div>
			<p class="mobile-card-desc">Operator credentials and deployment history overview.</p>
			<div class="halftone-corner"></div>
		</a>

		<!-- Wide Contact Section -->
		<a href="#contact" class="mobile-card mobile-card--contact">
			<div class="ref-label-mobile">COMMS_LINK_ENCRYPTED</div>
			<div class="contact-content">
				<div class="satellite-icon">
					<Icon name="satellite_alt" class="satellite-icon__glyph" />
				</div>
				<div class="contact-info">
					<h3 class="contact-title">CONTACT_OPERATOR</h3>
					<p class="secure-channel">SECURE_CHANNEL: 256-BIT_AES</p>
					<div class="signal-bar">
						<span class="signal-label">Incoming_Signal</span>
						<div class="signal-line-bg"></div>
						<div class="signal-line"></div>
						<div class="signal-waiting">Waiting for transmission...</div>
					</div>
				</div>
			</div>
		</a>

		<!-- Bottom Metadata Stats -->
		<div class="mobile-stats">
			<div class="stat-box">
				<div class="stat-label">Lat_Coord</div>
				<div class="stat-value">52.3676° N</div>
			</div>
			<div class="stat-box">
				<div class="stat-label">Long_Coord</div>
				<div class="stat-value">4.9041° E</div>
			</div>
			<div class="stat-box">
				<div class="stat-label">Uptime</div>
				<div class="stat-value status-ok">99.98%</div>
			</div>
		</div>
	</div>
</section>

<style>
	/* Desktop Hero Hub Layout */
	.hero-hub {
		display: flex;
		min-height: calc(100vh - 80px - 32px);
		padding: 1.5rem;
		gap: 1.5rem;
	}

	.hub-grid {
		width: 45%;
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		grid-template-rows: repeat(6, 1fr);
		gap: 1rem;
	}

	.graphic-area {
		flex: 1;
		display: flex;
	}

	/* Hub Cards */
	.hub-card {
		background: var(--surface-container-low);
		position: relative;
		overflow: hidden;
		cursor: pointer;
		transition: all 200ms ease;
		text-decoration: none;
		color: inherit;
	}

	.hub-card:hover {
		background: var(--surface-container-high);
	}

	.hub-card--projects {
		grid-column: span 2;
		grid-row: span 3;
		display: flex;
		flex-direction: column;
		padding: 2rem;
	}

	.hub-card--projects:hover .halftone-overlay {
		opacity: 0.1;
	}

	.hub-card--resume {
		grid-column: span 1;
		grid-row: span 2;
		padding: 1.5rem;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
	}

	.hub-card--small {
		grid-column: span 1;
		grid-row: span 2;
		padding: 1rem;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}

	.hub-card--status {
		grid-column: span 2;
		grid-row: span 1;
		border: 1px solid rgba(175, 135, 134, 0.15);
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 0 1.5rem;
		background: transparent;
	}

	/* Card Content */
	.ref-label {
		position: absolute;
		top: 0.5rem;
		right: 0.5rem;
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.6rem;
		color: var(--on-surface-variant);
		opacity: 0.4;
		animation: jitter 2.5s ease-in-out infinite;
		animation-delay: calc(var(--jitter-delay, 0) * 1s);
	}

	@keyframes jitter {
		0%, 100% { transform: translate(0, 0); }
		25% { transform: translate(0.5px, -0.5px); }
		50% { transform: translate(-0.5px, 0.5px); }
		75% { transform: translate(0.5px, 0.5px); }
	}

	.halftone-overlay {
		position: absolute;
		inset: 0;
		background: radial-gradient(circle, var(--primary) 1px, transparent 1px);
		background-size: 4px 4px;
		opacity: 0;
		transition: opacity 200ms ease;
	}

	.card-content {
		position: relative;
		z-index: 1;
		height: 100%;
		display: flex;
		flex-direction: column;
		justify-content: flex-end;
	}

	.card-header-content {
		margin-bottom: 0.5rem;
	}

	.status-indicator-pip {
		display: inline-block;
		width: 8px;
		height: 8px;
		background: var(--tertiary);
		box-shadow: 0 0 10px rgba(42, 229, 0, 0.5);
		margin-bottom: 0.5rem;
	}

	.card-title {
		font-size: 2.5rem;
		font-weight: 700;
		letter-spacing: -0.02em;
		text-transform: uppercase;
		margin: 0;
		color: var(--on-surface);
	}

	.card-title-sm {
		font-size: 1.25rem;
		font-weight: 700;
		letter-spacing: -0.01em;
		text-transform: uppercase;
		margin: 0.5rem 0 0;
		color: var(--on-surface);
	}

	.card-label {
		font-size: 0.65rem;
		font-weight: 700;
		text-transform: uppercase;
		letter-spacing: 0.15em;
	}

	.card-desc {
		color: var(--on-surface-variant);
		font-size: 0.875rem;
		max-width: 280px;
		line-height: 1.6;
		margin: 1rem 0;
	}

	.enter-link {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.65rem;
		text-transform: uppercase;
		letter-spacing: 0.15em;
		color: var(--primary);
		border-bottom: 1px solid rgba(255, 179, 178, 0.3);
		padding-bottom: 0.25rem;
	}

	.card-footer {
		margin-top: 1.5rem;
	}

	.card-content-vertical {
		position: relative;
		z-index: 1;
	}

	.card-content-center {
		text-align: center;
	}

	.card-meta {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.6rem;
		color: var(--on-surface-variant);
		opacity: 0.6;
		position: relative;
		z-index: 1;
	}

	:global(.resume-icon) {
		font-size: 1.5rem;
		color: var(--secondary);
		margin-bottom: 0.5rem;
	}

	:global(.contact-icon) {
		font-size: 1.5rem;
		color: var(--primary);
		display: block;
		margin-bottom: 0.25rem;
	}

	/* Status Bar */
	.status-info {
		display: flex;
		flex-direction: column;
		gap: 0.25rem;
	}

	.status-label {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.55rem;
		color: var(--on-surface-variant);
		text-transform: uppercase;
	}

	.status-value {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.75rem;
		color: var(--tertiary);
		animation: jitter 4s ease-in-out infinite;
	}

	.progress-bar {
		width: 8rem;
		height: 1px;
		background: rgba(175, 135, 134, 0.3);
		position: relative;
	}

	.progress-fill {
		position: absolute;
		top: 0;
		left: 0;
		height: 100%;
		background: var(--primary);
		box-shadow: 0 0 8px var(--primary);
		transition: width 0.5s ease;
	}

	/* HUD Corners */
	.hud-corner {
		position: absolute;
		width: 16px;
		height: 16px;
		border-color: rgba(175, 135, 134, 0.3);
	}

	.hud-corner-tl {
		top: 1rem;
		left: 1rem;
		border-top: 1px solid;
		border-left: 1px solid;
	}

	.hud-corner-br {
		bottom: 1rem;
		right: 1rem;
		border-bottom: 1px solid;
		border-right: 1px solid;
	}

	/* Graphic Area */
	.graphic-container {
		position: relative;
		width: 100%;
		height: 100%;
		background: var(--surface-container-lowest);
		overflow: hidden;
	}

	.graphic-bg {
		position: absolute;
		inset: 0;
		background: 
			linear-gradient(135deg, rgba(255, 179, 178, 0.05) 0%, transparent 50%),
			linear-gradient(225deg, rgba(211, 251, 255, 0.03) 0%, transparent 50%),
			var(--surface-container-lowest);
	}

	.gradient-overlay {
		position: absolute;
		inset: 0;
		background: linear-gradient(to top right, var(--background) 0%, transparent 50%, rgba(255, 179, 178, 0.1) 100%);
	}

	.hud-overlay {
		position: absolute;
		inset: 0;
		padding: 2rem;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		pointer-events: none;
	}

	.hud-top {
		display: flex;
		justify-content: space-between;
		align-items: flex-start;
	}

	.targeting-box {
		padding: 1rem;
		border-left: 1px solid rgba(255, 179, 178, 0.4);
		border-top: 1px solid rgba(255, 179, 178, 0.4);
	}

	.target-label {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.65rem;
		color: var(--primary);
		text-transform: uppercase;
		margin-bottom: 0.25rem;
	}

	.target-value {
		font-size: 1.5rem;
		font-weight: 700;
		color: var(--on-surface);
	}

	.coordinates {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.55rem;
		color: var(--on-surface-variant);
		text-align: right;
	}

	.hud-bottom {
		display: flex;
		justify-content: space-between;
		align-items: flex-end;
	}

	.manifesto-box {
		max-width: 280px;
		padding: 1rem;
		background: rgba(18, 19, 20, 0.8);
		backdrop-filter: blur(12px);
	}

	.manifesto-label {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.65rem;
		color: var(--secondary);
		text-transform: uppercase;
		margin-bottom: 0.5rem;
	}

	.manifesto-text {
		font-size: 0.75rem;
		color: var(--on-surface);
		opacity: 0.8;
		line-height: 1.6;
		margin: 0;
	}

	.ghost-text {
		font-family: 'JetBrains Mono', monospace;
		font-size: 3.5rem;
		color: var(--on-surface);
		opacity: 0.05;
		font-weight: 700;
		line-height: 1;
		user-select: none;
	}

	/* Targeting Reticle */
	.targeting-reticle {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		width: 12rem;
		height: 12rem;
		border: 1px solid rgba(175, 135, 134, 0.1);
		display: flex;
		align-items: center;
		justify-content: center;
		pointer-events: none;
	}

	.reticle-inner {
		position: relative;
		width: 4rem;
		height: 4rem;
	}

	.corner {
		position: absolute;
		width: 1rem;
		height: 1rem;
		border-color: var(--primary);
	}

	.corner-tl { top: 0; left: 0; border-top: 2px solid; border-left: 2px solid; }
	.corner-tr { top: 0; right: 0; border-top: 2px solid; border-right: 2px solid; }
	.corner-bl { bottom: 0; left: 0; border-bottom: 2px solid; border-left: 2px solid; }
	.corner-br { bottom: 0; right: 0; border-bottom: 2px solid; border-right: 2px solid; }

	.crosshair-h {
		position: absolute;
		top: 50%;
		left: 0;
		width: 100%;
		height: 1px;
		background: rgba(255, 179, 178, 0.1);
		transform: translateY(-50%);
	}

	.crosshair-v {
		position: absolute;
		top: 0;
		left: 50%;
		width: 1px;
		height: 100%;
		background: rgba(255, 179, 178, 0.1);
		transform: translateX(-50%);
	}

	/* Chromatic Effect */
	.chromatic {
		text-shadow: 2px 0 0 rgba(255, 0, 0, 0.5), -2px 0 0 rgba(0, 255, 255, 0.5);
	}

	/* Mobile Hero Layout */
	.hero-mobile {
		display: none;
		padding: 1rem;
		padding-bottom: 6rem;
		position: relative;
		overflow: hidden;
	}

	.ghost-bg {
		position: fixed;
		inset: 0;
		display: flex;
		align-items: center;
		justify-content: center;
		pointer-events: none;
		opacity: 0.03;
		z-index: 0;
	}

	.ghost-text-mobile {
		font-size: 8rem;
		font-weight: 900;
		transform: rotate(90deg);
		white-space: nowrap;
		letter-spacing: -0.05em;
		color: var(--on-surface);
	}

	.mobile-content {
		position: relative;
		z-index: 1;
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.mobile-card {
		background: var(--surface-container-low);
		padding: 1.5rem;
		position: relative;
		overflow: hidden;
		text-decoration: none;
		color: inherit;
	}

	.mobile-card--hero {
		min-height: 16rem;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
	}

	.hud-bracket-mobile {
		position: absolute;
		width: 12px;
		height: 12px;
		border-color: rgba(175, 135, 134, 0.5);
	}

	.hud-tl { top: 0; left: 0; border-top: 2px solid; border-left: 2px solid; }
	.hud-tr { top: 0; right: 0; border-top: 2px solid; border-right: 2px solid; }
	.hud-bl { bottom: 0; left: 0; border-bottom: 2px solid; border-left: 2px solid; }
	.hud-br { bottom: 0; right: 0; border-bottom: 2px solid; border-right: 2px solid; }

	.ref-label-mobile {
		position: absolute;
		top: 0.5rem;
		right: 0.5rem;
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.55rem;
		color: var(--on-surface-variant);
		opacity: 0.3;
		animation: jitter 3s ease-in-out infinite;
	}

	.active-badge {
		display: inline-block;
		padding: 0.25rem 0.5rem;
		background: var(--primary);
		color: var(--on-primary);
		font-size: 0.6rem;
		font-weight: 700;
		margin-bottom: 0.75rem;
	}

	.mobile-title {
		font-size: 2.5rem;
		font-weight: 900;
		letter-spacing: -0.03em;
		color: var(--primary);
		margin: 0;
	}

	.mobile-desc {
		font-size: 0.875rem;
		color: var(--on-surface-variant);
		max-width: 200px;
		margin-top: 0.5rem;
		line-height: 1.5;
	}

	.mobile-card-footer {
		display: flex;
		justify-content: space-between;
		align-items: flex-end;
		position: relative;
		z-index: 2;
	}

	.network-status {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}

	.status-row {
		display: flex;
		align-items: center;
		gap: 0.5rem;
	}

	.status-dot {
		width: 6px;
		height: 6px;
		background: var(--tertiary);
		box-shadow: 0 0 8px rgba(42, 229, 0, 0.5);
	}

	.status-dot-blue {
		background: var(--secondary);
		box-shadow: 0 0 8px rgba(211, 251, 255, 0.5);
	}

	.status-text {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.6rem;
		color: var(--on-surface-variant);
		text-transform: uppercase;
	}

	.init-button {
		display: inline-flex;
		align-items: center;
		gap: 0.5rem;
		padding: 0.75rem 1rem;
		background: var(--primary);
		color: var(--on-primary);
		font-weight: 700;
		font-size: 0.75rem;
	}

	:global(.init-button__icon) {
		font-size: 0.875rem;
	}

	.card-glow {
		position: absolute;
		right: -10%;
		top: 0;
		width: 50%;
		height: 100%;
		background: radial-gradient(circle, rgba(255, 179, 178, 0.1) 0%, transparent 70%);
		filter: blur(40px);
		pointer-events: none;
	}

	.mobile-card--resume {
		min-height: 12rem;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
	}

	.mobile-card-header {
		margin-bottom: 1rem;
	}

	.mobile-title-sm {
		font-size: 1.25rem;
		font-weight: 700;
		color: var(--secondary);
		margin: 0;
	}

	.title-underline {
		width: 2rem;
		height: 4px;
		background: var(--secondary);
		margin-top: 0.5rem;
	}

	.skill-bar {
		width: 100%;
		height: 4px;
		background: var(--surface-container-lowest);
		margin: 1rem 0;
	}

	.skill-fill {
		width: 75%;
		height: 100%;
		background: var(--secondary-container);
	}

	.mobile-card-desc {
		font-size: 0.7rem;
		color: var(--on-surface-variant);
		line-height: 1.5;
	}

	.halftone-corner {
		position: absolute;
		bottom: -1rem;
		right: -1rem;
		width: 6rem;
		height: 6rem;
		background: radial-gradient(circle, var(--secondary) 1px, transparent 1px);
		background-size: 4px 4px;
		opacity: 0.15;
		transform: rotate(12deg);
	}

	.mobile-card--contact {
		padding: 1.5rem;
	}

	.contact-content {
		display: flex;
		gap: 1rem;
	}

	.satellite-icon {
		width: 3rem;
		height: 3rem;
		background: rgba(228, 226, 227, 0.1);
		display: flex;
		align-items: center;
		justify-content: center;
		flex-shrink: 0;
	}

	:global(.satellite-icon__glyph) {
		font-size: 1.5rem;
		color: var(--primary);
	}

	.contact-info {
		flex: 1;
	}

	.contact-title {
		font-size: 1.125rem;
		font-weight: 700;
		color: var(--primary);
		margin: 0 0 0.5rem;
	}

	.secure-channel {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.65rem;
		color: var(--on-surface-variant);
		margin-bottom: 1rem;
	}

	.signal-bar {
		position: relative;
		padding-bottom: 0.5rem;
	}

	.signal-label {
		font-size: 0.55rem;
		text-transform: uppercase;
		color: var(--secondary);
		display: block;
		margin-bottom: 0.25rem;
	}

	.signal-line-bg {
		position: absolute;
		bottom: 0;
		left: 0;
		width: 100%;
		height: 1px;
		background: rgba(175, 135, 134, 0.3);
	}

	.signal-line {
		position: absolute;
		bottom: 0;
		left: 0;
		width: 25%;
		height: 1px;
		background: var(--secondary);
	}

	.signal-waiting {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.75rem;
		color: var(--on-surface);
		opacity: 0.6;
		padding: 0.5rem 0;
	}

	.mobile-stats {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		gap: 0.75rem;
		padding-top: 1rem;
	}

	.stat-box {
		background: var(--surface-container-low);
		padding: 0.75rem;
	}

	.stat-label {
		font-size: 0.55rem;
		text-transform: uppercase;
		color: var(--on-surface-variant);
		opacity: 0.5;
		margin-bottom: 0.25rem;
	}

	.stat-value {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.7rem;
		color: var(--on-surface);
		letter-spacing: -0.02em;
		animation: jitter 3.5s ease-in-out infinite;
	}

	.status-ok {
		color: var(--tertiary);
	}

	/* Responsive - Hide desktop on mobile, show mobile */
	@media (max-width: 1024px) {
		.hero-hub {
			display: none;
		}

		.hero-mobile {
			display: block;
		}
	}

	@media (max-width: 768px) {
		.hero-mobile {
			padding-top: 4rem;
		}

		.ghost-text-mobile {
			font-size: 5rem;
		}
	}
</style>
