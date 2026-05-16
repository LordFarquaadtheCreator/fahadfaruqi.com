import { describe, it, expect } from 'vitest';
import { match, longestCommonPrefix } from '$lib/shell/glob';

const pool = ['senior-eng', 'mid-eng', 'junior-eng', 'experience', 'skills'];

describe('match', () => {
  it('prefix match without wildcard', () => expect(match('exp', pool)).toEqual(['experience']));
  it('suffix wildcard',               () => expect(match('*eng', pool)).toEqual(['senior-eng','mid-eng','junior-eng']));
  it('prefix wildcard',               () => expect(match('mid*', pool)).toEqual(['mid-eng']));
  it('double wildcard',               () => expect(match('*i*', pool)).toEqual(['senior-eng','mid-eng','junior-eng','experience','skills']));
  it('star matches all',              () => expect(match('*', pool)).toHaveLength(pool.length));
  it('no match returns []',           () => expect(match('zzz', pool)).toEqual([]));
  it('case insensitive',              () => expect(match('EXP*', pool)).toEqual(['experience']));
});

describe('longestCommonPrefix', () => {
  it('single item returns itself',  () => expect(longestCommonPrefix(['mid-eng'])).toBe('mid-eng'));
  it('common prefix',               () => expect(longestCommonPrefix(['senior-eng','skills'])).toBe('s'));
  it('no common prefix',            () => expect(longestCommonPrefix(['abc','xyz'])).toBe(''));
  it('empty array',                 () => expect(longestCommonPrefix([])).toBe(''));
});
