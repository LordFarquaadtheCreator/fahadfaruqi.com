import { getNode } from '$lib/model/filesystem';

export type ResolveResult =
  | { ok: true;  path: string }
  | { ok: false; error: string };

/**
 * Resolve an input path against the current working directory.
 * Supports:
 *   - absolute paths  ~/resume/experience
 *   - relative paths  experience  or  ../blog
 *   - dot-dot         ..  or  ../../
 *   - tilde           ~
 *
 * Does NOT touch the filesystem — pure string math.
 * Call validate() after to confirm the resolved path exists.
 */
export function resolve(cwd: string, input: string): ResolveResult {
  if (!input || input === '.') return { ok: true, path: cwd };
  if (input === '~')           return { ok: true, path: '~' };

  // Build segment array from base
  const base = input.startsWith('~')
    ? []
    : (cwd === '~' ? [] : cwd.replace('~/', '').split('/'));

  const inputSegs = input.startsWith('~/')
    ? input.replace('~/', '').split('/')
    : input.split('/');

  const stack = [...base];

  for (const seg of inputSegs) {
    if (seg === '' || seg === '.') continue;
    if (seg === '..') {
      if (stack.length > 0) stack.pop();
      // already at root — silently stay
    } else {
      stack.push(seg);
    }
  }

  const resolved = stack.length === 0 ? '~' : `~/${stack.join('/')}`;
  return { ok: true, path: resolved };
}

/**
 * Resolve and validate — confirms the path exists in the filesystem.
 */
export function resolveAndValidate(cwd: string, input: string): ResolveResult {
  const result = resolve(cwd, input);
  if (!result.ok) return result;

  const node = getNode(result.path);
  if (!node) return { ok: false, error: `no such file or directory: ${input}` };

  return { ok: true, path: result.path };
}

/**
 * Return the parent of an absolute path.
 */
export function parent(path: string): string {
  if (path === '~') return '~';
  const segs = path.replace('~/', '').split('/');
  segs.pop();
  return segs.length === 0 ? '~' : `~/${segs.join('/')}`;
}
