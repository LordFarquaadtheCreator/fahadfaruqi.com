<script lang="ts">
	import Icon from './Icon.svelte';
	import faviconImg from '$lib/assets/favicon.svg';

	let {
		name,
		activeSection = 'ops'
	} = $props<{
		name: string;
		activeSection?: string;
	}>();

	const operatorId = $derived('PRO_' + name.slice(0, 4).toUpperCase().replace(/[^A-Z]/g, '') + '01');
	const sysId = $derived('0x' + Math.floor(Math.random() * 65535).toString(16).toUpperCase().padStart(4, '0'));
</script>

<!-- Mobile Top Header -->
<header class="mobile-header">
	<div class="mobile-header__left">
		<img src={faviconImg} alt="Logo" class="mobile-header__icon" />
		<h1 class="mobile-header__title">TERM_V.4.02</h1>
	</div>
	<div class="mobile-header__right">
		<span class="mobile-header__sysid">SYS_ID: {sysId}</span>
		<Icon name="signal_cellular_4_bar" class="mobile-header__signal" />
	</div>
</header>

<!-- Side Navigation Bar (Desktop) -->
<aside class="side-nav" aria-label="Primary">
	<div class="side-nav__top">
		<!-- Operator Profile -->
		<div class="operator-badge">
			<div class="operator-avatar">
				<img src={faviconImg} alt="Logo" class="operator-avatar__icon" />
			</div>
			<span class="operator-id tech-label">{operatorId}</span>
		</div>

		<!-- Navigation Icons -->
		<nav class="side-nav__icons">
			<a href="#top" class="nav-icon" class:active={activeSection === 'home'} aria-label="Home">
				<Icon name="radar" class="nav-icon__glyph" />
				<span class="nav-icon__label tech-label">HOME</span>
			</a>
			<a href="#about" class="nav-icon" class:active={activeSection === 'about'} aria-label="About">
				<Icon name="dataset" class="nav-icon__glyph" />
				<span class="nav-icon__label tech-label">ABOUT</span>
			</a>
			<a href="#resume" class="nav-icon" class:active={activeSection === 'resume'} aria-label="Resume">
				<Icon name="folder_open" class="nav-icon__glyph" />
				<span class="nav-icon__label tech-label">RESUME</span>
			</a>
			<a href="#contact" class="nav-icon" class:active={activeSection === 'contact'} aria-label="Contact">
				<Icon name="hub" class="nav-icon__glyph" />
				<span class="nav-icon__label tech-label">CONTACT</span>
			</a>
		</nav>
	</div>

	<div class="side-nav__bottom">
		<div class="status-indicator">
			<span class="status-pip status-active"></span>
			<span class="tech-label" style="color: var(--tertiary);">LIVE</span>
		</div>
	</div>
</aside>

<!-- Mobile Bottom Navigation -->
<nav class="mobile-nav">
	<a href="#top" class="mobile-nav__item" class:active={activeSection === 'home'}>
		<Icon name="grid_view" class="mobile-nav__icon" />
		<span class="mobile-nav__label">HOME</span>
	</a>
	<a href="#about" class="mobile-nav__item" class:active={activeSection === 'about'}>
		<Icon name="code" class="mobile-nav__icon" />
		<span class="mobile-nav__label">ABOUT</span>
	</a>
	<a href="#resume" class="mobile-nav__item" class:active={activeSection === 'resume'}>
		<Icon name="memory" class="mobile-nav__icon" />
		<span class="mobile-nav__label">RESUME</span>
	</a>
	<a href="#contact" class="mobile-nav__item" class:active={activeSection === 'contact'}>
		<Icon name="settings_input_component" class="mobile-nav__icon" />
		<span class="mobile-nav__label">CONTACT</span>
	</a>
</nav>

<style>
	/* Mobile Top Header */
	.mobile-header {
		display: none;
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		z-index: 50;
		justify-content: space-between;
		align-items: center;
		padding: 0.5rem 1rem;
		background: rgba(18, 19, 20, 0.9);
		backdrop-filter: blur(12px);
	}

	.mobile-header__left {
		display: flex;
		align-items: center;
		gap: 0.5rem;
	}

	.mobile-header__icon {
		width: 1.25rem;
		height: 1.25rem;
		object-fit: contain;
	}

	.mobile-header__title {
		font-family: 'Space Grotesk', sans-serif;
		font-size: 0.75rem;
		letter-spacing: 0.1em;
		text-transform: uppercase;
		color: var(--primary);
		font-weight: 700;
		margin: 0;
	}

	.mobile-header__right {
		display: flex;
		align-items: center;
		gap: 1rem;
	}

	.mobile-header__sysid {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.6rem;
		color: var(--on-surface-variant);
		letter-spacing: 0.1em;
		animation: jitter 3s ease-in-out infinite;
	}

	@keyframes jitter {
		0%, 100% { transform: translate(0, 0); }
		25% { transform: translate(0.5px, -0.5px); }
		50% { transform: translate(-0.5px, 0.5px); }
		75% { transform: translate(0.5px, 0.5px); }
	}

	:global(.mobile-header__signal) {
		font-size: 1rem;
		color: var(--primary);
	}

	/* Side Navigation */
	.side-nav {
		position: fixed;
		left: 0;
		top: 0;
		bottom: 0;
		width: 80px;
		z-index: 40;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		padding: 1.5rem 0;
		background: var(--surface-container-lowest);
		border-right: 1px solid rgba(255, 0, 60, 0.1);
	}

	.side-nav__top {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1.5rem;
	}

	.operator-badge {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 0.5rem;
	}

	.operator-avatar {
		width: 40px;
		height: 40px;
		display: flex;
		align-items: center;
		justify-content: center;
		background: var(--surface-container-high);
		border: 1px solid rgba(255, 0, 60, 0.2);
		color: var(--on-surface-variant);
	}

	.operator-avatar__icon {
		width: 24px;
		height: 24px;
		object-fit: contain;
	}

	.operator-id {
		color: var(--inverse-primary);
		font-weight: 700;
	}

	.side-nav__icons {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
		width: 100%;
		padding: 0 0.75rem;
	}

	.nav-icon {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 0.25rem;
		padding: 0.75rem 0.25rem;
		color: var(--on-surface-variant);
		transition: all 150ms ease;
		position: relative;
		cursor: pointer;
	}

	.nav-icon:hover {
		background: var(--surface-container-high);
		color: var(--primary);
	}

	.nav-icon.active {
		background: var(--inverse-primary);
		color: white;
	}

	.nav-icon.active::before {
		content: '';
		position: absolute;
		left: -0.75rem;
		top: 50%;
		transform: translateY(-50%);
		width: 3px;
		height: 60%;
		background: var(--primary);
	}

	:global(.nav-icon__glyph) {
		font-size: 20px;
	}

	.nav-icon__label {
		font-size: 0.5rem;
		letter-spacing: 0.08em;
	}

	.side-nav__bottom {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 0.5rem;
	}

	.status-indicator {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 0.25rem;
	}

	/* Mobile Bottom Navigation */
	.mobile-nav {
		display: none;
		position: fixed;
		bottom: 0;
		left: 0;
		right: 0;
		z-index: 50;
		height: 64px;
		background: var(--background);
		box-shadow: 0 -4px 20px rgba(255, 179, 178, 0.05);
		justify-content: space-around;
		align-items: stretch;
	}

	.mobile-nav__item {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 0.25rem;
		flex: 1;
		padding: 0.75rem;
		color: var(--on-surface-variant);
		opacity: 0.5;
		transition: all 150ms ease;
		text-decoration: none;
	}

	.mobile-nav__item:hover {
		color: var(--primary);
		opacity: 0.8;
		background: rgba(255, 179, 178, 0.05);
	}

	.mobile-nav__item.active {
		background: var(--primary);
		color: var(--on-primary);
		opacity: 1;
	}

	:global(.mobile-nav__icon) {
		font-size: 1.25rem;
	}

	.mobile-nav__label {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.6rem;
		letter-spacing: 0.1em;
	}

	/* Responsive - Mobile */
	@media (max-width: 1024px) {
		.mobile-header {
			display: flex;
		}

		.side-nav {
			display: none;
		}

		.mobile-nav {
			display: flex;
		}
	}

</style>
