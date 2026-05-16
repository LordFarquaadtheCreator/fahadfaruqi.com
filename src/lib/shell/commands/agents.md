# src/lib/shell/commands/

Command handlers. Each returns `OutputEntry` (or `DispatchResult` for `cd`).

- **index.ts** — `parse(raw)` tokenizes input (handles quotes). `dispatch(cwd, raw)` routes to handler. Returns `{ output, newPath? }`. `theme` command returns usage text (not error) on invalid/missing input.
- **cd.ts** — `cd(cwd, args)` navigates filesystem. Supports relative, absolute, `~`, `..`. Returns `DispatchResult` with `newPath`. Errors on non-existent or non-directory targets.
- **ls.ts** — `ls(cwd, args)` lists directory. Supports glob filter (e.g. `ls *eng`). Returns `list` output with `ListItem[]`. Uses `toListItem` for `FsFile` data.
- **cat.ts** — `cat(cwd, args)` reads file. Returns `text` output with formatted entry via `formatEntry`. Errors on directories or missing files.
- **help.ts** — `help()` returns static usage text for all commands and keybindings.
