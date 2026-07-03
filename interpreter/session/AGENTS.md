# interpreter/session/

Shell session state. One per terminal window.

## Files

### session.go

`Session` struct: `CWD` (current working directory), `Vars` (map), `History` (command strings).

`New(env)` — creates session with CWD set to `$HOME` (defaults to `~`).

Methods: `SetVar(k, v)`, `AppendHistory(input)`, `SetCWD(path)`.

## Relationship

- Created by `terminal/terminal.go` on `Spawn()` — one session per terminal window
- Passed to `command/` handlers as third argument
- `cd` command mutates `CWD`
- `echo` with `VAR=value` mutates `Vars`
- All commands append to `History`
