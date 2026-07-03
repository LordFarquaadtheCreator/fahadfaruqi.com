# src/lib/components/terminal/

Terminal app window content. Renders command history and input prompt.

## Files

### TerminalApp.svelte

Renders terminal from `TerminalContentVM`:
- **History**: List of `TerminalEntry` (input + output). Output is JSON from `response.Response` — parsed and rendered based on `type` field.
- **Input prompt**: Shows CWD, accepts user input, forwards to `os()?.terminalRun(id, input)`.
- **Tab completion**: Calls `os()?.terminalTab(id, input)` for candidates.

Props: `window: WindowVM`, `content: TerminalContentVM`.

## Relationship

- Rendered by `WindowLayer.svelte` when `win.content.kind === 'terminal'`
- Command execution routed through `window.os.terminalRun` → Go `shell.Run()` → `command/` handlers
- Output rendering: `response.Response` JSON parsed in component, displayed by type (LS, Cat, Error, etc.)
- CSS in `src/app.css` under `/* ── Terminal ── */` section
