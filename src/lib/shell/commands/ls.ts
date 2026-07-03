import type { OutputEntry } from '$lib/model/types';
import { getNode, listChildren } from '$lib/model/filesystem';
import { resolve } from '$lib/shell/resolver';
import { match } from '$lib/shell/glob';

/**
 * Format filenames into columns like real `ls` output.
 * Tries to fit as many columns as possible within target width.
 */
function formatColumns(names: string[], termWidth = 80): string {
  if (names.length === 0) return '';
  if (names.length === 1) return names[0];

  const maxLen = Math.max(...names.map(n => n.length));
  const colWidth = maxLen + 2; // min gap between columns

  const cols = Math.max(1, Math.floor(termWidth / colWidth));
  const rows = Math.ceil(names.length / cols);

  const lines: string[] = [];
  for (let r = 0; r < rows; r++) {
    const row: string[] = [];
    for (let c = 0; c < cols; c++) {
      const idx = r + c * rows;
      if (idx < names.length) {
        const name = names[idx];
        const padding = ' '.repeat(colWidth - name.length);
        row.push(name + padding);
      }
    }
    lines.push(row.join('').trimEnd());
  }

  return lines.join('\n');
}

export function ls(cwd: string, args: string[]): OutputEntry {
  const longFormat = args.includes('-l') || args.includes('-al') || args.includes('-la');
  const plainArgs = args.filter(a => !a.startsWith('-'));

  const targetPath = plainArgs.length > 0
    ? (() => { const r = resolve(cwd, plainArgs[0]); return r.ok ? r.path : null; })()
    : cwd;

  if (!targetPath) {
    return { type: 'error', payload: `no such directory: ${plainArgs[0]}` };
  }

  const node = getNode(targetPath);

  if (!node) {
    const pattern  = plainArgs[0] ?? '*';
    const children = listChildren(cwd);
    const matched  = match(pattern, children);

    if (matched.length === 0) {
      return { type: 'error', payload: `no match: ${pattern}` };
    }

    const names = matched.map(name => {
      const child = getNode(cwd === '~' ? `~/${name}` : `${cwd}/${name}`);
      return child?.kind === 'dir' ? name + '/' : name;
    });

    return { type: 'text', payload: formatColumns(names) };
  }

  if (node.kind === 'file') {
    return { type: 'text', payload: node.name };
  }

  const children = Array.from(node.children.values());

  if (children.length === 0) {
    return { type: 'text', payload: '' };
  }

  const names = children.map(child =>
    child.kind === 'dir' ? child.name + '/' : child.name
  );

  if (longFormat) {
    // Simple ls -l style with placeholder permissions
    const total = children.length;
    const lines = [`total ${total}`];
    for (const child of children) {
      const perms = child.kind === 'dir' ? 'drwxr-xr-x' : '-rw-r--r--';
      const size = child.kind === 'dir' ? '96' : '1024';
      const date = 'May 21 15:06';
      const name = child.kind === 'dir' ? child.name + '/' : child.name;
      lines.push(`${perms}  1 user  user  ${size.padStart(5)} ${date} ${name}`);
    }
    return { type: 'text', payload: lines.join('\n') };
  }

  return { type: 'text', payload: formatColumns(names) };
}
