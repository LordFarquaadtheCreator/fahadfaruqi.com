# fahadfaruqi.com — Agent Reference

macOS desktop simulation in the browser. SvelteKit 2 + Svelte 5 (Runes) + TypeScript frontend. Go WASM backend owns all logic. Strict MVVM.

## Commands

```bash
npm run dev        # dev server
npm run build      # static output → build/
npm run preview    # serve build/
npm test           # vitest (unit + component)
npm run check      # svelte-check + tsc

# WASM build (run from interpreter/ directory)
GOOS=js GOARCH=wasm go build -o ../static/interp.wasm .
```

## Stack

- **SvelteKit 2** / **Svelte 5** — Runes syntax (`$props`, `$state`, `$derived`, `$effect`)
- **TypeScript** — strict types, no `any`
- **Go WASM** — backend logic, filesystem, shell, window manager, all VMs
- **CSS Custom Properties** — macOS design system, no CSS framework
- **Vitest** + `@testing-library/svelte` — component and unit tests
- **GitHub Actions** — `.github/workflows/deploy.yaml` auto-deploys on push

## Architecture (Strict MVVM)

Go WASM backend (`interpreter/`) owns ALL logic: filesystem, parsing, zsh interpretation, window/Finder/Terminal/Viewer state, drag math, markdown rendering, focus, z-order, navigation history.

Svelte is a dumb view layer:
- Renders ViewModel JSON received from Go via `os.subscribe`
- Forwards raw input events (clicks, key, pointer deltas) to Go via `window.os.*`
- Holds zero domain state. No `$derived` business logic, no parsing, no math
- Only `$state` allowed: ephemeral UI (hover/focus, open menus, clock interval)

Every meaningful operation routes through Go.

## File Map

```
interpreter/              # Go WASM backend (see interpreter/AGENTS.md)
├── main.go               # WASM entry, window.os.* bindings, pushVM()
├── apps/                 # App registry + file associations (see apps/AGENTS.md)
├── command/              # zsh command handlers (see command/AGENTS.md)
├── desktop/              # Desktop icons VM (see desktop/AGENTS.md)
├── dock/                 # Dock VM (see dock/AGENTS.md)
├── finder/               # Finder manager + VM + markdown render (see finder/AGENTS.md)
├── fs/                   # In-memory filesystem (see fs/AGENTS.md)
├── menubar/              # Menu bar VM (see menubar/AGENTS.md)
├── response/             # Shell response types (see response/AGENTS.md)
├── session/              # Shell session state (see session/AGENTS.md)
├── shell/                # Tokenizer, expander, runner (see shell/AGENTS.md)
├── terminal/             # Terminal manager + VM (see terminal/AGENTS.md)
├── viewer/               # Preview app: image + text viewer (see viewer/AGENTS.md)
├── wm/                   # Window manager: position, z-order, focus (see wm/AGENTS.md)
└── world/                # World struct + boot from data.json (see world/AGENTS.md)

src/                      # SvelteKit frontend (see src/agents.md)
├── app.css               # Global CSS: macOS design system
├── app.html              # HTML shell, loads wasm_exec.js
├── routes/
│   ├── +layout.svelte    # Font imports, favicon
│   ├── +layout.ts        # prerender = true
│   └── +page.svelte      # Main desktop page, calls initOS()
├── lib/                  # Frontend library (see lib/AGENTS.md)
│   ├── os/               # VM types + OS bridge (see os/AGENTS.md)
│   ├── components/       # Svelte UI (see components/AGENTS.md)
│   │   ├── desktop/      # Menubar, dock, wallpaper, icons, window layer
│   │   ├── finder/       # Finder app
│   │   ├── terminal/     # Terminal app
│   │   ├── viewer/       # Preview app (image + text/markdown)
│   │   ├── window/       # Generic window chrome
│   │   └── placeholder/  # Splash for unimplemented apps
│   ├── data/             # Static data assets (images)
│   └── wasmLoader.js     # Loads Go WASM module

static/                   # Static assets (see static/AGENTS.md)
├── interp.wasm           # Go WASM binary
├── wasm_exec.js          # Go WASM runtime
├── icons/                # App and file type icons (PNG)
└── wallpapers/           # Desktop wallpapers
```

## Critical Conventions

- **0px border-radius site-wide** — enforced via `* { border-radius: 0 !important }` in `app.css`
- Svelte 5 Runes only. Never legacy `export let` / `$:` syntax
- VM types in `src/lib/os/types.ts` must match Go struct JSON tags exactly
- Rebuild WASM after any Go change
- `mix-blend-mode: difference` on menubar halves for auto-contrast against wallpaper
- Desktop icons are right-aligned (CSS `right` offset, not `left`)

## Deployment

GitHub Actions workflow: `npm ci` → `npm run build` → upload `build/` → deploy to GitHub Pages. No env vars needed.

## Testing

Two Vitest projects in `vite.config.ts`:
- `component` — jsdom, matches `*.svelte.{test,spec}.{js,ts}`
- `server` — node, matches `*.{test,spec}.{js,ts}` excluding svelte files
