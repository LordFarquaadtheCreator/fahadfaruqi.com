# interpreter/desktop/

Desktop icons view model. Shows top-level non-hidden items from home directory (`~`) as desktop icons.

## Files

### desktop.go

`DesktopVM` — contains `Icons []DesktopIconVM` and `Wallpaper` (the compressed wallpaper for the current theme).

`DesktopIconVM` — `Name`, `Path`, `Icon` (image path like `/icons/folder.png`), `X`, `Y`.

`BuildVM(filesystem, dark)` — reads `~` children, assigns icons (folder.png for dirs, file.png for files), positions in columns. X offset is from right edge (CSS uses `right` not `left`). 6 rows per column, 80px spacing. `dark` picks between `/wallpapers/compressed/wallpaper.avif` and `wallpaper-dark.avif`.

`isImageType(mime)` — checks if MIME type is image.

## Relationship

- Called by `main.go:buildVM()` on every state change
- Filesystem comes from `world.World.FS`
- Icons and wallpapers served from the CDN (`www/icons/`, `www/wallpapers/`)
- Frontend component: `src/lib/components/desktop/DesktopIcons.svelte`
