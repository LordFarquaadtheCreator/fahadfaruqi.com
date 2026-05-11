import { describe, expect, it } from 'vitest';
import { fireEvent, render, screen } from '@testing-library/svelte';
import portfolio from '$lib/data/portfolio';
import Page from './+page.svelte';

describe('portfolio page', () => {
	it('renders the PRD gateway slides from the JSON data source', async () => {
		render(Page);

		expect(document.body.textContent).toContain(portfolio.home.title);

		for (const slide of portfolio.home.slides) {
			expect(screen.getByRole('button', { name: new RegExp(slide.title) })).toBeTruthy();
		}

		await fireEvent.click(screen.getByRole('button', { name: new RegExp(portfolio.home.slides[1].title) }));
		expect(document.body.textContent).toContain(portfolio.home.slides[1].description);
	});
});
