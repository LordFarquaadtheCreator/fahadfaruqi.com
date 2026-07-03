# interpreter/world/

World struct and boot sequence. Ties all modules together.

## Files

### world.go

`World` struct — holds references to everything:
- `FS` — `fs.FileSystem`
- `Env` — `map[string]string` (HOME, USER, SHELL)
- `WM` — `*wm.Manager`
- `Finder` — `*finder.FinderManager`
- `Terminal` — `*terminal.TerminalManager`
- `Viewer` — `*viewer.ViewerManager`

Note: `Dock`, `MenuBar`, `Desktop` fields exist but are built on-demand in `main.go:buildVM()`, not stored persistently.

### load.go

`Load()` — boots everything from embedded `data.json`:
1. Unmarshals `data.json` into `map[string]jsonNode`
2. Creates `fs.Node` map with content, meta, hidden flag
3. Validates: no orphan children, no orphan nodes, hidden flag consistency
4. Creates `fs.FileSystem`, `wm.Manager`, `finder.FinderManager`, `viewer.ViewerManager`
5. Returns `World` struct

`data.json` is embedded via `//go:embed data.json`.

### data.json

Virtual filesystem definition. All files and directories. Paths use `~` as root. See `fs/AGENTS.md` for format.

## Relationship

- `main.go` calls `world.Load()` on WASM init
- `World` implements `terminal.CommandRunner` interface (delegates to `shell.Run()`)
- `World` is passed to every `command/` handler
- Validation panics on boot — bad `data.json` crashes WASM load
