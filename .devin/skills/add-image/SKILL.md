---
name: add-image
description: when the user needs to upload a folder to Pictures/
---

# Sync Pictures

Uploads a folder of images to the CDN and syncs them into `interpreter/world/data.json` as
`~/Pictures/<filename>` entries.

## When to use

- After adding or removing pictures from the source folder
- Before building WASM if pictures changed

## Steps

1. Run the sync script (needs `R2_ACCESS_KEY_ID` / `R2_SECRET_ACCESS_KEY`):

```bash
./scripts/sync-pictures.py [source-dir]   # defaults to ~/Pictures
```

2. Rebuild WASM — `data.json` is embedded in it:

```bash
bash build-wasm.sh
```

## What it does

- Scans the source directory for `.png`, `.jpg`, `.jpeg`, `.gif`, `.webp` files
- Uploads three objects per picture to `www/images/pictures/`:
  - `<name>` — the master, as it came off the camera (`max-age=604800`)
  - `thumbs/<name>` — 100×100 thumbnail for Finder icons, via `sips`
  - `compressed/<stem>.avif` — the copy the simulation renders, via `ffmpeg`
    (`libsvtav1 -crf 22 -preset 6`, longest side 1600px). Animated gifs are skipped and
    keep pointing at their master.
- Reads natural dimensions with `ffprobe`, EXIF rotation applied
- Removes stale `~/Pictures/*` entries from `data.json`
- Adds new file nodes with `imagePath`, `thumbPath`, `compressedPath`, `imgWidth`, `imgHeight`
- Updates the `~/Pictures` children array
- `data.json` remains the single source of truth — new pictures need no Go code change

## Rendering rules

- Finder icons render `thumbPath`
- Quick Look and the Finder's click prefetch render `compressedPath`
- The Viewer opens `imagePath` — the master

## Wallpapers

`./scripts/sync-wallpapers.py` does the same job for the two desktop wallpapers, reading the
masters back from the bucket and writing `www/wallpapers/compressed/<stem>.avif` at up to
3200px. `desktop.go` points at those.
