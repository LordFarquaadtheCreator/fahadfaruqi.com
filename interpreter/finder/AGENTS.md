# interpreter/finder/

Window manager and ViewModel builder for macOS-style Finder.

## API (window.os.*)

- `os.subscribe(cb, w, h)` — push VM JSON on every state change
- `os.setViewport(w, h)` — update bounds clamping
- `os.spawn(path)` — open new Finder window at path
- `os.close(id)` / `os.closeFocused()` — remove window
- `os.focus(id)` — bring to front
- `os.drag(id, dx, dy)` — move window (Go clamps bounds)
- `os.navigate(id, path)` — change cwd, push history
- `os.back(id)` / `os.forward(id)` — history navigation
- `os.setViewMode(id, mode)` — switch icon/list/column/gallery
- `os.selectEntry(id, name)` — highlight entry in folder
- `os.openEntry(id, path)` — double-click handler (dir = navigate, file = preview)

## ViewModel Schema

See `viewmodel.go` for full `VM`, `WindowVM`, `ContentVM` structs.

## Markdown Rendering

- `goldmark` + GFM extension for parsing
- `bluemonday` UGCPolicy for sanitization
- Images allowed: `/static/...`, `/images/...`, `http(s)://...`
- Relative/local image paths stripped

## Constraints

- Max 10 windows (FIFO eviction)
- Fixed-size windows (no resize)
- All files are `.md` only
- Read-only: no edit, move, delete
