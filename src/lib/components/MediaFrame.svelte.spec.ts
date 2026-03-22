import { describe, expect, it } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import MediaFrame from './MediaFrame.svelte';

describe('MediaFrame.svelte', () => {
	it('shows a placeholder when no image source is provided', async () => {
		render(MediaFrame, {
			src: '',
			alt: 'Missing image',
			label: 'Tobi Wealth'
		});

		expect(screen.getByLabelText('Tobi Wealth placeholder')).toBeTruthy();
	});
});
