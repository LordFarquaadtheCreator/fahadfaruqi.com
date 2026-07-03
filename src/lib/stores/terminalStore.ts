import { writable, get } from 'svelte/store';
import type { TerminalState, HistoryEntry, Hint } from '$lib/model/types';
import { loadInterpreter } from '$lib/wasmLoader.js';

import { profile } from '$lib/model/data';
import { complete } from '$lib/shell/completer';

const BOOT: HistoryEntry = {
  id:      'boot',
  command: '',
  path:    '~',
  output:  { type: 'text', payload: `${profile.name} — type help or press Tab to explore` },
};

function hintsFromStrings(cmds: string[]): Hint[] {
  return cmds.map(cmd => ({ label: cmd, cmd }));
}

let run: ((input: string) => string) | null = null;
let initPromise: Promise<void> | null = null;
const cmdQueue: string[] = [];

function createTerminalStore() {
  const initial: TerminalState = {
    path:                '~',
    history:             [BOOT],
    hints:               hintsFromStrings(['ls', 'cd resume', 'cat readme.txt', 'help', 'clear', 'theme']),
    theme:               'gruvbox-dark',
    cmdHistoryLog:       [],
    cmdHistoryCursor:    -1,
    completionCandidates: [],
  };

  const { subscribe, update, set } = writable<TerminalState>(initial);

  // Initialize WASM interpreter
  async function init() {
    console.log('terminalStore init called');
    if (initPromise) return initPromise;
    initPromise = (async () => {
      console.log('terminalStore loading WASM...');
      run = await loadInterpreter();
      console.log('terminalStore WASM loaded, run set:', run !== null);
      // Process queued commands
      while (cmdQueue.length > 0) {
        const cmd = cmdQueue.shift()!;
        processCommand(cmd);
      }
    })();
    return initPromise;
  }

  function processCommand(trimmed: string) {
    if (!run) return;

    update(state => {
      const resultStr = run!(trimmed);
      const result = JSON.parse(resultStr);

      // clear is a special signal
      if (result.type === 'clear') {
        return {
          ...state,
          history:             [],
          cmdHistoryLog:       [],
          cmdHistoryCursor:    -1,
          completionCandidates: [],
        };
      }

      // theme change
      const { verb, args } = (() => {
        const p = trimmed.split(/\s+/);
        return { verb: p[0], args: p.slice(1) };
      })();

      const nextTheme = verb === 'theme' && result.type === 'success'
        ? (args[0] as TerminalState['theme'])
        : state.theme;

      const nextPath = result.cwd ?? state.path;

      const entry: HistoryEntry = {
        id:      crypto.randomUUID(),
        command: trimmed,
        path:    state.path,
        output:  { type: result.type as any, payload: result.payload },
      };

      return {
        ...state,
        path:                nextPath,
        theme:               nextTheme,
        history:             [...state.history, entry],
        hints:               hintsFromStrings(['ls', 'cd resume', 'cat readme.txt', 'help', 'clear', 'theme']),
        cmdHistoryLog:       [trimmed, ...state.cmdHistoryLog].slice(0, 50),
        cmdHistoryCursor:    -1,
        completionCandidates: [],
      };
    });
  }

  // ── execute ───────────────────────────────────────────────────

  function execute(raw: string) {
    const trimmed = raw.trim();
    console.log('execute called:', trimmed, 'run:', run !== null);
    if (!trimmed) return;

    if (!run) {
      // Queue command and trigger init
      console.log('queuing command');
      cmdQueue.push(trimmed);
      init();
      return;
    }

    console.log('processing command');
    processCommand(trimmed);
  }

  // ── tabComplete ───────────────────────────────────────────────

  function tabComplete(input: string): string {
    const state = get({ subscribe });
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
    set(initial);
  }

  return { subscribe, execute, tabComplete, historyUp, historyDown, setTheme, reset, init };
}

export const terminal = createTerminalStore();
