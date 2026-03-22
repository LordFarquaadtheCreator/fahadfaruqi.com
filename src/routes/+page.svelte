<script lang="ts">
	import AboutSection from '$lib/components/AboutSection.svelte';
	import ContactSection from '$lib/components/ContactSection.svelte';
	import HeroSection from '$lib/components/HeroSection.svelte';
	import PhotoSection from '$lib/components/PhotoSection.svelte';
	import ResumeSection from '$lib/components/ResumeSection.svelte';
	import SiteNav from '$lib/components/SiteNav.svelte';
	import portfolio from '$lib/data/portfolio';

	const navLinks = [
		{ label: 'About', href: '#about' },
		{ label: 'Photo', href: '#photo' },
		{ label: 'Resume', href: '#resume' },
		{ label: 'Contact', href: '#contact' }
	];
</script>

<svelte:head>
	<title>{portfolio.profile.name} | Portfolio</title>
	<meta name="description" content={portfolio.profile.heroIntro} />
</svelte:head>

<SiteNav name={portfolio.profile.name} links={navLinks} socialLinks={portfolio.profile.socialLinks} />

<main>
	<div class="page-glow page-glow--warm" aria-hidden="true"></div>
	<div class="page-glow page-glow--cool" aria-hidden="true"></div>

	<HeroSection profile={portfolio.profile} resumeDownloadUrl={portfolio.resume.downloadUrl} />
	<AboutSection about={portfolio.about} />
	<PhotoSection profile={portfolio.profile} caption={portfolio.about.photoCaption} />
	<ResumeSection resume={portfolio.resume} />
	<ContactSection contact={portfolio.contact} />
</main>

<style>
	main {
		position: relative;
		overflow: clip;
		padding-bottom: 5rem;
	}

	.page-glow {
		position: absolute;
		z-index: -1;
		filter: blur(12px);
		opacity: 0.75;
		pointer-events: none;
	}

	.page-glow--warm {
		top: 7rem;
		right: -12rem;
		width: 34rem;
		height: 34rem;
		border-radius: 50%;
		background:
			radial-gradient(circle at 40% 40%, rgba(242, 170, 123, 0.55), transparent 56%),
			radial-gradient(circle at 70% 65%, rgba(222, 119, 74, 0.28), transparent 42%);
	}

	.page-glow--cool {
		top: 64rem;
		left: -14rem;
		width: 28rem;
		height: 28rem;
		border-radius: 50%;
		background:
			radial-gradient(circle at 35% 50%, rgba(83, 126, 158, 0.24), transparent 54%),
			radial-gradient(circle at 70% 30%, rgba(133, 82, 55, 0.18), transparent 50%);
	}
</style>
