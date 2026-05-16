# src/lib/components/

Svelte UI components for the terminal application.

- **Terminal.svelte** — Main container. Full-viewport flex layout, manages scroll, wires store to child components. Applies theme class from store.
- **History.svelte** — Renders scrollback. Loops `HistoryEntry` components. Pure presentational.
- **HistoryEntry.svelte** — Single command+output row. Shows prompt string, timestamp, delegates output rendering to `OutputRenderer`.
- **OutputRenderer.svelte** — Renders `OutputEntry` by type: `text` (pre), `list` (grid), `json` (tree), `error`/`success`/`warning` (colored text). No `cat` type — removed.
- **InputRow.svelte** — Command input. Handles key events: Enter (submit), Tab (complete), ArrowUp/Down (history). Shows prompt and cursor.
- **HintBar.svelte** — Clickable hint pills below input. Renders `hints` and `completionCandidates` from store. Touch-friendly on mobile.
