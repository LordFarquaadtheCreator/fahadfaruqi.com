<script lang="ts">
	let {
		name,
		activeSection = 'ops'
	} = $props<{
		name: string;
		activeSection?: string;
	}>();

	const operatorId = $derived('OPR_' + name.slice(0, 4).toUpperCase().replace(/[^A-Z]/g, '') + '01');
</script>

<!-- Side Navigation Bar -->
<aside class="side-nav" aria-label="Primary">
	<div class="side-nav__top">
		<!-- Operator Profile -->
		<div class="operator-badge">
			<div class="operator-avatar">
				<span class="material-symbols-outlined">person</span>
			</div>
			<span class="operator-id tech-label">{operatorId}</span>
		</div>

		<!-- Navigation Icons -->
		<nav class="side-nav__icons">
			<a href="#top" class="nav-icon" class:active={activeSection === 'ops'} aria-label="Operations">
				<span class="material-symbols-outlined">radar</span>
				<span class="nav-icon__label tech-label">OPS</span>
			</a>
			<a href="#about" class="nav-icon" class:active={activeSection === 'intel'} aria-label="Intel">
				<span class="material-symbols-outlined">dataset</span>
				<span class="nav-icon__label tech-label">INTEL</span>
			</a>
			<a href="#resume" class="nav-icon" class:active={activeSection === 'archive'} aria-label="Archive">
				<span class="material-symbols-outlined">folder_open</span>
				<span class="nav-icon__label tech-label">ARCHIVE</span>
			</a>
			<a href="#contact" class="nav-icon" class:active={activeSection === 'uplink'} aria-label="Comm Link">
				<span class="material-symbols-outlined">hub</span>
				<span class="nav-icon__label tech-label">UPLINK</span>
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

<style>
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

	.operator-avatar .material-symbols-outlined {
		font-size: 24px;
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

	.nav-icon .material-symbols-outlined {
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

	/* Hide on mobile */
	@media (max-width: 768px) {
		.side-nav {
			display: none;
		}
	}

	/* Material Symbols Base */
	:global(.material-symbols-outlined) {
		font-variation-settings: 'FILL' 0, 'wght' 400, 'GRAD' 0, 'opsz' 24;
		font-family: 'Material Symbols Outlined';
	}
</style>
