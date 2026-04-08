import { describe, expect, it } from 'vitest';
import { fireEvent, render, screen } from '@testing-library/svelte';
import portfolio from '$lib/data/portfolio';
import ContactSection from './ContactSection.svelte';

describe('ContactSection.svelte', () => {
	it('updates the draft mailto link when form inputs change', async () => {
		render(ContactSection, {
			contact: portfolio.contact
		});

		// Query inputs by their placeholder text
		const nameInput = screen.getByPlaceholderText('ENTER_DESIGNATION...');
		const emailInput = screen.getByPlaceholderText('ENTER_EMAIL...');
		const messageInput = screen.getByPlaceholderText('INPUT_DATA_PACKAGE...');

		await fireEvent.input(nameInput, { target: { value: 'Jane & Co' } });
		await fireEvent.input(emailInput, { target: { value: 'jane@example.com' } });
		await fireEvent.input(messageInput, { target: { value: 'Looking forward to chatting soon.' } });

		const draftLink = screen.getByTestId('contact-draft-link');

		expect(draftLink.getAttribute('href')).toContain('Jane%20%26%20Co');
		expect(draftLink.getAttribute('href')).toContain('Looking%20forward%20to%20chatting%20soon.');
	});
});
