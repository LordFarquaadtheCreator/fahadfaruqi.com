# fahadfaruqi.com — Tactical Archive Portfolio

A **classified terminal interface** portfolio built with **Svelte 5**, **SvelteKit**, and **TypeScript**. Designed with a "Tactical Archive" aesthetic — high-contrast dark mode, acidic signal colors, and hardware-inspired UI components. All content is managed through a single JSON file (`portfolio.json`).

---

## Design System: The Tactical Archive

**Creative Direction**: "The Tactical Archive" — positions the user as an operator of high-end, futuristic hardware. An aesthetic of **Graphic Realism** where digital interfaces bear the "scars" of physical existence — film grain, chromatic aberration, and halftone patterns.

### Color Palette

| Role | Color | Usage |
|------|-------|-------|
| Background | `#121314` | Dark canvas |
| Signal Red (Primary) | `#ffb3b2` / `#bf002a` | CTAs, critical actions |
| Electric Blue (Secondary) | `#d3fbff` / `#00eefc` | Data streams, navigation |
| Radioactive Green (Tertiary) | `#2ae500` | Active status, success |
| Surface Tiers | `#1b1c1d` → `#343536` | Tonal layering |

### Typography

- **Space Grotesk**: Display headlines, tactical labels
- **Work Sans**: Body text, readable content
- **JetBrains Mono**: Technical metadata, timestamps, code

### Key Design Rules

- **0px border radius** — hard corners only
- **No 1px borders** — use surface color shifts for structure
- **Film grain overlay** (5% opacity) on all surfaces
- **Chromatic aberration** on high-impact text
- **HUD brackets** on featured cards
- **Tonal layering** instead of drop shadows

---

## Tech Stack

| Layer | Technology |
|-------|------------|
| Framework | SvelteKit 2.x with Svelte 5 (Runes mode) |
| Language | TypeScript |
| Styling | CSS Custom Properties + Scoped Component Styles |
| Icons | Material Symbols Outlined |
| Build Output | Static HTML via `@sveltejs/adapter-static` |
| Testing | Vitest + `@testing-library/svelte` |
| Fonts | Space Grotesk + Work Sans + JetBrains Mono |
| Effects | Glitch hover animations, scroll-based nav highlighting |

---

## Architecture Overview

### File Structure

```
src/
├── app.css              # Global styles, CSS variables, utility classes
├── app.html             # HTML template
├── routes/
│   ├── +layout.svelte   # Root layout (fonts, favicon, meta)
│   ├── +layout.ts       # Prerender config
│   ├── +page.svelte     # Main page composing all sections
│   └── page.svelte.spec.ts  # Component tests
├── lib/
│   ├── index.ts         # Public exports
│   ├── actions/
│   │   └── reveal.ts    # Intersection Observer scroll reveal action
│   ├── components/      # All section components
│   │   ├── HeroSection.svelte      # Tactical dashboard
│   │   ├── AboutSection.svelte     # Intel dossier
│   │   ├── ResumeSection.svelte    # Operation logs
│   │   ├── ContactSection.svelte   # Uplink protocol
│   │   ├── SiteNav.svelte          # Side navigation with scroll highlighting
│   │   ├── TacticalCard.svelte     # Reusable card with HUD brackets
│   │   ├── TacticalSectionHeader.svelte  # Reusable section header
│   │   └── MediaFrame.svelte
│   ├── data/
│   │   ├── portfolio.json   # ALL content lives here
│   │   └── portfolio.ts   # Typed loader + type exports
│   ├── utils/
│   │   ├── mailto.ts      # Contact form mailto generator
│   │   └── mailto.spec.ts
│   └── assets/
│       └── favicon.svg
static/
├── images/
│   ├── companies/       # Company logos (SVG)
│   └── profile/         # Portrait images
└── robots.txt
```

### SvelteKit Route Structure

This is a **single-page app** with one route:
- `+page.svelte` — renders all sections sequentially
- `+layout.ts` exports `prerender = true` → generates static HTML at build time
- No server-side logic; everything runs client-side

---

## Content Management (JSON-Driven)

**All content is stored in `/src/lib/data/portfolio.json`**. No component code changes needed to update copy, add jobs, or change links.

### Data Structure

```typescript
interface PortfolioData {
  profile: {
    name: string;
    title: string;
    location: string;
    heroIntro: string;
    supportingIntro: string;
    portraitImage?: string;   // Optional path; empty treated as missing
    portraitAlt: string;
    socialLinks: PortfolioLink[];
  };
  about: {
    paragraphs: string[];
    education: EducationItem[];
    highlights: HighlightItem[];
    photoCaption: string;
  };
  resume: {
    downloadUrl?: string;       // Optional URL to resume PDF
    experiences: ExperienceItem[];  // Each with company, role, highlights, tools, skills
  };
  contact: {
    email: string;
    subjectPrefix: string;
    availability: string;
    links: PortfolioLink[];
  };
}
```

### Adding a New Job

Add to `resume.experiences` array in `portfolio.json`:

```json
{
  "company": "Company Name",
  "companyUrl": "https://...",
  "logo": "/images/companies/company.svg",
  "role": "Job Title",
  "location": "City, ST",
  "start": "Jan 2024",
  "end": "Present",
  "summary": "One-line description...",
  "highlights": ["Achievement 1", "Achievement 2"],
  "tools": ["React", "TypeScript"],
  "skills": ["frontend", "architecture"]
}
```

### Adding Images

- **Company logos**: Place SVG in `/static/images/companies/`
- **Portrait**: Set `profile.portraitImage` to path (e.g., `/images/profile/portrait.jpg`)
- **MediaFrame** shows a placeholder with initials if image fails to load or is empty

---

## Component System

### Section Components (Page-level)

| Component | Purpose | Data Prop |
|-----------|---------|-----------|
| `HeroSection` | Bento-grid nav hub (desktop + mobile layouts) | `profile`, `resumeDownloadUrl?` |
| `AboutSection` | Bio paragraphs, education records, highlights grid | `about` |
| `ResumeSection` | Work experience cards with scroll reveal | `resume` |
| `ContactSection` | Registry panel + mailto transmission form | `contact` |
| `SiteNav` | Fixed side nav (desktop) / top+bottom nav (mobile) | `name`, `activeSection` |
| `TacticalCard` | Card surface with optional HUD brackets, header, variants | various |
| `TacticalSectionHeader` | Section header with ghost background text | various |

> **Note**: `PhotoSection`, `SectionHeading`, and `MediaFrame` exist in the codebase but are not used in the current page. Do not add them back without review.

### Shared Components

| Component | Purpose |
|-----------|---------|
| `MediaFrame` | Image container with error fallback (logo or portrait variant) |
| `TacticalCard` | Card container with optional HUD brackets, header, variants |
| `TacticalSectionHeader` | Section header with ghost text, eyebrow, title, description |

### Svelte 5 Patterns

Components use the new **Runes** syntax:

```svelte
<script lang="ts">
  let { profile, resumeDownloadUrl } = $props<{
    profile: ProfileData;
    resumeDownloadUrl?: string;
  }>();
</script>
```

State management:
```svelte
let name = $state('');
let activeSection = $state('ops');

const operatorId = $derived('OPR_' + name.slice(0, 4).toUpperCase() + '01');
const mailtoHref = $derived(buildMailtoHref({ ... }));
```

---

## Styling Architecture

### CSS Custom Properties (Variables)

Defined in `app.css` (Tactical Archive dark theme):

```css
:root {
  --background: #121314;           /* Dark canvas */
  --surface-container-low: #1b1c1d; /* Surface tier 1 */
  --surface-container: #1f2021;     /* Surface tier 2 */
  --surface-container-high: #292a2b;/* Surface tier 3 */
  --on-surface: #e4e2e3;            /* Primary text */
  --on-surface-variant: #e9bcba;    /* Secondary text */
  --primary: #ffb3b2;               /* Signal red */
  --secondary: #d3fbff;             /* Electric blue */
  --tertiary: #2ae500;              /* Radioactive green */
  --outline: #af8786;               /* Structural lines */
}
```

### Global Utility Classes (`app.css`)

| Class | Purpose |
|-------|---------|
| `.section-shell` | Max-width container (1200px) |
| `.surface` | Tactical card surface style |
| `.eyebrow` | Technical label with mono font |
| `.button` | Hardware switch button style |
| `.button--ghost` | Outline button variant |
| `.chip` | Technical tag/pill style |
| `.reveal` | Scroll-reveal animation base class |
| `.chromatic` | Chromatic aberration text effect |
| `.tech-label` | JetBrains Mono uppercase label |
| `.status-pip` | Status indicator dot |
| `.hud-bracket-*` | HUD corner brackets |
| `.sr-only` | Screen reader only content |

### Component Styles

Each component has scoped `<style>` block using CSS nesting. Mobile breakpoints:
- `1024px` — Tablet (grid columns collapse)
- `900px` — Hero/dashboard stacks
- `768px` — Side nav hidden, single column layout

### Scroll Reveal Animation

The `reveal` Svelte action (`src/lib/actions/reveal.ts`) uses IntersectionObserver:
- Elements start with `opacity: 0` and `translateY(16px)`
- Triggered at 18% visibility
- Respects `prefers-reduced-motion`
- Apply with `use:reveal` attribute

### Scroll-Based Navigation

The side navigation automatically highlights the active section based on scroll position:
- Uses IntersectionObserver to track section visibility
- Maps section IDs to nav keys: `top→home`, `about→about`, `resume→resume`, `contact→contact`
- Updates `activeSection` prop on `SiteNav` component

### Hover Effects (Glitch Animation)

All information cards feature a tactical glitch effect on hover:
- Subtle shake animation (translate X/Y jitter)
- Color-coded glow: Red (Hero/About), Blue (Resume), Green (Contact)
- Border color intensifies
- Implemented via CSS `@keyframes glitch`

---

## Key Features

1. **Tactical Archive Aesthetic** — Dark high-contrast UI with acidic signal colors
2. **Content-First Architecture** — All copy in JSON, zero component edits for content updates
3. **Scroll-Based Navigation** — Side nav highlights active section automatically via IntersectionObserver
4. **Glitch Hover Effects** — Tactical shake animation on all interactive cards
5. **Contact Form** — Client-side form generates mailto: link with pre-filled subject/body
6. **Static Generation** — Pre-rendered HTML for fast loads and easy hosting
7. **Accessibility** — Semantic HTML, reduced-motion support, proper ARIA labels
8. **Mobile-First Responsive** — Breakpoints at 768px and 1024px
9. **Easter Egg** — Footer brand (`TACTICAL_ARCHIVE`) is clickable; 5 clicks triggers a redirect

---

## Development Workflow

### Install Dependencies

```bash
npm install
```

### Start Dev Server

```bash
npm run dev
# or
npm run dev -- --open
```

### Type Check

```bash
npm run check
npm run check:watch
```

### Run Tests

```bash
npm test              # Unit + component tests
npm run test:unit     # Vitest only
```

### Build for Production

```bash
npm run build
```

Output goes to `build/` directory (static HTML/CSS/JS).

### Preview Production Build

```bash
npm run preview
```

---

## Deployment

The site is configured for **static hosting** (Netlify, Vercel, GitHub Pages, etc.):

1. Build: `npm run build`
2. Deploy `build/` folder contents

The `svelte.config.js` uses `@sveltejs/adapter-static` which generates:
- `index.html` (prerendered)
- `_app/` (JS/CSS chunks)
- Static assets from `static/`

---

## Customization Guide

### Change Colors

Edit CSS variables in `src/app.css`:

```css
:root {
  --bg: #your-color;
  --accent: #your-accent;
}
```

### Change Fonts

Edit `src/routes/+layout.svelte`:
- Update Google Fonts `<link>`
- Update `font-family` in `app.css` body selector

### Add New Section

1. Create component in `src/lib/components/`
2. Import and add to `src/routes/+page.svelte`
3. Add data type to `portfolio.ts`
4. Add content to `portfolio.json`

### Add Resume Download

Add to `portfolio.json`:
```json
{
  "resume": {
    "downloadUrl": "/resume.pdf",
    "experiences": [...]
  }
}
```

Then place PDF in `static/resume.pdf`.

---

## Project Recreation

To recreate this exact project setup:

```bash
npx sv@latest create --template minimal --types ts \
  --add vitest sveltekit-adapter="adapter:static" \
  --install npm .
```

---

## Recent Changes

### v2.0 — Tactical Archive Redesign
- Complete visual overhaul to dark tactical theme
- Removed top navigation, moved brand to vertical side nav
- Added scroll-based navigation highlighting with IntersectionObserver
- Added glitch hover effects on all cards
- Created reusable `TacticalCard` and `TacticalSectionHeader` components
- Dynamic page title derived from `portfolio.json` profile name
- All card content data-driven from `portfolio.json`
- HeroSection desktop/mobile dual layouts (portrait area is WIP)
- Added easter egg: 5-click footer brand trigger

