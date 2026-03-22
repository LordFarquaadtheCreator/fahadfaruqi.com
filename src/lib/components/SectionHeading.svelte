<script lang="ts">
	let {
		eyebrow,
		title,
		copy = '',
		align = 'left'
	} = $props<{
		eyebrow: string;
		title: string;
		copy?: string;
		align?: 'left' | 'center';
	}>();
</script>

<header class:section-heading--center={align === 'center'} class="section-heading">
	<p class="eyebrow">{eyebrow}</p>
	<h2>{title}</h2>
	{#if copy}
		<p class="section-copy">{copy}</p>
	{/if}
</header>

<style>
	.section-heading {
		display: grid;
		gap: 0.4rem;
	}

	.section-heading h2 {
		margin: 0;
		font-size: clamp(2rem, 3vw, 3.4rem);
		line-height: 0.98;
		letter-spacing: -0.04em;
	}

	.section-heading--center {
		justify-items: center;
		text-align: center;
	}
</style>
