# src/lib/components/finder/

macOS Finder UI. Pure view layer.

## Components

- **Desktop.svelte** — Fixed-position host. Mounts `FinderWindow` per VM entry.
- **FinderWindow.svelte** — Chrome (title bar, traffic lights), drag handler.
- **FinderToolbar.svelte** — Back/forward, view-mode switcher, breadcrumb.
- **FinderSidebar.svelte** — Renders `win.sidebar`. Click → `os.navigate`.
- **content/** — One component per VM `content.kind`.
  - `IconGrid.svelte` — icon grid
  - `ListTable.svelte` — list rows
  - `ColumnView.svelte` — column browser
  - `GalleryView.svelte` — gallery preview
  - `MarkdownPane.svelte` — `{@html c.html}`

## Rules

- Every component renders props received from `finderState.vm`.
- Every interaction forwards events to `window.os.*`.
- Zero business logic, zero `$derived` view model computation.
- Rounded corners on `.finder-window` override global `border-radius: 0`.
