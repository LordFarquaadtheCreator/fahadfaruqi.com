import { describe, expect, it } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import Page from './+page.svelte';

describe('gothic landing page', () => {
	it('renders the four navigation panels', () => {
		render(Page);
		expect(screen.getByRole('link', { name: /ABOUT_ME/i })).toBeTruthy();
		expect(screen.getByRole('link', { name: /PROJECTS/i })).toBeTruthy();
		expect(screen.getByRole('link', { name: /RESUME/i })).toBeTruthy();
		expect(screen.getByRole('link', { name: /CONTACT/i })).toBeTruthy();
	});

	it('renders the gothic title', () => {
		render(Page);
		expect(document.body.textContent).toContain('Fahad');
		expect(document.body.textContent).toContain('Faruqi');
	});
});
