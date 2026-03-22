<script lang="ts">
	import { reveal } from '$lib/actions/reveal';
	import type { ProfileData } from '$lib/data/portfolio';
	import MediaFrame from './MediaFrame.svelte';
	import SectionHeading from './SectionHeading.svelte';

	let {
		profile,
		caption
	} = $props<{
		profile: ProfileData;
		caption: string;
	}>();
</script>

<section class="photo section-shell reveal" id="photo" use:reveal>
	<div class="photo__copy">
		<SectionHeading
			eyebrow="Portrait"
			title="A dedicated place for your photo."
			copy="The layout is already framed for a real portrait, with a graceful fallback until you drop the image path into the JSON data."
		/>
		<p>{caption}</p>
	</div>

	<div class="surface photo__frame">
		<MediaFrame
			src={profile.portraitImage}
			alt={profile.portraitAlt}
			label={profile.name}
			variant="portrait"
		/>
	</div>
</section>

<style>
	.photo {
		display: grid;
		grid-template-columns: minmax(0, 1fr) minmax(18rem, 24rem);
		gap: 1.5rem;
		align-items: center;
		padding: 2rem 0;
	}

	.photo__copy p:last-child {
		max-width: 34rem;
		margin: 1rem 0 0;
		color: var(--text-soft);
	}

	.photo__frame {
		padding: 1rem;
	}

	@media (max-width: 880px) {
		.photo {
			grid-template-columns: 1fr;
		}
	}
</style>
