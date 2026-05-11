<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import SiteNav from '$lib/components/SiteNav.svelte';
	import portfolio from '$lib/data/portfolio';
	import { theme } from '$lib/theme.svelte';
	import { onMount } from 'svelte';

	let { children } = $props();
	let emailCopied = $state(false);
	let themeSliderValue = $state(1); // 0=light, 1=system, 2=dark

	function handleThemeChange(e: Event) {
		const target = e.target as HTMLInputElement;
		const value = parseInt(target.value, 10);
		themeSliderValue = value;
		if (value === 0) theme.set('light');
		else if (value === 1) theme.set('system');
		else theme.set('dark');
	}

	async function copyEmail() {
		await navigator.clipboard.writeText(portfolio.contact.email);
		emailCopied = true;
		setTimeout(() => (emailCopied = false), 1800);
	}

	onMount(() => {
		theme.init();
		themeSliderValue = theme.choice === 'light' ? 0 : theme.choice === 'system' ? 1 : 2;
	});
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<meta name="theme-color" content="#faf8ff" media="(prefers-color-scheme: light)" />
	<meta name="theme-color" content="#131313" media="(prefers-color-scheme: dark)" />
</svelte:head>

<SiteNav name={portfolio.profile.name} navigation={portfolio.navigation} />
<div class="route-shell">
	{@render children()}
</div>

<footer class="site-footer">
	<p>Made in Queens, NYC (the best borough)</p>
	<div class="footer-links">
		<a href="https://github.com/fahadfaruqi" target="_blank" rel="noopener">GitHub</a>
		<a href="https://linkedin.com/in/fahadfaruqi" target="_blank" rel="noopener">LinkedIn</a>
		<button class="email-link" onclick={copyEmail} aria-label="Copy email">
			{emailCopied ? 'Copied!' : 'Email'}
		</button>
		<div class="theme-slider-container">
			<span class="theme-label">☀</span>
			<input
				type="range"
				min="0"
				max="2"
				step="1"
				value={themeSliderValue}
				oninput={handleThemeChange}
				aria-label="Theme slider: light, system, dark"
				class="theme-slider"
			/>
			<span class="theme-label">🌙</span>
		</div>
	</div>
</footer>

<style>
	.route-shell {
		min-height: 100vh;
		animation: route-enter 520ms cubic-bezier(0.22, 1, 0.36, 1);
	}

	@keyframes route-enter {
		from {
			opacity: 0;
			transform: translateY(12px);
			filter: blur(8px);
		}

		to {
			opacity: 1;
			transform: translateY(0);
			filter: blur(0);
		}
	}

	@media (prefers-reduced-motion: reduce) {
		.route-shell {
			animation: none;
		}
	}

	.site-footer {
		position: fixed;
		bottom: 0;
		left: 0;
		right: 0;
		background: var(--surface-container-lowest);
		border-top: 1px solid var(--hairline);
		padding: 1rem;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 0.5rem;
		z-index: 10;
	}

	.site-footer > p {
		font-family: var(--font-mono);
		font-size: 0.75rem;
		letter-spacing: -0.01em;
		color: var(--on-surface);
		opacity: 0.8;
		text-transform: uppercase;
	}

	.footer-links {
		display: flex;
		gap: 1.5rem;
	}

	.footer-links a,
	.footer-links button {
		font-family: var(--font-mono);
		font-size: 0.75rem;
		letter-spacing: -0.01em;
		color: var(--primary);
		text-decoration: underline;
		text-decoration-thickness: 1px;
		text-underline-offset: 4px;
		transition: color 300ms ease;
		background: none;
		border: none;
		cursor: pointer;
		padding: 0;
	}

	.footer-links a:hover,
	.footer-links button:hover {
		color: var(--on-primary-container);
	}

	.theme-slider-container {
		display: flex;
		align-items: center;
		gap: 0.5rem;
	}

	.theme-label {
		font-family: var(--font-mono);
		font-size: 0.75rem;
		opacity: 0.6;
		transition: opacity 200ms ease;
	}

	.theme-slider {
		width: 80px;
		height: 4px;
		-webkit-appearance: none;
		appearance: none;
		background: var(--hairline);
		border-radius: 0;
		outline: none;
		cursor: pointer;
	}

	.theme-slider::-webkit-slider-thumb {
		-webkit-appearance: none;
		appearance: none;
		width: 12px;
		height: 12px;
		background: var(--primary);
		border: none;
		border-radius: 0;
		cursor: pointer;
		transition: transform 150ms ease;
	}

	.theme-slider::-webkit-slider-thumb:hover {
		transform: scale(1.2);
	}

	.theme-slider::-moz-range-thumb {
		width: 12px;
		height: 12px;
		background: var(--primary);
		border: none;
		border-radius: 0;
		cursor: pointer;
		transition: transform 150ms ease;
	}

	.theme-slider::-moz-range-thumb:hover {
		transform: scale(1.2);
	}

	@media (min-width: 768px) {
		.site-footer {
			flex-direction: row;
			justify-content: space-between;
			padding: 1rem 2rem;
			max-width: 1440px;
			margin: 0 auto;
		}
	}

	@media (max-width: 767px) {
		.site-footer {
			padding: 0.75rem 1rem;
		}

		.footer-links {
			gap: 1rem;
			flex-wrap: wrap;
			justify-content: center;
		}
	}
</style>
