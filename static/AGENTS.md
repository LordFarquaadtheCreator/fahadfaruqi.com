# static/

Only site plumbing is served from here: `favicon.png`, `robots.txt`. Everything that
used to live here — `icons/`, `wallpapers/`, `images/`, `interp.wasm`, `wasm_exec.js` —
now comes from Cloudflare.

## Where assets actually are

The R2 bucket `assets`, custom domain `assets.fahadfaruqi.com`, prefix `www/`:

```
https://assets.fahadfaruqi.com/www/icons/finder.png
https://assets.fahadfaruqi.com/www/wallpapers/wallpaper.jpg
https://assets.fahadfaruqi.com/www/images/pictures/lucky_1.JPG
https://assets.fahadfaruqi.com/www/images/pictures/thumbs/lucky_1.JPG
https://assets.fahadfaruqi.com/www/images/pictures/compressed/lucky_1.avif
https://assets.fahadfaruqi.com/www/wallpapers/compressed/wallpaper.avif
https://assets.fahadfaruqi.com/www/interp.wasm
https://assets.fahadfaruqi.com/www/wasm_exec.js
```

The bucket is shared with art.fahadfaruqi.com. Its metadata Worker lists every object
that is not under `d/` (its own derivatives) or `www/` (this site), so keep new objects
under `www/` or they turn up as photographs in that gallery.

## Who builds the URLs

- Go VM: paths stay relative (`/icons/folder.png`, `/wallpapers/compressed/wallpaper.avif`) and
  `interpreter/assets.URL()` prefixes them on the way into a view model. The base is
  baked in by `build-wasm.sh` (`-ldflags "-X interpreter/assets.Base=…"`), overridable
  with `ASSETS_BASE`.
- Svelte: `src/lib/assets.js` (overridable with `VITE_ASSETS_BASE`) and `src/app.html`,
  which needs the base for the pre-hydration wasm load.

## Publishing

```bash
# interpreter + the wasm_exec.js of the toolchain that built it (CI does this on push)
R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... ./scripts/publish-wasm.py

# pictures, from a local source directory, uploaded and synced into data.json:
# master, 100px thumbnail, and the compressed copy the simulation renders
R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... ./scripts/sync-pictures.py [source-dir]

# the desktop wallpapers, re-encoded from the masters already in the bucket
R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... ./scripts/sync-wallpapers.py
```

Adding an app icon means: put the PNG in the bucket at `www/icons/<name>.png`, then
reference `assets.URL("/icons/<name>.png")` in `apps/registry.go`. The filename must
match exactly — nothing validates it.
