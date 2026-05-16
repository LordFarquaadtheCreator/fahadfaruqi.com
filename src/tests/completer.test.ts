import { describe, it, expect } from 'vitest';
import { complete, deriveHints } from '$lib/shell/completer';

describe('complete — verb completion', () => {
  it('c → cd, cat, clear',   () => expect(complete('~', 'c').candidates).toContain('cd'));
  it('ls → single, fills in',() => expect(complete('~', 'ls').completed).toContain('ls'));
});

describe('complete — path completion', () => {
  it('cd res → resume/',     () => expect(complete('~', 'cd res').completed).toBe('cd resume/'));
  it('cd → candidates',      () => expect(complete('~', 'cd ').candidates).toContain('resume'));
  it('ls soci at experience', () => expect(complete('~/resume/experience', 'ls soci').completed).toBe('ls socialode '));
  it('ls * matches all',     () => expect(complete('~/resume/experience', 'ls *').candidates.length).toBeGreaterThan(0));
  it('no match unchanged',   () => expect(complete('~', 'cd zzz').completed).toBe('cd zzz'));
});

describe('complete — theme argument', () => {
  it('theme gru → gruvbox-dark', () => expect(complete('~', 'theme gru').candidates).toContain('gruvbox-dark'));
  it('theme matrix → exact', () => expect(complete('~', 'theme matrix').completed).toBe('theme matrix '));
  it('theme zzz → no match', () => expect(complete('~', 'theme zzz').candidates).toEqual([]));
});

describe('deriveHints', () => {
  it('root has cd resume and cat blog', () => {
    const h = deriveHints('~');
    expect(h).toContain('cd resume');
    expect(h).toContain('cat blog');
  });
  it('experience dir has cat hints', () => {
    const h = deriveHints('~/resume/experience');
    expect(h.some(x => x.startsWith('cat'))).toBe(true);
  });
  it('root has no cd .. hint', () => {
    expect(deriveHints('~')).not.toContain('cd ..');
  });
  it('subdirectory has cd .. hint', () => {
    expect(deriveHints('~/resume')).toContain('cd ..');
  });
  it('always has theme hint', () => {
    expect(deriveHints('~')).toContain('theme');
  });
});
