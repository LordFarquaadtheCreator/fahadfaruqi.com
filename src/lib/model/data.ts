import type { ListItem } from './types';
import portfolio from '../data/portfolio.json';

function slugify(s: string): string {
  return s.toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/-+$/, '');
}

// ── Raw portfolio data ──────────────────────────────────────

export const profile = portfolio.profile;
export const contact = portfolio.contact;
export const blogInfo = portfolio.blog;

export const experiences = portfolio.resume.experiences.map(e => ({
  id: slugify(e.company),
  ...e,
}));

export const education = portfolio.resume.education.map(e => ({
  id: slugify(e.institution),
  ...e,
}));

export const projects = portfolio.resume.projects.map(p => ({
  id: slugify(p.name),
  ...p,
}));

export const toolkit = portfolio.resume.toolkit.map(g => ({
  id: slugify(g.title),
  ...g,
}));

// ── ListItem conversion for ls ──────────────────────────────

export function toListItem(raw: unknown): ListItem {
  const r = raw as Record<string, unknown>;
  if ('company' in r && typeof r.company === 'string') {
    return {
      id: r.id as string,
      primary: `${r.role} — ${r.company}`,
      secondary: r.summary as string,
      meta: `${r.start} – ${r.end}`,
    };
  }
  if ('institution' in r && typeof r.institution === 'string') {
    return {
      id: r.id as string,
      primary: `${r.degree}`,
      secondary: r.institution as string,
      meta: `${r.start} – ${r.end}`,
    };
  }
  if ('name' in r && typeof r.name === 'string' && 'role' in r) {
    return {
      id: r.id as string,
      primary: r.name as string,
      secondary: r.role as string,
      meta: `${r.start} – ${r.end}`,
    };
  }
  if ('title' in r && Array.isArray(r.items)) {
    return {
      id: r.id as string,
      primary: r.title as string,
      secondary: (r.items as string[]).join(', '),
    };
  }
  if ('name' in r && typeof r.name === 'string') {
    return {
      id: slugify(r.name as string),
      primary: r.name as string,
    };
  }
  if ('title' in r && typeof r.title === 'string') {
    return {
      id: ('id' in r && typeof r.id === 'string') ? r.id : slugify(r.title as string),
      primary: r.title as string,
      secondary: 'description' in r ? String(r.description) : undefined,
    };
  }
  if ('email' in r && typeof r.email === 'string') {
    return {
      id: ('id' in r && typeof r.id === 'string') ? r.id : 'contact',
      primary: 'contact',
      secondary: r.email as string,
    };
  }
  return {
    id: 'unknown',
    primary: String(r),
  };
}

// ── Text formatting for cat ─────────────────────────────────

export function formatEntry(raw: unknown): string {
  const r = raw as Record<string, unknown>;
  const lines: string[] = [];

  if ('company' in r && typeof r.company === 'string') {
    lines.push(`company    : ${r.company}`);
    if (r.companyUrl) lines.push(`url        : ${r.companyUrl}`);
    if (r.logo) lines.push(`logo       : ${r.logo}`);
    lines.push(`role       : ${r.role}`);
    lines.push(`location   : ${r.location}`);
    lines.push(`duration   : ${r.start} – ${r.end}`);
    lines.push('');
    lines.push(String(r.summary));
    lines.push('');
    lines.push('highlights :');
    for (const h of r.highlights as string[]) lines.push(`  - ${h}`);
    lines.push('');
    lines.push(`tools      : ${(r.tools as string[]).join(', ')}`);
    lines.push(`skills     : ${(r.skills as string[]).join(', ')}`);
    return lines.join('\n');
  }

  if ('institution' in r && typeof r.institution === 'string') {
    lines.push(`institution : ${r.institution}`);
    lines.push(`degree      : ${r.degree}`);
    lines.push(`location    : ${r.location}`);
    lines.push(`duration    : ${r.start} – ${r.end}`);
    lines.push('');
    lines.push('highlights :');
    for (const h of r.highlights as string[]) lines.push(`  - ${h}`);
    return lines.join('\n');
  }

  if ('name' in r && typeof r.name === 'string' && 'role' in r) {
    lines.push(`name    : ${r.name}`);
    lines.push(`role    : ${r.role}`);
    lines.push(`duration: ${r.start} – ${r.end}`);
    lines.push('');
    lines.push(String(r.summary));
    lines.push('');
    lines.push('highlights :');
    for (const h of r.highlights as string[]) lines.push(`  - ${h}`);
    lines.push('');
    lines.push(`tools   : ${(r.tools as string[]).join(', ')}`);
    return lines.join('\n');
  }

  if ('title' in r && Array.isArray(r.items)) {
    lines.push(String(r.title));
    lines.push('');
    for (const item of r.items as string[]) lines.push(`  - ${item}`);
    return lines.join('\n');
  }

  if ('name' in r && typeof r.name === 'string') {
    lines.push(`name  : ${r.name}`);
    if (r.title) lines.push(`title : ${r.title}`);
    if (r.location) lines.push(`loc   : ${r.location}`);
    if (r.email) lines.push(`email : ${r.email}`);
    if (r.description) lines.push('');
    if (r.description) lines.push(String(r.description));
    return lines.join('\n');
  }

  return JSON.stringify(r, null, 2);
}
