<script lang="ts">
	import AboutSection from '$lib/components/AboutSection.svelte';
	import ContactSection from '$lib/components/ContactSection.svelte';
	import HeroSection from '$lib/components/HeroSection.svelte';
	import ResumeSection from '$lib/components/ResumeSection.svelte';
	import SiteNav from '$lib/components/SiteNav.svelte';
	import portfolio from '$lib/data/portfolio';
	import { onMount } from 'svelte';

	let activeSection = $state('ops');

	onMount(() => {
		if (typeof window === 'undefined' || !('IntersectionObserver' in window)) return;

		const sections = [
			{ id: 'top', key: 'ops' },
			{ id: 'about', key: 'intel' },
			{ id: 'resume', key: 'archive' },
			{ id: 'contact', key: 'uplink' }
		];

		const observer = new IntersectionObserver(
			(entries) => {
				entries.forEach((entry) => {
					if (entry.isIntersecting) {
						const section = sections.find((s) => s.id === entry.target.id);
						if (section) activeSection = section.key;
					}
				});
			},
			{ threshold: 0.3, rootMargin: '-10% 0px -50% 0px' }
		);

		sections.forEach((s) => {
			const el = document.getElementById(s.id);
			if (el) observer.observe(el);
		});

		return () => observer.disconnect();
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
		<span class="tech-label" style="color: var(--inverse-primary);">
			©2499_TACTICAL_ARCHIVE_V.01
		</span>
		<div class="footer-links">
			<a href="#top" class="tech-label footer-link">DECRYPT</a>
			<a href="#resume" class="tech-label footer-link">SYSTEM_LOGS</a>
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
	}

	@media (max-width: 768px) {
		.tactical-footer {
			left: 0;
		}
	}
</style>
