# interpreter/fs/

In-memory filesystem. All nodes loaded from `world/data.json` at boot. Read-only — no write operations.

## Files

### node.go

`Node` struct: `Type` (Dir/File), `Children` (paths), `Content` (string), `Meta`, `Hidden` (bool).

`Meta` struct: `Size`, `Modified`, `ImagePath` (URL for images), `MimeType`.

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
  }
}
```
