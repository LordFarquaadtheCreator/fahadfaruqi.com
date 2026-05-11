import rawPortfolio from './portfolio.json';

export interface PortfolioLink {
	label: string;
	href: string;
}

export interface NavigationData {
	links: PortfolioLink[];
}

export interface ProfileData {
	name: string;
	title: string;
	location: string;
}

export interface HomeSlide {
	label: string;
	title: string;
	description: string;
	href: string;
	meta: string;
}

export interface HomeData {
	eyebrow: string;
	title: string;
	description: string;
	slides: HomeSlide[];
}

export interface ExperienceItem {
	company: string;
	companyUrl?: string;
	logo?: string;
	role: string;
	location: string;
	start: string;
	end: string;
	summary?: string;
	highlights: string[];
	tools: string[];
	skills: string[];
}

export interface ResumeData {
	downloadUrl?: string;
	title: string;
	description: string;
	experiences: ExperienceItem[];
}

export interface BlogData {
	title: string;
	description: string;
}

export interface ContactData {
	email: string;
}

export interface PortfolioData {
	navigation: NavigationData;
	profile: ProfileData;
	home: HomeData;
	resume: ResumeData;
	blog: BlogData;
	contact: ContactData;
}

type JsonRecord = Record<string, unknown>;

function asRecord(value: unknown, path: string): JsonRecord {
	if (!value || typeof value !== 'object' || Array.isArray(value)) {
		throw new Error(`Invalid portfolio data at ${path}: expected an object.`);
	}

	return value as JsonRecord;
}

function readString(value: unknown, path: string): string {
	if (typeof value !== 'string' || value.trim().length === 0) {
		throw new Error(`Invalid portfolio data at ${path}: expected a non-empty string.`);
	}

	return value;
}

function readOptionalString(value: unknown, path: string): string | undefined {
	if (value === undefined || value === null || value === '') {
		return undefined;
	}

	return readString(value, path);
}

function readArray(value: unknown, path: string): unknown[] {
	if (!Array.isArray(value)) {
		throw new Error(`Invalid portfolio data at ${path}: expected an array.`);
	}

	return value;
}

function readStringArray(value: unknown, path: string): string[] {
	return readArray(value, path).map((item, index) => readString(item, `${path}[${index}]`));
}

function readLinks(value: unknown, path: string): PortfolioLink[] {
	return readArray(value, path).map((entry, index) => {
		const record = asRecord(entry, `${path}[${index}]`);

		return {
			label: readString(record.label, `${path}[${index}].label`),
			href: readString(record.href, `${path}[${index}].href`)
		};
	});
}

function validatePortfolio(data: unknown): PortfolioData {
	const root = asRecord(data, 'portfolio');

	const navigationRecord = asRecord(root.navigation, 'navigation');
	const profileRecord = asRecord(root.profile, 'profile');
	const homeRecord = asRecord(root.home, 'home');
	const resumeRecord = asRecord(root.resume, 'resume');
	const blogRecord = asRecord(root.blog, 'blog');
	const contactRecord = asRecord(root.contact, 'contact');

	return {
		navigation: {
			links: readLinks(navigationRecord.links, 'navigation.links')
		},
		profile: {
			name: readString(profileRecord.name, 'profile.name'),
			title: readString(profileRecord.title, 'profile.title'),
			location: readString(profileRecord.location, 'profile.location')
		},
		home: {
			eyebrow: readString(homeRecord.eyebrow, 'home.eyebrow'),
			title: readString(homeRecord.title, 'home.title'),
			description: readString(homeRecord.description, 'home.description'),
			slides: readArray(homeRecord.slides, 'home.slides').map((entry, index) => {
				const record = asRecord(entry, `home.slides[${index}]`);

				return {
					label: readString(record.label, `home.slides[${index}].label`),
					title: readString(record.title, `home.slides[${index}].title`),
					description: readString(record.description, `home.slides[${index}].description`),
					href: readString(record.href, `home.slides[${index}].href`),
					meta: readString(record.meta, `home.slides[${index}].meta`)
				};
			})
		},
		resume: {
			downloadUrl: readOptionalString(resumeRecord.downloadUrl, 'resume.downloadUrl'),
			title: readString(resumeRecord.title, 'resume.title'),
			description: readString(resumeRecord.description, 'resume.description'),
			experiences: readArray(resumeRecord.experiences, 'resume.experiences').map((entry, index) => {
				const record = asRecord(entry, `resume.experiences[${index}]`);

				return {
					company: readString(record.company, `resume.experiences[${index}].company`),
					companyUrl: readOptionalString(record.companyUrl, `resume.experiences[${index}].companyUrl`),
					logo: readOptionalString(record.logo, `resume.experiences[${index}].logo`),
					role: readString(record.role, `resume.experiences[${index}].role`),
					location: readString(record.location, `resume.experiences[${index}].location`),
					start: readString(record.start, `resume.experiences[${index}].start`),
					end: readString(record.end, `resume.experiences[${index}].end`),
					summary: readOptionalString(record.summary, `resume.experiences[${index}].summary`),
					highlights: readStringArray(record.highlights, `resume.experiences[${index}].highlights`),
					tools: readStringArray(record.tools, `resume.experiences[${index}].tools`),
					skills: readStringArray(record.skills, `resume.experiences[${index}].skills`)
				};
			})
		},
		blog: {
			title: readString(blogRecord.title, 'blog.title'),
			description: readString(blogRecord.description, 'blog.description')
		},
		contact: {
			email: readString(contactRecord.email, 'contact.email')
		}
	};
}

const portfolio = validatePortfolio(rawPortfolio);

export default portfolio;
