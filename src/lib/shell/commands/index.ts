import type { OutputEntry } from '$lib/model/types';
import { cd, type CdResult } from './cd';
import { ls }   from './ls';
import { cat }  from './cat';
import { help } from './help';

export interface DispatchResult {
  output: OutputEntry;
  newPath?: string;
}

/**
 * Parse a raw input string into verb + args.
 * Handles quoted strings:  cat "my file"  →  ['cat', 'my file']
 */
export function parse(raw: string): { verb: string; args: string[] } {
  const tokens: string[] = [];
  let current = '';
  let inQuote = false;

  for (const char of raw.trim()) {
    if (char === '"' || char === "'") { inQuote = !inQuote; continue; }
    if (char === ' ' && !inQuote)     { if (current) { tokens.push(current); current = ''; } continue; }
    current += char;
  }
  if (current) tokens.push(current);

  const [verb = '', ...args] = tokens;
  return { verb: verb.toLowerCase(), args };
}

/**
 * Dispatch a parsed command to the correct handler.
 */
export function dispatch(cwd: string, raw: string): DispatchResult {
  const { verb, args } = parse(raw);

  switch (verb) {
    case 'cd':    return cd(cwd, args);
    case 'ls':    return { output: ls(cwd, args) };
    case 'cat':   return { output: cat(cwd, args) };
    case 'help':  return { output: help() };
    case 'clear': return { output: { type: 'text', payload: '__clear__' } };
    case 'theme': {
      const valid: string[] = [
        'gruvbox-dark', 'gruvbox-hard', 'gruvbox-soft',
        'one-dark', 'dracula', 'nord', 'monokai',
        'solarized-dark', 'tokyo-night',
        'cyberpunk', 'synthwave', 'matrix',
        'hotdogstand', 'campfire', 'vaporwave',
      ];
      if (!args[0] || !valid.includes(args[0])) {
        return {
          output: {
            type: 'text',
            payload: [
              'usage: theme <name>',
              '',
              `unrecognized: "${args[0] ?? '(none)'}"`,
              '',
              'available themes:',
              ...valid.map(v => `  ${v}`),
            ].join('\n'),
          },
        };
      }
      return { output: { type: 'success', payload: `theme set to ${args[0]}` } };
    }
    case '': return { output: { type: 'text', payload: '' } };
    default: {
      return {
        output: {
          type: 'error',
          payload: `command not found: ${verb}. type help for available commands.`,
        },
      };
    }
  }
}
