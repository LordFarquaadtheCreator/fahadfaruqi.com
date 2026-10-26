# interpreter/command/

zsh command handlers. Each command is a function matching `CommandFunc` signature: `func(args []string, w world.World, s *session.Session) response.Response`.

## Files

### registry.go

`Registry` map: command name → handler function. Currently registered: `ls`, `cd`, `cat`, `pwd`, `echo`, `env`, `history`, `clear`, `open`, `apps`, `help`, `head`, `tail`, `wc`, `grep`, `find`, `tree`, `stat`, `which`, `date`, `whoami`, `hostname`, `uname`.

`terminal/complete.go` offers these for tab completion through `shell.Commands()`, which reads `Registry`. `which` cannot read `Registry` (that would be an initialisation cycle), so `which.go` keeps `builtins`, and `utilities_test.go` fails if the two lists ever disagree.

### util.go

Shared helpers: `splitFlags` (switches vs operands), `hasFlag`, `intFlag`/`flagValue` (options that take a value: `-n 5`, `-n5`, `--lines=5`), `file` (resolve one path, with the error to report), `contentLines`, `textResult` (plain-text output, which the terminal renders as preformatted).

### Command Handlers

- `ls.go` — directory listing with flags (`-l`, `-a`, `-t`, `-S`, `-r`, `--color`, etc.)
- `cd.go` — change directory, resolves `~` and relative paths
- `cat.go` — file content reader, markdown files get rendered to HTML
- `head.go`, `tail.go` — first and last lines, `-n N`
- `wc.go` — line, word and byte counts, `-l|-w|-c`
- `grep.go` — line search, `-i` `-n` `-v` `-c`; silence when nothing matches
- `find.go` — walk a tree, `-name` takes a shell pattern
- `tree.go` — indented folder tree, bounded at depth 4 and 120 lines
- `stat.go` — what the filesystem knows about a path
- `which.go` — where a command comes from
- `date.go`, `sysinfo.go` — time, `whoami`, `hostname`, `uname [-a]`
- `pwd.go` — print working directory
- `echo.go` — echo arguments
- `env.go` — print environment variables
- `history.go` — command history
- `clear.go` — clear terminal
- `open.go` — open a file, a folder, or a URL in appropriate app (Finder for dirs, viewer for files, the browser tab for URLs)
- `apps.go` — list registered apps
- `help.go` — show available commands

## Adding a New Command

1. Create `<name>.go` with handler function matching `CommandFunc` signature
2. Register in `registry.go`: `"name": HandlerFunc`
3. Add the name to `builtins` in `which.go` (the test enforces it)
4. Rebuild WASM

## Relationship

- Called by `shell/run.go` after tokenization and variable expansion
- Return `response.Response` which is marshaled to JSON and sent to terminal frontend
- `open.go` directly calls `w.Finder.Spawn()`, `w.Viewer.SpawnForFile()` and `w.Terminal.OpenPage()` — bridges command layer to window manager
