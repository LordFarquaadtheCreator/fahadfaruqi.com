# src/lib/model/

Core types and data layer. No UI, no framework.

- **types.ts** — TypeScript definitions. `Theme` (27 theme names), `OutputType`, `OutputEntry`, `HistoryEntry`, `ListItem`, `JsonNode`, `FsNode` (`FsFile` | `FsDirectory`), `TerminalState`, `Hint`.
- **data.ts** — Imports `portfolio.json`. Exports typed data arrays: `experiences`, `education`, `projects`, `toolkit`, `profile`, `contact`, `blogInfo`. Helpers: `toListItem(raw)` converts `unknown` fs data to `ListItem`; `formatEntry(raw)` formats rich objects as readable text for `cat`.
- **filesystem.ts** — Builds virtual `FsNode` tree from portfolio data. Root `~` with children: `about`, `contact`, `blog`, `resume/` (`experience/`, `education/`, `projects/`, `toolkit/`). Exports `fs` and `getNode(path)`.
