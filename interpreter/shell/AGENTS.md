# interpreter/shell/

Shell pipeline: tokenize → expand variables → run command.

## Files

### run.go

`Run(input, world, session)` — main entry point. Trims input, handles empty, checks for variable assignment (`VAR=value`), tokenizes, expands variables, looks up command in `command.Registry`, executes.

### tokenize.go

Splits input string into tokens. Handles quoted strings (single and double quotes).

### expand.go

Variable expansion: replaces `$VAR` and `~` with values from session/env. `IsVarAssignment(input)` detects `VAR=value` patterns.

## Relationship

- Called by `main.go` via `interpRun(input)` binding
- Calls into `command/` registry to execute commands
- Returns marshaled `response.Response` JSON string
- `world.World` provides filesystem and window manager access
