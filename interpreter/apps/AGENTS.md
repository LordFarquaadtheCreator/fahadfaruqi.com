# interpreter/apps/

App registry and file association system. Central source of truth for what apps exist, their metadata, and which files open in which app.

## Files

### registry.go

Defines `AppType` constants: `TypeFinder`, `TypeTerminal`, `TypeViewer`, `TypePlaceholder`, `TypeEditor`.

`App` struct: `ID`, `Name`, `Icon` (path like `/icons/finder.png`), `Type`, `Width`, `Height`.

`Registry` map: all registered apps keyed by ID. Both `viewer-image` and `viewer-text` are named "Preview" with `preview.png` icon — they share `TypeViewer`.

`DockOrder` slice: controls which apps appear in dock and their order. Trash is added separately in `dock/dock.go`, not here.

`Get(id)` — lookup helper.

### associations.go

`Extensions` map: file extension → app ID. Currently:
- `.png`, `.jpeg`, `.jpg`, `.gif`, `.webp` → `viewer-image`
- `.md` → `viewer-text`
- `.txt` → `viewer-text`

`AppForFile(filename)` — resolves app ID from extension. Returns empty string if no association.

## Adding a New App

1. Add entry to `Registry` map with unique ID, name, icon path, type, default window size
2. Add ID to `DockOrder` slice (controls dock position)
3. Place icon PNG in `static/icons/<name>.png`
4. If app opens files by extension, add mapping to `Extensions`
5. Rebuild WASM

## Relationship

- `dock/` reads `DockOrder` to build dock VM
- `menubar/` reads `Registry` to resolve active app name from focused window
- `viewer/` uses `AppForFile()` to determine viewer type
- `main.go` uses `Get()` to resolve app metadata on spawn
