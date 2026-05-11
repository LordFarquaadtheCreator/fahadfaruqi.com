<script lang="ts">
	import ContactSection from '$lib/components/ContactSection.svelte';
	import GothicShell from '$lib/components/GothicShell.svelte';
	import portfolio from '$lib/data/portfolio';
</script>

<svelte:head>
	<title>Contact | Fahad Faruqi</title>
</svelte:head>

<GothicShell title="Fahad Faruqi">
	{#snippet children()}
		<div class="section-shell">
			<ContactSection contact={portfolio.contact} />
		</div>
	{/snippet}
</GothicShell>
