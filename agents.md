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
- **CSS Custom Properties** — tri-state theme (system/light/dark), no CSS framework
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
- **Theme system**: Tri-state (system/light/dark) via `data-theme` attribute on `<html>`. Slider in footer: 0=light, 1=system, 2=dark (left to right). State persisted in `localStorage.theme`. Inline script in `app.html` prevents FOUC by applying theme before paint. CSS uses `:root[data-theme='dark']` selectors with `@media (prefers-color-scheme: dark)` fallback for JS-disabled. Store in `src/lib/theme.svelte.ts`.

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
│   ├── theme.svelte.ts      # Theme store (system/light/dark tri-state)
│   ├── theme.spec.ts        # Theme store unit tests
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

- **Sticky header**: `SiteNav` component provides three-part slider navigation (Home, CV, Blog) with active state highlighting.
- **Home page**: Centered nav with title, navigation links, and social icons. No SiteNav on home.
- **Sub-routes**: SiteNav sticky at top, main content below.
- Footer fixed at bottom with location, social links, email copy, and theme toggle.

## Deployment

GitHub Actions workflow: `npm ci` → `npm run build` → upload `build/` → deploy to GitHub Pages. No env vars needed. `BASE_PATH` env var supported in `svelte.config.js` if needed for sub-path hosting.

## Architecture Rule (MVVM, strict)

This app is strict MVVM. The Go WASM backend (`/interpreter`) owns ALL logic:
filesystem, parsing, zsh interpretation, window/Finder state, drag math,
view-mode layout, markdown rendering, focus, z-order, navigation history.

Svelte is a dumb view layer. It:
- Renders ViewModel JSON received from Go.
- Forwards raw input events (clicks, key, pointer deltas) to Go via `window.os.*`.
- Holds zero domain state. No `$derived` business logic, no parsing, no math.

The only Svelte state allowed: ephemeral UI affordances (hover/focus rings)
and a single `$state` holding the latest VM snapshot.

Every meaningful operation routes through Go.

## File Map (Extended)

```
interpreter/
├── finder/         # Window manager + VM + markdown render (see interpreter/finder/AGENTS.md)
├── command/open.go # zsh `open` entry point
└── ...

src/lib/
├── finder/         # Read-only VM mirror (see src/lib/finder/AGENTS.md)
└── components/finder/  # Finder UI (see src/lib/components/finder/AGENTS.md)
```

## Testing

Two Vitest projects in `vite.config.ts`:
- `component` — jsdom, matches `*.svelte.{test,spec}.{js,ts}`
- `server` — node, matches `*.{test,spec}.{js,ts}` excluding svelte files
