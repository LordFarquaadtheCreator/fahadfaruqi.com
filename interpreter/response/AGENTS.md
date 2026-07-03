# interpreter/response/

JSON response types for shell command output. Commands return `Response` structs which are marshaled to JSON and sent to the terminal frontend.

## Files

### response.go

`Type` constants: `LS`, `Cat`, `PWD`, `Echo`, `Env`, `History`, `Clear`, `VarSet`, `Open`, `Apps`, `Completion`, `Err`.

`Response` struct: `Type`, `Payload` (interface{}), `CWD`.

Payload types:
- `LSPayload` — directory entries with flags
- `CatPayload` — file name + content
- `TextPayload` — simple text
- `EnvPayload` — env vars map
- `HistoryPayload` — command history
- `CompletionPayload` — tab completion candidates
- `ErrorPayload` — command name + error message

`Error(cmd, msg, cwd)` — helper to build error response.
`Marshal(r)` — JSON marshal response.

## Relationship

- Used by all `command/` handlers
- Marshaled by `shell/run.go` and returned to JS via `interpRun()`
- Terminal frontend parses `type` field to render appropriate output
