# src/lib/components/window/

Generic window chrome. Shared by all app windows — titlebar, traffic lights, resize handle.

## Files

### Window.svelte

Renders window container with:
- **Titlebar**: Traffic lights (close/minimize/maximize), centered title text
- **Content slot**: App-specific content rendered via Svelte snippet (`{@render children()}`)
- **Resize handle**: Bottom-right corner drag handle

Props: `window: WindowVM`, `focused: boolean`, event callbacks (`onclose`, `onfocus`, `ondrag`, `onresize`, `onminimize`, `onmaximize`), `children` snippet.

Local state: `dragging`, `resizing` booleans for pointer event tracking. Global pointer listeners attached via `$effect` when dragging/resizing.

## Relationship

- Rendered by `WindowLayer.svelte` for each window in `windows[]`
- Event callbacks forwarded to `window.os.*` by `WindowLayer.svelte`
- CSS in `src/app.css` under `/* ── Window ── */` section
- `.window.focused` class adds focused border/shadow
- Minimized windows are hidden via CSS (not removed from DOM)
