import type { OutputEntry, ListItem } from '$lib/model/types';
import { getNode, listChildren } from '$lib/model/filesystem';
import { toListItem } from '$lib/model/data';
import { resolve } from '$lib/shell/resolver';
import { match } from '$lib/shell/glob';

export function ls(cwd: string, args: string[]): OutputEntry {
  const targetPath = args.length > 0
    ? (() => { const r = resolve(cwd, args[0]); return r.ok ? r.path : null; })()
    : cwd;

  if (!targetPath) {
    return { type: 'error', payload: `no such directory: ${args[0]}` };
  }

  const node = getNode(targetPath);

  if (!node) {
    const pattern  = args[0] ?? '*';
    const children = listChildren(cwd);
    const matched  = match(pattern, children);

    if (matched.length === 0) {
      return { type: 'error', payload: `no match: ${pattern}` };
    }

    const items = matched.map(name => {
      const child = getNode(cwd === '~' ? `~/${name}` : `${cwd}/${name}`);
      return child?.kind === 'file'
        ? toListItem(child.data)
        : { id: name, primary: name + '/' } as ListItem;
    });

    return { type: 'list', payload: items };
  }

  if (node.kind === 'file') {
    return { type: 'list', payload: [toListItem(node.data)] };
  }

  const children = Array.from(node.children.values());

  if (children.length === 0) {
    return { type: 'text', payload: '(empty)' };
  }

  const items: ListItem[] = children.map(child =>
    child.kind === 'file'
      ? toListItem(child.data)
      : { id: child.name, primary: child.name + '/', secondary: `${child.children.size} items` }
  );

  return { type: 'list', payload: items };
}
