import { describe, it, expect, beforeEach, vi } from 'vitest';
import { theme } from './theme.svelte';

const mockLocalStorage = {
	getItem: vi.fn(),
	setItem: vi.fn(),
	removeItem: vi.fn(),
	clear: vi.fn()
};

const mockDocumentElement = {
	getAttribute: vi.fn(),
	setAttribute: vi.fn(),
	removeAttribute: vi.fn()
};

describe('theme store', () => {
	beforeEach(() => {
		mockLocalStorage.getItem.mockReturnValue(null);
		mockLocalStorage.setItem.mockClear();
		mockLocalStorage.removeItem.mockClear();
		mockLocalStorage.clear.mockClear();
		mockDocumentElement.getAttribute.mockReturnValue(null);
		mockDocumentElement.setAttribute.mockClear();
		mockDocumentElement.removeAttribute.mockClear();
		vi.stubGlobal('localStorage', mockLocalStorage);
		vi.stubGlobal('document', { documentElement: mockDocumentElement });
		vi.stubGlobal('window', {
			matchMedia: (q: string) => ({
				matches: q.includes('dark') ? false : false,
				media: q,
				addEventListener: vi.fn(),
				removeEventListener: vi.fn(),
				addListener: vi.fn(),
				removeListener: vi.fn(),
				dispatchEvent: vi.fn(),
				onchange: null
			})
		});

		theme.choice = 'system';
		theme.resolved = 'light';
	});

	it('defaults to system + light when no storage and OS light', () => {
		theme.init();
		expect(theme.choice).toBe('system');
		expect(theme.resolved).toBe('light');
		expect(mockDocumentElement.setAttribute).toHaveBeenCalledWith('data-theme', 'light');
	});

	it('reads stored dark', () => {
		mockLocalStorage.getItem.mockReturnValue('dark');
		theme.init();
		expect(theme.choice).toBe('dark');
		expect(theme.resolved).toBe('dark');
	});

	it('cycle: system → light → dark → system', () => {
		theme.init();
		theme.cycle();
		expect(theme.choice).toBe('light');
		theme.cycle();
		expect(theme.choice).toBe('dark');
		expect(mockLocalStorage.setItem).toHaveBeenCalledWith('theme', 'dark');
		theme.cycle();
		expect(theme.choice).toBe('system');
		expect(mockLocalStorage.removeItem).toHaveBeenCalledWith('theme');
	});

	it('set persists and applies', () => {
		theme.set('dark');
		expect(document.documentElement.setAttribute).toHaveBeenCalledWith('data-theme', 'dark');
		expect(mockLocalStorage.setItem).toHaveBeenCalledWith('theme', 'dark');
	});
});
