# src/

Source root. SvelteKit 2 + Svelte 5 (Runes) + TypeScript terminal application.

- **app.css** — Global design system. CSS custom properties for 32 terminal themes. `0px border-radius` enforced. Breakpoints for mobile. `.sr-only` utility.

## Adding New Themes

### Pattern
Each theme is a CSS class on `.terminal` that defines 12 custom properties:
- `--gb-bg`, `--gb-bg-hard` — Background colors (normal and darker variant)
- `--gb-fg`, `--gb-fg2`, `--gb-fg3` — Foreground text colors (primary, accent, dimmed)
- `--gb-border`, `--gb-border2` — Border colors (primary and secondary)
- `--gb-green` through `--gb-blue` — Semantic color roles

### Example
```css
/* ──── New Theme Name ──── */
.terminal.theme-name {
  --gb-bg:      #hex_value;
  --gb-bg-hard: #darker_hex;
  --gb-fg:      #primary_light;
  --gb-fg2:     #accent_color;
  --gb-fg3:     #dim_text;
  --gb-border:  #border_color;
  --gb-border2: #dark_border;
  /* ... etc for all semantic colors */
}
```

### Color Selection Guidelines
1. **Contrast first** — Ensure `--gb-fg` has sufficient brightness against `--gb-bg`
2. **Maintain hierarchy** — `--gb-fg > --gb-fg2 > --gb-fg3` in lightness
3. **Semantic mapping**:
   - `--gb-green` → success, terminals, nature
   - `--gb-aqua` → links, highlights, water/cyberpunk accents
   - `--gb-purple` → commands, magical, mysterious elements
   - `--gb-yellow` → warnings, sand, classic terminal yellow
   - `--gb-orange` → UI elements, notifications, warmth
   - `--gb-red` → errors, danger, critical states
   - `--gb-blue` → primary actions, sky, information

## Directory map

- `lib/components/` — Svelte UI (Terminal, InputRow, History, OutputRenderer, HintBar)
- `lib/data/` — Static JSON data (`portfolio.json`)
- `lib/model/` — Types (`types.ts`), data helpers (`data.ts`), virtual filesystem (`filesystem.ts`)
- `lib/shell/` — Shell logic: completer, resolver, glob, command handlers
- `lib/stores/` — Svelte writable store for terminal state
- `routes/` — SvelteKit pages (single-page terminal)
- `tests/` — Vitest unit and component tests
