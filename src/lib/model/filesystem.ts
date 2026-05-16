import type { FsDirectory, FsFile, FsNode } from './types';
import { profile, contact, blogInfo, experiences, education, projects, toolkit } from './data';

function makeDir(name: string, children: [string, FsNode][]): FsDirectory {
  return { kind: 'dir', name, children: new Map(children) };
}

function makeFile(name: string, data: unknown): FsFile {
  return { kind: 'file', name, data };
}

/**
 * The root of the virtual filesystem.
 * Built from portfolio.json.
 */
export const fs: FsDirectory = makeDir('~', [
  ['about',    makeFile('about',    profile)],
  ['contact',  makeFile('contact',  contact)],
  ['blog',     makeFile('blog',     blogInfo)],
  ['resume',   makeDir('resume', [
    ['experience', makeDir('experience', experiences.map(e => [e.id, makeFile(e.id, e)]))],
    ['education',  makeDir('education',  education.map(e => [e.id, makeFile(e.id, e)]))],
    ['projects',   makeDir('projects',   projects.map(p => [p.id, makeFile(p.id, p)]))],
    ['toolkit',    makeDir('toolkit',    toolkit.map(g => [g.id, makeFile(g.id, g)]))],
  ])],
]);

/**
 * Traverse the filesystem tree by absolute path segments.
 * Returns the node at that path, or null if not found.
 *
 * @param path - absolute path string e.g. '~/resume/experience'
 */
export function getNode(path: string): FsNode | null {
  const segments = path === '~' ? [] : path.replace('~/', '').split('/');
  let current: FsNode = fs;

  for (const seg of segments) {
    if (current.kind !== 'dir') return null;
    const child = current.children.get(seg);
    if (!child) return null;
    current = child;
  }

  return current;
}

/**
 * List the immediate children of a directory node.
 * Returns [] if the node is a file or doesn't exist.
 */
export function listChildren(path: string): string[] {
  const node = getNode(path);
  if (!node || node.kind !== 'dir') return [];
  return Array.from(node.children.keys());
}
