export interface BlogPost {
	slug: string;
	title: string;
	category: string;
	date: string;
	excerpt: string;
	image?: string;
	imageAlt: string;
	pullQuote: string;
	markdown: string;
	html: string;
}

const modules = import.meta.glob('/blogs/*.md', {
	eager: true,
	query: '?raw',
	import: 'default'
}) as Record<string, string>;

function escapeHtml(value: string): string {
	return value
		.replaceAll('&', '&amp;')
		.replaceAll('<', '&lt;')
		.replaceAll('>', '&gt;')
		.replaceAll('"', '&quot;')
		.replaceAll("'", '&#39;');
}

function slugFromPath(path: string): string {
	return path.split('/').pop()?.replace(/\.md$/, '') ?? '';
}

function readFrontmatter(source: string, slug: string): Omit<BlogPost, 'slug' | 'html'> {
	const match = /^---\n([\s\S]*?)\n---\n([\s\S]*)$/u.exec(source);

	if (!match) {
		throw new Error(`Blog post ${slug} is missing frontmatter.`);
	}

	const fields = new Map<string, string>();

	for (const line of match[1].split('\n')) {
		const separatorIndex = line.indexOf(':');
		if (separatorIndex === -1) continue;

		fields.set(line.slice(0, separatorIndex).trim(), line.slice(separatorIndex + 1).trim());
	}

	function required(key: string): string {
		const value = fields.get(key);
		if (!value) throw new Error(`Blog post ${slug} is missing ${key}.`);
		return value;
	}

	function optional(key: string): string | undefined {
		return fields.get(key);
	}

	return {
		title: required('title'),
		category: required('category'),
		date: required('date'),
		excerpt: required('excerpt'),
		image: optional('image'),
		imageAlt: required('imageAlt'),
		pullQuote: required('pullQuote'),
		markdown: match[2].trim()
	};
}

function markdownToHtml(markdown: string): string {
	const lines = markdown.split('\n');
	const html: string[] = [];
	let paragraph: string[] = [];
	let code: string[] | null = null;

	function flushParagraph() {
		if (paragraph.length === 0) return;
		html.push(`<p>${escapeHtml(paragraph.join(' '))}</p>`);
		paragraph = [];
	}

	for (const line of lines) {
		if (line.startsWith('```')) {
			if (code) {
				html.push(`<pre><code>${escapeHtml(code.join('\n'))}</code></pre>`);
				code = null;
			} else {
				flushParagraph();
				code = [];
			}
			continue;
		}

		if (code) {
			code.push(line);
			continue;
		}

		if (line.trim().length === 0) {
			flushParagraph();
			continue;
		}

		if (line.startsWith('## ')) {
			flushParagraph();
			html.push(`<h2>${escapeHtml(line.slice(3).trim())}</h2>`);
			continue;
		}

		if (line.startsWith('> ')) {
			flushParagraph();
			html.push(`<blockquote>${escapeHtml(line.slice(2).trim())}</blockquote>`);
			continue;
		}

		paragraph.push(line.trim());
	}

	flushParagraph();

	if (code) {
		html.push(`<pre><code>${escapeHtml(code.join('\n'))}</code></pre>`);
	}

	return html.join('\n');
}

export const blogPosts: BlogPost[] = Object.entries(modules)
	.map(([path, source]) => {
		const slug = slugFromPath(path);
		const post = readFrontmatter(source, slug);

		return {
			slug,
			...post,
			html: markdownToHtml(post.markdown)
		};
	})
	.sort((a, b) => dateValue(b.date) - dateValue(a.date));

function dateValue(value: string): number {
	const [month = '', year = '0'] = value.split(/\s+/u);
	const monthIndex = [
		'January',
		'February',
		'March',
		'April',
		'May',
		'June',
		'July',
		'August',
		'September',
		'October',
		'November',
		'December'
	].indexOf(month);

	return Number(year) * 12 + monthIndex;
}

export function getBlogPost(slug: string): BlogPost | undefined {
	return blogPosts.find((post) => post.slug === slug);
}
