# interpreter/desktop/

Desktop icons view model. Shows top-level non-hidden items from home directory (`~`) as desktop icons.

## Files

### desktop.go

`DesktopVM` — contains `Icons []DesktopIconVM`.

`DesktopIconVM` — `Name`, `Path`, `Icon` (image path like `/icons/folder.png`), `X`, `Y`.

`BuildVM(filesystem)` — reads `~` children, assigns icons (folder.png for dirs, file.png for files), positions in columns. X offset is from right edge (CSS uses `right` not `left`). 6 rows per column, 80px spacing.

`isImageType(mime)` — checks if MIME type is image.

## Relationship

- Called by `main.go:buildVM()` on every state change
- Filesystem comes from `world.World.FS`
- Icons served from `static/icons/`
- Frontend component: `src/lib/components/desktop/DesktopIcons.svelte`
