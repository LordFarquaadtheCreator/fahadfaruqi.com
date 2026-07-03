# src/lib/

Svelte frontend library. Dumb view layer — renders VM JSON from Go WASM, forwards events back via `window.os.*`.

## Structure

```
lib/
├── os/                    # VM types + OS bridge
│   ├── types.ts           # TypeScript interfaces matching Go VM structs
│   └── osStore.svelte.ts  # VM subscription store, initOS(), os() accessor
├── components/            # Svelte UI components (see components/AGENTS.md)
├── data/                  # Static data (portfolio.json)
├── wasmLoader.js          # Loads Go WASM module
└── terminalStore.ts       # Legacy terminal store (may be unused)
```

## Critical Rule

Svelte components MUST NOT:
- Own app state (use `$state` only for ephemeral UI: hover, focus, open menus)
- Derive view-model logic with `$derived`
- Parse or compute business logic
- Read filesystem directly

Svelte components MUST:
- Render props from `RootVM` types
- Forward user events to Go via `window.os.*` bindings
- Use Svelte 5 Runes (`$props`, `$state`, `$effect`) — never legacy `export let` / `$:`

## Data Flow

```
Go WASM → os.subscribe(json) → osStore parses → subscribers notified → Svelte re-renders
User click → window.os.<action>(args) → Go mutates state → pushVM() → cycle repeats
```
