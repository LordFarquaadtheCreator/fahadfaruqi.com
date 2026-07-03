# static/

Static assets served directly by SvelteKit. No build step — files served as-is.

## Structure

```
static/
├── interp.wasm        # Go WASM binary (built from interpreter/)
├── interp.v2.wasm     # Older WASM binary (legacy, can be cleaned up)
├── wasm_exec.js       # Go WASM runtime support (from Go SDK)
├── favicon.png        # Site favicon
├── robots.txt         # Robots
├── icons/             # App and file type icons (PNG)
│   ├── finder.png, iterm.png, preview.png, textedit.png
│   ├── obsidian.png, slack.png, spotify.png, steam.png
│   ├── devin.png, zed.png, xcode.png, figma.png, hermes.png
│   ├── affinity.png, whatsapp.png, imsg.png
│   ├── folder.png, file.png, trash.png
│   └── ... (see apps/registry.go for full list)
├── wallpapers/        # Desktop wallpaper images
│   ├── wallpaper.jpg  # Active wallpaper (referenced by Wallpaper.svelte)
│   └── ... (other wallpapers, not currently used)
└── images/            # General images (currently empty)
```

## WASM Build

```bash
cd interpreter/
GOOS=js GOARCH=wasm go build -o ../static/interp.wasm .
```

`wasm_exec.js` comes from Go SDK: `$(go env GOROOT)/misc/wasm/wasm_exec.js`. Copy here if Go version changes.

## Icon Naming

Icons referenced by path in Go `apps/registry.go` (e.g. `/icons/finder.png`). Must match exactly. Adding a new app = add icon PNG here + reference in registry.

## Relationship

- `wasmLoader.js` (in `src/lib/`) fetches `interp.wasm` and `wasm_exec.js`
- `Wallpaper.svelte` references `/wallpapers/wallpaper.jpg`
- Go VM sends icon paths as strings (e.g. `/icons/folder.png`) — Svelte renders `<img src={path}>`
