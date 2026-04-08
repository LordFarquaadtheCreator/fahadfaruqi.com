import { describe, expect, it } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import portfolio from '$lib/data/portfolio';
import Page from './+page.svelte';

describe('portfolio page', () => {
	it('renders every experience card from the JSON data source', async () => {
		render(Page);

		// Name is now displayed as FAHAD_FARUQI.SYS (uppercase with underscore)
		const transformedName = portfolio.profile.name.toUpperCase().replace(/ /g, '_');
		expect(document.body.textContent).toContain(transformedName);
		expect(screen.getAllByTestId('experience-card').length).toBe(portfolio.resume.experiences.length);
		expect(document.body.textContent).toContain(portfolio.resume.experiences[0]?.company);
		expect(document.body.textContent).toContain(
			portfolio.resume.experiences[portfolio.resume.experiences.length - 1]?.company
		);
	});
});
