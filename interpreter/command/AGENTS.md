# interpreter/command/

zsh command handlers. Each command is a function matching `CommandFunc` signature: `func(args []string, w world.World, s *session.Session) response.Response`.

## Files

### registry.go

`Registry` map: command name → handler function. Currently registered: `ls`, `cd`, `cat`, `pwd`, `echo`, `env`, `history`, `clear`, `open`, `apps`, `help`.

### Command Handlers

- `ls.go` — directory listing with flags (`-l`, `-a`, `-t`, `-S`, `-r`, `--color`, etc.)
- `cd.go` — change directory, resolves `~` and relative paths
- `cat.go` — file content reader, markdown files get rendered to HTML
- `pwd.go` — print working directory
- `echo.go` — echo arguments
- `env.go` — print environment variables
- `history.go` — command history
- `clear.go` — clear terminal
- `open.go` — open file or directory in appropriate app (Finder for dirs, Preview for files)
- `apps.go` — list registered apps
- `help.go` — show available commands

## Adding a New Command

1. Create `<name>.go` with handler function matching `CommandFunc` signature
2. Register in `registry.go`: `"name": HandlerFunc`
3. Rebuild WASM

## Relationship

- Called by `shell/run.go` after tokenization and variable expansion
- Return `response.Response` which is marshaled to JSON and sent to terminal frontend
- `open.go` directly calls `w.Finder.Spawn()` and `w.Viewer.SpawnForFile()` — bridges command layer to window manager
