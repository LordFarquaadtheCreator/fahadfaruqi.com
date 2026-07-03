// VM types matching Go RootVM schema. Svelte renders these, never mutates.

export interface RootVM {
  windows: WindowVM[];
  focusedId: string;
  viewportW: number;
  viewportH: number;
  dock: DockVM;
  menuBar: MenuBarVM;
  desktop: DesktopVM;
}

export interface WindowVM {
  id: string;
  appType: string;
  appId: string;
  title: string;
  x: number;
  y: number;
  z: number;
  width: number;
  height: number;
  minimized: boolean;
  maximized: boolean;
  ephemeral: boolean;
  content: FinderContentVM | TerminalContentVM | ViewerContentVM | EphemeralContentVM | SplashContentVM;
}

export interface FinderContentVM {
  kind: 'finder';
  viewMode: string;
  selectedEntry: string;
  canBack: boolean;
  canForward: boolean;
  sidebar: SidebarSectionVM[];
  breadcrumb: BreadcrumbItemVM[];
  content: FinderContent;
}

export interface SidebarSectionVM {
  title: string;
  items: SidebarItemVM[];
}

export interface SidebarItemVM {
  name: string;
  path: string;
  icon: string;
  active: boolean;
}

export interface FinderContent {
  kind: string;
  entries?: EntryVM[];
  html?: string;
}

export interface EntryVM {
  name: string;
  path: string;
  kind: string;
  icon: string;
  size: number;
  modified: string;
}

export interface BreadcrumbItemVM {
  name: string;
  path: string;
  icon: string;
}

export interface TerminalContentVM {
  kind: 'terminal';
  cwd: string;
  history: TerminalEntry[];
}

export interface TerminalEntry {
  input: string;
  output: string;
  cwd: string;
}

export interface ViewerContentVM {
  kind: 'viewer-image' | 'viewer-text';
  src?: string;
  name: string;
  content?: string;
  isMarkdown?: boolean;
}

export interface EphemeralContentVM {
  kind: 'ephemeral-image' | 'ephemeral-folder';
  src?: string;
  name: string;
  filePath: string;
  icon?: string;
}

export interface SplashContentVM {
  kind: 'splash';
  appId: string;
  name: string;
  icon: string;
}

export interface DockVM {
  items: DockItemVM[];
  trash: DockItemVM;
}

export interface DockItemVM {
  appId: string;
  name: string;
  icon: string;
  running: boolean;
  active: boolean;
}

export interface MenuBarVM {
  activeApp: string;
  menus: MenuVM[];
  clock: string;
  tray: TrayItemVM[];
}

export interface MenuVM {
  name: string;
  items: MenuItemVM[];
}

export interface MenuItemVM {
  label: string;
  shortcut?: string;
  separator?: boolean;
}

export interface TrayItemVM {
  icon: string;
  name: string;
}

export interface DesktopVM {
  icons: DesktopIconVM[];
  wallpaper: string;
}

export interface DesktopIconVM {
  name: string;
  path: string;
  icon: string;
  x: number;
  y: number;
}

// Shell response types (from interpRun / terminalRun)
export interface ShellResponse {
  type: string;
  payload: unknown;
  cwd: string;
}

export interface LsPayload {
  entries: FSEntry[];
  longFormat?: boolean;
  'human-readable'?: boolean;
  'show-hidden'?: boolean;
  recursive?: boolean;
  color?: string;
}

export interface FSEntry {
  name: string;
  type: string;
  size: number;
  modified: string;
  hidden: boolean;
  permission?: string;
  owner?: string;
  group?: string;
}

export interface CatPayload {
  name: string;
  content: string;
}

export interface TextPayload {
  text: string;
}

export interface EnvPayload {
  vars: Record<string, string>;
}

export interface HistoryPayload {
  entries: string[];
}

export interface ErrorPayload {
  command: string;
  message: string;
}
