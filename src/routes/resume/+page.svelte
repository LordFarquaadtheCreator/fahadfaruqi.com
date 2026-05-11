<script lang="ts">
	import GothicShell from '$lib/components/GothicShell.svelte';
	import ResumeSection from '$lib/components/ResumeSection.svelte';
	import portfolio from '$lib/data/portfolio';
</script>

<svelte:head>
	<title>Resume | Fahad Faruqi</title>
</svelte:head>

<GothicShell title="Fahad Faruqi">
	{#snippet children()}
		<div class="section-shell">
			<ResumeSection resume={portfolio.resume} />
		</div>
	{/snippet}
</GothicShell>
