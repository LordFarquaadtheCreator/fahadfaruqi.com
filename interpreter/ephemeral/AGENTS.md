# interpreter/ephemeral/

Quick Look — the preview that opens on a Finder selection when spacebar is pressed, and
closes on Escape, on the next spacebar, or on a click away. One preview at a time.

## Files

### ephemeral.go

`EphemeralState` — per-window: `FilePath`, `FileName`, `ImagePath` (the URL to render),
`IsDir`.

`EphemeralManager` — owns `states map[string]*EphemeralState`, wraps `wm.Manager` and
`fs.FileSystem`, and tracks `currentID`, the single open preview.

`SpawnForFile(filePath)` — closes the open preview first, then picks the window size:
- images: the same fit math as the viewer, on the file's natural dimensions
- folders: the app's default size

An image renders `Meta.CompressedPath` when the file has one — the master stays for the
viewer, which is where a deliberate open goes.

`Close()`, `CurrentID()`, `CurrentFilePath()`, `Cleanup(id)`, `GetState(id)`.

`BuildContentVM(win, em)` — builds `EphemeralContentVM`:
- Image: `kind: "ephemeral-image"`, `src`, `name`, `filePath`
- Folder: `kind: "ephemeral-folder"`, `icon`, `name`, `filePath`

`IsImageFile(filePath)` — extension check for image files.

## Relationship

- Driven by `main.go`: `window.os.ephemeralPreview(filePath)`, `ephemeralClose()`, and
  `ephemeralOpenPreview()` — the last one closes the preview and opens the file in the
  viewer
- `buildContent()` handles `apps.TypeEphemeral` by calling `BuildContentVM`
- Frontend component: `src/lib/components/viewer/EphemeralPreview.svelte`, rendered inside
  the generic window chrome; the "Open in Preview" button forwards to
  `os()?.ephemeralOpenPreview?.()`
- The spacebar handler lives in `src/routes/+page.svelte`, alongside the Escape handler
