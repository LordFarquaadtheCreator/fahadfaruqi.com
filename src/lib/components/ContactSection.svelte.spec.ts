import { describe, expect, it } from 'vitest';
import { fireEvent, render, screen } from '@testing-library/svelte';
import portfolio from '$lib/data/portfolio';
import ContactSection from './ContactSection.svelte';

describe('ContactSection.svelte', () => {
	it('updates the draft mailto link and keeps a direct email link available', async () => {
		render(ContactSection, {
			contact: portfolio.contact
		});

		await fireEvent.input(screen.getByLabelText('Your name'), {
			target: { value: 'Jane & Co' }
		});
		await fireEvent.input(screen.getByLabelText('Your email'), {
			target: { value: 'jane@example.com' }
		});
		await fireEvent.input(screen.getByLabelText('Your message'), {
			target: { value: 'Looking forward to chatting soon.' }
		});

		const draftLink = screen.getByTestId('contact-draft-link');
		const emailLink = screen.getByRole('link', { name: 'Email' });

		expect(draftLink.getAttribute('href')).toContain('Jane%20%26%20Co');
		expect(draftLink.getAttribute('href')).toContain('Looking%20forward%20to%20chatting%20soon.');
		expect(emailLink.getAttribute('href')).toBe(`mailto:${portfolio.contact.email}`);
	});
});
