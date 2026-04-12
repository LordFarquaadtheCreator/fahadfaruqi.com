<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.png';
	import '@fontsource/space-grotesk/400.css';
	import '@fontsource/space-grotesk/500.css';
	import '@fontsource/space-grotesk/600.css';
	import '@fontsource/space-grotesk/700.css';
	import '@fontsource/work-sans/400.css';
	import '@fontsource/work-sans/500.css';
	import '@fontsource/work-sans/600.css';
	import '@fontsource/jetbrains-mono/400.css';
	import '@fontsource/jetbrains-mono/700.css';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<meta name="theme-color" content="#121314" />
</svelte:head>

{@render children()}
