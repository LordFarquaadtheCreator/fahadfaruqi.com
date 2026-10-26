# interpreter/finder/

Finder window manager and ViewModel builder. Markdown rendering lives in `interpreter/markdown/`, because
the viewer and Quick Look need it too.

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
- `buildSidebar(path, fs)` — Favorites holds the home folders that exist, Locations holds the home; nothing is listed that would open an empty window
- `buildBreadcrumb(path)` — splits path into segments with icons
- Content entries from filesystem with icon paths; a `.sh` is `kind: script` with the terminal's icon, because it runs rather than opens

## API (window.os.*)

- `os.spawn(path)` — open new Finder window at path
- `os.finderNavigate(id, path)` — change cwd, push history
- `os.finderBack(id)` / `os.finderForward(id)` — history navigation
- `os.finderSetView(id, mode)` — switch icon/list/column/gallery
- `os.finderSelect(id, name)` — highlight entry
- `os.finderOpenEntry(id, path)` — double-click: a folder navigates in place, a file goes through `openPath` in `main.go`, the same path the desktop uses — a script runs in a terminal, an image opens its viewer, a markdown file opens the text viewer

## Constraints

- Max 10 windows (FIFO eviction via `wm.Manager`)
- Read-only: no edit, move, delete
