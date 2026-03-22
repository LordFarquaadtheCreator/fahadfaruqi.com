import { describe, expect, it } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import portfolio from '$lib/data/portfolio';
import Page from './+page.svelte';

describe('portfolio page', () => {
	it('renders every experience card from the JSON data source', async () => {
		render(Page);

		expect(screen.getByRole('heading', { level: 1, name: portfolio.profile.name })).toBeTruthy();
		expect(screen.getAllByTestId('experience-card').length).toBe(portfolio.resume.experiences.length);
		expect(document.body.textContent).toContain(portfolio.resume.experiences[0]?.company);
		expect(document.body.textContent).toContain(
			portfolio.resume.experiences[portfolio.resume.experiences.length - 1]?.company
		);
	});
});
