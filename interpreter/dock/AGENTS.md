# interpreter/dock/

Dock view model. Shows apps from `apps.DockOrder` with running/active state.

## Files

### dock.go

`DockVM` — `Items []DockItemVM` + `Trash DockItemVM`.

`DockItemVM` — `AppID`, `Name`, `Icon`, `Running`, `Active`.

`BuildVM(manager)` — iterates `apps.DockOrder`, checks which apps have open windows via `wm.Manager.Windows()`, marks focused window's app as `Active`. Trash is hardcoded with `/icons/trash.png`.

## Relationship

- Reads `apps.DockOrder` and `apps.Registry` for app metadata
- Reads `wm.Manager` for running/active state
- Called by `main.go:buildVM()` on every state change
- Frontend component: `src/lib/components/desktop/Dock.svelte`
