# src/lib/data — Agent Reference

## Single Source of Truth

Portfolio metadata, navigation, CV content, and the email address live in `portfolio.json`.
Blog articles live as Markdown files in `/blogs`.

`portfolio.ts` imports the JSON, validates it at module load, and exports typed data. Bad content should fail the build with a descriptive path.

## Content Shape

```typescript
interface PortfolioData {
  navigation: NavigationData;
  profile: ProfileData;
  home: HomeData;
  resume: ResumeData;
  blog: BlogData;
  contact: ContactData;
}

interface NavigationData {
  links: PortfolioLink[]; // Home, CV, Blog
}

interface HomeData {
  eyebrow: string;
  title: string;
  description: string;
  slides: HomeSlide[]; // Home, CV, Blog
}

interface ResumeData {
  downloadUrl?: string;
  title: string;
  description: string;
  experiences: ExperienceItem[];
}

interface BlogData {
  title: string;
  description: string;
}

interface ContactData {
  email: string;
}
```

## Blog Posts

Add presentation or production posts by creating `blogs/my-slug.md` with frontmatter:

```md
---
title: Example Post
category: Practice
date: May 2026
excerpt: Short card description.
imageAlt: Abstract editorial image description
codeTitle: code sample label
codeSnippet: const example = true;
pullQuote: A short pull quote.
---

Markdown body goes here.
```

## Validation Rules

- Required string fields must be non-empty.
- Optional fields: `downloadUrl`, `companyUrl`, `logo`, and `summary`.
- Arrays must exist and be arrays.
