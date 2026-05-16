export type Theme =
  | 'gruvbox-dark' | 'gruvbox-hard' | 'gruvbox-soft'
  | 'one-dark'
  | 'dracula'
  | 'nord'
  | 'monokai'
  | 'solarized-dark'
  | 'tokyo-night'
  | 'cyberpunk'
  | 'synthwave'
  | 'matrix'
  | 'hotdogstand'
  | 'campfire'
  | 'vaporwave';

export type OutputType = 'text' | 'success' | 'error' | 'warning' | 'json' | 'list';

export interface OutputEntry {
  type: OutputType;
  payload: string | JsonNode | ListItem[] | ListItem;
}

export interface HistoryEntry {
  id: string;
  command: string;
  path: string;
  output: OutputEntry;
}

export interface Hint {
  label: string;
  cmd: string;
}

export interface TerminalState {
  path: string;
  history: HistoryEntry[];
  hints: Hint[];
  theme: Theme;
  cmdHistoryLog: string[];
  cmdHistoryCursor: number;
  completionCandidates: string[];
}

export interface JsonNode {
  [key: string]: string | number | boolean | null | JsonNode | JsonNode[];
}

export interface ListItem {
  id: string;
  primary: string;
  secondary?: string;
  meta?: string;
}

// ── Virtual filesystem ────────────────────────────────────────

export interface FsFile {
  kind: 'file';
  name: string;
  data: unknown;
}

export interface FsDirectory {
  kind: 'dir';
  name: string;
  children: Map<string, FsNode>;
}

export type FsNode = FsDirectory | FsFile;
