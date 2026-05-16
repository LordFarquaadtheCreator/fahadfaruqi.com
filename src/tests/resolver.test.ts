import { describe, it, expect } from 'vitest';
import { resolve, resolveAndValidate, parent } from '$lib/shell/resolver';

function path(cwd: string, input: string): string {
  const r = resolve(cwd, input);
  if (!r.ok) throw new Error(r.error);
  return r.path;
}

describe('resolve', () => {
  it('relative from root',          () => expect(path('~', 'resume')).toBe('~/resume'));
  it('relative nested',             () => expect(path('~/resume', 'experience')).toBe('~/resume/experience'));
  it('dot-dot one level',           () => expect(path('~/resume', '..')).toBe('~'));
  it('dot-dot multi-level',         () => expect(path('~/resume/experience', '../..')).toBe('~'));
  it('dot-dot past root clamps',    () => expect(path('~', '../../..')).toBe('~'));
  it('absolute path',               () => expect(path('~/resume', '~/blog')).toBe('~/blog'));
  it('absolute with dot-dot',       () => expect(path('~', '~/resume/../blog')).toBe('~/blog'));
  it('tilde alone',                 () => expect(path('~/resume', '~')).toBe('~'));
  it('empty input stays at cwd',    () => expect(path('~/resume', '')).toBe('~/resume'));
  it('cross-directory shortcut',    () => expect(path('~/resume', '../blog')).toBe('~/blog'));
});

describe('resolveAndValidate', () => {
  it('valid path resolves',         () => expect(resolveAndValidate('~', 'resume').ok).toBe(true));
  it('invalid path errors',         () => expect(resolveAndValidate('~', 'nonexistent').ok).toBe(false));
  it('deep valid path resolves',    () => expect(resolveAndValidate('~', 'resume/experience').ok).toBe(true));
});

describe('parent', () => {
  it('from nested',                 () => expect(parent('~/resume/experience')).toBe('~/resume'));
  it('from one level',              () => expect(parent('~/resume')).toBe('~'));
  it('from root stays root',        () => expect(parent('~')).toBe('~'));
});
