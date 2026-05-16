# src/lib/shell/

Terminal shell logic. Pure functions, no UI.

- **completer.ts** — Tab completion engine. `deriveHints(cwd, input)` returns command/path suggestions. `complete(cwd, input)` expands to longest common prefix or returns candidates array.
- **resolver.ts** — Path resolution. `resolve(cwd, path)` normalizes relative/absolute paths, handles `~`, `..`, `.`. Returns resolved path string.
- **glob.ts** — Wildcard matching for `ls`. `matches(pattern, name)` supports `*` prefix/suffix/infix globs. Used by `ls` to filter directory listings.
