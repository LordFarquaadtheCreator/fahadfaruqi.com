# src/lib/components/desktop/

Desktop chrome components — menubar, dock, wallpaper, desktop icons, window layer.

## Files

### Desktop.svelte

Top-level desktop container. Renders Wallpaper, MenuBar, DesktopIcons, WindowLayer, Dock. Subscribes to VM via `osStore.subscribe()`.

### Wallpaper.svelte

Full-screen fixed background image. Default: `/wallpapers/wallpaper.jpg`. Accepts `src` prop override.

### MenuBar.svelte

Top menu bar. Renders:
- Left: Apple logo (SVG), active app name (bolder via `.app-name` class), menu items
- Right: tray icons (inline SVGs for search, controls, wifi, battery), live clock

Uses `mix-blend-mode: difference` on left/right halves for auto-contrast against wallpaper. Clock updates via `setInterval` every second (local `$state`, not from Go VM).

Tray icon SVGs are keyed by `item.icon` string: `search`, `controls`, `wifi`, `battery-full`.

### DesktopIcons.svelte

Right-aligned desktop icons. Positioned via `style="right: {icon.x}px; top: {icon.y}px;"`. Double-click forwards to `os()?.openEntry(icon.path)`.

### Dock.svelte

Bottom dock. Renders app icons from `dock.items`, trash icon. Shows running indicator dot and active highlight. Tooltips on hover. Click forwards to `os()?.launchApp(appId)`.

### WindowLayer.svelte

Renders all open windows. Dispatches to app-specific component based on `win.content.kind`. Passes window VM and event callbacks to `Window.svelte`.

## Relationship

- All read from `RootVM` via `osStore` subscription
- Event forwarding via `os()` accessor
- CSS in `src/app.css` under `/* ── Desktop ── */` and `/* ── MenuBar ── */` sections
