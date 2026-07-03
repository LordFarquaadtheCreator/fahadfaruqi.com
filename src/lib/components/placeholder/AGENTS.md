# src/lib/components/placeholder/

Splash screen for apps without full implementation. Shows app icon and "Coming Soon" message.

## Files

### SplashApp.svelte

Renders centered splash with app icon, name, and "Coming Soon" text.

Props: `content: SplashContentVM` (`appId`, `name`, `icon`).

## Relationship

- Rendered by `WindowLayer.svelte` when `win.content.kind === 'splash'`
- Used for `TypePlaceholder` apps (Obsidian, Slack, Spotify, Steam, Affinity, WhatsApp, Messages) and `TypeEditor` apps (Devin, Zed, Xcode, Figma, Hermes) until they get real implementations
- CSS in `src/app.css` under `/* ── Splash ── */` section
