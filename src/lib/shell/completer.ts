import { listChildren, getNode } from '$lib/model/filesystem';
import { resolve } from './resolver';
import { match, longestCommonPrefix } from './glob';

export interface CompletionResult {
  completed: string;        // the completed input string to put back in the input box
  candidates: string[];     // all matches (shown if > 1)
}

/**
 * Complete a partial command input.
 *
 * Handles two cases:
 * 1. Completing a verb (first token): match against known commands.
 * 2. Completing a path argument (second+ token): match against fs children.
 *
 * @param cwd     - current working directory
 * @param input   - the full current input string (what's in the text box)
 */
export function complete(cwd: string, input: string): CompletionResult {
  const tokens = input.split(/\s+/);
  const COMMANDS = ['cd', 'ls', 'cat', 'help', 'clear', 'theme'];

  // Completing the verb
  if (tokens.length === 1) {
    const partial = tokens[0];
    const candidates = match(partial, COMMANDS);
    const completed = candidates.length === 1
      ? candidates[0] + ' '
      : input.slice(0, -partial.length) + longestCommonPrefix(candidates);
    return { completed, candidates };
  }

  // Completing a path argument
  const verb    = tokens[0];
  const partial = tokens[tokens.length - 1];  // the token being typed

  // Split partial into directory part and filename part
  // e.g. 'resume/exp' → dir='~/resume', file='exp'
  // e.g. 'exp'        → dir=cwd,         file='exp'
  let dirPart: string;
  let filePart: string;

  if (partial.includes('/')) {
    const lastSlash = partial.lastIndexOf('/');
    const dirInput  = partial.slice(0, lastSlash) || '~';
    filePart        = partial.slice(lastSlash + 1);
    const resolved  = resolve(cwd, dirInput);
    dirPart         = resolved.ok ? resolved.path : cwd;
  } else {
    dirPart  = cwd;
    filePart = partial;
  }

  const children   = listChildren(dirPart);
  const candidates = match(filePart, children);

  if (candidates.length === 0) {
    return { completed: input, candidates: [] };
  }

  const fill = candidates.length === 1
    ? candidates[0]
    : longestCommonPrefix(candidates);

  // Rebuild the completed input string
  const prefix   = tokens.slice(0, -1).join(' ');
  const dirPrefix = partial.includes('/')
    ? partial.slice(0, partial.lastIndexOf('/') + 1)
    : '';

  // Append trailing / if completing a directory
  const resolvedPath = resolve(dirPart, fill);
  const node = resolvedPath.ok ? getNode(resolvedPath.path) : null;
  const suffix = node?.kind === 'dir' ? '/' : ' ';

  const completed = `${prefix} ${dirPrefix}${fill}${candidates.length === 1 ? suffix : ''}`.trimStart();
  return { completed, candidates };
}

/**
 * Derive hint pills from the current working directory.
 * Called after every cd/execute to keep hints in sync.
 */
export function deriveHints(cwd: string): string[] {
  const children = listChildren(cwd);
  const ALWAYS   = ['help', 'clear'];

  const childHints = children.map(c => {
    const node = getNode(cwd === '~' ? `~/${c}` : `${cwd}/${c}`);
    return node?.kind === 'dir' ? `cd ${c}` : `cat ${c}`;
  });

  // ls only if there are children to list
  const lsHint = children.length > 0 ? ['ls'] : [];

  // cd .. unless we're at root
  const upHint = cwd !== '~' ? ['cd ..'] : [];

  return [...childHints, ...lsHint, ...upHint, ...ALWAYS];
}
