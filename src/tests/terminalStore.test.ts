import { describe, it, expect, beforeEach } from 'vitest';
import { get } from 'svelte/store';
import { terminal } from '$lib/stores/terminalStore';

describe('execute — cd', () => {
  beforeEach(() => terminal.reset());

  it('navigates to resume', () => {
    terminal.execute('cd resume');
    expect(get(terminal).path).toBe('~/resume');
  });
  it('rejects invalid target at ~', () => {
    terminal.execute('cd projects');
    const last = get(terminal).history.at(-1)!;
    expect(last.output.type).toBe('error');
  });
  it('cd .. from ~/resume returns ~', () => {
    terminal.execute('cd resume');
    terminal.execute('cd ..');
    expect(get(terminal).path).toBe('~');
  });
  it('cd to file is error', () => {
    terminal.execute('cd resume');
    terminal.execute('cd experience');
    terminal.execute('cd socialode');
    const last = get(terminal).history.at(-1)!;
    expect(last.output.type).toBe('error');
  });
});

describe('execute — ls', () => {
  beforeEach(() => terminal.reset());

  it('ls at root lists sections', () => {
    terminal.execute('ls');
    const last = get(terminal).history.at(-1)!;
    expect(last.output.type).toBe('list');
    expect(Array.isArray(last.output.payload)).toBe(true);
  });
  it('ls with glob filters', () => {
    terminal.execute('cd resume');
    terminal.execute('cd experience');
    terminal.execute('ls *');
    const last = get(terminal).history.at(-1)!;
    expect(last.output.type).toBe('list');
    expect((last.output.payload as any[]).length).toBeGreaterThan(0);
  });
  it('ls nonexistent is error', () => {
    terminal.execute('ls zzz*');
    const last = get(terminal).history.at(-1)!;
    expect(last.output.type).toBe('error');
  });
});

describe('execute — cat', () => {
  beforeEach(() => terminal.reset());

  it('cat a file returns text', () => {
    terminal.execute('cd resume');
    terminal.execute('cd experience');
    terminal.execute('cat socialode');
    const last = get(terminal).history.at(-1)!;
    expect(last.output.type).toBe('text');
  });
  it('cat a dir is error', () => {
    terminal.execute('cat resume');
    const last = get(terminal).history.at(-1)!;
    expect(last.output.type).toBe('error');
  });
});

describe('execute — theme', () => {
  beforeEach(() => terminal.reset());

  it('sets valid theme', () => {
    terminal.execute('theme gruvbox-hard');
    expect(get(terminal).theme).toBe('gruvbox-hard');
  });
  it('rejects invalid theme', () => {
    terminal.execute('theme solarized');
    const last = get(terminal).history.at(-1)!;
    expect(last.output.type).toBe('error');
  });
});

describe('execute — clear', () => {
  beforeEach(() => terminal.reset());

  it('wipes history', () => {
    terminal.execute('cd resume');
    terminal.execute('clear');
    expect(get(terminal).history.length).toBe(0);
  });
});

describe('hints sync', () => {
  beforeEach(() => terminal.reset());

  it('updates hints after cd', () => {
    terminal.execute('cd resume');
    const hints = get(terminal).hints;
    expect(hints.some(h => h.cmd === 'ls')).toBe(true);
  });
});

describe('tabComplete', () => {
  beforeEach(() => terminal.reset());

  it('completes verb', () => {
    const completed = terminal.tabComplete('cd res');
    expect(completed).toBe('cd resume/');
  });
  it('stores candidates on multiple matches', () => {
    terminal.tabComplete('c');
    const candidates = get(terminal).completionCandidates;
    expect(candidates.length).toBeGreaterThan(1);
  });
});
