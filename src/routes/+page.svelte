<script lang="ts">
    import AboutSection from "$lib/components/AboutSection.svelte";
    import ContactSection from "$lib/components/ContactSection.svelte";
    import HeroSection from "$lib/components/HeroSection.svelte";
    import ResumeSection from "$lib/components/ResumeSection.svelte";
    import SiteNav from "$lib/components/SiteNav.svelte";
    import portfolio from "$lib/data/portfolio";
    import { onMount } from "svelte";

    let activeSection = $state("home");

    // Easter egg: Track clicks on the brand logo area
    let clickCount = $state(0);
    const MAX_CLICKS = 5;
    const EASTER_EGG_URL =
        "https://www.youtube.com/watch?v=TMCYSjSN_mw&pp=ygUXc2hhcnVrIGtoYW4gdGlnZXIgc2NlbmU%3D";
    let showRedirectOverlay = $state(false);
    let showSuccessMessage = $state(false);

    onMount(() => {
        if (
            typeof window === "undefined" ||
            !("IntersectionObserver" in window)
        )
            return;

        const sections = [
            { id: "top", key: "home" },
            { id: "about", key: "about" },
            { id: "resume", key: "resume" },
            { id: "contact", key: "contact" },
        ];

        const visibilityState = new Map<string, number>();
        let rafId: number | null = null;

        const updateActiveSection = () => {
            let bestKey = "home";
            let bestRatio = 0;

            for (const [key, ratio] of visibilityState) {
                if (ratio > bestRatio) {
                    bestRatio = ratio;
                    bestKey = key;
                }
            }

            if (bestRatio > 0) {
                activeSection = bestKey;
            }
        };

        const observer = new IntersectionObserver(
            (entries) => {
                for (const entry of entries) {
                    const section = sections.find(
                        (s) => s.id === entry.target.id,
                    );
                    if (section) {
                        visibilityState.set(
                            section.key,
                            entry.isIntersecting ? entry.intersectionRatio : 0,
                        );
                    }
                }

                if (rafId) cancelAnimationFrame(rafId);
                rafId = requestAnimationFrame(updateActiveSection);
            },
            {
                threshold: [0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1],
                rootMargin: "-50px 0px -50px 0px",
            },
        );

        sections.forEach((s) => {
            const el = document.getElementById(s.id);
            if (el) {
                visibilityState.set(s.key, 0);
                observer.observe(el);
            }
        });

        return () => {
            if (rafId) cancelAnimationFrame(rafId);
            observer.disconnect();
            visibilityState.clear();
        };
    });

    // Handle the hidden click trigger on brand logo area
    function handleBrandClick(event: MouseEvent | KeyboardEvent) {
        clickCount++;

        // Show subtle visual feedback (flash effect)
        const target = event.currentTarget as HTMLElement;
        if (target) {
            target.style.filter = "brightness(1.5)";
            setTimeout(() => {
                target.style.filter = "";
            }, 100);
        }

        if (clickCount >= MAX_CLICKS) {
            clickCount = 0;
            showSuccessMessage = true;

            // Wait for success message, then redirect
            setTimeout(() => {
                window.location.href = EASTER_EGG_URL;
            }, 1500);
        } else if (clickCount > 0) {
            // Show progress indicator (subtle dots appear below brand)
            target.style.cursor = "pointer";
            setTimeout(() => {
                target.style.cursor = "";
            }, 500);
        }

        event.stopPropagation();
    }
</script>

<SiteNav name={portfolio.profile.name} {activeSection} />

<main class="main-content">
    <div class="section-shell">
        <HeroSection
            profile={portfolio.profile}
            resumeDownloadUrl={portfolio.resume.downloadUrl}
        />
    </div>

    <div class="section-shell">
        <AboutSection about={portfolio.about} />
    </div>

    <div class="section-shell">
        <ResumeSection resume={portfolio.resume} />
    </div>

    <div class="section-shell">
        <ContactSection contact={portfolio.contact} />
    </div>
</main>

<!-- Footer -->
<footer class="tactical-footer">
    <div class="footer-content">
        <span
            class="footer-brand tech-label"
            onclick={handleBrandClick}
            onkeydown={(e) => {
                if (e.key === "Enter" || e.key === " ") handleBrandClick(e);
            }}
            role="button"
            aria-label="Secret mission trigger - click 5 times"
            tabindex="0"
            data-easter-egg-trigger
        >
            <span class="copyright">©2499_</span>
            <span class="brand-primary">TACTICAL_ARCHIVE</span>
            <span class="version">_V.4.02</span>
        </span>

        <!-- Progress dots for easter egg (hidden until clicked) -->
        {#if clickCount > 0}
            <div class="easter-egg-progress">
                {#each Array(5) as _, i}
                    <span class={i < clickCount ? "dot active" : "dot"}></span>
                {/each}
            </div>
        {/if}

        <div class="footer-links">
            <a href="#top" class="tech-label footer-link">DECRYPT</a>
            <a href="#resume" class="tech-label footer-link">SYSTEM_LOGS</a>
        </div>
    </div>
</footer>

<!-- Success Message Overlay -->
{#if showSuccessMessage}
    <div class="easter-egg-overlay">
        <div class="success-message">
            <span class="message-icon">🐯</span>
            <h2>MISSION ACCOMPLISHED!</h2>
            <p>Finding the hidden tiger...</p>
        </div>
    </div>
{/if}

<!-- Mobile Footer -->
<footer class="mobile-footer">
    <div class="mobile-footer-content">
        <span
            class="mobile-footer-brand"
            onclick={handleBrandClick}
            onkeydown={(e) => {
                if (e.key === "Enter" || e.key === " ") handleBrandClick(e);
            }}
            role="button"
            aria-label="Secret mission trigger - click 5 times"
            tabindex="0"
        >
            TACTICAL_ARCHIVE_V.4
        </span>
        <div class="mobile-footer-divider"></div>
        <div class="mobile-footer-links">
            <a href="#top" class="mobile-footer-link">HOME</a>
            <a href="#resume" class="mobile-footer-link">LOGS</a>
        </div>
    </div>
</footer>

<style>
    .tactical-footer {
        position: fixed;
        bottom: 0;
        left: 80px;
        right: 0;
        height: 32px;
        background: rgba(13, 14, 15, 0.8);
        backdrop-filter: blur(10px);
        border-top: 1px solid rgba(255, 0, 60, 0.1);
        display: flex;
        align-items: center;
        padding: 0 1.5rem;
        z-index: 30;
    }

    .footer-content {
        display: flex;
        justify-content: space-between;
        align-items: center;
        width: 100%;
        gap: 2rem;
    }

    .footer-brand {
        display: flex;
        align-items: center;
        gap: 0.25rem;
        cursor: pointer;
        user-select: none;
        position: relative;
        padding: 4px 8px;
        border-radius: 4px;
        transition:
            filter 100ms ease,
            transform 100ms ease;
    }

    .footer-brand:focus {
        outline: 2px solid var(--primary);
        outline-offset: 4px;
    }

    .footer-brand:hover {
        filter: brightness(1.2);
        transform: scale(1.02);
    }

    .copyright {
        color: var(--on-surface-variant);
        opacity: 0.5;
    }

    .brand-primary {
        color: var(--primary);
    }

    .version {
        color: var(--tertiary);
        opacity: 0.7;
    }

    /* Progress dots indicator */
    .easter-egg-progress {
        display: flex;
        gap: 4px;
        align-items: center;
        padding-left: 1rem;
        border-left: 1px solid rgba(255, 0, 60, 0.3);
    }

    .dot {
        width: 8px;
        height: 8px;
        border-radius: 50%;
        background: var(--on-surface-variant);
        opacity: 0.3;
        transition: all 200ms ease;
    }

    .dot.active {
        background: var(--primary);
        opacity: 1;
        box-shadow: 0 0 8px var(--primary);
    }

    .footer-links {
        display: flex;
        gap: 1.5rem;
    }

    .footer-link {
        opacity: 0.4;
        transition: opacity 150ms ease;
        color: var(--inverse-primary);
        text-decoration: none;
    }

    .footer-link:hover {
        opacity: 1;
        color: var(--primary);
    }

    /* Mobile Footer */
    .mobile-footer {
        display: none;
        padding: 3rem 1rem;
        margin-top: 2rem;
        border-top: 1px solid rgba(255, 179, 178, 0.1);
    }

    .mobile-footer-content {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 1rem;
    }

    .mobile-footer-brand {
        font-family: "Space Grotesk", sans-serif;
        font-size: 0.75rem;
        letter-spacing: 0.15em;
        color: var(--on-surface);
        opacity: 0.5;
        cursor: pointer;
        user-select: none;
    }

    .mobile-footer-brand:hover {
        opacity: 0.8;
    }

    .mobile-footer-brand:focus {
        outline: 2px solid var(--primary);
        outline-offset: 4px;
    }

    .mobile-footer-divider {
        width: 2rem;
        height: 1px;
        background: rgba(255, 179, 178, 0.3);
    }

    .mobile-footer-links {
        display: flex;
        gap: 2rem;
    }

    .mobile-footer-link {
        font-family: "JetBrains Mono", monospace;
        font-size: 0.65rem;
        letter-spacing: 0.1em;
        color: var(--on-surface-variant);
        opacity: 0.6;
        text-decoration: none;
        transition: all 150ms ease;
    }

    .mobile-footer-link:hover {
        color: var(--primary);
        opacity: 1;
    }

    /* Easter Egg Overlay */
    .easter-egg-overlay {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background: rgba(0, 0, 0, 0.95);
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 1000;
        animation: fadeIn 300ms ease;
    }

    .success-message {
        text-align: center;
        padding: 2rem;
        color: var(--primary);
        font-family: "JetBrains Mono", monospace;
    }

    .message-icon {
        font-size: 4rem;
        display: block;
        margin-bottom: 1rem;
        animation: bounce 500ms ease;
    }

    .success-message h2 {
        font-size: 1.5rem;
        margin: 0 0 0.5rem 0;
        text-transform: uppercase;
        letter-spacing: 0.2em;
        background: linear-gradient(90deg, var(--primary), var(--tertiary));
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
        background-clip: text;
    }

    .success-message p {
        color: var(--on-surface-variant);
        font-size: 0.9rem;
        margin: 0;
        animation: pulse 1s ease infinite;
    }

    /* Responsive */
    @media (max-width: 1024px) {
        .tactical-footer {
            display: none;
        }

        .mobile-footer {
            display: flex;
            justify-content: center;
        }
    }

    @keyframes fadeIn {
        from {
            opacity: 0;
        }
        to {
            opacity: 1;
        }
    }

    @keyframes bounce {
        0%,
        20%,
        50%,
        80%,
        100% {
            transform: translateY(0);
        }
        40% {
            transform: translateY(-20px);
        }
        60% {
            transform: translateY(-10px);
        }
    }

    @keyframes pulse {
        0%,
        100% {
            opacity: 1;
        }
        50% {
            opacity: 0.5;
        }
    }
</style>
