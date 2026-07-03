# src/lib/components/viewer/

Preview app components. Renders image and text/markdown file content.

## Files

### ImageViewer.svelte

Renders image from `ViewerContentVM`:
- Full-bleed image centered in window
- `src` from VM (image URL path)
- Zoom not implemented yet

Props: `content: ViewerContentVM`.

### TextViewer.svelte

Renders text/markdown from `ViewerContentVM`:
- **Toolbar**: 36px height, centered filename, share button (SVG icon)
- **Markdown mode** (`isMarkdown: true`): `{@html content.content}` in `.markdown-body` div. Content is pre-rendered HTML from Go (goldmark + bluemonday).
- **Plain text mode** (`isMarkdown: false`): `<pre>` with monospace, pre-wrap, word break

Props: `content: ViewerContentVM`.

## Relationship

- Rendered by `WindowLayer.svelte` based on `win.content.kind`:
  - `'viewer-image'` → `ImageViewer.svelte`
  - `'viewer-text'` → `TextViewer.svelte`
- Both share `TypeViewer` in Go — dispatch is by `ViewerContentVM.kind`, not `win.appType`
- Markdown HTML is generated server-side in Go (goldmark + bluemonday sanitization)
- CSS in `src/app.css` under `/* ── Viewer ── */` and `/* ── Markdown Body ── */` sections
