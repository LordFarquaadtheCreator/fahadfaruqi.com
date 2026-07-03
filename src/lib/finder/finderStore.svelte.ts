// VM types matching Go finder/viewmodel.go
export interface VM {
	windows: WindowVM[];
	focusedId: string;
	viewportW: number;
	viewportH: number;
}

export interface WindowVM {
	id: string;
	title: string;
	x: number;
	y: number;
	z: number;
	width: number;
	height: number;
	viewMode: string;
	selectedEntry: string;
	canBack: boolean;
	canForward: boolean;
	sidebar: SidebarItemVM[];
	content: ContentVM;
}

export interface SidebarItemVM {
	name: string;
	path: string;
	kind: string;
	level: number;
	active: boolean;
}

export interface ContentVM {
	kind: 'folder' | 'markdown';
	entries?: EntryVM[];
	html?: string;
}

export interface EntryVM {
	name: string;
	path: string;
	kind: 'dir' | 'md';
	size: number;
	modified: string;
}

// global window.os bindings
export interface OsAPI {
	subscribe(cb: (json: string) => void): () => void;
	setViewport(w: number, h: number): void;
	spawn(path: string): string | null;
	close(id: string): void;
	closeFocused(): void;
	focus(id: string): void;
	drag(id: string, dx: number, dy: number): void;
	navigate(id: string, path: string): string | null;
	back(id: string): void;
	forward(id: string): void;
	setViewMode(id: string, mode: string): void;
	selectEntry(id: string, name: string): void;
	openEntry(id: string, path: string): string | null;
}

declare global {
	interface Window {
		os: OsAPI;
	}
}

// Reactive VM snapshot from Go
export const finderState = $state<{ vm: VM | null }>({ vm: null });

export function initFinder(): (() => void) | void {
	if (typeof window === 'undefined') return;

	let unsub: (() => void) | null = null;
	let interval: ReturnType<typeof setInterval> | null = null;

	const setup = () => {
		if (!window.os) return;
		if (interval) {
			clearInterval(interval);
			interval = null;
		}

		unsub = window.os.subscribe((json: string) => {
			finderState.vm = JSON.parse(json) as VM;
		});

		window.os.setViewport(window.innerWidth, window.innerHeight);
	};

	if (window.os) {
		setup();
	} else {
		interval = setInterval(setup, 50);
	}

	const onResize = () => {
		if (window.os) window.os.setViewport(window.innerWidth, window.innerHeight);
	};
	window.addEventListener('resize', onResize);

	const onKey = (e: KeyboardEvent) => {
		if (e.key === 'Escape' && window.os) {
			window.os.closeFocused();
		}
	};
	window.addEventListener('keydown', onKey);

	return () => {
		if (interval) clearInterval(interval);
		if (unsub) unsub();
		window.removeEventListener('resize', onResize);
		window.removeEventListener('keydown', onKey);
	};
}
