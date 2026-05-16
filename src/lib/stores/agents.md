# src/lib/stores/

Svelte stores for terminal state management.

- **terminalStore.ts** — `createTerminalStore()` returns a Svelte writable store with methods:
  - `execute(raw)` — parses, dispatches, appends `HistoryEntry`, updates `cwd` if `cd`.
  - `reset()` — clears history, restores defaults.
  - `setTheme(name)` — validates and sets `TerminalState.theme`.
  - Boot message auto-generated from `profile.name` on init.
  - State shape: `{ cwd, theme, history, hints, completionCandidates }`.
