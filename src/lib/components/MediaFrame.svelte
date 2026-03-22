<script lang="ts">
	let {
		src = '',
		alt,
		label,
		variant = 'logo'
	} = $props<{
		src?: string;
		alt: string;
		label: string;
		variant?: 'logo' | 'portrait';
	}>();

	let failed = $state(false);

	$effect(() => {
		failed = !src;
	});

	function handleError() {
		failed = true;
	}

	function getInitials(value: string) {
		return value
			.split(/\s+/)
			.map((part) => part.slice(0, 1))
			.join('')
			.slice(0, 2)
			.toUpperCase();
	}
</script>

<div class:media-frame--portrait={variant === 'portrait'} class="media-frame">
	{#if failed}
		<div
			class="media-frame__placeholder"
			data-testid="media-placeholder"
			aria-label={`${label} placeholder`}
		>
			<span>{getInitials(label)}</span>
			<small>{label}</small>
		</div>
	{:else}
		<img class="media-frame__image" {src} {alt} loading="lazy" onerror={handleError} />
	{/if}
</div>

<style>
	.media-frame {
		aspect-ratio: 16 / 10;
		width: 100%;
		overflow: hidden;
		border: 1px solid rgba(34, 22, 13, 0.1);
		border-radius: 1.5rem;
		background:
			linear-gradient(140deg, rgba(255, 255, 255, 0.74), rgba(255, 255, 255, 0.44)),
			radial-gradient(circle at top right, rgba(179, 92, 46, 0.2), transparent 45%);
	}

	.media-frame--portrait {
		aspect-ratio: 4 / 5;
	}

	.media-frame__image {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.media-frame__placeholder {
		display: grid;
		place-items: center;
		height: 100%;
		padding: 1.5rem;
		text-align: center;
		background:
			radial-gradient(circle at 30% 30%, rgba(179, 92, 46, 0.18), transparent 34%),
			radial-gradient(circle at 78% 76%, rgba(78, 111, 136, 0.16), transparent 28%),
			linear-gradient(180deg, rgba(255, 255, 255, 0.82), rgba(246, 239, 230, 0.82));
	}

	.media-frame__placeholder span {
		display: inline-grid;
		place-items: center;
		width: clamp(4rem, 18vw, 6rem);
		height: clamp(4rem, 18vw, 6rem);
		border-radius: 1.5rem;
		background: rgba(255, 255, 255, 0.84);
		color: var(--accent-deep);
		font-size: clamp(1.4rem, 4vw, 2rem);
		font-weight: 700;
		letter-spacing: 0.08em;
		box-shadow: 0 14px 30px rgba(38, 24, 14, 0.08);
	}

	.media-frame__placeholder small {
		margin-top: 0.9rem;
		font-size: 0.92rem;
		font-weight: 600;
		color: var(--text-soft);
	}
</style>
