interface MailtoInput {
	email: string;
	subjectPrefix: string;
	name: string;
	fromEmail: string;
	message: string;
}

export function buildMailtoHref({
	email,
	subjectPrefix,
	name,
	fromEmail,
	message
}: MailtoInput): string {
	const safeName = name.trim() || 'Portfolio visitor';
	const safeFromEmail = fromEmail.trim() || 'No email provided';
	const safeMessage = message.trim() || 'Hi Fahad,\n\nI wanted to reach out through your portfolio.';
	const subject = `${subjectPrefix} - ${safeName}`;
	const body = [
		'Hi Fahad,',
		'',
		`Name: ${safeName}`,
		`Email: ${safeFromEmail}`,
		'',
		'Message:',
		safeMessage
	].join('\n');

	return `mailto:${email}?subject=${encodeURIComponent(subject)}&body=${encodeURIComponent(body)}`;
}
