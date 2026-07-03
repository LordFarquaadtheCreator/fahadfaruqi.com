# interpreter/

Go WASM module. Owns ALL domain logic per strict MVVM rule. Svelte is dumb view layer — renders VM JSON, forwards events back via `window.os.*`.

## Build

```bash
GOOS=js GOARCH=wasm go build -o ../static/interp.wasm .
```

Run from `interpreter/` directory. Output goes to `static/interp.wasm`.

## Architecture

`main.go` is WASM entry point. On load:
1. `world.Load()` boots filesystem from `world/data.json`, creates window manager, finder, viewer, terminal managers
2. Exposes `window.os.*` bindings (spawn, close, focus, drag, resize, navigate, etc.)
3. Exposes `interpRun(input)` for shell command execution
4. `pushVM()` serializes entire `RootVM` JSON and pushes to all JS subscribers via `os.subscribe` callbacks

Every state change calls `pushVM()` → full VM re-serialization → Svelte re-renders. No incremental updates.

## Module Map

| Module | Purpose | Key Files |
|---|---|---|
| `apps/` | App registry, file associations | `registry.go`, `associations.go` |
| `command/` | zsh command handlers | `registry.go`, `ls.go`, `cd.go`, `cat.go`, `open.go`, etc. |
| `desktop/` | Desktop icons VM | `desktop.go` |
| `dock/` | Dock VM with running/active state | `dock.go` |
| `finder/` | Finder window manager + VM + markdown render | `finder.go`, `viewmodel.go`, `render.go` |
| `fs/` | In-memory filesystem | `fs.go`, `node.go`, `resolve.go` |
| `menubar/` | Menu bar VM (menus, tray, clock) | `menubar.go` |
| `response/` | JSON response types for shell commands | `response.go` |
| `session/` | Shell session (cwd, vars, history) | `session.go` |
| `shell/` | Tokenizer, variable expander, command runner | `run.go`, `tokenize.go`, `expand.go` |
| `terminal/` | Terminal window manager + VM | `terminal.go` |
| `viewer/` | Preview app (image + text/markdown viewer) | `viewer.go` |
| `wm/` | Generic window manager (position, z-order, focus) | `manager.go`, `window.go`, `viewmodel.go` |
| `world/` | World struct + boots everything from data.json | `world.go`, `load.go`, `data.json` |

## Data Flow

```
User action (click/keypress)
  → window.os.<action>(args)    [JS → Go WASM]
  → Manager mutates state
  → pushVM()                     [Go serializes RootVM]
  → os.subscribe callback        [Go → JS]
  → Svelte re-renders            [JS reads VM, no logic]
```

## RootVM Schema

Built in `main.go:buildVM()`. Contains:
- `windows[]` — all open windows with content VMs
- `focusedId` — topmost window ID
- `dock` — dock items with running/active flags
- `menuBar` — menus, tray icons, clock, active app name
- `desktop` — desktop icons (right-aligned)

## Content Dispatch

`buildContent(win)` in `main.go` dispatches by `win.AppType`:
- `"finder"` → `finder.BuildContentVM`
- `"terminal"` → `terminal.BuildContentVM`
- `"viewer"` → `viewer.BuildContentVM`
- default → splash screen

**Gotcha**: `AppType` is the generic type (`"viewer"`), not the specific ID (`"viewer-image"`). Don't check for `"viewer-image"` in the switch — use `apps.TypeViewer`.

## Key Rule

Svelte never reads FS directly. All state flows through `window.os.subscribe`. No Svelte stores for app state. No `$derived` business logic in components.
