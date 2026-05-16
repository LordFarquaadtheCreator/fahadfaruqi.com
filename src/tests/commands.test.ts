import { describe, it, expect } from 'vitest';
import { dispatch, parse } from '$lib/shell/commands/index';

describe('parse', () => {
  it('simple verb + arg',    () => expect(parse('cd resume')).toEqual({ verb: 'cd', args: ['resume'] }));
  it('quoted arg',           () => expect(parse('cat "my file"')).toEqual({ verb: 'cat', args: ['my file'] }));
  it('verb only',            () => expect(parse('ls')).toEqual({ verb: 'ls', args: [] }));
  it('extra whitespace',     () => expect(parse('  cd  resume  ')).toEqual({ verb: 'cd', args: ['resume'] }));
});

describe('dispatch — cd', () => {
  it('valid cd returns newPath',     () => expect(dispatch('~', 'cd resume').newPath).toBe('~/resume'));
  it('cd .. goes up',                () => expect(dispatch('~/resume', 'cd ..').newPath).toBe('~'));
  it('cd to file is error',          () => expect(dispatch('~/resume/experience', 'cd socialode').output.type).toBe('error'));
  it('cd nonexistent is error',      () => expect(dispatch('~', 'cd fake').output.type).toBe('error'));
  it('cd with no arg goes home',     () => expect(dispatch('~/resume', 'cd').newPath).toBe('~'));
  it('cd absolute path works',       () => expect(dispatch('~', 'cd ~/resume/experience').newPath).toBe('~/resume/experience'));
});

describe('dispatch — ls', () => {
  it('ls at root lists sections',    () => expect(dispatch('~', 'ls').output.type).toBe('list'));
  it('ls with glob filters',         () => {
    const out = dispatch('~/resume/experience', 'ls *').output;
    expect(out.type).toBe('list');
    expect((out.payload as any[]).length).toBeGreaterThan(0);
  });
  it('ls nonexistent glob is error', () => expect(dispatch('~', 'ls zzz*').output.type).toBe('error'));
});

describe('dispatch — cat', () => {
  it('cat a file returns text',      () => expect(dispatch('~/resume/experience', 'cat socialode').output.type).toBe('text'));
  it('cat a dir is error',           () => expect(dispatch('~', 'cat resume').output.type).toBe('error'));
  it('cat no arg is error',          () => expect(dispatch('~', 'cat').output.type).toBe('error'));
});

describe('dispatch — unknown', () => {
  it('unknown command is error',     () => expect(dispatch('~', 'rm -rf /').output.type).toBe('error'));
});
