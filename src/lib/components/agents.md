# src/lib/components — Agent Reference

All components use Svelte 5 Runes. No legacy `export let` or `$:` syntax.

## Active Shared Components

Route pages own their page structure directly. Shared components should stay small.

| Component | Purpose |
|-----------|---------|
| `SiteNav.svelte` | Sticky editorial three-part slider navigation for Home, CV, and Blog. |
| `Icon.svelte` | Icon registry using unplugin-icons for Material Symbols. |

## Design Constraints

- The visual system is editorial: Playfair Display headlines, Montserrat body text, JetBrains Mono metadata, 1px rulework, paper/obsidian theme tokens, and 0px border radius.
- Light and dark variants follow browser or OS `prefers-color-scheme`.
- Keep the home page as a three-section slider only.
- Blog posts are sourced from `blogs/*.md`, not JSON.
- Preserve the editorial design direction and avoid reviving prior experimental interface treatments.
