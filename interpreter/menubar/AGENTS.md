# interpreter/menubar/

Menu bar view model. Top bar with Apple menu, app menus, tray icons, and clock.

## Files

### menubar.go

`MenuBarVM` — `ActiveApp`, `Menus[]`, `Clock`, `Tray[]`.

`MenuVM` — `Name`, `Items[]`. First menu has empty name (Apple logo).

`MenuItemVM` — `Label`, `Shortcut`, `Separator`.

`TrayItemVM` — `Icon` (string key), `Name` (tooltip).

`BuildVM(manager)` — resolves active app name from focused window, builds static menu structure, formats clock, defines tray items.

## Tray Icons

Tray items use string keys that map to inline SVGs in `MenuBar.svelte`:
- `search` — magnifying glass
- `controls` — Control Center (two toggle switches)
- `wifi` — full signal Wi-Fi arcs
- `battery-full` — full battery

Clock format from Go: `"Mon Jan 2 3:04 PM"`. Frontend overrides with live `setInterval` updating every second.

## Menu Structure

Static menus: Apple (empty name), active app name, File, Edit, View, Go, Window, Help. All items are display-only — no dropdown functionality yet.

## Relationship

- Reads `wm.Manager` for focused window → active app name
- Reads `apps.Registry` to resolve app display name
- Called by `main.go:buildVM()` on every state change
- Frontend component: `src/lib/components/desktop/MenuBar.svelte`
- CSS: `mix-blend-mode: difference` on left/right halves for auto contrast against wallpaper
