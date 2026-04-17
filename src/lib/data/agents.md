# src/lib/data — Agent Reference

## Single Source of Truth

**All portfolio content lives in `portfolio.json`.** Zero component edits needed for content updates.

`portfolio.ts` imports the JSON, runs `validatePortfolio()`, and exports the typed result as the default export. Validation throws at module load time — a malformed JSON field will crash the build immediately with a descriptive error path (e.g. `Invalid portfolio data at resume.experiences[0].role`).

## TypeScript Interfaces

```typescript
interface PortfolioData {
  profile:  ProfileData;
  about:    AboutData;
  resume:   ResumeData;
  contact:  ContactData;
}

interface ProfileData {
  name:             string;
  title:            string;
  location:         string;
  heroIntro:        string;   // Not currently rendered in HeroSection
  supportingIntro:  string;   // Not currently rendered in HeroSection
  portraitImage?:   string;   // Optional; empty string treated as missing
  portraitAlt:      string;
  socialLinks:      PortfolioLink[];
}

interface AboutData {
  paragraphs:   string[];
  education:    EducationItem[];
  highlights:   HighlightItem[];
  photoCaption: string;
}

interface EducationItem {
  institution: string;
  credential:  string;
  dates:       string;
  location:    string;
}

interface HighlightItem {
  eyebrow:     string;
  title:       string;
  description: string;
  href?:       string;       // Optional external link
  linkLabel?:  string;       // Required if href is set
}

interface ResumeData {
  downloadUrl?:  string;         // Optional URL to resume PDF
  experiences:   ExperienceItem[];
}

interface ExperienceItem {
  company:     string;
  companyUrl?: string;
  logo?:       string;       // Path under /static/, e.g. /images/companies/foo.svg
  role:        string;
  location:    string;
  start:       string;       // e.g. "Jan 2020"
  end:         string;       // e.g. "Present" or "Dec 2022"
  summary?:    string;
  highlights:  string[];
  tools:       string[];
  skills:      string[];
}

interface ContactData {
  email:          string;
  subjectPrefix:  string;
  availability:   string;
  links:          PortfolioLink[];
}

interface PortfolioLink {
  label: string;
  href:  string;
}
```

## Common Edits

**Add a job** — append to `resume.experiences`:
```json
{
  "company": "Acme Corp",
  "companyUrl": "https://acme.com",
  "logo": "/images/companies/acme.svg",
  "role": "Senior Engineer",
  "location": "Remote",
  "start": "Jan 2024",
  "end": "Present",
  "summary": "Optional one-liner.",
  "highlights": ["Built X", "Led Y"],
  "tools": ["React", "Go"],
  "skills": ["frontend", "backend"]
}
```

**Add a resume download** — set `resume.downloadUrl` to a URL or `/resume.pdf` (place PDF in `static/`).

**Add a highlight card** — append to `about.highlights`:
```json
{
  "eyebrow": "CATEGORY",
  "title": "Card Title",
  "description": "Supporting text.",
  "href": "https://...",
  "linkLabel": "View"
}
```

**Add a social link** — append to `profile.socialLinks`:
```json
{ "label": "GitHub", "href": "https://github.com/..." }
```

## Validation Rules

- All string fields must be non-empty (trimmed). Empty string → throws.
- Optional fields: `portraitImage`, `companyUrl`, `logo`, `summary`, `downloadUrl`, `href`, `linkLabel` — these accept `null`, `undefined`, or `""` and will be returned as `undefined`.
- Arrays must exist and be arrays; each element is fully validated with path info in the error message.
- `readStringArray` validates every item in array fields like `highlights`, `tools`, `skills`.
