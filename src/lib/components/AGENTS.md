# src/lib/components/

Svelte UI components. Dumb view layer — renders VM props, forwards events to `window.os.*`.

## Structure

```
components/
├── desktop/     # Desktop chrome: menubar, dock, wallpaper, icons, window layer
├── finder/      # Finder app window content
├── terminal/    # Terminal app window content
├── viewer/      # Preview app: image + text/markdown viewers
├── window/      # Generic window chrome (titlebar, traffic lights, resize)
└── placeholder/ # Splash screen for apps without implementation
```

## Component Conventions

- Svelte 5 Runes only: `$props`, `$state`, `$effect`. Never `export let` or `$:`.
- Props typed from `$lib/os/types.ts` interfaces
- Event forwarding: `os()?.<action>(args)` — never compute locally
- `$state` only for ephemeral UI (hover, focus, open menus, clock interval)
- No `$derived` for business logic — Go owns all derivations

## Window Content Dispatch

`WindowLayer.svelte` renders the right component based on `win.content.kind`:
- `'finder'` → `FinderApp.svelte`
- `'terminal'` → `TerminalApp.svelte`
- `'viewer-image'` → `ImageViewer.svelte`
- `'viewer-text'` → `TextViewer.svelte`
- `'splash'` → `SplashApp.svelte`

## CSS

All component styling lives in `src/app.css`, not scoped `<style>` blocks (except `Wallpaper.svelte`). Global classes: `.window`, `.titlebar`, `.traffic-lights`, `.menubar-*`, `.dock-*`, `.text-viewer-*`, `.markdown-body`, etc.
