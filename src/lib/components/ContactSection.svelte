<script lang="ts">
	import { reveal } from '$lib/actions/reveal';
	import type { ContactData } from '$lib/data/portfolio';
	import { buildMailtoHref } from '$lib/utils/mailto';
	import SectionHeading from './SectionHeading.svelte';

	let { contact } = $props<{ contact: ContactData }>();

	let name = $state('');
	let fromEmail = $state('');
	let message = $state('');

	const mailtoHref = $derived(
		buildMailtoHref({
			email: contact.email,
			subjectPrefix: contact.subjectPrefix,
			name,
			fromEmail,
			message
		})
	);
</script>

<section class="contact section-shell reveal" id="contact" use:reveal>
	<div class="surface contact__panel">
		<div class="contact__intro">
			<SectionHeading
				eyebrow="Contact"
				title="Let’s make something excellent."
				copy={contact.availability}
			/>

			<ul class="contact__links">
				{#each contact.links as link}
					<li>
						<a href={link.href} target={link.href.startsWith('http') ? '_blank' : undefined} rel="noreferrer">
							{link.label}
						</a>
					</li>
				{/each}
			</ul>
		</div>

		<form class="contact__form">
			<label>
				<span>Your name</span>
				<input bind:value={name} name="name" placeholder="What should I call you?" type="text" />
			</label>

			<label>
				<span>Your email</span>
				<input bind:value={fromEmail} name="email" placeholder="Where should I reply?" type="email" />
			</label>

			<label>
				<span>Your message</span>
				<textarea
					bind:value={message}
					name="message"
					rows="7"
					placeholder="Tell me a little about the role, team, or project."
				></textarea>
			</label>

			<div class="contact__actions">
				<a class="button" data-testid="contact-draft-link" href={mailtoHref}>
					Open email draft
				</a>
				<a class="button button--ghost" href={`mailto:${contact.email}`}>Email directly</a>
			</div>
		</form>
	</div>
</section>

<style>
	.contact {
		padding: 2rem 0 0;
	}

	.contact__panel {
		display: grid;
		grid-template-columns: minmax(0, 0.9fr) minmax(0, 1.1fr);
		gap: 1.4rem;
		padding: clamp(1.4rem, 3vw, 2.2rem);
	}

	.contact__intro {
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		gap: 1.2rem;
	}

	.contact__links {
		display: flex;
		flex-wrap: wrap;
		gap: 0.85rem;
		padding: 0;
		margin: 0;
		list-style: none;
	}

	.contact__links a {
		display: inline-flex;
		padding: 0.65rem 0.92rem;
		border: 1px solid var(--line);
		border-radius: var(--radius-sm);
		background: rgba(255, 255, 255, 0.66);
		font-weight: 600;
		color: var(--text-soft);
	}

	.contact__form {
		display: grid;
		gap: 1rem;
	}

	.contact__form label {
		display: grid;
		gap: 0.45rem;
	}

	.contact__form span {
		font-size: 0.92rem;
		font-weight: 700;
	}

	.contact__form input,
	.contact__form textarea {
		width: 100%;
		padding: 0.95rem 1rem;
		border: 1px solid rgba(45, 32, 23, 0.12);
		border-radius: 1.1rem;
		background: rgba(255, 255, 255, 0.8);
		color: var(--text);
	}

	.contact__form textarea {
		resize: vertical;
		min-height: 12rem;
	}

	.contact__actions {
		display: flex;
		flex-wrap: wrap;
		gap: 0.8rem;
		padding-top: 0.3rem;
	}

	@media (max-width: 900px) {
		.contact__panel {
			grid-template-columns: 1fr;
		}
	}
</style>
