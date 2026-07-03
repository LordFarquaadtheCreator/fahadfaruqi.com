# This is my Website
<img width="2317" height="1269" alt="Screenshot 2026-07-03 at 3 26 45 PM" src="https://github.com/user-attachments/assets/e9ab6853-4bf8-4e36-b294-4b582adcd087" />

# fahadfaruqi.com

A macOS desktop environment simulated entirely in the browser. Go WASM backend owns all logic — filesystem, shell, window management, state. Svelte renders the view model and forwards events back.

## Stack

- **SvelteKit 2** + **Svelte 5** (Runes) — dumb view layer, renders VM JSON, forwards events
- **Go WASM** — all domain logic: filesystem, zsh shell, window manager, Finder, Terminal, Viewer
- **TypeScript** — strict, VM types mirror Go structs
- **CSS Custom Properties** — macOS design system, no framework
- **Vitest** — unit + component tests

## Architecture

Strict MVVM. Go WASM owns ALL state and logic. Svelte holds zero domain state.

```
User action → window.os.<action>() → Go mutates state → pushVM() → JSON → osStore → Svelte re-renders
```

Every state change triggers full VM re-serialization. No incremental updates. No Svelte stores for app state — only ephemeral UI (hover, focus, clock).

## Project Structure

```
interpreter/              Go WASM backend
├── main.go               Entry point, window.os.* bindings, pushVM()
├── apps/                 App registry + file associations
├── command/              zsh command handlers (ls, cd, cat, open, etc.)
├── desktop/              Desktop icons VM
├── dock/                 Dock VM with running/active state
├── finder/               Finder window manager + VM + markdown render
├── fs/                   In-memory filesystem
├── menubar/              Menu bar VM (menus, tray, clock)
├── response/             Shell response types
├── session/              Shell session (cwd, vars, history)
├── shell/                Tokenizer, expander, runner
├── terminal/             Terminal window manager + VM + tab completion
├── viewer/               Preview app (image + text/markdown viewer)
├── wm/                   Window manager (position, z-order, focus, resize)
└── world/                World struct, boots FS + env from data.json

src/                      SvelteKit frontend
├── app.css               Global macOS design system
├── app.html              HTML shell, loads wasm_exec.js + WASM
├── routes/
│   ├── +layout.svelte    Renders <Desktop />, font imports
│   └── +page.svelte      Head only (title)
└── lib/
    ├── os/               VM types + OS bridge (osStore)
    ├── components/       Svelte UI (desktop, finder, terminal, viewer, window)
    ├── data/             Static data
    └── wasmLoader.js     Loads Go WASM module

static/                   Static assets
├── interp.wasm           Compiled Go WASM binary
├── wasm_exec.js          Go WASM runtime
├── icons/                App + file type icons
└── wallpapers/           Desktop wallpapers
```

## Commands

```bash
npm run dev        # dev server
npm run build      # static output → build/
npm run preview    # serve build/
npm test           # vitest (unit + component)
npm run check      # svelte-check + tsc

# Rebuild WASM after Go changes (from interpreter/)
GOOS=js GOARCH=wasm go build -o ../static/interp.wasm .
```

## Features

- **Finder** — browse in-memory filesystem, sidebar, breadcrumbs, icon/list view, markdown rendering
- **Terminal** — zsh-like shell with command history, tab completion, pipe support
- **Viewer** — image preview + text/markdown viewer
- **Window Manager** — drag, resize, minimize, maximize, focus z-order
- **Dock** — app launcher with running indicators, tooltips
- **Menu Bar** — Apple menu, app menus, tray icons, live clock, mix-blend-mode auto-contrast
- **Desktop Icons** — right-aligned, double-click to open

## Deployment

GitHub Actions: setup Go + Node → build WASM from `interpreter/` → `npm ci` → `npm run build` → deploy `build/` to GitHub Pages. Triggers on push to `main`. CI builds WASM from source — no need to commit `static/interp.wasm`.

## Conventions

- **0px border-radius site-wide** — enforced via `* { border-radius: 0 !important }`
- Svelte 5 Runes only — never legacy `export let` / `$:`
- VM types in `src/lib/os/types.ts` must match Go struct JSON tags exactly
- Rebuild WASM after any Go change
- Desktop icons right-aligned (CSS `right` offset, not `left`)
- `mix-blend-mode: difference` on menubar for auto-contrast against wallpaper

## License

MIT

