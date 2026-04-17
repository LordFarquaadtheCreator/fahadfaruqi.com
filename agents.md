# fahadfaruqi.com — Agent Reference

Static portfolio site. SvelteKit 2 + Svelte 5 (Runes mode) + TypeScript. Compiled to static HTML via `adapter-static`. Deployed to GitHub Pages on push to `main`.

## Commands

```bash
npm run dev        # dev server
npm run build      # static output → build/
npm run preview    # serve build/
npm test           # vitest (unit + component)
npm run check      # svelte-check + tsc
```

## Stack

- **SvelteKit 2** / **Svelte 5** — Runes syntax throughout (`$props`, `$state`, `$derived`, `$effect`)
- **TypeScript** — strict types, no `any`
- **CSS Custom Properties** — single dark theme, no CSS framework
- **unplugin-icons** — Material Symbols, bundled at build time via `Icon.svelte` registry
- **Vitest** + `@testing-library/svelte` — component and unit tests
- **GitHub Actions** — `.github/workflows/deploy.yaml` auto-deploys on `main` push

## Critical Conventions

- **All content lives in `src/lib/data/portfolio.json`**. Never hardcode copy in components.
- `portfolio.ts` validates and exports the JSON. Validation throws at load time — bad JSON crashes the build.
- **0px border-radius site-wide** — enforced via `* { border-radius: 0 !important }` in `app.css`.
- Components use Svelte 5 Runes only. Never use legacy `export let` / `$:` syntax.
- Icons must be registered in `Icon.svelte` before use. Adding a new icon = add import + registry entry there.
- `use:reveal` action on elements that should fade in on scroll. Requires `.reveal` CSS class in scope.
- `.tech-label` = JetBrains Mono uppercase label. `.chromatic` = chromatic aberration text effect. `.chip` = tag pill. These are global utility classes in `app.css`.

## File Map

```
src/
├── app.css                  # Design system: CSS vars, global utilities, breakpoints
├── app.html                 # HTML shell
├── routes/
│   ├── +layout.svelte       # Font imports, favicon, theme-color meta
│   ├── +layout.ts           # prerender = true
│   ├── +page.svelte         # Composes all sections + footer + easter egg
│   └── page.svelte.spec.ts  # Renders page, asserts experience cards match JSON
├── lib/
│   ├── index.ts             # Re-exports portfolio + types from data/portfolio.ts
│   ├── actions/
│   │   └── reveal.ts        # IntersectionObserver scroll-reveal Svelte action
│   ├── assets/
│   │   └── favicon.svg
│   ├── components/          # See src/lib/components/agents.md
│   ├── data/                # See src/lib/data/agents.md
│   └── utils/
│       ├── mailto.ts        # buildMailtoHref() — builds mailto: URL from form fields
│       └── mailto.spec.ts
static/
├── images/
│   ├── companies/           # Company logo SVGs referenced in portfolio.json
│   └── profile/             # Portrait images
└── robots.txt
```

## Layout & Navigation

- **Desktop (>1024px)**: Fixed 80px side nav on left (`SiteNav`). Main content `margin-left: 80px`. Fixed 32px footer bar at bottom.
- **Mobile (≤1024px)**: Fixed top header + fixed bottom nav bar. Side nav hidden. Tactical footer hidden; mobile footer shown inline.
- `activeSection` state in `+page.svelte` drives nav highlighting via `IntersectionObserver` watching `#top`, `#about`, `#resume`, `#contact`.
- Section keys: `home` / `about` / `resume` / `contact`.

## Easter Egg

Footer brand (`TACTICAL_ARCHIVE`) is a hidden button. 5 clicks → progress dots appear → redirect to a YouTube URL. Both desktop and mobile footers share the same `handleBrandClick` handler defined in `+page.svelte`.

## Deployment

GitHub Actions workflow: `npm ci` → `npm run build` → upload `build/` → deploy to GitHub Pages. No env vars needed. `BASE_PATH` env var supported in `svelte.config.js` if needed for sub-path hosting.

## Testing

Two Vitest projects in `vite.config.ts`:
- `component` — jsdom, matches `*.svelte.{test,spec}.{js,ts}`
- `server` — node, matches `*.{test,spec}.{js,ts}` excluding svelte files
