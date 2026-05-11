export type ThemeChoice = 'light' | 'dark' | 'system';
export type ResolvedTheme = 'light' | 'dark';

const STORAGE_KEY = 'theme';

function readChoice(): ThemeChoice {
	if (typeof localStorage === 'undefined') return 'system';
	const v = localStorage.getItem(STORAGE_KEY);
	return v === 'light' || v === 'dark' ? v : 'system';
}

function systemPref(): ResolvedTheme {
	if (typeof window === 'undefined') return 'light';
	return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
}

function resolve(choice: ThemeChoice): ResolvedTheme {
	return choice === 'system' ? systemPref() : choice;
}

function apply(resolved: ResolvedTheme) {
	if (typeof document === 'undefined') return;
	document.documentElement.setAttribute('data-theme', resolved);
}

function persist(choice: ThemeChoice) {
	if (typeof localStorage === 'undefined') return;
	if (choice === 'system') localStorage.removeItem(STORAGE_KEY);
	else localStorage.setItem(STORAGE_KEY, choice);
}

class ThemeStore {
	choice = $state<ThemeChoice>('system');
	resolved = $state<ResolvedTheme>('light');

	init() {
		this.choice = readChoice();
		this.resolved = resolve(this.choice);
		apply(this.resolved);

		const mq = window.matchMedia('(prefers-color-scheme: dark)');
		const onChange = () => {
			if (this.choice === 'system') {
				this.resolved = systemPref();
				apply(this.resolved);
			}
		};
		mq.addEventListener('change', onChange);
		return () => mq.removeEventListener('change', onChange);
	}

	set(choice: ThemeChoice) {
		this.choice = choice;
		this.resolved = resolve(choice);
		apply(this.resolved);
		persist(choice);
	}

	cycle() {
		const next: ThemeChoice =
			this.choice === 'system' ? 'light' : this.choice === 'light' ? 'dark' : 'system';
		this.set(next);
	}
}

export const theme = new ThemeStore();
