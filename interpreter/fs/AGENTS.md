# interpreter/fs/

In-memory filesystem. All nodes loaded from `world/data.json` at boot. Read-only — no write operations.

## Files

### node.go

`Node` struct: `Type` (Dir/File), `Children` (paths), `Content` (string), `Meta`, `Hidden` (bool).

`Meta` struct: `Size`, `Modified`, `ImagePath` (URL for the full-size image), `ThumbPath` (URL for the Finder thumbnail), `CompressedPath` (URL for the compressed copy the UI renders, empty when the file has none), `ImgWidth`, `ImgHeight`, `MimeType`.

`NodeType` constants: `Dir`, `File`.

### fs.go

`FileSystem` struct wraps `map[string]Node`.

Methods: `Stat(path)`, `Children(path, showHidden)`, `Exists(path)`, `IsDir(path)`, `IsFile(path)`.

### resolve.go

Path resolution: expands `~`, handles relative paths, resolves `.` and `..`.

## Relationship

- Loaded by `world/load.go` from `data.json`
- Used by `finder/`, `viewer/`, `desktop/`, `command/` packages
- All paths use `~` as home root (not `/home/user`)
- Hidden files have paths starting with `~/.`

## data.json Format

```json
{
  "~": {
    "type": "dir",
    "children": ["~/Documents", "~/notes.txt"],
    "hidden": false
  },
  "~/notes.txt": {
    "type": "file",
    "content": "file contents here",
    "mimeType": "text/plain",
    "hidden": false
  },
  "~/Pictures/sufi_12.JPG": {
    "type": "file",
    "content": "",
    "imagePath": "/images/pictures/sufi_12.JPG",
    "thumbPath": "/images/pictures/thumbs/sufi_12.JPG",
    "compressedPath": "/images/pictures/compressed/sufi_12.avif",
    "imgWidth": 717,
    "imgHeight": 978,
    "mimeType": "image/jpeg",
    "hidden": false
  }
}
```

The image fields are written by `scripts/sync-pictures.py`. Paths stay relative —
`assets.URL()` prefixes them in `world/load.go` — and the dimensions are the master's with
any EXIF rotation applied, which is the size a browser shows.
