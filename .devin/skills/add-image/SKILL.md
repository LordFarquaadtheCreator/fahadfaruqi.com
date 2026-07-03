---
name: add-image
description: when the user needs to upload a folder to Pictures/
---

# Sync Pictures

Syncs all images in `static/images/pictures/` into `interpreter/world/data.json` as `~/Pictures/<filename>` entries.

## When to use

- After adding or removing images from `static/images/pictures/`
- Before building WASM if pictures changed

## Steps

1. Run the sync script:

```bash
python3 scripts/sync-pictures.py
```

2. Rebuild WASM:

```bash
bash build-wasm.sh
```

## What it does

- Scans `static/images/pictures/` for `.png`, `.jpg`, `.jpeg`, `.gif`, `.webp` files
- Removes stale `~/Pictures/*` entries from `data.json`
- Adds new file nodes with `imagePath` pointing to `/images/pictures/<name>`
- Updates `~/Pictures` children array
- `data.json` remains single source of truth — no Go code changes

