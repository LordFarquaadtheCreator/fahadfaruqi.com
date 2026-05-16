# src/routes/

SvelteKit routes. This is a single-page terminal app.

- **+layout.svelte** — Root layout. Sets font imports, favicon, theme-color meta tag. Wraps all pages.
- **+layout.ts** — `export const prerender = true` for static site generation.
- **+page.svelte** — Landing page. Renders `<Terminal />` full-viewport. No padding, no centering. Terminal owns the entire screen.
