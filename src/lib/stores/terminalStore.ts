import { writable, get } from 'svelte/store';
import type { TerminalState, HistoryEntry, Hint } from '$lib/model/types';
import { dispatch }    from '$lib/shell/commands/index';
import { complete, deriveHints } from '$lib/shell/completer';

import { profile } from '$lib/model/data';

const BOOT: HistoryEntry = {
  id:      'boot',
  command: '',
  path:    '~',
  output:  { type: 'text', payload: `${profile.name} — type help or press Tab to explore` },
};

function hintsFromStrings(cmds: string[]): Hint[] {
  return cmds.map(cmd => ({ label: cmd, cmd }));
}

function createTerminalStore() {
  const initial: TerminalState = {
    path:                '~',
    history:             [BOOT],
    hints:               hintsFromStrings(deriveHints('~')),
    theme:               'gruvbox-dark',
    cmdHistoryLog:       [],
    cmdHistoryCursor:    -1,
    completionCandidates: [],
  };

  const { subscribe, update } = writable<TerminalState>(initial);

  // ── execute ───────────────────────────────────────────────────

  function execute(raw: string) {
    const trimmed = raw.trim();
    if (!trimmed) return;

    update(state => {
      const result = dispatch(state.path, trimmed);

      // clear is a special signal
      if (result.output.payload === '__clear__') {
        return {
          ...state,
          history:             [],
          cmdHistoryLog:       [],
          cmdHistoryCursor:    -1,
          completionCandidates: [],
        };
      }

      // theme change: extract new theme from the command
      const { verb, args } = (() => {
        const p = trimmed.split(/\s+/);
        return { verb: p[0], args: p.slice(1) };
      })();

      const nextTheme = verb === 'theme' && result.output.type === 'success'
        ? (args[0] as TerminalState['theme'])
        : state.theme;

      const nextPath = result.newPath ?? state.path;

      const entry: HistoryEntry = {
        id:      crypto.randomUUID(),
        command: trimmed,
        path:    state.path,
        output:  result.output,
      };

      return {
        ...state,
        path:                nextPath,
        theme:               nextTheme,
        history:             [...state.history, entry],
        hints:               hintsFromStrings(deriveHints(nextPath)),
        cmdHistoryLog:       [trimmed, ...state.cmdHistoryLog].slice(0, 50),
        cmdHistoryCursor:    -1,
        completionCandidates: [],
      };
    });
  }

  // ── tabComplete ───────────────────────────────────────────────

  function tabComplete(input: string): string {
    const state  = get({ subscribe });
    const result = complete(state.path, input);

    update(s => ({ ...s, completionCandidates: result.candidates }));

    return result.completed;
  }

  // ── history navigation ────────────────────────────────────────

  function historyUp(): string {
    let val = '';
    update(state => {
      const next = Math.min(state.cmdHistoryCursor + 1, state.cmdHistoryLog.length - 1);
      val = state.cmdHistoryLog[next] ?? '';
      return { ...state, cmdHistoryCursor: next };
    });
    return val;
  }

  function historyDown(): string {
    let val = '';
    update(state => {
      const next = Math.max(state.cmdHistoryCursor - 1, -1);
      val = next === -1 ? '' : state.cmdHistoryLog[next] ?? '';
      return { ...state, cmdHistoryCursor: next };
    });
    return val;
  }

  function setTheme(theme: TerminalState['theme']) {
    update(s => ({ ...s, theme }));
  }

  function reset() {
    update(() => initial);
  }

  return { subscribe, execute, tabComplete, historyUp, historyDown, setTheme, reset };
}

export const terminal = createTerminalStore();
