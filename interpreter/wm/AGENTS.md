# interpreter/wm/

Generic window manager. Owns window position, size, z-order, focus, minimize/maximize. App-specific state lives in separate packages (`finder/`, `terminal/`, `viewer/`).

## Files

### window.go

`Window` struct: `ID`, `AppType` (generic type string), `AppID` (specific app ID), `Title`, `X`, `Y`, `Width`, `Height`, `Z`, `Minimized`, `Maximized`.

`AppType` values: `"finder"`, `"terminal"`, `"viewer"`, `"placeholder"`, `"editor"`.

### manager.go

`Manager` struct: `windows map[string]*Window`, `order []string` (z-order, last = focused), `viewportW/H`, `notify func()`.

Methods: `Spawn(appType, appID, title, w, h)`, `Close(id)`, `CloseFocused()`, `Focus(id)`, `Drag(id, dx, dy)`, `Resize(id, w, h)`, `Minimize(id)`, `Maximize(id)`, `Restore(id)`, `Get(id)`, `FocusedID()`, `Order()`, `Windows()`, `Viewport()`, `AssignZ()`.

- Max 10 windows (FIFO eviction when exceeded)
- `clampWindow()` keeps at least 40px visible within viewport
- `SetNotify(n)` sets callback fired on every state change → triggers `pushVM()` in `main.go`

### viewmodel.go

`WindowVM` — JSON-serializable window for frontend. Built by `BuildVM(manager)` which iterates windows in z-order, assigns Z values, marshals to JSON.

## Relationship

- Created by `world/load.go`
- Used by all app packages (`finder/`, `terminal/`, `viewer/`) to spawn windows
- `dock/` and `menubar/` read focused/running state from it
- `main.go` sets notify callback → `pushVM()` on every change
- Frontend: `src/lib/components/window/Window.svelte` renders each window
