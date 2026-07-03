# interpreter/viewer/

Preview app — handles image and text/markdown file viewing. One viewer window per file.

## Files

### viewer.go

`ViewerState` — per-window: `FilePath`, `FileName`, `MimeType`, `Content`, `ImagePath`.

`ViewerManager` — owns `states map[string]*ViewerState`, wraps `wm.Manager` and `fs.FileSystem`.

`SpawnForFile(filePath)` — resolves app ID via `apps.AppForFile()`, reads file content from filesystem, determines MIME type by extension:
- `.png`, `.jpg`, `.jpeg`, `.gif`, `.webp` → image viewer
- `.md` → markdown (content stored as raw text, rendered to HTML by finder's `render.go`)
- `.txt` and everything else → plain text

`BuildContentVM(win, vm)` — builds `ViewerContentVM`:
- Image: `kind: "viewer-image"`, `src`, `name`
- Text: `kind: "viewer-text"`, `name`, `content`, `isMarkdown`

## Gotcha

`AppType` for viewer windows is `"viewer"` (generic `TypeViewer`), NOT `"viewer-image"` or `"viewer-text"`. The specific kind is in the `ViewerContentVM.Kind` field, not `win.AppType`. The `buildContent` dispatch in `main.go` checks `apps.TypeViewer`, not the specific IDs.

## Relationship

- Called by `main.go:buildContent()` when `win.AppType == "viewer"`
- Spawned by `command/open.go`, `main.go:openFile()`, and `main.go:finderOpenEntry()`
- Both `viewer-image` and `viewer-text` share `TypeViewer` and `preview.png` icon
- Frontend components: `src/lib/components/viewer/ImageViewer.svelte`, `src/lib/components/viewer/TextViewer.svelte`
- Markdown content is raw text from filesystem — HTML rendering happens in `finder/render.go` via `cat` command, not here
