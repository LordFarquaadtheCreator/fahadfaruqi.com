import type { OutputEntry } from '$lib/model/types';
import { getNode } from '$lib/model/filesystem';
import { formatEntry } from '$lib/model/data';
import { resolve } from '$lib/shell/resolver';

export function cat(cwd: string, args: string[]): OutputEntry {
  if (args.length === 0) {
    return { type: 'error', payload: 'usage: cat <file>' };
  }

  const result  = resolve(cwd, args[0]);
  if (!result.ok) return { type: 'error', payload: result.error };

  const node = getNode(result.path);

  if (!node) {
    return { type: 'error', payload: `no such file: ${args[0]}` };
  }

  if (node.kind === 'dir') {
    return { type: 'error', payload: `${args[0]} is a directory. use ls to list its contents.` };
  }

  return { type: 'text', payload: formatEntry(node.data) };
}
