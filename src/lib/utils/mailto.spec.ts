import { describe, expect, it } from 'vitest';
import { buildMailtoHref } from './mailto';

describe('buildMailtoHref', () => {
	it('encodes the subject and body content for a mail client draft', () => {
		const href = buildMailtoHref({
			email: 'fahadfaruqi1@gmail.com',
			subjectPrefix: 'Portfolio inquiry',
			name: 'Jane & Co',
			fromEmail: 'jane@example.com',
			message: 'Hello there!'
		});

		expect(href.startsWith('mailto:fahadfaruqi1@gmail.com')).toBe(true);
		expect(href).toContain('Jane%20%26%20Co');
		expect(href).toContain('jane%40example.com');
		expect(href).toContain('Hello%20there!');
	});
});
