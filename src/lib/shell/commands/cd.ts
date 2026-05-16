import type { OutputEntry } from '$lib/model/types';
import { resolveAndValidate } from '$lib/shell/resolver';
import { getNode } from '$lib/model/filesystem';

export interface CdResult {
  output: OutputEntry;
  newPath?: string;
}

export function cd(cwd: string, args: string[]): CdResult {
  if (args.length === 0) {
    return { output: { type: 'success', payload: 'moved to ~' }, newPath: '~' };
  }

  const target = args[0];
  const result = resolveAndValidate(cwd, target);

  if (!result.ok) {
    return { output: { type: 'error', payload: result.error } };
  }

  const node = getNode(result.path);

  if (node?.kind === 'file') {
    return {
      output: { type: 'error', payload: `not a directory: ${target}. use cat to read files.` },
    };
  }

  return {
    output: { type: 'success', payload: `${result.path}` },
    newPath: result.path,
  };
}
