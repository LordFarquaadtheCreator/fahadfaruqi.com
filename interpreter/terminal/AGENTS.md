# interpreter/terminal/

Terminal window manager and ViewModel builder. Each terminal window has its own session and command history.

## Files

### terminal.go

`CommandRunner` interface — implemented by `world.World` to avoid import cycle. Methods: `RunCommand(input, session)`, `TabComplete(input, cwd)`.

`TerminalState` — per-window: `Session *session.Session`, `History []TerminalEntry`.

`TerminalEntry` — `Input` string, `Output` string (JSON response).

`TerminalManager` — owns `states map[string]*TerminalState`, wraps `wm.Manager`.

Methods: `Spawn()`, `RunCommand(id, input)`, `TabComplete(id, input)`, `GetState(id)`, `Cleanup(id)`.

`BuildContentVM(win, tm)` — builds `TerminalContentVM` with `kind: "terminal"`, `cwd`, `history[]`.

## Relationship

- `world.World` implements `CommandRunner` via `shell.Run()` and `command/` handlers
- Each terminal window gets independent `session.Session` with its own CWD
- Called by `main.go:buildContent()` when `win.AppType == "terminal"`
- Frontend component: `src/lib/components/terminal/TerminalApp.svelte`
- Window title uses `app.Name` ("iTerm")
