# src/

Source root. SvelteKit 2 + Svelte 5 (Runes) + TypeScript terminal application.

- **app.css** — Global design system. CSS custom properties for 27 terminal themes. `0px border-radius` enforced. Breakpoints for mobile. `.sr-only` utility.
- **app.html** — HTML shell. Viewport meta, `data-theme` attribute on `<html>`, inline script prevents FOUC.

## Directory map

- `lib/components/` — Svelte UI (Terminal, InputRow, History, OutputRenderer, HintBar)
- `lib/data/` — Static JSON data (`portfolio.json`)
- `lib/model/` — Types (`types.ts`), data helpers (`data.ts`), virtual filesystem (`filesystem.ts`)
- `lib/shell/` — Shell logic: completer, resolver, glob, command handlers
- `lib/stores/` — Svelte writable store for terminal state
- `routes/` — SvelteKit pages (single-page terminal)
- `tests/` — Vitest unit and component tests
