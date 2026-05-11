<script lang="ts">
    import GothicNavPanel from "$lib/components/GothicNavPanel.svelte";
    import portfolio from "$lib/data/portfolio";
</script>

<svelte:head>
    <title>Fahad Faruqi</title>
</svelte:head>

<div class="gothic-portal">
    <!-- Top metadata strip -->
    <div class="top-strip">
        <span class="tech-label strip-left">{portfolio.profile.location} · SOFTWARE_ENGINEER</span>
        <div class="barcode-wrap" aria-hidden="true">
            <svg class="barcode" viewBox="0 0 80 18" xmlns="http://www.w3.org/2000/svg">
                {#each Array(38) as _, i}
                    <rect
                        x={i * 2.05 + (i % 4 === 0 ? 0.4 : 0)}
                        y="0"
                        width={i % 5 === 0 ? 1.5 : 0.9}
                        height="18"
                        fill="rgba(157,78,221,0.5)"
                    />
                {/each}
            </svg>
        </div>
    </div>

    <!-- Main poster grid -->
    <div class="poster-grid">
        <!-- Left column -->
        <div class="col-side col-left">
            <GothicNavPanel href="/about" label="ABOUT_ME" />
            <div class="col-deco">
                <span class="sparkle">✦</span>
                <span class="hindi-text">फ़हाद</span>
                <span class="sparkle small">✧</span>
            </div>
            <GothicNavPanel href="/projects" label="PROJECTS" />
        </div>

        <!-- Center column -->
        <div class="col-center">
            <div class="title-wrap">
                <h1 class="gothic-title">Fahad<br />Faruqi</h1>
            </div>
            <div class="hero-frame">
                <div class="hud-bracket-tl"></div>
                <div class="hud-bracket-tr"></div>
                <div class="hud-bracket-bl"></div>
                <div class="hud-bracket-br"></div>
                <div class="hero-img-wrap">
                    <div class="hero-placeholder">
                        <span class="placeholder-text tech-label">HERO_IMAGE_PENDING</span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Right column -->
        <div class="col-side col-right">
            <GothicNavPanel href="/resume" label="RESUME" />
            <div class="col-deco col-deco-right">
                <span class="media-controls">◄◄ ◄ ■ ► ►►</span>
                <p class="deco-quote">{portfolio.profile.title}</p>
            </div>
    			<GothicNavPanel href="/contact" label="CONTACT" />
		</div>
	</div>

	<!-- Bottom strip -->
	<div class="bottom-strip">
		<span class="advisory tech-label">CLASSIFIED_PORTFOLIO</span>
		<p class="bottom-quote">{portfolio.profile.title}</p>
		<span class="tech-label bottom-tag">NYC · 2025</span>
	</div>
</div>

<style>
	.gothic-portal {
		min-height: 100vh;
		background: #09090b;
		display: flex;
		flex-direction: column;
		overflow-x: hidden;
	}

	/* Top strip */
	.top-strip {
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 0.4rem 1.25rem;
		border-bottom: 1px solid rgba(157, 78, 221, 0.25);
		background: rgba(9, 9, 11, 0.8);
	}

	.strip-left {
		color: rgba(157, 78, 221, 0.6);
		font-size: 0.55rem;
	}

	.barcode-wrap {
		display: flex;
		align-items: center;
	}

	.barcode {
		width: 72px;
		height: 16px;
	}

	/* Three-column poster grid */
	.poster-grid {
		display: grid;
		grid-template-columns: 1fr 2fr 1fr;
		flex: 1;
		border-bottom: 1px solid rgba(157, 78, 221, 0.2);
	}

	/* Side columns */
	.col-side {
		display: flex;
		flex-direction: column;
		gap: 0;
		border-right: 1px solid rgba(157, 78, 221, 0.2);
	}

	.col-right {
		border-right: none;
		border-left: 1px solid rgba(157, 78, 221, 0.2);
	}

	.col-deco {
		flex: 1;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 1.5rem 1rem;
		gap: 0.75rem;
		border-top: 1px solid rgba(157, 78, 221, 0.15);
		border-bottom: 1px solid rgba(157, 78, 221, 0.15);
	}

	.col-deco-right {
		text-align: center;
	}

	.sparkle {
		font-size: 1.25rem;
		color: rgba(157, 78, 221, 0.6);
	}

	.sparkle.small {
		font-size: 0.85rem;
		opacity: 0.5;
	}

	.hindi-text {
		font-family: 'Space Grotesk', sans-serif;
		font-size: 2.5rem;
		font-weight: 700;
		color: rgba(157, 78, 221, 0.15);
		line-height: 1;
		letter-spacing: 0.05em;
	}

	.media-controls {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.6rem;
		color: rgba(157, 78, 221, 0.5);
		letter-spacing: 0.15em;
	}

	.deco-quote {
		font-family: 'Work Sans', sans-serif;
		font-style: italic;
		font-size: 0.65rem;
		color: rgba(232, 213, 245, 0.3);
		line-height: 1.5;
		text-align: center;
		padding: 0 0.5rem;
	}

	/* Center column */
	.col-center {
		position: relative;
		display: flex;
		flex-direction: column;
	}

	.title-wrap {
		position: relative;
		z-index: 2;
		padding: 1rem 1.25rem 0;
	}

	.gothic-title {
		font-family: 'Pirata One', cursive;
		font-size: clamp(3rem, 8vw, 7.5rem);
		font-weight: 400;
		line-height: 0.88;
		color: #c084fc;
		text-transform: capitalize;
		text-shadow:
			3px 0 rgba(255, 0, 255, 0.45),
			-3px 0 rgba(0, 255, 255, 0.45),
			0 0 40px rgba(157, 78, 221, 0.3);
		margin: 0;
	}

	.hero-frame {
		position: relative;
		flex: 1;
		margin: -0.5rem 1.25rem 1.25rem;
		min-height: 320px;
	}

	.hero-img-wrap {
		width: 100%;
		height: 100%;
		background: rgba(157, 78, 221, 0.08);
		border: 1px solid rgba(157, 78, 221, 0.3);
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.hero-placeholder {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 0.5rem;
		padding: 2rem;
		height: 100%;
		min-height: 320px;
		background: repeating-linear-gradient(
			45deg,
			rgba(157, 78, 221, 0.04) 0px,
			rgba(157, 78, 221, 0.04) 1px,
			transparent 1px,
			transparent 10px
		);
	}

	.placeholder-text {
		color: rgba(157, 78, 221, 0.35);
		font-size: 0.55rem;
		letter-spacing: 0.2em;
	}

	/* HUD bracket overrides for hero */
	.hero-frame :global(.hud-bracket-tl),
	.hero-frame :global(.hud-bracket-tr),
	.hero-frame :global(.hud-bracket-bl),
	.hero-frame :global(.hud-bracket-br) {
		border-color: rgba(157, 78, 221, 0.6);
		width: 16px;
		height: 16px;
		z-index: 3;
	}

	/* Bottom strip */
	.bottom-strip {
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 0.6rem 1.25rem;
		border-top: 1px solid rgba(157, 78, 221, 0.25);
		background: rgba(9, 9, 11, 0.9);
		gap: 1rem;
	}

	.advisory {
		color: rgba(157, 78, 221, 0.5);
		font-size: 0.55rem;
		white-space: nowrap;
	}

	.bottom-quote {
		font-family: 'Work Sans', sans-serif;
		font-style: italic;
		font-size: 0.65rem;
		color: rgba(232, 213, 245, 0.3);
		margin: 0;
		text-align: center;
		flex: 1;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	.bottom-tag {
		color: rgba(157, 78, 221, 0.5);
		font-size: 0.55rem;
		white-space: nowrap;
	}

	/* Mobile */
	@media (max-width: 768px) {
		.poster-grid {
			grid-template-columns: 1fr;
		}

		.col-side {
			border-right: none;
			border-left: none;
			border-top: 1px solid rgba(157, 78, 221, 0.2);
		}

		.col-right {
			border-left: none;
		}

		.col-deco {
			flex-direction: row;
			padding: 1rem;
		}

		.hindi-text {
			font-size: 1.5rem;
		}

		.gothic-title {
			font-size: clamp(2.5rem, 14vw, 4rem);
		}

		.col-left {
			display: grid;
			grid-template-columns: 1fr 1fr;
		}

		.col-left .col-deco {
			grid-column: 1 / -1;
		}

		.col-right {
			display: grid;
			grid-template-columns: 1fr 1fr;
		}

		.col-right .col-deco {
			grid-column: 1 / -1;
		}

		.bottom-quote {
			display: none;
		}
	}
</style>
