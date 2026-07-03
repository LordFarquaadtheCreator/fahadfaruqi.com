# src/

SvelteKit 2 + Svelte 5 (Runes) + TypeScript. macOS desktop simulation. Strict MVVM — Go WASM backend owns all logic, Svelte is dumb view layer.

## Architecture

This is a macOS desktop environment simulated in the browser. Go WASM backend (`interpreter/`) owns all state and logic. Svelte renders VM JSON and forwards events back.

```
Go WASM (interpreter/) → RootVM JSON → osStore → Svelte components → user events → window.os.* → Go
```

## Directory Map

- `app.css` — Global design system. macOS CSS variables, window/dock/menubar/finder/viewer styles. `0px border-radius` enforced.
- `app.html` — HTML shell, loads wasm_exec.js
- `lib/` — Frontend library (see `lib/AGENTS.md`)
  - `os/` — VM types + OS bridge to Go WASM
  - `components/` — Svelte UI components (see `components/AGENTS.md`)
    - `desktop/` — Menubar, dock, wallpaper, icons, window layer
    - `finder/` — Finder app
    - `terminal/` — Terminal app
    - `viewer/` — Preview app (image + text/markdown)
    - `window/` — Generic window chrome
    - `placeholder/` — Splash for unimplemented apps
  - `data/` — Static data assets (images)
  - `wasmLoader.js` — Loads Go WASM module
- `routes/` — SvelteKit pages
  - `+page.svelte` — Main desktop page, calls `initOS()` on mount
  - `+layout.svelte` — Font imports, favicon, theme-color meta

## CSS Sections in app.css

1. **Reset & Base** — box-sizing, font vars, scrollbar styling
2. **Desktop** — wallpaper, desktop icons
3. **MenuBar** — top bar, menu items, tray icons, mix-blend-mode
4. **Dock** — bottom dock, tooltips, running indicators
5. **Window** — window chrome, titlebar, traffic lights, resize
6. **Finder** — sidebar, toolbar, breadcrumb, entries
7. **Terminal** — terminal output, input prompt
8. **Viewer** — image viewer, text viewer toolbar
9. **Markdown Body** — rendered markdown styles (headings, code, tables, etc.)
10. **Splash** — placeholder app splash screen

## Critical Rules

- Svelte 5 Runes only (`$props`, `$state`, `$effect`). Never `export let` or `$:`.
- No business logic in Svelte. All state from Go VM, all events to Go via `window.os.*`.
- `$state` only for ephemeral UI (hover, focus, open menus, clock interval).
- VM types in `lib/os/types.ts` must match Go struct JSON tags exactly.
- Rebuild WASM after Go changes: `cd interpreter && GOOS=js GOARCH=wasm go build -o ../static/interp.wasm .`
