<script lang="ts">
	import { reveal } from '$lib/actions/reveal';
	import type { ContactData } from '$lib/data/portfolio';
	import { buildMailtoHref } from '$lib/utils/mailto';
	import Icon from './Icon.svelte';

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

	const packetSize = $derived((name.length + fromEmail.length + message.length) / 1024);
</script>

<section class="uplink-section" id="contact">
	<!-- Background Coordinates -->
	<div class="bg-coords" aria-hidden="true">
		<span class="coords-text">40.7128° N<br/>74.0060° W</span>
	</div>

	<!-- Section Header -->
	<div class="uplink-header reveal" use:reveal>
		<h2 class="uplink-title chromatic">Contact</h2>
	</div>

	<!-- Main Uplink Interface -->
	<div class="uplink-grid">
		<!-- Left Panel - Registry -->
		<div class="uplink-panel panel-registry reveal" use:reveal>
			<div class="hud-bracket-tl"></div>
			<div class="hud-bracket-tr"></div>
			<div class="hud-bracket-bl"></div>
			<div class="hud-bracket-br"></div>

			<div class="panel-header">
				<Icon name="database" class="panel-header__icon" style="color: var(--secondary);" />
				<span class="tech-label">COMMS_REGISTRY</span>
			</div>

			<div class="registry-data">
				<div class="registry-row">
					<span class="tech-label registry-key">OPERATOR_SIG</span>
					<span class="registry-value" style="color: var(--secondary);">#ARCHIVE_01</span>
				</div>
				<div class="registry-row">
					<span class="tech-label registry-key">UPLINK_STRENGTH</span>
					<span class="registry-value" style="color: var(--tertiary);">98.4%</span>
				</div>
				<div class="registry-row">
					<span class="tech-label registry-key">RESPONSE_ETA</span>
					<span class="registry-value" style="color: var(--primary);">24_HRS_MAX</span>
				</div>
			</div>

			<div class="feed-panel">
				<span class="tech-label feed-label">Global_Status_Feed</span>
				<div class="feed-content">
					<p><span class="feed-time">[14:02:11]</span> INITIALIZING HANDSHAKE...</p>
					<p><span class="feed-time">[14:02:12]</span> AUTH_TOKEN_VALIDATED</p>
					<p><span class="feed-time">[14:02:15]</span> PACKET_FILTER_BYPASS: OK</p>
					<p><span class="feed-time">[14:02:18]</span> NODE_09_RESPONDING</p>
					<p><span class="feed-time">[14:02:22]</span> SYNCING_TIME_PROTOCOLS...</p>
					<p class="feed-blink"><span class="feed-time">[14:02:25]</span> READY_FOR_INPUT_</p>
				</div>
			</div>
		</div>

		<!-- Right Panel - Transmission Form -->
		<div class="uplink-panel panel-transmission reveal" use:reveal>
			<div class="panel-header">
				<Icon name="send" class="panel-header__icon" style="color: var(--primary);" />
				<span class="tech-label">TRANSMISSION_FORM</span>
			</div>

			<form class="transmission-form">
				<div class="form-row">
					<div class="form-field">
						<label for="identity" class="field-label">
							<span class="tech-label">IDENTITY_STRING</span>
							<span class="required-tag" style="color: var(--primary);">REQUIRED</span>
						</label>
						<input
							id="identity"
							type="text"
							class="input-terminal"
							placeholder="ENTER_DESIGNATION..."
							bind:value={name}
						/>
					</div>

					<div class="form-field">
						<label for="email" class="field-label">
							<span class="tech-label">COMM_PROTOCOL</span>
							<span class="required-tag" style="color: var(--primary);">REQUIRED</span>
						</label>
						<input
							id="email"
							type="email"
							class="input-terminal"
							placeholder="ENTER_EMAIL..."
							bind:value={fromEmail}
						/>
					</div>
				</div>

				<div class="form-field">
					<label for="message" class="field-label">
						<span class="tech-label">TRANSMISSION_CONTENT</span>
						<span class="required-tag" style="color: var(--primary);">REQUIRED</span>
					</label>
					<textarea
						id="message"
						class="input-terminal"
						rows="5"
						placeholder="INPUT_DATA_PACKAGE..."
						bind:value={message}
					></textarea>
				</div>

				<div class="form-footer">
					<div class="transmission-meta">
						<div class="meta-item">
							<span class="status-pip status-active"></span>
							<span class="tech-label" style="color: var(--tertiary);">SIGNAL: SECURE</span>
						</div>
						<span class="tech-label" style="color: var(--outline);">|</span>
						<span class="tech-label">PACKET_SIZE: {packetSize.toFixed(1)}KB</span>
					</div>

					<a href={mailtoHref} class="button" data-testid="contact-draft-link">
						SEND_TRANSMISSION
					</a>
				</div>
			</form>

			<!-- Corner Coordinates -->
			<div class="corner-coords tech-label">
				X: 99.0029 // Y: 14.2881
			</div>
		</div>
	</div>

	<!-- Footer Archive Note -->
	<div class="archive-footer reveal" use:reveal>
		<div class="footer-left">
			<div class="qr-placeholder">
				<Icon name="qr_code" class="qr-placeholder__icon" />
			</div>
			<div class="footer-meta">
				<span class="tech-label">ENCRYPTED_SIGNATURE:</span>
				<span class="tech-label" style="color: var(--outline);">9f8a-221c-bd00-ee41</span>
				<span class="tech-label">ARCHIVE_VER: 0.9.2-BETA</span>
			</div>
		</div>
		<div class="footer-right">
			<span class="ghost-label">TACTICAL_CONTACT</span>
		</div>
	</div>
</section>

<style>
	.uplink-section {
		padding: 4rem 0;
		position: relative;
		overflow: hidden;
	}

	.bg-coords {
		position: absolute;
		top: 2rem;
		right: 0;
		opacity: 0.1;
		pointer-events: none;
	}

	.coords-text {
		font-family: 'JetBrains Mono', monospace;
		font-size: clamp(3rem, 8vw, 6rem);
		line-height: 1;
		color: var(--on-surface);
		user-select: none;
	}

	.uplink-header {
		margin-bottom: 3rem;
		position: relative;
	}

	.uplink-title {
		font-size: clamp(2rem, 5vw, 3.5rem);
		font-weight: 900;
		letter-spacing: -0.02em;
		line-height: 1;
		color: var(--on-surface);
		margin: 0;
	}

	.uplink-grid {
		display: grid;
		grid-template-columns: 1fr 2fr;
		gap: 1.5rem;
	}

	.uplink-panel {
		background: var(--surface-container-low);
		border: 1px solid rgba(175, 135, 134, 0.15);
		padding: 1.5rem;
		position: relative;
		transition: all 200ms ease;
	}

	.uplink-panel:hover {
		border-color: rgba(42, 229, 0, 0.3);
		box-shadow: 0 0 30px rgba(42, 229, 0, 0.08);
	}

	.panel-header {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		margin-bottom: 1.5rem;
		padding-bottom: 0.75rem;
		border-bottom: 1px solid rgba(175, 135, 134, 0.1);
	}

	:global(.panel-header__icon) {
		font-size: 20px;
	}

	/* Registry Panel */
	.registry-data {
		display: flex;
		flex-direction: column;
		gap: 0.75rem;
		margin-bottom: 1.5rem;
	}

	.registry-row {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding-bottom: 0.5rem;
		border-bottom: 1px solid var(--outline-variant);
	}

	.registry-key {
		color: var(--on-surface-variant);
	}

	.registry-value {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.7rem;
	}

	.feed-panel {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}

	.feed-label {
		color: var(--outline);
	}

	.feed-content {
		background: var(--surface-container-lowest);
		padding: 1rem;
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.65rem;
		color: var(--on-surface-variant);
		line-height: 1.8;
		max-height: 160px;
		overflow: hidden;
	}

	.feed-content p {
		margin: 0;
	}

	.feed-time {
		color: var(--error);
	}

	.feed-blink {
		animation: blink 1s step-end infinite;
	}

	@keyframes blink {
		50% { opacity: 0; }
	}

	/* Transmission Panel */
	.transmission-form {
		display: flex;
		flex-direction: column;
		gap: 1.25rem;
	}

	.form-row {
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		gap: 1rem;
	}

	.form-field {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}

	.field-label {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.required-tag {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.6rem;
	}

	.form-footer {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-top: 1rem;
		padding-top: 1rem;
		border-top: 1px solid rgba(175, 135, 134, 0.1);
		flex-wrap: wrap;
		gap: 1rem;
	}

	.transmission-meta {
		display: flex;
		align-items: center;
		gap: 0.75rem;
	}

	.meta-item {
		display: flex;
		align-items: center;
		gap: 0.5rem;
	}

	.corner-coords {
		position: absolute;
		bottom: 0;
		right: 0;
		background: var(--background);
		padding: 0.25rem 0.5rem;
		color: var(--outline-variant);
	}

	/* Archive Footer */
	.archive-footer {
		display: flex;
		justify-content: space-between;
		align-items: flex-end;
		margin-top: 2rem;
		padding-top: 1.5rem;
		border-top: 1px solid rgba(175, 135, 134, 0.1);
	}

	.footer-left {
		display: flex;
		gap: 1rem;
		align-items: center;
	}

	.qr-placeholder {
		width: 64px;
		height: 64px;
		display: flex;
		align-items: center;
		justify-content: center;
		background: var(--surface-container);
		border: 1px solid var(--outline-variant);
		opacity: 0.3;
	}

	:global(.qr-placeholder__icon) {
		font-size: 32px;
		color: var(--on-surface);
	}

	.footer-meta {
		display: flex;
		flex-direction: column;
		gap: 0.25rem;
	}

	.ghost-label {
		font-family: 'Space Grotesk', sans-serif;
		font-size: 2rem;
		font-weight: 900;
		color: var(--outline);
		opacity: 0.1;
		letter-spacing: -0.02em;
	}

	/* Responsive */
	@media (max-width: 1024px) {
		.uplink-grid {
			grid-template-columns: 1fr;
		}

		.uplink-title {
			font-size: clamp(3rem, 15vw, 6rem);
		}
	}

	@media (max-width: 768px) {
		.form-row {
			grid-template-columns: 1fr;
		}

		.bg-coords {
			display: none;
		}

		.archive-footer {
			flex-direction: column;
			gap: 1rem;
			align-items: flex-start;
		}

		.ghost-label {
			display: none;
		}
	}
</style>
