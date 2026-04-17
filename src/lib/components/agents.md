# src/lib/components — Agent Reference

All components use Svelte 5 Runes. No legacy `export let` or `$:` syntax.

## Active Page-Level Sections

Used in `src/routes/+page.svelte`.

| Component | Section ID | Prop | Purpose |
|-----------|------------|------|---------|
| `HeroSection.svelte` | `#top` / `#top-mobile` | `profile`, `resumeDownloadUrl?` | Bento-grid navigation hub. Desktop + mobile layouts toggled via `@media (max-width: 1024px)`. |
| `AboutSection.svelte` | `#about` | `about: AboutData` | Bio paragraphs, education records, highlights grid. 12-col CSS grid. |
| `ResumeSection.svelte` | `#resume` | `resume: ResumeData` | Work experience cards (`data-testid="experience-card"`). Scroll-revealed. Optional resume download link. |
| `ContactSection.svelte` | `#contact` | `contact: ContactData` | Registry panel + mailto transmission form. Form state is local `$state`; mailto href is `$derived` from `buildMailtoHref()`. |
| `SiteNav.svelte` | — | `name: string`, `activeSection?: string` | Desktop: fixed 80px left sidebar. Mobile: fixed top header + bottom nav bar. `activeSection` drives active highlight. Default fallback is `'ops'` but actual keys are `home/about/resume/contact`. |

## Design System Primitives

| Component | Purpose | Key Props |
|-----------|---------|-----------|
| `TacticalCard.svelte` | Reusable surface card with optional HUD brackets, header icon/label. Always `use:reveal`. | `variant: 'default'|'primary'|'highlight'|'stream'`, `hudBrackets`, `headerIcon`, `headerLabel`, `headerRight` (Snippet), `class` |
| `TacticalSectionHeader.svelte` | Section header with ghost background text. Always `use:reveal`. | `title` (required), `ghostText`, `eyebrow`, `description`, `headerRight` (Snippet), `variant: 'default'|'uplink'` |
| `Icon.svelte` | Typed icon wrapper over `unplugin-icons` Material Symbols. Silently renders nothing for unknown names. | `name: string`, `class?`, `style?`, `ariaLabel?` |

## Unused / Legacy Components

Not imported in `+page.svelte`. Do not add them back without discussion.

| Component | Notes |
|-----------|-------|
| `PhotoSection.svelte` | Portrait + caption display. Depends on `MediaFrame` and `SectionHeading`. |
| `SectionHeading.svelte` | Generic eyebrow + title + copy header. Only used by `PhotoSection`. |
| `MediaFrame.svelte` | Image container with initials fallback. Only used by `PhotoSection`. |

## Adding a New Icon

1. Import from `~icons/material-symbols/<name>` at the top of `Icon.svelte`
2. Add to the `icons` record with the snake_case key used in `name` prop
3. Use anywhere via `<Icon name="your_icon" />`

## Svelte 5 Patterns Used

```svelte
<!-- Props -->
let { foo, bar = 'default' } = $props<{ foo: string; bar?: string }>();

<!-- Reactive state -->
let value = $state('');

<!-- Derived -->
const computed = $derived(someExpression);

<!-- Side effects -->
$effect(() => { ...; return () => cleanup(); });

<!-- Snippets (render prop equivalent) -->
{#snippet headerRight()}...{/snippet}
{@render headerRight()}
```

## Scroll Reveal

Apply `use:reveal` to any element. Also add `.reveal` class (CSS in `app.css` sets initial opacity/transform). The action sets `data-revealed="true"` when element enters viewport at 18% threshold. Respects `prefers-reduced-motion`.

## HeroSection Notes

- `profile` prop declared but profile text fields (`name`, `heroIntro`, etc.) are NOT currently rendered. The graphic area has a placeholder (`<p>my image image placeholder</p>`). Portrait integration is incomplete.
- `resumeDownloadUrl` is passed through but not used inside `HeroSection` — the Resume nav card is a simple `<a href="#resume">` link, not a download.

## Design Constraints

- **0px border-radius everywhere** — `* { border-radius: 0 !important }` in `app.css`
- **Glitch hover** — `.tactical-card:hover` and `.op-card:hover` run a `@keyframes glitch` translate jitter
- **HUD brackets** (`.hud-bracket-tl/tr/bl/br`) — global utility in `app.css`, used via `hudBrackets` prop or directly in markup
- **Chromatic aberration** — `.chromatic` class adds red/cyan text-shadow offset
- **Jitter animation** — `@keyframes jitter` applied to `.ref-label`, `.status-value`, etc. for "live terminal" feel
- Color roles: `--primary` = signal red, `--secondary` = electric blue, `--tertiary` = radioactive green
