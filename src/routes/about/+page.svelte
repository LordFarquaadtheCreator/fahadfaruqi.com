<script lang="ts">
	import AboutSection from '$lib/components/AboutSection.svelte';
	import GothicShell from '$lib/components/GothicShell.svelte';
	import portfolio from '$lib/data/portfolio';
</script>

<svelte:head>
	<title>About | Fahad Faruqi</title>
</svelte:head>

<GothicShell title="Fahad Faruqi">
	{#snippet children()}
		<div class="section-shell">
			<AboutSection about={portfolio.about} />
		</div>
	{/snippet}
</GothicShell>
