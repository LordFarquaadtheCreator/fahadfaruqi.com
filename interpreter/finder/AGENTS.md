# interpreter/finder/

Finder window manager, ViewModel builder, and markdown renderer.

## Files

### finder.go

`FinderState` — per-window state: current path, navigation history (back/forward stacks), view mode, selected entry.

`FinderManager` — owns `states map[string]*FinderState`, wraps `wm.Manager` and `fs.FileSystem`.

Methods: `Spawn(path)`, `Navigate(id, path)`, `Back(id)`, `Forward(id)`, `SetViewMode(id, mode)`, `SelectEntry(id, name)`, `Cleanup(id)`.

### viewmodel.go

ViewModel structs:
- `FinderContentVM` — `kind`, `viewMode`, `selectedEntry`, `canBack`, `canForward`, `sidebar`, `breadcrumb`, `content`
- `SidebarSectionVM` — `title`, `items[]`
- `SidebarItemVM` — `name`, `path`, `icon`, `active`
- `BreadcrumbItemVM` — `name`, `path`, `icon`
- `EntryVM` — `name`, `path`, `kind`, `icon`, `size`, `modified`

`BuildContentVM(win, finder, fs)` — builds full finder content for a window:
- `buildSidebar()` — 3 hardcoded sections (Recents, Favorites, Locations) with icons
- `buildBreadcrumb(path)` — splits path into segments with icons
- Content entries from filesystem with icon paths

### render.go

Markdown rendering using `goldmark` + GFM extension + `bluemonday` UGCPolicy sanitization. Images allowed from `/static/`, `/images/`, `http(s)://`.

## API (window.os.*)

- `os.spawn(path)` — open new Finder window at path
- `os.finderNavigate(id, path)` — change cwd, push history
- `os.finderBack(id)` / `os.finderForward(id)` — history navigation
- `os.finderSetView(id, mode)` — switch icon/list/column/gallery
- `os.finderSelect(id, name)` — highlight entry
- `os.finderOpenEntry(id, path)` — double-click: dir = navigate, file = open in viewer

## Sidebar Structure

3 sections:
1. **No title**: Recents, Shared
2. **Favorites**: Applications, Desktop, Documents, Downloads
3. **Locations**: iCloud Drive, Home (farquaad), Network

Each item has `icon` path pointing to `/icons/` images.

## Constraints

- Max 10 windows (FIFO eviction via `wm.Manager`)
- Read-only: no edit, move, delete
- Sidebar items are hardcoded, not derived from filesystem
