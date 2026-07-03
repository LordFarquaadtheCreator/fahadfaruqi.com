import { writable } from 'svelte/store';

interface GoResponse {
	type: string;
	payload: any;
	cwd: string;
}

interface OutputEntry {
	type: 'text' | 'list' | 'json' | 'error' | 'success' | 'warning';
	payload: any;
}

interface HistoryEntry {
	id: string;
	command: string;
	path: string;
	output: OutputEntry;
}

let run: ((input: string) => string) | null = null;

const history = writable<HistoryEntry[]>([]);
const cwd = writable<string>('');
const ready = writable<boolean>(false);

export { history, cwd, ready };

function generateId(): string {
	return Date.now().toString(36) + Math.random().toString(36).substring(2);
}

function convertGoResponseToSvelte(response: GoResponse): HistoryEntry {
	const output: OutputEntry = {
		type: response.type as any,
		payload: response.payload
	};

	return {
		id: generateId(),
		command: '',
		path: response.cwd,
		output
	};
}

export async function init() {
	const { loadInterpreter } = await import('./wasmLoader.js');
	run = await loadInterpreter();
	
	const bootResult = run!('pwd');
	const bootResponse = JSON.parse(bootResult) as GoResponse;
	
	const bootEntry: HistoryEntry = {
		id: generateId(),
		command: '',
		path: bootResponse.cwd,
		output: {
			type: 'text',
			payload: 'Fahad Faruqi — type help or press Tab to explore'
		}
	};
	history.set([bootEntry]);
	cwd.set(bootResponse.cwd);
	ready.set(true);
}

export function dispatch(input: string) {
	if (!run) return;

	const raw = run(input);
	const response = JSON.parse(raw) as GoResponse;

	if (response.type === 'clear') {
		history.set([]);
	} else {
		const entry = convertGoResponseToSvelte(response);
		entry.command = input;
		history.update(h => [...h, entry]);
	}

	cwd.set(response.cwd);
}
