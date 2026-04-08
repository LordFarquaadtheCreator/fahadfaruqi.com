<script lang="ts">
	import AboutSection from '$lib/components/AboutSection.svelte';
	import ContactSection from '$lib/components/ContactSection.svelte';
	import HeroSection from '$lib/components/HeroSection.svelte';
	import ResumeSection from '$lib/components/ResumeSection.svelte';
	import SiteNav from '$lib/components/SiteNav.svelte';
	import portfolio from '$lib/data/portfolio';
	import { onMount } from 'svelte';

	let activeSection = $state('home');

	onMount(() => {
		if (typeof window === 'undefined' || !('IntersectionObserver' in window)) return;

		const sections = [
			{ id: 'top', key: 'home' },
			{ id: 'about', key: 'about' },
			{ id: 'resume', key: 'resume' },
			{ id: 'contact', key: 'contact' }
		];

		// Track visibility state for all sections
		const visibilityState = new Map<string, number>();

		let rafId: number | null = null;

		const updateActiveSection = () => {
			// Find the section with highest visibility that exceeds threshold
			let bestKey = 'home';
			let bestRatio = 0;

			for (const [key, ratio] of visibilityState) {
				if (ratio > bestRatio) {
					bestRatio = ratio;
					bestKey = key;
				}
			}

			if (bestRatio > 0) {
				activeSection = bestKey;
			}
		};

		const observer = new IntersectionObserver(
			(entries) => {
				// Update visibility state for each entry
				for (const entry of entries) {
					const section = sections.find((s) => s.id === entry.target.id);
					if (section) {
						visibilityState.set(section.key, entry.isIntersecting ? entry.intersectionRatio : 0);
					}
				}

				// Debounce the active section update
				if (rafId) cancelAnimationFrame(rafId);
				rafId = requestAnimationFrame(updateActiveSection);
			},
			{ threshold: [0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1], rootMargin: '-50px 0px -50px 0px' }
		);

		sections.forEach((s) => {
			const el = document.getElementById(s.id);
			if (el) {
				visibilityState.set(s.key, 0);
				observer.observe(el);
			}
		});

		return () => {
			if (rafId) cancelAnimationFrame(rafId);
			observer.disconnect();
			visibilityState.clear();
		};
	});
</script>

<svelte:head>
	<title>{portfolio.profile.name.toLowerCase().replace(/\s+/g, '')}.com</title>
	<meta name="description" content={portfolio.profile.heroIntro} />
</svelte:head>

<SiteNav name={portfolio.profile.name} {activeSection} />

<main class="main-content">
	<div class="section-shell">
		<HeroSection profile={portfolio.profile} resumeDownloadUrl={portfolio.resume.downloadUrl} />
	</div>

	<div class="section-shell">
		<AboutSection about={portfolio.about} />
	</div>

	<div class="section-shell">
		<ResumeSection resume={portfolio.resume} />
	</div>

	<div class="section-shell">
		<ContactSection contact={portfolio.contact} />
	</div>
</main>

<!-- Footer -->
<footer class="tactical-footer">
	<div class="footer-content">
		<span class="footer-brand tech-label">
			<span class="copyright">©2499_</span>
			<span class="brand-primary">TACTICAL_ARCHIVE</span>
			<span class="version">_V.4.02</span>
		</span>
		<div class="footer-links">
			<a href="#top" class="tech-label footer-link">DECRYPT</a>
			<a href="#resume" class="tech-label footer-link">SYSTEM_LOGS</a>
		</div>
	</div>
</footer>

<!-- Mobile Footer -->
<footer class="mobile-footer">
	<div class="mobile-footer-content">
		<span class="mobile-footer-brand">TACTICAL_ARCHIVE_V.4</span>
		<div class="mobile-footer-divider"></div>
		<div class="mobile-footer-links">
			<a href="#top" class="mobile-footer-link">HOME</a>
			<a href="#resume" class="mobile-footer-link">LOGS</a>
		</div>
	</div>
</footer>

<style>
	.tactical-footer {
		position: fixed;
		bottom: 0;
		left: 80px;
		right: 0;
		height: 32px;
		background: rgba(13, 14, 15, 0.8);
		backdrop-filter: blur(10px);
		border-top: 1px solid rgba(255, 0, 60, 0.1);
		display: flex;
		align-items: center;
		padding: 0 1.5rem;
		z-index: 30;
	}

	.footer-content {
		display: flex;
		justify-content: space-between;
		align-items: center;
		width: 100%;
	}

	.footer-brand {
		display: flex;
		align-items: center;
		gap: 0.25rem;
	}

	.copyright {
		color: var(--on-surface-variant);
		opacity: 0.5;
	}

	.brand-primary {
		color: var(--primary);
	}

	.version {
		color: var(--tertiary);
		opacity: 0.7;
	}

	.footer-links {
		display: flex;
		gap: 1.5rem;
	}

	.footer-link {
		opacity: 0.4;
		transition: opacity 150ms ease;
		color: var(--inverse-primary);
	}

	.footer-link:hover {
		opacity: 1;
		color: var(--primary);
	}

	/* Mobile Footer */
	.mobile-footer {
		display: none;
		padding: 3rem 1rem;
		margin-top: 2rem;
		border-top: 1px solid rgba(255, 179, 178, 0.1);
	}

	.mobile-footer-content {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1rem;
	}

	.mobile-footer-brand {
		font-family: 'Space Grotesk', sans-serif;
		font-size: 0.75rem;
		letter-spacing: 0.15em;
		color: var(--on-surface);
		opacity: 0.5;
	}

	.mobile-footer-divider {
		width: 2rem;
		height: 1px;
		background: rgba(255, 179, 178, 0.3);
	}

	.mobile-footer-links {
		display: flex;
		gap: 2rem;
	}

	.mobile-footer-link {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.65rem;
		letter-spacing: 0.1em;
		color: var(--on-surface-variant);
		opacity: 0.6;
		text-decoration: none;
		transition: all 150ms ease;
	}

	.mobile-footer-link:hover {
		color: var(--primary);
		opacity: 1;
	}

	/* Responsive - Desktop footer hidden on mobile, mobile footer shown */
	@media (max-width: 1024px) {
		.tactical-footer {
			display: none;
		}

		.mobile-footer {
			display: flex;
			justify-content: center;
		}
	}
</style>
