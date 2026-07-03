# src/lib/components/finder/

Finder app window content. Renders sidebar, toolbar, breadcrumb, and file entries.

## Files

### FinderApp.svelte

Renders Finder window content from `FinderContentVM`:
- **Sidebar**: 3 sections (Recents, Favorites, Locations) with icons. Click forwards to `os()?.finderNavigate(id, path)`.
- **Toolbar**: Back/forward arrows (disabled state from `canBack`/`canForward`), view mode buttons, search icon. Click forwards to `os()?.finderBack(id)` / `os()?.finderForward(id)`.
- **Breadcrumb**: Path segments with icons. Click forwards to `os()?.finderNavigate(id, path)`.
- **Content area**: File entries with icons, names, sizes, modified dates. Double-click forwards to `os()?.finderOpenEntry(id, path)`.

Props: `window: WindowVM`, `content: FinderContentVM`.

## Relationship

- Rendered by `WindowLayer.svelte` when `win.content.kind === 'finder'`
- All navigation routed through `window.os.*` — no local path state
- CSS in `src/app.css` under `/* ── Finder ── */` section
- Sidebar items, breadcrumbs, and entry icons all come from Go VM (icon paths like `/icons/folder.png`)
