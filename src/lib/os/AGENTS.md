# src/lib/os/

TypeScript VM types and OS bridge. Single source of truth for frontend rendering.

## Files

### types.ts

TypeScript interfaces mirroring Go VM structs in `interpreter/`. All fields match JSON tags exactly.

Key types:
- `RootVM` — top-level VM with windows, dock, menuBar, desktop
- `WindowVM` — window with `content` union (Finder | Terminal | Viewer | Splash)
- `FinderContentVM` — sidebar sections, breadcrumb, entries
- `TerminalContentVM` — cwd, history
- `ViewerContentVM` — `kind: 'viewer-image' | 'viewer-text'`, src, content, isMarkdown
- `DockVM`, `MenuBarVM`, `DesktopVM` — chrome VMs

### osStore.svelte.ts

VM subscription store. Not a Svelte store — plain module-level state.

- `initOS()` — called on app mount, hooks into `window.os.subscribe`, parses JSON, notifies subscribers
- `subscribe(fn)` — register callback for VM updates, returns unsubscribe
- `getVM()` — get current VM snapshot
- `setViewport(w, h)` — forward to `window.os.setViewport`
- `os()` — typed accessor for `window.os.*` API

## Relationship

- `types.ts` must stay in sync with Go struct JSON tags
- `osStore` is the only place that talks to `window.os`
- Components import types from here, call `os()` for event forwarding
- `+page.svelte` calls `initOS()` on mount
