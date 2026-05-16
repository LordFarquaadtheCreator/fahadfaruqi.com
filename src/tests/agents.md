# src/tests/

Vitest tests. Two projects in `vite.config.ts`: `server` (node, `*.test.ts`) and `component` (jsdom, `*.svelte.test.ts`).

- **commands.test.ts** — Tests `dispatch()` and `parse()`: valid commands, invalid commands, theme switching, `cd` relative/absolute, `ls` glob, `cat` file content, `clear` behavior.
- **completer.test.ts** — Tests `complete()` and `deriveHints()`: path completion, command hints, partial matches.
- **glob.test.ts** — Tests `matches()`: prefix, suffix, infix wildcards, exact match, no match.
- **resolver.test.ts** — Tests `resolve()`: relative, absolute, `~`, `..`, `.`, edge cases.
- **terminalStore.test.ts** — Tests store behavior: init, `execute`, `cd`, `theme`, `clear`, history cycling, boot message, invalid theme returns usage text.
- **outputRenderer.svelte.test.ts** — Component tests for `OutputRenderer.svelte`: renders `text`, `list`, `json`, `error`, `success`, `warning` types correctly.
