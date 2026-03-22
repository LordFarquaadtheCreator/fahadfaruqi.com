import rawPortfolio from './portfolio.json';

export interface PortfolioLink {
	label: string;
	href: string;
}

export interface Metric {
	value: string;
	label: string;
	detail: string;
}

export interface ProfileData {
	name: string;
	title: string;
	location: string;
	heroIntro: string;
	supportingIntro: string;
	portraitImage?: string;
	portraitAlt: string;
	socialLinks: PortfolioLink[];
	metrics: Metric[];
}

export interface EducationItem {
	institution: string;
	credential: string;
	dates: string;
	location: string;
}

export interface HighlightItem {
	eyebrow: string;
	title: string;
	description: string;
	href?: string;
	linkLabel?: string;
}

export interface AboutData {
	paragraphs: string[];
	education: EducationItem[];
	highlights: HighlightItem[];
	photoCaption: string;
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
	experiences: ExperienceItem[];
}

export interface ContactData {
	email: string;
	subjectPrefix: string;
	availability: string;
	links: PortfolioLink[];
}

export interface PortfolioData {
	profile: ProfileData;
	about: AboutData;
	resume: ResumeData;
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

	const profileRecord = asRecord(root.profile, 'profile');
	const aboutRecord = asRecord(root.about, 'about');
	const resumeRecord = asRecord(root.resume, 'resume');
	const contactRecord = asRecord(root.contact, 'contact');

	return {
		profile: {
			name: readString(profileRecord.name, 'profile.name'),
			title: readString(profileRecord.title, 'profile.title'),
			location: readString(profileRecord.location, 'profile.location'),
			heroIntro: readString(profileRecord.heroIntro, 'profile.heroIntro'),
			supportingIntro: readString(profileRecord.supportingIntro, 'profile.supportingIntro'),
			portraitImage: readOptionalString(profileRecord.portraitImage, 'profile.portraitImage'),
			portraitAlt: readString(profileRecord.portraitAlt, 'profile.portraitAlt'),
			socialLinks: readLinks(profileRecord.socialLinks, 'profile.socialLinks'),
			metrics: readArray(profileRecord.metrics, 'profile.metrics').map((entry, index) => {
				const record = asRecord(entry, `profile.metrics[${index}]`);

				return {
					value: readString(record.value, `profile.metrics[${index}].value`),
					label: readString(record.label, `profile.metrics[${index}].label`),
					detail: readString(record.detail, `profile.metrics[${index}].detail`)
				};
			})
		},
		about: {
			paragraphs: readStringArray(aboutRecord.paragraphs, 'about.paragraphs'),
			education: readArray(aboutRecord.education, 'about.education').map((entry, index) => {
				const record = asRecord(entry, `about.education[${index}]`);

				return {
					institution: readString(record.institution, `about.education[${index}].institution`),
					credential: readString(record.credential, `about.education[${index}].credential`),
					dates: readString(record.dates, `about.education[${index}].dates`),
					location: readString(record.location, `about.education[${index}].location`)
				};
			}),
			highlights: readArray(aboutRecord.highlights, 'about.highlights').map((entry, index) => {
				const record = asRecord(entry, `about.highlights[${index}]`);

				return {
					eyebrow: readString(record.eyebrow, `about.highlights[${index}].eyebrow`),
					title: readString(record.title, `about.highlights[${index}].title`),
					description: readString(record.description, `about.highlights[${index}].description`),
					href: readOptionalString(record.href, `about.highlights[${index}].href`),
					linkLabel: readOptionalString(record.linkLabel, `about.highlights[${index}].linkLabel`)
				};
			}),
			photoCaption: readString(aboutRecord.photoCaption, 'about.photoCaption')
		},
		resume: {
			downloadUrl: readOptionalString(resumeRecord.downloadUrl, 'resume.downloadUrl'),
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
		contact: {
			email: readString(contactRecord.email, 'contact.email'),
			subjectPrefix: readString(contactRecord.subjectPrefix, 'contact.subjectPrefix'),
			availability: readString(contactRecord.availability, 'contact.availability'),
			links: readLinks(contactRecord.links, 'contact.links')
		}
	};
}

const portfolio = validatePortfolio(rawPortfolio);

export default portfolio;
