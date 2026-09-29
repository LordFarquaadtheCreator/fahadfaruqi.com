# fahadfaruqi.com

A macOS desktop environment simulated entirely in the browser. Go WASM backend owns all logic — filesystem, shell, window management, state. Svelte renders the view model and forwards events back.

<img width="2317" height="1269" alt="Screenshot 2026-07-03 at 3 26 45 PM" src="https://github.com/user-attachments/assets/e9ab6853-4bf8-4e36-b294-4b582adcd087" />

## Stack

- **SvelteKit 2** + **Svelte 5** (Runes) — dumb view layer, renders VM JSON, forwards events
- **Go WASM** — all domain logic: filesystem, zsh shell, window manager, Finder, Terminal, Viewer, Quick Look
- **TypeScript** — strict, VM types mirror Go structs
- **CSS Custom Properties** — macOS design system, no framework
- **Vitest** — unit + component tests

## Architecture

Strict MVVM. Go WASM owns ALL state and logic. Svelte holds zero domain state.

```
User action → window.os.<action>() → Go mutates state → pushVM() → JSON → osStore → Svelte re-renders
```

Every state change triggers full VM re-serialization. No incremental updates. No Svelte stores for app state — only ephemeral UI (hover, focus, clock).

## Assets

No image, icon, wallpaper or wasm binary lives in this repo or in the build. They are served from Cloudflare — the R2 bucket `assets`, custom domain `assets.fahadfaruqi.com`, under the `www/` prefix — which keeps the Pages artifact around 100 KB instead of ~95 MB.

- **Origin**: `https://assets.fahadfaruqi.com/www`. Three places know it, all defaulting to the same value: `src/lib/assets.js` (Svelte; `VITE_ASSETS_BASE` overrides it), `src/app.html` (the pre-hydration wasm load), and the Go wasm (baked in by `build-wasm.sh` via `-ldflags "-X interpreter/assets.Base=..."`, `ASSETS_BASE` overrides it).
- **Go side**: `interpreter/world/data.json` and the VM builders keep relative paths (`/images/pictures/…`, `/icons/…`, `/wallpapers/…`) and `interpreter/assets.URL()` turns them into absolute URLs as they enter a view model.
- **R2 layout**: `www/images/pictures/…` (masters, with `thumbs/` and `compressed/` beside them), `www/images/companies/…`, `www/icons/…`, `www/wallpapers/…` (masters + `compressed/`), `www/interp.wasm`, `www/wasm_exec.js`. The bucket is shared with art.fahadfaruqi.com, whose Worker ignores anything under `www/` the same way it ignores its own `d/` derivatives.
- **Compressed copies**: every picture and wallpaper also exists as an AVIF copy — `compressed/<stem>.avif`, longest side 1600 px for pictures and 3200 px for wallpapers, encoded with `libsvtav1 -crf 22 -preset 6`. The Finder renders the 100 px thumbnail, Quick Look and the Finder's click prefetch render the compressed copy, and the Viewer opens the master.
- **Cache**: images `max-age=604800` (7 days) — masters, thumbnails and compressed copies alike; `interp.wasm` `max-age=300` — Cloudflare does not edge-cache `application/wasm`, so every visit fetches it once; `wasm_exec.js` `max-age=14400`, because Cloudflare raises browser TTLs below the zone's 4-hour Browser Cache TTL. Replacing an image *in place* therefore keeps serving the old bytes for up to 7 days — rename the file, or purge it in the Cloudflare dashboard.

## Project Structure

```
interpreter/              Go WASM backend
├── main.go               Entry point, window.os.* bindings, pushVM()
├── apps/                 App registry + file associations
├── assets/               CDN base injected at build time; assets.URL()
├── command/              zsh command handlers (ls, cd, cat, open, etc.)
├── desktop/              Desktop icons VM
├── dock/                 Dock VM with running/active state
├── ephemeral/            Quick Look preview manager (spacebar toggle)
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
├── app.html              HTML shell, loads wasm_exec.js + WASM from the CDN
├── routes/
│   ├── +layout.svelte    Renders <Desktop />, font imports
│   └── +page.svelte      Head only (title)
└── lib/
    ├── assets.js         CDN base + assetUrl()
    ├── os/               VM types + OS bridge (osStore)
    ├── components/       Svelte UI (desktop, finder, terminal, viewer, window)
    ├── data/             Static data
    └── wasmLoader.js     Loads Go WASM module

static/                   Site plumbing only: favicon.png, robots.txt, AGENTS.md
scripts/                  publish-wasm.py, sync-pictures.py, sync-wallpapers.py (all need R2 credentials)
```

## Commands

```bash
bun run dev          # dev server (rebuilds WASM via predev hook)
bun run dev:light    # dev server, force light theme
bun run dev:dark     # dev server, force dark theme
bun run build        # static output → build/ (rebuilds WASM via prebuild)
bun run preview      # serve build/
bun run test         # vitest (unit + component)
bun run check        # svelte-check + tsc

# Rebuild the wasm locally (also done by the predev/prebuild hooks)
./build-wasm.sh

# Publish the interpreter + its wasm_exec.js to the CDN (CI does this on every push)
R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... ./scripts/publish-wasm.py

# Upload pictures from a local directory and sync them into data.json
R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... ./scripts/sync-pictures.py [source-dir]

# Re-encode the desktop wallpapers from the masters already on the CDN
R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... ./scripts/sync-wallpapers.py
```

`sync-pictures.py` writes three objects per picture (master, thumbnail, compressed copy) and
needs ffmpeg on PATH for the AVIF encode; `sync-wallpapers.py` reads the masters back from
the bucket and writes only the compressed copy.

## Features

- **Finder** — browse in-memory filesystem, sidebar, breadcrumbs, icon/list view, markdown rendering
- **Terminal** — zsh-like shell with command history, tab completion, pipe support
- **Viewer** — image preview + text/markdown viewer
- **Quick Look** — ephemeral preview via spacebar, click-away close, folder support
- **Window Manager** — drag, resize, minimize, maximize, focus z-order
- **Dock** — app launcher with running indicators, tooltips
- **Menu Bar** — Apple menu, app menus, tray icons, live clock, mix-blend-mode auto-contrast
- **Desktop Icons** — right-aligned, double-click to open
- **Theming** — light/dark wallpaper driven from WASM via `os.setTheme`

## Virtual Filesystem

Boots from `interpreter/world/data.json`. Home (`~`) contains:
`About/`, `Contact/`, `Blog/`, `Resume/`, `Pictures/` — each with markdown or image content. Markdown files render to HTML in Viewer.

## Deployment

GitHub Actions, on push to `main`: setup Bun + Go (`go-version-file: interpreter/go.mod`, because the published `wasm_exec.js` must come from the toolchain that built the wasm) → `bun install` → `bun run build` (its `prebuild` hook compiles the wasm with `ASSETS_BASE`, the job-level env) → `scripts/publish-wasm.py` uploads `interp.wasm` + `wasm_exec.js` to the CDN using the `R2_ACCESS_KEY_ID` / `R2_SECRET_ACCESS_KEY` repo secrets → deploy `build/` to GitHub Pages.

## Conventions

- **Bun, not npm** — `bun install`, `bun run <script>`, `bunx <package>`. No npm, npx, yarn or pnpm.
- **0px border-radius site-wide** — enforced via `* { border-radius: 0 !important }`
- Svelte 5 Runes only — never legacy `export let` / `$:`
- VM types in `src/lib/os/types.ts` must match Go struct JSON tags exactly
- Rebuild WASM after any Go change, and publish it — the deployed site loads the wasm from the CDN, not from `build/`
- Never add images, icons or wallpapers under `static/` — they go to the CDN (see Assets)
- Desktop icons right-aligned (CSS `right` offset, not `left`)
- `mix-blend-mode: difference` on menubar for auto-contrast against wallpaper

## License

MIT
