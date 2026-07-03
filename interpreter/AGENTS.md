# interpreter/

Go WASM module. Owns all domain logic per MVVM rule.

## Modules

- `command/` — zsh command handlers (see command/registry.go)
- `finder/` — Window manager + ViewModel + markdown render (see finder/AGENTS.md)
- `fs/` — In-memory filesystem from data.json
- `response/` — JSON response types sent to JS
- `session/` — Shell session (cwd, vars, history)
- `shell/` — Tokenizer, expander, runner
- `world/` — Boots FS + env from data.json
- `main.go` — WASM entry. Exposes `interpRun` and `window.os.*` bindings.

## Build

```bash
./build-wasm.sh   # GOOS=js GOARCH=wasm go build -o static/interp.wasm
```

## Key Rule

Svelte never reads FS directly. All state flows through `window.os.subscribe`.
